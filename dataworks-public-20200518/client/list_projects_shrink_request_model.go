// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListProjectsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListProjectsShrinkRequest
	GetPageSize() *int32
	SetResourceManagerResourceGroupId(v string) *ListProjectsShrinkRequest
	GetResourceManagerResourceGroupId() *string
	SetTagsShrink(v string) *ListProjectsShrinkRequest
	GetTagsShrink() *string
}

type ListProjectsShrinkRequest struct {
	// The page number. Valid values: 1 to 100.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzbn7pti3zfa
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The tags to add to the workspace.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListProjectsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProjectsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProjectsShrinkRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ListProjectsShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ListProjectsShrinkRequest) SetPageNumber(v int32) *ListProjectsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetPageSize(v int32) *ListProjectsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetResourceManagerResourceGroupId(v string) *ListProjectsShrinkRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetTagsShrink(v string) *ListProjectsShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListProjectsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
