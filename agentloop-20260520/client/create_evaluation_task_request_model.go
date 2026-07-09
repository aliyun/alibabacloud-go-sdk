// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEvaluationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannel(v string) *CreateEvaluationTaskRequest
	GetChannel() *string
	SetConfig(v map[string]*string) *CreateEvaluationTaskRequest
	GetConfig() map[string]*string
	SetDataFilter(v string) *CreateEvaluationTaskRequest
	GetDataFilter() *string
	SetDataType(v string) *CreateEvaluationTaskRequest
	GetDataType() *string
	SetDescription(v string) *CreateEvaluationTaskRequest
	GetDescription() *string
	SetEvaluators(v []*Evaluator) *CreateEvaluationTaskRequest
	GetEvaluators() []*Evaluator
	SetRunStrategies(v *RunStrategies) *CreateEvaluationTaskRequest
	GetRunStrategies() *RunStrategies
	SetTags(v map[string]*string) *CreateEvaluationTaskRequest
	GetTags() map[string]*string
	SetTaskMode(v string) *CreateEvaluationTaskRequest
	GetTaskMode() *string
	SetTaskName(v string) *CreateEvaluationTaskRequest
	GetTaskName() *string
	SetClientToken(v string) *CreateEvaluationTaskRequest
	GetClientToken() *string
}

type CreateEvaluationTaskRequest struct {
	// The task source. If this parameter is not specified, the backend uses `default`.
	//
	// example:
	//
	// default
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// The data source and execution configuration. When `dataType` is set to `trace`, the backend automatically populates the SLS Project and sets `storeName` to `logstore-tracing`. For trace-level evaluation, set `dataScope` to `trace`.
	//
	// example:
	//
	// {"dataScope":"trace"}
	Config map[string]*string `json:"config,omitempty" xml:"config,omitempty"`
	// The filter conditions for evaluation data. This parameter supports a JSON object or a JSON string. Common fields include `query`, `provided`, `maxRecords`, and `samplingRate`.
	//
	// example:
	//
	// {"query":"serviceName=\\"checkout-service\\"","maxRecords":10,"samplingRate":100}
	DataFilter *string `json:"dataFilter,omitempty" xml:"dataFilter,omitempty"`
	// The data source type of the evaluation object. Set this parameter to `trace` for trace-based evaluation.
	//
	// example:
	//
	// trace
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// The description of the evaluation task.
	//
	// example:
	//
	// 评估线上 Agent 链路任务完成度
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The list of evaluator configurations. This parameter cannot be empty. Within the same task, `evaluatorRef` takes precedence as the unique identifier. Otherwise, `name` is used.
	//
	// example:
	//
	// [{"evaluatorRef":"Builtin.agent_task_completion","resultName":"agent_task_completion","resultType":"score","variableMapping":{"input":"trace.input","output":"trace.output","agent_trajectory":"trace.agent_trajectory"}}]
	Evaluators []*Evaluator `json:"evaluators,omitempty" xml:"evaluators,omitempty" type:"Repeated"`
	// The task execution strategy. This parameter supports a JSON object or a JSON string. Set this parameter to `backfill` for historical data backfill or `continuous` for continuous evaluation of new data.
	RunStrategies *RunStrategies `json:"runStrategies,omitempty" xml:"runStrategies,omitempty"`
	// The key-value pairs of task tags. You do not need to specify this parameter by default. Specify this parameter only when you want to associate or filter tasks by business tags.
	//
	// example:
	//
	// {"env":"prod","serviceId":"checkout-service","planId":"plan-20260703"}
	Tags map[string]*string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The evaluation task mode. Set this parameter to `batch` to create a persistent evaluation task.
	//
	// example:
	//
	// batch
	TaskMode *string `json:"taskMode,omitempty" xml:"taskMode,omitempty"`
	// The task name. The name must be unique among non-deleted tasks within the same user and AgentSpace. The name can be up to 256 characters in length.
	//
	// example:
	//
	// trace_task_completion_eval
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
	// The idempotency token. This query parameter is declared in CloudSpec, but the backend does not currently perform idempotency comparison.
	//
	// example:
	//
	// a1b2c3d4-1234-5678-90ab-cdef12345678
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateEvaluationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEvaluationTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateEvaluationTaskRequest) GetChannel() *string {
	return s.Channel
}

func (s *CreateEvaluationTaskRequest) GetConfig() map[string]*string {
	return s.Config
}

func (s *CreateEvaluationTaskRequest) GetDataFilter() *string {
	return s.DataFilter
}

func (s *CreateEvaluationTaskRequest) GetDataType() *string {
	return s.DataType
}

func (s *CreateEvaluationTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateEvaluationTaskRequest) GetEvaluators() []*Evaluator {
	return s.Evaluators
}

func (s *CreateEvaluationTaskRequest) GetRunStrategies() *RunStrategies {
	return s.RunStrategies
}

func (s *CreateEvaluationTaskRequest) GetTags() map[string]*string {
	return s.Tags
}

func (s *CreateEvaluationTaskRequest) GetTaskMode() *string {
	return s.TaskMode
}

func (s *CreateEvaluationTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateEvaluationTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateEvaluationTaskRequest) SetChannel(v string) *CreateEvaluationTaskRequest {
	s.Channel = &v
	return s
}

func (s *CreateEvaluationTaskRequest) SetConfig(v map[string]*string) *CreateEvaluationTaskRequest {
	s.Config = v
	return s
}

func (s *CreateEvaluationTaskRequest) SetDataFilter(v string) *CreateEvaluationTaskRequest {
	s.DataFilter = &v
	return s
}

func (s *CreateEvaluationTaskRequest) SetDataType(v string) *CreateEvaluationTaskRequest {
	s.DataType = &v
	return s
}

func (s *CreateEvaluationTaskRequest) SetDescription(v string) *CreateEvaluationTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateEvaluationTaskRequest) SetEvaluators(v []*Evaluator) *CreateEvaluationTaskRequest {
	s.Evaluators = v
	return s
}

func (s *CreateEvaluationTaskRequest) SetRunStrategies(v *RunStrategies) *CreateEvaluationTaskRequest {
	s.RunStrategies = v
	return s
}

func (s *CreateEvaluationTaskRequest) SetTags(v map[string]*string) *CreateEvaluationTaskRequest {
	s.Tags = v
	return s
}

func (s *CreateEvaluationTaskRequest) SetTaskMode(v string) *CreateEvaluationTaskRequest {
	s.TaskMode = &v
	return s
}

func (s *CreateEvaluationTaskRequest) SetTaskName(v string) *CreateEvaluationTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateEvaluationTaskRequest) SetClientToken(v string) *CreateEvaluationTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEvaluationTaskRequest) Validate() error {
	if s.Evaluators != nil {
		for _, item := range s.Evaluators {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RunStrategies != nil {
		if err := s.RunStrategies.Validate(); err != nil {
			return err
		}
	}
	return nil
}
