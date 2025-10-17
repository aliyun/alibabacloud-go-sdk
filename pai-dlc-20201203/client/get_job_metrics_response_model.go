// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetJobMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetJobMetricsResponse
	GetStatusCode() *int32
	SetBody(v *GetJobMetricsResponseBody) *GetJobMetricsResponse
	GetBody() *GetJobMetricsResponseBody
}

type GetJobMetricsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetJobMetricsResponse) GoString() string {
	return s.String()
}

func (s *GetJobMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetJobMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetJobMetricsResponse) GetBody() *GetJobMetricsResponseBody {
	return s.Body
}

func (s *GetJobMetricsResponse) SetHeaders(v map[string]*string) *GetJobMetricsResponse {
	s.Headers = v
	return s
}

func (s *GetJobMetricsResponse) SetStatusCode(v int32) *GetJobMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobMetricsResponse) SetBody(v *GetJobMetricsResponseBody) *GetJobMetricsResponse {
	s.Body = v
	return s
}

func (s *GetJobMetricsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
