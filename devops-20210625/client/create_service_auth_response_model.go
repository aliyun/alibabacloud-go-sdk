// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceAuthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateServiceAuthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateServiceAuthResponse
	GetStatusCode() *int32
	SetBody(v *CreateServiceAuthResponseBody) *CreateServiceAuthResponse
	GetBody() *CreateServiceAuthResponseBody
}

type CreateServiceAuthResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceAuthResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceAuthResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceAuthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateServiceAuthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateServiceAuthResponse) GetBody() *CreateServiceAuthResponseBody {
	return s.Body
}

func (s *CreateServiceAuthResponse) SetHeaders(v map[string]*string) *CreateServiceAuthResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceAuthResponse) SetStatusCode(v int32) *CreateServiceAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceAuthResponse) SetBody(v *CreateServiceAuthResponseBody) *CreateServiceAuthResponse {
	s.Body = v
	return s
}

func (s *CreateServiceAuthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
