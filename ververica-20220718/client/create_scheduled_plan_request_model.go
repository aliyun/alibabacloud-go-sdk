// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduledPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ScheduledPlan) *CreateScheduledPlanRequest
	GetBody() *ScheduledPlan
}

type CreateScheduledPlanRequest struct {
	Body *ScheduledPlan `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScheduledPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduledPlanRequest) GetBody() *ScheduledPlan {
	return s.Body
}

func (s *CreateScheduledPlanRequest) SetBody(v *ScheduledPlan) *CreateScheduledPlanRequest {
	s.Body = v
	return s
}

func (s *CreateScheduledPlanRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
