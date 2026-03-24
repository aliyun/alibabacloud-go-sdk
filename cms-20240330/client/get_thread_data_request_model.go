// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetThreadDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int64) *GetThreadDataRequest
	GetMaxResults() *int64
	SetNextToken(v string) *GetThreadDataRequest
	GetNextToken() *string
}

type GetThreadDataRequest struct {
	// The maximum number of session data entries to return. The maximum value is 100.
	//
	// example:
	//
	// 1
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// xxxxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetThreadDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetThreadDataRequest) GoString() string {
	return s.String()
}

func (s *GetThreadDataRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *GetThreadDataRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetThreadDataRequest) SetMaxResults(v int64) *GetThreadDataRequest {
	s.MaxResults = &v
	return s
}

func (s *GetThreadDataRequest) SetNextToken(v string) *GetThreadDataRequest {
	s.NextToken = &v
	return s
}

func (s *GetThreadDataRequest) Validate() error {
	return dara.Validate(s)
}
