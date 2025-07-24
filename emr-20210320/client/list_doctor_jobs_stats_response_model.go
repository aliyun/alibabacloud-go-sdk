// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorJobsStatsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDoctorJobsStatsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDoctorJobsStatsResponse
	GetStatusCode() *int32
	SetBody(v *ListDoctorJobsStatsResponseBody) *ListDoctorJobsStatsResponse
	GetBody() *ListDoctorJobsStatsResponseBody
}

type ListDoctorJobsStatsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDoctorJobsStatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDoctorJobsStatsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorJobsStatsResponse) GoString() string {
	return s.String()
}

func (s *ListDoctorJobsStatsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDoctorJobsStatsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDoctorJobsStatsResponse) GetBody() *ListDoctorJobsStatsResponseBody {
	return s.Body
}

func (s *ListDoctorJobsStatsResponse) SetHeaders(v map[string]*string) *ListDoctorJobsStatsResponse {
	s.Headers = v
	return s
}

func (s *ListDoctorJobsStatsResponse) SetStatusCode(v int32) *ListDoctorJobsStatsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDoctorJobsStatsResponse) SetBody(v *ListDoctorJobsStatsResponseBody) *ListDoctorJobsStatsResponse {
	s.Body = v
	return s
}

func (s *ListDoctorJobsStatsResponse) Validate() error {
	return dara.Validate(s)
}
