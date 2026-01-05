// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcesForTagOptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListResourcesForTagOptionRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResourcesForTagOptionRequest
	GetPageSize() *int32
	SetResourceType(v string) *ListResourcesForTagOptionRequest
	GetResourceType() *string
	SetTagOptionId(v string) *ListResourcesForTagOptionRequest
	GetTagOptionId() *string
}

type ListResourcesForTagOptionRequest struct {
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100 Minimum value: 1. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of resource that is associated with the tag option. Valid values:
	//
	// 	- product: product
	//
	// 	- Portfolio: product portfolio
	//
	// This parameter is required.
	//
	// example:
	//
	// Product
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the tag option.
	//
	// This parameter is required.
	//
	// example:
	//
	// tag-bp1u6mdf3d****
	TagOptionId *string `json:"TagOptionId,omitempty" xml:"TagOptionId,omitempty"`
}

func (s ListResourcesForTagOptionRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesForTagOptionRequest) GoString() string {
	return s.String()
}

func (s *ListResourcesForTagOptionRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourcesForTagOptionRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourcesForTagOptionRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourcesForTagOptionRequest) GetTagOptionId() *string {
	return s.TagOptionId
}

func (s *ListResourcesForTagOptionRequest) SetPageNumber(v int32) *ListResourcesForTagOptionRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesForTagOptionRequest) SetPageSize(v int32) *ListResourcesForTagOptionRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourcesForTagOptionRequest) SetResourceType(v string) *ListResourcesForTagOptionRequest {
	s.ResourceType = &v
	return s
}

func (s *ListResourcesForTagOptionRequest) SetTagOptionId(v string) *ListResourcesForTagOptionRequest {
	s.TagOptionId = &v
	return s
}

func (s *ListResourcesForTagOptionRequest) Validate() error {
	return dara.Validate(s)
}
