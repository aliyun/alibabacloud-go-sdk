// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialVerifyIntlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CredentialVerifyIntlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CredentialVerifyIntlResponse
	GetStatusCode() *int32
	SetBody(v *CredentialVerifyIntlResponseBody) *CredentialVerifyIntlResponse
	GetBody() *CredentialVerifyIntlResponseBody
}

type CredentialVerifyIntlResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CredentialVerifyIntlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CredentialVerifyIntlResponse) String() string {
	return dara.Prettify(s)
}

func (s CredentialVerifyIntlResponse) GoString() string {
	return s.String()
}

func (s *CredentialVerifyIntlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CredentialVerifyIntlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CredentialVerifyIntlResponse) GetBody() *CredentialVerifyIntlResponseBody {
	return s.Body
}

func (s *CredentialVerifyIntlResponse) SetHeaders(v map[string]*string) *CredentialVerifyIntlResponse {
	s.Headers = v
	return s
}

func (s *CredentialVerifyIntlResponse) SetStatusCode(v int32) *CredentialVerifyIntlResponse {
	s.StatusCode = &v
	return s
}

func (s *CredentialVerifyIntlResponse) SetBody(v *CredentialVerifyIntlResponseBody) *CredentialVerifyIntlResponse {
	s.Body = v
	return s
}

func (s *CredentialVerifyIntlResponse) Validate() error {
	return dara.Validate(s)
}
