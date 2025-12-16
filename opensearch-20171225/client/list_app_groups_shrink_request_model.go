// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppGroupsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListAppGroupsShrinkRequest
	GetInstanceId() *string
	SetName(v string) *ListAppGroupsShrinkRequest
	GetName() *string
	SetPageNumber(v int32) *ListAppGroupsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAppGroupsShrinkRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *ListAppGroupsShrinkRequest
	GetResourceGroupId() *string
	SetSortBy(v int32) *ListAppGroupsShrinkRequest
	GetSortBy() *int32
	SetTagsShrink(v string) *ListAppGroupsShrinkRequest
	GetTagsShrink() *string
	SetType(v string) *ListAppGroupsShrinkRequest
	GetType() *string
}

type ListAppGroupsShrinkRequest struct {
	// The ID of the instance. Exact match is used.
	//
	// example:
	//
	// ops-cn-xxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// my_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// "110123123"
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The method based on which applications are sorted. Valid values:
	//
	// 	- 0: sorts applications in descending order by creation time.
	//
	// 	- 1: sorts applications in descending order by modification time.
	//
	// Default value: 0.
	//
	// example:
	//
	// 0
	SortBy *int32 `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
	// The tags.
	TagsShrink *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The type of the application. Valid values:
	//
	// 	- standard: a High-performance Search Edition application.
	//
	// 	- enhanced: an Industry Algorithm Edition application.
	//
	// example:
	//
	// standard
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListAppGroupsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppGroupsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAppGroupsShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAppGroupsShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListAppGroupsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAppGroupsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppGroupsShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListAppGroupsShrinkRequest) GetSortBy() *int32 {
	return s.SortBy
}

func (s *ListAppGroupsShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ListAppGroupsShrinkRequest) GetType() *string {
	return s.Type
}

func (s *ListAppGroupsShrinkRequest) SetInstanceId(v string) *ListAppGroupsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAppGroupsShrinkRequest) SetName(v string) *ListAppGroupsShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListAppGroupsShrinkRequest) SetPageNumber(v int32) *ListAppGroupsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppGroupsShrinkRequest) SetPageSize(v int32) *ListAppGroupsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppGroupsShrinkRequest) SetResourceGroupId(v string) *ListAppGroupsShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListAppGroupsShrinkRequest) SetSortBy(v int32) *ListAppGroupsShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListAppGroupsShrinkRequest) SetTagsShrink(v string) *ListAppGroupsShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListAppGroupsShrinkRequest) SetType(v string) *ListAppGroupsShrinkRequest {
	s.Type = &v
	return s
}

func (s *ListAppGroupsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
