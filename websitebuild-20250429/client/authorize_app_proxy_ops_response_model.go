// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeAppProxyOpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthorizeAppProxyOpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthorizeAppProxyOpsResponse
	GetStatusCode() *int32
	SetBody(v *AuthorizeAppProxyOpsResponseBody) *AuthorizeAppProxyOpsResponse
	GetBody() *AuthorizeAppProxyOpsResponseBody
}

type AuthorizeAppProxyOpsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeAppProxyOpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeAppProxyOpsResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeAppProxyOpsResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeAppProxyOpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthorizeAppProxyOpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthorizeAppProxyOpsResponse) GetBody() *AuthorizeAppProxyOpsResponseBody {
	return s.Body
}

func (s *AuthorizeAppProxyOpsResponse) SetHeaders(v map[string]*string) *AuthorizeAppProxyOpsResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeAppProxyOpsResponse) SetStatusCode(v int32) *AuthorizeAppProxyOpsResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeAppProxyOpsResponse) SetBody(v *AuthorizeAppProxyOpsResponseBody) *AuthorizeAppProxyOpsResponse {
	s.Body = v
	return s
}

func (s *AuthorizeAppProxyOpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
