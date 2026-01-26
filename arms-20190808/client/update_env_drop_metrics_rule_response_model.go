// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnvDropMetricsRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEnvDropMetricsRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEnvDropMetricsRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEnvDropMetricsRuleResponseBody) *UpdateEnvDropMetricsRuleResponse
	GetBody() *UpdateEnvDropMetricsRuleResponseBody
}

type UpdateEnvDropMetricsRuleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEnvDropMetricsRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEnvDropMetricsRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnvDropMetricsRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateEnvDropMetricsRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEnvDropMetricsRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEnvDropMetricsRuleResponse) GetBody() *UpdateEnvDropMetricsRuleResponseBody {
	return s.Body
}

func (s *UpdateEnvDropMetricsRuleResponse) SetHeaders(v map[string]*string) *UpdateEnvDropMetricsRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateEnvDropMetricsRuleResponse) SetStatusCode(v int32) *UpdateEnvDropMetricsRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEnvDropMetricsRuleResponse) SetBody(v *UpdateEnvDropMetricsRuleResponseBody) *UpdateEnvDropMetricsRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateEnvDropMetricsRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
