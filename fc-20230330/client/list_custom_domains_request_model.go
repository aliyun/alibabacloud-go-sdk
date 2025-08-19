// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v int32) *ListCustomDomainsRequest
	GetLimit() *int32
	SetNextToken(v string) *ListCustomDomainsRequest
	GetNextToken() *string
	SetPrefix(v string) *ListCustomDomainsRequest
	GetPrefix() *string
}

type ListCustomDomainsRequest struct {
	// The number of custom domain names returned.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// MTIzNCNhYmM=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The domain name prefix.
	//
	// example:
	//
	// foo
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
}

func (s ListCustomDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListCustomDomainsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListCustomDomainsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCustomDomainsRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *ListCustomDomainsRequest) SetLimit(v int32) *ListCustomDomainsRequest {
	s.Limit = &v
	return s
}

func (s *ListCustomDomainsRequest) SetNextToken(v string) *ListCustomDomainsRequest {
	s.NextToken = &v
	return s
}

func (s *ListCustomDomainsRequest) SetPrefix(v string) *ListCustomDomainsRequest {
	s.Prefix = &v
	return s
}

func (s *ListCustomDomainsRequest) Validate() error {
	return dara.Validate(s)
}
