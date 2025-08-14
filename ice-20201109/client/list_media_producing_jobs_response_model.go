// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaProducingJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMediaProducingJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMediaProducingJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListMediaProducingJobsResponseBody) *ListMediaProducingJobsResponse
	GetBody() *ListMediaProducingJobsResponseBody
}

type ListMediaProducingJobsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMediaProducingJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMediaProducingJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMediaProducingJobsResponse) GoString() string {
	return s.String()
}

func (s *ListMediaProducingJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMediaProducingJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMediaProducingJobsResponse) GetBody() *ListMediaProducingJobsResponseBody {
	return s.Body
}

func (s *ListMediaProducingJobsResponse) SetHeaders(v map[string]*string) *ListMediaProducingJobsResponse {
	s.Headers = v
	return s
}

func (s *ListMediaProducingJobsResponse) SetStatusCode(v int32) *ListMediaProducingJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMediaProducingJobsResponse) SetBody(v *ListMediaProducingJobsResponseBody) *ListMediaProducingJobsResponse {
	s.Body = v
	return s
}

func (s *ListMediaProducingJobsResponse) Validate() error {
	return dara.Validate(s)
}
