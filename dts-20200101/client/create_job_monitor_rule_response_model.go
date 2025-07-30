// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobMonitorRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateJobMonitorRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateJobMonitorRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateJobMonitorRuleResponseBody) *CreateJobMonitorRuleResponse
	GetBody() *CreateJobMonitorRuleResponseBody
}

type CreateJobMonitorRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateJobMonitorRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateJobMonitorRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateJobMonitorRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateJobMonitorRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateJobMonitorRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateJobMonitorRuleResponse) GetBody() *CreateJobMonitorRuleResponseBody {
	return s.Body
}

func (s *CreateJobMonitorRuleResponse) SetHeaders(v map[string]*string) *CreateJobMonitorRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateJobMonitorRuleResponse) SetStatusCode(v int32) *CreateJobMonitorRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateJobMonitorRuleResponse) SetBody(v *CreateJobMonitorRuleResponseBody) *CreateJobMonitorRuleResponse {
	s.Body = v
	return s
}

func (s *CreateJobMonitorRuleResponse) Validate() error {
	return dara.Validate(s)
}
