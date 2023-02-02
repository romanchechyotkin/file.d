// Code generated by "./colgenerator"; DO NOT EDIT.

package clickhouse

import (
	"errors"

	"github.com/ClickHouse/ch-go/proto"
	insaneJSON "github.com/vitkovskii/insane-json"
)

type InsaneColInput interface {
	proto.ColInput
	Append(node *insaneJSON.StrictNode) error
	Reset()
}

var (
	ErrNodeIsNil = errors.New("node is nil, but column is not")
)

type ColStr struct {
	col      *proto.ColStr
	nullCol  proto.ColNullable[string]
	nullable bool
}

var _ InsaneColInput = (*ColStr)(nil)

func (t *ColStr) Append(node *insaneJSON.StrictNode) error {
	if node == nil {
		if !t.nullable {
			return ErrNodeIsNil
		}
		t.nullCol.Append(proto.Null[string]())
		return nil
	}

	val, err := node.AsString()
	if err != nil {
		return err
	}

	if t.nullable {
		t.nullCol.Append(proto.NewNullable(val))
		return nil
	}
	t.col.Append(val)

	return nil
}

func (t *ColStr) Reset() {
	t.col.Reset()
	t.nullCol.Reset()
}

func (t *ColStr) SetNullable() {
	t.nullable = true
}

func (t *ColStr) Type() proto.ColumnType {
	if t.nullable {
		return t.col.Type()
	}
	return t.nullCol.Type()
}

func (t *ColStr) Rows() int {
	if t.nullable {
		return t.col.Rows()
	}
	return t.nullCol.Rows()
}

func (t *ColStr) EncodeColumn(buffer *proto.Buffer) {
	if t.nullable {
		t.col.EncodeColumn(buffer)
	}
	t.nullCol.EncodeColumn(buffer)
}

type ColEnum struct {
	col *proto.ColEnum
}

var _ InsaneColInput = (*ColEnum)(nil)

func (t *ColEnum) Append(node *insaneJSON.StrictNode) error {
	if node == nil {
		return ErrNodeIsNil
	}

	val, err := node.AsString()
	if err != nil {
		return err
	}

	t.col.Append(val)

	return nil
}

func (t *ColEnum) Reset() {
	t.col.Reset()
}

func (t *ColEnum) SetNullable() {
}

func (t *ColEnum) Type() proto.ColumnType {
	return t.col.Type()
}

func (t *ColEnum) Rows() int {
	return t.col.Rows()
}

func (t *ColEnum) EncodeColumn(buffer *proto.Buffer) {
	t.col.EncodeColumn(buffer)
}

type ColInt8 struct {
	col      *proto.ColInt8
	nullCol  proto.ColNullable[int8]
	nullable bool
}

var _ InsaneColInput = (*ColInt8)(nil)

func (t *ColInt8) Append(node *insaneJSON.StrictNode) error {
	if node == nil {
		if !t.nullable {
			return ErrNodeIsNil
		}
		t.nullCol.Append(proto.Null[int8]())
		return nil
	}

	v, err := node.AsInt()
	if err != nil {
		return err
	}
	val := int8(v)
	if t.nullable {
		t.nullCol.Append(proto.NewNullable(val))
		return nil
	}
	t.col.Append(val)

	return nil
}

func (t *ColInt8) Reset() {
	t.col.Reset()
	t.nullCol.Reset()
}

func (t *ColInt8) SetNullable() {
	t.nullable = true
}

func (t *ColInt8) Type() proto.ColumnType {
	if t.nullable {
		return t.col.Type()
	}
	return t.nullCol.Type()
}

func (t *ColInt8) Rows() int {
	if t.nullable {
		return t.col.Rows()
	}
	return t.nullCol.Rows()
}

func (t *ColInt8) EncodeColumn(buffer *proto.Buffer) {
	if t.nullable {
		t.col.EncodeColumn(buffer)
	}
	t.nullCol.EncodeColumn(buffer)
}

type ColInt16 struct {
	col      *proto.ColInt16
	nullCol  proto.ColNullable[int16]
	nullable bool
}

var _ InsaneColInput = (*ColInt16)(nil)

func (t *ColInt16) Append(node *insaneJSON.StrictNode) error {
	if node == nil {
		if !t.nullable {
			return ErrNodeIsNil
		}
		t.nullCol.Append(proto.Null[int16]())
		return nil
	}

	v, err := node.AsInt()
	if err != nil {
		return err
	}
	val := int16(v)
	if t.nullable {
		t.nullCol.Append(proto.NewNullable(val))
		return nil
	}
	t.col.Append(val)

	return nil
}

