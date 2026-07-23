// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExperimentRunsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListExperimentRunsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListExperimentRunsResponseBody
	GetNextToken() *string
	SetPage(v int32) *ListExperimentRunsResponseBody
	GetPage() *int32
	SetPageSize(v int32) *ListExperimentRunsResponseBody
	GetPageSize() *int32
	SetRecords(v []*ExperimentRecord) *ListExperimentRunsResponseBody
	GetRecords() []*ExperimentRecord
	SetRequestId(v string) *ListExperimentRunsResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListExperimentRunsResponseBody
	GetTotal() *int32
}

type ListExperimentRunsResponseBody struct {
	// The maximum number of entries returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// Optional.
	//
	// example:
	//
	// eyJwYWdlIjoxfQ==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 0
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The list of experiment run records.
	//
	// example:
	//
	// [{"recordId": "exp-run-f6d419b0ed3d43a7b585948a55efc07b", "experimentPlanId": "exp-plan-0242d983f5d340fd8479cf2c19eb279e", "recordName": "arms_agent_experiment 2026/07/22 20:02:55", "planName": "arms_agent_experiment", "status": "evaluating", "progress": 100.0, "totalTasks": 40, "completedTasks": 40, "failedTasks": 0, "dataSourceType": "dataset-full", "datasetId": "arms_customer_agent_level1", "modelNames": ["qwen3.7-plus", "qwen3.7-max"], "evaluationTaskId": "eval-task-6bec93bfa03740dd86ce2bf1496e65fb", "executedAt": 1784721775379, "completedAt": 1784721811392}, {"recordId": "a5397261-6e6d-4e45-bf52-feb8686f7524", "experimentPlanId": "exp-plan-e95bff54685a4ae29ff3a834c1008a71", "recordName": "rca_benchmark_eval_experiment 2026/07/22 19:23:59", "planName": "rca_benchmark_eval_experiment", "status": "completed", "progress": 100.0, "totalTasks": 20, "completedTasks": 20, "failedTasks": 0, "dataSourceType": "dataset-full", "datasetId": "rca_benckmark_eval", "modelNames": [], "evaluationTaskId": "eval-task-b1395b3bdf3e4dc994d7dcde7a66da45", "executedAt": 1784719439255, "completedAt": 1784719989371}]
	Records []*ExperimentRecord `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 3FE4CD1E-FF41-56BE-B590-7A021D9C1524
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 100
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListExperimentRunsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExperimentRunsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExperimentRunsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListExperimentRunsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListExperimentRunsResponseBody) GetPage() *int32 {
	return s.Page
}

func (s *ListExperimentRunsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListExperimentRunsResponseBody) GetRecords() []*ExperimentRecord {
	return s.Records
}

func (s *ListExperimentRunsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExperimentRunsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListExperimentRunsResponseBody) SetMaxResults(v int32) *ListExperimentRunsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListExperimentRunsResponseBody) SetNextToken(v string) *ListExperimentRunsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListExperimentRunsResponseBody) SetPage(v int32) *ListExperimentRunsResponseBody {
	s.Page = &v
	return s
}

func (s *ListExperimentRunsResponseBody) SetPageSize(v int32) *ListExperimentRunsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListExperimentRunsResponseBody) SetRecords(v []*ExperimentRecord) *ListExperimentRunsResponseBody {
	s.Records = v
	return s
}

func (s *ListExperimentRunsResponseBody) SetRequestId(v string) *ListExperimentRunsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExperimentRunsResponseBody) SetTotal(v int32) *ListExperimentRunsResponseBody {
	s.Total = &v
	return s
}

func (s *ListExperimentRunsResponseBody) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
