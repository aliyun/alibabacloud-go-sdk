// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApsHiveWorkloadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeApsHiveWorkloadRequest
	GetDBClusterId() *string
	SetRegionId(v string) *DescribeApsHiveWorkloadRequest
	GetRegionId() *string
	SetWorkloadId(v string) *DescribeApsHiveWorkloadRequest
	GetWorkloadId() *string
}

type DescribeApsHiveWorkloadRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-*******
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

func (s DescribeApsHiveWorkloadRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsHiveWorkloadRequest) GoString() string {
	return s.String()
}

func (s *DescribeApsHiveWorkloadRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeApsHiveWorkloadRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApsHiveWorkloadRequest) GetWorkloadId() *string {
	return s.WorkloadId
}

func (s *DescribeApsHiveWorkloadRequest) SetDBClusterId(v string) *DescribeApsHiveWorkloadRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeApsHiveWorkloadRequest) SetRegionId(v string) *DescribeApsHiveWorkloadRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApsHiveWorkloadRequest) SetWorkloadId(v string) *DescribeApsHiveWorkloadRequest {
	s.WorkloadId = &v
	return s
}

func (s *DescribeApsHiveWorkloadRequest) Validate() error {
	return dara.Validate(s)
}
