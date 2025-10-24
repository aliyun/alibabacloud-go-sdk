// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListJobMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListJobMetricResponse
	GetStatusCode() *int32
	SetBody(v *ListJobMetricResponseBody) *ListJobMetricResponse
	GetBody() *ListJobMetricResponseBody
}

type ListJobMetricResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJobMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s ListJobMetricResponse) GoString() string {
	return s.String()
}

func (s *ListJobMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListJobMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListJobMetricResponse) GetBody() *ListJobMetricResponseBody {
	return s.Body
}

func (s *ListJobMetricResponse) SetHeaders(v map[string]*string) *ListJobMetricResponse {
	s.Headers = v
	return s
}

func (s *ListJobMetricResponse) SetStatusCode(v int32) *ListJobMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobMetricResponse) SetBody(v *ListJobMetricResponseBody) *ListJobMetricResponse {
	s.Body = v
	return s
}

func (s *ListJobMetricResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
