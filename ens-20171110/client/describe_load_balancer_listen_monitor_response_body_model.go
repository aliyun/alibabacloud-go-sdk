// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerListenMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLoadBalancerMonitorListenData(v []*DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) *DescribeLoadBalancerListenMonitorResponseBody
	GetLoadBalancerMonitorListenData() []*DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData
	SetRequestId(v string) *DescribeLoadBalancerListenMonitorResponseBody
	GetRequestId() *string
}

type DescribeLoadBalancerListenMonitorResponseBody struct {
	// The TCP/UDP monitoring data of the ELB instance.
	LoadBalancerMonitorListenData []*DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData `json:"LoadBalancerMonitorListenData,omitempty" xml:"LoadBalancerMonitorListenData,omitempty" type:"Repeated"`
	// Id of the request.
	//
	// example:
	//
	// C0003E8B-B930-4F59-ADC0-0E209A9012A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLoadBalancerListenMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerListenMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerListenMonitorResponseBody) GetLoadBalancerMonitorListenData() []*DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData {
	return s.LoadBalancerMonitorListenData
}

func (s *DescribeLoadBalancerListenMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLoadBalancerListenMonitorResponseBody) SetLoadBalancerMonitorListenData(v []*DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) *DescribeLoadBalancerListenMonitorResponseBody {
	s.LoadBalancerMonitorListenData = v
	return s
}

