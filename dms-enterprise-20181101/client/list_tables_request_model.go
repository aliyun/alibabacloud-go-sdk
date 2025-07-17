// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseId(v string) *ListTablesRequest
	GetDatabaseId() *string
	SetPageNumber(v int32) *ListTablesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTablesRequest
	GetPageSize() *int32
	SetReturnGuid(v bool) *ListTablesRequest
	GetReturnGuid() *bool
	SetSearchName(v string) *ListTablesRequest
	GetSearchName() *string
	SetTid(v int64) *ListTablesRequest
	GetTid() *int64
}

type ListTablesRequest struct {
	// The ID of the physical database. You can call the [ListDatabases](https://help.aliyun.com/document_detail/141873.html) or [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation to obtain the ID of the physical database.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1860****
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether to return the GUID of a table. Valid values:
	//
	// 	- **true**: returns the GUID of a table.
	//
	// 	- **false**: does not return the GUID of a table.
	//
	// example:
	//
	// true
	ReturnGuid *bool `json:"ReturnGuid,omitempty" xml:"ReturnGuid,omitempty"`
	// The name used to search for tables. Fuzzy search is supported.
	//
	// example:
	//
	// test
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTablesRequest) GoString() string {
	return s.String()
}

func (s *ListTablesRequest) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *ListTablesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTablesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTablesRequest) GetReturnGuid() *bool {
	return s.ReturnGuid
}

func (s *ListTablesRequest) GetSearchName() *string {
	return s.SearchName
}

func (s *ListTablesRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListTablesRequest) SetDatabaseId(v string) *ListTablesRequest {
	s.DatabaseId = &v
	return s
}

func (s *ListTablesRequest) SetPageNumber(v int32) *ListTablesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTablesRequest) SetPageSize(v int32) *ListTablesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTablesRequest) SetReturnGuid(v bool) *ListTablesRequest {
	s.ReturnGuid = &v
	return s
}

func (s *ListTablesRequest) SetSearchName(v string) *ListTablesRequest {
	s.SearchName = &v
	return s
}

func (s *ListTablesRequest) SetTid(v int64) *ListTablesRequest {
	s.Tid = &v
	return s
}

func (s *ListTablesRequest) Validate() error {
	return dara.Validate(s)
}
