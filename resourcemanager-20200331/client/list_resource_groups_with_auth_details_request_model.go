// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupsWithAuthDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayName(v string) *ListResourceGroupsWithAuthDetailsRequest
	GetDisplayName() *string
	SetIncludeTags(v bool) *ListResourceGroupsWithAuthDetailsRequest
	GetIncludeTags() *bool
	SetName(v string) *ListResourceGroupsWithAuthDetailsRequest
	GetName() *string
	SetPageNumber(v int32) *ListResourceGroupsWithAuthDetailsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResourceGroupsWithAuthDetailsRequest
	GetPageSize() *int32
	SetResourceGroupIds(v []*string) *ListResourceGroupsWithAuthDetailsRequest
	GetResourceGroupIds() []*string
	SetResourceRegionId(v string) *ListResourceGroupsWithAuthDetailsRequest
	GetResourceRegionId() *string
	SetResourceTypes(v []*ListResourceGroupsWithAuthDetailsRequestResourceTypes) *ListResourceGroupsWithAuthDetailsRequest
	GetResourceTypes() []*ListResourceGroupsWithAuthDetailsRequestResourceTypes
	SetStatus(v string) *ListResourceGroupsWithAuthDetailsRequest
	GetStatus() *string
	SetTag(v []*ListResourceGroupsWithAuthDetailsRequestTag) *ListResourceGroupsWithAuthDetailsRequest
	GetTag() []*ListResourceGroupsWithAuthDetailsRequestTag
}

type ListResourceGroupsWithAuthDetailsRequest struct {
	// The display name of the resource group. This parameter specifies a filter condition for the query. Fuzzy search is supported.
	//
	// The display name can be a maximum of 50 characters in length.
	//
	// example:
	//
	// TestRG-BVT1
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// Specifies whether to return the information of tags. Valid values:
	//
	// 	- false (default)
	//
	// 	- true
	//
	// >  If you set a tag filter condition, the tag information is returned regardless of the `IncludeTags` value.
	//
	// example:
	//
	// true
	IncludeTags *bool `json:"IncludeTags,omitempty" xml:"IncludeTags,omitempty"`
	// The identifier of the resource group. This parameter specifies a filter condition for the query. Fuzzy search is supported.
	//
	// The identifier can be a maximum of 50 characters in length and can contain letters, digits, and hyphens (-).
	//
	// example:
	//
	// prod-rg
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// The IDs of the resource groups that you want to query.
	ResourceGroupIds []*string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty" type:"Repeated"`
	// The ID of the region where the resource resides.
	//
	// example:
	//
	// cn-shenzhen
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The resource types.
	ResourceTypes []*ListResourceGroupsWithAuthDetailsRequestResourceTypes `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
	// The status of the resource group. This parameter specifies a filter condition for the query. Valid values:
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
	// The tags.
	Tag []*ListResourceGroupsWithAuthDetailsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListResourceGroupsWithAuthDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsWithAuthDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsWithAuthDetailsRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListResourceGroupsWithAuthDetailsRequest) GetIncludeTags() *bool {
	return s.IncludeTags
}

func (s *ListResourceGroupsWithAuthDetailsRequest) GetName() *string {
	return s.Name
}

func (s *ListResourceGroupsWithAuthDetailsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourceGroupsWithAuthDetailsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourceGroupsWithAuthDetailsRequest) GetResourceGroupIds() []*string {
	return s.ResourceGroupIds
}

func (s *ListResourceGroupsWithAuthDetailsRequest) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *ListResourceGroupsWithAuthDetailsRequest) GetResourceTypes() []*ListResourceGroupsWithAuthDetailsRequestResourceTypes {
	return s.ResourceTypes
}

func (s *ListResourceGroupsWithAuthDetailsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListResourceGroupsWithAuthDetailsRequest) GetTag() []*ListResourceGroupsWithAuthDetailsRequestTag {
	return s.Tag
}

func (s *ListResourceGroupsWithAuthDetailsRequest) SetDisplayName(v string) *ListResourceGroupsWithAuthDetailsRequest {
	s.DisplayName = &v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsRequest) SetIncludeTags(v bool) *ListResourceGroupsWithAuthDetailsRequest {
	s.IncludeTags = &v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsRequest) SetName(v string) *ListResourceGroupsWithAuthDetailsRequest {
	s.Name = &v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsRequest) SetPageNumber(v int32) *ListResourceGroupsWithAuthDetailsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsRequest) SetPageSize(v int32) *ListResourceGroupsWithAuthDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsRequest) SetResourceGroupIds(v []*string) *ListResourceGroupsWithAuthDetailsRequest {
	s.ResourceGroupIds = v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsRequest) SetResourceRegionId(v string) *ListResourceGroupsWithAuthDetailsRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsRequest) SetResourceTypes(v []*ListResourceGroupsWithAuthDetailsRequestResourceTypes) *ListResourceGroupsWithAuthDetailsRequest {
	s.ResourceTypes = v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsRequest) SetStatus(v string) *ListResourceGroupsWithAuthDetailsRequest {
	s.Status = &v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsRequest) SetTag(v []*ListResourceGroupsWithAuthDetailsRequestTag) *ListResourceGroupsWithAuthDetailsRequest {
	s.Tag = v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsRequest) Validate() error {
	if s.ResourceTypes != nil {
		for _, item := range s.ResourceTypes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type ListResourceGroupsWithAuthDetailsRequestResourceTypes struct {
	// The resource type.
	//
	// You can obtain the resource type from the **Resource type*	- column in [Services that work with Resource Group](https://help.aliyun.com/document_detail/94479.html).
	//
	// example:
	//
	// instance
	ResourceTypeCode *string `json:"ResourceTypeCode,omitempty" xml:"ResourceTypeCode,omitempty"`
	// The service code.
	//
	// You can obtain the code from the **Service code*	- column in [Services that work with Resource Group](https://help.aliyun.com/document_detail/94479.html).
	//
	// example:
	//
	// ecs
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s ListResourceGroupsWithAuthDetailsRequestResourceTypes) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsWithAuthDetailsRequestResourceTypes) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsWithAuthDetailsRequestResourceTypes) GetResourceTypeCode() *string {
	return s.ResourceTypeCode
}

func (s *ListResourceGroupsWithAuthDetailsRequestResourceTypes) GetService() *string {
	return s.Service
}

func (s *ListResourceGroupsWithAuthDetailsRequestResourceTypes) SetResourceTypeCode(v string) *ListResourceGroupsWithAuthDetailsRequestResourceTypes {
	s.ResourceTypeCode = &v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsRequestResourceTypes) SetService(v string) *ListResourceGroupsWithAuthDetailsRequestResourceTypes {
	s.Service = &v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsRequestResourceTypes) Validate() error {
	return dara.Validate(s)
}

type ListResourceGroupsWithAuthDetailsRequestTag struct {
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

func (s ListResourceGroupsWithAuthDetailsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsWithAuthDetailsRequestTag) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsWithAuthDetailsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListResourceGroupsWithAuthDetailsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListResourceGroupsWithAuthDetailsRequestTag) SetKey(v string) *ListResourceGroupsWithAuthDetailsRequestTag {
	s.Key = &v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsRequestTag) SetValue(v string) *ListResourceGroupsWithAuthDetailsRequestTag {
	s.Value = &v
	return s
}

func (s *ListResourceGroupsWithAuthDetailsRequestTag) Validate() error {
	return dara.Validate(s)
}
