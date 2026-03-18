// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeAuthenticationTokenHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *RevokeAuthenticationTokenHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *RevokeAuthenticationTokenHeaders
	GetAuthorization() *string
}

type RevokeAuthenticationTokenHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s RevokeAuthenticationTokenHeaders) String() string {
	return dara.Prettify(s)
}

func (s RevokeAuthenticationTokenHeaders) GoString() string {
	return s.String()
}

func (s *RevokeAuthenticationTokenHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *RevokeAuthenticationTokenHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *RevokeAuthenticationTokenHeaders) SetCommonHeaders(v map[string]*string) *RevokeAuthenticationTokenHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RevokeAuthenticationTokenHeaders) SetAuthorization(v string) *RevokeAuthenticationTokenHeaders {
	s.Authorization = &v
	return s
}

func (s *RevokeAuthenticationTokenHeaders) Validate() error {
	return dara.Validate(s)
}
