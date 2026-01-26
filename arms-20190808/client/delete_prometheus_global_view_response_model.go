// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrometheusGlobalViewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePrometheusGlobalViewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePrometheusGlobalViewResponse
	GetStatusCode() *int32
	SetBody(v *DeletePrometheusGlobalViewResponseBody) *DeletePrometheusGlobalViewResponse
	GetBody() *DeletePrometheusGlobalViewResponseBody
}

type DeletePrometheusGlobalViewResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePrometheusGlobalViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePrometheusGlobalViewResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePrometheusGlobalViewResponse) GoString() string {
	return s.String()
}

func (s *DeletePrometheusGlobalViewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePrometheusGlobalViewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePrometheusGlobalViewResponse) GetBody() *DeletePrometheusGlobalViewResponseBody {
	return s.Body
}

func (s *DeletePrometheusGlobalViewResponse) SetHeaders(v map[string]*string) *DeletePrometheusGlobalViewResponse {
	s.Headers = v
	return s
}

func (s *DeletePrometheusGlobalViewResponse) SetStatusCode(v int32) *DeletePrometheusGlobalViewResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePrometheusGlobalViewResponse) SetBody(v *DeletePrometheusGlobalViewResponseBody) *DeletePrometheusGlobalViewResponse {
	s.Body = v
	return s
}

func (s *DeletePrometheusGlobalViewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
