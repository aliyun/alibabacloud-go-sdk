// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNisTrafficRankingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFlowRankingList(v []*DescribeNisTrafficRankingResponseBodyFlowRankingList) *DescribeNisTrafficRankingResponseBody
	GetFlowRankingList() []*DescribeNisTrafficRankingResponseBodyFlowRankingList
	SetMaxResults(v int32) *DescribeNisTrafficRankingResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeNisTrafficRankingResponseBody
	GetNextToken() *string
	SetNisTrafficRankingId(v string) *DescribeNisTrafficRankingResponseBody
	GetNisTrafficRankingId() *string
	SetRequestId(v string) *DescribeNisTrafficRankingResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeNisTrafficRankingResponseBody
	GetStatus() *string
	SetTotalCount(v int32) *DescribeNisTrafficRankingResponseBody
	GetTotalCount() *int32
}

type DescribeNisTrafficRankingResponseBody struct {
	FlowRankingList []*DescribeNisTrafficRankingResponseBodyFlowRankingList `json:"FlowRankingList,omitempty" xml:"FlowRankingList,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// LoeJLhK0fsDqYoXkXieZUqB2vWnccJtVnsyKu9KxFFOMQxtV8XckOg5lk7F2bhC+
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// task-7619ecb1db9148bab9f4
	NisTrafficRankingId *string `json:"NisTrafficRankingId,omitempty" xml:"NisTrafficRankingId,omitempty"`
	// example:
	//
	// 4DAC4BE1-BEEA-5D84-BE06-E1B796F3B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Complete
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 72
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeNisTrafficRankingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNisTrafficRankingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNisTrafficRankingResponseBody) GetFlowRankingList() []*DescribeNisTrafficRankingResponseBodyFlowRankingList {
	return s.FlowRankingList
}

func (s *DescribeNisTrafficRankingResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeNisTrafficRankingResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeNisTrafficRankingResponseBody) GetNisTrafficRankingId() *string {
	return s.NisTrafficRankingId
}

func (s *DescribeNisTrafficRankingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNisTrafficRankingResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeNisTrafficRankingResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeNisTrafficRankingResponseBody) SetFlowRankingList(v []*DescribeNisTrafficRankingResponseBodyFlowRankingList) *DescribeNisTrafficRankingResponseBody {
	s.FlowRankingList = v
	return s
}

func (s *DescribeNisTrafficRankingResponseBody) SetMaxResults(v int32) *DescribeNisTrafficRankingResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBody) SetNextToken(v string) *DescribeNisTrafficRankingResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBody) SetNisTrafficRankingId(v string) *DescribeNisTrafficRankingResponseBody {
	s.NisTrafficRankingId = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBody) SetRequestId(v string) *DescribeNisTrafficRankingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBody) SetStatus(v string) *DescribeNisTrafficRankingResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBody) SetTotalCount(v int32) *DescribeNisTrafficRankingResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBody) Validate() error {
	if s.FlowRankingList != nil {
		for _, item := range s.FlowRankingList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNisTrafficRankingResponseBodyFlowRankingList struct {
	// example:
	//
	// ngw-ufwerthgvc*****
	BindingResourceId *string `json:"BindingResourceId,omitempty" xml:"BindingResourceId,omitempty"`
	// example:
	//
	// EIP_NAT
	BindingResourceType *string `json:"BindingResourceType,omitempty" xml:"BindingResourceType,omitempty"`
	// example:
	//
	// 100
	Bytes *float64 `json:"Bytes,omitempty" xml:"Bytes,omitempty"`
	// example:
	//
	// 0.2
	BytesRate *float64 `json:"BytesRate,omitempty" xml:"BytesRate,omitempty"`
	// example:
	//
	// 45102
	ClientAsn      *string `json:"ClientAsn,omitempty" xml:"ClientAsn,omitempty"`
	ClientCity     *string `json:"ClientCity,omitempty" xml:"ClientCity,omitempty"`
	ClientCountry  *string `json:"ClientCountry,omitempty" xml:"ClientCountry,omitempty"`
	ClientIsp      *string `json:"ClientIsp,omitempty" xml:"ClientIsp,omitempty"`
	ClientProvince *string `json:"ClientProvince,omitempty" xml:"ClientProvince,omitempty"`
	// example:
	//
	// 192.168.***.0
	DestinationIp *string `json:"DestinationIp,omitempty" xml:"DestinationIp,omitempty"`
	// example:
	//
	// 23
	DestinationPort *string `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty"`
	// example:
	//
	// cn-hangzhou
	DestinationRegionNo *string `json:"DestinationRegionNo,omitempty" xml:"DestinationRegionNo,omitempty"`
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// example:
	//
	// 0
	Dscp *string `json:"Dscp,omitempty" xml:"Dscp,omitempty"`
	// example:
	//
	// i-uf6i1zi6yhq7h***
	EcsId *string `json:"EcsId,omitempty" xml:"EcsId,omitempty"`
	// example:
	//
	// eip-fb6wzjl9hm****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// eni-8vbf2jxul***
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// example:
	//
	// 100
	Packets *float64 `json:"Packets,omitempty" xml:"Packets,omitempty"`
	// example:
	//
	// 4
	PacketsLostBlackhole *float64 `json:"PacketsLostBlackhole,omitempty" xml:"PacketsLostBlackhole,omitempty"`
	// example:
	//
	// 2
	PacketsLostNoRoute *float64 `json:"PacketsLostNoRoute,omitempty" xml:"PacketsLostNoRoute,omitempty"`
	// example:
	//
	// 7
	PacketsLostTTLExpired *float64 `json:"PacketsLostTTLExpired,omitempty" xml:"PacketsLostTTLExpired,omitempty"`
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// 118.31.***.86
	PublicIpAddress *string `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 2
	RoundTripTime *float64 `json:"RoundTripTime,omitempty" xml:"RoundTripTime,omitempty"`
	// example:
	//
	// 47.92.245.***
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// example:
	//
	// 5432
	SourcePort *string `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	// example:
	//
	// cn-hangzhou
	SourceRegionNo *string `json:"SourceRegionNo,omitempty" xml:"SourceRegionNo,omitempty"`
	// example:
	//
	// all
	TrafficPath *string `json:"TrafficPath,omitempty" xml:"TrafficPath,omitempty"`
	// example:
	//
	// tr-attach-bfde1cd4cj***
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// example:
	//
	// 1906814138****
	TransitRouterDestinationAccountId *string `json:"TransitRouterDestinationAccountId,omitempty" xml:"TransitRouterDestinationAccountId,omitempty"`
	// example:
	//
	// cn-hangzhou-j
	TransitRouterDestinationAvailableZone *string `json:"TransitRouterDestinationAvailableZone,omitempty" xml:"TransitRouterDestinationAvailableZone,omitempty"`
	// example:
	//
	// eni-fdbf2jxulm***
	TransitRouterDestinationNetworkInterface *string `json:"TransitRouterDestinationNetworkInterface,omitempty" xml:"TransitRouterDestinationNetworkInterface,omitempty"`
	// example:
	//
	// tr-attach-bfve1cd4cjp****
	TransitRouterDestinationResourceId *string `json:"TransitRouterDestinationResourceId,omitempty" xml:"TransitRouterDestinationResourceId,omitempty"`
	// example:
	//
	// vsw-2zeekevlhxpqxu****
	TransitRouterDestinationVSwitchId *string `json:"TransitRouterDestinationVSwitchId,omitempty" xml:"TransitRouterDestinationVSwitchId,omitempty"`
	// example:
	//
	// tr-2zefvwy2fz3444***
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// example:
	//
	// tr-attach-okvj1cd4cjp***
	TransitRouterPairAttachmentId *string `json:"TransitRouterPairAttachmentId,omitempty" xml:"TransitRouterPairAttachmentId,omitempty"`
	// example:
	//
	// 1906814138***
	TransitRouterSourceAccountId *string `json:"TransitRouterSourceAccountId,omitempty" xml:"TransitRouterSourceAccountId,omitempty"`
	// example:
	//
	// cn-hangzhou-j
	TransitRouterSourceAvailableZone *string `json:"TransitRouterSourceAvailableZone,omitempty" xml:"TransitRouterSourceAvailableZone,omitempty"`
	// example:
	//
	// eni-8vbf2jxulma***
	TransitRouterSourceNetworkInterface *string `json:"TransitRouterSourceNetworkInterface,omitempty" xml:"TransitRouterSourceNetworkInterface,omitempty"`
	// example:
	//
	// tr-attach-hvve1cd4cjpj***
	TransitRouterSourceResourceId *string `json:"TransitRouterSourceResourceId,omitempty" xml:"TransitRouterSourceResourceId,omitempty"`
	// example:
	//
	// vsw-ikfdkevlhxpqxuz****
	TransitRouterSourceVSwitchId *string `json:"TransitRouterSourceVSwitchId,omitempty" xml:"TransitRouterSourceVSwitchId,omitempty"`
	// example:
	//
	// vsw-2zeekevlh****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// vpc-m5ec6i0h5xss***
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeNisTrafficRankingResponseBodyFlowRankingList) String() string {
	return dara.Prettify(s)
}

