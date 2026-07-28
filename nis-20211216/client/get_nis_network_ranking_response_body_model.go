// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNisNetworkRankingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetNisNetworkRankingResponseBodyData) *GetNisNetworkRankingResponseBody
	GetData() []*GetNisNetworkRankingResponseBodyData
	SetRequestId(v string) *GetNisNetworkRankingResponseBody
	GetRequestId() *string
}

type GetNisNetworkRankingResponseBody struct {
	// The collection of cloud network metric ranking data.
	Data []*GetNisNetworkRankingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// A7F0D6EC-E19E-58AC-AC9F-08036763960F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNisNetworkRankingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNisNetworkRankingResponseBody) GoString() string {
	return s.String()
}

func (s *GetNisNetworkRankingResponseBody) GetData() []*GetNisNetworkRankingResponseBodyData {
	return s.Data
}

func (s *GetNisNetworkRankingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNisNetworkRankingResponseBody) SetData(v []*GetNisNetworkRankingResponseBodyData) *GetNisNetworkRankingResponseBody {
	s.Data = v
	return s
}

func (s *GetNisNetworkRankingResponseBody) SetRequestId(v string) *GetNisNetworkRankingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNisNetworkRankingResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetNisNetworkRankingResponseBodyData struct {
	// The number of concurrent connections.
	//
	// example:
	//
	// 66
	ActiveSessionCount *float64 `json:"ActiveSessionCount,omitempty" xml:"ActiveSessionCount,omitempty"`
	// The autonomous system number (ASN) of the client ISP.
	//
	// example:
	//
	// 129103
	Asn *string `json:"Asn,omitempty" xml:"Asn,omitempty"`
	// The transit router attachment ID.
	//
	// example:
	//
	// tr-sample*
	AttachmentId *string `json:"AttachmentId,omitempty" xml:"AttachmentId,omitempty"`
	// The bandwidth package instance ID.
	//
	// example:
	//
	// cbwp-sample*
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The traffic volume in bytes.
	//
	// example:
	//
	// 1024
	ByteCount *float64 `json:"ByteCount,omitempty" xml:"ByteCount,omitempty"`
	// The city where the client is located.
	//
	// example:
	//
	// Chengdu.
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The country where the client is located.
	//
	// example:
	//
	// China.
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The destination IP address.
	//
	// example:
	//
	// 2.2.XX.XX
	DestinationIp *string `json:"DestinationIp,omitempty" xml:"DestinationIp,omitempty"`
	// The destination ISP.
	//
	// example:
	//
	// Alibaba Cloud.
	DestinationIsp *string `json:"DestinationIsp,omitempty" xml:"DestinationIsp,omitempty"`
	// The destination port.
	//
	// example:
	//
	// 80
	DestinationPort *string `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty"`
	// The destination region ID.
	//
	// example:
	//
	// cn-hangzhou
	DestinationRegionNo *string `json:"DestinationRegionNo,omitempty" xml:"DestinationRegionNo,omitempty"`
	// The destination zone for probing.
	//
	// example:
	//
	// cn-hangzhou-b
	DestinationZone    *string `json:"DestinationZone,omitempty" xml:"DestinationZone,omitempty"`
	GlobalCountryCode  *string `json:"GlobalCountryCode,omitempty" xml:"GlobalCountryCode,omitempty"`
	GlobalProvinceCode *string `json:"GlobalProvinceCode,omitempty" xml:"GlobalProvinceCode,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 120.238.XX.XX
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The inbound bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 10
	InBps *float64 `json:"InBps,omitempty" xml:"InBps,omitempty"`
	// The inbound packet rate. Unit: packets per second.
	//
	// example:
	//
	// 3
	InPps *float64 `json:"InPps,omitempty" xml:"InPps,omitempty"`
	// The cloud resource instance ID corresponding to each scenario. For example, in the cross-region network traffic analysis scenario, this represents the CEN ID. In the public network scenario, this represents the EIP ID, ECS instance ID, or CLB ID.
	//
	// example:
	//
	// eip-sample*
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ISP of the client.
	//
	// example:
	//
	// China Mobile.
	Isp      *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	LineType *string `json:"LineType,omitempty" xml:"LineType,omitempty"`
	// The number of new connections per second.
	//
	// example:
	//
	// 18
	NewSessionPerSecond *float64 `json:"NewSessionPerSecond,omitempty" xml:"NewSessionPerSecond,omitempty"`
	// The outbound bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 88
	OutBps *float64 `json:"OutBps,omitempty" xml:"OutBps,omitempty"`
	// The outbound packet rate. Unit: packets per second.
	//
	// example:
	//
	// 8
	OutPps *float64 `json:"OutPps,omitempty" xml:"OutPps,omitempty"`
	// The number of traffic packets.
	//
	// example:
	//
	// 66
	PacketCount *float64 `json:"PacketCount,omitempty" xml:"PacketCount,omitempty"`
	// The network protocol.
	//
	// example:
	//
	// 6
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The province or state where the client is located.
	//
	// example:
	//
	// Sichuan.
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// The latency. Unit: ms.
	//
	// example:
	//
	// 23
	RTT *float64 `json:"RTT,omitempty" xml:"RTT,omitempty"`
	// The Alibaba Cloud region ID.
	//
	// example:
	//
	// cn-shenzhen
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The retransmission rate of TCP packets.
	//
	// example:
	//
	// 0.1
	RetransmitRate *float64 `json:"RetransmitRate,omitempty" xml:"RetransmitRate,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 42.120.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The source Internet Service Provider (ISP).
	//
	// example:
	//
	// China Mobile.
	SourceIsp *string `json:"SourceIsp,omitempty" xml:"SourceIsp,omitempty"`
	// The source port.
	//
	// example:
	//
	// 443
	SourcePort   *string `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	SourceRegion *string `json:"SourceRegion,omitempty" xml:"SourceRegion,omitempty"`
	// The source zone for probing.
	//
	// example:
	//
	// cn-hangzhou-a
	SourceZone *string `json:"SourceZone,omitempty" xml:"SourceZone,omitempty"`
	// The instance ID of the virtual border router (VBR).
	//
	// example:
	//
	// vbr-sample*
	VbrId *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
}

func (s GetNisNetworkRankingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetNisNetworkRankingResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNisNetworkRankingResponseBodyData) GetActiveSessionCount() *float64 {
	return s.ActiveSessionCount
}

func (s *GetNisNetworkRankingResponseBodyData) GetAsn() *string {
	return s.Asn
}

func (s *GetNisNetworkRankingResponseBodyData) GetAttachmentId() *string {
	return s.AttachmentId
}

func (s *GetNisNetworkRankingResponseBodyData) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *GetNisNetworkRankingResponseBodyData) GetByteCount() *float64 {
	return s.ByteCount
}

func (s *GetNisNetworkRankingResponseBodyData) GetCity() *string {
	return s.City
}

func (s *GetNisNetworkRankingResponseBodyData) GetCountry() *string {
	return s.Country
}

func (s *GetNisNetworkRankingResponseBodyData) GetDestinationIp() *string {
	return s.DestinationIp
}

func (s *GetNisNetworkRankingResponseBodyData) GetDestinationIsp() *string {
	return s.DestinationIsp
}

func (s *GetNisNetworkRankingResponseBodyData) GetDestinationPort() *string {
	return s.DestinationPort
}

func (s *GetNisNetworkRankingResponseBodyData) GetDestinationRegionNo() *string {
	return s.DestinationRegionNo
}

func (s *GetNisNetworkRankingResponseBodyData) GetDestinationZone() *string {
	return s.DestinationZone
}

func (s *GetNisNetworkRankingResponseBodyData) GetGlobalCountryCode() *string {
	return s.GlobalCountryCode
}

func (s *GetNisNetworkRankingResponseBodyData) GetGlobalProvinceCode() *string {
	return s.GlobalProvinceCode
}

func (s *GetNisNetworkRankingResponseBodyData) GetIP() *string {
	return s.IP
}

func (s *GetNisNetworkRankingResponseBodyData) GetInBps() *float64 {
	return s.InBps
}

func (s *GetNisNetworkRankingResponseBodyData) GetInPps() *float64 {
	return s.InPps
}

func (s *GetNisNetworkRankingResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetNisNetworkRankingResponseBodyData) GetIsp() *string {
	return s.Isp
}

func (s *GetNisNetworkRankingResponseBodyData) GetLineType() *string {
	return s.LineType
}

func (s *GetNisNetworkRankingResponseBodyData) GetNewSessionPerSecond() *float64 {
	return s.NewSessionPerSecond
}

func (s *GetNisNetworkRankingResponseBodyData) GetOutBps() *float64 {
	return s.OutBps
}

func (s *GetNisNetworkRankingResponseBodyData) GetOutPps() *float64 {
	return s.OutPps
}

func (s *GetNisNetworkRankingResponseBodyData) GetPacketCount() *float64 {
	return s.PacketCount
}

func (s *GetNisNetworkRankingResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *GetNisNetworkRankingResponseBodyData) GetProvince() *string {
	return s.Province
}

func (s *GetNisNetworkRankingResponseBodyData) GetRTT() *float64 {
	return s.RTT
}

func (s *GetNisNetworkRankingResponseBodyData) GetRegionNo() *string {
	return s.RegionNo
}

func (s *GetNisNetworkRankingResponseBodyData) GetRetransmitRate() *float64 {
	return s.RetransmitRate
}

func (s *GetNisNetworkRankingResponseBodyData) GetSourceIp() *string {
	return s.SourceIp
}

func (s *GetNisNetworkRankingResponseBodyData) GetSourceIsp() *string {
	return s.SourceIsp
}

func (s *GetNisNetworkRankingResponseBodyData) GetSourcePort() *string {
	return s.SourcePort
}

func (s *GetNisNetworkRankingResponseBodyData) GetSourceRegion() *string {
	return s.SourceRegion
}

func (s *GetNisNetworkRankingResponseBodyData) GetSourceZone() *string {
	return s.SourceZone
}

func (s *GetNisNetworkRankingResponseBodyData) GetVbrId() *string {
	return s.VbrId
}

func (s *GetNisNetworkRankingResponseBodyData) SetActiveSessionCount(v float64) *GetNisNetworkRankingResponseBodyData {
	s.ActiveSessionCount = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetAsn(v string) *GetNisNetworkRankingResponseBodyData {
	s.Asn = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetAttachmentId(v string) *GetNisNetworkRankingResponseBodyData {
	s.AttachmentId = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetBandwidthPackageId(v string) *GetNisNetworkRankingResponseBodyData {
	s.BandwidthPackageId = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetByteCount(v float64) *GetNisNetworkRankingResponseBodyData {
	s.ByteCount = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetCity(v string) *GetNisNetworkRankingResponseBodyData {
	s.City = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetCountry(v string) *GetNisNetworkRankingResponseBodyData {
	s.Country = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetDestinationIp(v string) *GetNisNetworkRankingResponseBodyData {
	s.DestinationIp = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetDestinationIsp(v string) *GetNisNetworkRankingResponseBodyData {
	s.DestinationIsp = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetDestinationPort(v string) *GetNisNetworkRankingResponseBodyData {
	s.DestinationPort = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetDestinationRegionNo(v string) *GetNisNetworkRankingResponseBodyData {
	s.DestinationRegionNo = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetDestinationZone(v string) *GetNisNetworkRankingResponseBodyData {
	s.DestinationZone = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetGlobalCountryCode(v string) *GetNisNetworkRankingResponseBodyData {
	s.GlobalCountryCode = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetGlobalProvinceCode(v string) *GetNisNetworkRankingResponseBodyData {
	s.GlobalProvinceCode = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetIP(v string) *GetNisNetworkRankingResponseBodyData {
	s.IP = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetInBps(v float64) *GetNisNetworkRankingResponseBodyData {
	s.InBps = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetInPps(v float64) *GetNisNetworkRankingResponseBodyData {
	s.InPps = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetInstanceId(v string) *GetNisNetworkRankingResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetIsp(v string) *GetNisNetworkRankingResponseBodyData {
	s.Isp = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetLineType(v string) *GetNisNetworkRankingResponseBodyData {
	s.LineType = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetNewSessionPerSecond(v float64) *GetNisNetworkRankingResponseBodyData {
	s.NewSessionPerSecond = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetOutBps(v float64) *GetNisNetworkRankingResponseBodyData {
	s.OutBps = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetOutPps(v float64) *GetNisNetworkRankingResponseBodyData {
	s.OutPps = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetPacketCount(v float64) *GetNisNetworkRankingResponseBodyData {
	s.PacketCount = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetProtocol(v string) *GetNisNetworkRankingResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetProvince(v string) *GetNisNetworkRankingResponseBodyData {
	s.Province = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetRTT(v float64) *GetNisNetworkRankingResponseBodyData {
	s.RTT = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetRegionNo(v string) *GetNisNetworkRankingResponseBodyData {
	s.RegionNo = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetRetransmitRate(v float64) *GetNisNetworkRankingResponseBodyData {
	s.RetransmitRate = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetSourceIp(v string) *GetNisNetworkRankingResponseBodyData {
	s.SourceIp = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetSourceIsp(v string) *GetNisNetworkRankingResponseBodyData {
	s.SourceIsp = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetSourcePort(v string) *GetNisNetworkRankingResponseBodyData {
	s.SourcePort = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetSourceRegion(v string) *GetNisNetworkRankingResponseBodyData {
	s.SourceRegion = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetSourceZone(v string) *GetNisNetworkRankingResponseBodyData {
	s.SourceZone = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) SetVbrId(v string) *GetNisNetworkRankingResponseBodyData {
	s.VbrId = &v
	return s
}

func (s *GetNisNetworkRankingResponseBodyData) Validate() error {
	return dara.Validate(s)
}
