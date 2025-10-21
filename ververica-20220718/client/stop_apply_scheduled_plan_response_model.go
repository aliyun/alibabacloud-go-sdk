// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopApplyScheduledPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopApplyScheduledPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopApplyScheduledPlanResponse
	GetStatusCode() *int32
	SetBody(v *StopApplyScheduledPlanResponseBody) *StopApplyScheduledPlanResponse
	GetBody() *StopApplyScheduledPlanResponseBody
}

type StopApplyScheduledPlanResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopApplyScheduledPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopApplyScheduledPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s StopApplyScheduledPlanResponse) GoString() string {
	return s.String()
}

func (s *StopApplyScheduledPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopApplyScheduledPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopApplyScheduledPlanResponse) GetBody() *StopApplyScheduledPlanResponseBody {
	return s.Body
}

func (s *StopApplyScheduledPlanResponse) SetHeaders(v map[string]*string) *StopApplyScheduledPlanResponse {
	s.Headers = v
	return s
}

func (s *StopApplyScheduledPlanResponse) SetStatusCode(v int32) *StopApplyScheduledPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *StopApplyScheduledPlanResponse) SetBody(v *StopApplyScheduledPlanResponseBody) *StopApplyScheduledPlanResponse {
	s.Body = v
	return s
}

func (s *StopApplyScheduledPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
