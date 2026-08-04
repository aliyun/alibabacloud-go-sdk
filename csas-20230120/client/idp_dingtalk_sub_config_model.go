// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIdpDingtalkSubConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v string) *IdpDingtalkSubConfig
	GetAppKey() *string
	SetAppSecret(v string) *IdpDingtalkSubConfig
	GetAppSecret() *string
	SetCorpId(v string) *IdpDingtalkSubConfig
	GetCorpId() *string
	SetEventAesKey(v string) *IdpDingtalkSubConfig
	GetEventAesKey() *string
	SetEventLabel(v string) *IdpDingtalkSubConfig
	GetEventLabel() *string
	SetEventVerifyToken(v string) *IdpDingtalkSubConfig
	GetEventVerifyToken() *string
	SetExclusive(v bool) *IdpDingtalkSubConfig
	GetExclusive() *bool
	SetOauth(v bool) *IdpDingtalkSubConfig
	GetOauth() *bool
	SetRedirectUri(v string) *IdpDingtalkSubConfig
	GetRedirectUri() *string
}

type IdpDingtalkSubConfig struct {
	// Your application\\"s unique identifier. You can get this identifier from the DingTalk Open Platform.
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// Your application\\"s secret key. You can get this key from the DingTalk Open Platform.
	AppSecret *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	// Your enterprise\\"s unique ID in DingTalk.
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	// The AES key used to decrypt the content of event callbacks. This ensures the confidentiality of the event data.
	EventAesKey *string `json:"EventAesKey,omitempty" xml:"EventAesKey,omitempty"`
	// A custom label for event subscriptions. This field is reserved for future use.
	EventLabel *string `json:"EventLabel,omitempty" xml:"EventLabel,omitempty"`
	// The token used to verify the authenticity of event callback requests from DingTalk.
	EventVerifyToken *string `json:"EventVerifyToken,omitempty" xml:"EventVerifyToken,omitempty"`
	// Specifies whether this identity provider is the exclusive login method. If set to `true`, other login methods are disabled.
	Exclusive *bool `json:"Exclusive,omitempty" xml:"Exclusive,omitempty"`
	// Specifies whether to enable the OAuth authentication flow.
	Oauth *bool `json:"Oauth,omitempty" xml:"Oauth,omitempty"`
	// The URL where the user is redirected after successful authorization. You must register this URL on the DingTalk Open Platform.
	RedirectUri *string `json:"RedirectUri,omitempty" xml:"RedirectUri,omitempty"`
}

func (s IdpDingtalkSubConfig) String() string {
	return dara.Prettify(s)
}

func (s IdpDingtalkSubConfig) GoString() string {
	return s.String()
}

func (s *IdpDingtalkSubConfig) GetAppKey() *string {
	return s.AppKey
}

func (s *IdpDingtalkSubConfig) GetAppSecret() *string {
	return s.AppSecret
}

func (s *IdpDingtalkSubConfig) GetCorpId() *string {
	return s.CorpId
}

func (s *IdpDingtalkSubConfig) GetEventAesKey() *string {
	return s.EventAesKey
}

func (s *IdpDingtalkSubConfig) GetEventLabel() *string {
	return s.EventLabel
}

func (s *IdpDingtalkSubConfig) GetEventVerifyToken() *string {
	return s.EventVerifyToken
}

func (s *IdpDingtalkSubConfig) GetExclusive() *bool {
	return s.Exclusive
}

func (s *IdpDingtalkSubConfig) GetOauth() *bool {
	return s.Oauth
}

func (s *IdpDingtalkSubConfig) GetRedirectUri() *string {
	return s.RedirectUri
}

func (s *IdpDingtalkSubConfig) SetAppKey(v string) *IdpDingtalkSubConfig {
	s.AppKey = &v
	return s
}

func (s *IdpDingtalkSubConfig) SetAppSecret(v string) *IdpDingtalkSubConfig {
	s.AppSecret = &v
	return s
}

func (s *IdpDingtalkSubConfig) SetCorpId(v string) *IdpDingtalkSubConfig {
	s.CorpId = &v
	return s
}

func (s *IdpDingtalkSubConfig) SetEventAesKey(v string) *IdpDingtalkSubConfig {
	s.EventAesKey = &v
	return s
}

func (s *IdpDingtalkSubConfig) SetEventLabel(v string) *IdpDingtalkSubConfig {
	s.EventLabel = &v
	return s
}

func (s *IdpDingtalkSubConfig) SetEventVerifyToken(v string) *IdpDingtalkSubConfig {
	s.EventVerifyToken = &v
	return s
}

func (s *IdpDingtalkSubConfig) SetExclusive(v bool) *IdpDingtalkSubConfig {
	s.Exclusive = &v
	return s
}

func (s *IdpDingtalkSubConfig) SetOauth(v bool) *IdpDingtalkSubConfig {
	s.Oauth = &v
	return s
}

func (s *IdpDingtalkSubConfig) SetRedirectUri(v string) *IdpDingtalkSubConfig {
	s.RedirectUri = &v
	return s
}

func (s *IdpDingtalkSubConfig) Validate() error {
	return dara.Validate(s)
}
