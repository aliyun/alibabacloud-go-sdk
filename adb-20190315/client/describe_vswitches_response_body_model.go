// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVSwitchesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVSwitchesResponseBody
	GetRequestId() *string
	SetVSwitches(v *DescribeVSwitchesResponseBodyVSwitches) *DescribeVSwitchesResponseBody
	GetVSwitches() *DescribeVSwitchesResponseBodyVSwitches
}

type DescribeVSwitchesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D65A809F-34CE-4550-9BC1-0ED21ETG380
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried vSwitches.
	VSwitches *DescribeVSwitchesResponseBodyVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Struct"`
}

func (s DescribeVSwitchesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVSwitchesResponseBody) GetVSwitches() *DescribeVSwitchesResponseBodyVSwitches {
	return s.VSwitches
}

func (s *DescribeVSwitchesResponseBody) SetRequestId(v string) *DescribeVSwitchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetVSwitches(v *DescribeVSwitchesResponseBodyVSwitches) *DescribeVSwitchesResponseBody {
	s.VSwitches = v
	return s
}

func (s *DescribeVSwitchesResponseBody) Validate() error {
	if s.VSwitches != nil {
		if err := s.VSwitches.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVSwitchesResponseBodyVSwitches struct {
	// The queried vSwitch.
	VSwitch []*DescribeVSwitchesResponseBodyVSwitchesVSwitch `json:"VSwitch,omitempty" xml:"VSwitch,omitempty" type:"Repeated"`
}

func (s DescribeVSwitchesResponseBodyVSwitches) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchesResponseBodyVSwitches) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBodyVSwitches) GetVSwitch() []*DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	return s.VSwitch
}

func (s *DescribeVSwitchesResponseBodyVSwitches) SetVSwitch(v []*DescribeVSwitchesResponseBodyVSwitchesVSwitch) *DescribeVSwitchesResponseBodyVSwitches {
	s.VSwitch = v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitches) Validate() error {
	if s.VSwitch != nil {
		for _, item := range s.VSwitch {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVSwitchesResponseBodyVSwitchesVSwitch struct {
	// The ID of the Resource Access Management (RAM) user.
	//
	// example:
	//
	// 195813423043****
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The ID of the user channel.
	//
	// example:
	//
	// 2****
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// The IPv4 CIDR block of the vSwitch.
	//
	// example:
	//
	// 172.16.0.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The time when the vSwitch was created.
	//
	// example:
	//
	// 2022-01-18T12:43:57Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the vSwitch was modified.
	//
	// example:
	//
	// 2022-01-22T12:43:57Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates whether the vSwitch is the default vSwitch. Valid values: **true**: The vSwitch is the default vSwitch. **false**: The vSwitch is not the default vSwitch.
	//
	// example:
	//
	// false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The zone ID of the vSwitch.
	//
	// example:
	//
	// cn-hangzhou-k
	IzNo *string `json:"IzNo,omitempty" xml:"IzNo,omitempty"`
	// The region ID of the vSwitch.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The state of the vSwitch. Valid values: **Pending**: the vSwitch is being configured. **Available**: the vSwitch is available.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-25bcdxs7pv1****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The name of the vSwitch.
	//
	// example:
	//
	// vswitch
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
}

func (s DescribeVSwitchesResponseBodyVSwitchesVSwitch) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchesResponseBodyVSwitchesVSwitch) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetAliUid() *string {
	return s.AliUid
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetBid() *string {
	return s.Bid
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetIzNo() *string {
	return s.IzNo
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetStatus() *string {
	return s.Status
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetAliUid(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.AliUid = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetBid(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.Bid = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetCidrBlock(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.CidrBlock = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetGmtCreate(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.GmtCreate = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetGmtModified(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.GmtModified = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetIsDefault(v bool) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.IsDefault = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetIzNo(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.IzNo = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetRegionNo(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.RegionNo = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetStatus(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.Status = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetVSwitchId(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetVSwitchName(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.VSwitchName = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) Validate() error {
	return dara.Validate(s)
}
