// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyInstanceLDAPAuthServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyInstanceLDAPAuthServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyInstanceLDAPAuthServerResponse
	GetStatusCode() *int32
	SetBody(v *VerifyInstanceLDAPAuthServerResponseBody) *VerifyInstanceLDAPAuthServerResponse
	GetBody() *VerifyInstanceLDAPAuthServerResponseBody
}

type VerifyInstanceLDAPAuthServerResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyInstanceLDAPAuthServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyInstanceLDAPAuthServerResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyInstanceLDAPAuthServerResponse) GoString() string {
	return s.String()
}

func (s *VerifyInstanceLDAPAuthServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyInstanceLDAPAuthServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyInstanceLDAPAuthServerResponse) GetBody() *VerifyInstanceLDAPAuthServerResponseBody {
	return s.Body
}

func (s *VerifyInstanceLDAPAuthServerResponse) SetHeaders(v map[string]*string) *VerifyInstanceLDAPAuthServerResponse {
	s.Headers = v
	return s
}

func (s *VerifyInstanceLDAPAuthServerResponse) SetStatusCode(v int32) *VerifyInstanceLDAPAuthServerResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyInstanceLDAPAuthServerResponse) SetBody(v *VerifyInstanceLDAPAuthServerResponseBody) *VerifyInstanceLDAPAuthServerResponse {
	s.Body = v
	return s
}

func (s *VerifyInstanceLDAPAuthServerResponse) Validate() error {
	return dara.Validate(s)
}
