// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKyuubiServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateKyuubiServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateKyuubiServiceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateKyuubiServiceResponseBody) *UpdateKyuubiServiceResponse
	GetBody() *UpdateKyuubiServiceResponseBody
}

type UpdateKyuubiServiceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKyuubiServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKyuubiServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateKyuubiServiceResponse) GoString() string {
	return s.String()
}

func (s *UpdateKyuubiServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateKyuubiServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateKyuubiServiceResponse) GetBody() *UpdateKyuubiServiceResponseBody {
	return s.Body
}

func (s *UpdateKyuubiServiceResponse) SetHeaders(v map[string]*string) *UpdateKyuubiServiceResponse {
	s.Headers = v
	return s
}

func (s *UpdateKyuubiServiceResponse) SetStatusCode(v int32) *UpdateKyuubiServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKyuubiServiceResponse) SetBody(v *UpdateKyuubiServiceResponseBody) *UpdateKyuubiServiceResponse {
	s.Body = v
	return s
}

func (s *UpdateKyuubiServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