func (t *ColInt16) Reset() {
	t.col.Reset()
	t.nullCol.Reset()
}

func (t *ColInt16) SetNullable() {
	t.nullable = true
}

func (t *ColInt16) Type() proto.ColumnType {
	if t.nullable {
		return t.col.Type()
	}
	return t.nullCol.Type()
}

func (t *ColInt16) Rows() int {
	if t.nullable {
		return t.col.Rows()
	}
	return t.nullCol.Rows()
}

func (t *ColInt16) EncodeColumn(buffer *proto.Buffer) {
	if t.nullable {
		t.col.EncodeColumn(buffer)
	}
	t.nullCol.EncodeColumn(buffer)
}

type ColInt32 struct {
	col      *proto.ColInt32
	nullCol  proto.ColNullable[int32]
	nullable bool
}

var _ InsaneColInput = (*ColInt32)(nil)

func (t *ColInt32) Append(node *insaneJSON.StrictNode) error {
	if node == nil {
		if !t.nullable {
			return ErrNodeIsNil
		}
		t.nullCol.Append(proto.Null[int32]())
		return nil
	}

	v, err := node.AsInt()
	if err != nil {
		return err
	}
	val := int32(v)
	if t.nullable {
		t.nullCol.Append(proto.NewNullable(val))
		return nil
	}
	t.col.Append(val)

	return nil
}

func (t *ColInt32) Reset() {
	t.col.Reset()
	t.nullCol.Reset()
}

func (t *ColInt32) SetNullable() {
	t.nullable = true
}

func (t *ColInt32) Type() proto.ColumnType {
	if t.nullable {
		return t.col.Type()
	}
	return t.nullCol.Type()
}

func (t *ColInt32) Rows() int {
	if t.nullable {
		return t.col.Rows()
	}
	return t.nullCol.Rows()
}

func (t *ColInt32) EncodeColumn(buffer *proto.Buffer) {
	if t.nullable {
		t.col.EncodeColumn(buffer)
	}
	t.nullCol.EncodeColumn(buffer)
}

type ColInt64 struct {
	col      *proto.ColInt64
	nullCol  proto.ColNullable[int64]
	nullable bool
}

var _ InsaneColInput = (*ColInt64)(nil)

func (t *ColInt64) Append(node *insaneJSON.StrictNode) error {
	if node == nil {
		if !t.nullable {
			return ErrNodeIsNil
		}
		t.nullCol.Append(proto.Null[int64]())
		return nil
	}

	v, err := node.AsInt()
	if err != nil {
		return err
	}
	val := int64(v)
	if t.nullable {
		t.nullCol.Append(proto.NewNullable(val))
		return nil
	}
	t.col.Append(val)

	return nil
}

func (t *ColInt64) Reset() {
	t.col.Reset()
	t.nullCol.Reset()
}

func (t *ColInt64) SetNullable() {
	t.nullable = true
}

func (t *ColInt64) Type() proto.ColumnType {
	if t.nullable {
		return t.col.Type()
	}
	return t.nullCol.Type()
}

func (t *ColInt64) Rows() int {
	if t.nullable {
		return t.col.Rows()
	}
	return t.nullCol.Rows()
}

func (t *ColInt64) EncodeColumn(buffer *proto.Buffer) {
	if t.nullable {
		t.col.EncodeColumn(buffer)
	}
	t.nullCol.EncodeColumn(buffer)
}

type ColInt128 struct {
	col      *proto.ColInt128
	nullCol  proto.ColNullable[proto.Int128]
	nullable bool
}

var _ InsaneColInput = (*ColInt128)(nil)

func (t *ColInt128) Append(node *insaneJSON.StrictNode) error {
	if node == nil {
		if !t.nullable {
			return ErrNodeIsNil
		}
		t.nullCol.Append(proto.Null[proto.Int128]())
		return nil
	}

	v, err := node.AsInt()
	if err != nil {
		return err
	}
	val := proto.Int128FromInt(v)
	if t.nullable {
		t.nullCol.Append(proto.NewNullable(val))
		return nil
	}
	t.col.Append(val)

	return nil
}

func (t *ColInt128) Reset() {
	t.col.Reset()
	t.nullCol.Reset()
}

