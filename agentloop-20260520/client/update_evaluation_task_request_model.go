// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEvaluationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v map[string]*string) *UpdateEvaluationTaskRequest
	GetConfig() map[string]*string
	SetDataFilter(v string) *UpdateEvaluationTaskRequest
	GetDataFilter() *string
	SetDescription(v string) *UpdateEvaluationTaskRequest
	GetDescription() *string
	SetEvaluators(v []*Evaluator) *UpdateEvaluationTaskRequest
	GetEvaluators() []*Evaluator
	SetRunStrategies(v *RunStrategies) *UpdateEvaluationTaskRequest
	GetRunStrategies() *RunStrategies
	SetStatus(v string) *UpdateEvaluationTaskRequest
	GetStatus() *string
	SetTags(v map[string]*string) *UpdateEvaluationTaskRequest
	GetTags() map[string]*string
	SetClientToken(v string) *UpdateEvaluationTaskRequest
	GetClientToken() *string
}

type UpdateEvaluationTaskRequest struct {
	// The new task configuration. Some fields that are set during creation cannot be modified.
	//
	// example:
	//
	// {"dataScope":"trace"}
	Config map[string]*string `json:"config,omitempty" xml:"config,omitempty"`
	// The filter condition for evaluation data. A JSON object or JSON string is supported.
	//
	// example:
	//
	// {"query":"serviceName=\\"checkout-service\\" AND status=\\"OK\\"","maxRecords":10,"samplingRate":50}
	DataFilter *string `json:"dataFilter,omitempty" xml:"dataFilter,omitempty"`
	// The description of the evaluation task.
	//
	// example:
	//
	// 更新后的链路 Trace 任务完成度评估
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The new list of evaluator configurations. When specified, this list entirely replaces the existing evaluator list of the task, and the system re-validates evaluator uniqueness and variable mappings.
	//
	// example:
	//
	// [{"evaluatorRef":"Builtin.agent_task_completion","resultName":"agent_task_completion","resultType":"score","variableMapping":{"input":"trace.input","output":"trace.output","agent_trajectory":"trace.agent_trajectory"}}]
	Evaluators []*Evaluator `json:"evaluators,omitempty" xml:"evaluators,omitempty" type:"Repeated"`
	// The new task execution strategies. A JSON object or JSON string is supported. If the task is in `Completed`, `Terminated`, or `Failed` status and the new strategy enables backfill or continuous mode, the backend resets the task to `Pending` and triggers orchestration.
	RunStrategies *RunStrategies `json:"runStrategies,omitempty" xml:"runStrategies,omitempty"`
	// The task status. Currently the backend only allows users to explicitly set this to `Terminated`. Other statuses are managed by the system.
	//
	// example:
	//
	// Terminated
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The key-value pairs of task tags. You do not need to specify this parameter by default. Specify this parameter only when you want to associate or filter tasks by business tags.
	//
	// example:
	//
	// {"env":"prod","serviceId":"checkout-service","planId":"plan-20260703"}
	Tags map[string]*string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The idempotency token. CloudSpec declares this query parameter, but the backend does not currently perform idempotency comparison.
	//
	// example:
	//
	// a1b2c3d4-1234-5678-90ab-cdef12345678
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateEvaluationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEvaluationTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateEvaluationTaskRequest) GetConfig() map[string]*string {
	return s.Config
}

func (s *UpdateEvaluationTaskRequest) GetDataFilter() *string {
	return s.DataFilter
}

func (s *UpdateEvaluationTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateEvaluationTaskRequest) GetEvaluators() []*Evaluator {
	return s.Evaluators
}

func (s *UpdateEvaluationTaskRequest) GetRunStrategies() *RunStrategies {
	return s.RunStrategies
}

func (s *UpdateEvaluationTaskRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateEvaluationTaskRequest) GetTags() map[string]*string {
	return s.Tags
}

func (s *UpdateEvaluationTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateEvaluationTaskRequest) SetConfig(v map[string]*string) *UpdateEvaluationTaskRequest {
	s.Config = v
	return s
}

func (s *UpdateEvaluationTaskRequest) SetDataFilter(v string) *UpdateEvaluationTaskRequest {
	s.DataFilter = &v
	return s
}

func (s *UpdateEvaluationTaskRequest) SetDescription(v string) *UpdateEvaluationTaskRequest {
	s.Description = &v
	return s
}

func (s *UpdateEvaluationTaskRequest) SetEvaluators(v []*Evaluator) *UpdateEvaluationTaskRequest {
	s.Evaluators = v
	return s
}

func (s *UpdateEvaluationTaskRequest) SetRunStrategies(v *RunStrategies) *UpdateEvaluationTaskRequest {
	s.RunStrategies = v
	return s
}

func (s *UpdateEvaluationTaskRequest) SetStatus(v string) *UpdateEvaluationTaskRequest {
	s.Status = &v
	return s
}

func (s *UpdateEvaluationTaskRequest) SetTags(v map[string]*string) *UpdateEvaluationTaskRequest {
	s.Tags = v
	return s
}

func (s *UpdateEvaluationTaskRequest) SetClientToken(v string) *UpdateEvaluationTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateEvaluationTaskRequest) Validate() error {
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
