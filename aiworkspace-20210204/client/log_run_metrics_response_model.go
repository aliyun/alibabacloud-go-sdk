// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLogRunMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LogRunMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LogRunMetricsResponse
	GetStatusCode() *int32
	SetBody(v *LogRunMetricsResponseBody) *LogRunMetricsResponse
	GetBody() *LogRunMetricsResponseBody
}

type LogRunMetricsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LogRunMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LogRunMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s LogRunMetricsResponse) GoString() string {
	return s.String()
}

func (s *LogRunMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LogRunMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LogRunMetricsResponse) GetBody() *LogRunMetricsResponseBody {
	return s.Body
}

func (s *LogRunMetricsResponse) SetHeaders(v map[string]*string) *LogRunMetricsResponse {
	s.Headers = v
	return s
}

func (s *LogRunMetricsResponse) SetStatusCode(v int32) *LogRunMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *LogRunMetricsResponse) SetBody(v *LogRunMetricsResponseBody) *LogRunMetricsResponse {
	s.Body = v
	return s
}

func (s *LogRunMetricsResponse) Validate() error {
	return dara.Validate(s)
}
