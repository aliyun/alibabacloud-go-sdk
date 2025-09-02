// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizExtKey(v string) *ListResourceGroupsShrinkRequest
	GetBizExtKey() *string
	SetKeyword(v string) *ListResourceGroupsShrinkRequest
	GetKeyword() *string
	SetResourceGroupType(v int32) *ListResourceGroupsShrinkRequest
	GetResourceGroupType() *int32
	SetResourceManagerResourceGroupId(v string) *ListResourceGroupsShrinkRequest
	GetResourceManagerResourceGroupId() *string
	SetTagsShrink(v string) *ListResourceGroupsShrinkRequest
	GetTagsShrink() *string
}

type ListResourceGroupsShrinkRequest struct {
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
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListResourceGroupsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsShrinkRequest) GetBizExtKey() *string {
	return s.BizExtKey
}

func (s *ListResourceGroupsShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListResourceGroupsShrinkRequest) GetResourceGroupType() *int32 {
	return s.ResourceGroupType
}

func (s *ListResourceGroupsShrinkRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ListResourceGroupsShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ListResourceGroupsShrinkRequest) SetBizExtKey(v string) *ListResourceGroupsShrinkRequest {
	s.BizExtKey = &v
	return s
}

func (s *ListResourceGroupsShrinkRequest) SetKeyword(v string) *ListResourceGroupsShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListResourceGroupsShrinkRequest) SetResourceGroupType(v int32) *ListResourceGroupsShrinkRequest {
	s.ResourceGroupType = &v
	return s
}

func (s *ListResourceGroupsShrinkRequest) SetResourceManagerResourceGroupId(v string) *ListResourceGroupsShrinkRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ListResourceGroupsShrinkRequest) SetTagsShrink(v string) *ListResourceGroupsShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListResourceGroupsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
