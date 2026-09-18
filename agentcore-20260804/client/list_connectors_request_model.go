// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConnectorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListConnectorsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListConnectorsRequest
	GetNextToken() *string
}

type ListConnectorsRequest struct {
	// The page size. The current version does not support this parameter.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The next page token. The current version does not support this parameter.
	//
	// example:
	//
	// dGVzdA==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListConnectorsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConnectorsRequest) GoString() string {
	return s.String()
}

func (s *ListConnectorsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListConnectorsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListConnectorsRequest) SetMaxResults(v int32) *ListConnectorsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListConnectorsRequest) SetNextToken(v string) *ListConnectorsRequest {
	s.NextToken = &v
	return s
}

func (s *ListConnectorsRequest) Validate() error {
	return dara.Validate(s)
}
