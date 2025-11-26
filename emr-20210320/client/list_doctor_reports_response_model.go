// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorReportsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDoctorReportsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDoctorReportsResponse
	GetStatusCode() *int32
	SetBody(v *ListDoctorReportsResponseBody) *ListDoctorReportsResponse
	GetBody() *ListDoctorReportsResponseBody
}

type ListDoctorReportsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDoctorReportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDoctorReportsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorReportsResponse) GoString() string {
	return s.String()
}

func (s *ListDoctorReportsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDoctorReportsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDoctorReportsResponse) GetBody() *ListDoctorReportsResponseBody {
	return s.Body
}

func (s *ListDoctorReportsResponse) SetHeaders(v map[string]*string) *ListDoctorReportsResponse {
	s.Headers = v
	return s
}

func (s *ListDoctorReportsResponse) SetStatusCode(v int32) *ListDoctorReportsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDoctorReportsResponse) SetBody(v *ListDoctorReportsResponseBody) *ListDoctorReportsResponse {
	s.Body = v
	return s
}

func (s *ListDoctorReportsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
