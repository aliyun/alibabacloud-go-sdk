// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceMetricsResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceMetricsResponseBody) *GetInstanceMetricsResponse
	GetBody() *GetInstanceMetricsResponseBody
}

type GetInstanceMetricsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceMetricsResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceMetricsResponse) GetBody() *GetInstanceMetricsResponseBody {
	return s.Body
}

func (s *GetInstanceMetricsResponse) SetHeaders(v map[string]*string) *GetInstanceMetricsResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceMetricsResponse) SetStatusCode(v int32) *GetInstanceMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceMetricsResponse) SetBody(v *GetInstanceMetricsResponseBody) *GetInstanceMetricsResponse {
	s.Body = v
	return s
}

func (s *GetInstanceMetricsResponse) Validate() error {
	return dara.Validate(s)
}
