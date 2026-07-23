// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExperimentPlansResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListExperimentPlansResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListExperimentPlansResponseBody
	GetNextToken() *string
	SetPage(v int32) *ListExperimentPlansResponseBody
	GetPage() *int32
	SetPageSize(v int32) *ListExperimentPlansResponseBody
	GetPageSize() *int32
	SetPlans(v []*ExperimentPlanData) *ListExperimentPlansResponseBody
	GetPlans() []*ExperimentPlanData
	SetRequestId(v string) *ListExperimentPlansResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListExperimentPlansResponseBody
	GetTotal() *int32
}

type ListExperimentPlansResponseBody struct {
	// The maximum number of entries returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token for the next page.
	//
	// example:
	//
	// eyJsYXN0SWQiOjEwMX0=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 0
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The list of experiment plan summaries.
	//
	// example:
	//
	// [{"planId": "exp-plan-e95bff54685a4ae29ff3a834c1008a71", "planName": "rca_benchmark_eval_experiment", "experimentType": "offline", "description": "", "status": "pending", "datasetId": "rca_benckmark_eval", "querySql": "", "experimentCount": 5, "createdAt": 1784612365000, "updatedAt": 1784619562000}, {"planId": "exp-plan-0242d983f5d340fd8479cf2c19eb279e", "planName": "arms_agent_experiment", "experimentType": "online", "description": "", "status": "stopped", "datasetId": "arms_customer_agent_level1", "querySql": "", "experimentCount": 4, "createdAt": 1784257858000, "updatedAt": 1784721811000}, {"planId": "b7f0ad3d-3765-446a-a744-ab64ab8bf386", "planName": "arms_customer_agent_plan", "experimentType": "offline", "description": "", "status": "stopped", "datasetId": "arms_customer_agent_level1", "querySql": "where \\"input\\" LIKE \\"%探针%\\"", "experimentCount": 65, "createdAt": 1782310430000, "updatedAt": 1784692254000}]
	Plans []*ExperimentPlanData `json:"plans,omitempty" xml:"plans,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 3FE4CD1E-FF41-56BE-B590-7A021D9C1524
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of records that match the filter criteria.
	//
	// example:
	//
	// 6
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListExperimentPlansResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExperimentPlansResponseBody) GoString() string {
	return s.String()
}

func (s *ListExperimentPlansResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListExperimentPlansResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListExperimentPlansResponseBody) GetPage() *int32 {
	return s.Page
}

func (s *ListExperimentPlansResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListExperimentPlansResponseBody) GetPlans() []*ExperimentPlanData {
	return s.Plans
}

func (s *ListExperimentPlansResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExperimentPlansResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListExperimentPlansResponseBody) SetMaxResults(v int32) *ListExperimentPlansResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListExperimentPlansResponseBody) SetNextToken(v string) *ListExperimentPlansResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListExperimentPlansResponseBody) SetPage(v int32) *ListExperimentPlansResponseBody {
	s.Page = &v
	return s
}

func (s *ListExperimentPlansResponseBody) SetPageSize(v int32) *ListExperimentPlansResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListExperimentPlansResponseBody) SetPlans(v []*ExperimentPlanData) *ListExperimentPlansResponseBody {
	s.Plans = v
	return s
}

func (s *ListExperimentPlansResponseBody) SetRequestId(v string) *ListExperimentPlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExperimentPlansResponseBody) SetTotal(v int32) *ListExperimentPlansResponseBody {
	s.Total = &v
	return s
}

func (s *ListExperimentPlansResponseBody) Validate() error {
	if s.Plans != nil {
		for _, item := range s.Plans {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
