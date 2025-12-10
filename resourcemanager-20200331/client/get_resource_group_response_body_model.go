// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetResourceGroupResponseBody
	GetRequestId() *string
	SetResourceGroup(v *GetResourceGroupResponseBodyResourceGroup) *GetResourceGroupResponseBody
	GetResourceGroup() *GetResourceGroupResponseBodyResourceGroup
}

type GetResourceGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2D69A58F-345C-4FDE-88E4-BF5189484043
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the resource group.
	ResourceGroup *GetResourceGroupResponseBodyResourceGroup `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty" type:"Struct"`
}

func (s GetResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceGroupResponseBody) GetResourceGroup() *GetResourceGroupResponseBodyResourceGroup {
	return s.ResourceGroup
}

func (s *GetResourceGroupResponseBody) SetRequestId(v string) *GetResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceGroupResponseBody) SetResourceGroup(v *GetResourceGroupResponseBodyResourceGroup) *GetResourceGroupResponseBody {
	s.ResourceGroup = v
	return s
}

func (s *GetResourceGroupResponseBody) Validate() error {
	if s.ResourceGroup != nil {
		if err := s.ResourceGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResourceGroupResponseBodyResourceGroup struct {
	// The ID of the Alibaba Cloud account to which the resource group belongs.
	//
	// example:
	//
	// 123456789****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The time when the resource group was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The display name of the resource group.
	//
	// example:
	//
	// my-project
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-9gLOoK****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The identifier of the resource group.
	//
	// example:
	//
	// my-project
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the resource group in all regions.
	RegionStatuses *GetResourceGroupResponseBodyResourceGroupRegionStatuses `json:"RegionStatuses,omitempty" xml:"RegionStatuses,omitempty" type:"Struct"`
	// The status of the resource group. Valid values:
	//
	// 	- Creating: The resource group is being created.
	//
	// 	- OK: The resource group is created.
	//
	// 	- PendingDelete: The resource group is waiting to be deleted.
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags that are added to the resource group.
	Tags *GetResourceGroupResponseBodyResourceGroupTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s GetResourceGroupResponseBodyResourceGroup) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupResponseBodyResourceGroup) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResponseBodyResourceGroup) GetAccountId() *string {
	return s.AccountId
}

func (s *GetResourceGroupResponseBodyResourceGroup) GetCreateDate() *string {
	return s.CreateDate
}

func (s *GetResourceGroupResponseBodyResourceGroup) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetResourceGroupResponseBodyResourceGroup) GetId() *string {
	return s.Id
}

func (s *GetResourceGroupResponseBodyResourceGroup) GetName() *string {
	return s.Name
}

func (s *GetResourceGroupResponseBodyResourceGroup) GetRegionStatuses() *GetResourceGroupResponseBodyResourceGroupRegionStatuses {
	return s.RegionStatuses
}

func (s *GetResourceGroupResponseBodyResourceGroup) GetStatus() *string {
	return s.Status
}

func (s *GetResourceGroupResponseBodyResourceGroup) GetTags() *GetResourceGroupResponseBodyResourceGroupTags {
	return s.Tags
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetAccountId(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.AccountId = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetCreateDate(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.CreateDate = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetDisplayName(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.DisplayName = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetId(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.Id = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetName(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.Name = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetRegionStatuses(v *GetResourceGroupResponseBodyResourceGroupRegionStatuses) *GetResourceGroupResponseBodyResourceGroup {
	s.RegionStatuses = v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetStatus(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.Status = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetTags(v *GetResourceGroupResponseBodyResourceGroupTags) *GetResourceGroupResponseBodyResourceGroup {
	s.Tags = v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) Validate() error {
	if s.RegionStatuses != nil {
		if err := s.RegionStatuses.Validate(); err != nil {
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

type GetResourceGroupResponseBodyResourceGroupRegionStatuses struct {
	RegionStatus []*GetResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus `json:"RegionStatus,omitempty" xml:"RegionStatus,omitempty" type:"Repeated"`
}

func (s GetResourceGroupResponseBodyResourceGroupRegionStatuses) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupResponseBodyResourceGroupRegionStatuses) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResponseBodyResourceGroupRegionStatuses) GetRegionStatus() []*GetResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus {
	return s.RegionStatus
}

func (s *GetResourceGroupResponseBodyResourceGroupRegionStatuses) SetRegionStatus(v []*GetResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) *GetResourceGroupResponseBodyResourceGroupRegionStatuses {
	s.RegionStatus = v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroupRegionStatuses) Validate() error {
	if s.RegionStatus != nil {
		for _, item := range s.RegionStatus {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus struct {
	// The region ID.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the resource group. Valid values:
	//
	// 	- Creating: The resource group is being created.
	//
	// 	- OK: The resource group is created.
	//
	// 	- PendingDelete: The resource group is waiting to be deleted.
	//
	// 	- Deleting: The resource group is being deleted.
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) GetRegionId() *string {
	return s.RegionId
}

func (s *GetResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) GetStatus() *string {
	return s.Status
}

func (s *GetResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) SetRegionId(v string) *GetResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus {
	s.RegionId = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) SetStatus(v string) *GetResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus {
	s.Status = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) Validate() error {
	return dara.Validate(s)
}

type GetResourceGroupResponseBodyResourceGroupTags struct {
	Tag []*GetResourceGroupResponseBodyResourceGroupTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s GetResourceGroupResponseBodyResourceGroupTags) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupResponseBodyResourceGroupTags) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResponseBodyResourceGroupTags) GetTag() []*GetResourceGroupResponseBodyResourceGroupTagsTag {
	return s.Tag
}

func (s *GetResourceGroupResponseBodyResourceGroupTags) SetTag(v []*GetResourceGroupResponseBodyResourceGroupTagsTag) *GetResourceGroupResponseBodyResourceGroupTags {
	s.Tag = v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroupTags) Validate() error {
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

type GetResourceGroupResponseBodyResourceGroupTagsTag struct {
	// The tag key.
	//
	// example:
	//
	// k1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// v1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetResourceGroupResponseBodyResourceGroupTagsTag) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupResponseBodyResourceGroupTagsTag) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResponseBodyResourceGroupTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *GetResourceGroupResponseBodyResourceGroupTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *GetResourceGroupResponseBodyResourceGroupTagsTag) SetTagKey(v string) *GetResourceGroupResponseBodyResourceGroupTagsTag {
	s.TagKey = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroupTagsTag) SetTagValue(v string) *GetResourceGroupResponseBodyResourceGroupTagsTag {
	s.TagValue = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroupTagsTag) Validate() error {
	return dara.Validate(s)
}
