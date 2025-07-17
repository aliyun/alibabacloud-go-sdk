// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrometheusGlobalViewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePrometheusGlobalViewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePrometheusGlobalViewResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePrometheusGlobalViewResponseBody) *UpdatePrometheusGlobalViewResponse
	GetBody() *UpdatePrometheusGlobalViewResponseBody
}

type UpdatePrometheusGlobalViewResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePrometheusGlobalViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePrometheusGlobalViewResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusGlobalViewResponse) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusGlobalViewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePrometheusGlobalViewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePrometheusGlobalViewResponse) GetBody() *UpdatePrometheusGlobalViewResponseBody {
	return s.Body
}

func (s *UpdatePrometheusGlobalViewResponse) SetHeaders(v map[string]*string) *UpdatePrometheusGlobalViewResponse {
	s.Headers = v
	return s
}

func (s *UpdatePrometheusGlobalViewResponse) SetStatusCode(v int32) *UpdatePrometheusGlobalViewResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePrometheusGlobalViewResponse) SetBody(v *UpdatePrometheusGlobalViewResponseBody) *UpdatePrometheusGlobalViewResponse {
	s.Body = v
	return s
}

func (s *UpdatePrometheusGlobalViewResponse) Validate() error {
	return dara.Validate(s)
}
