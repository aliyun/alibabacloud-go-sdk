// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationPlansForRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPlans(v []*ListOperationPlansForRegionResponseBodyPlans) *ListOperationPlansForRegionResponseBody
	GetPlans() []*ListOperationPlansForRegionResponseBodyPlans
}

type ListOperationPlansForRegionResponseBody struct {
	Plans []*ListOperationPlansForRegionResponseBodyPlans `json:"plans,omitempty" xml:"plans,omitempty" type:"Repeated"`
}

func (s ListOperationPlansForRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOperationPlansForRegionResponseBody) GoString() string {
	return s.String()
}

func (s *ListOperationPlansForRegionResponseBody) GetPlans() []*ListOperationPlansForRegionResponseBodyPlans {
	return s.Plans
}

func (s *ListOperationPlansForRegionResponseBody) SetPlans(v []*ListOperationPlansForRegionResponseBodyPlans) *ListOperationPlansForRegionResponseBody {
	s.Plans = v
	return s
}

func (s *ListOperationPlansForRegionResponseBody) Validate() error {
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

type ListOperationPlansForRegionResponseBodyPlans struct {
	// example:
	//
	// c29ced64b3dfe4f33b57ca0aa9f68****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// example:
	//
	// 2023-11-21T20:01:22+08:00
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// example:
	//
	// 2023-11-22T18:00:00+08:00
	EndTime *string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// example:
	//
	// P-655c9c127e0e6603ef00****
	PlanId *string `json:"plan_id,omitempty" xml:"plan_id,omitempty"`
	// example:
	//
	// 2023-11-22T15:18:00+08:00
	StartTime *string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// example:
	//
	// Scheduled
	State       *string                                                  `json:"state,omitempty" xml:"state,omitempty"`
	StateReason *ListOperationPlansForRegionResponseBodyPlansStateReason `json:"state_reason,omitempty" xml:"state_reason,omitempty" type:"Struct"`
	// example:
	//
	// c29ced64b3dfe4f33b57ca0aa9f68****
	TargetId *string `json:"target_id,omitempty" xml:"target_id,omitempty"`
	// example:
	//
	// cluster
	TargetType *string `json:"target_type,omitempty" xml:"target_type,omitempty"`
	// example:
	//
	// T-681ac448b23ced010600****
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
	// example:
	//
	// CLUSTER_UPGRADE_MASTER
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListOperationPlansForRegionResponseBodyPlans) String() string {
	return dara.Prettify(s)
}

func (s ListOperationPlansForRegionResponseBodyPlans) GoString() string {
	return s.String()
}

func (s *ListOperationPlansForRegionResponseBodyPlans) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListOperationPlansForRegionResponseBodyPlans) GetCreated() *string {
	return s.Created
}

func (s *ListOperationPlansForRegionResponseBodyPlans) GetEndTime() *string {
	return s.EndTime
}

func (s *ListOperationPlansForRegionResponseBodyPlans) GetPlanId() *string {
	return s.PlanId
}

func (s *ListOperationPlansForRegionResponseBodyPlans) GetStartTime() *string {
	return s.StartTime
}

func (s *ListOperationPlansForRegionResponseBodyPlans) GetState() *string {
	return s.State
}

func (s *ListOperationPlansForRegionResponseBodyPlans) GetStateReason() *ListOperationPlansForRegionResponseBodyPlansStateReason {
	return s.StateReason
}

func (s *ListOperationPlansForRegionResponseBodyPlans) GetTargetId() *string {
	return s.TargetId
}

func (s *ListOperationPlansForRegionResponseBodyPlans) GetTargetType() *string {
	return s.TargetType
}

func (s *ListOperationPlansForRegionResponseBodyPlans) GetTaskId() *string {
	return s.TaskId
}

func (s *ListOperationPlansForRegionResponseBodyPlans) GetType() *string {
	return s.Type
}

func (s *ListOperationPlansForRegionResponseBodyPlans) SetClusterId(v string) *ListOperationPlansForRegionResponseBodyPlans {
	s.ClusterId = &v
	return s
}

func (s *ListOperationPlansForRegionResponseBodyPlans) SetCreated(v string) *ListOperationPlansForRegionResponseBodyPlans {
	s.Created = &v
	return s
}

func (s *ListOperationPlansForRegionResponseBodyPlans) SetEndTime(v string) *ListOperationPlansForRegionResponseBodyPlans {
	s.EndTime = &v
	return s
}

func (s *ListOperationPlansForRegionResponseBodyPlans) SetPlanId(v string) *ListOperationPlansForRegionResponseBodyPlans {
	s.PlanId = &v
	return s
}

func (s *ListOperationPlansForRegionResponseBodyPlans) SetStartTime(v string) *ListOperationPlansForRegionResponseBodyPlans {
	s.StartTime = &v
	return s
}

func (s *ListOperationPlansForRegionResponseBodyPlans) SetState(v string) *ListOperationPlansForRegionResponseBodyPlans {
	s.State = &v
	return s
}

func (s *ListOperationPlansForRegionResponseBodyPlans) SetStateReason(v *ListOperationPlansForRegionResponseBodyPlansStateReason) *ListOperationPlansForRegionResponseBodyPlans {
	s.StateReason = v
	return s
}

func (s *ListOperationPlansForRegionResponseBodyPlans) SetTargetId(v string) *ListOperationPlansForRegionResponseBodyPlans {
	s.TargetId = &v
	return s
}

func (s *ListOperationPlansForRegionResponseBodyPlans) SetTargetType(v string) *ListOperationPlansForRegionResponseBodyPlans {
	s.TargetType = &v
	return s
}

func (s *ListOperationPlansForRegionResponseBodyPlans) SetTaskId(v string) *ListOperationPlansForRegionResponseBodyPlans {
	s.TaskId = &v
	return s
}

func (s *ListOperationPlansForRegionResponseBodyPlans) SetType(v string) *ListOperationPlansForRegionResponseBodyPlans {
	s.Type = &v
	return s
}

func (s *ListOperationPlansForRegionResponseBodyPlans) Validate() error {
	if s.StateReason != nil {
		if err := s.StateReason.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListOperationPlansForRegionResponseBodyPlansStateReason struct {
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ListOperationPlansForRegionResponseBodyPlansStateReason) String() string {
	return dara.Prettify(s)
}

func (s ListOperationPlansForRegionResponseBodyPlansStateReason) GoString() string {
	return s.String()
}

func (s *ListOperationPlansForRegionResponseBodyPlansStateReason) GetCode() *string {
	return s.Code
}

func (s *ListOperationPlansForRegionResponseBodyPlansStateReason) GetMessage() *string {
	return s.Message
}

func (s *ListOperationPlansForRegionResponseBodyPlansStateReason) SetCode(v string) *ListOperationPlansForRegionResponseBodyPlansStateReason {
	s.Code = &v
	return s
}

func (s *ListOperationPlansForRegionResponseBodyPlansStateReason) SetMessage(v string) *ListOperationPlansForRegionResponseBodyPlansStateReason {
	s.Message = &v
	return s
}

func (s *ListOperationPlansForRegionResponseBodyPlansStateReason) Validate() error {
	return dara.Validate(s)
}
