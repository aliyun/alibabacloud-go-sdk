// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorApplicationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDoctorApplicationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDoctorApplicationsResponse
	GetStatusCode() *int32
	SetBody(v *ListDoctorApplicationsResponseBody) *ListDoctorApplicationsResponse
	GetBody() *ListDoctorApplicationsResponseBody
}

type ListDoctorApplicationsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDoctorApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDoctorApplicationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorApplicationsResponse) GoString() string {
	return s.String()
}

func (s *ListDoctorApplicationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDoctorApplicationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDoctorApplicationsResponse) GetBody() *ListDoctorApplicationsResponseBody {
	return s.Body
}

func (s *ListDoctorApplicationsResponse) SetHeaders(v map[string]*string) *ListDoctorApplicationsResponse {
	s.Headers = v
	return s
}

func (s *ListDoctorApplicationsResponse) SetStatusCode(v int32) *ListDoctorApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDoctorApplicationsResponse) SetBody(v *ListDoctorApplicationsResponseBody) *ListDoctorApplicationsResponse {
	s.Body = v
	return s
}

func (s *ListDoctorApplicationsResponse) Validate() error {
	return dara.Validate(s)
}
