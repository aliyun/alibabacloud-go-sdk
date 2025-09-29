// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialSubmitIntlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CredentialSubmitIntlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CredentialSubmitIntlResponse
	GetStatusCode() *int32
	SetBody(v *CredentialSubmitIntlResponseBody) *CredentialSubmitIntlResponse
	GetBody() *CredentialSubmitIntlResponseBody
}

type CredentialSubmitIntlResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CredentialSubmitIntlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CredentialSubmitIntlResponse) String() string {
	return dara.Prettify(s)
}

func (s CredentialSubmitIntlResponse) GoString() string {
	return s.String()
}

func (s *CredentialSubmitIntlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CredentialSubmitIntlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CredentialSubmitIntlResponse) GetBody() *CredentialSubmitIntlResponseBody {
	return s.Body
}

func (s *CredentialSubmitIntlResponse) SetHeaders(v map[string]*string) *CredentialSubmitIntlResponse {
	s.Headers = v
	return s
}

func (s *CredentialSubmitIntlResponse) SetStatusCode(v int32) *CredentialSubmitIntlResponse {
	s.StatusCode = &v
	return s
}

func (s *CredentialSubmitIntlResponse) SetBody(v *CredentialSubmitIntlResponseBody) *CredentialSubmitIntlResponse {
	s.Body = v
	return s
}

func (s *CredentialSubmitIntlResponse) Validate() error {
	return dara.Validate(s)
}
