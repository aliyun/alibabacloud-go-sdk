// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRdsVSwitchsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRdsVSwitchsResponseBody
	GetRequestId() *string
	SetVSwitches(v *DescribeRdsVSwitchsResponseBodyVSwitches) *DescribeRdsVSwitchsResponseBody
	GetVSwitches() *DescribeRdsVSwitchsResponseBodyVSwitches
}

type DescribeRdsVSwitchsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// B0BA251D-6854-5A99-A964-3C**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details of the vSwitches.
	VSwitches *DescribeRdsVSwitchsResponseBodyVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Struct"`
}

func (s DescribeRdsVSwitchsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsVSwitchsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRdsVSwitchsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRdsVSwitchsResponseBody) GetVSwitches() *DescribeRdsVSwitchsResponseBodyVSwitches {
	return s.VSwitches
}

func (s *DescribeRdsVSwitchsResponseBody) SetRequestId(v string) *DescribeRdsVSwitchsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBody) SetVSwitches(v *DescribeRdsVSwitchsResponseBodyVSwitches) *DescribeRdsVSwitchsResponseBody {
	s.VSwitches = v
	return s
}

func (s *DescribeRdsVSwitchsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRdsVSwitchsResponseBodyVSwitches struct {
	// Details of the vSwitch.
	VSwitch []*DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch `json:"VSwitch,omitempty" xml:"VSwitch,omitempty" type:"Repeated"`
}

func (s DescribeRdsVSwitchsResponseBodyVSwitches) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsVSwitchsResponseBodyVSwitches) GoString() string {
	return s.String()
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitches) GetVSwitch() []*DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	return s.VSwitch
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitches) SetVSwitch(v []*DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) *DescribeRdsVSwitchsResponseBodyVSwitches {
	s.VSwitch = v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitches) Validate() error {
	return dara.Validate(s)
}

type DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch struct {
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
	// The CIDR block of the vSwitch.
	//
	// example:
	//
	// 192.**.**.0/24
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
	// false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The ID of the zone.
	//
	// example:
	//
	// cn-hangzhou-h
	IzNo *string `json:"IzNo,omitempty" xml:"IzNo,omitempty"`
	// An invalid parameter. It is no longer returned when you call this operation.
	//
	// example:
	//
	// null
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The state of the vSwitch. If **Available*	- is returned, the vSwitch is available.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the vSwitch.
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

func (s DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) GoString() string {
	return s.String()
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) GetAliUid() *string {
	return s.AliUid
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) GetBid() *string {
	return s.Bid
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) GetIzNo() *string {
	return s.IzNo
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) GetStatus() *string {
	return s.Status
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetAliUid(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.AliUid = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetBid(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.Bid = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetCidrBlock(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.CidrBlock = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetGmtCreate(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.GmtCreate = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetGmtModified(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.GmtModified = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetIsDefault(v bool) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.IsDefault = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetIzNo(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.IzNo = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetRegionNo(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.RegionNo = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetStatus(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.Status = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetVSwitchId(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.VSwitchId = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) SetVSwitchName(v string) *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch {
	s.VSwitchName = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseBodyVSwitchesVSwitch) Validate() error {
	return dara.Validate(s)
}
