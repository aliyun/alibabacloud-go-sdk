// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFirewallVswitchResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeFirewallVswitchResourcesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeFirewallVswitchResourcesResponseBody
	GetTotalCount() *int32
	SetVswitchList(v []*DescribeFirewallVswitchResourcesResponseBodyVswitchList) *DescribeFirewallVswitchResourcesResponseBody
	GetVswitchList() []*DescribeFirewallVswitchResourcesResponseBodyVswitchList
}

type DescribeFirewallVswitchResourcesResponseBody struct {
	// example:
	//
	// A81E99DF-07CF-5EE4-966A-9FF9F2F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount  *int32                                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VswitchList []*DescribeFirewallVswitchResourcesResponseBodyVswitchList `json:"VswitchList,omitempty" xml:"VswitchList,omitempty" type:"Repeated"`
}

func (s DescribeFirewallVswitchResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallVswitchResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFirewallVswitchResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFirewallVswitchResourcesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeFirewallVswitchResourcesResponseBody) GetVswitchList() []*DescribeFirewallVswitchResourcesResponseBodyVswitchList {
	return s.VswitchList
}

func (s *DescribeFirewallVswitchResourcesResponseBody) SetRequestId(v string) *DescribeFirewallVswitchResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFirewallVswitchResourcesResponseBody) SetTotalCount(v int32) *DescribeFirewallVswitchResourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeFirewallVswitchResourcesResponseBody) SetVswitchList(v []*DescribeFirewallVswitchResourcesResponseBodyVswitchList) *DescribeFirewallVswitchResourcesResponseBody {
	s.VswitchList = v
	return s
}

func (s *DescribeFirewallVswitchResourcesResponseBody) Validate() error {
	if s.VswitchList != nil {
		for _, item := range s.VswitchList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFirewallVswitchResourcesResponseBodyVswitchList struct {
	// example:
	//
	// 10
	AvailableIpCount *string `json:"AvailableIpCount,omitempty" xml:"AvailableIpCount,omitempty"`
	// example:
	//
	// 192.168.0.XX/16
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// example:
	//
	// []
	Detail       *string                                                                `json:"Detail,omitempty" xml:"Detail,omitempty"`
	FirewallList []*DescribeFirewallVswitchResourcesResponseBodyVswitchListFirewallList `json:"FirewallList,omitempty" xml:"FirewallList,omitempty" type:"Repeated"`
	// example:
	//
	// vtb-uf6ml7rgw5gzzdr****
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	// example:
	//
	// Custom
	RouteTableType *string `json:"RouteTableType,omitempty" xml:"RouteTableType,omitempty"`
	// example:
	//
	// open
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// vsw-bp10qla9mgi42eo****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// example:
	//
	// vsw-test
	VswitchName *string `json:"VswitchName,omitempty" xml:"VswitchName,omitempty"`
	// example:
	//
	// cn-shanghai-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeFirewallVswitchResourcesResponseBodyVswitchList) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallVswitchResourcesResponseBodyVswitchList) GoString() string {
	return s.String()
}

func (s *DescribeFirewallVswitchResourcesResponseBodyVswitchList) GetAvailableIpCount() *string {
	return s.AvailableIpCount
}

func (s *DescribeFirewallVswitchResourcesResponseBodyVswitchList) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeFirewallVswitchResourcesResponseBodyVswitchList) GetDetail() *string {
	return s.Detail
}

func (s *DescribeFirewallVswitchResourcesResponseBodyVswitchList) GetFirewallList() []*DescribeFirewallVswitchResourcesResponseBodyVswitchListFirewallList {
	return s.FirewallList
}

func (s *DescribeFirewallVswitchResourcesResponseBodyVswitchList) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeFirewallVswitchResourcesResponseBodyVswitchList) GetRouteTableType() *string {
	return s.RouteTableType
}

func (s *DescribeFirewallVswitchResourcesResponseBodyVswitchList) GetStatus() *string {
	return s.Status
}

func (s *DescribeFirewallVswitchResourcesResponseBodyVswitchList) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeFirewallVswitchResourcesResponseBodyVswitchList) GetVswitchName() *string {
	return s.VswitchName
}

func (s *DescribeFirewallVswitchResourcesResponseBodyVswitchList) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeFirewallVswitchResourcesResponseBodyVswitchList) SetAvailableIpCount(v string) *DescribeFirewallVswitchResourcesResponseBodyVswitchList {
	s.AvailableIpCount = &v
	return s
}

