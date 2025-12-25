// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListThreadsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*ListThreadsRequestFilter) *ListThreadsRequest
	GetFilter() []*ListThreadsRequestFilter
	SetMaxResults(v int64) *ListThreadsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListThreadsRequest
	GetNextToken() *string
	SetStatus(v string) *ListThreadsRequest
	GetStatus() *string
	SetThreadId(v string) *ListThreadsRequest
	GetThreadId() *string
}

type ListThreadsRequest struct {
	Filter []*ListThreadsRequestFilter `json:"filter,omitempty" xml:"filter,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// xxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// thread_id01
	ThreadId *string `json:"threadId,omitempty" xml:"threadId,omitempty"`
}

func (s ListThreadsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListThreadsRequest) GoString() string {
	return s.String()
}

func (s *ListThreadsRequest) GetFilter() []*ListThreadsRequestFilter {
	return s.Filter
}

func (s *ListThreadsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListThreadsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListThreadsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListThreadsRequest) GetThreadId() *string {
	return s.ThreadId
}

func (s *ListThreadsRequest) SetFilter(v []*ListThreadsRequestFilter) *ListThreadsRequest {
	s.Filter = v
	return s
}

func (s *ListThreadsRequest) SetMaxResults(v int64) *ListThreadsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListThreadsRequest) SetNextToken(v string) *ListThreadsRequest {
	s.NextToken = &v
	return s
}

func (s *ListThreadsRequest) SetStatus(v string) *ListThreadsRequest {
	s.Status = &v
	return s
}

func (s *ListThreadsRequest) SetThreadId(v string) *ListThreadsRequest {
	s.ThreadId = &v
	return s
}

func (s *ListThreadsRequest) Validate() error {
	if s.Filter != nil {
		for _, item := range s.Filter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListThreadsRequestFilter struct {
	// This parameter is required.
	//
	// example:
	//
	// dukang-oxs-pre-obeqi
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{\\"Type\\": \\"primaryDoc\\", \\"Id\\": \\"01ZM8y7\\", \\"Name\\": \\"key\\"}, {\\"Type\\": \\"text\\", \\"Id\\": \\"mHe1U1b\\", \\"Name\\": \\"value\\"}, {\\"Type\\": \\"text\\", \\"Id\\": \\"rRIhpBs\\", \\"Name\\": u\\"\\u4f5c\\u7528\\"}]
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListThreadsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListThreadsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListThreadsRequestFilter) GetKey() *string {
	return s.Key
}

func (s *ListThreadsRequestFilter) GetValue() *string {
	return s.Value
}

func (s *ListThreadsRequestFilter) SetKey(v string) *ListThreadsRequestFilter {
	s.Key = &v
	return s
}

func (s *ListThreadsRequestFilter) SetValue(v string) *ListThreadsRequestFilter {
	s.Value = &v
	return s
}

func (s *ListThreadsRequestFilter) Validate() error {
	return dara.Validate(s)
}
