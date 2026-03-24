// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListThreadsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterShrink(v string) *ListThreadsShrinkRequest
	GetFilterShrink() *string
	SetMaxResults(v int64) *ListThreadsShrinkRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListThreadsShrinkRequest
	GetNextToken() *string
	SetStatus(v string) *ListThreadsShrinkRequest
	GetStatus() *string
	SetThreadId(v string) *ListThreadsShrinkRequest
	GetThreadId() *string
}

type ListThreadsShrinkRequest struct {
	// The filter conditions for the query. If you do not specify this parameter, all threads in the instance are queried.
	FilterShrink *string `json:"filter,omitempty" xml:"filter,omitempty"`
	// The maximum number of results to return. The maximum value is 200.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The paging token.
	//
	// example:
	//
	// xxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The session status.
	//
	// example:
	//
	// active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The session ID.
	//
	// example:
	//
	// thread-123123
	ThreadId *string `json:"threadId,omitempty" xml:"threadId,omitempty"`
}

func (s ListThreadsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListThreadsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListThreadsShrinkRequest) GetFilterShrink() *string {
	return s.FilterShrink
}

func (s *ListThreadsShrinkRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListThreadsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListThreadsShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ListThreadsShrinkRequest) GetThreadId() *string {
	return s.ThreadId
}

func (s *ListThreadsShrinkRequest) SetFilterShrink(v string) *ListThreadsShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *ListThreadsShrinkRequest) SetMaxResults(v int64) *ListThreadsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListThreadsShrinkRequest) SetNextToken(v string) *ListThreadsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListThreadsShrinkRequest) SetStatus(v string) *ListThreadsShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListThreadsShrinkRequest) SetThreadId(v string) *ListThreadsShrinkRequest {
	s.ThreadId = &v
	return s
}

func (s *ListThreadsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
