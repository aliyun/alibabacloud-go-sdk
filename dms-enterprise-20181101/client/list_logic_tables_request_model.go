// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogicTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseId(v string) *ListLogicTablesRequest
	GetDatabaseId() *string
	SetPageNumber(v int32) *ListLogicTablesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListLogicTablesRequest
	GetPageSize() *int32
	SetReturnGuid(v bool) *ListLogicTablesRequest
	GetReturnGuid() *bool
	SetSearchName(v string) *ListLogicTablesRequest
	GetSearchName() *string
	SetTid(v int64) *ListLogicTablesRequest
	GetTid() *int64
}

type ListLogicTablesRequest struct {
	// The ID of the logical database.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
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
	// Specifies whether to return the GUID of the table.
	//
	// example:
	//
	// true
	ReturnGuid *bool `json:"ReturnGuid,omitempty" xml:"ReturnGuid,omitempty"`
	// The keyword that is used to search for the logical tables. Prefix match is supported.
	//
	// example:
	//
	// test
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	// The ID of the tenant.
	//
	// > To view the tenant ID, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// -1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListLogicTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLogicTablesRequest) GoString() string {
	return s.String()
}

func (s *ListLogicTablesRequest) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *ListLogicTablesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLogicTablesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLogicTablesRequest) GetReturnGuid() *bool {
	return s.ReturnGuid
}

func (s *ListLogicTablesRequest) GetSearchName() *string {
	return s.SearchName
}

func (s *ListLogicTablesRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListLogicTablesRequest) SetDatabaseId(v string) *ListLogicTablesRequest {
	s.DatabaseId = &v
	return s
}

func (s *ListLogicTablesRequest) SetPageNumber(v int32) *ListLogicTablesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLogicTablesRequest) SetPageSize(v int32) *ListLogicTablesRequest {
	s.PageSize = &v
	return s
}

func (s *ListLogicTablesRequest) SetReturnGuid(v bool) *ListLogicTablesRequest {
	s.ReturnGuid = &v
	return s
}

func (s *ListLogicTablesRequest) SetSearchName(v string) *ListLogicTablesRequest {
	s.SearchName = &v
	return s
}

func (s *ListLogicTablesRequest) SetTid(v int64) *ListLogicTablesRequest {
	s.Tid = &v
	return s
}

func (s *ListLogicTablesRequest) Validate() error {
	return dara.Validate(s)
}
