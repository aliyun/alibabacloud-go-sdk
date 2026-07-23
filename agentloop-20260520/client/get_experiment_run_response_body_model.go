// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperimentRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecord(v *ExperimentRecord) *GetExperimentRunResponseBody
	GetRecord() *ExperimentRecord
	SetRegionId(v string) *GetExperimentRunResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetExperimentRunResponseBody
	GetRequestId() *string
}

type GetExperimentRunResponseBody struct {
	// The experiment run record details. Fields with null values are not returned.
	//
	// example:
	//
	// {"recordId": "exp-run-f6d419b0ed3d43a7b585948a55efc07b", "experimentPlanId": "exp-plan-0242d983f5d340fd8479cf2c19eb279e", "recordName": "arms_agent_experiment 2026/07/22 20:02:55", "planName": "arms_agent_experiment", "status": "evaluating", "totalTasks": 40, "completedTasks": 40, "failedTasks": 0, "progress": 100.0, "executedAt": 1784721775379, "completedAt": 1784721811392, "dataSourceType": "dataset-full", "datasetId": "arms_customer_agent_level1", "modelNames": ["qwen3.7-plus", "qwen3.7-max"], "evaluationTaskId": "eval-task-6bec93bfa03740dd86ce2bf1496e65fb"}
	Record *ExperimentRecord `json:"record,omitempty" xml:"record,omitempty"`
	// The region ID.
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
}

func (s GetExperimentRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentRunResponseBody) GoString() string {
	return s.String()
}

func (s *GetExperimentRunResponseBody) GetRecord() *ExperimentRecord {
	return s.Record
}

func (s *GetExperimentRunResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetExperimentRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExperimentRunResponseBody) SetRecord(v *ExperimentRecord) *GetExperimentRunResponseBody {
	s.Record = v
	return s
}

func (s *GetExperimentRunResponseBody) SetRegionId(v string) *GetExperimentRunResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetExperimentRunResponseBody) SetRequestId(v string) *GetExperimentRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExperimentRunResponseBody) Validate() error {
	if s.Record != nil {
		if err := s.Record.Validate(); err != nil {
			return err
		}
	}
	return nil
}
