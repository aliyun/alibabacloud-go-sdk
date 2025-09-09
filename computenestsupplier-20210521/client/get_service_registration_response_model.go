// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceRegistrationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceRegistrationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceRegistrationResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceRegistrationResponseBody) *GetServiceRegistrationResponse
	GetBody() *GetServiceRegistrationResponseBody
}

type GetServiceRegistrationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceRegistrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceRegistrationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceRegistrationResponse) GoString() string {
	return s.String()
}

func (s *GetServiceRegistrationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceRegistrationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceRegistrationResponse) GetBody() *GetServiceRegistrationResponseBody {
	return s.Body
}

func (s *GetServiceRegistrationResponse) SetHeaders(v map[string]*string) *GetServiceRegistrationResponse {
	s.Headers = v
	return s
}

func (s *GetServiceRegistrationResponse) SetStatusCode(v int32) *GetServiceRegistrationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceRegistrationResponse) SetBody(v *GetServiceRegistrationResponseBody) *GetServiceRegistrationResponse {
	s.Body = v
	return s
}

func (s *GetServiceRegistrationResponse) Validate() error {
	return dara.Validate(s)
}