func (t *ColInt128) SetNullable() {
	t.nullable = true
}

func (t *ColInt128) Type() proto.ColumnType {
	if t.nullable {
		return t.col.Type()
	}
	return t.nullCol.Type()
}

func (t *ColInt128) Rows() int {
	if t.nullable {
		return t.col.Rows()
	}
	return t.nullCol.Rows()
}

func (t *ColInt128) EncodeColumn(buffer *proto.Buffer) {
	if t.nullable {
		t.col.EncodeColumn(buffer)
	}
	t.nullCol.EncodeColumn(buffer)
}

type ColInt256 struct {
	col      *proto.ColInt256
	nullCol  proto.ColNullable[proto.Int256]
	nullable bool
}

var _ InsaneColInput = (*ColInt256)(nil)

func (t *ColInt256) Append(node *insaneJSON.StrictNode) error {
	if node == nil {
		if !t.nullable {
			return ErrNodeIsNil
		}
		t.nullCol.Append(proto.Null[proto.Int256]())
		return nil
	}

	v, err := node.AsInt()
	if err != nil {
		return err
	}
	val := proto.Int256FromInt(v)
	if t.nullable {
		t.nullCol.Append(proto.NewNullable(val))
		return nil
	}
	t.col.Append(val)

	return nil
}

func (t *ColInt256) Reset() {
	t.col.Reset()
	t.nullCol.Reset()
}

func (t *ColInt256) SetNullable() {
	t.nullable = true
}

func (t *ColInt256) Type() proto.ColumnType {
	if t.nullable {
		return t.col.Type()
	}
	return t.nullCol.Type()
}

func (t *ColInt256) Rows() int {
	if t.nullable {
		return t.col.Rows()
	}
	return t.nullCol.Rows()
}

func (t *ColInt256) EncodeColumn(buffer *proto.Buffer) {
	if t.nullable {
		t.col.EncodeColumn(buffer)
	}
	t.nullCol.EncodeColumn(buffer)
}

type ColUInt8 struct {
	col      *proto.ColUInt8
	nullCol  proto.ColNullable[uint8]
	nullable bool
}

var _ InsaneColInput = (*ColUInt8)(nil)

func (t *ColUInt8) Append(node *insaneJSON.StrictNode) error {
	if node == nil {
		if !t.nullable {
			return ErrNodeIsNil
		}
		t.nullCol.Append(proto.Null[uint8]())
		return nil
	}

	v, err := node.AsInt()
	if err != nil {
		return err
	}
	val := uint8(v)
	if t.nullable {
		t.nullCol.Append(proto.NewNullable(val))
		return nil
	}
	t.col.Append(val)

	return nil
}

func (t *ColUInt8) Reset() {
	t.col.Reset()
	t.nullCol.Reset()
}

func (t *ColUInt8) SetNullable() {
	t.nullable = true
}

func (t *ColUInt8) Type() proto.ColumnType {
	if t.nullable {
		return t.col.Type()
	}
	return t.nullCol.Type()
}

func (t *ColUInt8) Rows() int {
	if t.nullable {
		return t.col.Rows()
	}
	return t.nullCol.Rows()
}

func (t *ColUInt8) EncodeColumn(buffer *proto.Buffer) {
	if t.nullable {
		t.col.EncodeColumn(buffer)
	}
	t.nullCol.EncodeColumn(buffer)
}

type ColUInt16 struct {
	col      *proto.ColUInt16
	nullCol  proto.ColNullable[uint16]
	nullable bool
}

var _ InsaneColInput = (*ColUInt16)(nil)

func (t *ColUInt16) Append(node *insaneJSON.StrictNode) error {
	if node == nil {
		if !t.nullable {
			return ErrNodeIsNil
		}
		t.nullCol.Append(proto.Null[uint16]())
		return nil
	}

	v, err := node.AsInt()
	if err != nil {
		return err
	}
	val := uint16(v)
	if t.nullable {
		t.nullCol.Append(proto.NewNullable(val))
		return nil
	}
	t.col.Append(val)

	return nil
}

func (t *ColUInt16) Reset() {
	t.col.Reset()
	t.nullCol.Reset()
}

func (t *ColUInt16) SetNullable() {
	t.nullable = true
}

func (t *ColUInt16) Type() proto.ColumnType {
	if t.nullable {
		return t.col.Type()
	}
	return t.nullCol.Type()
}

