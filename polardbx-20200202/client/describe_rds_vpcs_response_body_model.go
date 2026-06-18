// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRdsVpcsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRdsVpcsResponseBody
	GetRequestId() *string
	SetVpcs(v *DescribeRdsVpcsResponseBodyVpcs) *DescribeRdsVpcsResponseBody
	GetVpcs() *DescribeRdsVpcsResponseBodyVpcs
}

type DescribeRdsVpcsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4C1A7CFB-9F73-5041-8C74-27626E58985A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The VPC information.
	Vpcs *DescribeRdsVpcsResponseBodyVpcs `json:"Vpcs,omitempty" xml:"Vpcs,omitempty" type:"Struct"`
}

func (s DescribeRdsVpcsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsVpcsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRdsVpcsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRdsVpcsResponseBody) GetVpcs() *DescribeRdsVpcsResponseBodyVpcs {
	return s.Vpcs
}

func (s *DescribeRdsVpcsResponseBody) SetRequestId(v string) *DescribeRdsVpcsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRdsVpcsResponseBody) SetVpcs(v *DescribeRdsVpcsResponseBodyVpcs) *DescribeRdsVpcsResponseBody {
	s.Vpcs = v
	return s
}

func (s *DescribeRdsVpcsResponseBody) Validate() error {
	if s.Vpcs != nil {
		if err := s.Vpcs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRdsVpcsResponseBodyVpcs struct {
	// The VPC information.
	Vpc []*DescribeRdsVpcsResponseBodyVpcsVpc `json:"Vpc,omitempty" xml:"Vpc,omitempty" type:"Repeated"`
}

func (s DescribeRdsVpcsResponseBodyVpcs) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsVpcsResponseBodyVpcs) GoString() string {
	return s.String()
}

func (s *DescribeRdsVpcsResponseBodyVpcs) GetVpc() []*DescribeRdsVpcsResponseBodyVpcsVpc {
	return s.Vpc
}

func (s *DescribeRdsVpcsResponseBodyVpcs) SetVpc(v []*DescribeRdsVpcsResponseBodyVpcsVpc) *DescribeRdsVpcsResponseBodyVpcs {
	s.Vpc = v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcs) Validate() error {
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

type DescribeRdsVpcsResponseBodyVpcsVpc struct {
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// 314****36
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// 26842
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// The IPv4 CIDR block of the VPC.
	//
	// example:
	//
	// 47.121.246.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The time when the VPC was created.
	//
	// example:
	//
	// Mon May 12 10:13:14 CST 2025
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the VPC was last modified.
	//
	// example:
	//
	// 2024-09-29 17:47:56
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates whether the VPC is the default VPC. Valid values: **true**: The VPC is the default VPC. **false**: The VPC is not the default VPC.
	//
	// example:
	//
	// False
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The region in which the VPC resides.
	//
	// example:
	//
	// cn-shanghai
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The VPC status.
	//
	// example:
	//
	// True
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The vSwitch information of the VPC.
	VSwitchs []*DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs `json:"VSwitchs,omitempty" xml:"VSwitchs,omitempty" type:"Repeated"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-2zekldxxy28nobay7o1f3
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the VPC.
	//
	// example:
	//
	// cloud_auto_test
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeRdsVpcsResponseBodyVpcsVpc) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsVpcsResponseBodyVpcsVpc) GoString() string {
	return s.String()
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) GetAliUid() *string {
	return s.AliUid
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) GetBid() *string {
	return s.Bid
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) GetStatus() *string {
	return s.Status
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) GetVSwitchs() []*DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs {
	return s.VSwitchs
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetAliUid(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.AliUid = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetBid(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.Bid = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetCidrBlock(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.CidrBlock = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetGmtCreate(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.GmtCreate = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetGmtModified(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.GmtModified = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetIsDefault(v bool) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.IsDefault = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetRegionNo(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.RegionNo = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetStatus(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.Status = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetVSwitchs(v []*DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.VSwitchs = v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetVpcId(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) SetVpcName(v string) *DescribeRdsVpcsResponseBodyVpcsVpc {
	s.VpcName = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpc) Validate() error {
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

type DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs struct {
	// The vSwitch CIDR block.
	//
	// example:
	//
	// 10.0.0.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The time when the vSwitch was created. The value is a UNIX timestamp accurate to milliseconds.
	//
	// example:
	//
	// 1574156944000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the vSwitch was last modified. The value is a UNIX timestamp.
	//
	// example:
	//
	// 2025-09-03T19:19:13+08:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates whether the vSwitch is the default vSwitch.
	//
	// example:
	//
	// false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-beijing-g
	IzNo *string `json:"IzNo,omitempty" xml:"IzNo,omitempty"`
	// The vSwitch status.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-8vbkwmpn****4nrd639ih
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The name of the vSwitch.
	//
	// example:
	//
	// cloud_auto_g
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
}

func (s DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) GoString() string {
	return s.String()
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) GetIzNo() *string {
	return s.IzNo
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) GetStatus() *string {
	return s.Status
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) SetCidrBlock(v string) *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs {
	s.CidrBlock = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) SetGmtCreate(v string) *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs {
	s.GmtCreate = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) SetGmtModified(v string) *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs {
	s.GmtModified = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) SetIsDefault(v bool) *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs {
	s.IsDefault = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) SetIzNo(v string) *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs {
	s.IzNo = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) SetStatus(v string) *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs {
	s.Status = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) SetVSwitchId(v string) *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs {
	s.VSwitchId = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) SetVSwitchName(v string) *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs {
	s.VSwitchName = &v
	return s
}

func (s *DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs) Validate() error {
	return dara.Validate(s)
}
