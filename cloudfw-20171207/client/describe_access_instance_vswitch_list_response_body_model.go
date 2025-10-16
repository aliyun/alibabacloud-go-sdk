// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessInstanceVSwitchListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *DescribeAccessInstanceVSwitchListResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeAccessInstanceVSwitchListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAccessInstanceVSwitchListResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeAccessInstanceVSwitchListResponseBody
	GetTotalCount() *int64
	SetZones(v []*DescribeAccessInstanceVSwitchListResponseBodyZones) *DescribeAccessInstanceVSwitchListResponseBody
	GetZones() []*DescribeAccessInstanceVSwitchListResponseBodyZones
}

type DescribeAccessInstanceVSwitchListResponseBody struct {
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 450D47F5-956E-543E-8502-2F71C8C54E72
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Zones      []*DescribeAccessInstanceVSwitchListResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s DescribeAccessInstanceVSwitchListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessInstanceVSwitchListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessInstanceVSwitchListResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeAccessInstanceVSwitchListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAccessInstanceVSwitchListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccessInstanceVSwitchListResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeAccessInstanceVSwitchListResponseBody) GetZones() []*DescribeAccessInstanceVSwitchListResponseBodyZones {
	return s.Zones
}

func (s *DescribeAccessInstanceVSwitchListResponseBody) SetPageNo(v int32) *DescribeAccessInstanceVSwitchListResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeAccessInstanceVSwitchListResponseBody) SetPageSize(v int32) *DescribeAccessInstanceVSwitchListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessInstanceVSwitchListResponseBody) SetRequestId(v string) *DescribeAccessInstanceVSwitchListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccessInstanceVSwitchListResponseBody) SetTotalCount(v int64) *DescribeAccessInstanceVSwitchListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAccessInstanceVSwitchListResponseBody) SetZones(v []*DescribeAccessInstanceVSwitchListResponseBodyZones) *DescribeAccessInstanceVSwitchListResponseBody {
	s.Zones = v
	return s
}

func (s *DescribeAccessInstanceVSwitchListResponseBody) Validate() error {
	if s.Zones != nil {
		for _, item := range s.Zones {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAccessInstanceVSwitchListResponseBodyZones struct {
	VSwitchList []*DescribeAccessInstanceVSwitchListResponseBodyZonesVSwitchList `json:"VSwitchList,omitempty" xml:"VSwitchList,omitempty" type:"Repeated"`
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAccessInstanceVSwitchListResponseBodyZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessInstanceVSwitchListResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeAccessInstanceVSwitchListResponseBodyZones) GetVSwitchList() []*DescribeAccessInstanceVSwitchListResponseBodyZonesVSwitchList {
	return s.VSwitchList
}

func (s *DescribeAccessInstanceVSwitchListResponseBodyZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAccessInstanceVSwitchListResponseBodyZones) SetVSwitchList(v []*DescribeAccessInstanceVSwitchListResponseBodyZonesVSwitchList) *DescribeAccessInstanceVSwitchListResponseBodyZones {
	s.VSwitchList = v
	return s
}

func (s *DescribeAccessInstanceVSwitchListResponseBodyZones) SetZoneId(v string) *DescribeAccessInstanceVSwitchListResponseBodyZones {
	s.ZoneId = &v
	return s
}

func (s *DescribeAccessInstanceVSwitchListResponseBodyZones) Validate() error {
	if s.VSwitchList != nil {
		for _, item := range s.VSwitchList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAccessInstanceVSwitchListResponseBodyZonesVSwitchList struct {
	// example:
	//
	// 254
	AvailableIpAddressCount *int64 `json:"AvailableIpAddressCount,omitempty" xml:"AvailableIpAddressCount,omitempty"`
	// example:
	//
	// 192.168.0.0/16
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// example:
	//
	// false
	FirewallVSwitch *bool `json:"FirewallVSwitch,omitempty" xml:"FirewallVSwitch,omitempty"`
	// example:
	//
	// vsw-qzeaol304m***
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// ManagedVSW
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	// example:
	//
	// vpc-uf6b5lyul0x******
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeAccessInstanceVSwitchListResponseBodyZonesVSwitchList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessInstanceVSwitchListResponseBodyZonesVSwitchList) GoString() string {
	return s.String()
}

func (s *DescribeAccessInstanceVSwitchListResponseBodyZonesVSwitchList) GetAvailableIpAddressCount() *int64 {
	return s.AvailableIpAddressCount
}

func (s *DescribeAccessInstanceVSwitchListResponseBodyZonesVSwitchList) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeAccessInstanceVSwitchListResponseBodyZonesVSwitchList) GetFirewallVSwitch() *bool {
	return s.FirewallVSwitch
}

func (s *DescribeAccessInstanceVSwitchListResponseBodyZonesVSwitchList) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeAccessInstanceVSwitchListResponseBodyZonesVSwitchList) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *DescribeAccessInstanceVSwitchListResponseBodyZonesVSwitchList) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeAccessInstanceVSwitchListResponseBodyZonesVSwitchList) SetAvailableIpAddressCount(v int64) *DescribeAccessInstanceVSwitchListResponseBodyZonesVSwitchList {
	s.AvailableIpAddressCount = &v
	return s
}

func (s *DescribeAccessInstanceVSwitchListResponseBodyZonesVSwitchList) SetCidrBlock(v string) *DescribeAccessInstanceVSwitchListResponseBodyZonesVSwitchList {
	s.CidrBlock = &v
	return s
}

func (s *DescribeAccessInstanceVSwitchListResponseBodyZonesVSwitchList) SetFirewallVSwitch(v bool) *DescribeAccessInstanceVSwitchListResponseBodyZonesVSwitchList {
	s.FirewallVSwitch = &v
	return s
}

func (s *DescribeAccessInstanceVSwitchListResponseBodyZonesVSwitchList) SetVSwitchId(v string) *DescribeAccessInstanceVSwitchListResponseBodyZonesVSwitchList {
	s.VSwitchId = &v
	return s
}

func (s *DescribeAccessInstanceVSwitchListResponseBodyZonesVSwitchList) SetVSwitchName(v string) *DescribeAccessInstanceVSwitchListResponseBodyZonesVSwitchList {
	s.VSwitchName = &v
	return s
}

func (s *DescribeAccessInstanceVSwitchListResponseBodyZonesVSwitchList) SetVpcId(v string) *DescribeAccessInstanceVSwitchListResponseBodyZonesVSwitchList {
	s.VpcId = &v
	return s
}

func (s *DescribeAccessInstanceVSwitchListResponseBodyZonesVSwitchList) Validate() error {
	return dara.Validate(s)
}
