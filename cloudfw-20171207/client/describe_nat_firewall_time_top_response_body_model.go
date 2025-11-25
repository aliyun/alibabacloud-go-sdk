// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatFirewallTimeTopResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataCount(v int64) *DescribeNatFirewallTimeTopResponseBody
	GetDataCount() *int64
	SetDataList(v []*DescribeNatFirewallTimeTopResponseBodyDataList) *DescribeNatFirewallTimeTopResponseBody
	GetDataList() []*DescribeNatFirewallTimeTopResponseBodyDataList
	SetRequestId(v string) *DescribeNatFirewallTimeTopResponseBody
	GetRequestId() *string
	SetTrafficTime(v string) *DescribeNatFirewallTimeTopResponseBody
	GetTrafficTime() *string
}

type DescribeNatFirewallTimeTopResponseBody struct {
	// example:
	//
	// 2
	DataCount *int64                                            `json:"DataCount,omitempty" xml:"DataCount,omitempty"`
	DataList  []*DescribeNatFirewallTimeTopResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// example:
	//
	// C5DDD596-1191-5F36-A504-8733045A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1656923760
	TrafficTime *string `json:"TrafficTime,omitempty" xml:"TrafficTime,omitempty"`
}

func (s DescribeNatFirewallTimeTopResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallTimeTopResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallTimeTopResponseBody) GetDataCount() *int64 {
	return s.DataCount
}

func (s *DescribeNatFirewallTimeTopResponseBody) GetDataList() []*DescribeNatFirewallTimeTopResponseBodyDataList {
	return s.DataList
}

func (s *DescribeNatFirewallTimeTopResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNatFirewallTimeTopResponseBody) GetTrafficTime() *string {
	return s.TrafficTime
}

func (s *DescribeNatFirewallTimeTopResponseBody) SetDataCount(v int64) *DescribeNatFirewallTimeTopResponseBody {
	s.DataCount = &v
	return s
}

func (s *DescribeNatFirewallTimeTopResponseBody) SetDataList(v []*DescribeNatFirewallTimeTopResponseBodyDataList) *DescribeNatFirewallTimeTopResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeNatFirewallTimeTopResponseBody) SetRequestId(v string) *DescribeNatFirewallTimeTopResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNatFirewallTimeTopResponseBody) SetTrafficTime(v string) *DescribeNatFirewallTimeTopResponseBody {
	s.TrafficTime = &v
	return s
}

