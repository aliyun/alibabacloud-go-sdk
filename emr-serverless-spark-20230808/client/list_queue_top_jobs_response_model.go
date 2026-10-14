// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueueTopJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQueueTopJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQueueTopJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListQueueTopJobsResponseBody) *ListQueueTopJobsResponse
	GetBody() *ListQueueTopJobsResponseBody
}

type ListQueueTopJobsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQueueTopJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQueueTopJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQueueTopJobsResponse) GoString() string {
	return s.String()
}

func (s *ListQueueTopJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQueueTopJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQueueTopJobsResponse) GetBody() *ListQueueTopJobsResponseBody {
	return s.Body
}

func (s *ListQueueTopJobsResponse) SetHeaders(v map[string]*string) *ListQueueTopJobsResponse {
	s.Headers = v
	return s
}

func (s *ListQueueTopJobsResponse) SetStatusCode(v int32) *ListQueueTopJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQueueTopJobsResponse) SetBody(v *ListQueueTopJobsResponseBody) *ListQueueTopJobsResponse {
	s.Body = v
	return s
}

func (s *ListQueueTopJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
