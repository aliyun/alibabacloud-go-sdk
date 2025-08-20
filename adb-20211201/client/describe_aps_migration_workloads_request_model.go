// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApsMigrationWorkloadsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeApsMigrationWorkloadsRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeApsMigrationWorkloadsRequest
	GetEndTime() *string
	SetOssLocation(v string) *DescribeApsMigrationWorkloadsRequest
	GetOssLocation() *string
	SetPageNumber(v int32) *DescribeApsMigrationWorkloadsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApsMigrationWorkloadsRequest
	GetPageSize() *int32
	SetStartTime(v string) *DescribeApsMigrationWorkloadsRequest
	GetStartTime() *string
	SetWorkloadName(v string) *DescribeApsMigrationWorkloadsRequest
	GetWorkloadName() *string
}

type DescribeApsMigrationWorkloadsRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// amv-*******
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query.
	//
	// example:
	//
	// 2021-07-20T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The Object Storage Service (OSS) URL.
	//
	// example:
	//
	// oss://******
	OssLocation *string `json:"OssLocation,omitempty" xml:"OssLocation,omitempty"`
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
	// The start of the time range to query.
	//
	// example:
	//
	// 2021-06-20T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the workload.
	//
	// example:
	//
	// test
	WorkloadName *string `json:"WorkloadName,omitempty" xml:"WorkloadName,omitempty"`
}

func (s DescribeApsMigrationWorkloadsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsMigrationWorkloadsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApsMigrationWorkloadsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeApsMigrationWorkloadsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeApsMigrationWorkloadsRequest) GetOssLocation() *string {
	return s.OssLocation
}

func (s *DescribeApsMigrationWorkloadsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApsMigrationWorkloadsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApsMigrationWorkloadsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeApsMigrationWorkloadsRequest) GetWorkloadName() *string {
	return s.WorkloadName
}

func (s *DescribeApsMigrationWorkloadsRequest) SetDBClusterId(v string) *DescribeApsMigrationWorkloadsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeApsMigrationWorkloadsRequest) SetEndTime(v string) *DescribeApsMigrationWorkloadsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeApsMigrationWorkloadsRequest) SetOssLocation(v string) *DescribeApsMigrationWorkloadsRequest {
	s.OssLocation = &v
	return s
}

func (s *DescribeApsMigrationWorkloadsRequest) SetPageNumber(v int32) *DescribeApsMigrationWorkloadsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApsMigrationWorkloadsRequest) SetPageSize(v int32) *DescribeApsMigrationWorkloadsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApsMigrationWorkloadsRequest) SetStartTime(v string) *DescribeApsMigrationWorkloadsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeApsMigrationWorkloadsRequest) SetWorkloadName(v string) *DescribeApsMigrationWorkloadsRequest {
	s.WorkloadName = &v
	return s
}

func (s *DescribeApsMigrationWorkloadsRequest) Validate() error {
	return dara.Validate(s)
}
