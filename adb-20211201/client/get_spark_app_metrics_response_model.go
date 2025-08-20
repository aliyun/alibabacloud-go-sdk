// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkAppMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSparkAppMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSparkAppMetricsResponse
	GetStatusCode() *int32
	SetBody(v *GetSparkAppMetricsResponseBody) *GetSparkAppMetricsResponse
	GetBody() *GetSparkAppMetricsResponseBody
}

type GetSparkAppMetricsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSparkAppMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSparkAppMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSparkAppMetricsResponse) GoString() string {
	return s.String()
}

func (s *GetSparkAppMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSparkAppMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSparkAppMetricsResponse) GetBody() *GetSparkAppMetricsResponseBody {
	return s.Body
}

func (s *GetSparkAppMetricsResponse) SetHeaders(v map[string]*string) *GetSparkAppMetricsResponse {
	s.Headers = v
	return s
}

func (s *GetSparkAppMetricsResponse) SetStatusCode(v int32) *GetSparkAppMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSparkAppMetricsResponse) SetBody(v *GetSparkAppMetricsResponseBody) *GetSparkAppMetricsResponse {
	s.Body = v
	return s
}

func (s *GetSparkAppMetricsResponse) Validate() error {
	return dara.Validate(s)
}
