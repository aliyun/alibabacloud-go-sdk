// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSimulationTaskListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSimulationTaskListResponseBody
	GetRequestId() *string
	SetCurrentPage(v int32) *DescribeSimulationTaskListResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeSimulationTaskListResponseBody
	GetPageSize() *int32
	SetResultObject(v bool) *DescribeSimulationTaskListResponseBody
	GetResultObject() *bool
	SetTotalItem(v int32) *DescribeSimulationTaskListResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeSimulationTaskListResponseBody
	GetTotalPage() *int32
}

type DescribeSimulationTaskListResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Page size, with a default value of 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Return object
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
	// Total count.
	//
	// example:
	//
	// 6
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total pages
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeSimulationTaskListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSimulationTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSimulationTaskListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSimulationTaskListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSimulationTaskListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSimulationTaskListResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeSimulationTaskListResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeSimulationTaskListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeSimulationTaskListResponseBody) SetRequestId(v string) *DescribeSimulationTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSimulationTaskListResponseBody) SetCurrentPage(v int32) *DescribeSimulationTaskListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSimulationTaskListResponseBody) SetPageSize(v int32) *DescribeSimulationTaskListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSimulationTaskListResponseBody) SetResultObject(v bool) *DescribeSimulationTaskListResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeSimulationTaskListResponseBody) SetTotalItem(v int32) *DescribeSimulationTaskListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeSimulationTaskListResponseBody) SetTotalPage(v int32) *DescribeSimulationTaskListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeSimulationTaskListResponseBody) Validate() error {
	return dara.Validate(s)
}
