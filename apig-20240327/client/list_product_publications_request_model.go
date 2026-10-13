// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductPublicationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListProductPublicationsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListProductPublicationsRequest
	GetPageSize() *int32
	SetPortalName(v string) *ListProductPublicationsRequest
	GetPortalName() *string
}

type ListProductPublicationsRequest struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The name of the portal to which the API product is published.
	//
	// example:
	//
	// developer-portal
	PortalName *string `json:"portalName,omitempty" xml:"portalName,omitempty"`
}

func (s ListProductPublicationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProductPublicationsRequest) GoString() string {
	return s.String()
}

func (s *ListProductPublicationsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProductPublicationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProductPublicationsRequest) GetPortalName() *string {
	return s.PortalName
}

func (s *ListProductPublicationsRequest) SetPageNumber(v int32) *ListProductPublicationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProductPublicationsRequest) SetPageSize(v int32) *ListProductPublicationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListProductPublicationsRequest) SetPortalName(v string) *ListProductPublicationsRequest {
	s.PortalName = &v
	return s
}

func (s *ListProductPublicationsRequest) Validate() error {
	return dara.Validate(s)
}
