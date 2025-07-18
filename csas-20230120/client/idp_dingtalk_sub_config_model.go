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
	SetRedirectUri(v string) *IdpDingtalkSubConfig
	GetRedirectUri() *string
}

type IdpDingtalkSubConfig struct {
	AppKey           *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	AppSecret        *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	CorpId           *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	EventAesKey      *string `json:"EventAesKey,omitempty" xml:"EventAesKey,omitempty"`
	EventLabel       *string `json:"EventLabel,omitempty" xml:"EventLabel,omitempty"`
	EventVerifyToken *string `json:"EventVerifyToken,omitempty" xml:"EventVerifyToken,omitempty"`
	Exclusive        *bool   `json:"Exclusive,omitempty" xml:"Exclusive,omitempty"`
	RedirectUri      *string `json:"RedirectUri,omitempty" xml:"RedirectUri,omitempty"`
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

func (s *IdpDingtalkSubConfig) SetRedirectUri(v string) *IdpDingtalkSubConfig {
	s.RedirectUri = &v
	return s
}

func (s *IdpDingtalkSubConfig) Validate() error {
	return dara.Validate(s)
}
