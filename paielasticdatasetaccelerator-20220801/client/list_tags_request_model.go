// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *ListTagsRequest
	GetKey() *string
	SetOrder(v string) *ListTagsRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListTagsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTagsRequest
	GetPageSize() *int32
	SetResourceId(v string) *ListTagsRequest
	GetResourceId() *string
	SetResourceType(v string) *ListTagsRequest
	GetResourceType() *string
	SetSortBy(v string) *ListTagsRequest
	GetSortBy() *string
	SetValue(v string) *ListTagsRequest
	GetValue() *string
}

type ListTagsRequest struct {
	// example:
	//
	// dataset_version
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// Desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// inst-my1tk3jggooi5uwks5
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// v0.1.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagsRequest) GoString() string {
	return s.String()
}

func (s *ListTagsRequest) GetKey() *string {
	return s.Key
}

func (s *ListTagsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListTagsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTagsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTagsRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListTagsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListTagsRequest) GetValue() *string {
	return s.Value
}

func (s *ListTagsRequest) SetKey(v string) *ListTagsRequest {
	s.Key = &v
	return s
}

func (s *ListTagsRequest) SetOrder(v string) *ListTagsRequest {
	s.Order = &v
	return s
}

func (s *ListTagsRequest) SetPageNumber(v int32) *ListTagsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTagsRequest) SetPageSize(v int32) *ListTagsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagsRequest) SetResourceId(v string) *ListTagsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListTagsRequest) SetResourceType(v string) *ListTagsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagsRequest) SetSortBy(v string) *ListTagsRequest {
	s.SortBy = &v
	return s
}

func (s *ListTagsRequest) SetValue(v string) *ListTagsRequest {
	s.Value = &v
	return s
}

func (s *ListTagsRequest) Validate() error {
	return dara.Validate(s)
}
