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
	JobPlans   []*JobPlan `json:"JobPlans,omitempty" xml:"JobPlans,omitempty" type:"Repeated"`
	RequestId  *string    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
