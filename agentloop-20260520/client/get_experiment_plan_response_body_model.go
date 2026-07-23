// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperimentPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v int64) *GetExperimentPlanResponseBody
	GetCreatedAt() *int64
	SetDatasetId(v string) *GetExperimentPlanResponseBody
	GetDatasetId() *string
	SetDescription(v string) *GetExperimentPlanResponseBody
	GetDescription() *string
	SetEvaluators(v []*Evaluator) *GetExperimentPlanResponseBody
	GetEvaluators() []*Evaluator
	SetExperimentType(v string) *GetExperimentPlanResponseBody
	GetExperimentType() *string
	SetExperiments(v []*ExperimentConfig) *GetExperimentPlanResponseBody
	GetExperiments() []*ExperimentConfig
	SetInput(v map[string]interface{}) *GetExperimentPlanResponseBody
	GetInput() map[string]interface{}
	SetPlanId(v string) *GetExperimentPlanResponseBody
	GetPlanId() *string
	SetPlanName(v string) *GetExperimentPlanResponseBody
	GetPlanName() *string
	SetQuerySql(v string) *GetExperimentPlanResponseBody
	GetQuerySql() *string
	SetRequestId(v string) *GetExperimentPlanResponseBody
	GetRequestId() *string
	SetSelectedItemIds(v []*string) *GetExperimentPlanResponseBody
	GetSelectedItemIds() []*string
	SetStatus(v string) *GetExperimentPlanResponseBody
	GetStatus() *string
	SetUpdatedAt(v int64) *GetExperimentPlanResponseBody
	GetUpdatedAt() *int64
}

type GetExperimentPlanResponseBody struct {
	// The creation time, in millisecond Unix timestamp.
	//
	// example:
	//
	// 1782816000000
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The associated dataset ID.
	//
	// example:
	//
	// arms_customer_agent_level1
	DatasetId *string `json:"datasetId,omitempty" xml:"datasetId,omitempty"`
	// The description.
	//
	// example:
	//
	// 对比 checkout Agent 基线与优化版本
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The list of evaluators.
	//
	// example:
	//
	// [{"evaluatorRef": "Builtin.agent_task_completion"}]
	Evaluators []*Evaluator `json:"evaluators,omitempty" xml:"evaluators,omitempty" type:"Repeated"`
	// The experiment type.
	//
	// example:
	//
	// online
	ExperimentType *string `json:"experimentType,omitempty" xml:"experimentType,omitempty"`
	// The list of experiment configurations.
	//
	// example:
	//
	// [{"label": "A", "name": "baseline", "modelName": "qwen-max"}]
	Experiments []*ExperimentConfig `json:"experiments,omitempty" xml:"experiments,omitempty" type:"Repeated"`
	// Optional.
	//
	// example:
	//
	// {"question": "如何退款？"}
	Input map[string]interface{} `json:"input,omitempty" xml:"input,omitempty"`
	// The experiment plan ID.
	//
	// example:
	//
	// exp-plan-0242d983f5d340fd8479cf2c19eb279e
	PlanId *string `json:"planId,omitempty" xml:"planId,omitempty"`
	// The experiment plan name.
	//
	// example:
	//
	// arms_agent_experiment
	PlanName *string `json:"planName,omitempty" xml:"planName,omitempty"`
	// The custom query SQL clause in partial dataset mode.
	//
	// example:
	//
	// status=\\"OK\\"
	QuerySql *string `json:"querySql,omitempty" xml:"querySql,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3FE4CD1E-FF41-56BE-B590-7A021D9C1524
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The list of selected data item IDs in partial dataset mode.
	//
	// example:
	//
	// ["019ef4d5-a0f0-7114-832d-5542d771cd8c"]
	SelectedItemIds []*string `json:"selectedItemIds,omitempty" xml:"selectedItemIds,omitempty" type:"Repeated"`
	// The plan status.
	//
	// example:
	//
	// stopped
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The update time, in millisecond Unix timestamp.
	//
	// example:
	//
	// 1782816600000
	UpdatedAt *int64 `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s GetExperimentPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentPlanResponseBody) GoString() string {
	return s.String()
}

func (s *GetExperimentPlanResponseBody) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *GetExperimentPlanResponseBody) GetDatasetId() *string {
	return s.DatasetId
}

func (s *GetExperimentPlanResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetExperimentPlanResponseBody) GetEvaluators() []*Evaluator {
	return s.Evaluators
}

func (s *GetExperimentPlanResponseBody) GetExperimentType() *string {
	return s.ExperimentType
}

func (s *GetExperimentPlanResponseBody) GetExperiments() []*ExperimentConfig {
	return s.Experiments
}

func (s *GetExperimentPlanResponseBody) GetInput() map[string]interface{} {
	return s.Input
}

func (s *GetExperimentPlanResponseBody) GetPlanId() *string {
	return s.PlanId
}

func (s *GetExperimentPlanResponseBody) GetPlanName() *string {
	return s.PlanName
}

func (s *GetExperimentPlanResponseBody) GetQuerySql() *string {
	return s.QuerySql
}

func (s *GetExperimentPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExperimentPlanResponseBody) GetSelectedItemIds() []*string {
	return s.SelectedItemIds
}

func (s *GetExperimentPlanResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetExperimentPlanResponseBody) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *GetExperimentPlanResponseBody) SetCreatedAt(v int64) *GetExperimentPlanResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *GetExperimentPlanResponseBody) SetDatasetId(v string) *GetExperimentPlanResponseBody {
	s.DatasetId = &v
	return s
}

func (s *GetExperimentPlanResponseBody) SetDescription(v string) *GetExperimentPlanResponseBody {
	s.Description = &v
	return s
}

func (s *GetExperimentPlanResponseBody) SetEvaluators(v []*Evaluator) *GetExperimentPlanResponseBody {
	s.Evaluators = v
	return s
}

func (s *GetExperimentPlanResponseBody) SetExperimentType(v string) *GetExperimentPlanResponseBody {
	s.ExperimentType = &v
	return s
}

func (s *GetExperimentPlanResponseBody) SetExperiments(v []*ExperimentConfig) *GetExperimentPlanResponseBody {
	s.Experiments = v
	return s
}

func (s *GetExperimentPlanResponseBody) SetInput(v map[string]interface{}) *GetExperimentPlanResponseBody {
	s.Input = v
	return s
}

func (s *GetExperimentPlanResponseBody) SetPlanId(v string) *GetExperimentPlanResponseBody {
	s.PlanId = &v
	return s
}

func (s *GetExperimentPlanResponseBody) SetPlanName(v string) *GetExperimentPlanResponseBody {
	s.PlanName = &v
	return s
}

func (s *GetExperimentPlanResponseBody) SetQuerySql(v string) *GetExperimentPlanResponseBody {
	s.QuerySql = &v
	return s
}

func (s *GetExperimentPlanResponseBody) SetRequestId(v string) *GetExperimentPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExperimentPlanResponseBody) SetSelectedItemIds(v []*string) *GetExperimentPlanResponseBody {
	s.SelectedItemIds = v
	return s
}

func (s *GetExperimentPlanResponseBody) SetStatus(v string) *GetExperimentPlanResponseBody {
	s.Status = &v
	return s
}

func (s *GetExperimentPlanResponseBody) SetUpdatedAt(v int64) *GetExperimentPlanResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *GetExperimentPlanResponseBody) Validate() error {
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
