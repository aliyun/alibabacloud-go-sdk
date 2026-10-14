// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIdpWeixin2SubConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *IdpWeixin2SubConfig
	GetAgentId() *string
	SetAppSchema(v string) *IdpWeixin2SubConfig
	GetAppSchema() *string
	SetAppSecret(v string) *IdpWeixin2SubConfig
	GetAppSecret() *string
	SetCorpId(v string) *IdpWeixin2SubConfig
	GetCorpId() *string
	SetEventAesKey(v string) *IdpWeixin2SubConfig
	GetEventAesKey() *string
	SetEventLabel(v string) *IdpWeixin2SubConfig
	GetEventLabel() *string
	SetEventVerifyToken(v string) *IdpWeixin2SubConfig
	GetEventVerifyToken() *string
	SetRedirectUri(v string) *IdpWeixin2SubConfig
	GetRedirectUri() *string
}

type IdpWeixin2SubConfig struct {
	// example:
	//
	// 1000002
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// wwautha1b2c3****
	AppSchema *string `json:"AppSchema,omitempty" xml:"AppSchema,omitempty"`
	// example:
	//
	// s1e2c3r4e5****
	AppSecret *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	// example:
	//
	// wwa1b2c3d4****
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	// example:
	//
	// a1e2s3k4e5****
	EventAesKey *string `json:"EventAesKey,omitempty" xml:"EventAesKey,omitempty"`
	// example:
	//
	// https://auth.example.com/callback/wework/a1b2****
	EventLabel *string `json:"EventLabel,omitempty" xml:"EventLabel,omitempty"`
	// example:
	//
	// v1t2o3k4e5****
	EventVerifyToken *string `json:"EventVerifyToken,omitempty" xml:"EventVerifyToken,omitempty"`
	// example:
	//
	// https://auth.example.com/callback
	RedirectUri *string `json:"RedirectUri,omitempty" xml:"RedirectUri,omitempty"`
}

func (s IdpWeixin2SubConfig) String() string {
	return dara.Prettify(s)
}

func (s IdpWeixin2SubConfig) GoString() string {
	return s.String()
}

func (s *IdpWeixin2SubConfig) GetAgentId() *string {
	return s.AgentId
}

func (s *IdpWeixin2SubConfig) GetAppSchema() *string {
	return s.AppSchema
}

func (s *IdpWeixin2SubConfig) GetAppSecret() *string {
	return s.AppSecret
}

func (s *IdpWeixin2SubConfig) GetCorpId() *string {
	return s.CorpId
}

func (s *IdpWeixin2SubConfig) GetEventAesKey() *string {
	return s.EventAesKey
}

func (s *IdpWeixin2SubConfig) GetEventLabel() *string {
	return s.EventLabel
}

func (s *IdpWeixin2SubConfig) GetEventVerifyToken() *string {
	return s.EventVerifyToken
}

func (s *IdpWeixin2SubConfig) GetRedirectUri() *string {
	return s.RedirectUri
}

func (s *IdpWeixin2SubConfig) SetAgentId(v string) *IdpWeixin2SubConfig {
	s.AgentId = &v
	return s
}

func (s *IdpWeixin2SubConfig) SetAppSchema(v string) *IdpWeixin2SubConfig {
	s.AppSchema = &v
	return s
}

func (s *IdpWeixin2SubConfig) SetAppSecret(v string) *IdpWeixin2SubConfig {
	s.AppSecret = &v
	return s
}

func (s *IdpWeixin2SubConfig) SetCorpId(v string) *IdpWeixin2SubConfig {
	s.CorpId = &v
	return s
}

func (s *IdpWeixin2SubConfig) SetEventAesKey(v string) *IdpWeixin2SubConfig {
	s.EventAesKey = &v
	return s
}

func (s *IdpWeixin2SubConfig) SetEventLabel(v string) *IdpWeixin2SubConfig {
	s.EventLabel = &v
	return s
}

func (s *IdpWeixin2SubConfig) SetEventVerifyToken(v string) *IdpWeixin2SubConfig {
	s.EventVerifyToken = &v
	return s
}

func (s *IdpWeixin2SubConfig) SetRedirectUri(v string) *IdpWeixin2SubConfig {
	s.RedirectUri = &v
	return s
}

func (s *IdpWeixin2SubConfig) Validate() error {
	return dara.Validate(s)
}
