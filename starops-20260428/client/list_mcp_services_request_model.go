// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcpServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListMcpServicesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListMcpServicesRequest
	GetNextToken() *string
}

type ListMcpServicesRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListMcpServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMcpServicesRequest) GoString() string {
	return s.String()
}

func (s *ListMcpServicesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMcpServicesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMcpServicesRequest) SetMaxResults(v int32) *ListMcpServicesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMcpServicesRequest) SetNextToken(v string) *ListMcpServicesRequest {
	s.NextToken = &v
	return s
}

func (s *ListMcpServicesRequest) Validate() error {
	return dara.Validate(s)
}
