// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVSwitchListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeVSwitchListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVSwitchListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeVSwitchListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVSwitchListResponseBody
	GetTotalCount() *int32
	SetVSwitchs(v []*DescribeVSwitchListResponseBodyVSwitchs) *DescribeVSwitchListResponseBody
	GetVSwitchs() []*DescribeVSwitchListResponseBodyVSwitchs
}

type DescribeVSwitchListResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 6A2EE5B4-CC9F-46E1-A747-E43BC9******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VSwitchs   []*DescribeVSwitchListResponseBodyVSwitchs `json:"VSwitchs,omitempty" xml:"VSwitchs,omitempty" type:"Repeated"`
}

func (s DescribeVSwitchListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVSwitchListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVSwitchListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVSwitchListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVSwitchListResponseBody) GetVSwitchs() []*DescribeVSwitchListResponseBodyVSwitchs {
	return s.VSwitchs
}

func (s *DescribeVSwitchListResponseBody) SetPageNumber(v int32) *DescribeVSwitchListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVSwitchListResponseBody) SetPageSize(v int32) *DescribeVSwitchListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVSwitchListResponseBody) SetRequestId(v string) *DescribeVSwitchListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVSwitchListResponseBody) SetTotalCount(v int32) *DescribeVSwitchListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVSwitchListResponseBody) SetVSwitchs(v []*DescribeVSwitchListResponseBodyVSwitchs) *DescribeVSwitchListResponseBody {
	s.VSwitchs = v
	return s
}

func (s *DescribeVSwitchListResponseBody) Validate() error {
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

type DescribeVSwitchListResponseBodyVSwitchs struct {
	// example:
	//
	// 1
	AvailableIpAddressCount *int64 `json:"AvailableIpAddressCount,omitempty" xml:"AvailableIpAddressCount,omitempty"`
	// example:
	//
	// 172.16.0.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// example:
	//
	// vSwitchDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// example:
	//
	// cn-hangzhou-b
	IzNo *string `json:"IzNo,omitempty" xml:"IzNo,omitempty"`
	// example:
	//
	// 177563751276****
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// rg-************
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// Public
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// vsw-25bcdxs7pv1****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// vSwitch
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	// example:
	//
	// vpc-bp1vbkkyt7apvy4j*****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeVSwitchListResponseBodyVSwitchs) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchListResponseBodyVSwitchs) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchListResponseBodyVSwitchs) GetAvailableIpAddressCount() *int64 {
	return s.AvailableIpAddressCount
}

func (s *DescribeVSwitchListResponseBodyVSwitchs) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeVSwitchListResponseBodyVSwitchs) GetDescription() *string {
	return s.Description
}

func (s *DescribeVSwitchListResponseBodyVSwitchs) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeVSwitchListResponseBodyVSwitchs) GetIzNo() *string {
	return s.IzNo
}

func (s *DescribeVSwitchListResponseBodyVSwitchs) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeVSwitchListResponseBodyVSwitchs) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeVSwitchListResponseBodyVSwitchs) GetShareType() *string {
	return s.ShareType
}

func (s *DescribeVSwitchListResponseBodyVSwitchs) GetStatus() *string {
	return s.Status
}

func (s *DescribeVSwitchListResponseBodyVSwitchs) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeVSwitchListResponseBodyVSwitchs) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *DescribeVSwitchListResponseBodyVSwitchs) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVSwitchListResponseBodyVSwitchs) SetAvailableIpAddressCount(v int64) *DescribeVSwitchListResponseBodyVSwitchs {
	s.AvailableIpAddressCount = &v
	return s
}

func (s *DescribeVSwitchListResponseBodyVSwitchs) SetCidrBlock(v string) *DescribeVSwitchListResponseBodyVSwitchs {
	s.CidrBlock = &v
	return s
}

func (s *DescribeVSwitchListResponseBodyVSwitchs) SetDescription(v string) *DescribeVSwitchListResponseBodyVSwitchs {
	s.Description = &v
	return s
}

func (s *DescribeVSwitchListResponseBodyVSwitchs) SetIsDefault(v bool) *DescribeVSwitchListResponseBodyVSwitchs {
	s.IsDefault = &v
	return s
}

func (s *DescribeVSwitchListResponseBodyVSwitchs) SetIzNo(v string) *DescribeVSwitchListResponseBodyVSwitchs {
	s.IzNo = &v
	return s
}

func (s *DescribeVSwitchListResponseBodyVSwitchs) SetOwnerId(v string) *DescribeVSwitchListResponseBodyVSwitchs {
	s.OwnerId = &v
	return s
}

func (s *DescribeVSwitchListResponseBodyVSwitchs) SetResourceGroupId(v string) *DescribeVSwitchListResponseBodyVSwitchs {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVSwitchListResponseBodyVSwitchs) SetShareType(v string) *DescribeVSwitchListResponseBodyVSwitchs {
	s.ShareType = &v
	return s
}

func (s *DescribeVSwitchListResponseBodyVSwitchs) SetStatus(v string) *DescribeVSwitchListResponseBodyVSwitchs {
	s.Status = &v
	return s
}

func (s *DescribeVSwitchListResponseBodyVSwitchs) SetVSwitchId(v string) *DescribeVSwitchListResponseBodyVSwitchs {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVSwitchListResponseBodyVSwitchs) SetVSwitchName(v string) *DescribeVSwitchListResponseBodyVSwitchs {
	s.VSwitchName = &v
	return s
}

func (s *DescribeVSwitchListResponseBodyVSwitchs) SetVpcId(v string) *DescribeVSwitchListResponseBodyVSwitchs {
	s.VpcId = &v
	return s
}

func (s *DescribeVSwitchListResponseBodyVSwitchs) Validate() error {
	return dara.Validate(s)
}
