// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListResourceTagsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResourceTagsRequest
	GetPageSize() *int32
	SetResourceType(v string) *ListResourceTagsRequest
	GetResourceType() *string
}

type ListResourceTagsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListResourceTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTagsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceTagsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourceTagsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourceTagsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourceTagsRequest) SetPageNumber(v int32) *ListResourceTagsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourceTagsRequest) SetPageSize(v int32) *ListResourceTagsRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourceTagsRequest) SetResourceType(v string) *ListResourceTagsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListResourceTagsRequest) Validate() error {
	return dara.Validate(s)
}
