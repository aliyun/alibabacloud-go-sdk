// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitUrgentDemandPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPlanId(v string) *SubmitUrgentDemandPlanRequest
	GetPlanId() *string
	SetUserId(v string) *SubmitUrgentDemandPlanRequest
	GetUserId() *string
}

type SubmitUrgentDemandPlanRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 111223
	PlanId *string `json:"planId,omitempty" xml:"planId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 262940
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SubmitUrgentDemandPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitUrgentDemandPlanRequest) GoString() string {
	return s.String()
}

func (s *SubmitUrgentDemandPlanRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *SubmitUrgentDemandPlanRequest) GetUserId() *string {
	return s.UserId
}

func (s *SubmitUrgentDemandPlanRequest) SetPlanId(v string) *SubmitUrgentDemandPlanRequest {
	s.PlanId = &v
	return s
}

func (s *SubmitUrgentDemandPlanRequest) SetUserId(v string) *SubmitUrgentDemandPlanRequest {
	s.UserId = &v
	return s
}

func (s *SubmitUrgentDemandPlanRequest) Validate() error {
	return dara.Validate(s)
}
