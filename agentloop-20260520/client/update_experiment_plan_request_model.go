// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExperimentPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v string) *UpdateExperimentPlanRequest
	GetDatasetId() *string
	SetDatasetProject(v string) *UpdateExperimentPlanRequest
	GetDatasetProject() *string
	SetDescription(v string) *UpdateExperimentPlanRequest
	GetDescription() *string
	SetEvaluators(v []*Evaluator) *UpdateExperimentPlanRequest
	GetEvaluators() []*Evaluator
	SetExperimentType(v string) *UpdateExperimentPlanRequest
	GetExperimentType() *string
	SetExperiments(v []*ExperimentConfig) *UpdateExperimentPlanRequest
	GetExperiments() []*ExperimentConfig
	SetInput(v map[string]interface{}) *UpdateExperimentPlanRequest
	GetInput() map[string]interface{}
	SetPipelineName(v string) *UpdateExperimentPlanRequest
	GetPipelineName() *string
	SetPlanName(v string) *UpdateExperimentPlanRequest
	GetPlanName() *string
	SetQuerySql(v string) *UpdateExperimentPlanRequest
	GetQuerySql() *string
	SetSelectedItemIds(v []*string) *UpdateExperimentPlanRequest
	GetSelectedItemIds() []*string
}

type UpdateExperimentPlanRequest struct {
	// The associated dataset ID.
	//
	// example:
	//
	// rca_benckmark_eval
	DatasetId *string `json:"datasetId,omitempty" xml:"datasetId,omitempty"`
	// Optional.
	//
	// example:
	//
	// agentspace-project
	DatasetProject *string `json:"datasetProject,omitempty" xml:"datasetProject,omitempty"`
	// The description.
	//
	// example:
	//
	// rca_benchmark_eval_experiment offline experiment.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The evaluator list. Omitting this field means no modification. Passing an empty array clears the list.
	//
	// example:
	//
	// [{"evaluatorRef": "Builtin.agent_task_completion"}]
	Evaluators []*Evaluator `json:"evaluators,omitempty" xml:"evaluators,omitempty" type:"Repeated"`
	// The experiment type.
	//
	// example:
	//
	// OFFLINE
	ExperimentType *string `json:"experimentType,omitempty" xml:"experimentType,omitempty"`
	// The experiment configuration list. When passed, the entire list is replaced. The number of items must be 1 to 5.
	//
	// example:
	//
	// [{"label": "A", "name": "baseline", "modelName": "qwen-max"}]
	Experiments []*ExperimentConfig `json:"experiments,omitempty" xml:"experiments,omitempty" type:"Repeated"`
	// Optional.
	//
	// example:
	//
	// {"question": "How do I get a refund?"}
	Input        map[string]interface{} `json:"input,omitempty" xml:"input,omitempty"`
	PipelineName *string                `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
	// The experiment plan name.
	//
	// example:
	//
	// rca_benchmark_eval_experiment
	PlanName *string `json:"planName,omitempty" xml:"planName,omitempty"`
	// The custom query SQL clause in partial dataset mode.
	//
	// example:
	//
	// level > 2
	QuerySql *string `json:"querySql,omitempty" xml:"querySql,omitempty"`
	// The list of selected data item IDs in partial dataset mode. Passing an empty array clears the list.
	//
	// example:
	//
	// []
	SelectedItemIds []*string `json:"selectedItemIds,omitempty" xml:"selectedItemIds,omitempty" type:"Repeated"`
}

func (s UpdateExperimentPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateExperimentPlanRequest) GoString() string {
	return s.String()
}

func (s *UpdateExperimentPlanRequest) GetDatasetId() *string {
	return s.DatasetId
}

func (s *UpdateExperimentPlanRequest) GetDatasetProject() *string {
	return s.DatasetProject
}

func (s *UpdateExperimentPlanRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateExperimentPlanRequest) GetEvaluators() []*Evaluator {
	return s.Evaluators
}

func (s *UpdateExperimentPlanRequest) GetExperimentType() *string {
	return s.ExperimentType
}

func (s *UpdateExperimentPlanRequest) GetExperiments() []*ExperimentConfig {
	return s.Experiments
}

func (s *UpdateExperimentPlanRequest) GetInput() map[string]interface{} {
	return s.Input
}

func (s *UpdateExperimentPlanRequest) GetPipelineName() *string {
	return s.PipelineName
}

func (s *UpdateExperimentPlanRequest) GetPlanName() *string {
	return s.PlanName
}

func (s *UpdateExperimentPlanRequest) GetQuerySql() *string {
	return s.QuerySql
}

func (s *UpdateExperimentPlanRequest) GetSelectedItemIds() []*string {
	return s.SelectedItemIds
}

func (s *UpdateExperimentPlanRequest) SetDatasetId(v string) *UpdateExperimentPlanRequest {
	s.DatasetId = &v
	return s
}

func (s *UpdateExperimentPlanRequest) SetDatasetProject(v string) *UpdateExperimentPlanRequest {
	s.DatasetProject = &v
	return s
}

func (s *UpdateExperimentPlanRequest) SetDescription(v string) *UpdateExperimentPlanRequest {
	s.Description = &v
	return s
}

func (s *UpdateExperimentPlanRequest) SetEvaluators(v []*Evaluator) *UpdateExperimentPlanRequest {
	s.Evaluators = v
	return s
}

func (s *UpdateExperimentPlanRequest) SetExperimentType(v string) *UpdateExperimentPlanRequest {
	s.ExperimentType = &v
	return s
}

func (s *UpdateExperimentPlanRequest) SetExperiments(v []*ExperimentConfig) *UpdateExperimentPlanRequest {
	s.Experiments = v
	return s
}

func (s *UpdateExperimentPlanRequest) SetInput(v map[string]interface{}) *UpdateExperimentPlanRequest {
	s.Input = v
	return s
}

func (s *UpdateExperimentPlanRequest) SetPipelineName(v string) *UpdateExperimentPlanRequest {
	s.PipelineName = &v
	return s
}

func (s *UpdateExperimentPlanRequest) SetPlanName(v string) *UpdateExperimentPlanRequest {
	s.PlanName = &v
	return s
}

func (s *UpdateExperimentPlanRequest) SetQuerySql(v string) *UpdateExperimentPlanRequest {
	s.QuerySql = &v
	return s
}

func (s *UpdateExperimentPlanRequest) SetSelectedItemIds(v []*string) *UpdateExperimentPlanRequest {
	s.SelectedItemIds = v
	return s
}

func (s *UpdateExperimentPlanRequest) Validate() error {
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
