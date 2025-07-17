// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAssetTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ListDataAssetTagsRequest
	GetCategory() *string
	SetKey(v string) *ListDataAssetTagsRequest
	GetKey() *string
	SetPageNumber(v int32) *ListDataAssetTagsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataAssetTagsRequest
	GetPageSize() *int32
}

type ListDataAssetTagsRequest struct {
	// The type of the tag. Valid values:
	//
	// 	- Normal
	//
	// 	- System
	//
	// example:
	//
	// Normal
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
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
}

func (s ListDataAssetTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataAssetTagsRequest) GoString() string {
	return s.String()
}

func (s *ListDataAssetTagsRequest) GetCategory() *string {
	return s.Category
}

func (s *ListDataAssetTagsRequest) GetKey() *string {
	return s.Key
}

func (s *ListDataAssetTagsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataAssetTagsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataAssetTagsRequest) SetCategory(v string) *ListDataAssetTagsRequest {
	s.Category = &v
	return s
}

func (s *ListDataAssetTagsRequest) SetKey(v string) *ListDataAssetTagsRequest {
	s.Key = &v
	return s
}

func (s *ListDataAssetTagsRequest) SetPageNumber(v int32) *ListDataAssetTagsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataAssetTagsRequest) SetPageSize(v int32) *ListDataAssetTagsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataAssetTagsRequest) Validate() error {
	return dara.Validate(s)
}
