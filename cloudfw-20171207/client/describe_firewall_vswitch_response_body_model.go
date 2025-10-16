// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFirewallVSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeFirewallVSwitchResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeFirewallVSwitchResponseBody
	GetTotalCount() *int32
	SetVswitchList(v []*DescribeFirewallVSwitchResponseBodyVswitchList) *DescribeFirewallVSwitchResponseBody
	GetVswitchList() []*DescribeFirewallVSwitchResponseBodyVswitchList
}

type DescribeFirewallVSwitchResponseBody struct {
	// example:
	//
	// A1562A68-99FA-5D6B-BD5B-2F959F25****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20
	TotalCount  *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VswitchList []*DescribeFirewallVSwitchResponseBodyVswitchList `json:"VswitchList,omitempty" xml:"VswitchList,omitempty" type:"Repeated"`
}

func (s DescribeFirewallVSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallVSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFirewallVSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFirewallVSwitchResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeFirewallVSwitchResponseBody) GetVswitchList() []*DescribeFirewallVSwitchResponseBodyVswitchList {
	return s.VswitchList
}

func (s *DescribeFirewallVSwitchResponseBody) SetRequestId(v string) *DescribeFirewallVSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFirewallVSwitchResponseBody) SetTotalCount(v int32) *DescribeFirewallVSwitchResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeFirewallVSwitchResponseBody) SetVswitchList(v []*DescribeFirewallVSwitchResponseBodyVswitchList) *DescribeFirewallVSwitchResponseBody {
	s.VswitchList = v
	return s
}

func (s *DescribeFirewallVSwitchResponseBody) Validate() error {
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

type DescribeFirewallVSwitchResponseBodyVswitchList struct {
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
	// 10
	FirewallCount *string                                                       `json:"FirewallCount,omitempty" xml:"FirewallCount,omitempty"`
	FirewallList  []*DescribeFirewallVSwitchResponseBodyVswitchListFirewallList `json:"FirewallList,omitempty" xml:"FirewallList,omitempty" type:"Repeated"`
	// example:
	//
	// 184480249330****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// example:
	//
	// vpc-2zeez7gymz5r4pi****am
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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

func (s DescribeFirewallVSwitchResponseBodyVswitchList) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallVSwitchResponseBodyVswitchList) GoString() string {
	return s.String()
}

func (s *DescribeFirewallVSwitchResponseBodyVswitchList) GetAvailableIpCount() *string {
	return s.AvailableIpCount
}

func (s *DescribeFirewallVSwitchResponseBodyVswitchList) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeFirewallVSwitchResponseBodyVswitchList) GetFirewallCount() *string {
	return s.FirewallCount
}

func (s *DescribeFirewallVSwitchResponseBodyVswitchList) GetFirewallList() []*DescribeFirewallVSwitchResponseBodyVswitchListFirewallList {
	return s.FirewallList
}

func (s *DescribeFirewallVSwitchResponseBodyVswitchList) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeFirewallVSwitchResponseBodyVswitchList) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeFirewallVSwitchResponseBodyVswitchList) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeFirewallVSwitchResponseBodyVswitchList) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeFirewallVSwitchResponseBodyVswitchList) GetVswitchName() *string {
	return s.VswitchName
}

func (s *DescribeFirewallVSwitchResponseBodyVswitchList) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeFirewallVSwitchResponseBodyVswitchList) SetAvailableIpCount(v string) *DescribeFirewallVSwitchResponseBodyVswitchList {
	s.AvailableIpCount = &v
	return s
}

func (s *DescribeFirewallVSwitchResponseBodyVswitchList) SetCidrBlock(v string) *DescribeFirewallVSwitchResponseBodyVswitchList {
	s.CidrBlock = &v
	return s
}

func (s *DescribeFirewallVSwitchResponseBodyVswitchList) SetFirewallCount(v string) *DescribeFirewallVSwitchResponseBodyVswitchList {
	s.FirewallCount = &v
	return s
}

func (s *DescribeFirewallVSwitchResponseBodyVswitchList) SetFirewallList(v []*DescribeFirewallVSwitchResponseBodyVswitchListFirewallList) *DescribeFirewallVSwitchResponseBodyVswitchList {
	s.FirewallList = v
	return s
}

func (s *DescribeFirewallVSwitchResponseBodyVswitchList) SetMemberUid(v string) *DescribeFirewallVSwitchResponseBodyVswitchList {
	s.MemberUid = &v
	return s
}

func (s *DescribeFirewallVSwitchResponseBodyVswitchList) SetRegionNo(v string) *DescribeFirewallVSwitchResponseBodyVswitchList {
	s.RegionNo = &v
	return s
}

func (s *DescribeFirewallVSwitchResponseBodyVswitchList) SetVpcId(v string) *DescribeFirewallVSwitchResponseBodyVswitchList {
	s.VpcId = &v
	return s
}

func (s *DescribeFirewallVSwitchResponseBodyVswitchList) SetVswitchId(v string) *DescribeFirewallVSwitchResponseBodyVswitchList {
	s.VswitchId = &v
	return s
}

func (s *DescribeFirewallVSwitchResponseBodyVswitchList) SetVswitchName(v string) *DescribeFirewallVSwitchResponseBodyVswitchList {
	s.VswitchName = &v
	return s
}

func (s *DescribeFirewallVSwitchResponseBodyVswitchList) SetZoneId(v string) *DescribeFirewallVSwitchResponseBodyVswitchList {
	s.ZoneId = &v
	return s
}

func (s *DescribeFirewallVSwitchResponseBodyVswitchList) Validate() error {
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

type DescribeFirewallVSwitchResponseBodyVswitchListFirewallList struct {
	// example:
	//
	// vfw-tr-37145c8f5ede45e9****
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	// example:
	//
	// test-Firewall
	FirewallName *string `json:"FirewallName,omitempty" xml:"FirewallName,omitempty"`
}

func (s DescribeFirewallVSwitchResponseBodyVswitchListFirewallList) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallVSwitchResponseBodyVswitchListFirewallList) GoString() string {
	return s.String()
}

func (s *DescribeFirewallVSwitchResponseBodyVswitchListFirewallList) GetFirewallId() *string {
	return s.FirewallId
}

func (s *DescribeFirewallVSwitchResponseBodyVswitchListFirewallList) GetFirewallName() *string {
	return s.FirewallName
}

func (s *DescribeFirewallVSwitchResponseBodyVswitchListFirewallList) SetFirewallId(v string) *DescribeFirewallVSwitchResponseBodyVswitchListFirewallList {
	s.FirewallId = &v
	return s
}

func (s *DescribeFirewallVSwitchResponseBodyVswitchListFirewallList) SetFirewallName(v string) *DescribeFirewallVSwitchResponseBodyVswitchListFirewallList {
	s.FirewallName = &v
	return s
}

func (s *DescribeFirewallVSwitchResponseBodyVswitchListFirewallList) Validate() error {
	return dara.Validate(s)
}
