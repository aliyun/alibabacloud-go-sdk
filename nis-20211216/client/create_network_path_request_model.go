// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkPathRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkPathDescription(v string) *CreateNetworkPathRequest
	GetNetworkPathDescription() *string
	SetNetworkPathName(v string) *CreateNetworkPathRequest
	GetNetworkPathName() *string
	SetProtocol(v string) *CreateNetworkPathRequest
	GetProtocol() *string
	SetRegionId(v string) *CreateNetworkPathRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateNetworkPathRequest
	GetResourceGroupId() *string
	SetSourceId(v string) *CreateNetworkPathRequest
	GetSourceId() *string
	SetSourceIpAddress(v string) *CreateNetworkPathRequest
	GetSourceIpAddress() *string
	SetSourcePort(v int32) *CreateNetworkPathRequest
	GetSourcePort() *int32
	SetSourceType(v string) *CreateNetworkPathRequest
	GetSourceType() *string
	SetTag(v []*CreateNetworkPathRequestTag) *CreateNetworkPathRequest
	GetTag() []*CreateNetworkPathRequestTag
	SetTargetId(v string) *CreateNetworkPathRequest
	GetTargetId() *string
	SetTargetIpAddress(v string) *CreateNetworkPathRequest
	GetTargetIpAddress() *string
	SetTargetPort(v int32) *CreateNetworkPathRequest
	GetTargetPort() *int32
	SetTargetType(v string) *CreateNetworkPathRequest
	GetTargetType() *string
}

type CreateNetworkPathRequest struct {
	// The description of the network path.
	//
	// example:
	//
	// Analyze the path from ECS to ECS
	NetworkPathDescription *string `json:"NetworkPathDescription,omitempty" xml:"NetworkPathDescription,omitempty"`
	// The name of the network path.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs2PublicIp
	NetworkPathName *string `json:"NetworkPathName,omitempty" xml:"NetworkPathName,omitempty"`
	// The protocol type. Valid values:
	//
	// 	- **tcp**: Transmission Control Protocol (TCP)
	//
	// 	- **udp**: User Datagram Protocol (UDP)
	//
	// 	- **icmp**: Internet Control Message Protocol (ICMP)
	//
	// example:
	//
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The region ID of the network path that you want to create.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfm27qsxjj****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the source resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-2zef4ngqfarepyun****
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 172.17.XX.XX
	SourceIpAddress *string `json:"SourceIpAddress,omitempty" xml:"SourceIpAddress,omitempty"`
	// The source port.
	//
	// example:
	//
	// 443
	SourcePort *int32 `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	// The type of the source resource. Valid values:
	//
	// 	- **ecs**: the Elastic Compute Service (ECS) instance
	//
	// 	- **internetIp**: the public IP address
	//
	// 	- **vsw**: the vSwitch
	//
	// 	- **vpn**: the VPN gateway
	//
	// 	- **vbr**: the virtual border router (VBR)
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The tags to add to the resource.
	Tag []*CreateNetworkPathRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the destination resource.
	//
	// example:
	//
	// i-bp13d0e064gubm****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The destination IP address.
	//
	// example:
	//
	// 192.168.0.210
	TargetIpAddress *string `json:"TargetIpAddress,omitempty" xml:"TargetIpAddress,omitempty"`
	// The destination port.
	//
	// example:
	//
	// 80
	TargetPort *int32 `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
	// The type of the destination resource. Valid values:
	//
	// 	- **ecs**: the ECS instance
	//
	// 	- **internetIp**: the public IP address
	//
	// 	- **vsw**: the vSwitch
	//
	// 	- **vpn**: the VPN gateway
	//
	// 	- **vbr**: the VBR
	//
	// 	- **clb**: the Classic Load Balancer (CLB) instance
	//
	// example:
	//
	// ecs
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s CreateNetworkPathRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkPathRequest) GoString() string {
	return s.String()
}

func (s *CreateNetworkPathRequest) GetNetworkPathDescription() *string {
	return s.NetworkPathDescription
}

func (s *CreateNetworkPathRequest) GetNetworkPathName() *string {
	return s.NetworkPathName
}

func (s *CreateNetworkPathRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateNetworkPathRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateNetworkPathRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateNetworkPathRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateNetworkPathRequest) GetSourceIpAddress() *string {
	return s.SourceIpAddress
}

func (s *CreateNetworkPathRequest) GetSourcePort() *int32 {
	return s.SourcePort
}

func (s *CreateNetworkPathRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateNetworkPathRequest) GetTag() []*CreateNetworkPathRequestTag {
	return s.Tag
}

func (s *CreateNetworkPathRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *CreateNetworkPathRequest) GetTargetIpAddress() *string {
	return s.TargetIpAddress
}

func (s *CreateNetworkPathRequest) GetTargetPort() *int32 {
	return s.TargetPort
}

func (s *CreateNetworkPathRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateNetworkPathRequest) SetNetworkPathDescription(v string) *CreateNetworkPathRequest {
	s.NetworkPathDescription = &v
	return s
}

func (s *CreateNetworkPathRequest) SetNetworkPathName(v string) *CreateNetworkPathRequest {
	s.NetworkPathName = &v
	return s
}

func (s *CreateNetworkPathRequest) SetProtocol(v string) *CreateNetworkPathRequest {
	s.Protocol = &v
	return s
}

func (s *CreateNetworkPathRequest) SetRegionId(v string) *CreateNetworkPathRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNetworkPathRequest) SetResourceGroupId(v string) *CreateNetworkPathRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateNetworkPathRequest) SetSourceId(v string) *CreateNetworkPathRequest {
	s.SourceId = &v
	return s
}

func (s *CreateNetworkPathRequest) SetSourceIpAddress(v string) *CreateNetworkPathRequest {
	s.SourceIpAddress = &v
	return s
}

func (s *CreateNetworkPathRequest) SetSourcePort(v int32) *CreateNetworkPathRequest {
	s.SourcePort = &v
	return s
}

func (s *CreateNetworkPathRequest) SetSourceType(v string) *CreateNetworkPathRequest {
	s.SourceType = &v
	return s
}

func (s *CreateNetworkPathRequest) SetTag(v []*CreateNetworkPathRequestTag) *CreateNetworkPathRequest {
	s.Tag = v
	return s
}

func (s *CreateNetworkPathRequest) SetTargetId(v string) *CreateNetworkPathRequest {
	s.TargetId = &v
	return s
}

func (s *CreateNetworkPathRequest) SetTargetIpAddress(v string) *CreateNetworkPathRequest {
	s.TargetIpAddress = &v
	return s
}

func (s *CreateNetworkPathRequest) SetTargetPort(v int32) *CreateNetworkPathRequest {
	s.TargetPort = &v
	return s
}

func (s *CreateNetworkPathRequest) SetTargetType(v string) *CreateNetworkPathRequest {
	s.TargetType = &v
	return s
}

func (s *CreateNetworkPathRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateNetworkPathRequestTag struct {
	// The key of tag N to add to the resource. The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	//
	// You can add up to 20 tags in each call.
	//
	// example:
	//
	// role
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// ops
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateNetworkPathRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkPathRequestTag) GoString() string {
	return s.String()
}

func (s *CreateNetworkPathRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateNetworkPathRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateNetworkPathRequestTag) SetKey(v string) *CreateNetworkPathRequestTag {
	s.Key = &v
	return s
}

func (s *CreateNetworkPathRequestTag) SetValue(v string) *CreateNetworkPathRequestTag {
	s.Value = &v
	return s
}

func (s *CreateNetworkPathRequestTag) Validate() error {
	return dara.Validate(s)
}
