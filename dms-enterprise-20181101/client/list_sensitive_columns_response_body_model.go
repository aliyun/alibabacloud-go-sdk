// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSensitiveColumnsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListSensitiveColumnsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListSensitiveColumnsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListSensitiveColumnsResponseBody
	GetRequestId() *string
	SetSensitiveColumnList(v *ListSensitiveColumnsResponseBodySensitiveColumnList) *ListSensitiveColumnsResponseBody
	GetSensitiveColumnList() *ListSensitiveColumnsResponseBodySensitiveColumnList
	SetSuccess(v bool) *ListSensitiveColumnsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListSensitiveColumnsResponseBody
	GetTotalCount() *int64
}

type ListSensitiveColumnsResponseBody struct {
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 09D82FD7-F87F-59EF-AA82-AEF71B09E306
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The sensitive fields.
	SensitiveColumnList *ListSensitiveColumnsResponseBodySensitiveColumnList `json:"SensitiveColumnList,omitempty" xml:"SensitiveColumnList,omitempty" type:"Struct"`
	// Indicates whether the request is successful. Valid values:
	//
	// - true: The request is successful.
	//
	// - false: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSensitiveColumnsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSensitiveColumnsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSensitiveColumnsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListSensitiveColumnsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListSensitiveColumnsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSensitiveColumnsResponseBody) GetSensitiveColumnList() *ListSensitiveColumnsResponseBodySensitiveColumnList {
	return s.SensitiveColumnList
}

func (s *ListSensitiveColumnsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSensitiveColumnsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListSensitiveColumnsResponseBody) SetErrorCode(v string) *ListSensitiveColumnsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSensitiveColumnsResponseBody) SetErrorMessage(v string) *ListSensitiveColumnsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListSensitiveColumnsResponseBody) SetRequestId(v string) *ListSensitiveColumnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSensitiveColumnsResponseBody) SetSensitiveColumnList(v *ListSensitiveColumnsResponseBodySensitiveColumnList) *ListSensitiveColumnsResponseBody {
	s.SensitiveColumnList = v
	return s
}

func (s *ListSensitiveColumnsResponseBody) SetSuccess(v bool) *ListSensitiveColumnsResponseBody {
	s.Success = &v
	return s
}

func (s *ListSensitiveColumnsResponseBody) SetTotalCount(v int64) *ListSensitiveColumnsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSensitiveColumnsResponseBody) Validate() error {
	if s.SensitiveColumnList != nil {
		if err := s.SensitiveColumnList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSensitiveColumnsResponseBodySensitiveColumnList struct {
	SensitiveColumn []*ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn `json:"SensitiveColumn,omitempty" xml:"SensitiveColumn,omitempty" type:"Repeated"`
}

func (s ListSensitiveColumnsResponseBodySensitiveColumnList) String() string {
	return dara.Prettify(s)
}

func (s ListSensitiveColumnsResponseBodySensitiveColumnList) GoString() string {
	return s.String()
}

func (s *ListSensitiveColumnsResponseBodySensitiveColumnList) GetSensitiveColumn() []*ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn {
	return s.SensitiveColumn
}

func (s *ListSensitiveColumnsResponseBodySensitiveColumnList) SetSensitiveColumn(v []*ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn) *ListSensitiveColumnsResponseBodySensitiveColumnList {
	s.SensitiveColumn = v
	return s
}

func (s *ListSensitiveColumnsResponseBodySensitiveColumnList) Validate() error {
	if s.SensitiveColumn != nil {
		for _, item := range s.SensitiveColumn {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn struct {
	// The number of sensitive fields.
	//
	// example:
	//
	// 1
	ColumnCount *int64 `json:"ColumnCount,omitempty" xml:"ColumnCount,omitempty"`
	// The name of the field.
	//
	// example:
	//
	// test_column
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The type of the de-identification algorithm. Valid values:
	//
	// 	- DEFAULT: All characters are masked. This is the default value.
	//
	// 	- FIX_POS: The characters at specific positions are masked.
	//
	// 	- FIX_CHAR: Specific characters are masked.
	//
	// example:
	//
	// DEFAULT
	FunctionType *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// test_schema
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The sensitivity level of the field. Valid values:
	//
	// 	- SENSITIVE
	//
	// 	- CONFIDENTIAL
	//
	// example:
	//
	// SENSITIVE
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn) String() string {
	return dara.Prettify(s)
}

func (s ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn) GoString() string {
	return s.String()
}

func (s *ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn) GetColumnCount() *int64 {
	return s.ColumnCount
}

func (s *ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn) GetColumnName() *string {
	return s.ColumnName
}

func (s *ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn) GetFunctionType() *string {
	return s.FunctionType
}

func (s *ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn) GetSecurityLevel() *string {
	return s.SecurityLevel
}

func (s *ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn) GetTableName() *string {
	return s.TableName
}

func (s *ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn) SetColumnCount(v int64) *ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn {
	s.ColumnCount = &v
	return s
}

func (s *ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn) SetColumnName(v string) *ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn {
	s.ColumnName = &v
	return s
}

func (s *ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn) SetFunctionType(v string) *ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn {
	s.FunctionType = &v
	return s
}

func (s *ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn) SetSchemaName(v string) *ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn {
	s.SchemaName = &v
	return s
}

func (s *ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn) SetSecurityLevel(v string) *ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn {
	s.SecurityLevel = &v
	return s
}

func (s *ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn) SetTableName(v string) *ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn {
	s.TableName = &v
	return s
}

func (s *ListSensitiveColumnsResponseBodySensitiveColumnListSensitiveColumn) Validate() error {
	return dara.Validate(s)
}
