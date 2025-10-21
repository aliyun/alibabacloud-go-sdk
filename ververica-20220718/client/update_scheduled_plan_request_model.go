// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduledPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ScheduledPlan) *UpdateScheduledPlanRequest
	GetBody() *ScheduledPlan
}

type UpdateScheduledPlanRequest struct {
	Body *ScheduledPlan `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateScheduledPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduledPlanRequest) GoString() string {
	return s.String()
}

func (s *UpdateScheduledPlanRequest) GetBody() *ScheduledPlan {
	return s.Body
}

func (s *UpdateScheduledPlanRequest) SetBody(v *ScheduledPlan) *UpdateScheduledPlanRequest {
	s.Body = v
	return s
}

func (s *UpdateScheduledPlanRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
