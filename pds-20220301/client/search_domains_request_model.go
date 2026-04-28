// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v int64) *SearchDomainsRequest
	GetLimit() *int64
	SetMarker(v string) *SearchDomainsRequest
	GetMarker() *string
	SetName(v string) *SearchDomainsRequest
	GetName() *string
	SetOrderBy(v string) *SearchDomainsRequest
	GetOrderBy() *string
}

type SearchDomainsRequest struct {
	// The maximum number of results to return. Valid values: 1 to 100. Default value: 100.
	//
	// The number of returned results must be less than or equal to the specified number.
	//
	// example:
	//
	// 50
	Limit *int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker.\\
	//
	// By default, this parameter is empty.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The name of the domain. Fuzzy search is supported.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The sorting rule. Set the value to created_at, which specifies that the results are sorted based on the time when the domain was created.
	//
	// example:
	//
	// created_at
	OrderBy *string `json:"order_by,omitempty" xml:"order_by,omitempty"`
}

func (s SearchDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchDomainsRequest) GoString() string {
	return s.String()
}

func (s *SearchDomainsRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *SearchDomainsRequest) GetMarker() *string {
	return s.Marker
}

func (s *SearchDomainsRequest) GetName() *string {
	return s.Name
}

func (s *SearchDomainsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *SearchDomainsRequest) SetLimit(v int64) *SearchDomainsRequest {
	s.Limit = &v
	return s
}

func (s *SearchDomainsRequest) SetMarker(v string) *SearchDomainsRequest {
	s.Marker = &v
	return s
}

func (s *SearchDomainsRequest) SetName(v string) *SearchDomainsRequest {
	s.Name = &v
	return s
}

func (s *SearchDomainsRequest) SetOrderBy(v string) *SearchDomainsRequest {
	s.OrderBy = &v
	return s
}

func (s *SearchDomainsRequest) Validate() error {
	return dara.Validate(s)
}
