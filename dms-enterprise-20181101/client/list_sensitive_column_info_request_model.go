// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSensitiveColumnInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumnName(v string) *ListSensitiveColumnInfoRequest
	GetColumnName() *string
	SetInstanceId(v int32) *ListSensitiveColumnInfoRequest
	GetInstanceId() *int32
	SetPageNumber(v int32) *ListSensitiveColumnInfoRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSensitiveColumnInfoRequest
	GetPageSize() *int32
	SetSchemaName(v string) *ListSensitiveColumnInfoRequest
	GetSchemaName() *string
	SetTableName(v string) *ListSensitiveColumnInfoRequest
	GetTableName() *string
	SetTid(v int64) *ListSensitiveColumnInfoRequest
	GetTid() *int64
}

type ListSensitiveColumnInfoRequest struct {
	// example:
	//
	// test_column
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 183****
	InstanceId *int32 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// test_schema
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListSensitiveColumnInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSensitiveColumnInfoRequest) GoString() string {
	return s.String()
}

func (s *ListSensitiveColumnInfoRequest) GetColumnName() *string {
	return s.ColumnName
}

func (s *ListSensitiveColumnInfoRequest) GetInstanceId() *int32 {
	return s.InstanceId
}

func (s *ListSensitiveColumnInfoRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSensitiveColumnInfoRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSensitiveColumnInfoRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListSensitiveColumnInfoRequest) GetTableName() *string {
	return s.TableName
}

func (s *ListSensitiveColumnInfoRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListSensitiveColumnInfoRequest) SetColumnName(v string) *ListSensitiveColumnInfoRequest {
	s.ColumnName = &v
	return s
}

func (s *ListSensitiveColumnInfoRequest) SetInstanceId(v int32) *ListSensitiveColumnInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSensitiveColumnInfoRequest) SetPageNumber(v int32) *ListSensitiveColumnInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSensitiveColumnInfoRequest) SetPageSize(v int32) *ListSensitiveColumnInfoRequest {
	s.PageSize = &v
	return s
}

func (s *ListSensitiveColumnInfoRequest) SetSchemaName(v string) *ListSensitiveColumnInfoRequest {
	s.SchemaName = &v
	return s
}

func (s *ListSensitiveColumnInfoRequest) SetTableName(v string) *ListSensitiveColumnInfoRequest {
	s.TableName = &v
	return s
}

func (s *ListSensitiveColumnInfoRequest) SetTid(v int64) *ListSensitiveColumnInfoRequest {
	s.Tid = &v
	return s
}

func (s *ListSensitiveColumnInfoRequest) Validate() error {
	return dara.Validate(s)
}
