// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *AuthorizeRequest
	GetClientId() *string
	SetHideConsent(v bool) *AuthorizeRequest
	GetHideConsent() *bool
	SetLoginType(v string) *AuthorizeRequest
	GetLoginType() *string
	SetRedirectUri(v string) *AuthorizeRequest
	GetRedirectUri() *string
	SetResponseType(v string) *AuthorizeRequest
	GetResponseType() *string
	SetScope(v []*string) *AuthorizeRequest
	GetScope() []*string
	SetState(v string) *AuthorizeRequest
	GetState() *string
}

type AuthorizeRequest struct {
	// The application ID returned when the application was created.
	//
	// This parameter is required.
	//
	// example:
	//
	// 47eUHhrzgWBvlLWj
	ClientId *string `json:"client_id,omitempty" xml:"client_id,omitempty"`
	// Specifies whether to hide the consent page.
	//
	// example:
	//
	// true
	HideConsent *bool `json:"hide_consent,omitempty" xml:"hide_consent,omitempty"`
	// The authentication method. Valid values:
	//
	// 	- default: all logon methods that are integrated on the default logon page provided by Drive and Photo Service.
	//
	// 	- ding: logs on by scanning a DingTalk QR code.
	//
	// 	- ding_sns: logs on by entering a DingTalk account and its password.
	//
	// 	- ram: logs on as an Alibaba Cloud Resource Access Management (RAM) user.
	//
	// 	- wechat: logs on by scanning a WeCom QR code.
	//
	// 	- wechat_app: logs on without authentication in WeCom.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	LoginType *string `json:"login_type,omitempty" xml:"login_type,omitempty"`
	// The callback URL specified when the application was created.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://www.aliyunpds.com/sign/callback
	RedirectUri *string `json:"redirect_uri,omitempty" xml:"redirect_uri,omitempty"`
	// The format in which to return the response. Set the value to code.
	//
	// This parameter is required.
	//
	// example:
	//
	// code
	ResponseType *string `json:"response_type,omitempty" xml:"response_type,omitempty"`
	// The requested permissions. By default, all permissions are requested.
	Scope []*string `json:"scope,omitempty" xml:"scope,omitempty" type:"Repeated"`
	// The user-defined parameter to return in the callback URL after the requested permissions are granted.
	//
	// example:
	//
	// customdata
	State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s AuthorizeRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeRequest) GetClientId() *string {
	return s.ClientId
}

func (s *AuthorizeRequest) GetHideConsent() *bool {
	return s.HideConsent
}

func (s *AuthorizeRequest) GetLoginType() *string {
	return s.LoginType
}

func (s *AuthorizeRequest) GetRedirectUri() *string {
	return s.RedirectUri
}

func (s *AuthorizeRequest) GetResponseType() *string {
	return s.ResponseType
}

func (s *AuthorizeRequest) GetScope() []*string {
	return s.Scope
}

func (s *AuthorizeRequest) GetState() *string {
	return s.State
}

func (s *AuthorizeRequest) SetClientId(v string) *AuthorizeRequest {
	s.ClientId = &v
	return s
}

func (s *AuthorizeRequest) SetHideConsent(v bool) *AuthorizeRequest {
	s.HideConsent = &v
	return s
}

func (s *AuthorizeRequest) SetLoginType(v string) *AuthorizeRequest {
	s.LoginType = &v
	return s
}

func (s *AuthorizeRequest) SetRedirectUri(v string) *AuthorizeRequest {
	s.RedirectUri = &v
	return s
}

func (s *AuthorizeRequest) SetResponseType(v string) *AuthorizeRequest {
	s.ResponseType = &v
	return s
}

func (s *AuthorizeRequest) SetScope(v []*string) *AuthorizeRequest {
	s.Scope = v
	return s
}

func (s *AuthorizeRequest) SetState(v string) *AuthorizeRequest {
	s.State = &v
	return s
}

func (s *AuthorizeRequest) Validate() error {
	return dara.Validate(s)
}
