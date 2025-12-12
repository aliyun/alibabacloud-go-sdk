// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAppointmentRestartInstanceNodeListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelAppointmentRestartInstanceNodeListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelAppointmentRestartInstanceNodeListResponse
	GetStatusCode() *int32
	SetBody(v *CancelAppointmentRestartInstanceNodeListResponseBody) *CancelAppointmentRestartInstanceNodeListResponse
	GetBody() *CancelAppointmentRestartInstanceNodeListResponseBody
}

type CancelAppointmentRestartInstanceNodeListResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelAppointmentRestartInstanceNodeListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelAppointmentRestartInstanceNodeListResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelAppointmentRestartInstanceNodeListResponse) GoString() string {
	return s.String()
}

func (s *CancelAppointmentRestartInstanceNodeListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelAppointmentRestartInstanceNodeListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelAppointmentRestartInstanceNodeListResponse) GetBody() *CancelAppointmentRestartInstanceNodeListResponseBody {
	return s.Body
}

func (s *CancelAppointmentRestartInstanceNodeListResponse) SetHeaders(v map[string]*string) *CancelAppointmentRestartInstanceNodeListResponse {
	s.Headers = v
	return s
}

func (s *CancelAppointmentRestartInstanceNodeListResponse) SetStatusCode(v int32) *CancelAppointmentRestartInstanceNodeListResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelAppointmentRestartInstanceNodeListResponse) SetBody(v *CancelAppointmentRestartInstanceNodeListResponseBody) *CancelAppointmentRestartInstanceNodeListResponse {
	s.Body = v
	return s
}

func (s *CancelAppointmentRestartInstanceNodeListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
