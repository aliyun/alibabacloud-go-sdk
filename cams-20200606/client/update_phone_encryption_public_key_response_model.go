// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePhoneEncryptionPublicKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePhoneEncryptionPublicKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePhoneEncryptionPublicKeyResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePhoneEncryptionPublicKeyResponseBody) *UpdatePhoneEncryptionPublicKeyResponse
	GetBody() *UpdatePhoneEncryptionPublicKeyResponseBody
}

type UpdatePhoneEncryptionPublicKeyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePhoneEncryptionPublicKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePhoneEncryptionPublicKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePhoneEncryptionPublicKeyResponse) GoString() string {
	return s.String()
}

func (s *UpdatePhoneEncryptionPublicKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePhoneEncryptionPublicKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePhoneEncryptionPublicKeyResponse) GetBody() *UpdatePhoneEncryptionPublicKeyResponseBody {
	return s.Body
}

func (s *UpdatePhoneEncryptionPublicKeyResponse) SetHeaders(v map[string]*string) *UpdatePhoneEncryptionPublicKeyResponse {
	s.Headers = v
	return s
}

func (s *UpdatePhoneEncryptionPublicKeyResponse) SetStatusCode(v int32) *UpdatePhoneEncryptionPublicKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePhoneEncryptionPublicKeyResponse) SetBody(v *UpdatePhoneEncryptionPublicKeyResponseBody) *UpdatePhoneEncryptionPublicKeyResponse {
	s.Body = v
	return s
}

func (s *UpdatePhoneEncryptionPublicKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
