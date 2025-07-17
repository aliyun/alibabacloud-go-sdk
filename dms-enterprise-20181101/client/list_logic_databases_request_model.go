// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogicDatabasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListLogicDatabasesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListLogicDatabasesRequest
	GetPageSize() *int32
	SetTid(v int64) *ListLogicDatabasesRequest
	GetTid() *int64
}

type ListLogicDatabasesRequest struct {
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
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://www.alibabacloud.com/help/en/data-management-service/latest/getuseractivetenant) operation to query the tenant ID.
	//
	// example:
	//
	// 3422
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListLogicDatabasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLogicDatabasesRequest) GoString() string {
	return s.String()
}

func (s *ListLogicDatabasesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLogicDatabasesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLogicDatabasesRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListLogicDatabasesRequest) SetPageNumber(v int32) *ListLogicDatabasesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLogicDatabasesRequest) SetPageSize(v int32) *ListLogicDatabasesRequest {
	s.PageSize = &v
	return s
}

func (s *ListLogicDatabasesRequest) SetTid(v int64) *ListLogicDatabasesRequest {
	s.Tid = &v
	return s
}

func (s *ListLogicDatabasesRequest) Validate() error {
	return dara.Validate(s)
}
