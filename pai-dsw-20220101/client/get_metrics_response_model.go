// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMetricsResponse
	GetStatusCode() *int32
	SetBody(v *GetMetricsResponseBody) *GetMetricsResponse
	GetBody() *GetMetricsResponseBody
}

type GetMetricsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMetricsResponse) GoString() string {
	return s.String()
}

func (s *GetMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMetricsResponse) GetBody() *GetMetricsResponseBody {
	return s.Body
}

func (s *GetMetricsResponse) SetHeaders(v map[string]*string) *GetMetricsResponse {
	s.Headers = v
	return s
}

func (s *GetMetricsResponse) SetStatusCode(v int32) *GetMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetricsResponse) SetBody(v *GetMetricsResponseBody) *GetMetricsResponse {
	s.Body = v
	return s
}

func (s *GetMetricsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
