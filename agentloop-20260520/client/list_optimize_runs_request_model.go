// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOptimizeRunsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListOptimizeRunsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListOptimizeRunsRequest
	GetNextToken() *string
	SetStatus(v string) *ListOptimizeRunsRequest
	GetStatus() *string
	SetWindowStartTimeBegin(v int64) *ListOptimizeRunsRequest
	GetWindowStartTimeBegin() *int64
	SetWindowStartTimeEnd(v int64) *ListOptimizeRunsRequest
	GetWindowStartTimeEnd() *int64
}

type ListOptimizeRunsRequest struct {
	// The maximum number of entries to return. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// MTIzNDU2Nzg5MA==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The status used to filter results.
	//
	// example:
	//
	// Succeeded
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The lower bound of the window start time in UNIX seconds.
	//
	// example:
	//
	// 1786432688000
	WindowStartTimeBegin *int64 `json:"windowStartTimeBegin,omitempty" xml:"windowStartTimeBegin,omitempty"`
	// The upper bound of the window start time in UNIX seconds.
	//
	// example:
	//
	// 1786432690000
	WindowStartTimeEnd *int64 `json:"windowStartTimeEnd,omitempty" xml:"windowStartTimeEnd,omitempty"`
}

func (s ListOptimizeRunsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOptimizeRunsRequest) GoString() string {
	return s.String()
}

func (s *ListOptimizeRunsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListOptimizeRunsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListOptimizeRunsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListOptimizeRunsRequest) GetWindowStartTimeBegin() *int64 {
	return s.WindowStartTimeBegin
}

func (s *ListOptimizeRunsRequest) GetWindowStartTimeEnd() *int64 {
	return s.WindowStartTimeEnd
}

func (s *ListOptimizeRunsRequest) SetMaxResults(v int32) *ListOptimizeRunsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListOptimizeRunsRequest) SetNextToken(v string) *ListOptimizeRunsRequest {
	s.NextToken = &v
	return s
}

func (s *ListOptimizeRunsRequest) SetStatus(v string) *ListOptimizeRunsRequest {
	s.Status = &v
	return s
}

func (s *ListOptimizeRunsRequest) SetWindowStartTimeBegin(v int64) *ListOptimizeRunsRequest {
	s.WindowStartTimeBegin = &v
	return s
}

func (s *ListOptimizeRunsRequest) SetWindowStartTimeEnd(v int64) *ListOptimizeRunsRequest {
	s.WindowStartTimeEnd = &v
	return s
}

func (s *ListOptimizeRunsRequest) Validate() error {
	return dara.Validate(s)
}
