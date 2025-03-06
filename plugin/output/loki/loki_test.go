package loki

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPluginLabels(t *testing.T) {
	type testCase struct {
		name        string
		expectedLen int
		lables      []Label
	}

	tests := []testCase{
		{
			name:        "one label",
			expectedLen: 1,
			lables: []Label{
				{
					Label: "label1",
					Value: "value1",
				},
			},
		},
		{
			name:        "two labels",
			expectedLen: 2,
			lables: []Label{
				{
					Label: "label1",
					Value: "value1",
				},
				{
					Label: "label2",
					Value: "value2",
				},
			},
		},
		{
			name:        "two lables; repetetive label key",
			expectedLen: 1,
			lables: []Label{
				{
					Label: "label1",
					Value: "value1",
				},
				{
					Label: "label1",
					Value: "value2",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pl := &Plugin{
				config: &Config{
					Labels: tt.lables,
				},
			}

			resultLabelsMap := pl.labels()
			require.Len(t, resultLabelsMap, tt.expectedLen)
		})
	}
}

func TestPlugSetAuthHeaders(t *testing.T) {
	type testCase struct {
		name         string
		tenantID     string
		expectTenant bool

		username    string
		password    string
		expectBasic bool

		bearer       string
		expectBearer bool
	}

	tests := []testCase{
		{
			name:         "only tenant",
			tenantID:     "tenant",
			expectTenant: true,
		},
		{
			name:         "only tenant",
			tenantID:     "tenant",
			expectTenant: true,
			username:     "username",
			password:     "tenant",
			expectBasic:  true,
		},
		{
			name:         "bearer token",
			tenantID:     "tenant",
			expectTenant: true,
			expectBasic:  false,
			bearer:       "token",
			expectBearer: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pl := &Plugin{
				config: &Config{
					TenantID:     tt.tenantID,
					AuthUsername: tt.username,
					AuthPassword: tt.password,
					BearerToken:  tt.bearer,
				},
			}

			req, err := http.NewRequest(http.MethodPost, "url", nil)
			assert.NoError(t, err)

			pl.setAuthenticationHeaders(req)

			if tt.expectTenant {
				require.Equal(t, pl.config.TenantID, req.Header.Get("X-Scope-OrgID"))
			}
			if tt.expectBasic {
				require.Equal(t, fmt.Sprintf("Basic %s:%s", pl.config.AuthUsername, pl.config.AuthPassword), req.Header.Get("Authorization"))
			}
			if tt.expectBearer {
				require.Equal(t, fmt.Sprintf("Bearer %s", pl.config.BearerToken), req.Header.Get("Authorization"))
			}

		})
	}
}

func TestIsUnixNanoFormat(t *testing.T) {
	type args struct {
		name         string
		ts           string
		expectResult bool
	}

	testCases := []args{
		{
			name:         "valid",
			ts:           "1700000000000000000",
			expectResult: true,
		},
		{
			name:         "Too Large",
			ts:           "9999999999999999999", // 20 November 2286
			expectResult: false,
		},
		{
			name:         "Valid Timestamp",
			ts:           "1609459200123456789", // 1 January 2021
			expectResult: true,
		},
		{
			name:         "Invalid Non-Numeric",
			ts:           "hello123456789",
			expectResult: false,
		},
		{
			name:         "now",
			ts:           fmt.Sprintf("%d", time.Now().UnixNano()),
			expectResult: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			p := &Plugin{}

			result := p.isUnixNanoFormat(tt.ts)
			if result != tt.expectResult {
				t.Errorf("isUnixNano(%s) = %v, want %v", tt.ts, result, tt.expectResult)
			}
		})
	}
}
