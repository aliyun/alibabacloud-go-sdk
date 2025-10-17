// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDoctorApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDoctorApplicationResponse
	GetStatusCode() *int32
	SetBody(v *GetDoctorApplicationResponseBody) *GetDoctorApplicationResponse
	GetBody() *GetDoctorApplicationResponseBody
}

type GetDoctorApplicationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDoctorApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDoctorApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorApplicationResponse) GoString() string {
	return s.String()
}

func (s *GetDoctorApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDoctorApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDoctorApplicationResponse) GetBody() *GetDoctorApplicationResponseBody {
	return s.Body
}

func (s *GetDoctorApplicationResponse) SetHeaders(v map[string]*string) *GetDoctorApplicationResponse {
	s.Headers = v
	return s
}

func (s *GetDoctorApplicationResponse) SetStatusCode(v int32) *GetDoctorApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDoctorApplicationResponse) SetBody(v *GetDoctorApplicationResponseBody) *GetDoctorApplicationResponse {
	s.Body = v
	return s
}

func (s *GetDoctorApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
