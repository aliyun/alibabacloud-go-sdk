// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelServiceRegistrationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelServiceRegistrationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelServiceRegistrationResponse
	GetStatusCode() *int32
	SetBody(v *CancelServiceRegistrationResponseBody) *CancelServiceRegistrationResponse
	GetBody() *CancelServiceRegistrationResponseBody
}

type CancelServiceRegistrationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelServiceRegistrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelServiceRegistrationResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelServiceRegistrationResponse) GoString() string {
	return s.String()
}

func (s *CancelServiceRegistrationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelServiceRegistrationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelServiceRegistrationResponse) GetBody() *CancelServiceRegistrationResponseBody {
	return s.Body
}

func (s *CancelServiceRegistrationResponse) SetHeaders(v map[string]*string) *CancelServiceRegistrationResponse {
	s.Headers = v
	return s
}

func (s *CancelServiceRegistrationResponse) SetStatusCode(v int32) *CancelServiceRegistrationResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelServiceRegistrationResponse) SetBody(v *CancelServiceRegistrationResponseBody) *CancelServiceRegistrationResponse {
	s.Body = v
	return s
}

func (s *CancelServiceRegistrationResponse) Validate() error {
	return dara.Validate(s)
}
