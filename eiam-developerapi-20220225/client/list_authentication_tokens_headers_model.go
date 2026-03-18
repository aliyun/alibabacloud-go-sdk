// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthenticationTokensHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListAuthenticationTokensHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *ListAuthenticationTokensHeaders
	GetAuthorization() *string
}

type ListAuthenticationTokensHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListAuthenticationTokensHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListAuthenticationTokensHeaders) GoString() string {
	return s.String()
}

func (s *ListAuthenticationTokensHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListAuthenticationTokensHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListAuthenticationTokensHeaders) SetCommonHeaders(v map[string]*string) *ListAuthenticationTokensHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListAuthenticationTokensHeaders) SetAuthorization(v string) *ListAuthenticationTokensHeaders {
	s.Authorization = &v
	return s
}

func (s *ListAuthenticationTokensHeaders) Validate() error {
	return dara.Validate(s)
}
