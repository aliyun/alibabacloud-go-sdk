// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListThreadsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int64) *ListThreadsResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *ListThreadsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListThreadsResponseBody
	GetRequestId() *string
	SetThreadId(v string) *ListThreadsResponseBody
	GetThreadId() *string
	SetThreads(v []*ListThreadsResponseBodyThreads) *ListThreadsResponseBody
	GetThreads() []*ListThreadsResponseBodyThreads
	SetTotal(v int64) *ListThreadsResponseBody
	GetTotal() *int64
}

type ListThreadsResponseBody struct {
	// example:
	//
	// 20
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// xxxxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CD8BA7D6-995D-578D-9941-78B0FECD14B5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// thread_id01
	ThreadId *string                           `json:"threadId,omitempty" xml:"threadId,omitempty"`
	Threads  []*ListThreadsResponseBodyThreads `json:"threads,omitempty" xml:"threads,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListThreadsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListThreadsResponseBody) GoString() string {
	return s.String()
}

func (s *ListThreadsResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListThreadsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListThreadsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListThreadsResponseBody) GetThreadId() *string {
	return s.ThreadId
}

func (s *ListThreadsResponseBody) GetThreads() []*ListThreadsResponseBodyThreads {
	return s.Threads
}

func (s *ListThreadsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListThreadsResponseBody) SetMaxResults(v int64) *ListThreadsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListThreadsResponseBody) SetNextToken(v string) *ListThreadsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListThreadsResponseBody) SetRequestId(v string) *ListThreadsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListThreadsResponseBody) SetThreadId(v string) *ListThreadsResponseBody {
	s.ThreadId = &v
	return s
}

func (s *ListThreadsResponseBody) SetThreads(v []*ListThreadsResponseBodyThreads) *ListThreadsResponseBody {
	s.Threads = v
	return s
}

func (s *ListThreadsResponseBody) SetTotal(v int64) *ListThreadsResponseBody {
	s.Total = &v
	return s
}

func (s *ListThreadsResponseBody) Validate() error {
	if s.Threads != nil {
		for _, item := range s.Threads {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListThreadsResponseBodyThreads struct {
	// example:
	//
	// 2025-04-22T12:46:34Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// test
	DigitalEmployeeName *string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty"`
	// example:
	//
	// active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// thread-t71rrw-1s7r7z9gu0v2s
	ThreadId *string `json:"threadId,omitempty" xml:"threadId,omitempty"`
	// example:
	//
	// test
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 2025-06-24 10:05:21
	UpdateTime *string                                  `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	Variables  *ListThreadsResponseBodyThreadsVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Struct"`
	// example:
	//
	// *
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListThreadsResponseBodyThreads) String() string {
	return dara.Prettify(s)
}

func (s ListThreadsResponseBodyThreads) GoString() string {
	return s.String()
}

func (s *ListThreadsResponseBodyThreads) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListThreadsResponseBodyThreads) GetDigitalEmployeeName() *string {
	return s.DigitalEmployeeName
}

func (s *ListThreadsResponseBodyThreads) GetStatus() *string {
	return s.Status
}

func (s *ListThreadsResponseBodyThreads) GetThreadId() *string {
	return s.ThreadId
}

func (s *ListThreadsResponseBodyThreads) GetTitle() *string {
	return s.Title
}

func (s *ListThreadsResponseBodyThreads) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListThreadsResponseBodyThreads) GetVariables() *ListThreadsResponseBodyThreadsVariables {
	return s.Variables
}

func (s *ListThreadsResponseBodyThreads) GetVersion() *int64 {
	return s.Version
}

func (s *ListThreadsResponseBodyThreads) SetCreateTime(v string) *ListThreadsResponseBodyThreads {
	s.CreateTime = &v
	return s
}

func (s *ListThreadsResponseBodyThreads) SetDigitalEmployeeName(v string) *ListThreadsResponseBodyThreads {
	s.DigitalEmployeeName = &v
	return s
}

func (s *ListThreadsResponseBodyThreads) SetStatus(v string) *ListThreadsResponseBodyThreads {
	s.Status = &v
	return s
}

func (s *ListThreadsResponseBodyThreads) SetThreadId(v string) *ListThreadsResponseBodyThreads {
	s.ThreadId = &v
	return s
}

func (s *ListThreadsResponseBodyThreads) SetTitle(v string) *ListThreadsResponseBodyThreads {
	s.Title = &v
	return s
}

func (s *ListThreadsResponseBodyThreads) SetUpdateTime(v string) *ListThreadsResponseBodyThreads {
	s.UpdateTime = &v
	return s
}

func (s *ListThreadsResponseBodyThreads) SetVariables(v *ListThreadsResponseBodyThreadsVariables) *ListThreadsResponseBodyThreads {
	s.Variables = v
	return s
}

func (s *ListThreadsResponseBodyThreads) SetVersion(v int64) *ListThreadsResponseBodyThreads {
	s.Version = &v
	return s
}

func (s *ListThreadsResponseBodyThreads) Validate() error {
	if s.Variables != nil {
		if err := s.Variables.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListThreadsResponseBodyThreadsVariables struct {
	// SLS project。
	//
	// example:
	//
	// ciec_rl_ads
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// example:
	//
	// hd1-np
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListThreadsResponseBodyThreadsVariables) String() string {
	return dara.Prettify(s)
}

func (s ListThreadsResponseBodyThreadsVariables) GoString() string {
	return s.String()
}

func (s *ListThreadsResponseBodyThreadsVariables) GetProject() *string {
	return s.Project
}

func (s *ListThreadsResponseBodyThreadsVariables) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListThreadsResponseBodyThreadsVariables) SetProject(v string) *ListThreadsResponseBodyThreadsVariables {
	s.Project = &v
	return s
}

func (s *ListThreadsResponseBodyThreadsVariables) SetWorkspace(v string) *ListThreadsResponseBodyThreadsVariables {
	s.Workspace = &v
	return s
}

func (s *ListThreadsResponseBodyThreadsVariables) Validate() error {
	return dara.Validate(s)
}
