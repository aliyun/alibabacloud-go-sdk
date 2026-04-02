// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrometheusVirtualInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePrometheusVirtualInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePrometheusVirtualInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DeletePrometheusVirtualInstanceResponseBody) *DeletePrometheusVirtualInstanceResponse
	GetBody() *DeletePrometheusVirtualInstanceResponseBody
}

type DeletePrometheusVirtualInstanceResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePrometheusVirtualInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePrometheusVirtualInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePrometheusVirtualInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeletePrometheusVirtualInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePrometheusVirtualInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePrometheusVirtualInstanceResponse) GetBody() *DeletePrometheusVirtualInstanceResponseBody {
	return s.Body
}

func (s *DeletePrometheusVirtualInstanceResponse) SetHeaders(v map[string]*string) *DeletePrometheusVirtualInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeletePrometheusVirtualInstanceResponse) SetStatusCode(v int32) *DeletePrometheusVirtualInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePrometheusVirtualInstanceResponse) SetBody(v *DeletePrometheusVirtualInstanceResponseBody) *DeletePrometheusVirtualInstanceResponse {
	s.Body = v
	return s
}

func (s *DeletePrometheusVirtualInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
