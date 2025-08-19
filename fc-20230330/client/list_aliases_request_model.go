// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAliasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v int32) *ListAliasesRequest
	GetLimit() *int32
	SetNextToken(v string) *ListAliasesRequest
	GetNextToken() *string
	SetPrefix(v string) *ListAliasesRequest
	GetPrefix() *string
}

type ListAliasesRequest struct {
	// The number of aliases returned.
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
	// The alias prefix.
	//
	// example:
	//
	// my-alias
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
}

func (s ListAliasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAliasesRequest) GoString() string {
	return s.String()
}

func (s *ListAliasesRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListAliasesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAliasesRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *ListAliasesRequest) SetLimit(v int32) *ListAliasesRequest {
	s.Limit = &v
	return s
}

func (s *ListAliasesRequest) SetNextToken(v string) *ListAliasesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAliasesRequest) SetPrefix(v string) *ListAliasesRequest {
	s.Prefix = &v
	return s
}

func (s *ListAliasesRequest) Validate() error {
	return dara.Validate(s)
}
