// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExperimentPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPlanId(v string) *DeleteExperimentPlanResponseBody
	GetPlanId() *string
	SetRequestId(v string) *DeleteExperimentPlanResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteExperimentPlanResponseBody
	GetStatus() *string
}

type DeleteExperimentPlanResponseBody struct {
	// The experiment plan ID.
	//
	// example:
	//
	// exp-plan-aa1a66b074bc42aa8696c73c7dc9b718
	PlanId *string `json:"planId,omitempty" xml:"planId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3FE4CD1E-FF41-56BE-B590-7A021D9C1524
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The deletion result. The value is deleted if the operation is successful.
	//
	// example:
	//
	// deleted
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DeleteExperimentPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteExperimentPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExperimentPlanResponseBody) GetPlanId() *string {
	return s.PlanId
}

func (s *DeleteExperimentPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteExperimentPlanResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteExperimentPlanResponseBody) SetPlanId(v string) *DeleteExperimentPlanResponseBody {
	s.PlanId = &v
	return s
}

func (s *DeleteExperimentPlanResponseBody) SetRequestId(v string) *DeleteExperimentPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExperimentPlanResponseBody) SetStatus(v string) *DeleteExperimentPlanResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteExperimentPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
