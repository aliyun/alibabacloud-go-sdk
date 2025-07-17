// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvCustomJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEnvCustomJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEnvCustomJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListEnvCustomJobsResponseBody) *ListEnvCustomJobsResponse
	GetBody() *ListEnvCustomJobsResponseBody
}

type ListEnvCustomJobsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEnvCustomJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnvCustomJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEnvCustomJobsResponse) GoString() string {
	return s.String()
}

func (s *ListEnvCustomJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEnvCustomJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEnvCustomJobsResponse) GetBody() *ListEnvCustomJobsResponseBody {
	return s.Body
}

func (s *ListEnvCustomJobsResponse) SetHeaders(v map[string]*string) *ListEnvCustomJobsResponse {
	s.Headers = v
	return s
}

func (s *ListEnvCustomJobsResponse) SetStatusCode(v int32) *ListEnvCustomJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnvCustomJobsResponse) SetBody(v *ListEnvCustomJobsResponseBody) *ListEnvCustomJobsResponse {
	s.Body = v
	return s
}

func (s *ListEnvCustomJobsResponse) Validate() error {
	return dara.Validate(s)
}
