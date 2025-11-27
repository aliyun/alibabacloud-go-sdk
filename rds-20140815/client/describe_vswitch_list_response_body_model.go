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
	SetVSwitches(v []*DescribeVSwitchListResponseBodyVSwitches) *DescribeVSwitchListResponseBody
	GetVSwitches() []*DescribeVSwitchListResponseBodyVSwitches
}

type DescribeVSwitchListResponseBody struct {
	PageNumber *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VSwitches  []*DescribeVSwitchListResponseBodyVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
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

func (s *DescribeVSwitchListResponseBody) GetVSwitches() []*DescribeVSwitchListResponseBodyVSwitches {
	return s.VSwitches
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

func (s *DescribeVSwitchListResponseBody) SetVSwitches(v []*DescribeVSwitchListResponseBodyVSwitches) *DescribeVSwitchListResponseBody {
	s.VSwitches = v
	return s
}

func (s *DescribeVSwitchListResponseBody) Validate() error {
	if s.VSwitches != nil {
		for _, item := range s.VSwitches {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVSwitchListResponseBodyVSwitches struct {
	AvailableIpAddressCount *string `json:"AvailableIpAddressCount,omitempty" xml:"AvailableIpAddressCount,omitempty"`
	CidrBlock               *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate               *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	IsDefault               *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	IzNo                    *string `json:"IzNo,omitempty" xml:"IzNo,omitempty"`
	ShareType               *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VSwitchId               *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VSwitchName             *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
}

func (s DescribeVSwitchListResponseBodyVSwitches) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchListResponseBodyVSwitches) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchListResponseBodyVSwitches) GetAvailableIpAddressCount() *string {
	return s.AvailableIpAddressCount
}

func (s *DescribeVSwitchListResponseBodyVSwitches) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeVSwitchListResponseBodyVSwitches) GetDescription() *string {
	return s.Description
}

func (s *DescribeVSwitchListResponseBodyVSwitches) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeVSwitchListResponseBodyVSwitches) GetIsDefault() *string {
	return s.IsDefault
}

func (s *DescribeVSwitchListResponseBodyVSwitches) GetIzNo() *string {
	return s.IzNo
}

func (s *DescribeVSwitchListResponseBodyVSwitches) GetShareType() *string {
	return s.ShareType
}

func (s *DescribeVSwitchListResponseBodyVSwitches) GetStatus() *string {
	return s.Status
}

func (s *DescribeVSwitchListResponseBodyVSwitches) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeVSwitchListResponseBodyVSwitches) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *DescribeVSwitchListResponseBodyVSwitches) SetAvailableIpAddressCount(v string) *DescribeVSwitchListResponseBodyVSwitches {
	s.AvailableIpAddressCount = &v
	return s
}

func (s *DescribeVSwitchListResponseBodyVSwitches) SetCidrBlock(v string) *DescribeVSwitchListResponseBodyVSwitches {
	s.CidrBlock = &v
	return s
}

func (s *DescribeVSwitchListResponseBodyVSwitches) SetDescription(v string) *DescribeVSwitchListResponseBodyVSwitches {
	s.Description = &v
	return s
}

func (s *DescribeVSwitchListResponseBodyVSwitches) SetGmtCreate(v string) *DescribeVSwitchListResponseBodyVSwitches {
	s.GmtCreate = &v
	return s
}

func (s *DescribeVSwitchListResponseBodyVSwitches) SetIsDefault(v string) *DescribeVSwitchListResponseBodyVSwitches {
	s.IsDefault = &v
	return s
}

func (s *DescribeVSwitchListResponseBodyVSwitches) SetIzNo(v string) *DescribeVSwitchListResponseBodyVSwitches {
	s.IzNo = &v
	return s
}

func (s *DescribeVSwitchListResponseBodyVSwitches) SetShareType(v string) *DescribeVSwitchListResponseBodyVSwitches {
	s.ShareType = &v
	return s
}

func (s *DescribeVSwitchListResponseBodyVSwitches) SetStatus(v string) *DescribeVSwitchListResponseBodyVSwitches {
	s.Status = &v
	return s
}

func (s *DescribeVSwitchListResponseBodyVSwitches) SetVSwitchId(v string) *DescribeVSwitchListResponseBodyVSwitches {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVSwitchListResponseBodyVSwitches) SetVSwitchName(v string) *DescribeVSwitchListResponseBodyVSwitches {
	s.VSwitchName = &v
	return s
}

func (s *DescribeVSwitchListResponseBodyVSwitches) Validate() error {
	return dara.Validate(s)
}
