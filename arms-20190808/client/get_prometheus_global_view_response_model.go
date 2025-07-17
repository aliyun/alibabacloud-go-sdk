// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPrometheusGlobalViewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPrometheusGlobalViewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPrometheusGlobalViewResponse
	GetStatusCode() *int32
	SetBody(v *GetPrometheusGlobalViewResponseBody) *GetPrometheusGlobalViewResponse
	GetBody() *GetPrometheusGlobalViewResponseBody
}

type GetPrometheusGlobalViewResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPrometheusGlobalViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPrometheusGlobalViewResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPrometheusGlobalViewResponse) GoString() string {
	return s.String()
}

func (s *GetPrometheusGlobalViewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPrometheusGlobalViewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPrometheusGlobalViewResponse) GetBody() *GetPrometheusGlobalViewResponseBody {
	return s.Body
}

func (s *GetPrometheusGlobalViewResponse) SetHeaders(v map[string]*string) *GetPrometheusGlobalViewResponse {
	s.Headers = v
	return s
}

func (s *GetPrometheusGlobalViewResponse) SetStatusCode(v int32) *GetPrometheusGlobalViewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPrometheusGlobalViewResponse) SetBody(v *GetPrometheusGlobalViewResponseBody) *GetPrometheusGlobalViewResponse {
	s.Body = v
	return s
}

func (s *GetPrometheusGlobalViewResponse) Validate() error {
	return dara.Validate(s)
}
