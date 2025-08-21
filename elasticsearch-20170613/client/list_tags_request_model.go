// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageSize(v int32) *ListTagsRequest
	GetPageSize() *int32
	SetResourceType(v string) *ListTagsRequest
	GetResourceType() *string
}

type ListTagsRequest struct {
	// The return results.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The tag value of the ENI.
	//
	// example:
	//
	// INSTANCE
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagsRequest) GoString() string {
	return s.String()
}

func (s *ListTagsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTagsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagsRequest) SetPageSize(v int32) *ListTagsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagsRequest) SetResourceType(v string) *ListTagsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagsRequest) Validate() error {
	return dara.Validate(s)
}
