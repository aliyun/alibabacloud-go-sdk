// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTunnelMetricDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTunnelMetricDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTunnelMetricDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryTunnelMetricDetailResponseBody) *QueryTunnelMetricDetailResponse
	GetBody() *QueryTunnelMetricDetailResponseBody
}

type QueryTunnelMetricDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTunnelMetricDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTunnelMetricDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTunnelMetricDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryTunnelMetricDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTunnelMetricDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTunnelMetricDetailResponse) GetBody() *QueryTunnelMetricDetailResponseBody {
	return s.Body
}

func (s *QueryTunnelMetricDetailResponse) SetHeaders(v map[string]*string) *QueryTunnelMetricDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryTunnelMetricDetailResponse) SetStatusCode(v int32) *QueryTunnelMetricDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTunnelMetricDetailResponse) SetBody(v *QueryTunnelMetricDetailResponseBody) *QueryTunnelMetricDetailResponse {
	s.Body = v
	return s
}

func (s *QueryTunnelMetricDetailResponse) Validate() error {
	return dara.Validate(s)
}
