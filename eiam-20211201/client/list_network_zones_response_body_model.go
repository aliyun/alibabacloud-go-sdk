// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworkZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkZones(v []*ListNetworkZonesResponseBodyNetworkZones) *ListNetworkZonesResponseBody
	GetNetworkZones() []*ListNetworkZonesResponseBodyNetworkZones
	SetNextToken(v string) *ListNetworkZonesResponseBody
	GetNextToken() *string
	SetPreviousToken(v string) *ListNetworkZonesResponseBody
	GetPreviousToken() *string
	SetRequestId(v string) *ListNetworkZonesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListNetworkZonesResponseBody
	GetTotalCount() *int64
}

type ListNetworkZonesResponseBody struct {
	NetworkZones []*ListNetworkZonesResponseBodyNetworkZones `json:"NetworkZones,omitempty" xml:"NetworkZones,omitempty" type:"Repeated"`
	// 本次调用返回的查询凭证（Token）值，用于下一次翻页查询。
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// PTxxxxxexample
	PreviousToken *string `json:"PreviousToken,omitempty" xml:"PreviousToken,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNetworkZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkZonesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNetworkZonesResponseBody) GetNetworkZones() []*ListNetworkZonesResponseBodyNetworkZones {
	return s.NetworkZones
}

func (s *ListNetworkZonesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNetworkZonesResponseBody) GetPreviousToken() *string {
	return s.PreviousToken
}

func (s *ListNetworkZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNetworkZonesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListNetworkZonesResponseBody) SetNetworkZones(v []*ListNetworkZonesResponseBodyNetworkZones) *ListNetworkZonesResponseBody {
	s.NetworkZones = v
	return s
}

func (s *ListNetworkZonesResponseBody) SetNextToken(v string) *ListNetworkZonesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNetworkZonesResponseBody) SetPreviousToken(v string) *ListNetworkZonesResponseBody {
	s.PreviousToken = &v
	return s
}

func (s *ListNetworkZonesResponseBody) SetRequestId(v string) *ListNetworkZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNetworkZonesResponseBody) SetTotalCount(v int64) *ListNetworkZonesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListNetworkZonesResponseBody) Validate() error {
	if s.NetworkZones != nil {
		for _, item := range s.NetworkZones {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNetworkZonesResponseBodyNetworkZones struct {
	// IDaaS EIAM 网络区域描述
	//
	// example:
	//
	// test
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
	// network_m223wbvc3sn3uakfnxvhbxxxxx
	NetworkZoneId *string `json:"NetworkZoneId,omitempty" xml:"NetworkZoneId,omitempty"`
	// IDaaS EIAM 网络区域名称
	//
	// example:
	//
	// test_name
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
	// vpc-bp1usdmfqcgoy5ebxxxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListNetworkZonesResponseBodyNetworkZones) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkZonesResponseBodyNetworkZones) GoString() string {
	return s.String()
}

func (s *ListNetworkZonesResponseBodyNetworkZones) GetDescription() *string {
	return s.Description
}

func (s *ListNetworkZonesResponseBodyNetworkZones) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListNetworkZonesResponseBodyNetworkZones) GetIpv4Cidrs() []*string {
	return s.Ipv4Cidrs
}

func (s *ListNetworkZonesResponseBodyNetworkZones) GetIpv6Cidrs() []*string {
	return s.Ipv6Cidrs
}

func (s *ListNetworkZonesResponseBodyNetworkZones) GetNetworkZoneId() *string {
	return s.NetworkZoneId
}

func (s *ListNetworkZonesResponseBodyNetworkZones) GetNetworkZoneName() *string {
	return s.NetworkZoneName
}

func (s *ListNetworkZonesResponseBodyNetworkZones) GetNetworkZoneType() *string {
	return s.NetworkZoneType
}

func (s *ListNetworkZonesResponseBodyNetworkZones) GetVpcId() *string {
	return s.VpcId
}

func (s *ListNetworkZonesResponseBodyNetworkZones) SetDescription(v string) *ListNetworkZonesResponseBodyNetworkZones {
	s.Description = &v
	return s
}

func (s *ListNetworkZonesResponseBodyNetworkZones) SetInstanceId(v string) *ListNetworkZonesResponseBodyNetworkZones {
	s.InstanceId = &v
	return s
}

func (s *ListNetworkZonesResponseBodyNetworkZones) SetIpv4Cidrs(v []*string) *ListNetworkZonesResponseBodyNetworkZones {
	s.Ipv4Cidrs = v
	return s
}

func (s *ListNetworkZonesResponseBodyNetworkZones) SetIpv6Cidrs(v []*string) *ListNetworkZonesResponseBodyNetworkZones {
	s.Ipv6Cidrs = v
	return s
}

func (s *ListNetworkZonesResponseBodyNetworkZones) SetNetworkZoneId(v string) *ListNetworkZonesResponseBodyNetworkZones {
	s.NetworkZoneId = &v
	return s
}

func (s *ListNetworkZonesResponseBodyNetworkZones) SetNetworkZoneName(v string) *ListNetworkZonesResponseBodyNetworkZones {
	s.NetworkZoneName = &v
	return s
}

func (s *ListNetworkZonesResponseBodyNetworkZones) SetNetworkZoneType(v string) *ListNetworkZonesResponseBodyNetworkZones {
	s.NetworkZoneType = &v
	return s
}

func (s *ListNetworkZonesResponseBodyNetworkZones) SetVpcId(v string) *ListNetworkZonesResponseBodyNetworkZones {
	s.VpcId = &v
	return s
}

func (s *ListNetworkZonesResponseBodyNetworkZones) Validate() error {
	return dara.Validate(s)
}
