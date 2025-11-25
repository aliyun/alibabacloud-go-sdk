// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopKyuubiServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopKyuubiServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopKyuubiServiceResponse
	GetStatusCode() *int32
	SetBody(v *StopKyuubiServiceResponseBody) *StopKyuubiServiceResponse
	GetBody() *StopKyuubiServiceResponseBody
}

type StopKyuubiServiceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopKyuubiServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopKyuubiServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s StopKyuubiServiceResponse) GoString() string {
	return s.String()
}

func (s *StopKyuubiServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopKyuubiServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopKyuubiServiceResponse) GetBody() *StopKyuubiServiceResponseBody {
	return s.Body
}

func (s *StopKyuubiServiceResponse) SetHeaders(v map[string]*string) *StopKyuubiServiceResponse {
	s.Headers = v
	return s
}

func (s *StopKyuubiServiceResponse) SetStatusCode(v int32) *StopKyuubiServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopKyuubiServiceResponse) SetBody(v *StopKyuubiServiceResponseBody) *StopKyuubiServiceResponse {
	s.Body = v
	return s
}

func (s *StopKyuubiServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
