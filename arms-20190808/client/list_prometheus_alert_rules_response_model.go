// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusAlertRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPrometheusAlertRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPrometheusAlertRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListPrometheusAlertRulesResponseBody) *ListPrometheusAlertRulesResponse
	GetBody() *ListPrometheusAlertRulesResponseBody
}

type ListPrometheusAlertRulesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrometheusAlertRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrometheusAlertRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusAlertRulesResponse) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPrometheusAlertRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPrometheusAlertRulesResponse) GetBody() *ListPrometheusAlertRulesResponseBody {
	return s.Body
}

func (s *ListPrometheusAlertRulesResponse) SetHeaders(v map[string]*string) *ListPrometheusAlertRulesResponse {
	s.Headers = v
	return s
}

func (s *ListPrometheusAlertRulesResponse) SetStatusCode(v int32) *ListPrometheusAlertRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrometheusAlertRulesResponse) SetBody(v *ListPrometheusAlertRulesResponseBody) *ListPrometheusAlertRulesResponse {
	s.Body = v
	return s
}

func (s *ListPrometheusAlertRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
