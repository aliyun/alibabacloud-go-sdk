// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuditPageListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeAuditPageListResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeAuditPageListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAuditPageListResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribeAuditPageListResponseBody
	GetResultObject() *bool
	SetTotalItem(v int32) *DescribeAuditPageListResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int64) *DescribeAuditPageListResponseBody
	GetTotalPage() *int64
}

type DescribeAuditPageListResponseBody struct {
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Number of records per page, default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Return object
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
	// Total items
	//
	// example:
	//
	// 6
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total pages
	//
	// example:
	//
	// 4
	TotalPage *int64 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeAuditPageListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditPageListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuditPageListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeAuditPageListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAuditPageListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAuditPageListResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeAuditPageListResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeAuditPageListResponseBody) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *DescribeAuditPageListResponseBody) SetCurrentPage(v int32) *DescribeAuditPageListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAuditPageListResponseBody) SetPageSize(v int32) *DescribeAuditPageListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAuditPageListResponseBody) SetRequestId(v string) *DescribeAuditPageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuditPageListResponseBody) SetResultObject(v bool) *DescribeAuditPageListResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeAuditPageListResponseBody) SetTotalItem(v int32) *DescribeAuditPageListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeAuditPageListResponseBody) SetTotalPage(v int64) *DescribeAuditPageListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeAuditPageListResponseBody) Validate() error {
	return dara.Validate(s)
}
