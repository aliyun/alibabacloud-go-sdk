// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListColumnsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetColumnList(v *ListColumnsResponseBodyColumnList) *ListColumnsResponseBody
	GetColumnList() *ListColumnsResponseBodyColumnList
	SetErrorCode(v string) *ListColumnsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListColumnsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListColumnsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListColumnsResponseBody
	GetSuccess() *bool
}

type ListColumnsResponseBody struct {
	// The details about columns.
	ColumnList *ListColumnsResponseBodyColumnList `json:"ColumnList,omitempty" xml:"ColumnList,omitempty" type:"Struct"`
	// The error code returned.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0AD9AC55-5873-474A-9F33-4285806A3619
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

func (s ListColumnsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListColumnsResponseBody) GoString() string {
	return s.String()
}

func (s *ListColumnsResponseBody) GetColumnList() *ListColumnsResponseBodyColumnList {
	return s.ColumnList
}

func (s *ListColumnsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListColumnsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListColumnsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListColumnsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListColumnsResponseBody) SetColumnList(v *ListColumnsResponseBodyColumnList) *ListColumnsResponseBody {
	s.ColumnList = v
	return s
}

func (s *ListColumnsResponseBody) SetErrorCode(v string) *ListColumnsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListColumnsResponseBody) SetErrorMessage(v string) *ListColumnsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListColumnsResponseBody) SetRequestId(v string) *ListColumnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListColumnsResponseBody) SetSuccess(v bool) *ListColumnsResponseBody {
	s.Success = &v
	return s
}

func (s *ListColumnsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListColumnsResponseBodyColumnList struct {
	Column []*ListColumnsResponseBodyColumnListColumn `json:"Column,omitempty" xml:"Column,omitempty" type:"Repeated"`
}

func (s ListColumnsResponseBodyColumnList) String() string {
	return dara.Prettify(s)
}

func (s ListColumnsResponseBodyColumnList) GoString() string {
	return s.String()
}

func (s *ListColumnsResponseBodyColumnList) GetColumn() []*ListColumnsResponseBodyColumnListColumn {
	return s.Column
}

func (s *ListColumnsResponseBodyColumnList) SetColumn(v []*ListColumnsResponseBodyColumnListColumn) *ListColumnsResponseBodyColumnList {
	s.Column = v
	return s
}

func (s *ListColumnsResponseBodyColumnList) Validate() error {
	return dara.Validate(s)
}

type ListColumnsResponseBodyColumnListColumn struct {
	// Indicates whether the column is an auto-increment column. Valid values:
	//
	// 	- true: The column is an auto-increment column.
	//
	// 	- false: The column is not an auto-increment column.
	//
	// example:
	//
	// false
	AutoIncrement *bool `json:"AutoIncrement,omitempty" xml:"AutoIncrement,omitempty"`
	// The ID of the column.
	//
	// example:
	//
	// 62589****
	ColumnId *string `json:"ColumnId,omitempty" xml:"ColumnId,omitempty"`
	// The name of the column.
	//
	// example:
	//
	// name
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The data type of the column.
	//
	// example:
	//
	// varchar
	ColumnType *string `json:"ColumnType,omitempty" xml:"ColumnType,omitempty"`
	// The length of the field.
	//
	// example:
	//
	// 4
	DataLength *int64 `json:"DataLength,omitempty" xml:"DataLength,omitempty"`
	// The number of valid digits for the field.
	//
	// example:
	//
	// 0
	DataPrecision *int32 `json:"DataPrecision,omitempty" xml:"DataPrecision,omitempty"`
	// The number of decimal places for the field.
	//
	// example:
	//
	// 0
	DataScale *int32 `json:"DataScale,omitempty" xml:"DataScale,omitempty"`
	// The default value of the column.
	//
	// example:
	//
	// def_value
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// The description of the column.
	//
	// example:
	//
	// test
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
	// DEFAULT
	FunctionType *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	// Indicates whether the column can be empty. Valid values:
	//
	// 	- **true**: The column can be empty.
	//
	// 	- **false**: The column cannot be empty.
	//
	// example:
	//
	// false
	Nullable *bool `json:"Nullable,omitempty" xml:"Nullable,omitempty"`
	// The security level of the column. Valid values:
	//
	// 	- INNER: The column is an internal column but not sensitive.
	//
	// 	- SENSITIVE: The column is a sensitive column.
	//
	// 	- CONFIDENTIAL: The column is a confidential column.
	//
	// > For more information, see [Sensitivity levels of fields](https://help.aliyun.com/document_detail/66091.html).
	//
	// example:
	//
	// INNER
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	// Indicates whether the column is a sensitive column. Valid values:
	//
	// 	- **true**: The column is a sensitive column.
	//
	// 	- **false**: The column is not a sensitive column.
	//
	// example:
	//
	// false
	Sensitive *bool `json:"Sensitive,omitempty" xml:"Sensitive,omitempty"`
}

func (s ListColumnsResponseBodyColumnListColumn) String() string {
	return dara.Prettify(s)
}

func (s ListColumnsResponseBodyColumnListColumn) GoString() string {
	return s.String()
}

func (s *ListColumnsResponseBodyColumnListColumn) GetAutoIncrement() *bool {
	return s.AutoIncrement
}

func (s *ListColumnsResponseBodyColumnListColumn) GetColumnId() *string {
	return s.ColumnId
}

func (s *ListColumnsResponseBodyColumnListColumn) GetColumnName() *string {
	return s.ColumnName
}

func (s *ListColumnsResponseBodyColumnListColumn) GetColumnType() *string {
	return s.ColumnType
}

func (s *ListColumnsResponseBodyColumnListColumn) GetDataLength() *int64 {
	return s.DataLength
}

func (s *ListColumnsResponseBodyColumnListColumn) GetDataPrecision() *int32 {
	return s.DataPrecision
}

func (s *ListColumnsResponseBodyColumnListColumn) GetDataScale() *int32 {
	return s.DataScale
}

func (s *ListColumnsResponseBodyColumnListColumn) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *ListColumnsResponseBodyColumnListColumn) GetDescription() *string {
	return s.Description
}

