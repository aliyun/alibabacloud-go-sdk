// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrometheusInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePrometheusInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePrometheusInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DeletePrometheusInstanceResponseBody) *DeletePrometheusInstanceResponse
	GetBody() *DeletePrometheusInstanceResponseBody
}

type DeletePrometheusInstanceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePrometheusInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePrometheusInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePrometheusInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeletePrometheusInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePrometheusInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePrometheusInstanceResponse) GetBody() *DeletePrometheusInstanceResponseBody {
	return s.Body
}

func (s *DeletePrometheusInstanceResponse) SetHeaders(v map[string]*string) *DeletePrometheusInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeletePrometheusInstanceResponse) SetStatusCode(v int32) *DeletePrometheusInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePrometheusInstanceResponse) SetBody(v *DeletePrometheusInstanceResponseBody) *DeletePrometheusInstanceResponse {
	s.Body = v
	return s
}

func (s *DeletePrometheusInstanceResponse) Validate() error {
	return dara.Validate(s)
}
