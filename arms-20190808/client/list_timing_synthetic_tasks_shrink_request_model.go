// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTimingSyntheticTasksShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListTimingSyntheticTasksShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListTimingSyntheticTasksShrinkRequest
	GetResourceGroupId() *string
	SetSearchShrink(v string) *ListTimingSyntheticTasksShrinkRequest
	GetSearchShrink() *string
	SetTagsShrink(v string) *ListTimingSyntheticTasksShrinkRequest
	GetTagsShrink() *string
}

type ListTimingSyntheticTasksShrinkRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxyexli2****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The search keyword.
	SearchShrink *string `json:"Search,omitempty" xml:"Search,omitempty"`
	// The tags.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListTimingSyntheticTasksShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTimingSyntheticTasksShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTimingSyntheticTasksShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTimingSyntheticTasksShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListTimingSyntheticTasksShrinkRequest) GetSearchShrink() *string {
	return s.SearchShrink
}

func (s *ListTimingSyntheticTasksShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ListTimingSyntheticTasksShrinkRequest) SetRegionId(v string) *ListTimingSyntheticTasksShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListTimingSyntheticTasksShrinkRequest) SetResourceGroupId(v string) *ListTimingSyntheticTasksShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTimingSyntheticTasksShrinkRequest) SetSearchShrink(v string) *ListTimingSyntheticTasksShrinkRequest {
	s.SearchShrink = &v
	return s
}

func (s *ListTimingSyntheticTasksShrinkRequest) SetTagsShrink(v string) *ListTimingSyntheticTasksShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListTimingSyntheticTasksShrinkRequest) Validate() error {
	return dara.Validate(s)
}
