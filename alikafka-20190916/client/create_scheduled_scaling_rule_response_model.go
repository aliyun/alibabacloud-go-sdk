// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduledScalingRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateScheduledScalingRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateScheduledScalingRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateScheduledScalingRuleResponseBody) *CreateScheduledScalingRuleResponse
	GetBody() *CreateScheduledScalingRuleResponseBody
}

type CreateScheduledScalingRuleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScheduledScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScheduledScalingRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateScheduledScalingRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateScheduledScalingRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateScheduledScalingRuleResponse) GetBody() *CreateScheduledScalingRuleResponseBody {
	return s.Body
}

func (s *CreateScheduledScalingRuleResponse) SetHeaders(v map[string]*string) *CreateScheduledScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateScheduledScalingRuleResponse) SetStatusCode(v int32) *CreateScheduledScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScheduledScalingRuleResponse) SetBody(v *CreateScheduledScalingRuleResponseBody) *CreateScheduledScalingRuleResponse {
	s.Body = v
	return s
}

func (s *CreateScheduledScalingRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
