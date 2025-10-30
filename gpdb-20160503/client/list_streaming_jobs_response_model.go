// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStreamingJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListStreamingJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListStreamingJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListStreamingJobsResponseBody) *ListStreamingJobsResponse
	GetBody() *ListStreamingJobsResponseBody
}

type ListStreamingJobsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStreamingJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStreamingJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListStreamingJobsResponse) GoString() string {
	return s.String()
}

func (s *ListStreamingJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListStreamingJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListStreamingJobsResponse) GetBody() *ListStreamingJobsResponseBody {
	return s.Body
}

func (s *ListStreamingJobsResponse) SetHeaders(v map[string]*string) *ListStreamingJobsResponse {
	s.Headers = v
	return s
}

func (s *ListStreamingJobsResponse) SetStatusCode(v int32) *ListStreamingJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStreamingJobsResponse) SetBody(v *ListStreamingJobsResponseBody) *ListStreamingJobsResponse {
	s.Body = v
	return s
}

func (s *ListStreamingJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
