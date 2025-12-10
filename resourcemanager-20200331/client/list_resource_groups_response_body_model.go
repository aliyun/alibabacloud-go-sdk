// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListResourceGroupsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResourceGroupsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListResourceGroupsResponseBody
	GetRequestId() *string
	SetResourceGroups(v *ListResourceGroupsResponseBodyResourceGroups) *ListResourceGroupsResponseBody
	GetResourceGroups() *ListResourceGroupsResponseBodyResourceGroups
	SetTotalCount(v int32) *ListResourceGroupsResponseBody
	GetTotalCount() *int32
}

type ListResourceGroupsResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4B450CA1-36E8-4AA2-8461-86B42BF4CC4E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the resource groups.
	ResourceGroups *ListResourceGroupsResponseBodyResourceGroups `json:"ResourceGroups,omitempty" xml:"ResourceGroups,omitempty" type:"Struct"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourceGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourceGroupsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourceGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceGroupsResponseBody) GetResourceGroups() *ListResourceGroupsResponseBodyResourceGroups {
	return s.ResourceGroups
}

func (s *ListResourceGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListResourceGroupsResponseBody) SetPageNumber(v int32) *ListResourceGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListResourceGroupsResponseBody) SetPageSize(v int32) *ListResourceGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListResourceGroupsResponseBody) SetRequestId(v string) *ListResourceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceGroupsResponseBody) SetResourceGroups(v *ListResourceGroupsResponseBodyResourceGroups) *ListResourceGroupsResponseBody {
	s.ResourceGroups = v
	return s
}

func (s *ListResourceGroupsResponseBody) SetTotalCount(v int32) *ListResourceGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListResourceGroupsResponseBody) Validate() error {
	if s.ResourceGroups != nil {
		if err := s.ResourceGroups.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListResourceGroupsResponseBodyResourceGroups struct {
	ResourceGroup []*ListResourceGroupsResponseBodyResourceGroupsResourceGroup `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty" type:"Repeated"`
}

func (s ListResourceGroupsResponseBodyResourceGroups) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsResponseBodyResourceGroups) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBodyResourceGroups) GetResourceGroup() []*ListResourceGroupsResponseBodyResourceGroupsResourceGroup {
	return s.ResourceGroup
}

func (s *ListResourceGroupsResponseBodyResourceGroups) SetResourceGroup(v []*ListResourceGroupsResponseBodyResourceGroupsResourceGroup) *ListResourceGroupsResponseBodyResourceGroups {
	s.ResourceGroup = v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroups) Validate() error {
	if s.ResourceGroup != nil {
		for _, item := range s.ResourceGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourceGroupsResponseBodyResourceGroupsResourceGroup struct {
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
	Tags *ListResourceGroupsResponseBodyResourceGroupsResourceGroupTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s ListResourceGroupsResponseBodyResourceGroupsResourceGroup) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsResponseBodyResourceGroupsResourceGroup) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroup) GetAccountId() *string {
	return s.AccountId
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroup) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroup) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroup) GetId() *string {
	return s.Id
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroup) GetName() *string {
	return s.Name
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroup) GetStatus() *string {
	return s.Status
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroup) GetTags() *ListResourceGroupsResponseBodyResourceGroupsResourceGroupTags {
	return s.Tags
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroup) SetAccountId(v string) *ListResourceGroupsResponseBodyResourceGroupsResourceGroup {
	s.AccountId = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroup) SetCreateDate(v string) *ListResourceGroupsResponseBodyResourceGroupsResourceGroup {
	s.CreateDate = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroup) SetDisplayName(v string) *ListResourceGroupsResponseBodyResourceGroupsResourceGroup {
	s.DisplayName = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroup) SetId(v string) *ListResourceGroupsResponseBodyResourceGroupsResourceGroup {
	s.Id = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroup) SetName(v string) *ListResourceGroupsResponseBodyResourceGroupsResourceGroup {
	s.Name = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroup) SetStatus(v string) *ListResourceGroupsResponseBodyResourceGroupsResourceGroup {
	s.Status = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroup) SetTags(v *ListResourceGroupsResponseBodyResourceGroupsResourceGroupTags) *ListResourceGroupsResponseBodyResourceGroupsResourceGroup {
	s.Tags = v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroup) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListResourceGroupsResponseBodyResourceGroupsResourceGroupTags struct {
	Tag []*ListResourceGroupsResponseBodyResourceGroupsResourceGroupTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListResourceGroupsResponseBodyResourceGroupsResourceGroupTags) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsResponseBodyResourceGroupsResourceGroupTags) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroupTags) GetTag() []*ListResourceGroupsResponseBodyResourceGroupsResourceGroupTagsTag {
	return s.Tag
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroupTags) SetTag(v []*ListResourceGroupsResponseBodyResourceGroupsResourceGroupTagsTag) *ListResourceGroupsResponseBodyResourceGroupsResourceGroupTags {
	s.Tag = v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroupTags) Validate() error {
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

type ListResourceGroupsResponseBodyResourceGroupsResourceGroupTagsTag struct {
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

func (s ListResourceGroupsResponseBodyResourceGroupsResourceGroupTagsTag) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsResponseBodyResourceGroupsResourceGroupTagsTag) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroupTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroupTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroupTagsTag) SetTagKey(v string) *ListResourceGroupsResponseBodyResourceGroupsResourceGroupTagsTag {
	s.TagKey = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroupTagsTag) SetTagValue(v string) *ListResourceGroupsResponseBodyResourceGroupsResourceGroupTagsTag {
	s.TagValue = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupsResourceGroupTagsTag) Validate() error {
	return dara.Validate(s)
}
