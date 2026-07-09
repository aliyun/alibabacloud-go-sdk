// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEvaluationTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvaluationTasks(v []*ListEvaluationTasksResponseBodyEvaluationTasks) *ListEvaluationTasksResponseBody
	GetEvaluationTasks() []*ListEvaluationTasksResponseBodyEvaluationTasks
	SetMaxResults(v int32) *ListEvaluationTasksResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListEvaluationTasksResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListEvaluationTasksResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListEvaluationTasksResponseBody
	GetTotalCount() *int32
}

type ListEvaluationTasksResponseBody struct {
	// The list of evaluation task summaries.
	//
	// example:
	//
	// [{"taskId":"eval-task-8b36f2e2b1f94f9c91ce7a4b0f6d9c25","taskName":"trace_task_completion_eval","taskMode":"batch","dataType":"trace","status":"Running"}]
	EvaluationTasks []*ListEvaluationTasksResponseBodyEvaluationTasks `json:"evaluationTasks,omitempty" xml:"evaluationTasks,omitempty" type:"Repeated"`
	// The number of entries per page used in this request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token for the next page. An empty value indicates that no more pages are available.
	//
	// example:
	//
	// eyJsYXN0SWQiOjEwMSwib2Zmc2V0IjoyMH0=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3FE4CD1E-FF41-56BE-B590-7A021D9C1524
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of records. The total count is returned only on the first page. This value may be empty on subsequent pages.
	//
	// example:
	//
	// 126
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListEvaluationTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListEvaluationTasksResponseBody) GetEvaluationTasks() []*ListEvaluationTasksResponseBodyEvaluationTasks {
	return s.EvaluationTasks
}

func (s *ListEvaluationTasksResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListEvaluationTasksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEvaluationTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEvaluationTasksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListEvaluationTasksResponseBody) SetEvaluationTasks(v []*ListEvaluationTasksResponseBodyEvaluationTasks) *ListEvaluationTasksResponseBody {
	s.EvaluationTasks = v
	return s
}

