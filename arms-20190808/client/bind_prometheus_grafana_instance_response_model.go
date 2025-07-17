// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindPrometheusGrafanaInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindPrometheusGrafanaInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindPrometheusGrafanaInstanceResponse
	GetStatusCode() *int32
	SetBody(v *BindPrometheusGrafanaInstanceResponseBody) *BindPrometheusGrafanaInstanceResponse
	GetBody() *BindPrometheusGrafanaInstanceResponseBody
}

type BindPrometheusGrafanaInstanceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindPrometheusGrafanaInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindPrometheusGrafanaInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s BindPrometheusGrafanaInstanceResponse) GoString() string {
	return s.String()
}

func (s *BindPrometheusGrafanaInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindPrometheusGrafanaInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindPrometheusGrafanaInstanceResponse) GetBody() *BindPrometheusGrafanaInstanceResponseBody {
	return s.Body
}

func (s *BindPrometheusGrafanaInstanceResponse) SetHeaders(v map[string]*string) *BindPrometheusGrafanaInstanceResponse {
	s.Headers = v
	return s
}

func (s *BindPrometheusGrafanaInstanceResponse) SetStatusCode(v int32) *BindPrometheusGrafanaInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *BindPrometheusGrafanaInstanceResponse) SetBody(v *BindPrometheusGrafanaInstanceResponseBody) *BindPrometheusGrafanaInstanceResponse {
	s.Body = v
	return s
}

func (s *BindPrometheusGrafanaInstanceResponse) Validate() error {
	return dara.Validate(s)
}
