// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExperimentPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v string) *CreateExperimentPlanRequest
	GetDatasetId() *string
	SetDescription(v string) *CreateExperimentPlanRequest
	GetDescription() *string
	SetEvaluators(v []*Evaluator) *CreateExperimentPlanRequest
	GetEvaluators() []*Evaluator
	SetExperimentType(v string) *CreateExperimentPlanRequest
	GetExperimentType() *string
	SetExperiments(v []*ExperimentConfig) *CreateExperimentPlanRequest
	GetExperiments() []*ExperimentConfig
	SetInput(v map[string]interface{}) *CreateExperimentPlanRequest
	GetInput() map[string]interface{}
	SetPipelineName(v string) *CreateExperimentPlanRequest
	GetPipelineName() *string
	SetPlanName(v string) *CreateExperimentPlanRequest
	GetPlanName() *string
	SetQuerySql(v string) *CreateExperimentPlanRequest
	GetQuerySql() *string
	SetSelectedItemIds(v []*string) *CreateExperimentPlanRequest
	GetSelectedItemIds() []*string
}

type CreateExperimentPlanRequest struct {
	// The ID of the associated dataset. If this parameter is not specified, the execution phase processes in simple mode.
	//
	// example:
	//
	// rca_benckmark_eval
	DatasetId *string `json:"datasetId,omitempty" xml:"datasetId,omitempty"`
	// The description of the experiment plan.
	//
	// example:
	//
	// rca_benchmark_eval_experiment offline experiment.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The list of evaluators. After configuration, evaluation can be automatically triggered when the experiment completes.
	//
	// example:
	//
	// [{"evaluatorRef": "Builtin.agent_correctness", "name": "Builtin.agent_correctness", "type": "AGENT", "resultName": "Builtin.agent_correctness", "resultType": "score", "variableMapping": {"input": "experiment_input", "output": "experiment_output", "expected_output": "dataset.ground_truth_json"}, "filters": {"query": "", "sample": "100"}, "config": {"variables": [], "prompt": ""}}, {"evaluatorRef": "rca-toxicity-safety-accuracy", "name": "rca-toxicity-safety-accuracy", "type": "AGENT", "resultName": "rca-toxicity-safety-accuracy", "resultType": "score", "variableMapping": {"input": "experiment_input", "output": "experiment_output", "question": "dataset.question", "expected_output": "dataset.ground_truth_json", "payload_json": "dataset.payload_json"}, "filters": {"query": "", "sample": "100"}, "config": {"variables": [], "prompt": ""}}]
	Evaluators []*Evaluator `json:"evaluators,omitempty" xml:"evaluators,omitempty" type:"Repeated"`
	// The experiment type. Set this parameter to `OFFLINE` or `ONLINE`.
	//
	// This parameter is required.
	//
	// example:
	//
	// OFFLINE
	ExperimentType *string `json:"experimentType,omitempty" xml:"experimentType,omitempty"`
	// The list of experiment configurations. A maximum of 5 configurations are supported. For offline experiments, this parameter can be omitted or set to an empty array. For online experiments, at least one configuration is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"label": "A", "name": "experimentA", "modelName": "qwen3.7-plus", "modelProvider": "dashscope", "modelParameters": {"temperature": 0.7, "topP": 0.8, "presencePenalty": 0.0, "frequencyPenalty": 0.0}, "promptTemplate": [{"role": "system", "content": "You are an Alibaba Cloud ARMS product Q&A bot"}, {"role": "user", "content": "{{input}}"}]}, {"label": "B", "name": "experimentB", "modelName": "qwen3.7-max", "modelProvider": "dashscope", "modelParameters": {"temperature": 0.7, "topP": 0.8, "presencePenalty": 0.0, "frequencyPenalty": 0.0}, "promptTemplate": [{"role": "system", "content": "You are an Alibaba Cloud ARMS product Q&A bot"}, {"role": "user", "content": "{{input}}"}]}]
	Experiments []*ExperimentConfig `json:"experiments,omitempty" xml:"experiments,omitempty" type:"Repeated"`
	// Optional.
	//
	// example:
	//
	// {"question": "How do I get a refund?"}
	Input        map[string]interface{} `json:"input,omitempty" xml:"input,omitempty"`
	PipelineName *string                `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
	// The experiment plan name. The name must be unique within the same AgentSpace and account.
	//
	// This parameter is required.
	//
	// example:
	//
	// rca_benchmark_eval_experiment
	PlanName *string `json:"planName,omitempty" xml:"planName,omitempty"`
	// The custom query SQL clause in partial dataset mode. This parameter can be used when `selectedItemIds` is empty.
	//
	// example:
	//
	// status=\\"OK\\"
	QuerySql *string `json:"querySql,omitempty" xml:"querySql,omitempty"`
	// The list of selected data item IDs in partial dataset mode. This parameter must be used together with `datasetId`.
	//
	// example:
	//
	// ["019ef4d5-a0f0-7114-832d-5542d771cd8c", "019f1729-be9b-7769-a006-8e98023ad7ad"]
	SelectedItemIds []*string `json:"selectedItemIds,omitempty" xml:"selectedItemIds,omitempty" type:"Repeated"`
}

func (s CreateExperimentPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateExperimentPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateExperimentPlanRequest) GetDatasetId() *string {
	return s.DatasetId
}

func (s *CreateExperimentPlanRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateExperimentPlanRequest) GetEvaluators() []*Evaluator {
	return s.Evaluators
}

func (s *CreateExperimentPlanRequest) GetExperimentType() *string {
	return s.ExperimentType
}

func (s *CreateExperimentPlanRequest) GetExperiments() []*ExperimentConfig {
	return s.Experiments
}

func (s *CreateExperimentPlanRequest) GetInput() map[string]interface{} {
	return s.Input
}

func (s *CreateExperimentPlanRequest) GetPipelineName() *string {
	return s.PipelineName
}

func (s *CreateExperimentPlanRequest) GetPlanName() *string {
	return s.PlanName
}

func (s *CreateExperimentPlanRequest) GetQuerySql() *string {
	return s.QuerySql
}

func (s *CreateExperimentPlanRequest) GetSelectedItemIds() []*string {
	return s.SelectedItemIds
}

func (s *CreateExperimentPlanRequest) SetDatasetId(v string) *CreateExperimentPlanRequest {
	s.DatasetId = &v
	return s
}

func (s *CreateExperimentPlanRequest) SetDescription(v string) *CreateExperimentPlanRequest {
	s.Description = &v
	return s
}

func (s *CreateExperimentPlanRequest) SetEvaluators(v []*Evaluator) *CreateExperimentPlanRequest {
	s.Evaluators = v
	return s
}

func (s *CreateExperimentPlanRequest) SetExperimentType(v string) *CreateExperimentPlanRequest {
	s.ExperimentType = &v
	return s
}

func (s *CreateExperimentPlanRequest) SetExperiments(v []*ExperimentConfig) *CreateExperimentPlanRequest {
	s.Experiments = v
	return s
}

func (s *CreateExperimentPlanRequest) SetInput(v map[string]interface{}) *CreateExperimentPlanRequest {
	s.Input = v
	return s
}

func (s *CreateExperimentPlanRequest) SetPipelineName(v string) *CreateExperimentPlanRequest {
	s.PipelineName = &v
	return s
}

func (s *CreateExperimentPlanRequest) SetPlanName(v string) *CreateExperimentPlanRequest {
	s.PlanName = &v
	return s
}

func (s *CreateExperimentPlanRequest) SetQuerySql(v string) *CreateExperimentPlanRequest {
	s.QuerySql = &v
	return s
}

func (s *CreateExperimentPlanRequest) SetSelectedItemIds(v []*string) *CreateExperimentPlanRequest {
	s.SelectedItemIds = v
	return s
}

func (s *CreateExperimentPlanRequest) Validate() error {
	if s.Evaluators != nil {
		for _, item := range s.Evaluators {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Experiments != nil {
		for _, item := range s.Experiments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
