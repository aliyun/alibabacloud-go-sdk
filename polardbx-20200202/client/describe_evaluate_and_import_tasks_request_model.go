// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEvaluateAndImportTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeEvaluateAndImportTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeEvaluateAndImportTasksRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeEvaluateAndImportTasksRequest
	GetRegionId() *string
}

type DescribeEvaluateAndImportTasksRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 0
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeEvaluateAndImportTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEvaluateAndImportTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeEvaluateAndImportTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeEvaluateAndImportTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEvaluateAndImportTasksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEvaluateAndImportTasksRequest) SetPageNumber(v int32) *DescribeEvaluateAndImportTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEvaluateAndImportTasksRequest) SetPageSize(v int32) *DescribeEvaluateAndImportTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEvaluateAndImportTasksRequest) SetRegionId(v string) *DescribeEvaluateAndImportTasksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEvaluateAndImportTasksRequest) Validate() error {
	return dara.Validate(s)
}
