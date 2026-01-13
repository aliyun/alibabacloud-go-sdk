// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *ListCustomDomainsRequest
	GetDomainName() *string
	SetPageNumber(v int32) *ListCustomDomainsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomDomainsRequest
	GetPageSize() *int32
	SetResourceName(v string) *ListCustomDomainsRequest
	GetResourceName() *string
	SetResourceType(v string) *ListCustomDomainsRequest
	GetResourceType() *string
}

type ListCustomDomainsRequest struct {
	// example:
	//
	// mydomain.com
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// agent-1
	ResourceName *string `json:"resourceName,omitempty" xml:"resourceName,omitempty"`
	// example:
	//
	// runtime
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListCustomDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListCustomDomainsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *ListCustomDomainsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomDomainsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomDomainsRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListCustomDomainsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListCustomDomainsRequest) SetDomainName(v string) *ListCustomDomainsRequest {
	s.DomainName = &v
	return s
}

func (s *ListCustomDomainsRequest) SetPageNumber(v int32) *ListCustomDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCustomDomainsRequest) SetPageSize(v int32) *ListCustomDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomDomainsRequest) SetResourceName(v string) *ListCustomDomainsRequest {
	s.ResourceName = &v
	return s
}

func (s *ListCustomDomainsRequest) SetResourceType(v string) *ListCustomDomainsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListCustomDomainsRequest) Validate() error {
	return dara.Validate(s)
}
