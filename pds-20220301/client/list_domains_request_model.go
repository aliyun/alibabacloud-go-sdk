// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v int64) *ListDomainsRequest
	GetLimit() *int64
	SetMarker(v string) *ListDomainsRequest
	GetMarker() *string
	SetParentDomainId(v string) *ListDomainsRequest
	GetParentDomainId() *string
	SetServiceCode(v string) *ListDomainsRequest
	GetServiceCode() *string
}

type ListDomainsRequest struct {
	// The maximum number of results to return. Valid values: 1 to 100. Default value: 50.
	//
	// example:
	//
	// 60
	Limit *int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The ID of the parent domain.
	//
	// example:
	//
	// bj1
	ParentDomainId *string `json:"parent_domain_id,omitempty" xml:"parent_domain_id,omitempty"`
	ServiceCode    *string `json:"service_code,omitempty" xml:"service_code,omitempty"`
}

func (s ListDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListDomainsRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *ListDomainsRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListDomainsRequest) GetParentDomainId() *string {
	return s.ParentDomainId
}

func (s *ListDomainsRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *ListDomainsRequest) SetLimit(v int64) *ListDomainsRequest {
	s.Limit = &v
	return s
}

func (s *ListDomainsRequest) SetMarker(v string) *ListDomainsRequest {
	s.Marker = &v
	return s
}

func (s *ListDomainsRequest) SetParentDomainId(v string) *ListDomainsRequest {
	s.ParentDomainId = &v
	return s
}

func (s *ListDomainsRequest) SetServiceCode(v string) *ListDomainsRequest {
	s.ServiceCode = &v
	return s
}

func (s *ListDomainsRequest) Validate() error {
	return dara.Validate(s)
}
