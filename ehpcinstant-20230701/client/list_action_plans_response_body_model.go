// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListActionPlansResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActionPlans(v []*ListActionPlansResponseBodyActionPlans) *ListActionPlansResponseBody
	GetActionPlans() []*ListActionPlansResponseBodyActionPlans
	SetMaxResults(v int32) *ListActionPlansResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListActionPlansResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListActionPlansResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListActionPlansResponseBody
	GetTotalCount() *int32
}

type ListActionPlansResponseBody struct {
	// The list of execution plan results.
	ActionPlans []*ListActionPlansResponseBodyActionPlans `json:"ActionPlans,omitempty" xml:"ActionPlans,omitempty" type:"Repeated"`
	// The maximum number of records returned in this request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Indicates the read position returned by the current call. An empty value means all data has been read.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1d2db86scXXXXXXXXXX
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total data count under the current request conditions (optional; not returned by default).
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListActionPlansResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListActionPlansResponseBody) GoString() string {
	return s.String()
}

func (s *ListActionPlansResponseBody) GetActionPlans() []*ListActionPlansResponseBodyActionPlans {
	return s.ActionPlans
}

func (s *ListActionPlansResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListActionPlansResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListActionPlansResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListActionPlansResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListActionPlansResponseBody) SetActionPlans(v []*ListActionPlansResponseBodyActionPlans) *ListActionPlansResponseBody {
	s.ActionPlans = v
	return s
}

func (s *ListActionPlansResponseBody) SetMaxResults(v int32) *ListActionPlansResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListActionPlansResponseBody) SetNextToken(v string) *ListActionPlansResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListActionPlansResponseBody) SetRequestId(v string) *ListActionPlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListActionPlansResponseBody) SetTotalCount(v int32) *ListActionPlansResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListActionPlansResponseBody) Validate() error {
	if s.ActionPlans != nil {
		for _, item := range s.ActionPlans {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListActionPlansResponseBodyActionPlans struct {
	// The ID of the execution plan.
	//
	// example:
	//
	// ap-hz036ubmx2qmw93k****
	ActionPlanId *string `json:"ActionPlanId,omitempty" xml:"ActionPlanId,omitempty"`
	// The name of the execution plan.
	//
	// example:
	//
	// TestActionPlan
	ActionPlanName *string `json:"ActionPlanName,omitempty" xml:"ActionPlanName,omitempty"`
	// The time when the execution plan was created.
	//
	// example:
	//
	// 2025-08-10 17:58:24
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The status of the execution plan. The possible values are as follows:
	//
	// 	- Active Instant tasks are dynamically managed only when the execution plan is in the Active state.
	//
	// 	- Inactive Instant tasks are no longer managed by execution plans in the Inactive state.
	//
	// 	- Deleting: The execution plan is being deleted. You cannot modify the parameters of an execution plan in this state.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the execution plan was last modified. The time follows the ISO 8601 standard and UTC +0. The format is yyyy-MM-ddTHH:mmZ.
	//
	// example:
	//
	// 2025-08-10 17:58:24
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListActionPlansResponseBodyActionPlans) String() string {
	return dara.Prettify(s)
}

func (s ListActionPlansResponseBodyActionPlans) GoString() string {
	return s.String()
}

func (s *ListActionPlansResponseBodyActionPlans) GetActionPlanId() *string {
	return s.ActionPlanId
}

func (s *ListActionPlansResponseBodyActionPlans) GetActionPlanName() *string {
	return s.ActionPlanName
}

func (s *ListActionPlansResponseBodyActionPlans) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListActionPlansResponseBodyActionPlans) GetStatus() *string {
	return s.Status
}

func (s *ListActionPlansResponseBodyActionPlans) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListActionPlansResponseBodyActionPlans) SetActionPlanId(v string) *ListActionPlansResponseBodyActionPlans {
	s.ActionPlanId = &v
	return s
}

func (s *ListActionPlansResponseBodyActionPlans) SetActionPlanName(v string) *ListActionPlansResponseBodyActionPlans {
	s.ActionPlanName = &v
	return s
}

func (s *ListActionPlansResponseBodyActionPlans) SetCreateTime(v string) *ListActionPlansResponseBodyActionPlans {
	s.CreateTime = &v
	return s
}

func (s *ListActionPlansResponseBodyActionPlans) SetStatus(v string) *ListActionPlansResponseBodyActionPlans {
	s.Status = &v
	return s
}

func (s *ListActionPlansResponseBodyActionPlans) SetUpdateTime(v string) *ListActionPlansResponseBodyActionPlans {
	s.UpdateTime = &v
	return s
}

func (s *ListActionPlansResponseBodyActionPlans) Validate() error {
	return dara.Validate(s)
}