func (s *DescribeLoadBalancerListenMonitorResponseBody) SetRequestId(v string) *DescribeLoadBalancerListenMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoadBalancerListenMonitorResponseBody) Validate() error {
	if s.LoadBalancerMonitorListenData != nil {
		for _, item := range s.LoadBalancerMonitorListenData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData struct {
	// The number of active connections.
	//
	// example:
	//
	// 80285
	ActConns *string `json:"ActConns,omitempty" xml:"ActConns,omitempty"`
	// The business time.
	//
	// example:
	//
	// 2024-01-15 16:03:00
	BizTime *string `json:"BizTime,omitempty" xml:"BizTime,omitempty"`
	// The number of new connections.
	//
	// example:
	//
	// 37150
	Conns *string `json:"Conns,omitempty" xml:"Conns,omitempty"`
	// The number of dropped connections.
	//
	// example:
	//
	// 10
	DropConns *string `json:"DropConns,omitempty" xml:"DropConns,omitempty"`
	// The ID of the node to which the ELB instance belongs.
	//
	// example:
	//
	// cn-dongguan-9
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The number of inactive connections.
	//
	// example:
	//
	// 16322
	InActConns *string `json:"InActConns,omitempty" xml:"InActConns,omitempty"`
	// The inbound traffic.
	//
	// example:
	//
	// 67532
	InBytes *string `json:"InBytes,omitempty" xml:"InBytes,omitempty"`
	// The dropped inbound traffic.
	//
	// example:
	//
	// 324
	InDropBytes *string `json:"InDropBytes,omitempty" xml:"InDropBytes,omitempty"`
	// The number of dropped inbound packets.
	//
	// example:
	//
	// 27
	InDropPkts *string `json:"InDropPkts,omitempty" xml:"InDropPkts,omitempty"`
	// The number of inbound packets.
	//
	// example:
	//
	// 12
	InPkts *string `json:"InPkts,omitempty" xml:"InPkts,omitempty"`
	// The number of unavailable servers that are attached to the monitored ELB instance.
	//
	// example:
	//
	// 0
	InValidRsNum *string `json:"InValidRsNum,omitempty" xml:"InValidRsNum,omitempty"`
	// The ID of the ELB instance.
	//
	// example:
	//
	// lb-5q73cv04zeyh43lh74lp4****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The outbound traffic.
	//
	// example:
	//
	// 5155487
	OutBytes *string `json:"OutBytes,omitempty" xml:"OutBytes,omitempty"`
	// The dropped outbound traffic.
	//
	// example:
	//
	// 0
	OutDropBytes *string `json:"OutDropBytes,omitempty" xml:"OutDropBytes,omitempty"`
	// The number of dropped outbound packets.
	//
	// example:
	//
	// 76
	OutDropPkts *string `json:"OutDropPkts,omitempty" xml:"OutDropPkts,omitempty"`
	// The number of outbound packets.
	//
	// example:
	//
	// 34
	OutPkts *string `json:"OutPkts,omitempty" xml:"OutPkts,omitempty"`
	// The network protocol.
	//
	// example:
	//
	// tcp
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The VIP port of the ELB instance.
	//
	// example:
	//
	// 80
	VPort *string `json:"VPort,omitempty" xml:"VPort,omitempty"`
	// The number of available servers that are attached to the monitored ELB instance.
	//
	// example:
	//
	// 2
	ValidRsNum *string `json:"ValidRsNum,omitempty" xml:"ValidRsNum,omitempty"`
	// The virtual IP address (VIP) of the instance.
	//
	// example:
	//
	// 10.8.*.*
	Vip *string `json:"Vip,omitempty" xml:"Vip,omitempty"`
	// The ID of the tunnel.
	//
	// example:
	//
	// 53284
	Vni *string `json:"Vni,omitempty" xml:"Vni,omitempty"`
}

func (s DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) GetActConns() *string {
	return s.ActConns
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) GetBizTime() *string {
	return s.BizTime
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) GetConns() *string {
	return s.Conns
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) GetDropConns() *string {
	return s.DropConns
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) GetInActConns() *string {
	return s.InActConns
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) GetInBytes() *string {
	return s.InBytes
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) GetInDropBytes() *string {
	return s.InDropBytes
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) GetInDropPkts() *string {
	return s.InDropPkts
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) GetInPkts() *string {
	return s.InPkts
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) GetInValidRsNum() *string {
	return s.InValidRsNum
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) GetOutBytes() *string {
	return s.OutBytes
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) GetOutDropBytes() *string {
	return s.OutDropBytes
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) GetOutDropPkts() *string {
	return s.OutDropPkts
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) GetOutPkts() *string {
	return s.OutPkts
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) GetProto() *string {
	return s.Proto
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) GetVPort() *string {
	return s.VPort
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) GetValidRsNum() *string {
	return s.ValidRsNum
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) GetVip() *string {
	return s.Vip
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) GetVni() *string {
	return s.Vni
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) SetActConns(v string) *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData {
	s.ActConns = &v
	return s
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) SetBizTime(v string) *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData {
	s.BizTime = &v
	return s
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) SetConns(v string) *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData {
	s.Conns = &v
	return s
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) SetDropConns(v string) *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData {
	s.DropConns = &v
	return s
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) SetEnsRegionId(v string) *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) SetInActConns(v string) *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData {
	s.InActConns = &v
	return s
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) SetInBytes(v string) *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData {
	s.InBytes = &v
	return s
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) SetInDropBytes(v string) *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData {
	s.InDropBytes = &v
	return s
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) SetInDropPkts(v string) *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData {
	s.InDropPkts = &v
	return s
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) SetInPkts(v string) *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData {
	s.InPkts = &v
	return s
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) SetInValidRsNum(v string) *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData {
	s.InValidRsNum = &v
	return s
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) SetLoadBalancerId(v string) *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) SetOutBytes(v string) *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData {
	s.OutBytes = &v
	return s
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) SetOutDropBytes(v string) *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData {
	s.OutDropBytes = &v
	return s
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) SetOutDropPkts(v string) *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData {
	s.OutDropPkts = &v
	return s
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) SetOutPkts(v string) *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData {
	s.OutPkts = &v
	return s
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) SetProto(v string) *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData {
	s.Proto = &v
	return s
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) SetVPort(v string) *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData {
	s.VPort = &v
	return s
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) SetValidRsNum(v string) *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData {
	s.ValidRsNum = &v
	return s
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) SetVip(v string) *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData {
	s.Vip = &v
	return s
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) SetVni(v string) *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData {
	s.Vni = &v
	return s
}

func (s *DescribeLoadBalancerListenMonitorResponseBodyLoadBalancerMonitorListenData) Validate() error {
	return dara.Validate(s)
}
