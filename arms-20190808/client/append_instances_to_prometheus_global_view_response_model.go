// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppendInstancesToPrometheusGlobalViewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AppendInstancesToPrometheusGlobalViewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AppendInstancesToPrometheusGlobalViewResponse
	GetStatusCode() *int32
	SetBody(v *AppendInstancesToPrometheusGlobalViewResponseBody) *AppendInstancesToPrometheusGlobalViewResponse
	GetBody() *AppendInstancesToPrometheusGlobalViewResponseBody
}

type AppendInstancesToPrometheusGlobalViewResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AppendInstancesToPrometheusGlobalViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AppendInstancesToPrometheusGlobalViewResponse) String() string {
	return dara.Prettify(s)
}

func (s AppendInstancesToPrometheusGlobalViewResponse) GoString() string {
	return s.String()
}

func (s *AppendInstancesToPrometheusGlobalViewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AppendInstancesToPrometheusGlobalViewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AppendInstancesToPrometheusGlobalViewResponse) GetBody() *AppendInstancesToPrometheusGlobalViewResponseBody {
	return s.Body
}

func (s *AppendInstancesToPrometheusGlobalViewResponse) SetHeaders(v map[string]*string) *AppendInstancesToPrometheusGlobalViewResponse {
	s.Headers = v
	return s
}

func (s *AppendInstancesToPrometheusGlobalViewResponse) SetStatusCode(v int32) *AppendInstancesToPrometheusGlobalViewResponse {
	s.StatusCode = &v
	return s
}

func (s *AppendInstancesToPrometheusGlobalViewResponse) SetBody(v *AppendInstancesToPrometheusGlobalViewResponseBody) *AppendInstancesToPrometheusGlobalViewResponse {
	s.Body = v
	return s
}

func (s *AppendInstancesToPrometheusGlobalViewResponse) Validate() error {
	return dara.Validate(s)
}
