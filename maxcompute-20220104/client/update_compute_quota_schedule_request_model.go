// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeQuotaScheduleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v []*UpdateComputeQuotaScheduleRequestBody) *UpdateComputeQuotaScheduleRequest
	GetBody() []*UpdateComputeQuotaScheduleRequestBody
	SetScheduleTimezone(v string) *UpdateComputeQuotaScheduleRequest
	GetScheduleTimezone() *string
}

type UpdateComputeQuotaScheduleRequest struct {
	// The request body parameters.
	Body             []*UpdateComputeQuotaScheduleRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	ScheduleTimezone *string                                  `json:"scheduleTimezone,omitempty" xml:"scheduleTimezone,omitempty"`
}

func (s UpdateComputeQuotaScheduleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeQuotaScheduleRequest) GoString() string {
	return s.String()
}

func (s *UpdateComputeQuotaScheduleRequest) GetBody() []*UpdateComputeQuotaScheduleRequestBody {
	return s.Body
}

func (s *UpdateComputeQuotaScheduleRequest) GetScheduleTimezone() *string {
	return s.ScheduleTimezone
}

func (s *UpdateComputeQuotaScheduleRequest) SetBody(v []*UpdateComputeQuotaScheduleRequestBody) *UpdateComputeQuotaScheduleRequest {
	s.Body = v
	return s
}

func (s *UpdateComputeQuotaScheduleRequest) SetScheduleTimezone(v string) *UpdateComputeQuotaScheduleRequest {
	s.ScheduleTimezone = &v
	return s
}

func (s *UpdateComputeQuotaScheduleRequest) Validate() error {
	if s.Body != nil {
		for _, item := range s.Body {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateComputeQuotaScheduleRequestBody struct {
	// The value of effective condition.
	Condition *UpdateComputeQuotaScheduleRequestBodyCondition `json:"condition,omitempty" xml:"condition,omitempty" type:"Struct"`
	// The name of the quota plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// planA
	Plan *string `json:"plan,omitempty" xml:"plan,omitempty"`
	// The type of the quota plan.
	//
	// 	Notice: Currently, only daily is supported.</notice>
	//
	// This parameter is required.
	//
	// example:
	//
	// daily
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateComputeQuotaScheduleRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeQuotaScheduleRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateComputeQuotaScheduleRequestBody) GetCondition() *UpdateComputeQuotaScheduleRequestBodyCondition {
	return s.Condition
}

func (s *UpdateComputeQuotaScheduleRequestBody) GetPlan() *string {
	return s.Plan
}

func (s *UpdateComputeQuotaScheduleRequestBody) GetType() *string {
	return s.Type
}

func (s *UpdateComputeQuotaScheduleRequestBody) SetCondition(v *UpdateComputeQuotaScheduleRequestBodyCondition) *UpdateComputeQuotaScheduleRequestBody {
	s.Condition = v
	return s
}

func (s *UpdateComputeQuotaScheduleRequestBody) SetPlan(v string) *UpdateComputeQuotaScheduleRequestBody {
	s.Plan = &v
	return s
}

func (s *UpdateComputeQuotaScheduleRequestBody) SetType(v string) *UpdateComputeQuotaScheduleRequestBody {
	s.Type = &v
	return s
}

func (s *UpdateComputeQuotaScheduleRequestBody) Validate() error {
	if s.Condition != nil {
		if err := s.Condition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateComputeQuotaScheduleRequestBodyCondition struct {
	// The start time when the quota plan takes effect.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10:00
	At *string `json:"at,omitempty" xml:"at,omitempty"`
}

func (s UpdateComputeQuotaScheduleRequestBodyCondition) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeQuotaScheduleRequestBodyCondition) GoString() string {
	return s.String()
}

func (s *UpdateComputeQuotaScheduleRequestBodyCondition) GetAt() *string {
	return s.At
}

func (s *UpdateComputeQuotaScheduleRequestBodyCondition) SetAt(v string) *UpdateComputeQuotaScheduleRequestBodyCondition {
	s.At = &v
	return s
}

func (s *UpdateComputeQuotaScheduleRequestBodyCondition) Validate() error {
	return dara.Validate(s)
}
