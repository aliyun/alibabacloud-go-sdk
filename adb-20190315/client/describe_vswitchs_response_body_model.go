// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVSwitchsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVSwitchsResponseBody
	GetRequestId() *string
	SetVSwitches(v *DescribeVSwitchsResponseBodyVSwitches) *DescribeVSwitchsResponseBody
	GetVSwitches() *DescribeVSwitchsResponseBodyVSwitches
}

type DescribeVSwitchsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1B892DA2-9ABC-5CC0-AD73-405479C3FA53
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried vSwitches.
	VSwitches *DescribeVSwitchsResponseBodyVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Struct"`
}

func (s DescribeVSwitchsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVSwitchsResponseBody) GetVSwitches() *DescribeVSwitchsResponseBodyVSwitches {
	return s.VSwitches
}

func (s *DescribeVSwitchsResponseBody) SetRequestId(v string) *DescribeVSwitchsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVSwitchsResponseBody) SetVSwitches(v *DescribeVSwitchsResponseBodyVSwitches) *DescribeVSwitchsResponseBody {
	s.VSwitches = v
	return s
}

func (s *DescribeVSwitchsResponseBody) Validate() error {
	if s.VSwitches != nil {
		if err := s.VSwitches.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVSwitchsResponseBodyVSwitches struct {
	// The queried vSwitch.
	VSwitch []*DescribeVSwitchsResponseBodyVSwitchesVSwitch `json:"VSwitch,omitempty" xml:"VSwitch,omitempty" type:"Repeated"`
}

func (s DescribeVSwitchsResponseBodyVSwitches) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchsResponseBodyVSwitches) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchsResponseBodyVSwitches) GetVSwitch() []*DescribeVSwitchsResponseBodyVSwitchesVSwitch {
	return s.VSwitch
}

func (s *DescribeVSwitchsResponseBodyVSwitches) SetVSwitch(v []*DescribeVSwitchsResponseBodyVSwitchesVSwitch) *DescribeVSwitchsResponseBodyVSwitches {
	s.VSwitch = v
	return s
}

func (s *DescribeVSwitchsResponseBodyVSwitches) Validate() error {
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

type DescribeVSwitchsResponseBodyVSwitchesVSwitch struct {
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
	// 26842
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
	// true
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
	// VSW-Test-hangzhou-I
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
}

func (s DescribeVSwitchsResponseBodyVSwitchesVSwitch) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchsResponseBodyVSwitchesVSwitch) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchsResponseBodyVSwitchesVSwitch) GetAliUid() *string {
	return s.AliUid
}

func (s *DescribeVSwitchsResponseBodyVSwitchesVSwitch) GetBid() *string {
	return s.Bid
}

func (s *DescribeVSwitchsResponseBodyVSwitchesVSwitch) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeVSwitchsResponseBodyVSwitchesVSwitch) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeVSwitchsResponseBodyVSwitchesVSwitch) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeVSwitchsResponseBodyVSwitchesVSwitch) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeVSwitchsResponseBodyVSwitchesVSwitch) GetIzNo() *string {
	return s.IzNo
}

func (s *DescribeVSwitchsResponseBodyVSwitchesVSwitch) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeVSwitchsResponseBodyVSwitchesVSwitch) GetStatus() *string {
	return s.Status
}

func (s *DescribeVSwitchsResponseBodyVSwitchesVSwitch) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeVSwitchsResponseBodyVSwitchesVSwitch) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *DescribeVSwitchsResponseBodyVSwitchesVSwitch) SetAliUid(v string) *DescribeVSwitchsResponseBodyVSwitchesVSwitch {
	s.AliUid = &v
	return s
}

func (s *DescribeVSwitchsResponseBodyVSwitchesVSwitch) SetBid(v string) *DescribeVSwitchsResponseBodyVSwitchesVSwitch {
	s.Bid = &v
	return s
}

func (s *DescribeVSwitchsResponseBodyVSwitchesVSwitch) SetCidrBlock(v string) *DescribeVSwitchsResponseBodyVSwitchesVSwitch {
	s.CidrBlock = &v
	return s
}

func (s *DescribeVSwitchsResponseBodyVSwitchesVSwitch) SetGmtCreate(v string) *DescribeVSwitchsResponseBodyVSwitchesVSwitch {
	s.GmtCreate = &v
	return s
}

func (s *DescribeVSwitchsResponseBodyVSwitchesVSwitch) SetGmtModified(v string) *DescribeVSwitchsResponseBodyVSwitchesVSwitch {
	s.GmtModified = &v
	return s
}

func (s *DescribeVSwitchsResponseBodyVSwitchesVSwitch) SetIsDefault(v bool) *DescribeVSwitchsResponseBodyVSwitchesVSwitch {
	s.IsDefault = &v
	return s
}

func (s *DescribeVSwitchsResponseBodyVSwitchesVSwitch) SetIzNo(v string) *DescribeVSwitchsResponseBodyVSwitchesVSwitch {
	s.IzNo = &v
	return s
}

func (s *DescribeVSwitchsResponseBodyVSwitchesVSwitch) SetRegionNo(v string) *DescribeVSwitchsResponseBodyVSwitchesVSwitch {
	s.RegionNo = &v
	return s
}

func (s *DescribeVSwitchsResponseBodyVSwitchesVSwitch) SetStatus(v string) *DescribeVSwitchsResponseBodyVSwitchesVSwitch {
	s.Status = &v
	return s
}

func (s *DescribeVSwitchsResponseBodyVSwitchesVSwitch) SetVSwitchId(v string) *DescribeVSwitchsResponseBodyVSwitchesVSwitch {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVSwitchsResponseBodyVSwitchesVSwitch) SetVSwitchName(v string) *DescribeVSwitchsResponseBodyVSwitchesVSwitch {
	s.VSwitchName = &v
	return s
}

func (s *DescribeVSwitchsResponseBodyVSwitchesVSwitch) Validate() error {
	return dara.Validate(s)
}
