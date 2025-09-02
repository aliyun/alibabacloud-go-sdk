// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTableAddColumnRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumn(v []*UpdateTableAddColumnRequestColumn) *UpdateTableAddColumnRequest
	GetColumn() []*UpdateTableAddColumnRequestColumn
	SetTableGuid(v string) *UpdateTableAddColumnRequest
	GetTableGuid() *string
}

type UpdateTableAddColumnRequest struct {
	// The fields.
	//
	// This parameter is required.
	Column []*UpdateTableAddColumnRequestColumn `json:"Column,omitempty" xml:"Column,omitempty" type:"Repeated"`
	// The globally unique identifier (GUID) of the MaxCompute table. Specify the GUID in the odps.projectName.tableName format.
	//
	// This parameter is required.
	//
	// example:
	//
	// odps.engine_name.table_name
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
}

func (s UpdateTableAddColumnRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableAddColumnRequest) GoString() string {
	return s.String()
}

func (s *UpdateTableAddColumnRequest) GetColumn() []*UpdateTableAddColumnRequestColumn {
	return s.Column
}

func (s *UpdateTableAddColumnRequest) GetTableGuid() *string {
	return s.TableGuid
}

func (s *UpdateTableAddColumnRequest) SetColumn(v []*UpdateTableAddColumnRequestColumn) *UpdateTableAddColumnRequest {
	s.Column = v
	return s
}

func (s *UpdateTableAddColumnRequest) SetTableGuid(v string) *UpdateTableAddColumnRequest {
	s.TableGuid = &v
	return s
}

func (s *UpdateTableAddColumnRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateTableAddColumnRequestColumn struct {
	// The name of the field.
	//
	// This parameter is required.
	//
	// example:
	//
	// abc
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The display name of the field.
	//
	// example:
	//
	// Chinese
	ColumnNameCn *string `json:"ColumnNameCn,omitempty" xml:"ColumnNameCn,omitempty"`
	// The type of the field. For more information, see MaxCompute field types.
	//
	// This parameter is required.
	//
	// example:
	//
	// string
	ColumnType *string `json:"ColumnType,omitempty" xml:"ColumnType,omitempty"`
	// The comment of the field.
	//
	// example:
	//
	// Remarks
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
}

func (s UpdateTableAddColumnRequestColumn) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableAddColumnRequestColumn) GoString() string {
	return s.String()
}

func (s *UpdateTableAddColumnRequestColumn) GetColumnName() *string {
	return s.ColumnName
}

func (s *UpdateTableAddColumnRequestColumn) GetColumnNameCn() *string {
	return s.ColumnNameCn
}

func (s *UpdateTableAddColumnRequestColumn) GetColumnType() *string {
	return s.ColumnType
}

func (s *UpdateTableAddColumnRequestColumn) GetComment() *string {
	return s.Comment
}

func (s *UpdateTableAddColumnRequestColumn) SetColumnName(v string) *UpdateTableAddColumnRequestColumn {
	s.ColumnName = &v
	return s
}

func (s *UpdateTableAddColumnRequestColumn) SetColumnNameCn(v string) *UpdateTableAddColumnRequestColumn {
	s.ColumnNameCn = &v
	return s
}

func (s *UpdateTableAddColumnRequestColumn) SetColumnType(v string) *UpdateTableAddColumnRequestColumn {
	s.ColumnType = &v
	return s
}

func (s *UpdateTableAddColumnRequestColumn) SetComment(v string) *UpdateTableAddColumnRequestColumn {
	s.Comment = &v
	return s
}

func (s *UpdateTableAddColumnRequestColumn) Validate() error {
	return dara.Validate(s)
}
