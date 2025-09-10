// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrometheusViewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePrometheusViewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePrometheusViewResponse
	GetStatusCode() *int32
	SetBody(v *DeletePrometheusViewResponseBody) *DeletePrometheusViewResponse
	GetBody() *DeletePrometheusViewResponseBody
}

type DeletePrometheusViewResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePrometheusViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePrometheusViewResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePrometheusViewResponse) GoString() string {
	return s.String()
}

func (s *DeletePrometheusViewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePrometheusViewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePrometheusViewResponse) GetBody() *DeletePrometheusViewResponseBody {
	return s.Body
}

func (s *DeletePrometheusViewResponse) SetHeaders(v map[string]*string) *DeletePrometheusViewResponse {
	s.Headers = v
	return s
}

func (s *DeletePrometheusViewResponse) SetStatusCode(v int32) *DeletePrometheusViewResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePrometheusViewResponse) SetBody(v *DeletePrometheusViewResponseBody) *DeletePrometheusViewResponse {
	s.Body = v
	return s
}

func (s *DeletePrometheusViewResponse) Validate() error {
	return dara.Validate(s)
}
