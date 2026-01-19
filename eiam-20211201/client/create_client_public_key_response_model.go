// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClientPublicKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateClientPublicKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateClientPublicKeyResponse
	GetStatusCode() *int32
	SetBody(v *CreateClientPublicKeyResponseBody) *CreateClientPublicKeyResponse
	GetBody() *CreateClientPublicKeyResponseBody
}

type CreateClientPublicKeyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateClientPublicKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateClientPublicKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateClientPublicKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateClientPublicKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateClientPublicKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateClientPublicKeyResponse) GetBody() *CreateClientPublicKeyResponseBody {
	return s.Body
}

func (s *CreateClientPublicKeyResponse) SetHeaders(v map[string]*string) *CreateClientPublicKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateClientPublicKeyResponse) SetStatusCode(v int32) *CreateClientPublicKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClientPublicKeyResponse) SetBody(v *CreateClientPublicKeyResponseBody) *CreateClientPublicKeyResponse {
	s.Body = v
	return s
}

func (s *CreateClientPublicKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
