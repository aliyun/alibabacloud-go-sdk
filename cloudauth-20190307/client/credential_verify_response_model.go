// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialVerifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CredentialVerifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CredentialVerifyResponse
	GetStatusCode() *int32
	SetBody(v *CredentialVerifyResponseBody) *CredentialVerifyResponse
	GetBody() *CredentialVerifyResponseBody
}

type CredentialVerifyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CredentialVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CredentialVerifyResponse) String() string {
	return dara.Prettify(s)
}

func (s CredentialVerifyResponse) GoString() string {
	return s.String()
}

func (s *CredentialVerifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CredentialVerifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CredentialVerifyResponse) GetBody() *CredentialVerifyResponseBody {
	return s.Body
}

func (s *CredentialVerifyResponse) SetHeaders(v map[string]*string) *CredentialVerifyResponse {
	s.Headers = v
	return s
}

func (s *CredentialVerifyResponse) SetStatusCode(v int32) *CredentialVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *CredentialVerifyResponse) SetBody(v *CredentialVerifyResponseBody) *CredentialVerifyResponse {
	s.Body = v
	return s
}

func (s *CredentialVerifyResponse) Validate() error {
	return dara.Validate(s)
}
