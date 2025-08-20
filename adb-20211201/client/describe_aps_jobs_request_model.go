// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApsJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApsJobName(v string) *DescribeApsJobsRequest
	GetApsJobName() *string
	SetCreateTimeEnd(v string) *DescribeApsJobsRequest
	GetCreateTimeEnd() *string
	SetCreateTimeStart(v string) *DescribeApsJobsRequest
	GetCreateTimeStart() *string
	SetPageNumber(v int32) *DescribeApsJobsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApsJobsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeApsJobsRequest
	GetRegionId() *string
}

type DescribeApsJobsRequest struct {
	// The name of the APS job.
	//
	// example:
	//
	// aps-xxxxx
	ApsJobName *string `json:"ApsJobName,omitempty" xml:"ApsJobName,omitempty"`
	// The end of the time range to query.
	//
	// example:
	//
	// 2022-04-23T01:10Z
	CreateTimeEnd *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// The beginning of the time range to query.
	//
	// example:
	//
	// 2022-03-23T01:10Z
	CreateTimeStart *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeApsJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsJobsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApsJobsRequest) GetApsJobName() *string {
	return s.ApsJobName
}

func (s *DescribeApsJobsRequest) GetCreateTimeEnd() *string {
	return s.CreateTimeEnd
}

func (s *DescribeApsJobsRequest) GetCreateTimeStart() *string {
	return s.CreateTimeStart
}

func (s *DescribeApsJobsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApsJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApsJobsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApsJobsRequest) SetApsJobName(v string) *DescribeApsJobsRequest {
	s.ApsJobName = &v
	return s
}

func (s *DescribeApsJobsRequest) SetCreateTimeEnd(v string) *DescribeApsJobsRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *DescribeApsJobsRequest) SetCreateTimeStart(v string) *DescribeApsJobsRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *DescribeApsJobsRequest) SetPageNumber(v int32) *DescribeApsJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApsJobsRequest) SetPageSize(v int32) *DescribeApsJobsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApsJobsRequest) SetRegionId(v string) *DescribeApsJobsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApsJobsRequest) Validate() error {
	return dara.Validate(s)
}
