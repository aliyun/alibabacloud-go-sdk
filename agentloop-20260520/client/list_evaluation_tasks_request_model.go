// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEvaluationTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpace(v string) *ListEvaluationTasksRequest
	GetAgentSpace() *string
	SetChannel(v string) *ListEvaluationTasksRequest
	GetChannel() *string
	SetDataType(v string) *ListEvaluationTasksRequest
	GetDataType() *string
	SetMaxResults(v int32) *ListEvaluationTasksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListEvaluationTasksRequest
	GetNextToken() *string
	SetStatus(v string) *ListEvaluationTasksRequest
	GetStatus() *string
	SetTaskMode(v string) *ListEvaluationTasksRequest
	GetTaskMode() *string
	SetTaskName(v string) *ListEvaluationTasksRequest
	GetTaskName() *string
}

type ListEvaluationTasksRequest struct {
	// The AgentSpace name.
	//
	// example:
	//
	// prod-agentspace
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// The filter condition for the task source. If this parameter is not specified, tasks from the default source are queried.
	//
	// example:
	//
	// default
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// The data source type of the evaluation object. Set this parameter to `trace` for trace-based evaluation.
	//
	// example:
	//
	// trace
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token for the next page, obtained from the previous response.
	//
	// example:
	//
	// eyJsYXN0SWQiOjEyMywib2Zmc2V0IjoyMH0=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The filter condition for the evaluation task status.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The evaluation task mode. If this parameter is not specified, the default value is `batch`.
	//
	// example:
	//
	// batch
	TaskMode *string `json:"taskMode,omitempty" xml:"taskMode,omitempty"`
	// The fuzzy match condition for the task name.
	//
	// example:
	//
	// trace_task_completion_eval
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s ListEvaluationTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationTasksRequest) GoString() string {
	return s.String()
}

func (s *ListEvaluationTasksRequest) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *ListEvaluationTasksRequest) GetChannel() *string {
	return s.Channel
}

func (s *ListEvaluationTasksRequest) GetDataType() *string {
	return s.DataType
}

func (s *ListEvaluationTasksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListEvaluationTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEvaluationTasksRequest) GetStatus() *string {
	return s.Status
}

func (s *ListEvaluationTasksRequest) GetTaskMode() *string {
	return s.TaskMode
}

func (s *ListEvaluationTasksRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *ListEvaluationTasksRequest) SetAgentSpace(v string) *ListEvaluationTasksRequest {
	s.AgentSpace = &v
	return s
}

func (s *ListEvaluationTasksRequest) SetChannel(v string) *ListEvaluationTasksRequest {
	s.Channel = &v
	return s
}

func (s *ListEvaluationTasksRequest) SetDataType(v string) *ListEvaluationTasksRequest {
	s.DataType = &v
	return s
}

func (s *ListEvaluationTasksRequest) SetMaxResults(v int32) *ListEvaluationTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListEvaluationTasksRequest) SetNextToken(v string) *ListEvaluationTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListEvaluationTasksRequest) SetStatus(v string) *ListEvaluationTasksRequest {
	s.Status = &v
	return s
}

func (s *ListEvaluationTasksRequest) SetTaskMode(v string) *ListEvaluationTasksRequest {
	s.TaskMode = &v
	return s
}

func (s *ListEvaluationTasksRequest) SetTaskName(v string) *ListEvaluationTasksRequest {
	s.TaskName = &v
	return s
}

func (s *ListEvaluationTasksRequest) Validate() error {
	return dara.Validate(s)
}
