// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeEndpointAclResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthorizeEndpointAclResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthorizeEndpointAclResponse
	GetStatusCode() *int32
	SetBody(v *AuthorizeEndpointAclResponseBody) *AuthorizeEndpointAclResponse
	GetBody() *AuthorizeEndpointAclResponseBody
}

type AuthorizeEndpointAclResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeEndpointAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeEndpointAclResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeEndpointAclResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeEndpointAclResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthorizeEndpointAclResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthorizeEndpointAclResponse) GetBody() *AuthorizeEndpointAclResponseBody {
	return s.Body
}

func (s *AuthorizeEndpointAclResponse) SetHeaders(v map[string]*string) *AuthorizeEndpointAclResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeEndpointAclResponse) SetStatusCode(v int32) *AuthorizeEndpointAclResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeEndpointAclResponse) SetBody(v *AuthorizeEndpointAclResponseBody) *AuthorizeEndpointAclResponse {
	s.Body = v
	return s
}

func (s *AuthorizeEndpointAclResponse) Validate() error {
	return dara.Validate(s)
}
