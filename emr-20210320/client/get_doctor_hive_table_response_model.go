// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorHiveTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDoctorHiveTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDoctorHiveTableResponse
	GetStatusCode() *int32
	SetBody(v *GetDoctorHiveTableResponseBody) *GetDoctorHiveTableResponse
	GetBody() *GetDoctorHiveTableResponseBody
}

type GetDoctorHiveTableResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDoctorHiveTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDoctorHiveTableResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableResponse) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDoctorHiveTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDoctorHiveTableResponse) GetBody() *GetDoctorHiveTableResponseBody {
	return s.Body
}

func (s *GetDoctorHiveTableResponse) SetHeaders(v map[string]*string) *GetDoctorHiveTableResponse {
	s.Headers = v
	return s
}

func (s *GetDoctorHiveTableResponse) SetStatusCode(v int32) *GetDoctorHiveTableResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDoctorHiveTableResponse) SetBody(v *GetDoctorHiveTableResponseBody) *GetDoctorHiveTableResponse {
	s.Body = v
	return s
}

func (s *GetDoctorHiveTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
