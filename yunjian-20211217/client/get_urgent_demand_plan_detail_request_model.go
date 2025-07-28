// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUrgentDemandPlanDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPlanId(v string) *GetUrgentDemandPlanDetailRequest
	GetPlanId() *string
}

type GetUrgentDemandPlanDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 111223
	PlanId *string `json:"planId,omitempty" xml:"planId,omitempty"`
}

func (s GetUrgentDemandPlanDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUrgentDemandPlanDetailRequest) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandPlanDetailRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *GetUrgentDemandPlanDetailRequest) SetPlanId(v string) *GetUrgentDemandPlanDetailRequest {
	s.PlanId = &v
	return s
}

func (s *GetUrgentDemandPlanDetailRequest) Validate() error {
	return dara.Validate(s)
}
