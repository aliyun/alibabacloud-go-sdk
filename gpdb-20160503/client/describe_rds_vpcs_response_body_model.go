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
	// The ID of the request.
	//
	// example:
	//
	// 53924AF0-1628-5AA2-9C95-D4**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried VPCs.
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
	return dara.Validate(s)
}

type DescribeRdsVpcsResponseBodyVpcs struct {
	// The queried VPC.
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
	return dara.Validate(s)
}

type DescribeRdsVpcsResponseBodyVpcsVpc struct {
	// An invalid parameter. It is no longer returned when you call this operation.
	//
	// example:
	//
	// null
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// An invalid parameter. It is no longer returned when you call this operation.
	//
	// example:
	//
	// null
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// The CIDR block of the VPC.
	//
	// example:
	//
	// 10.**.**.0/8
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// An invalid parameter. It is no longer returned when you call this operation.
	//
	// example:
	//
	// null
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// An invalid parameter. It is no longer returned when you call this operation.
	//
	// example:
	//
	// null
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates whether the VPC is the default VPC. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The state of the VPC. If **Available*	- is returned, the VPC is available.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The queried vSwitches.
	VSwitchs []*DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs `json:"VSwitchs,omitempty" xml:"VSwitchs,omitempty" type:"Repeated"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp*******************
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the VPC.
	//
	// example:
	//
	// vpc-name
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
	return dara.Validate(s)
}

type DescribeRdsVpcsResponseBodyVpcsVpcVSwitchs struct {
	// The CIDR block of the vSwitch.
	//
	// example:
	//
	// 10.**.**.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// An invalid parameter. It is no longer returned when you call this operation.
	//
	// example:
	//
	// null
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// An invalid parameter. It is no longer returned when you call this operation.
	//
	// example:
	//
	// null
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates whether the vSwitch is the default vSwitch. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The ID of the zone to which the vSwitch belongs.
	//
	// example:
	//
	// cn-hangzhou-i
	IzNo *string `json:"IzNo,omitempty" xml:"IzNo,omitempty"`
	// The state of the vSwitch. If **Available*	- is returned, the vSwitch is available.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp*******************
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The name of the vSwitch.
	//
	// example:
	//
	// vsw-name
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
