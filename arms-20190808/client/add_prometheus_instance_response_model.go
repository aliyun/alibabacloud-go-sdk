// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPrometheusInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddPrometheusInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddPrometheusInstanceResponse
	GetStatusCode() *int32
	SetBody(v *AddPrometheusInstanceResponseBody) *AddPrometheusInstanceResponse
	GetBody() *AddPrometheusInstanceResponseBody
}

type AddPrometheusInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddPrometheusInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddPrometheusInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s AddPrometheusInstanceResponse) GoString() string {
	return s.String()
}

func (s *AddPrometheusInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddPrometheusInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddPrometheusInstanceResponse) GetBody() *AddPrometheusInstanceResponseBody {
	return s.Body
}

func (s *AddPrometheusInstanceResponse) SetHeaders(v map[string]*string) *AddPrometheusInstanceResponse {
	s.Headers = v
	return s
}

func (s *AddPrometheusInstanceResponse) SetStatusCode(v int32) *AddPrometheusInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddPrometheusInstanceResponse) SetBody(v *AddPrometheusInstanceResponseBody) *AddPrometheusInstanceResponse {
	s.Body = v
	return s
}

func (s *AddPrometheusInstanceResponse) Validate() error {
	return dara.Validate(s)
}
