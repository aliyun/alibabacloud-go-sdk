// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeBackupEncryptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthorizeBackupEncryptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthorizeBackupEncryptionResponse
	GetStatusCode() *int32
	SetBody(v *AuthorizeBackupEncryptionResponseBody) *AuthorizeBackupEncryptionResponse
	GetBody() *AuthorizeBackupEncryptionResponseBody
}

type AuthorizeBackupEncryptionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeBackupEncryptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeBackupEncryptionResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeBackupEncryptionResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeBackupEncryptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthorizeBackupEncryptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthorizeBackupEncryptionResponse) GetBody() *AuthorizeBackupEncryptionResponseBody {
	return s.Body
}

func (s *AuthorizeBackupEncryptionResponse) SetHeaders(v map[string]*string) *AuthorizeBackupEncryptionResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeBackupEncryptionResponse) SetStatusCode(v int32) *AuthorizeBackupEncryptionResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeBackupEncryptionResponse) SetBody(v *AuthorizeBackupEncryptionResponseBody) *AuthorizeBackupEncryptionResponse {
	s.Body = v
	return s
}

func (s *AuthorizeBackupEncryptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