func (t *ColUInt16) Rows() int {
	if t.nullable {
		return t.col.Rows()
	}
	return t.nullCol.Rows()
}

func (t *ColUInt16) EncodeColumn(buffer *proto.Buffer) {
	if t.nullable {
		t.col.EncodeColumn(buffer)
	}
	t.nullCol.EncodeColumn(buffer)
}

type ColUInt32 struct {
	col      *proto.ColUInt32
	nullCol  proto.ColNullable[uint32]
	nullable bool
}

var _ InsaneColInput = (*ColUInt32)(nil)

func (t *ColUInt32) Append(node *insaneJSON.StrictNode) error {
	if node == nil {
		if !t.nullable {
			return ErrNodeIsNil
		}
		t.nullCol.Append(proto.Null[uint32]())
		return nil
	}

	v, err := node.AsInt()
	if err != nil {
		return err
	}
	val := uint32(v)
	if t.nullable {
		t.nullCol.Append(proto.NewNullable(val))
		return nil
	}
	t.col.Append(val)

	return nil
}

func (t *ColUInt32) Reset() {
	t.col.Reset()
	t.nullCol.Reset()
}

func (t *ColUInt32) SetNullable() {
	t.nullable = true
}

func (t *ColUInt32) Type() proto.ColumnType {
	if t.nullable {
		return t.col.Type()
	}
	return t.nullCol.Type()
}

func (t *ColUInt32) Rows() int {
	if t.nullable {
		return t.col.Rows()
	}
	return t.nullCol.Rows()
}

func (t *ColUInt32) EncodeColumn(buffer *proto.Buffer) {
	if t.nullable {
		t.col.EncodeColumn(buffer)
	}
	t.nullCol.EncodeColumn(buffer)
}

type ColUInt64 struct {
	col      *proto.ColUInt64
	nullCol  proto.ColNullable[uint64]
	nullable bool
}

var _ InsaneColInput = (*ColUInt64)(nil)

func (t *ColUInt64) Append(node *insaneJSON.StrictNode) error {
	if node == nil {
		if !t.nullable {
			return ErrNodeIsNil
		}
		t.nullCol.Append(proto.Null[uint64]())
		return nil
	}

	v, err := node.AsInt()
	if err != nil {
		return err
	}
	val := uint64(v)
	if t.nullable {
		t.nullCol.Append(proto.NewNullable(val))
		return nil
	}
	t.col.Append(val)

	return nil
}

func (t *ColUInt64) Reset() {
	t.col.Reset()
	t.nullCol.Reset()
}

func (t *ColUInt64) SetNullable() {
	t.nullable = true
}

func (t *ColUInt64) Type() proto.ColumnType {
	if t.nullable {
		return t.col.Type()
	}
	return t.nullCol.Type()
}

func (t *ColUInt64) Rows() int {
	if t.nullable {
		return t.col.Rows()
	}
	return t.nullCol.Rows()
}

func (t *ColUInt64) EncodeColumn(buffer *proto.Buffer) {
	if t.nullable {
		t.col.EncodeColumn(buffer)
	}
	t.nullCol.EncodeColumn(buffer)
}

type ColUInt128 struct {
	col      *proto.ColUInt128
	nullCol  proto.ColNullable[proto.UInt128]
	nullable bool
}

var _ InsaneColInput = (*ColUInt128)(nil)

func (t *ColUInt128) Append(node *insaneJSON.StrictNode) error {
	if node == nil {
		if !t.nullable {
			return ErrNodeIsNil
		}
		t.nullCol.Append(proto.Null[proto.UInt128]())
		return nil
	}

	v, err := node.AsInt()
	if err != nil {
		return err
	}
	val := proto.UInt128FromInt(v)
	if t.nullable {
		t.nullCol.Append(proto.NewNullable(val))
		return nil
	}
	t.col.Append(val)

	return nil
}

func (t *ColUInt128) Reset() {
	t.col.Reset()
	t.nullCol.Reset()
}

func (t *ColUInt128) SetNullable() {
	t.nullable = true
}

func (t *ColUInt128) Type() proto.ColumnType {
	if t.nullable {
		return t.col.Type()
	}
	return t.nullCol.Type()
}

func (t *ColUInt128) Rows() int {
	if t.nullable {
		return t.col.Rows()
	}
	return t.nullCol.Rows()
}

