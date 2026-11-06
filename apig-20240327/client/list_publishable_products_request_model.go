// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublishableProductsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListPublishableProductsRequest
	GetName() *string
	SetPageNumber(v int32) *ListPublishableProductsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPublishableProductsRequest
	GetPageSize() *int32
}

type ListPublishableProductsRequest struct {
	// The name of the API product. This corresponds to the name field in the response list items.
	//
	// example:
	//
	// weather-query-api
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListPublishableProductsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPublishableProductsRequest) GoString() string {
	return s.String()
}

func (s *ListPublishableProductsRequest) GetName() *string {
	return s.Name
}

func (s *ListPublishableProductsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPublishableProductsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPublishableProductsRequest) SetName(v string) *ListPublishableProductsRequest {
	s.Name = &v
	return s
}

func (s *ListPublishableProductsRequest) SetPageNumber(v int32) *ListPublishableProductsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPublishableProductsRequest) SetPageSize(v int32) *ListPublishableProductsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPublishableProductsRequest) Validate() error {
	return dara.Validate(s)
}
