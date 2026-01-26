// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusAlertTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPrometheusAlertTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPrometheusAlertTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListPrometheusAlertTemplatesResponseBody) *ListPrometheusAlertTemplatesResponse
	GetBody() *ListPrometheusAlertTemplatesResponseBody
}

type ListPrometheusAlertTemplatesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrometheusAlertTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrometheusAlertTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusAlertTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPrometheusAlertTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPrometheusAlertTemplatesResponse) GetBody() *ListPrometheusAlertTemplatesResponseBody {
	return s.Body
}

func (s *ListPrometheusAlertTemplatesResponse) SetHeaders(v map[string]*string) *ListPrometheusAlertTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListPrometheusAlertTemplatesResponse) SetStatusCode(v int32) *ListPrometheusAlertTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrometheusAlertTemplatesResponse) SetBody(v *ListPrometheusAlertTemplatesResponseBody) *ListPrometheusAlertTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListPrometheusAlertTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
