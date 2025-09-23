// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBiddingDocRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTimeEnd(v string) *ListBiddingDocRequest
	GetCreateTimeEnd() *string
	SetCreateTimeStart(v string) *ListBiddingDocRequest
	GetCreateTimeStart() *string
	SetCurrent(v int32) *ListBiddingDocRequest
	GetCurrent() *int32
	SetMaxResults(v int32) *ListBiddingDocRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListBiddingDocRequest
	GetNextToken() *string
	SetSize(v int32) *ListBiddingDocRequest
	GetSize() *int32
	SetSkip(v int32) *ListBiddingDocRequest
	GetSkip() *int32
	SetTaskName(v string) *ListBiddingDocRequest
	GetTaskName() *string
	SetTaskStatus(v int32) *ListBiddingDocRequest
	GetTaskStatus() *int32
	SetWorkspaceId(v string) *ListBiddingDocRequest
	GetWorkspaceId() *string
}

type ListBiddingDocRequest struct {
	// example:
	//
	// 2023-03-18 02:00:00
	CreateTimeEnd *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// example:
	//
	// 2023-02-19 07:28:11
	CreateTimeStart *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// example:
	//
	// 1
	Current *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	// example:
	//
	// null
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// null
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// null
	Skip     *int32  `json:"Skip,omitempty" xml:"Skip,omitempty"`
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// 0-waiting、1-running、2-success、3-pause、4-fail
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListBiddingDocRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBiddingDocRequest) GoString() string {
	return s.String()
}

func (s *ListBiddingDocRequest) GetCreateTimeEnd() *string {
	return s.CreateTimeEnd
}

func (s *ListBiddingDocRequest) GetCreateTimeStart() *string {
	return s.CreateTimeStart
}

func (s *ListBiddingDocRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *ListBiddingDocRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListBiddingDocRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListBiddingDocRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListBiddingDocRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListBiddingDocRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *ListBiddingDocRequest) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *ListBiddingDocRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListBiddingDocRequest) SetCreateTimeEnd(v string) *ListBiddingDocRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *ListBiddingDocRequest) SetCreateTimeStart(v string) *ListBiddingDocRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *ListBiddingDocRequest) SetCurrent(v int32) *ListBiddingDocRequest {
	s.Current = &v
	return s
}

func (s *ListBiddingDocRequest) SetMaxResults(v int32) *ListBiddingDocRequest {
	s.MaxResults = &v
	return s
}

func (s *ListBiddingDocRequest) SetNextToken(v string) *ListBiddingDocRequest {
	s.NextToken = &v
	return s
}

func (s *ListBiddingDocRequest) SetSize(v int32) *ListBiddingDocRequest {
	s.Size = &v
	return s
}

func (s *ListBiddingDocRequest) SetSkip(v int32) *ListBiddingDocRequest {
	s.Skip = &v
	return s
}

func (s *ListBiddingDocRequest) SetTaskName(v string) *ListBiddingDocRequest {
	s.TaskName = &v
	return s
}

func (s *ListBiddingDocRequest) SetTaskStatus(v int32) *ListBiddingDocRequest {
	s.TaskStatus = &v
	return s
}

func (s *ListBiddingDocRequest) SetWorkspaceId(v string) *ListBiddingDocRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListBiddingDocRequest) Validate() error {
	return dara.Validate(s)
}
