// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrometheusAlertRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePrometheusAlertRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePrometheusAlertRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreatePrometheusAlertRuleResponseBody) *CreatePrometheusAlertRuleResponse
	GetBody() *CreatePrometheusAlertRuleResponseBody
}

type CreatePrometheusAlertRuleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePrometheusAlertRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePrometheusAlertRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePrometheusAlertRuleResponse) GoString() string {
	return s.String()
}

func (s *CreatePrometheusAlertRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePrometheusAlertRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePrometheusAlertRuleResponse) GetBody() *CreatePrometheusAlertRuleResponseBody {
	return s.Body
}

func (s *CreatePrometheusAlertRuleResponse) SetHeaders(v map[string]*string) *CreatePrometheusAlertRuleResponse {
	s.Headers = v
	return s
}

func (s *CreatePrometheusAlertRuleResponse) SetStatusCode(v int32) *CreatePrometheusAlertRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponse) SetBody(v *CreatePrometheusAlertRuleResponseBody) *CreatePrometheusAlertRuleResponse {
	s.Body = v
	return s
}

func (s *CreatePrometheusAlertRuleResponse) Validate() error {
	return dara.Validate(s)
}
