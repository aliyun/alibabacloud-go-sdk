// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeAuthenticationTokenByConsumerHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *RevokeAuthenticationTokenByConsumerHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *RevokeAuthenticationTokenByConsumerHeaders
	GetAuthorization() *string
}

type RevokeAuthenticationTokenByConsumerHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s RevokeAuthenticationTokenByConsumerHeaders) String() string {
	return dara.Prettify(s)
}

func (s RevokeAuthenticationTokenByConsumerHeaders) GoString() string {
	return s.String()
}

func (s *RevokeAuthenticationTokenByConsumerHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *RevokeAuthenticationTokenByConsumerHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *RevokeAuthenticationTokenByConsumerHeaders) SetCommonHeaders(v map[string]*string) *RevokeAuthenticationTokenByConsumerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RevokeAuthenticationTokenByConsumerHeaders) SetAuthorization(v string) *RevokeAuthenticationTokenByConsumerHeaders {
	s.Authorization = &v
	return s
}

func (s *RevokeAuthenticationTokenByConsumerHeaders) Validate() error {
	return dara.Validate(s)
}
