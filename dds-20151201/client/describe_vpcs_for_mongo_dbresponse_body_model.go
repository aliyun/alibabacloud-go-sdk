// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcsForMongoDBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeVpcsForMongoDBResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVpcsForMongoDBResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeVpcsForMongoDBResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVpcsForMongoDBResponseBody
	GetTotalCount() *int32
	SetVpcs(v []*DescribeVpcsForMongoDBResponseBodyVpcs) *DescribeVpcsForMongoDBResponseBody
	GetVpcs() []*DescribeVpcsForMongoDBResponseBodyVpcs
}

type DescribeVpcsForMongoDBResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 53924AF0-1628-5AA2-9C95-D4**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 4
	TotalCount *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Vpcs       []*DescribeVpcsForMongoDBResponseBodyVpcs `json:"Vpcs,omitempty" xml:"Vpcs,omitempty" type:"Repeated"`
}

func (s DescribeVpcsForMongoDBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsForMongoDBResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcsForMongoDBResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVpcsForMongoDBResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVpcsForMongoDBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcsForMongoDBResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVpcsForMongoDBResponseBody) GetVpcs() []*DescribeVpcsForMongoDBResponseBodyVpcs {
	return s.Vpcs
}

func (s *DescribeVpcsForMongoDBResponseBody) SetPageNumber(v int32) *DescribeVpcsForMongoDBResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVpcsForMongoDBResponseBody) SetPageSize(v int32) *DescribeVpcsForMongoDBResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcsForMongoDBResponseBody) SetRequestId(v string) *DescribeVpcsForMongoDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcsForMongoDBResponseBody) SetTotalCount(v int32) *DescribeVpcsForMongoDBResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpcsForMongoDBResponseBody) SetVpcs(v []*DescribeVpcsForMongoDBResponseBodyVpcs) *DescribeVpcsForMongoDBResponseBody {
	s.Vpcs = v
	return s
}

func (s *DescribeVpcsForMongoDBResponseBody) Validate() error {
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

type DescribeVpcsForMongoDBResponseBodyVpcs struct {
	// example:
	//
	// null
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// example:
	//
	// null
	Bid       *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// example:
	//
	// null
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// null
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// example:
	//
	// Available
	Status   *string                                           `json:"Status,omitempty" xml:"Status,omitempty"`
	VSwitchs []*DescribeVpcsForMongoDBResponseBodyVpcsVSwitchs `json:"VSwitchs,omitempty" xml:"VSwitchs,omitempty" type:"Repeated"`
	// VPC ID。
	//
	// example:
	//
	// vpc-2zep2pepkthg5ueal****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// vpc-name
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeVpcsForMongoDBResponseBodyVpcs) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsForMongoDBResponseBodyVpcs) GoString() string {
	return s.String()
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcs) GetAliUid() *string {
	return s.AliUid
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcs) GetBid() *string {
	return s.Bid
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcs) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcs) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcs) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcs) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcs) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcs) GetStatus() *string {
	return s.Status
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcs) GetVSwitchs() []*DescribeVpcsForMongoDBResponseBodyVpcsVSwitchs {
	return s.VSwitchs
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcs) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcs) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcs) SetAliUid(v string) *DescribeVpcsForMongoDBResponseBodyVpcs {
	s.AliUid = &v
	return s
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcs) SetBid(v string) *DescribeVpcsForMongoDBResponseBodyVpcs {
	s.Bid = &v
	return s
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcs) SetCidrBlock(v string) *DescribeVpcsForMongoDBResponseBodyVpcs {
	s.CidrBlock = &v
	return s
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcs) SetGmtCreate(v string) *DescribeVpcsForMongoDBResponseBodyVpcs {
	s.GmtCreate = &v
	return s
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcs) SetGmtModified(v string) *DescribeVpcsForMongoDBResponseBodyVpcs {
	s.GmtModified = &v
	return s
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcs) SetIsDefault(v bool) *DescribeVpcsForMongoDBResponseBodyVpcs {
	s.IsDefault = &v
	return s
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcs) SetRegionNo(v string) *DescribeVpcsForMongoDBResponseBodyVpcs {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcs) SetStatus(v string) *DescribeVpcsForMongoDBResponseBodyVpcs {
	s.Status = &v
	return s
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcs) SetVSwitchs(v []*DescribeVpcsForMongoDBResponseBodyVpcsVSwitchs) *DescribeVpcsForMongoDBResponseBodyVpcs {
	s.VSwitchs = v
	return s
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcs) SetVpcId(v string) *DescribeVpcsForMongoDBResponseBodyVpcs {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcs) SetVpcName(v string) *DescribeVpcsForMongoDBResponseBodyVpcs {
	s.VpcName = &v
	return s
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcs) Validate() error {
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

type DescribeVpcsForMongoDBResponseBodyVpcsVSwitchs struct {
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// example:
	//
	// null
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// null
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// example:
	//
	// cn-hangzhou-i
	IzNo *string `json:"IzNo,omitempty" xml:"IzNo,omitempty"`
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// VSwitch ID。
	//
	// example:
	//
	// vsw-bp*******************
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// vsw-name
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
}

func (s DescribeVpcsForMongoDBResponseBodyVpcsVSwitchs) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsForMongoDBResponseBodyVpcsVSwitchs) GoString() string {
	return s.String()
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcsVSwitchs) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcsVSwitchs) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcsVSwitchs) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcsVSwitchs) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcsVSwitchs) GetIzNo() *string {
	return s.IzNo
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcsVSwitchs) GetStatus() *string {
	return s.Status
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcsVSwitchs) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcsVSwitchs) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcsVSwitchs) SetCidrBlock(v string) *DescribeVpcsForMongoDBResponseBodyVpcsVSwitchs {
	s.CidrBlock = &v
	return s
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcsVSwitchs) SetGmtCreate(v string) *DescribeVpcsForMongoDBResponseBodyVpcsVSwitchs {
	s.GmtCreate = &v
	return s
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcsVSwitchs) SetGmtModified(v string) *DescribeVpcsForMongoDBResponseBodyVpcsVSwitchs {
	s.GmtModified = &v
	return s
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcsVSwitchs) SetIsDefault(v bool) *DescribeVpcsForMongoDBResponseBodyVpcsVSwitchs {
	s.IsDefault = &v
	return s
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcsVSwitchs) SetIzNo(v string) *DescribeVpcsForMongoDBResponseBodyVpcsVSwitchs {
	s.IzNo = &v
	return s
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcsVSwitchs) SetStatus(v string) *DescribeVpcsForMongoDBResponseBodyVpcsVSwitchs {
	s.Status = &v
	return s
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcsVSwitchs) SetVSwitchId(v string) *DescribeVpcsForMongoDBResponseBodyVpcsVSwitchs {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcsVSwitchs) SetVSwitchName(v string) *DescribeVpcsForMongoDBResponseBodyVpcsVSwitchs {
	s.VSwitchName = &v
	return s
}

func (s *DescribeVpcsForMongoDBResponseBodyVpcsVSwitchs) Validate() error {
	return dara.Validate(s)
}
