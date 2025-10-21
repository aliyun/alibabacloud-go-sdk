// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduledPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateScheduledPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateScheduledPlanResponse
	GetStatusCode() *int32
	SetBody(v *CreateScheduledPlanResponseBody) *CreateScheduledPlanResponse
	GetBody() *CreateScheduledPlanResponseBody
}

type CreateScheduledPlanResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScheduledPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScheduledPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateScheduledPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateScheduledPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateScheduledPlanResponse) GetBody() *CreateScheduledPlanResponseBody {
	return s.Body
}

func (s *CreateScheduledPlanResponse) SetHeaders(v map[string]*string) *CreateScheduledPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateScheduledPlanResponse) SetStatusCode(v int32) *CreateScheduledPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScheduledPlanResponse) SetBody(v *CreateScheduledPlanResponseBody) *CreateScheduledPlanResponse {
	s.Body = v
	return s
}

func (s *CreateScheduledPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