func (s DescribeNisTrafficRankingResponseBodyFlowRankingList) GoString() string {
	return s.String()
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetBindingResourceId() *string {
	return s.BindingResourceId
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetBindingResourceType() *string {
	return s.BindingResourceType
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetBytes() *float64 {
	return s.Bytes
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetBytesRate() *float64 {
	return s.BytesRate
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetClientAsn() *string {
	return s.ClientAsn
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetClientCity() *string {
	return s.ClientCity
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetClientCountry() *string {
	return s.ClientCountry
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetClientIsp() *string {
	return s.ClientIsp
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetClientProvince() *string {
	return s.ClientProvince
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetDestinationIp() *string {
	return s.DestinationIp
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetDestinationPort() *string {
	return s.DestinationPort
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetDestinationRegionNo() *string {
	return s.DestinationRegionNo
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetDirection() *string {
	return s.Direction
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetDscp() *string {
	return s.Dscp
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetEcsId() *string {
	return s.EcsId
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetPackets() *float64 {
	return s.Packets
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetPacketsLostBlackhole() *float64 {
	return s.PacketsLostBlackhole
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetPacketsLostNoRoute() *float64 {
	return s.PacketsLostNoRoute
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetPacketsLostTTLExpired() *float64 {
	return s.PacketsLostTTLExpired
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetPublicIpAddress() *string {
	return s.PublicIpAddress
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetRoundTripTime() *float64 {
	return s.RoundTripTime
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetSourcePort() *string {
	return s.SourcePort
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetSourceRegionNo() *string {
	return s.SourceRegionNo
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetTrafficPath() *string {
	return s.TrafficPath
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetTransitRouterDestinationAccountId() *string {
	return s.TransitRouterDestinationAccountId
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetTransitRouterDestinationAvailableZone() *string {
	return s.TransitRouterDestinationAvailableZone
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetTransitRouterDestinationNetworkInterface() *string {
	return s.TransitRouterDestinationNetworkInterface
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetTransitRouterDestinationResourceId() *string {
	return s.TransitRouterDestinationResourceId
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetTransitRouterDestinationVSwitchId() *string {
	return s.TransitRouterDestinationVSwitchId
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetTransitRouterPairAttachmentId() *string {
	return s.TransitRouterPairAttachmentId
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetTransitRouterSourceAccountId() *string {
	return s.TransitRouterSourceAccountId
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetTransitRouterSourceAvailableZone() *string {
	return s.TransitRouterSourceAvailableZone
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetTransitRouterSourceNetworkInterface() *string {
	return s.TransitRouterSourceNetworkInterface
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetTransitRouterSourceResourceId() *string {
	return s.TransitRouterSourceResourceId
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetTransitRouterSourceVSwitchId() *string {
	return s.TransitRouterSourceVSwitchId
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetBindingResourceId(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.BindingResourceId = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetBindingResourceType(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.BindingResourceType = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetBytes(v float64) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.Bytes = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetBytesRate(v float64) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.BytesRate = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetClientAsn(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.ClientAsn = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetClientCity(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.ClientCity = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetClientCountry(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.ClientCountry = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetClientIsp(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.ClientIsp = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetClientProvince(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.ClientProvince = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetDestinationIp(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.DestinationIp = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetDestinationPort(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.DestinationPort = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetDestinationRegionNo(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.DestinationRegionNo = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetDirection(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.Direction = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetDscp(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.Dscp = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetEcsId(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.EcsId = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetInstanceId(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.InstanceId = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetNetworkInterfaceId(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.NetworkInterfaceId = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetPackets(v float64) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.Packets = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetPacketsLostBlackhole(v float64) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.PacketsLostBlackhole = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetPacketsLostNoRoute(v float64) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.PacketsLostNoRoute = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetPacketsLostTTLExpired(v float64) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.PacketsLostTTLExpired = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetProtocol(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.Protocol = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetPublicIpAddress(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.PublicIpAddress = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetRegionId(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.RegionId = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetRoundTripTime(v float64) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.RoundTripTime = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetSourceIp(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.SourceIp = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetSourcePort(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.SourcePort = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetSourceRegionNo(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.SourceRegionNo = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetTrafficPath(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.TrafficPath = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetTransitRouterAttachmentId(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetTransitRouterDestinationAccountId(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.TransitRouterDestinationAccountId = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetTransitRouterDestinationAvailableZone(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.TransitRouterDestinationAvailableZone = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetTransitRouterDestinationNetworkInterface(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.TransitRouterDestinationNetworkInterface = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetTransitRouterDestinationResourceId(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.TransitRouterDestinationResourceId = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetTransitRouterDestinationVSwitchId(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.TransitRouterDestinationVSwitchId = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetTransitRouterId(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.TransitRouterId = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetTransitRouterPairAttachmentId(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.TransitRouterPairAttachmentId = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetTransitRouterSourceAccountId(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.TransitRouterSourceAccountId = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetTransitRouterSourceAvailableZone(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.TransitRouterSourceAvailableZone = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetTransitRouterSourceNetworkInterface(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.TransitRouterSourceNetworkInterface = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetTransitRouterSourceResourceId(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.TransitRouterSourceResourceId = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetTransitRouterSourceVSwitchId(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.TransitRouterSourceVSwitchId = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetVSwitchId(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.VSwitchId = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) SetVpcId(v string) *DescribeNisTrafficRankingResponseBodyFlowRankingList {
	s.VpcId = &v
	return s
}

func (s *DescribeNisTrafficRankingResponseBodyFlowRankingList) Validate() error {
	return dara.Validate(s)
}
