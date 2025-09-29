// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialGetResultIntlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CredentialGetResultIntlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CredentialGetResultIntlResponse
	GetStatusCode() *int32
	SetBody(v *CredentialGetResultIntlResponseBody) *CredentialGetResultIntlResponse
	GetBody() *CredentialGetResultIntlResponseBody
}

type CredentialGetResultIntlResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CredentialGetResultIntlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CredentialGetResultIntlResponse) String() string {
	return dara.Prettify(s)
}

func (s CredentialGetResultIntlResponse) GoString() string {
	return s.String()
}

func (s *CredentialGetResultIntlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CredentialGetResultIntlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CredentialGetResultIntlResponse) GetBody() *CredentialGetResultIntlResponseBody {
	return s.Body
}

func (s *CredentialGetResultIntlResponse) SetHeaders(v map[string]*string) *CredentialGetResultIntlResponse {
	s.Headers = v
	return s
}

func (s *CredentialGetResultIntlResponse) SetStatusCode(v int32) *CredentialGetResultIntlResponse {
	s.StatusCode = &v
	return s
}

func (s *CredentialGetResultIntlResponse) SetBody(v *CredentialGetResultIntlResponseBody) *CredentialGetResultIntlResponse {
	s.Body = v
	return s
}

func (s *CredentialGetResultIntlResponse) Validate() error {
	return dara.Validate(s)
}
