// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStreamingJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeStreamingJobRequest
	GetDBInstanceId() *string
	SetJobId(v int32) *DescribeStreamingJobRequest
	GetJobId() *int32
	SetRegionId(v string) *DescribeStreamingJobRequest
	GetRegionId() *string
}

type DescribeStreamingJobRequest struct {
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp10g78o9807yv9h3
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	JobId *int32 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) interface to view available region IDs.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeStreamingJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStreamingJobRequest) GoString() string {
	return s.String()
}

func (s *DescribeStreamingJobRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeStreamingJobRequest) GetJobId() *int32 {
	return s.JobId
}

func (s *DescribeStreamingJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeStreamingJobRequest) SetDBInstanceId(v string) *DescribeStreamingJobRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeStreamingJobRequest) SetJobId(v int32) *DescribeStreamingJobRequest {
	s.JobId = &v
	return s
}

func (s *DescribeStreamingJobRequest) SetRegionId(v string) *DescribeStreamingJobRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeStreamingJobRequest) Validate() error {
	return dara.Validate(s)
}
