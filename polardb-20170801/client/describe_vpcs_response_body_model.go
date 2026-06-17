// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeVpcsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVpcsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeVpcsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVpcsResponseBody
	GetTotalCount() *int32
	SetVpcs(v []*DescribeVpcsResponseBodyVpcs) *DescribeVpcsResponseBody
	GetVpcs() []*DescribeVpcsResponseBodyVpcs
}

type DescribeVpcsResponseBody struct {
	// The page number of the returned page. The default value is 1.
	//
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Valid values:
	//
	// - **30**
	//
	// - **50**
	//
	// - **100**
	//
	//   **30**
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 11FDB5A0-84F5-5361-B729-5770B0AEB9D5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 50
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// A list of VPCs.
	Vpcs []*DescribeVpcsResponseBodyVpcs `json:"Vpcs,omitempty" xml:"Vpcs,omitempty" type:"Repeated"`
}

func (s DescribeVpcsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVpcsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVpcsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVpcsResponseBody) GetVpcs() []*DescribeVpcsResponseBodyVpcs {
	return s.Vpcs
}

func (s *DescribeVpcsResponseBody) SetPageNumber(v int32) *DescribeVpcsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVpcsResponseBody) SetPageSize(v int32) *DescribeVpcsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcsResponseBody) SetRequestId(v string) *DescribeVpcsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcsResponseBody) SetTotalCount(v int32) *DescribeVpcsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpcsResponseBody) SetVpcs(v []*DescribeVpcsResponseBodyVpcs) *DescribeVpcsResponseBody {
	s.Vpcs = v
	return s
}

func (s *DescribeVpcsResponseBody) Validate() error {
	if s.Vpcs != nil {
		for _, item := range s.Vpcs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpcsResponseBodyVpcs struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 1868512340232755
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// Indicates whether the account is an Alibaba Finance Cloud account, an Alibaba Gov Cloud account, or a public cloud account.
	//
	// example:
	//
	// 26842
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// The IPv4 CIDR block of the VPC.
	//
	// example:
	//
	// 57.100.6.59/32
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The time when the VPC was created.
	//
	// example:
	//
	// 2021-04-18T15:02:37Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the VPC was last modified.
	//
	// example:
	//
	// 2021-04-18T15:02:37Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates whether the VPC is the default VPC. Valid values:
	//
	// - **true**: The VPC is the default VPC.
	//
	// - **false**: The VPC is not the default VPC.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The ID of the region to which the VPC belongs.
	//
	// example:
	//
	// cn-chengdu-wt97-a01
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The status of the VPC. Valid values:
	//
	// - `Pending`: The VPC is being configured.
	//
	// - `Available`: The VPC is active.
	//
	// example:
	//
	// Pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// A list of vSwitches.
	VSwitchs []*DescribeVpcsResponseBodyVpcsVSwitchs `json:"VSwitchs,omitempty" xml:"VSwitchs,omitempty" type:"Repeated"`
	// The ID of the VPC.
	//
	// > You can call the [DescribeVpcs](https://help.aliyun.com/document_detail/35739.html) operation to query the details of VPCs.
	//
	// example:
	//
	// vpc-bp16efwqjzyumc23c647v
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the VPC.
	//
	// example:
	//
	// vpc-e2e-10341f3
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeVpcsResponseBodyVpcs) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsResponseBodyVpcs) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponseBodyVpcs) GetAliUid() *string {
	return s.AliUid
}

func (s *DescribeVpcsResponseBodyVpcs) GetBid() *string {
	return s.Bid
}

func (s *DescribeVpcsResponseBodyVpcs) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeVpcsResponseBodyVpcs) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeVpcsResponseBodyVpcs) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeVpcsResponseBodyVpcs) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeVpcsResponseBodyVpcs) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeVpcsResponseBodyVpcs) GetStatus() *string {
	return s.Status
}

func (s *DescribeVpcsResponseBodyVpcs) GetVSwitchs() []*DescribeVpcsResponseBodyVpcsVSwitchs {
	return s.VSwitchs
}

