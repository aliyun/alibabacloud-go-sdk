// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApsProgressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeApsProgressRequest
	GetDBClusterId() *string
	SetRegionId(v string) *DescribeApsProgressRequest
	GetRegionId() *string
	SetWorkloadId(v string) *DescribeApsProgressRequest
	GetWorkloadId() *string
}

type DescribeApsProgressRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-******
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// aps-******
	WorkloadId *string `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
}

func (s DescribeApsProgressRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsProgressRequest) GoString() string {
	return s.String()
}

func (s *DescribeApsProgressRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeApsProgressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApsProgressRequest) GetWorkloadId() *string {
	return s.WorkloadId
}

func (s *DescribeApsProgressRequest) SetDBClusterId(v string) *DescribeApsProgressRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeApsProgressRequest) SetRegionId(v string) *DescribeApsProgressRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApsProgressRequest) SetWorkloadId(v string) *DescribeApsProgressRequest {
	s.WorkloadId = &v
	return s
}

func (s *DescribeApsProgressRequest) Validate() error {
	return dara.Validate(s)
}
