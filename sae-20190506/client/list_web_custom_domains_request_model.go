// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWebCustomDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ListWebCustomDomainsRequest
	GetApplicationId() *string
	SetLimit(v int32) *ListWebCustomDomainsRequest
	GetLimit() *int32
	SetNamespaceId(v string) *ListWebCustomDomainsRequest
	GetNamespaceId() *string
	SetNextToken(v string) *ListWebCustomDomainsRequest
	GetNextToken() *string
	SetPrefix(v string) *ListWebCustomDomainsRequest
	GetPrefix() *string
}

type ListWebCustomDomainsRequest struct {
	// The application ID.
	//
	// example:
	//
	// 7e41aff0-9eca-45c9-ac48-675e09******
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The number of custom domain names returned.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The namespace ID.
	//
	// example:
	//
	// cn-hangzhou
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// A2RN
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The prefix of the custom domain name that you want to query.
	//
	// example:
	//
	// remoteresult
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
}

func (s ListWebCustomDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWebCustomDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListWebCustomDomainsRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListWebCustomDomainsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListWebCustomDomainsRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListWebCustomDomainsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWebCustomDomainsRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *ListWebCustomDomainsRequest) SetApplicationId(v string) *ListWebCustomDomainsRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListWebCustomDomainsRequest) SetLimit(v int32) *ListWebCustomDomainsRequest {
	s.Limit = &v
	return s
}

func (s *ListWebCustomDomainsRequest) SetNamespaceId(v string) *ListWebCustomDomainsRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListWebCustomDomainsRequest) SetNextToken(v string) *ListWebCustomDomainsRequest {
	s.NextToken = &v
	return s
}

func (s *ListWebCustomDomainsRequest) SetPrefix(v string) *ListWebCustomDomainsRequest {
	s.Prefix = &v
	return s
}

func (s *ListWebCustomDomainsRequest) Validate() error {
	return dara.Validate(s)
}
