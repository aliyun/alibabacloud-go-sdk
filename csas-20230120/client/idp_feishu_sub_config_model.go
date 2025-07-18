// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIdpFeishuSubConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *IdpFeishuSubConfig
	GetAppId() *string
	SetAppSecret(v string) *IdpFeishuSubConfig
	GetAppSecret() *string
	SetCorpId(v string) *IdpFeishuSubConfig
	GetCorpId() *string
	SetEventAesKey(v string) *IdpFeishuSubConfig
	GetEventAesKey() *string
	SetEventLabel(v string) *IdpFeishuSubConfig
	GetEventLabel() *string
	SetEventVerifyToken(v string) *IdpFeishuSubConfig
	GetEventVerifyToken() *string
	SetRedirectUri(v string) *IdpFeishuSubConfig
	GetRedirectUri() *string
}

type IdpFeishuSubConfig struct {
	AppId            *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppSecret        *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	CorpId           *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	EventAesKey      *string `json:"EventAesKey,omitempty" xml:"EventAesKey,omitempty"`
	EventLabel       *string `json:"EventLabel,omitempty" xml:"EventLabel,omitempty"`
	EventVerifyToken *string `json:"EventVerifyToken,omitempty" xml:"EventVerifyToken,omitempty"`
	RedirectUri      *string `json:"RedirectUri,omitempty" xml:"RedirectUri,omitempty"`
}

func (s IdpFeishuSubConfig) String() string {
	return dara.Prettify(s)
}

func (s IdpFeishuSubConfig) GoString() string {
	return s.String()
}

func (s *IdpFeishuSubConfig) GetAppId() *string {
	return s.AppId
}

func (s *IdpFeishuSubConfig) GetAppSecret() *string {
	return s.AppSecret
}

func (s *IdpFeishuSubConfig) GetCorpId() *string {
	return s.CorpId
}

func (s *IdpFeishuSubConfig) GetEventAesKey() *string {
	return s.EventAesKey
}

func (s *IdpFeishuSubConfig) GetEventLabel() *string {
	return s.EventLabel
}

func (s *IdpFeishuSubConfig) GetEventVerifyToken() *string {
	return s.EventVerifyToken
}

func (s *IdpFeishuSubConfig) GetRedirectUri() *string {
	return s.RedirectUri
}

func (s *IdpFeishuSubConfig) SetAppId(v string) *IdpFeishuSubConfig {
	s.AppId = &v
	return s
}

func (s *IdpFeishuSubConfig) SetAppSecret(v string) *IdpFeishuSubConfig {
	s.AppSecret = &v
	return s
}

func (s *IdpFeishuSubConfig) SetCorpId(v string) *IdpFeishuSubConfig {
	s.CorpId = &v
	return s
}

func (s *IdpFeishuSubConfig) SetEventAesKey(v string) *IdpFeishuSubConfig {
	s.EventAesKey = &v
	return s
}

func (s *IdpFeishuSubConfig) SetEventLabel(v string) *IdpFeishuSubConfig {
	s.EventLabel = &v
	return s
}

func (s *IdpFeishuSubConfig) SetEventVerifyToken(v string) *IdpFeishuSubConfig {
	s.EventVerifyToken = &v
	return s
}

func (s *IdpFeishuSubConfig) SetRedirectUri(v string) *IdpFeishuSubConfig {
	s.RedirectUri = &v
	return s
}

func (s *IdpFeishuSubConfig) Validate() error {
	return dara.Validate(s)
}
