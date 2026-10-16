// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayName(v string) *ListResourceGroupsRequest
	GetDisplayName() *string
	SetIncludeTags(v bool) *ListResourceGroupsRequest
	GetIncludeTags() *bool
	SetName(v string) *ListResourceGroupsRequest
	GetName() *string
	SetPageNumber(v int32) *ListResourceGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResourceGroupsRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *ListResourceGroupsRequest
	GetResourceGroupId() *string
	SetResourceGroupIds(v []*string) *ListResourceGroupsRequest
	GetResourceGroupIds() []*string
	SetStatus(v string) *ListResourceGroupsRequest
	GetStatus() *string
	SetTag(v []*ListResourceGroupsRequestTag) *ListResourceGroupsRequest
	GetTag() []*ListResourceGroupsRequestTag
}

type ListResourceGroupsRequest struct {
	// The display name of the resource group. This parameter is used as a filter condition. Fuzzy match is supported.
	//
	// The value cannot exceed 50 characters in length.
	//
	// example:
	//
	// my-project
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// Specifies whether to return tag information. Valid values:
	//
	// - false (default): Tag information is not returned.
	//
	// - true: Tag information is returned.
	//
	// > If you specify tag-based filter conditions, tag information is always returned regardless of the value of IncludeTags.
	//
	// example:
	//
	// false
	IncludeTags *bool `json:"IncludeTags,omitempty" xml:"IncludeTags,omitempty"`
	// The unique identifier of the resource group. This parameter is used as a filter condition. Fuzzy match is supported.
	//
	// The value cannot exceed 50 characters in length and can contain letters, digits, and hyphens (-).
	//
	// example:
	//
	// my-project
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number.
	//
	// Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The resource group ID. This parameter is used as a filter condition.
	//
	// The value cannot exceed 18 characters in length and must start with `rg-`.
	//
	// > This parameter has been incorporated into the ResourceGroupIds parameter. If you set both ResourceGroupId and ResourceGroupIds, the value of ResourceGroupIds takes precedence.
	//
	// example:
	//
	// rg-9gLOoK****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of resource group IDs. This parameter is used as a filter condition.
	//
	// You can specify up to 100 resource group IDs.
	//
	// > If you set both ResourceGroupId and ResourceGroupIds, the value of ResourceGroupIds takes precedence.
	ResourceGroupIds []*string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty" type:"Repeated"`
	// The resource group status. This parameter is used as a filter condition. Valid values:
	//
	// - Creating: The resource group is being created.
	//
	// - OK: The resource group is created.
	//
	// - PendingDelete: The resource group is pending deletion.
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags. This parameter is used as a filter condition.
	//
	// You can specify up to 20 tags.
	Tag []*ListResourceGroupsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListResourceGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListResourceGroupsRequest) GetIncludeTags() *bool {
	return s.IncludeTags
}

func (s *ListResourceGroupsRequest) GetName() *string {
	return s.Name
}

func (s *ListResourceGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourceGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourceGroupsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListResourceGroupsRequest) GetResourceGroupIds() []*string {
	return s.ResourceGroupIds
}

func (s *ListResourceGroupsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListResourceGroupsRequest) GetTag() []*ListResourceGroupsRequestTag {
	return s.Tag
}

func (s *ListResourceGroupsRequest) SetDisplayName(v string) *ListResourceGroupsRequest {
	s.DisplayName = &v
	return s
}

func (s *ListResourceGroupsRequest) SetIncludeTags(v bool) *ListResourceGroupsRequest {
	s.IncludeTags = &v
	return s
}

func (s *ListResourceGroupsRequest) SetName(v string) *ListResourceGroupsRequest {
	s.Name = &v
	return s
}

func (s *ListResourceGroupsRequest) SetPageNumber(v int32) *ListResourceGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourceGroupsRequest) SetPageSize(v int32) *ListResourceGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourceGroupsRequest) SetResourceGroupId(v string) *ListResourceGroupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListResourceGroupsRequest) SetResourceGroupIds(v []*string) *ListResourceGroupsRequest {
	s.ResourceGroupIds = v
	return s
}

func (s *ListResourceGroupsRequest) SetStatus(v string) *ListResourceGroupsRequest {
	s.Status = &v
	return s
}

func (s *ListResourceGroupsRequest) SetTag(v []*ListResourceGroupsRequestTag) *ListResourceGroupsRequest {
	s.Tag = v
	return s
}

func (s *ListResourceGroupsRequest) Validate() error {
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

type ListResourceGroupsRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListResourceGroupsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsRequestTag) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListResourceGroupsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListResourceGroupsRequestTag) SetKey(v string) *ListResourceGroupsRequestTag {
	s.Key = &v
	return s
}

func (s *ListResourceGroupsRequestTag) SetValue(v string) *ListResourceGroupsRequestTag {
	s.Value = &v
	return s
}

func (s *ListResourceGroupsRequestTag) Validate() error {
	return dara.Validate(s)
}
