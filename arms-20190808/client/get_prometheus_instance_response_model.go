// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPrometheusInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPrometheusInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPrometheusInstanceResponse
	GetStatusCode() *int32
	SetBody(v *GetPrometheusInstanceResponseBody) *GetPrometheusInstanceResponse
	GetBody() *GetPrometheusInstanceResponseBody
}

type GetPrometheusInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPrometheusInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPrometheusInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPrometheusInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetPrometheusInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPrometheusInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPrometheusInstanceResponse) GetBody() *GetPrometheusInstanceResponseBody {
	return s.Body
}

func (s *GetPrometheusInstanceResponse) SetHeaders(v map[string]*string) *GetPrometheusInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetPrometheusInstanceResponse) SetStatusCode(v int32) *GetPrometheusInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPrometheusInstanceResponse) SetBody(v *GetPrometheusInstanceResponseBody) *GetPrometheusInstanceResponse {
	s.Body = v
	return s
}

func (s *GetPrometheusInstanceResponse) Validate() error {
	return dara.Validate(s)
}