func (s *DescribeFirewallVswitchResourcesResponseBodyVswitchList) SetCidrBlock(v string) *DescribeFirewallVswitchResourcesResponseBodyVswitchList {
	s.CidrBlock = &v
	return s
}

func (s *DescribeFirewallVswitchResourcesResponseBodyVswitchList) SetDetail(v string) *DescribeFirewallVswitchResourcesResponseBodyVswitchList {
	s.Detail = &v
	return s
}

func (s *DescribeFirewallVswitchResourcesResponseBodyVswitchList) SetFirewallList(v []*DescribeFirewallVswitchResourcesResponseBodyVswitchListFirewallList) *DescribeFirewallVswitchResourcesResponseBodyVswitchList {
	s.FirewallList = v
	return s
}

func (s *DescribeFirewallVswitchResourcesResponseBodyVswitchList) SetRouteTableId(v string) *DescribeFirewallVswitchResourcesResponseBodyVswitchList {
	s.RouteTableId = &v
	return s
}

func (s *DescribeFirewallVswitchResourcesResponseBodyVswitchList) SetRouteTableType(v string) *DescribeFirewallVswitchResourcesResponseBodyVswitchList {
	s.RouteTableType = &v
	return s
}

func (s *DescribeFirewallVswitchResourcesResponseBodyVswitchList) SetStatus(v string) *DescribeFirewallVswitchResourcesResponseBodyVswitchList {
	s.Status = &v
	return s
}

func (s *DescribeFirewallVswitchResourcesResponseBodyVswitchList) SetVswitchId(v string) *DescribeFirewallVswitchResourcesResponseBodyVswitchList {
	s.VswitchId = &v
	return s
}

func (s *DescribeFirewallVswitchResourcesResponseBodyVswitchList) SetVswitchName(v string) *DescribeFirewallVswitchResourcesResponseBodyVswitchList {
	s.VswitchName = &v
	return s
}

func (s *DescribeFirewallVswitchResourcesResponseBodyVswitchList) SetZoneId(v string) *DescribeFirewallVswitchResourcesResponseBodyVswitchList {
	s.ZoneId = &v
	return s
}

func (s *DescribeFirewallVswitchResourcesResponseBodyVswitchList) Validate() error {
	if s.FirewallList != nil {
		for _, item := range s.FirewallList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFirewallVswitchResourcesResponseBodyVswitchListFirewallList struct {
	// example:
	//
	// vfw-tr-37145c8f5ede45e9****
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	// example:
	//
	// test-Firewall
	FirewallName *string `json:"FirewallName,omitempty" xml:"FirewallName,omitempty"`
	// example:
	//
	// NatFirewall
	FirewallType *string `json:"FirewallType,omitempty" xml:"FirewallType,omitempty"`
}

func (s DescribeFirewallVswitchResourcesResponseBodyVswitchListFirewallList) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallVswitchResourcesResponseBodyVswitchListFirewallList) GoString() string {
	return s.String()
}

func (s *DescribeFirewallVswitchResourcesResponseBodyVswitchListFirewallList) GetFirewallId() *string {
	return s.FirewallId
}

func (s *DescribeFirewallVswitchResourcesResponseBodyVswitchListFirewallList) GetFirewallName() *string {
	return s.FirewallName
}

func (s *DescribeFirewallVswitchResourcesResponseBodyVswitchListFirewallList) GetFirewallType() *string {
	return s.FirewallType
}

func (s *DescribeFirewallVswitchResourcesResponseBodyVswitchListFirewallList) SetFirewallId(v string) *DescribeFirewallVswitchResourcesResponseBodyVswitchListFirewallList {
	s.FirewallId = &v
	return s
}

func (s *DescribeFirewallVswitchResourcesResponseBodyVswitchListFirewallList) SetFirewallName(v string) *DescribeFirewallVswitchResourcesResponseBodyVswitchListFirewallList {
	s.FirewallName = &v
	return s
}

func (s *DescribeFirewallVswitchResourcesResponseBodyVswitchListFirewallList) SetFirewallType(v string) *DescribeFirewallVswitchResourcesResponseBodyVswitchListFirewallList {
	s.FirewallType = &v
	return s
}

func (s *DescribeFirewallVswitchResourcesResponseBodyVswitchListFirewallList) Validate() error {
	return dara.Validate(s)
}
