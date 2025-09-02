// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListProjectsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListProjectsRequest
	GetPageSize() *int32
	SetResourceManagerResourceGroupId(v string) *ListProjectsRequest
	GetResourceManagerResourceGroupId() *string
	SetTags(v []*ListProjectsRequestTags) *ListProjectsRequest
	GetTags() []*ListProjectsRequestTags
}

type ListProjectsRequest struct {
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
	Tags []*ListProjectsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListProjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProjectsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProjectsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ListProjectsRequest) GetTags() []*ListProjectsRequestTags {
	return s.Tags
}

func (s *ListProjectsRequest) SetPageNumber(v int32) *ListProjectsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectsRequest) SetPageSize(v int32) *ListProjectsRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectsRequest) SetResourceManagerResourceGroupId(v string) *ListProjectsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ListProjectsRequest) SetTags(v []*ListProjectsRequestTags) *ListProjectsRequest {
	s.Tags = v
	return s
}

func (s *ListProjectsRequest) Validate() error {
	return dara.Validate(s)
}

type ListProjectsRequestTags struct {
	// The key of tag N to add to the workspace.
	//
	// example:
	//
	// Env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the workspace.
	//
	// example:
	//
	// Test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListProjectsRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsRequestTags) GoString() string {
	return s.String()
}

func (s *ListProjectsRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListProjectsRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListProjectsRequestTags) SetKey(v string) *ListProjectsRequestTags {
	s.Key = &v
	return s
}

func (s *ListProjectsRequestTags) SetValue(v string) *ListProjectsRequestTags {
	s.Value = &v
	return s
}

func (s *ListProjectsRequestTags) Validate() error {
	return dara.Validate(s)
}
