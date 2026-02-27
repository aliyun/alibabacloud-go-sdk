// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHaVipsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHaVips(v *DescribeHaVipsResponseBodyHaVips) *DescribeHaVipsResponseBody
	GetHaVips() *DescribeHaVipsResponseBodyHaVips
	SetPageNumber(v int32) *DescribeHaVipsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHaVipsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeHaVipsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeHaVipsResponseBody
	GetTotalCount() *int32
}

type DescribeHaVipsResponseBody struct {
	HaVips *DescribeHaVipsResponseBodyHaVips `json:"HaVips,omitempty" xml:"HaVips,omitempty" type:"Struct"`
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
	// 33E480C5-B46F-4CA5-B6FD-D77C746E86AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHaVipsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHaVipsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHaVipsResponseBody) GetHaVips() *DescribeHaVipsResponseBodyHaVips {
	return s.HaVips
}

func (s *DescribeHaVipsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHaVipsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHaVipsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHaVipsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeHaVipsResponseBody) SetHaVips(v *DescribeHaVipsResponseBodyHaVips) *DescribeHaVipsResponseBody {
	s.HaVips = v
	return s
}

func (s *DescribeHaVipsResponseBody) SetPageNumber(v int32) *DescribeHaVipsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeHaVipsResponseBody) SetPageSize(v int32) *DescribeHaVipsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeHaVipsResponseBody) SetRequestId(v string) *DescribeHaVipsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHaVipsResponseBody) SetTotalCount(v int32) *DescribeHaVipsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeHaVipsResponseBody) Validate() error {
	if s.HaVips != nil {
		if err := s.HaVips.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeHaVipsResponseBodyHaVips struct {
	HaVip []*DescribeHaVipsResponseBodyHaVipsHaVip `json:"HaVip,omitempty" xml:"HaVip,omitempty" type:"Repeated"`
}

func (s DescribeHaVipsResponseBodyHaVips) String() string {
	return dara.Prettify(s)
}

func (s DescribeHaVipsResponseBodyHaVips) GoString() string {
	return s.String()
}

func (s *DescribeHaVipsResponseBodyHaVips) GetHaVip() []*DescribeHaVipsResponseBodyHaVipsHaVip {
	return s.HaVip
}

func (s *DescribeHaVipsResponseBodyHaVips) SetHaVip(v []*DescribeHaVipsResponseBodyHaVipsHaVip) *DescribeHaVipsResponseBodyHaVips {
	s.HaVip = v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVips) Validate() error {
	if s.HaVip != nil {
		for _, item := range s.HaVip {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHaVipsResponseBodyHaVipsHaVip struct {
	AssociatedEipAddresses *DescribeHaVipsResponseBodyHaVipsHaVipAssociatedEipAddresses `json:"AssociatedEipAddresses,omitempty" xml:"AssociatedEipAddresses,omitempty" type:"Struct"`
	AssociatedInstanceType *string                                                      `json:"AssociatedInstanceType,omitempty" xml:"AssociatedInstanceType,omitempty"`
	AssociatedInstances    *DescribeHaVipsResponseBodyHaVipsHaVipAssociatedInstances    `json:"AssociatedInstances,omitempty" xml:"AssociatedInstances,omitempty" type:"Struct"`
	ChargeType             *string                                                      `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CreateTime             *string                                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description            *string                                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	HaVipId                *string                                                      `json:"HaVipId,omitempty" xml:"HaVipId,omitempty"`
	IpAddress              *string                                                      `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	MasterInstanceId       *string                                                      `json:"MasterInstanceId,omitempty" xml:"MasterInstanceId,omitempty"`
	Name                   *string                                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId               *string                                                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId        *string                                                      `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Status                 *string                                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                   *DescribeHaVipsResponseBodyHaVipsHaVipTags                   `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	VSwitchId              *string                                                      `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                  *string                                                      `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeHaVipsResponseBodyHaVipsHaVip) String() string {
	return dara.Prettify(s)
}

func (s DescribeHaVipsResponseBodyHaVipsHaVip) GoString() string {
	return s.String()
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) GetAssociatedEipAddresses() *DescribeHaVipsResponseBodyHaVipsHaVipAssociatedEipAddresses {
	return s.AssociatedEipAddresses
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) GetAssociatedInstanceType() *string {
	return s.AssociatedInstanceType
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) GetAssociatedInstances() *DescribeHaVipsResponseBodyHaVipsHaVipAssociatedInstances {
	return s.AssociatedInstances
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) GetDescription() *string {
	return s.Description
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) GetHaVipId() *string {
	return s.HaVipId
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) GetIpAddress() *string {
	return s.IpAddress
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) GetMasterInstanceId() *string {
	return s.MasterInstanceId
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) GetName() *string {
	return s.Name
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) GetStatus() *string {
	return s.Status
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) GetTags() *DescribeHaVipsResponseBodyHaVipsHaVipTags {
	return s.Tags
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) SetAssociatedEipAddresses(v *DescribeHaVipsResponseBodyHaVipsHaVipAssociatedEipAddresses) *DescribeHaVipsResponseBodyHaVipsHaVip {
	s.AssociatedEipAddresses = v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) SetAssociatedInstanceType(v string) *DescribeHaVipsResponseBodyHaVipsHaVip {
	s.AssociatedInstanceType = &v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) SetAssociatedInstances(v *DescribeHaVipsResponseBodyHaVipsHaVipAssociatedInstances) *DescribeHaVipsResponseBodyHaVipsHaVip {
	s.AssociatedInstances = v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) SetChargeType(v string) *DescribeHaVipsResponseBodyHaVipsHaVip {
	s.ChargeType = &v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) SetCreateTime(v string) *DescribeHaVipsResponseBodyHaVipsHaVip {
	s.CreateTime = &v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) SetDescription(v string) *DescribeHaVipsResponseBodyHaVipsHaVip {
	s.Description = &v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) SetHaVipId(v string) *DescribeHaVipsResponseBodyHaVipsHaVip {
	s.HaVipId = &v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) SetIpAddress(v string) *DescribeHaVipsResponseBodyHaVipsHaVip {
	s.IpAddress = &v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) SetMasterInstanceId(v string) *DescribeHaVipsResponseBodyHaVipsHaVip {
	s.MasterInstanceId = &v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) SetName(v string) *DescribeHaVipsResponseBodyHaVipsHaVip {
	s.Name = &v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) SetRegionId(v string) *DescribeHaVipsResponseBodyHaVipsHaVip {
	s.RegionId = &v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) SetResourceGroupId(v string) *DescribeHaVipsResponseBodyHaVipsHaVip {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) SetStatus(v string) *DescribeHaVipsResponseBodyHaVipsHaVip {
	s.Status = &v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) SetTags(v *DescribeHaVipsResponseBodyHaVipsHaVipTags) *DescribeHaVipsResponseBodyHaVipsHaVip {
	s.Tags = v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) SetVSwitchId(v string) *DescribeHaVipsResponseBodyHaVipsHaVip {
	s.VSwitchId = &v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) SetVpcId(v string) *DescribeHaVipsResponseBodyHaVipsHaVip {
	s.VpcId = &v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVip) Validate() error {
	if s.AssociatedEipAddresses != nil {
		if err := s.AssociatedEipAddresses.Validate(); err != nil {
			return err
		}
	}
	if s.AssociatedInstances != nil {
		if err := s.AssociatedInstances.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeHaVipsResponseBodyHaVipsHaVipAssociatedEipAddresses struct {
	AssociatedEipAddresse []*string `json:"associatedEipAddresse,omitempty" xml:"associatedEipAddresse,omitempty" type:"Repeated"`
}

func (s DescribeHaVipsResponseBodyHaVipsHaVipAssociatedEipAddresses) String() string {
	return dara.Prettify(s)
}

func (s DescribeHaVipsResponseBodyHaVipsHaVipAssociatedEipAddresses) GoString() string {
	return s.String()
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVipAssociatedEipAddresses) GetAssociatedEipAddresse() []*string {
	return s.AssociatedEipAddresse
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVipAssociatedEipAddresses) SetAssociatedEipAddresse(v []*string) *DescribeHaVipsResponseBodyHaVipsHaVipAssociatedEipAddresses {
	s.AssociatedEipAddresse = v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVipAssociatedEipAddresses) Validate() error {
	return dara.Validate(s)
}

type DescribeHaVipsResponseBodyHaVipsHaVipAssociatedInstances struct {
	AssociatedInstance []*string `json:"associatedInstance,omitempty" xml:"associatedInstance,omitempty" type:"Repeated"`
}

func (s DescribeHaVipsResponseBodyHaVipsHaVipAssociatedInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeHaVipsResponseBodyHaVipsHaVipAssociatedInstances) GoString() string {
	return s.String()
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVipAssociatedInstances) GetAssociatedInstance() []*string {
	return s.AssociatedInstance
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVipAssociatedInstances) SetAssociatedInstance(v []*string) *DescribeHaVipsResponseBodyHaVipsHaVipAssociatedInstances {
	s.AssociatedInstance = v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVipAssociatedInstances) Validate() error {
	return dara.Validate(s)
}

type DescribeHaVipsResponseBodyHaVipsHaVipTags struct {
	Tag []*DescribeHaVipsResponseBodyHaVipsHaVipTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeHaVipsResponseBodyHaVipsHaVipTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeHaVipsResponseBodyHaVipsHaVipTags) GoString() string {
	return s.String()
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVipTags) GetTag() []*DescribeHaVipsResponseBodyHaVipsHaVipTagsTag {
	return s.Tag
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVipTags) SetTag(v []*DescribeHaVipsResponseBodyHaVipsHaVipTagsTag) *DescribeHaVipsResponseBodyHaVipsHaVipTags {
	s.Tag = v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVipTags) Validate() error {
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

type DescribeHaVipsResponseBodyHaVipsHaVipTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHaVipsResponseBodyHaVipsHaVipTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeHaVipsResponseBodyHaVipsHaVipTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVipTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVipTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVipTagsTag) SetKey(v string) *DescribeHaVipsResponseBodyHaVipsHaVipTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVipTagsTag) SetValue(v string) *DescribeHaVipsResponseBodyHaVipsHaVipTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVipsHaVipTagsTag) Validate() error {
	return dara.Validate(s)
}
