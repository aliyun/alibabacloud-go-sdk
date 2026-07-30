// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOAuthAuthorizationSessionHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetOAuthAuthorizationSessionHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *GetOAuthAuthorizationSessionHeaders
	GetAuthorization() *string
}

type GetOAuthAuthorizationSessionHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The authentication information. Format: Bearer ${access_token}.
	//
	// > Enter the Access Token issued by IDaaS.
	//
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetOAuthAuthorizationSessionHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetOAuthAuthorizationSessionHeaders) GoString() string {
	return s.String()
}

func (s *GetOAuthAuthorizationSessionHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetOAuthAuthorizationSessionHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetOAuthAuthorizationSessionHeaders) SetCommonHeaders(v map[string]*string) *GetOAuthAuthorizationSessionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOAuthAuthorizationSessionHeaders) SetAuthorization(v string) *GetOAuthAuthorizationSessionHeaders {
	s.Authorization = &v
	return s
}

func (s *GetOAuthAuthorizationSessionHeaders) Validate() error {
	return dara.Validate(s)
}
