// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPortalPublicationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListPortalPublicationsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPortalPublicationsRequest
	GetPageSize() *int32
	SetProductName(v string) *ListPortalPublicationsRequest
	GetProductName() *string
}

type ListPortalPublicationsRequest struct {
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
	// The API product name. This corresponds to the productName field in the response list items.
	//
	// example:
	//
	// model-service
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
}

func (s ListPortalPublicationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPortalPublicationsRequest) GoString() string {
	return s.String()
}

func (s *ListPortalPublicationsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPortalPublicationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPortalPublicationsRequest) GetProductName() *string {
	return s.ProductName
}

func (s *ListPortalPublicationsRequest) SetPageNumber(v int32) *ListPortalPublicationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPortalPublicationsRequest) SetPageSize(v int32) *ListPortalPublicationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPortalPublicationsRequest) SetProductName(v string) *ListPortalPublicationsRequest {
	s.ProductName = &v
	return s
}

func (s *ListPortalPublicationsRequest) Validate() error {
	return dara.Validate(s)
}