func (s *DescribeVpcsResponseBodyVpcs) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcsResponseBodyVpcs) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeVpcsResponseBodyVpcs) SetAliUid(v string) *DescribeVpcsResponseBodyVpcs {
	s.AliUid = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcs) SetBid(v string) *DescribeVpcsResponseBodyVpcs {
	s.Bid = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcs) SetCidrBlock(v string) *DescribeVpcsResponseBodyVpcs {
	s.CidrBlock = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcs) SetGmtCreate(v string) *DescribeVpcsResponseBodyVpcs {
	s.GmtCreate = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcs) SetGmtModified(v string) *DescribeVpcsResponseBodyVpcs {
	s.GmtModified = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcs) SetIsDefault(v bool) *DescribeVpcsResponseBodyVpcs {
	s.IsDefault = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcs) SetRegionNo(v string) *DescribeVpcsResponseBodyVpcs {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcs) SetStatus(v string) *DescribeVpcsResponseBodyVpcs {
	s.Status = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcs) SetVSwitchs(v []*DescribeVpcsResponseBodyVpcsVSwitchs) *DescribeVpcsResponseBodyVpcs {
	s.VSwitchs = v
	return s
}

func (s *DescribeVpcsResponseBodyVpcs) SetVpcId(v string) *DescribeVpcsResponseBodyVpcs {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcs) SetVpcName(v string) *DescribeVpcsResponseBodyVpcs {
	s.VpcName = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcs) Validate() error {
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

type DescribeVpcsResponseBodyVpcsVSwitchs struct {
	// The IPv4 CIDR block of the vSwitch.
	//
	// example:
	//
	// 47.118.126.0/25
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The time when the vSwitch was created.
	//
	// example:
	//
	// 2021-04-18T15:02:37Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the vSwitch was last modified.
	//
	// example:
	//
	// 2021-04-18T15:02:37Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates whether the vSwitch is the default vSwitch. Valid values:
	//
	// - **true**: The vSwitch is the default vSwitch.
	//
	// - **false**: The vSwitch is not the default vSwitch.
	//
	// example:
	//
	// false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The zone to which the vSwitch belongs.
	//
	// example:
	//
	// cn-shenzhen-f
	IzNo *string `json:"IzNo,omitempty" xml:"IzNo,omitempty"`
	// The status of the vSwitch. Valid values:
	//
	// - **Pending**: The vSwitch is being configured.
	//
	// - **Available**: The vSwitch is active.
	//
	// example:
	//
	// Pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-uf6fus5py6hbvxqwzwnk8
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The name of the vSwitch.
	//
	// example:
	//
	// default-sw
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
}

func (s DescribeVpcsResponseBodyVpcsVSwitchs) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsResponseBodyVpcsVSwitchs) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponseBodyVpcsVSwitchs) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeVpcsResponseBodyVpcsVSwitchs) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeVpcsResponseBodyVpcsVSwitchs) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeVpcsResponseBodyVpcsVSwitchs) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeVpcsResponseBodyVpcsVSwitchs) GetIzNo() *string {
	return s.IzNo
}

func (s *DescribeVpcsResponseBodyVpcsVSwitchs) GetStatus() *string {
	return s.Status
}

func (s *DescribeVpcsResponseBodyVpcsVSwitchs) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeVpcsResponseBodyVpcsVSwitchs) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *DescribeVpcsResponseBodyVpcsVSwitchs) SetCidrBlock(v string) *DescribeVpcsResponseBodyVpcsVSwitchs {
	s.CidrBlock = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVSwitchs) SetGmtCreate(v string) *DescribeVpcsResponseBodyVpcsVSwitchs {
	s.GmtCreate = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVSwitchs) SetGmtModified(v string) *DescribeVpcsResponseBodyVpcsVSwitchs {
	s.GmtModified = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVSwitchs) SetIsDefault(v bool) *DescribeVpcsResponseBodyVpcsVSwitchs {
	s.IsDefault = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVSwitchs) SetIzNo(v string) *DescribeVpcsResponseBodyVpcsVSwitchs {
	s.IzNo = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVSwitchs) SetStatus(v string) *DescribeVpcsResponseBodyVpcsVSwitchs {
	s.Status = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVSwitchs) SetVSwitchId(v string) *DescribeVpcsResponseBodyVpcsVSwitchs {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVSwitchs) SetVSwitchName(v string) *DescribeVpcsResponseBodyVpcsVSwitchs {
	s.VSwitchName = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVSwitchs) Validate() error {
	return dara.Validate(s)
}
