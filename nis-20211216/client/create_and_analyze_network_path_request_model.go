// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAndAnalyzeNetworkPathRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProtocol(v string) *CreateAndAnalyzeNetworkPathRequest
	GetProtocol() *string
	SetRegionId(v string) *CreateAndAnalyzeNetworkPathRequest
	GetRegionId() *string
	SetSourceId(v string) *CreateAndAnalyzeNetworkPathRequest
	GetSourceId() *string
	SetSourceIpAddress(v string) *CreateAndAnalyzeNetworkPathRequest
	GetSourceIpAddress() *string
	SetSourcePort(v int32) *CreateAndAnalyzeNetworkPathRequest
	GetSourcePort() *int32
	SetSourceType(v string) *CreateAndAnalyzeNetworkPathRequest
	GetSourceType() *string
	SetTargetId(v string) *CreateAndAnalyzeNetworkPathRequest
	GetTargetId() *string
	SetTargetIpAddress(v string) *CreateAndAnalyzeNetworkPathRequest
	GetTargetIpAddress() *string
	SetTargetPort(v int32) *CreateAndAnalyzeNetworkPathRequest
	GetTargetPort() *int32
	SetTargetType(v string) *CreateAndAnalyzeNetworkPathRequest
	GetTargetType() *string
}

type CreateAndAnalyzeNetworkPathRequest struct {
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
	// The ID of the region for which you want to initiate a task for analyzing network reachability.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the source resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-uf62y8khhbkbdrp6****
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 192.168.XX.XX
	SourceIpAddress *string `json:"SourceIpAddress,omitempty" xml:"SourceIpAddress,omitempty"`
	// The source port.
	//
	// example:
	//
	// 0
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
	// The ID of the destination resource.
	//
	// example:
	//
	// i-m5eactvw7wtpktv5****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The destination IP address.
	//
	// example:
	//
	// 172.50.XX.XX
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

func (s CreateAndAnalyzeNetworkPathRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAndAnalyzeNetworkPathRequest) GoString() string {
	return s.String()
}

func (s *CreateAndAnalyzeNetworkPathRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateAndAnalyzeNetworkPathRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAndAnalyzeNetworkPathRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateAndAnalyzeNetworkPathRequest) GetSourceIpAddress() *string {
	return s.SourceIpAddress
}

func (s *CreateAndAnalyzeNetworkPathRequest) GetSourcePort() *int32 {
	return s.SourcePort
}

func (s *CreateAndAnalyzeNetworkPathRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateAndAnalyzeNetworkPathRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *CreateAndAnalyzeNetworkPathRequest) GetTargetIpAddress() *string {
	return s.TargetIpAddress
}

func (s *CreateAndAnalyzeNetworkPathRequest) GetTargetPort() *int32 {
	return s.TargetPort
}

func (s *CreateAndAnalyzeNetworkPathRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateAndAnalyzeNetworkPathRequest) SetProtocol(v string) *CreateAndAnalyzeNetworkPathRequest {
	s.Protocol = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathRequest) SetRegionId(v string) *CreateAndAnalyzeNetworkPathRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathRequest) SetSourceId(v string) *CreateAndAnalyzeNetworkPathRequest {
	s.SourceId = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathRequest) SetSourceIpAddress(v string) *CreateAndAnalyzeNetworkPathRequest {
	s.SourceIpAddress = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathRequest) SetSourcePort(v int32) *CreateAndAnalyzeNetworkPathRequest {
	s.SourcePort = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathRequest) SetSourceType(v string) *CreateAndAnalyzeNetworkPathRequest {
	s.SourceType = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathRequest) SetTargetId(v string) *CreateAndAnalyzeNetworkPathRequest {
	s.TargetId = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathRequest) SetTargetIpAddress(v string) *CreateAndAnalyzeNetworkPathRequest {
	s.TargetIpAddress = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathRequest) SetTargetPort(v int32) *CreateAndAnalyzeNetworkPathRequest {
	s.TargetPort = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathRequest) SetTargetType(v string) *CreateAndAnalyzeNetworkPathRequest {
	s.TargetType = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathRequest) Validate() error {
	return dara.Validate(s)
}
