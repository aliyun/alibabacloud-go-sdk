// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCInstanceDdosCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDdosRegionId(v string) *DescribeRCInstanceDdosCountRequest
	GetDdosRegionId() *string
	SetInstanceType(v string) *DescribeRCInstanceDdosCountRequest
	GetInstanceType() *string
	SetRegionId(v string) *DescribeRCInstanceDdosCountRequest
	GetRegionId() *string
}

type DescribeRCInstanceDdosCountRequest struct {
	// The region ID of the assets that are assigned public IP addresses to query.
	//
	// example:
	//
	// cn-beijing
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	// The instance type of the assets that are assigned public IP addresses to query. Set the value to **ecs**.
	//
	// example:
	//
	// ecs
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The region ID of the RDS Custom instance.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRCInstanceDdosCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceDdosCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceDdosCountRequest) GetDdosRegionId() *string {
	return s.DdosRegionId
}

func (s *DescribeRCInstanceDdosCountRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeRCInstanceDdosCountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCInstanceDdosCountRequest) SetDdosRegionId(v string) *DescribeRCInstanceDdosCountRequest {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeRCInstanceDdosCountRequest) SetInstanceType(v string) *DescribeRCInstanceDdosCountRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeRCInstanceDdosCountRequest) SetRegionId(v string) *DescribeRCInstanceDdosCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRCInstanceDdosCountRequest) Validate() error {
	return dara.Validate(s)
}
