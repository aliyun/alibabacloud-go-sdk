// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyScheduledPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyScheduledPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyScheduledPlanResponse
	GetStatusCode() *int32
	SetBody(v *ApplyScheduledPlanResponseBody) *ApplyScheduledPlanResponse
	GetBody() *ApplyScheduledPlanResponseBody
}

type ApplyScheduledPlanResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyScheduledPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyScheduledPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyScheduledPlanResponse) GoString() string {
	return s.String()
}

func (s *ApplyScheduledPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyScheduledPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyScheduledPlanResponse) GetBody() *ApplyScheduledPlanResponseBody {
	return s.Body
}

func (s *ApplyScheduledPlanResponse) SetHeaders(v map[string]*string) *ApplyScheduledPlanResponse {
	s.Headers = v
	return s
}

func (s *ApplyScheduledPlanResponse) SetStatusCode(v int32) *ApplyScheduledPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyScheduledPlanResponse) SetBody(v *ApplyScheduledPlanResponseBody) *ApplyScheduledPlanResponse {
	s.Body = v
	return s
}

func (s *ApplyScheduledPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
