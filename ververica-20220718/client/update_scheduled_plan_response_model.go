// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduledPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateScheduledPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateScheduledPlanResponse
	GetStatusCode() *int32
	SetBody(v *UpdateScheduledPlanResponseBody) *UpdateScheduledPlanResponse
	GetBody() *UpdateScheduledPlanResponseBody
}

type UpdateScheduledPlanResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateScheduledPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateScheduledPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduledPlanResponse) GoString() string {
	return s.String()
}

func (s *UpdateScheduledPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateScheduledPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateScheduledPlanResponse) GetBody() *UpdateScheduledPlanResponseBody {
	return s.Body
}

func (s *UpdateScheduledPlanResponse) SetHeaders(v map[string]*string) *UpdateScheduledPlanResponse {
	s.Headers = v
	return s
}

func (s *UpdateScheduledPlanResponse) SetStatusCode(v int32) *UpdateScheduledPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateScheduledPlanResponse) SetBody(v *UpdateScheduledPlanResponseBody) *UpdateScheduledPlanResponse {
	s.Body = v
	return s
}

func (s *UpdateScheduledPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
