// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportPrometheusRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportPrometheusRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportPrometheusRulesResponse
	GetStatusCode() *int32
	SetBody(v *ImportPrometheusRulesResponseBody) *ImportPrometheusRulesResponse
	GetBody() *ImportPrometheusRulesResponseBody
}

type ImportPrometheusRulesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportPrometheusRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportPrometheusRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportPrometheusRulesResponse) GoString() string {
	return s.String()
}

func (s *ImportPrometheusRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportPrometheusRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportPrometheusRulesResponse) GetBody() *ImportPrometheusRulesResponseBody {
	return s.Body
}

func (s *ImportPrometheusRulesResponse) SetHeaders(v map[string]*string) *ImportPrometheusRulesResponse {
	s.Headers = v
	return s
}

func (s *ImportPrometheusRulesResponse) SetStatusCode(v int32) *ImportPrometheusRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportPrometheusRulesResponse) SetBody(v *ImportPrometheusRulesResponseBody) *ImportPrometheusRulesResponse {
	s.Body = v
	return s
}

func (s *ImportPrometheusRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
