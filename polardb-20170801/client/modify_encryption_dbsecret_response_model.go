// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEncryptionDBSecretResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyEncryptionDBSecretResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyEncryptionDBSecretResponse
	GetStatusCode() *int32
	SetBody(v *ModifyEncryptionDBSecretResponseBody) *ModifyEncryptionDBSecretResponse
	GetBody() *ModifyEncryptionDBSecretResponseBody
}

type ModifyEncryptionDBSecretResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyEncryptionDBSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyEncryptionDBSecretResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyEncryptionDBSecretResponse) GoString() string {
	return s.String()
}

func (s *ModifyEncryptionDBSecretResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyEncryptionDBSecretResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyEncryptionDBSecretResponse) GetBody() *ModifyEncryptionDBSecretResponseBody {
	return s.Body
}

func (s *ModifyEncryptionDBSecretResponse) SetHeaders(v map[string]*string) *ModifyEncryptionDBSecretResponse {
	s.Headers = v
	return s
}

func (s *ModifyEncryptionDBSecretResponse) SetStatusCode(v int32) *ModifyEncryptionDBSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyEncryptionDBSecretResponse) SetBody(v *ModifyEncryptionDBSecretResponseBody) *ModifyEncryptionDBSecretResponse {
	s.Body = v
	return s
}

func (s *ModifyEncryptionDBSecretResponse) Validate() error {
	return dara.Validate(s)
}
