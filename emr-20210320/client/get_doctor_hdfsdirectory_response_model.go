// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorHDFSDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDoctorHDFSDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDoctorHDFSDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *GetDoctorHDFSDirectoryResponseBody) *GetDoctorHDFSDirectoryResponse
	GetBody() *GetDoctorHDFSDirectoryResponseBody
}

type GetDoctorHDFSDirectoryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDoctorHDFSDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDoctorHDFSDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSDirectoryResponse) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDoctorHDFSDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDoctorHDFSDirectoryResponse) GetBody() *GetDoctorHDFSDirectoryResponseBody {
	return s.Body
}

func (s *GetDoctorHDFSDirectoryResponse) SetHeaders(v map[string]*string) *GetDoctorHDFSDirectoryResponse {
	s.Headers = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponse) SetStatusCode(v int32) *GetDoctorHDFSDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDoctorHDFSDirectoryResponse) SetBody(v *GetDoctorHDFSDirectoryResponseBody) *GetDoctorHDFSDirectoryResponse {
	s.Body = v
	return s
}

func (s *GetDoctorHDFSDirectoryResponse) Validate() error {
	return dara.Validate(s)
}
