// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataProcessJobAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *DescribeDataProcessJobAttributeRequest
	GetJobId() *string
	SetRegionId(v string) *DescribeDataProcessJobAttributeRequest
	GetRegionId() *string
}

type DescribeDataProcessJobAttributeRequest struct {
	// The ID of the data processing job to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// job-123456
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the region where the job is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDataProcessJobAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataProcessJobAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataProcessJobAttributeRequest) GetJobId() *string {
	return s.JobId
}

func (s *DescribeDataProcessJobAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDataProcessJobAttributeRequest) SetJobId(v string) *DescribeDataProcessJobAttributeRequest {
	s.JobId = &v
	return s
}

func (s *DescribeDataProcessJobAttributeRequest) SetRegionId(v string) *DescribeDataProcessJobAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDataProcessJobAttributeRequest) Validate() error {
	return dara.Validate(s)
}
