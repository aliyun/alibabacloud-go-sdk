// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOAuthAuthorizationSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSessionUri(v string) *GetOAuthAuthorizationSessionRequest
	GetSessionUri() *string
}

type GetOAuthAuthorizationSessionRequest struct {
	// The authorization session URI.
	//
	// > Returned by the FetchOAuthAuthenticationToken call.
	//
	// This parameter is required.
	//
	// example:
	//
	// urn:ietf:params:oauth:request_uri:atpoas_01l6ljnvrpc5niakl3gj3amxxxxxx
	SessionUri *string `json:"sessionUri,omitempty" xml:"sessionUri,omitempty"`
}

func (s GetOAuthAuthorizationSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOAuthAuthorizationSessionRequest) GoString() string {
	return s.String()
}

func (s *GetOAuthAuthorizationSessionRequest) GetSessionUri() *string {
	return s.SessionUri
}

func (s *GetOAuthAuthorizationSessionRequest) SetSessionUri(v string) *GetOAuthAuthorizationSessionRequest {
	s.SessionUri = &v
	return s
}

func (s *GetOAuthAuthorizationSessionRequest) Validate() error {
	return dara.Validate(s)
}
