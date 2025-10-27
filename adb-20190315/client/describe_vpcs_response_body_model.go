// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVpcsResponseBody
	GetRequestId() *string
	SetVpcs(v *DescribeVpcsResponseBodyVpcs) *DescribeVpcsResponseBody
	GetVpcs() *DescribeVpcsResponseBodyVpcs
}

type DescribeVpcsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 78BDC895-F7C0-5961-92BE-F1C3D12B4BB5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried VPCs.
	Vpcs *DescribeVpcsResponseBodyVpcs `json:"Vpcs,omitempty" xml:"Vpcs,omitempty" type:"Struct"`
}

func (s DescribeVpcsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcsResponseBody) GetVpcs() *DescribeVpcsResponseBodyVpcs {
	return s.Vpcs
}

func (s *DescribeVpcsResponseBody) SetRequestId(v string) *DescribeVpcsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcsResponseBody) SetVpcs(v *DescribeVpcsResponseBodyVpcs) *DescribeVpcsResponseBody {
	s.Vpcs = v
	return s
}

func (s *DescribeVpcsResponseBody) Validate() error {
	if s.Vpcs != nil {
		if err := s.Vpcs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVpcsResponseBodyVpcs struct {
	// The queried VPC.
	Vpc []*DescribeVpcsResponseBodyVpcsVpc `json:"Vpc,omitempty" xml:"Vpc,omitempty" type:"Repeated"`
}

func (s DescribeVpcsResponseBodyVpcs) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsResponseBodyVpcs) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponseBodyVpcs) GetVpc() []*DescribeVpcsResponseBodyVpcsVpc {
	return s.Vpc
}

func (s *DescribeVpcsResponseBodyVpcs) SetVpc(v []*DescribeVpcsResponseBodyVpcsVpc) *DescribeVpcsResponseBodyVpcs {
	s.Vpc = v
	return s
}

func (s *DescribeVpcsResponseBodyVpcs) Validate() error {
	if s.Vpc != nil {
		for _, item := range s.Vpc {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpcsResponseBodyVpcsVpc struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 1431771514176727
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The user channel ID.
	//
	// example:
	//
	// 26842
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// The IPv4 CIDR block of the VPC.
	//
	// example:
	//
	// 192.168.0.0/16
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The time when the VPC was created.
	//
	// example:
	//
	// 1693217052000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the VPC was modified.
	//
	// example:
	//
	// 1724639118000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates whether the VPC is the default VPC in the region. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The region in which the VPC resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The status of the VPC. Valid values:
	//
	// 	- Pending
	//
	// 	- Available
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The vSwitches in the VPC.
	VSwitchs []*DescribeVpcsResponseBodyVpcsVpcVSwitchs `json:"VSwitchs,omitempty" xml:"VSwitchs,omitempty" type:"Repeated"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-uf6m0r5pihw1r79od6990
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the VPC.
	//
	// example:
	//
	// vpc1
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeVpcsResponseBodyVpcsVpc) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsResponseBodyVpcsVpc) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetAliUid() *string {
	return s.AliUid
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetBid() *string {
	return s.Bid
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetStatus() *string {
	return s.Status
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetVSwitchs() []*DescribeVpcsResponseBodyVpcsVpcVSwitchs {
	return s.VSwitchs
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetAliUid(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.AliUid = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetBid(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.Bid = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetCidrBlock(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.CidrBlock = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetGmtCreate(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.GmtCreate = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetGmtModified(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.GmtModified = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetIsDefault(v bool) *DescribeVpcsResponseBodyVpcsVpc {
	s.IsDefault = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetRegionNo(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetStatus(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.Status = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetVSwitchs(v []*DescribeVpcsResponseBodyVpcsVpcVSwitchs) *DescribeVpcsResponseBodyVpcsVpc {
	s.VSwitchs = v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetVpcId(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetVpcName(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.VpcName = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) Validate() error {
	if s.VSwitchs != nil {
		for _, item := range s.VSwitchs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpcsResponseBodyVpcsVpcVSwitchs struct {
	// The IPv4 CIDR block of the vSwitch.
	//
	// example:
	//
	// 172.17.0.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The time when the vSwitch was created.
	//
	// example:
	//
	// 1549012834000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the vSwitch was modified.
	//
	// example:
	//
	// 1731031910000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates whether the vSwitch is the default vSwitch. Valid values: **true*	- **false**
	//
	// example:
	//
	// false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The zone ID of the vSwitch.
	//
	// example:
	//
	// cn-hangzhou-j
	IzNo *string `json:"IzNo,omitempty" xml:"IzNo,omitempty"`
	// The status of the vSwitch. Valid values:
	//
	// 	- Pending
	//
	// 	- Available
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-8vbxk6ij0yz16bu4l3ijj
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The name of the vSwitch.
	//
	// example:
	//
	// vs1
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
}

func (s DescribeVpcsResponseBodyVpcsVpcVSwitchs) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsResponseBodyVpcsVpcVSwitchs) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponseBodyVpcsVpcVSwitchs) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeVpcsResponseBodyVpcsVpcVSwitchs) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeVpcsResponseBodyVpcsVpcVSwitchs) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeVpcsResponseBodyVpcsVpcVSwitchs) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeVpcsResponseBodyVpcsVpcVSwitchs) GetIzNo() *string {
	return s.IzNo
}

func (s *DescribeVpcsResponseBodyVpcsVpcVSwitchs) GetStatus() *string {
	return s.Status
}

func (s *DescribeVpcsResponseBodyVpcsVpcVSwitchs) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeVpcsResponseBodyVpcsVpcVSwitchs) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *DescribeVpcsResponseBodyVpcsVpcVSwitchs) SetCidrBlock(v string) *DescribeVpcsResponseBodyVpcsVpcVSwitchs {
	s.CidrBlock = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpcVSwitchs) SetGmtCreate(v string) *DescribeVpcsResponseBodyVpcsVpcVSwitchs {
	s.GmtCreate = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpcVSwitchs) SetGmtModified(v string) *DescribeVpcsResponseBodyVpcsVpcVSwitchs {
	s.GmtModified = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpcVSwitchs) SetIsDefault(v bool) *DescribeVpcsResponseBodyVpcsVpcVSwitchs {
	s.IsDefault = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpcVSwitchs) SetIzNo(v string) *DescribeVpcsResponseBodyVpcsVpcVSwitchs {
	s.IzNo = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpcVSwitchs) SetStatus(v string) *DescribeVpcsResponseBodyVpcsVpcVSwitchs {
	s.Status = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpcVSwitchs) SetVSwitchId(v string) *DescribeVpcsResponseBodyVpcsVpcVSwitchs {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpcVSwitchs) SetVSwitchName(v string) *DescribeVpcsResponseBodyVpcsVpcVSwitchs {
	s.VSwitchName = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpcVSwitchs) Validate() error {
	return dara.Validate(s)
}
