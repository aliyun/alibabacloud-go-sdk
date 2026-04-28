// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *AuthorizeShrinkRequest
	GetClientId() *string
	SetHideConsent(v bool) *AuthorizeShrinkRequest
	GetHideConsent() *bool
	SetLoginType(v string) *AuthorizeShrinkRequest
	GetLoginType() *string
	SetRedirectUri(v string) *AuthorizeShrinkRequest
	GetRedirectUri() *string
	SetResponseType(v string) *AuthorizeShrinkRequest
	GetResponseType() *string
	SetScopeShrink(v string) *AuthorizeShrinkRequest
	GetScopeShrink() *string
	SetState(v string) *AuthorizeShrinkRequest
	GetState() *string
}

type AuthorizeShrinkRequest struct {
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
	ScopeShrink *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The user-defined parameter to return in the callback URL after the requested permissions are granted.
	//
	// example:
	//
	// customdata
	State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s AuthorizeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeShrinkRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeShrinkRequest) GetClientId() *string {
	return s.ClientId
}

func (s *AuthorizeShrinkRequest) GetHideConsent() *bool {
	return s.HideConsent
}

func (s *AuthorizeShrinkRequest) GetLoginType() *string {
	return s.LoginType
}

func (s *AuthorizeShrinkRequest) GetRedirectUri() *string {
	return s.RedirectUri
}

func (s *AuthorizeShrinkRequest) GetResponseType() *string {
	return s.ResponseType
}

func (s *AuthorizeShrinkRequest) GetScopeShrink() *string {
	return s.ScopeShrink
}

func (s *AuthorizeShrinkRequest) GetState() *string {
	return s.State
}

func (s *AuthorizeShrinkRequest) SetClientId(v string) *AuthorizeShrinkRequest {
	s.ClientId = &v
	return s
}

func (s *AuthorizeShrinkRequest) SetHideConsent(v bool) *AuthorizeShrinkRequest {
	s.HideConsent = &v
	return s
}

func (s *AuthorizeShrinkRequest) SetLoginType(v string) *AuthorizeShrinkRequest {
	s.LoginType = &v
	return s
}

func (s *AuthorizeShrinkRequest) SetRedirectUri(v string) *AuthorizeShrinkRequest {
	s.RedirectUri = &v
	return s
}

func (s *AuthorizeShrinkRequest) SetResponseType(v string) *AuthorizeShrinkRequest {
	s.ResponseType = &v
	return s
}

func (s *AuthorizeShrinkRequest) SetScopeShrink(v string) *AuthorizeShrinkRequest {
	s.ScopeShrink = &v
	return s
}

func (s *AuthorizeShrinkRequest) SetState(v string) *AuthorizeShrinkRequest {
	s.State = &v
	return s
}

func (s *AuthorizeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
