// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCClusterConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeRCClusterConfigRequest
	GetRegionId() *string
	SetTemporaryDurationMinutes(v int32) *DescribeRCClusterConfigRequest
	GetTemporaryDurationMinutes() *int32
	SetVpcId(v string) *DescribeRCClusterConfigRequest
	GetVpcId() *string
}

type DescribeRCClusterConfigRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The validity period of the temporary KubeConfig. Unit: minutes. Valid values: 15 (15 minutes) to 4320 (3 days).
	//
	// > If this parameter is not specified, the system automatically determines a longer validity period. The specific expiration time is indicated by the value of the `expiration` field in the response.
	//
	// example:
	//
	// 20
	TemporaryDurationMinutes *int32 `json:"TemporaryDurationMinutes,omitempty" xml:"TemporaryDurationMinutes,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// > Reserved parameter.
	//
	// example:
	//
	// None
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeRCClusterConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCClusterConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeRCClusterConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCClusterConfigRequest) GetTemporaryDurationMinutes() *int32 {
	return s.TemporaryDurationMinutes
}

func (s *DescribeRCClusterConfigRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeRCClusterConfigRequest) SetRegionId(v string) *DescribeRCClusterConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRCClusterConfigRequest) SetTemporaryDurationMinutes(v int32) *DescribeRCClusterConfigRequest {
	s.TemporaryDurationMinutes = &v
	return s
}

func (s *DescribeRCClusterConfigRequest) SetVpcId(v string) *DescribeRCClusterConfigRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeRCClusterConfigRequest) Validate() error {
	return dara.Validate(s)
}
