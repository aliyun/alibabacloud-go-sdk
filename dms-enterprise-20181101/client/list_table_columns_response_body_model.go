// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTableColumnsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetColumnList(v *ListTableColumnsResponseBodyColumnList) *ListTableColumnsResponseBody
	GetColumnList() *ListTableColumnsResponseBodyColumnList
	SetErrorCode(v string) *ListTableColumnsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListTableColumnsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListTableColumnsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTableColumnsResponseBody
	GetSuccess() *bool
}

type ListTableColumnsResponseBody struct {
	// The details about fields in the table.
	ColumnList *ListTableColumnsResponseBodyColumnList `json:"ColumnList,omitempty" xml:"ColumnList,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListTableColumnsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTableColumnsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTableColumnsResponseBody) GetColumnList() *ListTableColumnsResponseBodyColumnList {
	return s.ColumnList
}

func (s *ListTableColumnsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListTableColumnsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListTableColumnsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTableColumnsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTableColumnsResponseBody) SetColumnList(v *ListTableColumnsResponseBodyColumnList) *ListTableColumnsResponseBody {
	s.ColumnList = v
	return s
}

func (s *ListTableColumnsResponseBody) SetErrorCode(v string) *ListTableColumnsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTableColumnsResponseBody) SetErrorMessage(v string) *ListTableColumnsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListTableColumnsResponseBody) SetRequestId(v string) *ListTableColumnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTableColumnsResponseBody) SetSuccess(v bool) *ListTableColumnsResponseBody {
	s.Success = &v
	return s
}

func (s *ListTableColumnsResponseBody) Validate() error {
	if s.ColumnList != nil {
		if err := s.ColumnList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTableColumnsResponseBodyColumnList struct {
	Column []*ListTableColumnsResponseBodyColumnListColumn `json:"Column,omitempty" xml:"Column,omitempty" type:"Repeated"`
}

func (s ListTableColumnsResponseBodyColumnList) String() string {
	return dara.Prettify(s)
}

func (s ListTableColumnsResponseBodyColumnList) GoString() string {
	return s.String()
}

func (s *ListTableColumnsResponseBodyColumnList) GetColumn() []*ListTableColumnsResponseBodyColumnListColumn {
	return s.Column
}

func (s *ListTableColumnsResponseBodyColumnList) SetColumn(v []*ListTableColumnsResponseBodyColumnListColumn) *ListTableColumnsResponseBodyColumnList {
	s.Column = v
	return s
}

func (s *ListTableColumnsResponseBodyColumnList) Validate() error {
	if s.Column != nil {
		for _, item := range s.Column {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTableColumnsResponseBodyColumnListColumn struct {
	// Indicates whether the field is an auto-increment field. Valid values:
	//
	// 	- true: The field is an auto-increment field.
	//
	// 	- false: The field is not an auto-increment field.
	//
	// example:
	//
	// false
	AutoIncrement *bool `json:"AutoIncrement,omitempty" xml:"AutoIncrement,omitempty"`
	// The ID of the field.
	//
	// example:
	//
	// 12345
	ColumnId *string `json:"ColumnId,omitempty" xml:"ColumnId,omitempty"`
	// The field name.
	//
	// example:
	//
	// c1
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The data type of the field.
	//
	// example:
	//
	// varchar
	ColumnType *string `json:"ColumnType,omitempty" xml:"ColumnType,omitempty"`
	// The length of the field.
	//
	// example:
	//
	// 32
	DataLength *int64 `json:"DataLength,omitempty" xml:"DataLength,omitempty"`
	// The number of valid digits for the column.
	//
	// example:
	//
	// 0
	DataPrecision *int32 `json:"DataPrecision,omitempty" xml:"DataPrecision,omitempty"`
	// The number of decimal places of the field data.
	//
	// example:
	//
	// 0
	DataScale *int32 `json:"DataScale,omitempty" xml:"DataScale,omitempty"`
	// The default value of the column.
	//
	// example:
	//
	// aaa
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// The description of the field.
	//
	// example:
	//
	// column desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of the masking algorithm that is used for the field. Valid values:
	//
	// 	- null: No masking algorithm is used.
	//
	// 	- DEFAULT: A full masking algorithm is used.
	//
	// 	- FIX_POS: The fixed position is masked.
	//
	// 	- FIX_CHAR: The fixed characters are replaced.
	//
	// example:
	//
	// NULL
	FunctionType *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	// Indicates whether the field can be empty. Valid values:
	//
	// 	- true: The field can be empty.
	//
	// 	- false: The field cannot be empty.
	//
	// example:
	//
	// true
	Nullable *bool `json:"Nullable,omitempty" xml:"Nullable,omitempty"`
	// The security level of the field. Valid values:
	//
	// 	- INNER: The field is an internal field but not sensitive.
	//
	// 	- SENSITIVE: The field is sensitive.
	//
	// 	- CONFIDENTIAL: The field is a confidential column.
	//
	// example:
	//
	// INNER
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	// Indicates whether the field is a sensitive column. Valid values:
	//
	// 	- true: The field is a sensitive field.
	//
	// 	- false: The field is not a sensitive field.
	//
	// example:
	//
	// false
	Sensitive *bool `json:"Sensitive,omitempty" xml:"Sensitive,omitempty"`
}

func (s ListTableColumnsResponseBodyColumnListColumn) String() string {
	return dara.Prettify(s)
}

func (s ListTableColumnsResponseBodyColumnListColumn) GoString() string {
	return s.String()
}

func (s *ListTableColumnsResponseBodyColumnListColumn) GetAutoIncrement() *bool {
	return s.AutoIncrement
}

func (s *ListTableColumnsResponseBodyColumnListColumn) GetColumnId() *string {
	return s.ColumnId
}

func (s *ListTableColumnsResponseBodyColumnListColumn) GetColumnName() *string {
	return s.ColumnName
}

func (s *ListTableColumnsResponseBodyColumnListColumn) GetColumnType() *string {
	return s.ColumnType
}

func (s *ListTableColumnsResponseBodyColumnListColumn) GetDataLength() *int64 {
	return s.DataLength
}

func (s *ListTableColumnsResponseBodyColumnListColumn) GetDataPrecision() *int32 {
	return s.DataPrecision
}

func (s *ListTableColumnsResponseBodyColumnListColumn) GetDataScale() *int32 {
	return s.DataScale
}

func (s *ListTableColumnsResponseBodyColumnListColumn) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *ListTableColumnsResponseBodyColumnListColumn) GetDescription() *string {
	return s.Description
}

func (s *ListTableColumnsResponseBodyColumnListColumn) GetFunctionType() *string {
	return s.FunctionType
}

func (s *ListTableColumnsResponseBodyColumnListColumn) GetNullable() *bool {
	return s.Nullable
}

func (s *ListTableColumnsResponseBodyColumnListColumn) GetSecurityLevel() *string {
	return s.SecurityLevel
}

func (s *ListTableColumnsResponseBodyColumnListColumn) GetSensitive() *bool {
	return s.Sensitive
}

func (s *ListTableColumnsResponseBodyColumnListColumn) SetAutoIncrement(v bool) *ListTableColumnsResponseBodyColumnListColumn {
	s.AutoIncrement = &v
	return s
}

func (s *ListTableColumnsResponseBodyColumnListColumn) SetColumnId(v string) *ListTableColumnsResponseBodyColumnListColumn {
	s.ColumnId = &v
	return s
}

func (s *ListTableColumnsResponseBodyColumnListColumn) SetColumnName(v string) *ListTableColumnsResponseBodyColumnListColumn {
	s.ColumnName = &v
	return s
}

func (s *ListTableColumnsResponseBodyColumnListColumn) SetColumnType(v string) *ListTableColumnsResponseBodyColumnListColumn {
	s.ColumnType = &v
	return s
}

func (s *ListTableColumnsResponseBodyColumnListColumn) SetDataLength(v int64) *ListTableColumnsResponseBodyColumnListColumn {
	s.DataLength = &v
	return s
}

func (s *ListTableColumnsResponseBodyColumnListColumn) SetDataPrecision(v int32) *ListTableColumnsResponseBodyColumnListColumn {
	s.DataPrecision = &v
	return s
}

func (s *ListTableColumnsResponseBodyColumnListColumn) SetDataScale(v int32) *ListTableColumnsResponseBodyColumnListColumn {
	s.DataScale = &v
	return s
}

func (s *ListTableColumnsResponseBodyColumnListColumn) SetDefaultValue(v string) *ListTableColumnsResponseBodyColumnListColumn {
	s.DefaultValue = &v
	return s
}

func (s *ListTableColumnsResponseBodyColumnListColumn) SetDescription(v string) *ListTableColumnsResponseBodyColumnListColumn {
	s.Description = &v
	return s
}

func (s *ListTableColumnsResponseBodyColumnListColumn) SetFunctionType(v string) *ListTableColumnsResponseBodyColumnListColumn {
	s.FunctionType = &v
	return s
}

func (s *ListTableColumnsResponseBodyColumnListColumn) SetNullable(v bool) *ListTableColumnsResponseBodyColumnListColumn {
	s.Nullable = &v
	return s
}

func (s *ListTableColumnsResponseBodyColumnListColumn) SetSecurityLevel(v string) *ListTableColumnsResponseBodyColumnListColumn {
	s.SecurityLevel = &v
	return s
}

func (s *ListTableColumnsResponseBodyColumnListColumn) SetSensitive(v bool) *ListTableColumnsResponseBodyColumnListColumn {
	s.Sensitive = &v
	return s
}

func (s *ListTableColumnsResponseBodyColumnListColumn) Validate() error {
	return dara.Validate(s)
}
