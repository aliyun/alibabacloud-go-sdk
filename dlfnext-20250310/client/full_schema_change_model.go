// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFullSchemaChange interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *FullSchemaChange
	GetAction() *string
	SetComment(v string) *FullSchemaChange
	GetComment() *string
	SetDataType(v *FullDataType) *FullSchemaChange
	GetDataType() *FullDataType
	SetFieldNames(v []*string) *FullSchemaChange
	GetFieldNames() []*string
	SetKeepNullability(v bool) *FullSchemaChange
	GetKeepNullability() *bool
	SetKey(v string) *FullSchemaChange
	GetKey() *string
	SetMove(v *Move) *FullSchemaChange
	GetMove() *Move
	SetNewComment(v string) *FullSchemaChange
	GetNewComment() *string
	SetNewDataType(v *FullDataType) *FullSchemaChange
	GetNewDataType() *FullDataType
	SetNewName(v string) *FullSchemaChange
	GetNewName() *string
	SetNewNullability(v bool) *FullSchemaChange
	GetNewNullability() *bool
	SetValue(v string) *FullSchemaChange
	GetValue() *string
}

type FullSchemaChange struct {
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// required in UpdateComment/AddColumn
	Comment  *string       `json:"comment,omitempty" xml:"comment,omitempty"`
	DataType *FullDataType `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// required in AddColumn/RenameColumn/DropColumn/UpdateColumnComment/UpdateColumnType/UpdateColumnNullability
	FieldNames []*string `json:"fieldNames,omitempty" xml:"fieldNames,omitempty" type:"Repeated"`
	// required in UpdateColumnType
	KeepNullability *bool `json:"keepNullability,omitempty" xml:"keepNullability,omitempty"`
	// required in SetOption/RemoveOption
	Key  *string `json:"key,omitempty" xml:"key,omitempty"`
	Move *Move   `json:"move,omitempty" xml:"move,omitempty"`
	// required in UpdateColumnComment
	NewComment  *string       `json:"newComment,omitempty" xml:"newComment,omitempty"`
	NewDataType *FullDataType `json:"newDataType,omitempty" xml:"newDataType,omitempty"`
	// required in RenameColumn
	NewName *string `json:"newName,omitempty" xml:"newName,omitempty"`
	// required in UpdateColumnNullability
	NewNullability *bool `json:"newNullability,omitempty" xml:"newNullability,omitempty"`
	// required in SetOption
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s FullSchemaChange) String() string {
	return dara.Prettify(s)
}

func (s FullSchemaChange) GoString() string {
	return s.String()
}

func (s *FullSchemaChange) GetAction() *string {
	return s.Action
}

func (s *FullSchemaChange) GetComment() *string {
	return s.Comment
}

func (s *FullSchemaChange) GetDataType() *FullDataType {
	return s.DataType
}

func (s *FullSchemaChange) GetFieldNames() []*string {
	return s.FieldNames
}

func (s *FullSchemaChange) GetKeepNullability() *bool {
	return s.KeepNullability
}

func (s *FullSchemaChange) GetKey() *string {
	return s.Key
}

func (s *FullSchemaChange) GetMove() *Move {
	return s.Move
}

func (s *FullSchemaChange) GetNewComment() *string {
	return s.NewComment
}

func (s *FullSchemaChange) GetNewDataType() *FullDataType {
	return s.NewDataType
}

func (s *FullSchemaChange) GetNewName() *string {
	return s.NewName
}

func (s *FullSchemaChange) GetNewNullability() *bool {
	return s.NewNullability
}

func (s *FullSchemaChange) GetValue() *string {
	return s.Value
}

func (s *FullSchemaChange) SetAction(v string) *FullSchemaChange {
	s.Action = &v
	return s
}

func (s *FullSchemaChange) SetComment(v string) *FullSchemaChange {
	s.Comment = &v
	return s
}

func (s *FullSchemaChange) SetDataType(v *FullDataType) *FullSchemaChange {
	s.DataType = v
	return s
}

func (s *FullSchemaChange) SetFieldNames(v []*string) *FullSchemaChange {
	s.FieldNames = v
	return s
}

func (s *FullSchemaChange) SetKeepNullability(v bool) *FullSchemaChange {
	s.KeepNullability = &v
	return s
}

func (s *FullSchemaChange) SetKey(v string) *FullSchemaChange {
	s.Key = &v
	return s
}

func (s *FullSchemaChange) SetMove(v *Move) *FullSchemaChange {
	s.Move = v
	return s
}

func (s *FullSchemaChange) SetNewComment(v string) *FullSchemaChange {
	s.NewComment = &v
	return s
}

func (s *FullSchemaChange) SetNewDataType(v *FullDataType) *FullSchemaChange {
	s.NewDataType = v
	return s
}

func (s *FullSchemaChange) SetNewName(v string) *FullSchemaChange {
	s.NewName = &v
	return s
}

func (s *FullSchemaChange) SetNewNullability(v bool) *FullSchemaChange {
	s.NewNullability = &v
	return s
}

func (s *FullSchemaChange) SetValue(v string) *FullSchemaChange {
	s.Value = &v
	return s
}

func (s *FullSchemaChange) Validate() error {
	return dara.Validate(s)
}
