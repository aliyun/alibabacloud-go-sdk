// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorHiveDatabaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDoctorHiveDatabaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDoctorHiveDatabaseResponse
	GetStatusCode() *int32
	SetBody(v *GetDoctorHiveDatabaseResponseBody) *GetDoctorHiveDatabaseResponse
	GetBody() *GetDoctorHiveDatabaseResponseBody
}

type GetDoctorHiveDatabaseResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDoctorHiveDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDoctorHiveDatabaseResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseResponse) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDoctorHiveDatabaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDoctorHiveDatabaseResponse) GetBody() *GetDoctorHiveDatabaseResponseBody {
	return s.Body
}

func (s *GetDoctorHiveDatabaseResponse) SetHeaders(v map[string]*string) *GetDoctorHiveDatabaseResponse {
	s.Headers = v
	return s
}

func (s *GetDoctorHiveDatabaseResponse) SetStatusCode(v int32) *GetDoctorHiveDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDoctorHiveDatabaseResponse) SetBody(v *GetDoctorHiveDatabaseResponseBody) *GetDoctorHiveDatabaseResponse {
	s.Body = v
	return s
}

func (s *GetDoctorHiveDatabaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