func (s *DescribeNatFirewallTimeTopResponseBody) Validate() error {
	if s.DataList != nil {
		for _, item := range s.DataList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNatFirewallTimeTopResponseBodyDataList struct {
	// example:
	//
	// 187
	InBps *int64 `json:"InBps,omitempty" xml:"InBps,omitempty"`
	// example:
	//
	// 10.66.231.XXX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// vfw-tr-7a9c8901ed394****
	NatFirewallId *string `json:"NatFirewallId,omitempty" xml:"NatFirewallId,omitempty"`
	// example:
	//
	// vfw-test
	NatFirewallName *string `json:"NatFirewallName,omitempty" xml:"NatFirewallName,omitempty"`
	// example:
	//
	// ngw-uf6pnry5vpawb****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// example:
	//
	// ngw-test
	NatGatewayName *string `json:"NatGatewayName,omitempty" xml:"NatGatewayName,omitempty"`
	// example:
	//
	// 27
	NewConn *string `json:"NewConn,omitempty" xml:"NewConn,omitempty"`
	// example:
	//
	// 45
	OutBps *int64 `json:"OutBps,omitempty" xml:"OutBps,omitempty"`
	// example:
	//
	// cn-qingdao
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// example:
	//
	// i-bp130nn8h6157dir****
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// example:
	//
	// test
	ResourceInstanceName *string `json:"ResourceInstanceName,omitempty" xml:"ResourceInstanceName,omitempty"`
	// example:
	//
	// 27
	SessionCount *string `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	// example:
	//
	// 232
	TotalBps *int64 `json:"TotalBps,omitempty" xml:"TotalBps,omitempty"`
}

func (s DescribeNatFirewallTimeTopResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallTimeTopResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallTimeTopResponseBodyDataList) GetInBps() *int64 {
	return s.InBps
}

func (s *DescribeNatFirewallTimeTopResponseBodyDataList) GetIp() *string {
	return s.Ip
}

func (s *DescribeNatFirewallTimeTopResponseBodyDataList) GetNatFirewallId() *string {
	return s.NatFirewallId
}

func (s *DescribeNatFirewallTimeTopResponseBodyDataList) GetNatFirewallName() *string {
	return s.NatFirewallName
}

func (s *DescribeNatFirewallTimeTopResponseBodyDataList) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeNatFirewallTimeTopResponseBodyDataList) GetNatGatewayName() *string {
	return s.NatGatewayName
}

func (s *DescribeNatFirewallTimeTopResponseBodyDataList) GetNewConn() *string {
	return s.NewConn
}

func (s *DescribeNatFirewallTimeTopResponseBodyDataList) GetOutBps() *int64 {
	return s.OutBps
}

func (s *DescribeNatFirewallTimeTopResponseBodyDataList) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeNatFirewallTimeTopResponseBodyDataList) GetResourceInstanceId() *string {
	return s.ResourceInstanceId
}

func (s *DescribeNatFirewallTimeTopResponseBodyDataList) GetResourceInstanceName() *string {
	return s.ResourceInstanceName
}

func (s *DescribeNatFirewallTimeTopResponseBodyDataList) GetSessionCount() *string {
	return s.SessionCount
}

func (s *DescribeNatFirewallTimeTopResponseBodyDataList) GetTotalBps() *int64 {
	return s.TotalBps
}

func (s *DescribeNatFirewallTimeTopResponseBodyDataList) SetInBps(v int64) *DescribeNatFirewallTimeTopResponseBodyDataList {
	s.InBps = &v
	return s
}

func (s *DescribeNatFirewallTimeTopResponseBodyDataList) SetIp(v string) *DescribeNatFirewallTimeTopResponseBodyDataList {
	s.Ip = &v
	return s
}

func (s *DescribeNatFirewallTimeTopResponseBodyDataList) SetNatFirewallId(v string) *DescribeNatFirewallTimeTopResponseBodyDataList {
	s.NatFirewallId = &v
	return s
}

func (s *DescribeNatFirewallTimeTopResponseBodyDataList) SetNatFirewallName(v string) *DescribeNatFirewallTimeTopResponseBodyDataList {
	s.NatFirewallName = &v
	return s
}

func (s *DescribeNatFirewallTimeTopResponseBodyDataList) SetNatGatewayId(v string) *DescribeNatFirewallTimeTopResponseBodyDataList {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeNatFirewallTimeTopResponseBodyDataList) SetNatGatewayName(v string) *DescribeNatFirewallTimeTopResponseBodyDataList {
	s.NatGatewayName = &v
	return s
}

func (s *DescribeNatFirewallTimeTopResponseBodyDataList) SetNewConn(v string) *DescribeNatFirewallTimeTopResponseBodyDataList {
	s.NewConn = &v
	return s
}

func (s *DescribeNatFirewallTimeTopResponseBodyDataList) SetOutBps(v int64) *DescribeNatFirewallTimeTopResponseBodyDataList {
	s.OutBps = &v
	return s
}

func (s *DescribeNatFirewallTimeTopResponseBodyDataList) SetRegionNo(v string) *DescribeNatFirewallTimeTopResponseBodyDataList {
	s.RegionNo = &v
	return s
}

func (s *DescribeNatFirewallTimeTopResponseBodyDataList) SetResourceInstanceId(v string) *DescribeNatFirewallTimeTopResponseBodyDataList {
	s.ResourceInstanceId = &v
	return s
}

func (s *DescribeNatFirewallTimeTopResponseBodyDataList) SetResourceInstanceName(v string) *DescribeNatFirewallTimeTopResponseBodyDataList {
	s.ResourceInstanceName = &v
	return s
}

func (s *DescribeNatFirewallTimeTopResponseBodyDataList) SetSessionCount(v string) *DescribeNatFirewallTimeTopResponseBodyDataList {
	s.SessionCount = &v
	return s
}

func (s *DescribeNatFirewallTimeTopResponseBodyDataList) SetTotalBps(v int64) *DescribeNatFirewallTimeTopResponseBodyDataList {
	s.TotalBps = &v
	return s
}

func (s *DescribeNatFirewallTimeTopResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
