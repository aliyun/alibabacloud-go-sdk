// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrometheusAlertRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePrometheusAlertRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePrometheusAlertRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeletePrometheusAlertRuleResponseBody) *DeletePrometheusAlertRuleResponse
	GetBody() *DeletePrometheusAlertRuleResponseBody
}

type DeletePrometheusAlertRuleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePrometheusAlertRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePrometheusAlertRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePrometheusAlertRuleResponse) GoString() string {
	return s.String()
}

func (s *DeletePrometheusAlertRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePrometheusAlertRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePrometheusAlertRuleResponse) GetBody() *DeletePrometheusAlertRuleResponseBody {
	return s.Body
}

func (s *DeletePrometheusAlertRuleResponse) SetHeaders(v map[string]*string) *DeletePrometheusAlertRuleResponse {
	s.Headers = v
	return s
}

func (s *DeletePrometheusAlertRuleResponse) SetStatusCode(v int32) *DeletePrometheusAlertRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePrometheusAlertRuleResponse) SetBody(v *DeletePrometheusAlertRuleResponseBody) *DeletePrometheusAlertRuleResponse {
	s.Body = v
	return s
}

func (s *DeletePrometheusAlertRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