func (s *ListColumnsResponseBodyColumnListColumn) GetFunctionType() *string {
	return s.FunctionType
}

func (s *ListColumnsResponseBodyColumnListColumn) GetNullable() *bool {
	return s.Nullable
}

func (s *ListColumnsResponseBodyColumnListColumn) GetSecurityLevel() *string {
	return s.SecurityLevel
}

func (s *ListColumnsResponseBodyColumnListColumn) GetSensitive() *bool {
	return s.Sensitive
}

func (s *ListColumnsResponseBodyColumnListColumn) SetAutoIncrement(v bool) *ListColumnsResponseBodyColumnListColumn {
	s.AutoIncrement = &v
	return s
}

func (s *ListColumnsResponseBodyColumnListColumn) SetColumnId(v string) *ListColumnsResponseBodyColumnListColumn {
	s.ColumnId = &v
	return s
}

func (s *ListColumnsResponseBodyColumnListColumn) SetColumnName(v string) *ListColumnsResponseBodyColumnListColumn {
	s.ColumnName = &v
	return s
}

func (s *ListColumnsResponseBodyColumnListColumn) SetColumnType(v string) *ListColumnsResponseBodyColumnListColumn {
	s.ColumnType = &v
	return s
}

func (s *ListColumnsResponseBodyColumnListColumn) SetDataLength(v int64) *ListColumnsResponseBodyColumnListColumn {
	s.DataLength = &v
	return s
}

func (s *ListColumnsResponseBodyColumnListColumn) SetDataPrecision(v int32) *ListColumnsResponseBodyColumnListColumn {
	s.DataPrecision = &v
	return s
}

func (s *ListColumnsResponseBodyColumnListColumn) SetDataScale(v int32) *ListColumnsResponseBodyColumnListColumn {
	s.DataScale = &v
	return s
}

func (s *ListColumnsResponseBodyColumnListColumn) SetDefaultValue(v string) *ListColumnsResponseBodyColumnListColumn {
	s.DefaultValue = &v
	return s
}

func (s *ListColumnsResponseBodyColumnListColumn) SetDescription(v string) *ListColumnsResponseBodyColumnListColumn {
	s.Description = &v
	return s
}

func (s *ListColumnsResponseBodyColumnListColumn) SetFunctionType(v string) *ListColumnsResponseBodyColumnListColumn {
	s.FunctionType = &v
	return s
}

func (s *ListColumnsResponseBodyColumnListColumn) SetNullable(v bool) *ListColumnsResponseBodyColumnListColumn {
	s.Nullable = &v
	return s
}

func (s *ListColumnsResponseBodyColumnListColumn) SetSecurityLevel(v string) *ListColumnsResponseBodyColumnListColumn {
	s.SecurityLevel = &v
	return s
}

func (s *ListColumnsResponseBodyColumnListColumn) SetSensitive(v bool) *ListColumnsResponseBodyColumnListColumn {
	s.Sensitive = &v
	return s
}

func (s *ListColumnsResponseBodyColumnListColumn) Validate() error {
	return dara.Validate(s)
}
