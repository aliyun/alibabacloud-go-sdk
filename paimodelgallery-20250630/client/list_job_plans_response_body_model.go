// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobPlansResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobPlans(v []*JobPlan) *ListJobPlansResponseBody
	GetJobPlans() []*JobPlan
	SetRequestId(v string) *ListJobPlansResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListJobPlansResponseBody
	GetTotalCount() *int32
}

type ListJobPlansResponseBody struct {
	// The list of job plans.
	JobPlans []*JobPlan `json:"JobPlans,omitempty" xml:"JobPlans,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82-9624-EC2B1779848E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListJobPlansResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobPlansResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobPlansResponseBody) GetJobPlans() []*JobPlan {
	return s.JobPlans
}

func (s *ListJobPlansResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJobPlansResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListJobPlansResponseBody) SetJobPlans(v []*JobPlan) *ListJobPlansResponseBody {
	s.JobPlans = v
	return s
}

func (s *ListJobPlansResponseBody) SetRequestId(v string) *ListJobPlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobPlansResponseBody) SetTotalCount(v int32) *ListJobPlansResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListJobPlansResponseBody) Validate() error {
	if s.JobPlans != nil {
		for _, item := range s.JobPlans {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
