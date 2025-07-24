// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDoctorJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDoctorJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListDoctorJobsResponseBody) *ListDoctorJobsResponse
	GetBody() *ListDoctorJobsResponseBody
}

type ListDoctorJobsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDoctorJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDoctorJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorJobsResponse) GoString() string {
	return s.String()
}

func (s *ListDoctorJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDoctorJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDoctorJobsResponse) GetBody() *ListDoctorJobsResponseBody {
	return s.Body
}

func (s *ListDoctorJobsResponse) SetHeaders(v map[string]*string) *ListDoctorJobsResponse {
	s.Headers = v
	return s
}

func (s *ListDoctorJobsResponse) SetStatusCode(v int32) *ListDoctorJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDoctorJobsResponse) SetBody(v *ListDoctorJobsResponseBody) *ListDoctorJobsResponse {
	s.Body = v
	return s
}

func (s *ListDoctorJobsResponse) Validate() error {
	return dara.Validate(s)
}
