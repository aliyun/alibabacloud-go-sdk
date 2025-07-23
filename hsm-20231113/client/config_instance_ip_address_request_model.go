// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigInstanceIpAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ConfigInstanceIpAddressRequest
	GetInstanceId() *string
	SetIp(v string) *ConfigInstanceIpAddressRequest
	GetIp() *string
	SetRegionId(v string) *ConfigInstanceIpAddressRequest
	GetRegionId() *string
	SetVSwitchId(v string) *ConfigInstanceIpAddressRequest
	GetVSwitchId() *string
	SetVpcId(v string) *ConfigInstanceIpAddressRequest
	GetVpcId() *string
}

type ConfigInstanceIpAddressRequest struct {
	// The ID of the HSM.
	//
	// This parameter is required.
	//
	// example:
	//
	// hsm-cn-mp90fxef****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The endpoint of the VPC to which the HMS belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the vSwitch to which the HMS belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-u7gb0qahu****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC to which the HMS belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-lmkmivmo6****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ConfigInstanceIpAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigInstanceIpAddressRequest) GoString() string {
	return s.String()
}

func (s *ConfigInstanceIpAddressRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ConfigInstanceIpAddressRequest) GetIp() *string {
	return s.Ip
}

func (s *ConfigInstanceIpAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConfigInstanceIpAddressRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ConfigInstanceIpAddressRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ConfigInstanceIpAddressRequest) SetInstanceId(v string) *ConfigInstanceIpAddressRequest {
	s.InstanceId = &v
	return s
}

func (s *ConfigInstanceIpAddressRequest) SetIp(v string) *ConfigInstanceIpAddressRequest {
	s.Ip = &v
	return s
}

func (s *ConfigInstanceIpAddressRequest) SetRegionId(v string) *ConfigInstanceIpAddressRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigInstanceIpAddressRequest) SetVSwitchId(v string) *ConfigInstanceIpAddressRequest {
	s.VSwitchId = &v
	return s
}

func (s *ConfigInstanceIpAddressRequest) SetVpcId(v string) *ConfigInstanceIpAddressRequest {
	s.VpcId = &v
	return s
}

func (s *ConfigInstanceIpAddressRequest) Validate() error {
	return dara.Validate(s)
}
