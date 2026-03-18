// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchOAuthAuthenticationTokenHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FetchOAuthAuthenticationTokenHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *FetchOAuthAuthenticationTokenHeaders
	GetAuthorization() *string
}

type FetchOAuthAuthenticationTokenHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s FetchOAuthAuthenticationTokenHeaders) String() string {
	return dara.Prettify(s)
}

func (s FetchOAuthAuthenticationTokenHeaders) GoString() string {
	return s.String()
}

func (s *FetchOAuthAuthenticationTokenHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FetchOAuthAuthenticationTokenHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *FetchOAuthAuthenticationTokenHeaders) SetCommonHeaders(v map[string]*string) *FetchOAuthAuthenticationTokenHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FetchOAuthAuthenticationTokenHeaders) SetAuthorization(v string) *FetchOAuthAuthenticationTokenHeaders {
	s.Authorization = &v
	return s
}

func (s *FetchOAuthAuthenticationTokenHeaders) Validate() error {
	return dara.Validate(s)
}
