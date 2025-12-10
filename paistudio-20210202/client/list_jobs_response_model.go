// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListJobsResponseBody) *ListJobsResponse
	GetBody() *ListJobsResponseBody
}

type ListJobsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponse) GoString() string {
	return s.String()
}

func (s *ListJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListJobsResponse) GetBody() *ListJobsResponseBody {
	return s.Body
}

func (s *ListJobsResponse) SetHeaders(v map[string]*string) *ListJobsResponse {
	s.Headers = v
	return s
}

func (s *ListJobsResponse) SetStatusCode(v int32) *ListJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobsResponse) SetBody(v *ListJobsResponseBody) *ListJobsResponse {
	s.Body = v
	return s
}

func (s *ListJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
