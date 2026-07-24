// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListComputeJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListComputeJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListComputeJobsResponseBody) *ListComputeJobsResponse
	GetBody() *ListComputeJobsResponseBody
}

type ListComputeJobsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListComputeJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListComputeJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListComputeJobsResponse) GoString() string {
	return s.String()
}

func (s *ListComputeJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListComputeJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListComputeJobsResponse) GetBody() *ListComputeJobsResponseBody {
	return s.Body
}

func (s *ListComputeJobsResponse) SetHeaders(v map[string]*string) *ListComputeJobsResponse {
	s.Headers = v
	return s
}

func (s *ListComputeJobsResponse) SetStatusCode(v int32) *ListComputeJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListComputeJobsResponse) SetBody(v *ListComputeJobsResponseBody) *ListComputeJobsResponse {
	s.Body = v
	return s
}

func (s *ListComputeJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
