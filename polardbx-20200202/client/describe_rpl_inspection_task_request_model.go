// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRplInspectionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFailPageNumber(v int32) *DescribeRplInspectionTaskRequest
	GetFailPageNumber() *int32
	SetFailPageSize(v int32) *DescribeRplInspectionTaskRequest
	GetFailPageSize() *int32
	SetRegionId(v string) *DescribeRplInspectionTaskRequest
	GetRegionId() *string
	SetSlinkTaskId(v string) *DescribeRplInspectionTaskRequest
	GetSlinkTaskId() *string
	SetSuccessPageNumber(v int64) *DescribeRplInspectionTaskRequest
	GetSuccessPageNumber() *int64
	SetSuccessPageSize(v int64) *DescribeRplInspectionTaskRequest
	GetSuccessPageSize() *int64
}

type DescribeRplInspectionTaskRequest struct {
	// example:
	//
	// 1
	FailPageNumber *int32 `json:"FailPageNumber,omitempty" xml:"FailPageNumber,omitempty"`
	// example:
	//
	// 10
	FailPageSize *int32 `json:"FailPageSize,omitempty" xml:"FailPageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// etx-szr2rr6i*****
	SlinkTaskId *string `json:"SlinkTaskId,omitempty" xml:"SlinkTaskId,omitempty"`
	// example:
	//
	// 1
	SuccessPageNumber *int64 `json:"SuccessPageNumber,omitempty" xml:"SuccessPageNumber,omitempty"`
	// example:
	//
	// 15
	SuccessPageSize *int64 `json:"SuccessPageSize,omitempty" xml:"SuccessPageSize,omitempty"`
}

func (s DescribeRplInspectionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRplInspectionTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeRplInspectionTaskRequest) GetFailPageNumber() *int32 {
	return s.FailPageNumber
}

func (s *DescribeRplInspectionTaskRequest) GetFailPageSize() *int32 {
	return s.FailPageSize
}

func (s *DescribeRplInspectionTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRplInspectionTaskRequest) GetSlinkTaskId() *string {
	return s.SlinkTaskId
}

func (s *DescribeRplInspectionTaskRequest) GetSuccessPageNumber() *int64 {
	return s.SuccessPageNumber
}

func (s *DescribeRplInspectionTaskRequest) GetSuccessPageSize() *int64 {
	return s.SuccessPageSize
}

func (s *DescribeRplInspectionTaskRequest) SetFailPageNumber(v int32) *DescribeRplInspectionTaskRequest {
	s.FailPageNumber = &v
	return s
}

func (s *DescribeRplInspectionTaskRequest) SetFailPageSize(v int32) *DescribeRplInspectionTaskRequest {
	s.FailPageSize = &v
	return s
}

func (s *DescribeRplInspectionTaskRequest) SetRegionId(v string) *DescribeRplInspectionTaskRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRplInspectionTaskRequest) SetSlinkTaskId(v string) *DescribeRplInspectionTaskRequest {
	s.SlinkTaskId = &v
	return s
}

func (s *DescribeRplInspectionTaskRequest) SetSuccessPageNumber(v int64) *DescribeRplInspectionTaskRequest {
	s.SuccessPageNumber = &v
	return s
}

func (s *DescribeRplInspectionTaskRequest) SetSuccessPageSize(v int64) *DescribeRplInspectionTaskRequest {
	s.SuccessPageSize = &v
	return s
}

func (s *DescribeRplInspectionTaskRequest) Validate() error {
	return dara.Validate(s)
}
