// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStreamingJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ListStreamingJobsRequest
	GetDBInstanceId() *string
	SetPageNumber(v int32) *ListStreamingJobsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListStreamingJobsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListStreamingJobsRequest
	GetRegionId() *string
}

type ListStreamingJobsRequest struct {
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp10g78o9807yv9h3
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Page number, greater than 0 and not exceeding the maximum value of Integer. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of records per page, with the following options:
	//
	// - **30*	- - **50*	- - **100**
	//
	// Default value: 30.
	//
	// example:
	//
	// kafka
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListStreamingJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStreamingJobsRequest) GoString() string {
	return s.String()
}

func (s *ListStreamingJobsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListStreamingJobsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListStreamingJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListStreamingJobsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListStreamingJobsRequest) SetDBInstanceId(v string) *ListStreamingJobsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ListStreamingJobsRequest) SetPageNumber(v int32) *ListStreamingJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStreamingJobsRequest) SetPageSize(v int32) *ListStreamingJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListStreamingJobsRequest) SetRegionId(v string) *ListStreamingJobsRequest {
	s.RegionId = &v
	return s
}

func (s *ListStreamingJobsRequest) Validate() error {
	return dara.Validate(s)
}
