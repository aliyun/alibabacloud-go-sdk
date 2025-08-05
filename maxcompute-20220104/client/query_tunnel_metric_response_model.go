// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTunnelMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTunnelMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTunnelMetricResponse
	GetStatusCode() *int32
	SetBody(v *QueryTunnelMetricResponseBody) *QueryTunnelMetricResponse
	GetBody() *QueryTunnelMetricResponseBody
}

type QueryTunnelMetricResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTunnelMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTunnelMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTunnelMetricResponse) GoString() string {
	return s.String()
}

func (s *QueryTunnelMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTunnelMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTunnelMetricResponse) GetBody() *QueryTunnelMetricResponseBody {
	return s.Body
}

func (s *QueryTunnelMetricResponse) SetHeaders(v map[string]*string) *QueryTunnelMetricResponse {
	s.Headers = v
	return s
}

func (s *QueryTunnelMetricResponse) SetStatusCode(v int32) *QueryTunnelMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTunnelMetricResponse) SetBody(v *QueryTunnelMetricResponseBody) *QueryTunnelMetricResponse {
	s.Body = v
	return s
}

func (s *QueryTunnelMetricResponse) Validate() error {
	return dara.Validate(s)
}
