// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEvaluationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpace(v string) *GetEvaluationTaskResponseBody
	GetAgentSpace() *string
	SetChannel(v string) *GetEvaluationTaskResponseBody
	GetChannel() *string
	SetConfig(v map[string]*string) *GetEvaluationTaskResponseBody
	GetConfig() map[string]*string
	SetCreatedAt(v int64) *GetEvaluationTaskResponseBody
	GetCreatedAt() *int64
	SetDataFilter(v string) *GetEvaluationTaskResponseBody
	GetDataFilter() *string
	SetDataType(v string) *GetEvaluationTaskResponseBody
	GetDataType() *string
	SetDescription(v string) *GetEvaluationTaskResponseBody
	GetDescription() *string
	SetEvaluators(v []*Evaluator) *GetEvaluationTaskResponseBody
	GetEvaluators() []*Evaluator
	SetRegionId(v string) *GetEvaluationTaskResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetEvaluationTaskResponseBody
	GetRequestId() *string
	SetRunStrategyConfig(v *RunStrategies) *GetEvaluationTaskResponseBody
	GetRunStrategyConfig() *RunStrategies
	SetStatus(v string) *GetEvaluationTaskResponseBody
	GetStatus() *string
	SetTags(v map[string]*string) *GetEvaluationTaskResponseBody
	GetTags() map[string]*string
	SetTaskId(v string) *GetEvaluationTaskResponseBody
	GetTaskId() *string
	SetTaskMode(v string) *GetEvaluationTaskResponseBody
	GetTaskMode() *string
	SetTaskName(v string) *GetEvaluationTaskResponseBody
	GetTaskName() *string
	SetUpdatedAt(v int64) *GetEvaluationTaskResponseBody
	GetUpdatedAt() *int64
}

type GetEvaluationTaskResponseBody struct {
	// The AgentSpace name.
	//
	// example:
	//
	// prod-agentspace
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// The task source.
	//
	// example:
	//
	// default
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// The data source and execution configuration. Tasks with `dataType=trace` typically include `project`, `storeName`, and `dataScope` fields populated by the backend.
	//
	// example:
	//
	// {"project":"agentspace-project","storeName":"logstore-tracing","dataScope":"trace"}
	Config map[string]*string `json:"config,omitempty" xml:"config,omitempty"`
	// The creation time, in seconds-level UNIX timestamp.
	//
	// example:
	//
	// 1782816000
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The filter conditions for evaluation data, returned by the backend as a JSON string.
	//
	// example:
	//
	// {"query":"serviceName=\\"checkout-service\\"","maxRecords":10,"samplingRate":100}
	DataFilter *string `json:"dataFilter,omitempty" xml:"dataFilter,omitempty"`
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
	// The list of evaluator configurations.
	//
	// example:
	//
	// [{"evaluatorRef":"Builtin.agent_task_completion","resultName":"agent_task_completion","resultType":"score","variableMapping":{"input":"trace.input","output":"trace.output","agent_trajectory":"trace.agent_trajectory"}}]
	Evaluators []*Evaluator `json:"evaluators,omitempty" xml:"evaluators,omitempty" type:"Repeated"`
	// The region to which the task belongs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3FE4CD1E-FF41-56BE-B590-7A021D9C1524
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The structured configuration of the run strategy, including the parsed backfill strategy and continuous evaluation strategy.
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
	// The key-value pairs of task tags. Empty if not set.
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

func (s GetEvaluationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEvaluationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetEvaluationTaskResponseBody) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *GetEvaluationTaskResponseBody) GetChannel() *string {
	return s.Channel
}

func (s *GetEvaluationTaskResponseBody) GetConfig() map[string]*string {
	return s.Config
}

func (s *GetEvaluationTaskResponseBody) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *GetEvaluationTaskResponseBody) GetDataFilter() *string {
	return s.DataFilter
}

func (s *GetEvaluationTaskResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *GetEvaluationTaskResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetEvaluationTaskResponseBody) GetEvaluators() []*Evaluator {
	return s.Evaluators
}

func (s *GetEvaluationTaskResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetEvaluationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEvaluationTaskResponseBody) GetRunStrategyConfig() *RunStrategies {
	return s.RunStrategyConfig
}

func (s *GetEvaluationTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetEvaluationTaskResponseBody) GetTags() map[string]*string {
	return s.Tags
}

func (s *GetEvaluationTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetEvaluationTaskResponseBody) GetTaskMode() *string {
	return s.TaskMode
}

func (s *GetEvaluationTaskResponseBody) GetTaskName() *string {
	return s.TaskName
}

func (s *GetEvaluationTaskResponseBody) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *GetEvaluationTaskResponseBody) SetAgentSpace(v string) *GetEvaluationTaskResponseBody {
	s.AgentSpace = &v
	return s
}

func (s *GetEvaluationTaskResponseBody) SetChannel(v string) *GetEvaluationTaskResponseBody {
	s.Channel = &v
	return s
}

func (s *GetEvaluationTaskResponseBody) SetConfig(v map[string]*string) *GetEvaluationTaskResponseBody {
	s.Config = v
	return s
}

func (s *GetEvaluationTaskResponseBody) SetCreatedAt(v int64) *GetEvaluationTaskResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *GetEvaluationTaskResponseBody) SetDataFilter(v string) *GetEvaluationTaskResponseBody {
	s.DataFilter = &v
	return s
}

func (s *GetEvaluationTaskResponseBody) SetDataType(v string) *GetEvaluationTaskResponseBody {
	s.DataType = &v
	return s
}

func (s *GetEvaluationTaskResponseBody) SetDescription(v string) *GetEvaluationTaskResponseBody {
	s.Description = &v
	return s
}

func (s *GetEvaluationTaskResponseBody) SetEvaluators(v []*Evaluator) *GetEvaluationTaskResponseBody {
	s.Evaluators = v
	return s
}

func (s *GetEvaluationTaskResponseBody) SetRegionId(v string) *GetEvaluationTaskResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetEvaluationTaskResponseBody) SetRequestId(v string) *GetEvaluationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEvaluationTaskResponseBody) SetRunStrategyConfig(v *RunStrategies) *GetEvaluationTaskResponseBody {
	s.RunStrategyConfig = v
	return s
}

func (s *GetEvaluationTaskResponseBody) SetStatus(v string) *GetEvaluationTaskResponseBody {
	s.Status = &v
	return s
}

func (s *GetEvaluationTaskResponseBody) SetTags(v map[string]*string) *GetEvaluationTaskResponseBody {
	s.Tags = v
	return s
}

func (s *GetEvaluationTaskResponseBody) SetTaskId(v string) *GetEvaluationTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetEvaluationTaskResponseBody) SetTaskMode(v string) *GetEvaluationTaskResponseBody {
	s.TaskMode = &v
	return s
}

func (s *GetEvaluationTaskResponseBody) SetTaskName(v string) *GetEvaluationTaskResponseBody {
	s.TaskName = &v
	return s
}

func (s *GetEvaluationTaskResponseBody) SetUpdatedAt(v int64) *GetEvaluationTaskResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *GetEvaluationTaskResponseBody) Validate() error {
	if s.Evaluators != nil {
		for _, item := range s.Evaluators {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RunStrategyConfig != nil {
		if err := s.RunStrategyConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
