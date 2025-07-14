// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWebApplicationRevisionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v int32) *ListWebApplicationRevisionsRequest
	GetLimit() *int32
	SetNamespaceId(v string) *ListWebApplicationRevisionsRequest
	GetNamespaceId() *string
	SetNextToken(v string) *ListWebApplicationRevisionsRequest
	GetNextToken() *string
}

type ListWebApplicationRevisionsRequest struct {
	// The number of applications returned.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The namespace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// A2RN
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListWebApplicationRevisionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWebApplicationRevisionsRequest) GoString() string {
	return s.String()
}

func (s *ListWebApplicationRevisionsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListWebApplicationRevisionsRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListWebApplicationRevisionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWebApplicationRevisionsRequest) SetLimit(v int32) *ListWebApplicationRevisionsRequest {
	s.Limit = &v
	return s
}

func (s *ListWebApplicationRevisionsRequest) SetNamespaceId(v string) *ListWebApplicationRevisionsRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListWebApplicationRevisionsRequest) SetNextToken(v string) *ListWebApplicationRevisionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListWebApplicationRevisionsRequest) Validate() error {
	return dara.Validate(s)
}
