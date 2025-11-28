// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeMemberBizChainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthorizeMemberBizChainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthorizeMemberBizChainResponse
	GetStatusCode() *int32
	SetBody(v *AuthorizeMemberBizChainResponseBody) *AuthorizeMemberBizChainResponse
	GetBody() *AuthorizeMemberBizChainResponseBody
}

type AuthorizeMemberBizChainResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeMemberBizChainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeMemberBizChainResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeMemberBizChainResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeMemberBizChainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthorizeMemberBizChainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthorizeMemberBizChainResponse) GetBody() *AuthorizeMemberBizChainResponseBody {
	return s.Body
}

func (s *AuthorizeMemberBizChainResponse) SetHeaders(v map[string]*string) *AuthorizeMemberBizChainResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeMemberBizChainResponse) SetStatusCode(v int32) *AuthorizeMemberBizChainResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeMemberBizChainResponse) SetBody(v *AuthorizeMemberBizChainResponseBody) *AuthorizeMemberBizChainResponse {
	s.Body = v
	return s
}

func (s *AuthorizeMemberBizChainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
