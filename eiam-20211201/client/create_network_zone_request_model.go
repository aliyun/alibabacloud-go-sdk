// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkZoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateNetworkZoneRequest
	GetClientToken() *string
	SetDescription(v string) *CreateNetworkZoneRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateNetworkZoneRequest
	GetInstanceId() *string
	SetIpv4Cidrs(v []*string) *CreateNetworkZoneRequest
	GetIpv4Cidrs() []*string
	SetIpv6Cidrs(v []*string) *CreateNetworkZoneRequest
	GetIpv6Cidrs() []*string
	SetNetworkZoneName(v string) *CreateNetworkZoneRequest
	GetNetworkZoneName() *string
	SetNetworkZoneType(v string) *CreateNetworkZoneRequest
	GetNetworkZoneType() *string
	SetVpcId(v string) *CreateNetworkZoneRequest
	GetVpcId() *string
}

type CreateNetworkZoneRequest struct {
	// 保证请求幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符，且不能超过64个字符。
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 网络区域描述
	//
	// example:
	//
	// 测试描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 网络区域ipv4Cidr
	//
	// example:
	//
	// 0.0.0.0/0
	Ipv4Cidrs []*string `json:"Ipv4Cidrs,omitempty" xml:"Ipv4Cidrs,omitempty" type:"Repeated"`
	// 网络区域ipv6Cidr
	//
	// example:
	//
	// ::/0
	Ipv6Cidrs []*string `json:"Ipv6Cidrs,omitempty" xml:"Ipv6Cidrs,omitempty" type:"Repeated"`
	// 网络区域名称
	//
	// This parameter is required.
	//
	// example:
	//
	// IPV4Test
	NetworkZoneName *string `json:"NetworkZoneName,omitempty" xml:"NetworkZoneName,omitempty"`
	// 网络区域类型
	//
	// This parameter is required.
	//
	// example:
	//
	// arn:alibaba:idaas:network:zone:classic
	NetworkZoneType *string `json:"NetworkZoneType,omitempty" xml:"NetworkZoneType,omitempty"`
	// 专有网络VpcId
	//
	// example:
	//
	// vpc_xxxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateNetworkZoneRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkZoneRequest) GoString() string {
	return s.String()
}

func (s *CreateNetworkZoneRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateNetworkZoneRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateNetworkZoneRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateNetworkZoneRequest) GetIpv4Cidrs() []*string {
	return s.Ipv4Cidrs
}

func (s *CreateNetworkZoneRequest) GetIpv6Cidrs() []*string {
	return s.Ipv6Cidrs
}

func (s *CreateNetworkZoneRequest) GetNetworkZoneName() *string {
	return s.NetworkZoneName
}

func (s *CreateNetworkZoneRequest) GetNetworkZoneType() *string {
	return s.NetworkZoneType
}

func (s *CreateNetworkZoneRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateNetworkZoneRequest) SetClientToken(v string) *CreateNetworkZoneRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateNetworkZoneRequest) SetDescription(v string) *CreateNetworkZoneRequest {
	s.Description = &v
	return s
}

func (s *CreateNetworkZoneRequest) SetInstanceId(v string) *CreateNetworkZoneRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateNetworkZoneRequest) SetIpv4Cidrs(v []*string) *CreateNetworkZoneRequest {
	s.Ipv4Cidrs = v
	return s
}

func (s *CreateNetworkZoneRequest) SetIpv6Cidrs(v []*string) *CreateNetworkZoneRequest {
	s.Ipv6Cidrs = v
	return s
}

func (s *CreateNetworkZoneRequest) SetNetworkZoneName(v string) *CreateNetworkZoneRequest {
	s.NetworkZoneName = &v
	return s
}

func (s *CreateNetworkZoneRequest) SetNetworkZoneType(v string) *CreateNetworkZoneRequest {
	s.NetworkZoneType = &v
	return s
}

func (s *CreateNetworkZoneRequest) SetVpcId(v string) *CreateNetworkZoneRequest {
	s.VpcId = &v
	return s
}

func (s *CreateNetworkZoneRequest) Validate() error {
	return dara.Validate(s)
}
