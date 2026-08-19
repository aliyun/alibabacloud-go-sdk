// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExperimentPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *UpdateExperimentPlanResponseBody
	GetMessage() *string
	SetPlanId(v string) *UpdateExperimentPlanResponseBody
	GetPlanId() *string
	SetRequestId(v string) *UpdateExperimentPlanResponseBody
	GetRequestId() *string
	SetStatus(v string) *UpdateExperimentPlanResponseBody
	GetStatus() *string
}

type UpdateExperimentPlanResponseBody struct {
	// The prompt message.
	//
	// example:
	//
	// Experiment plan updated successfully.
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
	// The update result. The value is updated if the operation is successful.
	//
	// example:
	//
	// updated
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateExperimentPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateExperimentPlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExperimentPlanResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateExperimentPlanResponseBody) GetPlanId() *string {
	return s.PlanId
}

func (s *UpdateExperimentPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateExperimentPlanResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateExperimentPlanResponseBody) SetMessage(v string) *UpdateExperimentPlanResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateExperimentPlanResponseBody) SetPlanId(v string) *UpdateExperimentPlanResponseBody {
	s.PlanId = &v
	return s
}

func (s *UpdateExperimentPlanResponseBody) SetRequestId(v string) *UpdateExperimentPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateExperimentPlanResponseBody) SetStatus(v string) *UpdateExperimentPlanResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateExperimentPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
