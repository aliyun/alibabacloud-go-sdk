// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDoctorJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDoctorJobResponse
	GetStatusCode() *int32
	SetBody(v *GetDoctorJobResponseBody) *GetDoctorJobResponse
	GetBody() *GetDoctorJobResponseBody
}

type GetDoctorJobResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDoctorJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDoctorJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorJobResponse) GoString() string {
	return s.String()
}

func (s *GetDoctorJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDoctorJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDoctorJobResponse) GetBody() *GetDoctorJobResponseBody {
	return s.Body
}

func (s *GetDoctorJobResponse) SetHeaders(v map[string]*string) *GetDoctorJobResponse {
	s.Headers = v
	return s
}

func (s *GetDoctorJobResponse) SetStatusCode(v int32) *GetDoctorJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDoctorJobResponse) SetBody(v *GetDoctorJobResponseBody) *GetDoctorJobResponse {
	s.Body = v
	return s
}

func (s *GetDoctorJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
