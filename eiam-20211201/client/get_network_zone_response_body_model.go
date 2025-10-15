// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetworkZoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkZone(v *GetNetworkZoneResponseBodyNetworkZone) *GetNetworkZoneResponseBody
	GetNetworkZone() *GetNetworkZoneResponseBodyNetworkZone
	SetRequestId(v string) *GetNetworkZoneResponseBody
	GetRequestId() *string
}

type GetNetworkZoneResponseBody struct {
	NetworkZone *GetNetworkZoneResponseBodyNetworkZone `json:"NetworkZone,omitempty" xml:"NetworkZone,omitempty" type:"Struct"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNetworkZoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkZoneResponseBody) GoString() string {
	return s.String()
}

func (s *GetNetworkZoneResponseBody) GetNetworkZone() *GetNetworkZoneResponseBodyNetworkZone {
	return s.NetworkZone
}

func (s *GetNetworkZoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNetworkZoneResponseBody) SetNetworkZone(v *GetNetworkZoneResponseBodyNetworkZone) *GetNetworkZoneResponseBody {
	s.NetworkZone = v
	return s
}

func (s *GetNetworkZoneResponseBody) SetRequestId(v string) *GetNetworkZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNetworkZoneResponseBody) Validate() error {
	if s.NetworkZone != nil {
		if err := s.NetworkZone.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNetworkZoneResponseBodyNetworkZone struct {
	// IDaaS EIAM 网络区域描述
	//
	// example:
	//
	// test_description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 实例ID。
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Ipv4Cidrs  []*string `json:"Ipv4Cidrs,omitempty" xml:"Ipv4Cidrs,omitempty" type:"Repeated"`
	Ipv6Cidrs  []*string `json:"Ipv6Cidrs,omitempty" xml:"Ipv6Cidrs,omitempty" type:"Repeated"`
	// IDaaS EIAM 网络区域Id
	//
	// example:
	//
	// network_m6fbr2bcbcadu3bcdpgzcxxxxx
	NetworkZoneId *string `json:"NetworkZoneId,omitempty" xml:"NetworkZoneId,omitempty"`
	// IDaaS EIAM 网络区域名称
	//
	// example:
	//
	// test4_name
	NetworkZoneName *string `json:"NetworkZoneName,omitempty" xml:"NetworkZoneName,omitempty"`
	// IDaaS EIAM 网络区域类型
	//
	// example:
	//
	// arn:alibaba:idaas:network:zone:classic
	NetworkZoneType *string `json:"NetworkZoneType,omitempty" xml:"NetworkZoneType,omitempty"`
	// IDaaS EIAM 专有网络VpcId
	//
	// example:
	//
	// vpc-25w8wxxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetNetworkZoneResponseBodyNetworkZone) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkZoneResponseBodyNetworkZone) GoString() string {
	return s.String()
}

func (s *GetNetworkZoneResponseBodyNetworkZone) GetDescription() *string {
	return s.Description
}

func (s *GetNetworkZoneResponseBodyNetworkZone) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetNetworkZoneResponseBodyNetworkZone) GetIpv4Cidrs() []*string {
	return s.Ipv4Cidrs
}

func (s *GetNetworkZoneResponseBodyNetworkZone) GetIpv6Cidrs() []*string {
	return s.Ipv6Cidrs
}

func (s *GetNetworkZoneResponseBodyNetworkZone) GetNetworkZoneId() *string {
	return s.NetworkZoneId
}

func (s *GetNetworkZoneResponseBodyNetworkZone) GetNetworkZoneName() *string {
	return s.NetworkZoneName
}

func (s *GetNetworkZoneResponseBodyNetworkZone) GetNetworkZoneType() *string {
	return s.NetworkZoneType
}

func (s *GetNetworkZoneResponseBodyNetworkZone) GetVpcId() *string {
	return s.VpcId
}

func (s *GetNetworkZoneResponseBodyNetworkZone) SetDescription(v string) *GetNetworkZoneResponseBodyNetworkZone {
	s.Description = &v
	return s
}

func (s *GetNetworkZoneResponseBodyNetworkZone) SetInstanceId(v string) *GetNetworkZoneResponseBodyNetworkZone {
	s.InstanceId = &v
	return s
}

func (s *GetNetworkZoneResponseBodyNetworkZone) SetIpv4Cidrs(v []*string) *GetNetworkZoneResponseBodyNetworkZone {
	s.Ipv4Cidrs = v
	return s
}

func (s *GetNetworkZoneResponseBodyNetworkZone) SetIpv6Cidrs(v []*string) *GetNetworkZoneResponseBodyNetworkZone {
	s.Ipv6Cidrs = v
	return s
}

func (s *GetNetworkZoneResponseBodyNetworkZone) SetNetworkZoneId(v string) *GetNetworkZoneResponseBodyNetworkZone {
	s.NetworkZoneId = &v
	return s
}

func (s *GetNetworkZoneResponseBodyNetworkZone) SetNetworkZoneName(v string) *GetNetworkZoneResponseBodyNetworkZone {
	s.NetworkZoneName = &v
	return s
}

func (s *GetNetworkZoneResponseBodyNetworkZone) SetNetworkZoneType(v string) *GetNetworkZoneResponseBodyNetworkZone {
	s.NetworkZoneType = &v
	return s
}

func (s *GetNetworkZoneResponseBodyNetworkZone) SetVpcId(v string) *GetNetworkZoneResponseBodyNetworkZone {
	s.VpcId = &v
	return s
}

func (s *GetNetworkZoneResponseBodyNetworkZone) Validate() error {
	return dara.Validate(s)
}
