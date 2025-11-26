// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorHDFSDirectoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDoctorHDFSDirectoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDoctorHDFSDirectoriesResponse
	GetStatusCode() *int32
	SetBody(v *ListDoctorHDFSDirectoriesResponseBody) *ListDoctorHDFSDirectoriesResponse
	GetBody() *ListDoctorHDFSDirectoriesResponseBody
}

type ListDoctorHDFSDirectoriesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDoctorHDFSDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDoctorHDFSDirectoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHDFSDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *ListDoctorHDFSDirectoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDoctorHDFSDirectoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDoctorHDFSDirectoriesResponse) GetBody() *ListDoctorHDFSDirectoriesResponseBody {
	return s.Body
}

func (s *ListDoctorHDFSDirectoriesResponse) SetHeaders(v map[string]*string) *ListDoctorHDFSDirectoriesResponse {
	s.Headers = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponse) SetStatusCode(v int32) *ListDoctorHDFSDirectoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponse) SetBody(v *ListDoctorHDFSDirectoriesResponseBody) *ListDoctorHDFSDirectoriesResponse {
	s.Body = v
	return s
}

func (s *ListDoctorHDFSDirectoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
