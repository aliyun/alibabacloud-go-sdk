// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOssV2ResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *DescribeOssV2ResultRequest
	GetBucket() *string
	SetCurrentPage(v int32) *DescribeOssV2ResultRequest
	GetCurrentPage() *int32
	SetEndDate(v string) *DescribeOssV2ResultRequest
	GetEndDate() *string
	SetPageSize(v int32) *DescribeOssV2ResultRequest
	GetPageSize() *int32
	SetRiskLevel(v string) *DescribeOssV2ResultRequest
	GetRiskLevel() *string
	SetStartDate(v string) *DescribeOssV2ResultRequest
	GetStartDate() *string
	SetTaskName(v string) *DescribeOssV2ResultRequest
	GetTaskName() *string
}

type DescribeOssV2ResultRequest struct {
	// example:
	//
	// buckect_test
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 2024-09-14 16:08:38
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// example:
	//
	// 2024-09-14 16:08:38
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	TaskName  *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DescribeOssV2ResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssV2ResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeOssV2ResultRequest) GetBucket() *string {
	return s.Bucket
}

func (s *DescribeOssV2ResultRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeOssV2ResultRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeOssV2ResultRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeOssV2ResultRequest) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeOssV2ResultRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeOssV2ResultRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeOssV2ResultRequest) SetBucket(v string) *DescribeOssV2ResultRequest {
	s.Bucket = &v
	return s
}

func (s *DescribeOssV2ResultRequest) SetCurrentPage(v int32) *DescribeOssV2ResultRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOssV2ResultRequest) SetEndDate(v string) *DescribeOssV2ResultRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeOssV2ResultRequest) SetPageSize(v int32) *DescribeOssV2ResultRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOssV2ResultRequest) SetRiskLevel(v string) *DescribeOssV2ResultRequest {
	s.RiskLevel = &v
	return s
}

func (s *DescribeOssV2ResultRequest) SetStartDate(v string) *DescribeOssV2ResultRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeOssV2ResultRequest) SetTaskName(v string) *DescribeOssV2ResultRequest {
	s.TaskName = &v
	return s
}

func (s *DescribeOssV2ResultRequest) Validate() error {
	return dara.Validate(s)
}
