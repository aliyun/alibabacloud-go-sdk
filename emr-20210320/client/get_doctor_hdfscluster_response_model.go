// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorHDFSClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDoctorHDFSClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDoctorHDFSClusterResponse
	GetStatusCode() *int32
	SetBody(v *GetDoctorHDFSClusterResponseBody) *GetDoctorHDFSClusterResponse
	GetBody() *GetDoctorHDFSClusterResponseBody
}

type GetDoctorHDFSClusterResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDoctorHDFSClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDoctorHDFSClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHDFSClusterResponse) GoString() string {
	return s.String()
}

func (s *GetDoctorHDFSClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDoctorHDFSClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDoctorHDFSClusterResponse) GetBody() *GetDoctorHDFSClusterResponseBody {
	return s.Body
}

func (s *GetDoctorHDFSClusterResponse) SetHeaders(v map[string]*string) *GetDoctorHDFSClusterResponse {
	s.Headers = v
	return s
}

func (s *GetDoctorHDFSClusterResponse) SetStatusCode(v int32) *GetDoctorHDFSClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDoctorHDFSClusterResponse) SetBody(v *GetDoctorHDFSClusterResponseBody) *GetDoctorHDFSClusterResponse {
	s.Body = v
	return s
}

func (s *GetDoctorHDFSClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
