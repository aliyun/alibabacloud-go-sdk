// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndOfCreateTime(v string) *ListWorkflowTasksRequest
	GetEndOfCreateTime() *string
	SetKeyText(v string) *ListWorkflowTasksRequest
	GetKeyText() *string
	SetMaxResults(v int32) *ListWorkflowTasksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListWorkflowTasksRequest
	GetNextToken() *string
	SetStartOfCreateTime(v string) *ListWorkflowTasksRequest
	GetStartOfCreateTime() *string
	SetWorkflowId(v string) *ListWorkflowTasksRequest
	GetWorkflowId() *string
	SetWorkflowName(v string) *ListWorkflowTasksRequest
	GetWorkflowName() *string
}

type ListWorkflowTasksRequest struct {
	// The end of the time range for filtering tasks by their creation time. Supports querying data from the last 90 days only.
	//
	// example:
	//
	// 2025-07-15T00:00:00Z
	EndOfCreateTime *string `json:"EndOfCreateTime,omitempty" xml:"EndOfCreateTime,omitempty"`
	// A keyword for fuzzy matching against the TaskInput, such as a file name or Media ID. Max length: 32 characters.
	KeyText *string `json:"KeyText,omitempty" xml:"KeyText,omitempty"`
	// The maximum number of media workflow instances to return. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request.
	//
	// example:
	//
	// **************VRpbWUQARgBIpcBCgkA1bUtaAAAAAAKiQEDhAAAADFTMzg2NTY2NjU2MzM3NjU2NjYyMzkzMTYyMzI2MjYzNjY2**********
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The start of the time range for filtering tasks by their creation time. Supports querying data from the last 90 days only.
	//
	// example:
	//
	// 2025-07-12T00:00:00Z
	StartOfCreateTime *string `json:"StartOfCreateTime,omitempty" xml:"StartOfCreateTime,omitempty"`
	// The ID of the workflow template.[](https://ims.console.aliyun.com/settings/workflow/list)
	//
	// example:
	//
	// ******b4fb044839815d4f2cd8******
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
	// The name of the workflow template.
	//
	// example:
	//
	// example-workflow-****
	WorkflowName *string `json:"WorkflowName,omitempty" xml:"WorkflowName,omitempty"`
}

func (s ListWorkflowTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowTasksRequest) GoString() string {
	return s.String()
}

func (s *ListWorkflowTasksRequest) GetEndOfCreateTime() *string {
	return s.EndOfCreateTime
}

func (s *ListWorkflowTasksRequest) GetKeyText() *string {
	return s.KeyText
}

func (s *ListWorkflowTasksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkflowTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkflowTasksRequest) GetStartOfCreateTime() *string {
	return s.StartOfCreateTime
}

func (s *ListWorkflowTasksRequest) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *ListWorkflowTasksRequest) GetWorkflowName() *string {
	return s.WorkflowName
}

func (s *ListWorkflowTasksRequest) SetEndOfCreateTime(v string) *ListWorkflowTasksRequest {
	s.EndOfCreateTime = &v
	return s
}

func (s *ListWorkflowTasksRequest) SetKeyText(v string) *ListWorkflowTasksRequest {
	s.KeyText = &v
	return s
}

func (s *ListWorkflowTasksRequest) SetMaxResults(v int32) *ListWorkflowTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkflowTasksRequest) SetNextToken(v string) *ListWorkflowTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkflowTasksRequest) SetStartOfCreateTime(v string) *ListWorkflowTasksRequest {
	s.StartOfCreateTime = &v
	return s
}

func (s *ListWorkflowTasksRequest) SetWorkflowId(v string) *ListWorkflowTasksRequest {
	s.WorkflowId = &v
	return s
}

func (s *ListWorkflowTasksRequest) SetWorkflowName(v string) *ListWorkflowTasksRequest {
	s.WorkflowName = &v
	return s
}

func (s *ListWorkflowTasksRequest) Validate() error {
	return dara.Validate(s)
}
