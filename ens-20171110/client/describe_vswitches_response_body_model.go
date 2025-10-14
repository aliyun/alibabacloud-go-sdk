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
	SetVSwitches(v *DescribeVSwitchesResponseBodyVSwitches) *DescribeVSwitchesResponseBody
	GetVSwitches() *DescribeVSwitchesResponseBodyVSwitches
}

type DescribeVSwitchesResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The information about the vSwitches. For more information, see the array of vSwitches in the response examples in the JSON format.
	VSwitches *DescribeVSwitchesResponseBodyVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Struct"`
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

func (s *DescribeVSwitchesResponseBody) GetVSwitches() *DescribeVSwitchesResponseBodyVSwitches {
	return s.VSwitches
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
	// The IPv4 CIDR block of the vSwitch.
	//
	// example:
	//
	// 10.21.224.0/22
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The time when the VPC was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-06-16T06:33:15Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The description of the vSwitch.
	//
	// example:
	//
	// VSwitchDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the ENS node.
	//
	// example:
	//
	// cn-xian-unicom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The number of available IP addresses.
	//
	// example:
	//
	// 1024
	FreeIpCount *int64 `json:"FreeIpCount,omitempty" xml:"FreeIpCount,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-25cdvfeq58pl****
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The status of the vSwitch. Valid values:
	//
	// 	- Pending
	//
	// 	- Available
	//
	// example:
	//
	// Pending
	Status *string                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   *DescribeVSwitchesResponseBodyVSwitchesVSwitchTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-5m9xhlq8oh***
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The name of the vSwitch.
	//
	// example:
	//
	// testVSwitchName
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
}

func (s DescribeVSwitchesResponseBodyVSwitchesVSwitch) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchesResponseBodyVSwitchesVSwitch) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetDescription() *string {
	return s.Description
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetFreeIpCount() *int64 {
	return s.FreeIpCount
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetNetworkId() *string {
	return s.NetworkId
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetStatus() *string {
	return s.Status
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetTags() *DescribeVSwitchesResponseBodyVSwitchesVSwitchTags {
	return s.Tags
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetCidrBlock(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.CidrBlock = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetCreatedTime(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.CreatedTime = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetDescription(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.Description = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetEnsRegionId(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetFreeIpCount(v int64) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.FreeIpCount = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetNetworkId(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.NetworkId = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetStatus(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.Status = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitch) SetTags(v *DescribeVSwitchesResponseBodyVSwitchesVSwitchTags) *DescribeVSwitchesResponseBodyVSwitchesVSwitch {
	s.Tags = v
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
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVSwitchesResponseBodyVSwitchesVSwitchTags struct {
	Tag []*DescribeVSwitchesResponseBodyVSwitchesVSwitchTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeVSwitchesResponseBodyVSwitchesVSwitchTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchesResponseBodyVSwitchesVSwitchTags) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitchTags) GetTag() []*DescribeVSwitchesResponseBodyVSwitchesVSwitchTagsTag {
	return s.Tag
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitchTags) SetTag(v []*DescribeVSwitchesResponseBodyVSwitchesVSwitchTagsTag) *DescribeVSwitchesResponseBodyVSwitchesVSwitchTags {
	s.Tag = v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitchTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVSwitchesResponseBodyVSwitchesVSwitchTagsTag struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Deprecated
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// Deprecated
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	// The request error rate.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVSwitchesResponseBodyVSwitchesVSwitchTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchesResponseBodyVSwitchesVSwitchTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitchTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitchTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitchTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitchTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitchTagsTag) SetKey(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitchTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitchTagsTag) SetTagKey(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitchTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitchTagsTag) SetTagValue(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitchTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitchTagsTag) SetValue(v string) *DescribeVSwitchesResponseBodyVSwitchesVSwitchTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeVSwitchesResponseBodyVSwitchesVSwitchTagsTag) Validate() error {
	return dara.Validate(s)
}
