// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataProcessJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *DescribeDataProcessJobsRequest
	GetJobId() *string
	SetPageNumber(v int32) *DescribeDataProcessJobsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDataProcessJobsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDataProcessJobsRequest
	GetRegionId() *string
	SetSourceId(v string) *DescribeDataProcessJobsRequest
	GetSourceId() *string
}

type DescribeDataProcessJobsRequest struct {
	// The ID of the data processing job used for filtering.
	//
	// example:
	//
	// job-123456
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the job resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the source PolarFS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pfs-123456
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
}

func (s DescribeDataProcessJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataProcessJobsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataProcessJobsRequest) GetJobId() *string {
	return s.JobId
}

func (s *DescribeDataProcessJobsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDataProcessJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDataProcessJobsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDataProcessJobsRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *DescribeDataProcessJobsRequest) SetJobId(v string) *DescribeDataProcessJobsRequest {
	s.JobId = &v
	return s
}

func (s *DescribeDataProcessJobsRequest) SetPageNumber(v int32) *DescribeDataProcessJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataProcessJobsRequest) SetPageSize(v int32) *DescribeDataProcessJobsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataProcessJobsRequest) SetRegionId(v string) *DescribeDataProcessJobsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDataProcessJobsRequest) SetSourceId(v string) *DescribeDataProcessJobsRequest {
	s.SourceId = &v
	return s
}

func (s *DescribeDataProcessJobsRequest) Validate() error {
	return dara.Validate(s)
}
