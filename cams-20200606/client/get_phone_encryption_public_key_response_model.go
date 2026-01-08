// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhoneEncryptionPublicKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPhoneEncryptionPublicKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPhoneEncryptionPublicKeyResponse
	GetStatusCode() *int32
	SetBody(v *GetPhoneEncryptionPublicKeyResponseBody) *GetPhoneEncryptionPublicKeyResponse
	GetBody() *GetPhoneEncryptionPublicKeyResponseBody
}

type GetPhoneEncryptionPublicKeyResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPhoneEncryptionPublicKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPhoneEncryptionPublicKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPhoneEncryptionPublicKeyResponse) GoString() string {
	return s.String()
}

func (s *GetPhoneEncryptionPublicKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPhoneEncryptionPublicKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPhoneEncryptionPublicKeyResponse) GetBody() *GetPhoneEncryptionPublicKeyResponseBody {
	return s.Body
}

func (s *GetPhoneEncryptionPublicKeyResponse) SetHeaders(v map[string]*string) *GetPhoneEncryptionPublicKeyResponse {
	s.Headers = v
	return s
}

func (s *GetPhoneEncryptionPublicKeyResponse) SetStatusCode(v int32) *GetPhoneEncryptionPublicKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPhoneEncryptionPublicKeyResponse) SetBody(v *GetPhoneEncryptionPublicKeyResponseBody) *GetPhoneEncryptionPublicKeyResponse {
	s.Body = v
	return s
}

func (s *GetPhoneEncryptionPublicKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
