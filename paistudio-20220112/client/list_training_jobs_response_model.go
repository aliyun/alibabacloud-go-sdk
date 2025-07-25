// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrainingJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTrainingJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTrainingJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListTrainingJobsResponseBody) *ListTrainingJobsResponse
	GetBody() *ListTrainingJobsResponseBody
}

type ListTrainingJobsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTrainingJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTrainingJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobsResponse) GoString() string {
	return s.String()
}

func (s *ListTrainingJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTrainingJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTrainingJobsResponse) GetBody() *ListTrainingJobsResponseBody {
	return s.Body
}

func (s *ListTrainingJobsResponse) SetHeaders(v map[string]*string) *ListTrainingJobsResponse {
	s.Headers = v
	return s
}

func (s *ListTrainingJobsResponse) SetStatusCode(v int32) *ListTrainingJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrainingJobsResponse) SetBody(v *ListTrainingJobsResponseBody) *ListTrainingJobsResponse {
	s.Body = v
	return s
}

func (s *ListTrainingJobsResponse) Validate() error {
	return dara.Validate(s)
}
