// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrometheusAlertRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePrometheusAlertRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePrometheusAlertRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePrometheusAlertRuleResponseBody) *UpdatePrometheusAlertRuleResponse
	GetBody() *UpdatePrometheusAlertRuleResponseBody
}

type UpdatePrometheusAlertRuleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePrometheusAlertRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePrometheusAlertRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusAlertRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusAlertRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePrometheusAlertRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePrometheusAlertRuleResponse) GetBody() *UpdatePrometheusAlertRuleResponseBody {
	return s.Body
}

func (s *UpdatePrometheusAlertRuleResponse) SetHeaders(v map[string]*string) *UpdatePrometheusAlertRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdatePrometheusAlertRuleResponse) SetStatusCode(v int32) *UpdatePrometheusAlertRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponse) SetBody(v *UpdatePrometheusAlertRuleResponseBody) *UpdatePrometheusAlertRuleResponse {
	s.Body = v
	return s
}

func (s *UpdatePrometheusAlertRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
