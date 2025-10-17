// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVSwitchesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeVSwitchesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVSwitchesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeVSwitchesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVSwitchesResponseBody
	GetTotalCount() *int32
	SetVSwitchs(v []*DescribeVSwitchesResponseBodyVSwitchs) *DescribeVSwitchesResponseBody
	GetVSwitchs() []*DescribeVSwitchesResponseBodyVSwitchs
}

type DescribeVSwitchesResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9A572171-4E27-40D1-BD36-D26C9E71E29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The vSwitches.
	VSwitchs []*DescribeVSwitchesResponseBodyVSwitchs `json:"VSwitchs,omitempty" xml:"VSwitchs,omitempty" type:"Repeated"`
}

func (s DescribeVSwitchesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVSwitchesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVSwitchesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVSwitchesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVSwitchesResponseBody) GetVSwitchs() []*DescribeVSwitchesResponseBodyVSwitchs {
	return s.VSwitchs
}

func (s *DescribeVSwitchesResponseBody) SetPageNumber(v int32) *DescribeVSwitchesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetPageSize(v int32) *DescribeVSwitchesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetRequestId(v string) *DescribeVSwitchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetTotalCount(v int32) *DescribeVSwitchesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVSwitchesResponseBody) SetVSwitchs(v []*DescribeVSwitchesResponseBodyVSwitchs) *DescribeVSwitchesResponseBody {
	s.VSwitchs = v
	return s
}

func (s *DescribeVSwitchesResponseBody) Validate() error {
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

type DescribeVSwitchesResponseBodyVSwitchs struct {
	// The number of available IP addresses in the vSwitch.
	//
	// example:
	//
	// 1
	AvailableIpAddressCount *int64 `json:"AvailableIpAddressCount,omitempty" xml:"AvailableIpAddressCount,omitempty"`
	// The IPv4 CIDR block of the vSwitch.
	//
	// example:
	//
	// 172.16.0.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The description of the vSwitch.
	//
	// example:
	//
	// vSwitchDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	// The zone to which the NAT gateway belongs.
	//
	// example:
	//
	// cn-hangzhou-b
	IzNo *string `json:"IzNo,omitempty" xml:"IzNo,omitempty"`
	// The status of the vSwitch. Valid values:
	//
	// 	- **Pending**: The vSwitch is being configured.
	//
	// 	- **Available**: The vSwitch is available.
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
	// vSwitch
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
}

func (s DescribeVSwitchesResponseBodyVSwitchs) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchesResponseBodyVSwitchs) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBodyVSwitchs) GetAvailableIpAddressCount() *int64 {
	return s.AvailableIpAddressCount
}

func (s *DescribeVSwitchesResponseBodyVSwitchs) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeVSwitchesResponseBodyVSwitchs) GetDescription() *string {
	return s.Description
}

func (s *DescribeVSwitchesResponseBodyVSwitchs) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeVSwitchesResponseBodyVSwitchs) GetIzNo() *string {
	return s.IzNo
}

func (s *DescribeVSwitchesResponseBodyVSwitchs) GetStatus() *string {
	return s.Status
}

func (s *DescribeVSwitchesResponseBodyVSwitchs) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeVSwitchesResponseBodyVSwitchs) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *DescribeVSwitchesResponseBodyVSwitchs) SetAvailableIpAddressCount(v int64) *DescribeVSwitchesResponseBodyVSwitchs {
	s.AvailableIpAddressCount = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchs) SetCidrBlock(v string) *DescribeVSwitchesResponseBodyVSwitchs {
	s.CidrBlock = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchs) SetDescription(v string) *DescribeVSwitchesResponseBodyVSwitchs {
	s.Description = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchs) SetIsDefault(v bool) *DescribeVSwitchesResponseBodyVSwitchs {
	s.IsDefault = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchs) SetIzNo(v string) *DescribeVSwitchesResponseBodyVSwitchs {
	s.IzNo = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchs) SetStatus(v string) *DescribeVSwitchesResponseBodyVSwitchs {
	s.Status = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchs) SetVSwitchId(v string) *DescribeVSwitchesResponseBodyVSwitchs {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchs) SetVSwitchName(v string) *DescribeVSwitchesResponseBodyVSwitchs {
	s.VSwitchName = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchs) Validate() error {
	return dara.Validate(s)
}
