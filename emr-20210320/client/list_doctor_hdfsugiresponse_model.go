// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorHDFSUGIResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDoctorHDFSUGIResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDoctorHDFSUGIResponse
	GetStatusCode() *int32
	SetBody(v *ListDoctorHDFSUGIResponseBody) *ListDoctorHDFSUGIResponse
	GetBody() *ListDoctorHDFSUGIResponseBody
}

type ListDoctorHDFSUGIResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDoctorHDFSUGIResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDoctorHDFSUGIResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSUGIResponse) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSUGIResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDoctorHDFSUGIResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDoctorHDFSUGIResponse) GetBody() *ListDoctorHDFSUGIResponseBody {
	return s.Body
}

func (s *ListDoctorHDFSUGIResponse) SetHeaders(v map[string]*string) *ListDoctorHDFSUGIResponse {
	s.Headers = v
	return s
}

func (s *ListDoctorHDFSUGIResponse) SetStatusCode(v int32) *ListDoctorHDFSUGIResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDoctorHDFSUGIResponse) SetBody(v *ListDoctorHDFSUGIResponseBody) *ListDoctorHDFSUGIResponse {
	s.Body = v
	return s
}

func (s *ListDoctorHDFSUGIResponse) Validate() error {
	return dara.Validate(s)
}
