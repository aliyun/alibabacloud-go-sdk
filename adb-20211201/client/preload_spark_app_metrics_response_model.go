// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreloadSparkAppMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PreloadSparkAppMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PreloadSparkAppMetricsResponse
	GetStatusCode() *int32
	SetBody(v *PreloadSparkAppMetricsResponseBody) *PreloadSparkAppMetricsResponse
	GetBody() *PreloadSparkAppMetricsResponseBody
}

type PreloadSparkAppMetricsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreloadSparkAppMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreloadSparkAppMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s PreloadSparkAppMetricsResponse) GoString() string {
	return s.String()
}

func (s *PreloadSparkAppMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PreloadSparkAppMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PreloadSparkAppMetricsResponse) GetBody() *PreloadSparkAppMetricsResponseBody {
	return s.Body
}

func (s *PreloadSparkAppMetricsResponse) SetHeaders(v map[string]*string) *PreloadSparkAppMetricsResponse {
	s.Headers = v
	return s
}

func (s *PreloadSparkAppMetricsResponse) SetStatusCode(v int32) *PreloadSparkAppMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *PreloadSparkAppMetricsResponse) SetBody(v *PreloadSparkAppMetricsResponseBody) *PreloadSparkAppMetricsResponse {
	s.Body = v
	return s
}

func (s *PreloadSparkAppMetricsResponse) Validate() error {
	return dara.Validate(s)
}
