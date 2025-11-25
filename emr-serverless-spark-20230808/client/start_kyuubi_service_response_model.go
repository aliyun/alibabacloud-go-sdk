// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartKyuubiServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartKyuubiServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartKyuubiServiceResponse
	GetStatusCode() *int32
	SetBody(v *StartKyuubiServiceResponseBody) *StartKyuubiServiceResponse
	GetBody() *StartKyuubiServiceResponseBody
}

type StartKyuubiServiceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartKyuubiServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartKyuubiServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s StartKyuubiServiceResponse) GoString() string {
	return s.String()
}

func (s *StartKyuubiServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartKyuubiServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartKyuubiServiceResponse) GetBody() *StartKyuubiServiceResponseBody {
	return s.Body
}

func (s *StartKyuubiServiceResponse) SetHeaders(v map[string]*string) *StartKyuubiServiceResponse {
	s.Headers = v
	return s
}

func (s *StartKyuubiServiceResponse) SetStatusCode(v int32) *StartKyuubiServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartKyuubiServiceResponse) SetBody(v *StartKyuubiServiceResponseBody) *StartKyuubiServiceResponse {
	s.Body = v
	return s
}

func (s *StartKyuubiServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
