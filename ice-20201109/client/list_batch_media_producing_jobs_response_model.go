// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBatchMediaProducingJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBatchMediaProducingJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBatchMediaProducingJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListBatchMediaProducingJobsResponseBody) *ListBatchMediaProducingJobsResponse
	GetBody() *ListBatchMediaProducingJobsResponseBody
}

type ListBatchMediaProducingJobsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBatchMediaProducingJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBatchMediaProducingJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBatchMediaProducingJobsResponse) GoString() string {
	return s.String()
}

func (s *ListBatchMediaProducingJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBatchMediaProducingJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBatchMediaProducingJobsResponse) GetBody() *ListBatchMediaProducingJobsResponseBody {
	return s.Body
}

func (s *ListBatchMediaProducingJobsResponse) SetHeaders(v map[string]*string) *ListBatchMediaProducingJobsResponse {
	s.Headers = v
	return s
}

func (s *ListBatchMediaProducingJobsResponse) SetStatusCode(v int32) *ListBatchMediaProducingJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBatchMediaProducingJobsResponse) SetBody(v *ListBatchMediaProducingJobsResponseBody) *ListBatchMediaProducingJobsResponse {
	s.Body = v
	return s
}

func (s *ListBatchMediaProducingJobsResponse) Validate() error {
	return dara.Validate(s)
}
