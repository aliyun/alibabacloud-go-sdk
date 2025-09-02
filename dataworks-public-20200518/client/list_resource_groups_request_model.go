// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizExtKey(v string) *ListResourceGroupsRequest
	GetBizExtKey() *string
	SetKeyword(v string) *ListResourceGroupsRequest
	GetKeyword() *string
	SetResourceGroupType(v int32) *ListResourceGroupsRequest
	GetResourceGroupType() *int32
	SetResourceManagerResourceGroupId(v string) *ListResourceGroupsRequest
	GetResourceManagerResourceGroupId() *string
	SetTags(v []*ListResourceGroupsRequestTags) *ListResourceGroupsRequest
	GetTags() []*ListResourceGroupsRequestTags
}

type ListResourceGroupsRequest struct {
	// The category of the resource group. Valid values:
	//
	// 	- default (default): shared resource group
	//
	// 	- single: exclusive resource group
	//
	// example:
	//
	// default
	BizExtKey *string `json:"BizExtKey,omitempty" xml:"BizExtKey,omitempty"`
	// The keyword that is used for fuzzy match by resource group name and identifier.
	//
	// example:
	//
	// abc
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The type of the resource group that you want to query. Valid values:
	//
	// 	- 0: DataWorks
	//
	// 	- 1: scheduling
	//
	// 	- 2: MaxCompute
	//
	// 	- 3: Platform for AI (PAI)
	//
	// 	- 4: Data Integration
	//
	// 	- 7: exclusive resource group for scheduling (An ID is generated for the purchased resource when you purchase an exclusive resource group for scheduling.)
	//
	// 	- 9: DataService Studio
	//
	// 	- Default value: 1
	//
	// If the value indicates a compute engine, the resource groups to query are the ones that were created when you purchased the compute engine.
	//
	// example:
	//
	// 3
	ResourceGroupType *int32 `json:"ResourceGroupType,omitempty" xml:"ResourceGroupType,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzbn7pti3zfa
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The tags.
	Tags []*ListResourceGroupsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListResourceGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsRequest) GetBizExtKey() *string {
	return s.BizExtKey
}

func (s *ListResourceGroupsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListResourceGroupsRequest) GetResourceGroupType() *int32 {
	return s.ResourceGroupType
}

func (s *ListResourceGroupsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ListResourceGroupsRequest) GetTags() []*ListResourceGroupsRequestTags {
	return s.Tags
}

func (s *ListResourceGroupsRequest) SetBizExtKey(v string) *ListResourceGroupsRequest {
	s.BizExtKey = &v
	return s
}

func (s *ListResourceGroupsRequest) SetKeyword(v string) *ListResourceGroupsRequest {
	s.Keyword = &v
	return s
}

func (s *ListResourceGroupsRequest) SetResourceGroupType(v int32) *ListResourceGroupsRequest {
	s.ResourceGroupType = &v
	return s
}

func (s *ListResourceGroupsRequest) SetResourceManagerResourceGroupId(v string) *ListResourceGroupsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ListResourceGroupsRequest) SetTags(v []*ListResourceGroupsRequestTags) *ListResourceGroupsRequest {
	s.Tags = v
	return s
}

func (s *ListResourceGroupsRequest) Validate() error {
	return dara.Validate(s)
}

type ListResourceGroupsRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// Env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// Test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListResourceGroupsRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsRequestTags) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListResourceGroupsRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListResourceGroupsRequestTags) SetKey(v string) *ListResourceGroupsRequestTags {
	s.Key = &v
	return s
}

func (s *ListResourceGroupsRequestTags) SetValue(v string) *ListResourceGroupsRequestTags {
	s.Value = &v
	return s
}

func (s *ListResourceGroupsRequestTags) Validate() error {
	return dara.Validate(s)
}
