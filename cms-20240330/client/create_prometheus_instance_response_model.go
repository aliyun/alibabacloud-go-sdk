// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrometheusInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePrometheusInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePrometheusInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreatePrometheusInstanceResponseBody) *CreatePrometheusInstanceResponse
	GetBody() *CreatePrometheusInstanceResponseBody
}

type CreatePrometheusInstanceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePrometheusInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePrometheusInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePrometheusInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreatePrometheusInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePrometheusInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePrometheusInstanceResponse) GetBody() *CreatePrometheusInstanceResponseBody {
	return s.Body
}

func (s *CreatePrometheusInstanceResponse) SetHeaders(v map[string]*string) *CreatePrometheusInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreatePrometheusInstanceResponse) SetStatusCode(v int32) *CreatePrometheusInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePrometheusInstanceResponse) SetBody(v *CreatePrometheusInstanceResponseBody) *CreatePrometheusInstanceResponse {
	s.Body = v
	return s
}

func (s *CreatePrometheusInstanceResponse) Validate() error {
	return dara.Validate(s)
}
