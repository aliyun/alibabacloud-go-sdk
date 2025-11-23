// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSensitiveColumnsDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListSensitiveColumnsDetailResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListSensitiveColumnsDetailResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListSensitiveColumnsDetailResponseBody
	GetRequestId() *string
	SetSensitiveColumnsDetailList(v *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailList) *ListSensitiveColumnsDetailResponseBody
	GetSensitiveColumnsDetailList() *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailList
	SetSuccess(v bool) *ListSensitiveColumnsDetailResponseBody
	GetSuccess() *bool
}

type ListSensitiveColumnsDetailResponseBody struct {
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
	// 7629888F-C9FB-4D16-A7D3-B443FE06FBD4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the sensitive field.
	SensitiveColumnsDetailList *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailList `json:"SensitiveColumnsDetailList,omitempty" xml:"SensitiveColumnsDetailList,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSensitiveColumnsDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSensitiveColumnsDetailResponseBody) GoString() string {
	return s.String()
}

func (s *ListSensitiveColumnsDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListSensitiveColumnsDetailResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListSensitiveColumnsDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSensitiveColumnsDetailResponseBody) GetSensitiveColumnsDetailList() *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailList {
	return s.SensitiveColumnsDetailList
}

func (s *ListSensitiveColumnsDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSensitiveColumnsDetailResponseBody) SetErrorCode(v string) *ListSensitiveColumnsDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSensitiveColumnsDetailResponseBody) SetErrorMessage(v string) *ListSensitiveColumnsDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListSensitiveColumnsDetailResponseBody) SetRequestId(v string) *ListSensitiveColumnsDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSensitiveColumnsDetailResponseBody) SetSensitiveColumnsDetailList(v *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailList) *ListSensitiveColumnsDetailResponseBody {
	s.SensitiveColumnsDetailList = v
	return s
}

func (s *ListSensitiveColumnsDetailResponseBody) SetSuccess(v bool) *ListSensitiveColumnsDetailResponseBody {
	s.Success = &v
	return s
}

func (s *ListSensitiveColumnsDetailResponseBody) Validate() error {
	if s.SensitiveColumnsDetailList != nil {
		if err := s.SensitiveColumnsDetailList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailList struct {
	SensitiveColumnsDetail []*ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail `json:"SensitiveColumnsDetail,omitempty" xml:"SensitiveColumnsDetail,omitempty" type:"Repeated"`
}

func (s ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailList) String() string {
	return dara.Prettify(s)
}

func (s ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailList) GoString() string {
	return s.String()
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailList) GetSensitiveColumnsDetail() []*ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail {
	return s.SensitiveColumnsDetail
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailList) SetSensitiveColumnsDetail(v []*ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailList {
	s.SensitiveColumnsDetail = v
	return s
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailList) Validate() error {
	if s.SensitiveColumnsDetail != nil {
		for _, item := range s.SensitiveColumnsDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail struct {
	// The description of the field.
	//
	// example:
	//
	// test
	ColumnDescription *string `json:"ColumnDescription,omitempty" xml:"ColumnDescription,omitempty"`
	// The name of the field.
	//
	// example:
	//
	// ColumnName_test
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The data type of the field.
	//
	// example:
	//
	// varchar(32)
	ColumnType *string `json:"ColumnType,omitempty" xml:"ColumnType,omitempty"`
	// The ID of the database.
	//
	// example:
	//
	// 1860****
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The type of the database.
	//
	// example:
	//
	// mysql
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The type of the environment to which the database belongs.
	//
	// example:
	//
	// product
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// Indicates whether the database is a logical database. Valid values:
	//
	// 	- **true**: The database is a logical database.
	//
	// 	- **false**: The database is not a logical database.
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// SchemaName_test
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name that is used to search for the database.
	//
	// example:
	//
	// test@xxx:3306
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) String() string {
	return dara.Prettify(s)
}

func (s ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) GoString() string {
	return s.String()
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) GetColumnDescription() *string {
	return s.ColumnDescription
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) GetColumnName() *string {
	return s.ColumnName
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) GetColumnType() *string {
	return s.ColumnType
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) GetDbId() *int64 {
	return s.DbId
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) GetDbType() *string {
	return s.DbType
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) GetEnvType() *string {
	return s.EnvType
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) GetLogic() *bool {
	return s.Logic
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) GetSearchName() *string {
	return s.SearchName
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) GetTableName() *string {
	return s.TableName
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) SetColumnDescription(v string) *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail {
	s.ColumnDescription = &v
	return s
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) SetColumnName(v string) *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail {
	s.ColumnName = &v
	return s
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) SetColumnType(v string) *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail {
	s.ColumnType = &v
	return s
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) SetDbId(v int64) *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail {
	s.DbId = &v
	return s
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) SetDbType(v string) *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail {
	s.DbType = &v
	return s
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) SetEnvType(v string) *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail {
	s.EnvType = &v
	return s
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) SetLogic(v bool) *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail {
	s.Logic = &v
	return s
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) SetSchemaName(v string) *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail {
	s.SchemaName = &v
	return s
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) SetSearchName(v string) *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail {
	s.SearchName = &v
	return s
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) SetTableName(v string) *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail {
	s.TableName = &v
	return s
}

func (s *ListSensitiveColumnsDetailResponseBodySensitiveColumnsDetailListSensitiveColumnsDetail) Validate() error {
	return dara.Validate(s)
}
