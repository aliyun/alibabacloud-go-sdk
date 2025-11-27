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
	// The validity period of the temporary kubeconfig file. Unit: minutes. Valid values: 15 to 4320.
	//
	// >  If you do not specify this parameter, the system specifies a longer validity period. The validity period is returned in the `expiration` parameter.
	//
	// example:
	//
	// 20
	TemporaryDurationMinutes *int32 `json:"TemporaryDurationMinutes,omitempty" xml:"TemporaryDurationMinutes,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// >  This is a reserved parameter.
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
