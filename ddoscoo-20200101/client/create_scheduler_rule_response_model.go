// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSchedulerRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSchedulerRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSchedulerRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateSchedulerRuleResponseBody) *CreateSchedulerRuleResponse
	GetBody() *CreateSchedulerRuleResponseBody
}

type CreateSchedulerRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSchedulerRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSchedulerRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSchedulerRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateSchedulerRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSchedulerRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSchedulerRuleResponse) GetBody() *CreateSchedulerRuleResponseBody {
	return s.Body
}

func (s *CreateSchedulerRuleResponse) SetHeaders(v map[string]*string) *CreateSchedulerRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateSchedulerRuleResponse) SetStatusCode(v int32) *CreateSchedulerRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSchedulerRuleResponse) SetBody(v *CreateSchedulerRuleResponseBody) *CreateSchedulerRuleResponse {
	s.Body = v
	return s
}

func (s *CreateSchedulerRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
