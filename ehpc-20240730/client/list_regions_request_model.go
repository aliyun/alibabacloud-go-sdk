// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListRegionsRequest
	GetAcceptLanguage() *string
	SetMaxResults(v int32) *ListRegionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListRegionsRequest
	GetNextToken() *string
	SetSpecCode(v string) *ListRegionsRequest
	GetSpecCode() *string
}

type ListRegionsRequest struct {
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0axxxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// Next
	SpecCode *string `json:"SpecCode,omitempty" xml:"SpecCode,omitempty"`
}

func (s ListRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRegionsRequest) GoString() string {
	return s.String()
}

func (s *ListRegionsRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListRegionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRegionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRegionsRequest) GetSpecCode() *string {
	return s.SpecCode
}

func (s *ListRegionsRequest) SetAcceptLanguage(v string) *ListRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListRegionsRequest) SetMaxResults(v int32) *ListRegionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRegionsRequest) SetNextToken(v string) *ListRegionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListRegionsRequest) SetSpecCode(v string) *ListRegionsRequest {
	s.SpecCode = &v
	return s
}

func (s *ListRegionsRequest) Validate() error {
	return dara.Validate(s)
}
