// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallManualVSwitchListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVpcFirewallManualVSwitchListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVpcFirewallManualVSwitchListResponseBody
	GetTotalCount() *int32
	SetVSwitchList(v []*DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList) *DescribeVpcFirewallManualVSwitchListResponseBody
	GetVSwitchList() []*DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList
}

type DescribeVpcFirewallManualVSwitchListResponseBody struct {
	// example:
	//
	// 6EED3674-74E7-54DC-8FD4-6A374133****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3
	TotalCount  *int32                                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VSwitchList []*DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList `json:"VSwitchList,omitempty" xml:"VSwitchList,omitempty" type:"Repeated"`
}

func (s DescribeVpcFirewallManualVSwitchListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallManualVSwitchListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallManualVSwitchListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcFirewallManualVSwitchListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVpcFirewallManualVSwitchListResponseBody) GetVSwitchList() []*DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList {
	return s.VSwitchList
}

func (s *DescribeVpcFirewallManualVSwitchListResponseBody) SetRequestId(v string) *DescribeVpcFirewallManualVSwitchListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallManualVSwitchListResponseBody) SetTotalCount(v int32) *DescribeVpcFirewallManualVSwitchListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpcFirewallManualVSwitchListResponseBody) SetVSwitchList(v []*DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList) *DescribeVpcFirewallManualVSwitchListResponseBody {
	s.VSwitchList = v
	return s
}

func (s *DescribeVpcFirewallManualVSwitchListResponseBody) Validate() error {
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

type DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList struct {
	// example:
	//
	// 122167357026****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// example:
	//
	// 10
	AvailableIpAddressCount *int64 `json:"AvailableIpAddressCount,omitempty" xml:"AvailableIpAddressCount,omitempty"`
	// example:
	//
	// 10.0.31.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// example:
	//
	// 137578716100****
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// cn-beijing
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// example:
	//
	// vtb-2zet3gyk01o07so****
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// vsw-bp10d0kcp907721z****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// vsw-test
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	// example:
	//
	// vpc-uf62kq7c364sil2z2****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList) GetAvailableIpAddressCount() *int64 {
	return s.AvailableIpAddressCount
}

func (s *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList) GetStatus() *string {
	return s.Status
}

func (s *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList) SetAliUid(v int64) *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList {
	s.AliUid = &v
	return s
}

func (s *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList) SetAvailableIpAddressCount(v int64) *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList {
	s.AvailableIpAddressCount = &v
	return s
}

func (s *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList) SetCidrBlock(v string) *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList {
	s.CidrBlock = &v
	return s
}

func (s *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList) SetOwnerId(v int64) *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList) SetRegionNo(v string) *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList) SetRouteTableId(v string) *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList {
	s.RouteTableId = &v
	return s
}

func (s *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList) SetStatus(v string) *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList {
	s.Status = &v
	return s
}

func (s *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList) SetVSwitchId(v string) *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList) SetVSwitchName(v string) *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList {
	s.VSwitchName = &v
	return s
}

func (s *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList) SetVpcId(v string) *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList) SetZoneId(v string) *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList {
	s.ZoneId = &v
	return s
}

func (s *DescribeVpcFirewallManualVSwitchListResponseBodyVSwitchList) Validate() error {
	return dara.Validate(s)
}
