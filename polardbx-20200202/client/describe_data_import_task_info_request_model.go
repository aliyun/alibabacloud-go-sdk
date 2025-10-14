// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataImportTaskInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFailPageNumber(v int32) *DescribeDataImportTaskInfoRequest
	GetFailPageNumber() *int32
	SetFailPageSize(v int32) *DescribeDataImportTaskInfoRequest
	GetFailPageSize() *int32
	SetRegionId(v string) *DescribeDataImportTaskInfoRequest
	GetRegionId() *string
	SetSlinkTaskId(v string) *DescribeDataImportTaskInfoRequest
	GetSlinkTaskId() *string
	SetSuccessPageNumber(v int64) *DescribeDataImportTaskInfoRequest
	GetSuccessPageNumber() *int64
	SetSuccessPageSize(v int64) *DescribeDataImportTaskInfoRequest
	GetSuccessPageSize() *int64
}

type DescribeDataImportTaskInfoRequest struct {
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
	// This parameter is required.
	//
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

func (s DescribeDataImportTaskInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataImportTaskInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataImportTaskInfoRequest) GetFailPageNumber() *int32 {
	return s.FailPageNumber
}

func (s *DescribeDataImportTaskInfoRequest) GetFailPageSize() *int32 {
	return s.FailPageSize
}

func (s *DescribeDataImportTaskInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDataImportTaskInfoRequest) GetSlinkTaskId() *string {
	return s.SlinkTaskId
}

func (s *DescribeDataImportTaskInfoRequest) GetSuccessPageNumber() *int64 {
	return s.SuccessPageNumber
}

func (s *DescribeDataImportTaskInfoRequest) GetSuccessPageSize() *int64 {
	return s.SuccessPageSize
}

func (s *DescribeDataImportTaskInfoRequest) SetFailPageNumber(v int32) *DescribeDataImportTaskInfoRequest {
	s.FailPageNumber = &v
	return s
}

func (s *DescribeDataImportTaskInfoRequest) SetFailPageSize(v int32) *DescribeDataImportTaskInfoRequest {
	s.FailPageSize = &v
	return s
}

func (s *DescribeDataImportTaskInfoRequest) SetRegionId(v string) *DescribeDataImportTaskInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDataImportTaskInfoRequest) SetSlinkTaskId(v string) *DescribeDataImportTaskInfoRequest {
	s.SlinkTaskId = &v
	return s
}

func (s *DescribeDataImportTaskInfoRequest) SetSuccessPageNumber(v int64) *DescribeDataImportTaskInfoRequest {
	s.SuccessPageNumber = &v
	return s
}

func (s *DescribeDataImportTaskInfoRequest) SetSuccessPageSize(v int64) *DescribeDataImportTaskInfoRequest {
	s.SuccessPageSize = &v
	return s
}

func (s *DescribeDataImportTaskInfoRequest) Validate() error {
	return dara.Validate(s)
}
