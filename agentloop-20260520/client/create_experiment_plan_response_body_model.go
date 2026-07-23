// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExperimentPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *CreateExperimentPlanResponseBody
	GetMessage() *string
	SetPlanId(v string) *CreateExperimentPlanResponseBody
	GetPlanId() *string
	SetRequestId(v string) *CreateExperimentPlanResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateExperimentPlanResponseBody
	GetStatus() *string
}

type CreateExperimentPlanResponseBody struct {
	// The message.
	//
	// example:
	//
	// Experiment plan created successfully.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The experiment plan ID.
	//
	// example:
	//
	// exp-plan-e95bff54685a4ae29ff3a834c1008a71
	PlanId *string `json:"planId,omitempty" xml:"planId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3FE4CD1E-FF41-56BE-B590-7A021D9C1524
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The creation result. The value is `created` if the operation is successful.
	//
	// example:
	//
	// created
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateExperimentPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateExperimentPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExperimentPlanResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateExperimentPlanResponseBody) GetPlanId() *string {
	return s.PlanId
}

func (s *CreateExperimentPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateExperimentPlanResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateExperimentPlanResponseBody) SetMessage(v string) *CreateExperimentPlanResponseBody {
	s.Message = &v
	return s
}

func (s *CreateExperimentPlanResponseBody) SetPlanId(v string) *CreateExperimentPlanResponseBody {
	s.PlanId = &v
	return s
}

func (s *CreateExperimentPlanResponseBody) SetRequestId(v string) *CreateExperimentPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExperimentPlanResponseBody) SetStatus(v string) *CreateExperimentPlanResponseBody {
	s.Status = &v
	return s
}

func (s *CreateExperimentPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