func (s *ListEvaluationTasksResponseBody) SetMaxResults(v int32) *ListEvaluationTasksResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListEvaluationTasksResponseBody) SetNextToken(v string) *ListEvaluationTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListEvaluationTasksResponseBody) SetRequestId(v string) *ListEvaluationTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEvaluationTasksResponseBody) SetTotalCount(v int32) *ListEvaluationTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEvaluationTasksResponseBody) Validate() error {
	if s.EvaluationTasks != nil {
		for _, item := range s.EvaluationTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEvaluationTasksResponseBodyEvaluationTasks struct {
	// The data source and execution configuration summary.
	//
	// example:
	//
	// {"storeName":"logstore-tracing","dataScope":"trace"}
	Config map[string]*string `json:"config,omitempty" xml:"config,omitempty"`
	// The creation time, in seconds-level UNIX timestamp.
	//
	// example:
	//
	// 1782816000
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The data source type of the evaluation object.
	//
	// example:
	//
	// trace
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// The evaluation task description.
	//
	// example:
	//
	// 评估线上 Agent 链路任务完成度
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The evaluator configuration summary, in JSON string format.
	//
	// example:
	//
	// [{"evaluatorRef":"Builtin.agent_task_completion"}]
	Evaluators *string `json:"evaluators,omitempty" xml:"evaluators,omitempty"`
	// The structured run strategy configuration, including the parsed backfill strategy and continuous evaluation strategy.
	//
	// example:
	//
	// {"backfill":{"enabled":true,"startTime":1782816000000,"endTime":1782902400000},"continuous":{"enabled":true,"intervalUnit":"HOUR","intervalValue":1,"dataDelayMinutes":5}}
	RunStrategyConfig *RunStrategies `json:"runStrategyConfig,omitempty" xml:"runStrategyConfig,omitempty"`
	// The evaluation task status.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The key-value pairs of task tags. This parameter is empty if no tags are set.
	//
	// example:
	//
	// {"serviceId":"checkout-service","env":"prod"}
	Tags map[string]*string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The evaluation task ID.
	//
	// example:
	//
	// eval-task-8b36f2e2b1f94f9c91ce7a4b0f6d9c25
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// The evaluation task mode.
	//
	// example:
	//
	// batch
	TaskMode *string `json:"taskMode,omitempty" xml:"taskMode,omitempty"`
	// The task name.
	//
	// example:
	//
	// trace_task_completion_eval
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
	// The last update time, in seconds-level UNIX timestamp.
	//
	// example:
	//
	// 1782816600
	UpdatedAt *int64 `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s ListEvaluationTasksResponseBodyEvaluationTasks) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationTasksResponseBodyEvaluationTasks) GoString() string {
	return s.String()
}

func (s *ListEvaluationTasksResponseBodyEvaluationTasks) GetConfig() map[string]*string {
	return s.Config
}

func (s *ListEvaluationTasksResponseBodyEvaluationTasks) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *ListEvaluationTasksResponseBodyEvaluationTasks) GetDataType() *string {
	return s.DataType
}

func (s *ListEvaluationTasksResponseBodyEvaluationTasks) GetDescription() *string {
	return s.Description
}

func (s *ListEvaluationTasksResponseBodyEvaluationTasks) GetEvaluators() *string {
	return s.Evaluators
}

func (s *ListEvaluationTasksResponseBodyEvaluationTasks) GetRunStrategyConfig() *RunStrategies {
	return s.RunStrategyConfig
}

func (s *ListEvaluationTasksResponseBodyEvaluationTasks) GetStatus() *string {
	return s.Status
}

func (s *ListEvaluationTasksResponseBodyEvaluationTasks) GetTags() map[string]*string {
	return s.Tags
}

func (s *ListEvaluationTasksResponseBodyEvaluationTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *ListEvaluationTasksResponseBodyEvaluationTasks) GetTaskMode() *string {
	return s.TaskMode
}

func (s *ListEvaluationTasksResponseBodyEvaluationTasks) GetTaskName() *string {
	return s.TaskName
}

func (s *ListEvaluationTasksResponseBodyEvaluationTasks) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *ListEvaluationTasksResponseBodyEvaluationTasks) SetConfig(v map[string]*string) *ListEvaluationTasksResponseBodyEvaluationTasks {
	s.Config = v
	return s
}

func (s *ListEvaluationTasksResponseBodyEvaluationTasks) SetCreatedAt(v int64) *ListEvaluationTasksResponseBodyEvaluationTasks {
	s.CreatedAt = &v
	return s
}

func (s *ListEvaluationTasksResponseBodyEvaluationTasks) SetDataType(v string) *ListEvaluationTasksResponseBodyEvaluationTasks {
	s.DataType = &v
	return s
}

func (s *ListEvaluationTasksResponseBodyEvaluationTasks) SetDescription(v string) *ListEvaluationTasksResponseBodyEvaluationTasks {
	s.Description = &v
	return s
}

func (s *ListEvaluationTasksResponseBodyEvaluationTasks) SetEvaluators(v string) *ListEvaluationTasksResponseBodyEvaluationTasks {
	s.Evaluators = &v
	return s
}

func (s *ListEvaluationTasksResponseBodyEvaluationTasks) SetRunStrategyConfig(v *RunStrategies) *ListEvaluationTasksResponseBodyEvaluationTasks {
	s.RunStrategyConfig = v
	return s
}

func (s *ListEvaluationTasksResponseBodyEvaluationTasks) SetStatus(v string) *ListEvaluationTasksResponseBodyEvaluationTasks {
	s.Status = &v
	return s
}

func (s *ListEvaluationTasksResponseBodyEvaluationTasks) SetTags(v map[string]*string) *ListEvaluationTasksResponseBodyEvaluationTasks {
	s.Tags = v
	return s
}

func (s *ListEvaluationTasksResponseBodyEvaluationTasks) SetTaskId(v string) *ListEvaluationTasksResponseBodyEvaluationTasks {
	s.TaskId = &v
	return s
}

func (s *ListEvaluationTasksResponseBodyEvaluationTasks) SetTaskMode(v string) *ListEvaluationTasksResponseBodyEvaluationTasks {
	s.TaskMode = &v
	return s
}

func (s *ListEvaluationTasksResponseBodyEvaluationTasks) SetTaskName(v string) *ListEvaluationTasksResponseBodyEvaluationTasks {
	s.TaskName = &v
	return s
}

func (s *ListEvaluationTasksResponseBodyEvaluationTasks) SetUpdatedAt(v int64) *ListEvaluationTasksResponseBodyEvaluationTasks {
	s.UpdatedAt = &v
	return s
}

func (s *ListEvaluationTasksResponseBodyEvaluationTasks) Validate() error {
	if s.RunStrategyConfig != nil {
		if err := s.RunStrategyConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