func (t *ColUInt128) EncodeColumn(buffer *proto.Buffer) {
	if t.nullable {
		t.col.EncodeColumn(buffer)
	}
	t.nullCol.EncodeColumn(buffer)
}

type ColUInt256 struct {
	col      *proto.ColUInt256
	nullCol  proto.ColNullable[proto.UInt256]
	nullable bool
}

var _ InsaneColInput = (*ColUInt256)(nil)

func (t *ColUInt256) Append(node *insaneJSON.StrictNode) error {
	if node == nil {
		if !t.nullable {
			return ErrNodeIsNil
		}
		t.nullCol.Append(proto.Null[proto.UInt256]())
		return nil
	}

	v, err := node.AsInt()
	if err != nil {
		return err
	}
	val := proto.UInt256FromInt(v)
	if t.nullable {
		t.nullCol.Append(proto.NewNullable(val))
		return nil
	}
	t.col.Append(val)

	return nil
}

func (t *ColUInt256) Reset() {
	t.col.Reset()
	t.nullCol.Reset()
}

func (t *ColUInt256) SetNullable() {
	t.nullable = true
}

func (t *ColUInt256) Type() proto.ColumnType {
	if t.nullable {
		return t.col.Type()
	}
	return t.nullCol.Type()
}

func (t *ColUInt256) Rows() int {
	if t.nullable {
		return t.col.Rows()
	}
	return t.nullCol.Rows()
}

func (t *ColUInt256) EncodeColumn(buffer *proto.Buffer) {
	if t.nullable {
		t.col.EncodeColumn(buffer)
	}
	t.nullCol.EncodeColumn(buffer)
}

type ColFloat32 struct {
	col      *proto.ColFloat32
	nullCol  proto.ColNullable[float32]
	nullable bool
}

var _ InsaneColInput = (*ColFloat32)(nil)

func (t *ColFloat32) Append(node *insaneJSON.StrictNode) error {
	if node == nil {
		if !t.nullable {
			return ErrNodeIsNil
		}
		t.nullCol.Append(proto.Null[float32]())
		return nil
	}

	v, err := node.AsInt()
	if err != nil {
		return err
	}
	val := float32(v)
	if t.nullable {
		t.nullCol.Append(proto.NewNullable(val))
		return nil
	}
	t.col.Append(val)

	return nil
}

func (t *ColFloat32) Reset() {
	t.col.Reset()
	t.nullCol.Reset()
}

func (t *ColFloat32) SetNullable() {
	t.nullable = true
}

func (t *ColFloat32) Type() proto.ColumnType {
	if t.nullable {
		return t.col.Type()
	}
	return t.nullCol.Type()
}

func (t *ColFloat32) Rows() int {
	if t.nullable {
		return t.col.Rows()
	}
	return t.nullCol.Rows()
}

func (t *ColFloat32) EncodeColumn(buffer *proto.Buffer) {
	if t.nullable {
		t.col.EncodeColumn(buffer)
	}
	t.nullCol.EncodeColumn(buffer)
}

type ColFloat64 struct {
	col      *proto.ColFloat64
	nullCol  proto.ColNullable[float64]
	nullable bool
}

var _ InsaneColInput = (*ColFloat64)(nil)

func (t *ColFloat64) Append(node *insaneJSON.StrictNode) error {
	if node == nil {
		if !t.nullable {
			return ErrNodeIsNil
		}
		t.nullCol.Append(proto.Null[float64]())
		return nil
	}

	v, err := node.AsInt()
	if err != nil {
		return err
	}
	val := float64(v)
	if t.nullable {
		t.nullCol.Append(proto.NewNullable(val))
		return nil
	}
	t.col.Append(val)

	return nil
}

func (t *ColFloat64) Reset() {
	t.col.Reset()
	t.nullCol.Reset()
}

func (t *ColFloat64) SetNullable() {
	t.nullable = true
}

func (t *ColFloat64) Type() proto.ColumnType {
	if t.nullable {
		return t.col.Type()
	}
	return t.nullCol.Type()
}

func (t *ColFloat64) Rows() int {
	if t.nullable {
		return t.col.Rows()
	}
	return t.nullCol.Rows()
}

func (t *ColFloat64) EncodeColumn(buffer *proto.Buffer) {
	if t.nullable {
		t.col.EncodeColumn(buffer)
	}
	t.nullCol.EncodeColumn(buffer)
}
