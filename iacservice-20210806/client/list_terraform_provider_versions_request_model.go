// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTerraformProviderVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListTerraformProviderVersionsRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListTerraformProviderVersionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTerraformProviderVersionsRequest
	GetNextToken() *string
	SetUsage(v string) *ListTerraformProviderVersionsRequest
	GetUsage() *string
}

type ListTerraformProviderVersionsRequest struct {
	// The keyword for searching versions. Fuzzy match is supported.
	//
	// example:
	//
	// key
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The maximum number of records to retrieve in a single request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token for the next page. A value of null indicates that no more pages are available.
	//
	// example:
	//
	// ar9lHbaidqWequN5orMba54dD+2BRJj42DLT6GrZysw=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The usage. Set to Explorer to retrieve meta information.
	//
	// example:
	//
	// Explorer
	Usage *string `json:"usage,omitempty" xml:"usage,omitempty"`
}

func (s ListTerraformProviderVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTerraformProviderVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListTerraformProviderVersionsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListTerraformProviderVersionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTerraformProviderVersionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTerraformProviderVersionsRequest) GetUsage() *string {
	return s.Usage
}

func (s *ListTerraformProviderVersionsRequest) SetKeyword(v string) *ListTerraformProviderVersionsRequest {
	s.Keyword = &v
	return s
}

func (s *ListTerraformProviderVersionsRequest) SetMaxResults(v int32) *ListTerraformProviderVersionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTerraformProviderVersionsRequest) SetNextToken(v string) *ListTerraformProviderVersionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTerraformProviderVersionsRequest) SetUsage(v string) *ListTerraformProviderVersionsRequest {
	s.Usage = &v
	return s
}

func (s *ListTerraformProviderVersionsRequest) Validate() error {
	return dara.Validate(s)
}
