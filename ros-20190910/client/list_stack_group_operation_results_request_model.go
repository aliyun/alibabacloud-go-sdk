// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStackGroupOperationResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperationId(v string) *ListStackGroupOperationResultsRequest
	GetOperationId() *string
	SetPageNumber(v int64) *ListStackGroupOperationResultsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListStackGroupOperationResultsRequest
	GetPageSize() *int64
	SetRegionId(v string) *ListStackGroupOperationResultsRequest
	GetRegionId() *string
}

type ListStackGroupOperationResultsRequest struct {
	// The ID of the operation.
	//
	// You can call the [ListStackGroupOperations](https://help.aliyun.com/document_detail/151342.html) operation to query the operation ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6da106ca-1784-4a6f-a7e1-e723863d****
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The number of the page to return.
	//
	// 	- Pages start from page 1.
	//
	// 	- Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// 	- Valid values: 1 to 50.
	//
	// 	- Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the stack group.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListStackGroupOperationResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStackGroupOperationResultsRequest) GoString() string {
	return s.String()
}

func (s *ListStackGroupOperationResultsRequest) GetOperationId() *string {
	return s.OperationId
}

func (s *ListStackGroupOperationResultsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListStackGroupOperationResultsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListStackGroupOperationResultsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListStackGroupOperationResultsRequest) SetOperationId(v string) *ListStackGroupOperationResultsRequest {
	s.OperationId = &v
	return s
}

func (s *ListStackGroupOperationResultsRequest) SetPageNumber(v int64) *ListStackGroupOperationResultsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStackGroupOperationResultsRequest) SetPageSize(v int64) *ListStackGroupOperationResultsRequest {
	s.PageSize = &v
	return s
}

func (s *ListStackGroupOperationResultsRequest) SetRegionId(v string) *ListStackGroupOperationResultsRequest {
	s.RegionId = &v
	return s
}

func (s *ListStackGroupOperationResultsRequest) Validate() error {
	return dara.Validate(s)
}
