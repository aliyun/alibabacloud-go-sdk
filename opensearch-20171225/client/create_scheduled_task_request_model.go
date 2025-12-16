// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduledTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ScheduledTask) *CreateScheduledTaskRequest
	GetBody() *ScheduledTask
}

type CreateScheduledTaskRequest struct {
	// The request body.
	Body *ScheduledTask `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScheduledTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskRequest) GetBody() *ScheduledTask {
	return s.Body
}

func (s *CreateScheduledTaskRequest) SetBody(v *ScheduledTask) *CreateScheduledTaskRequest {
	s.Body = v
	return s
}

func (s *CreateScheduledTaskRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
