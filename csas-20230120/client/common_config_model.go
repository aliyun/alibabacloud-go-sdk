// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommonConfig interface {
	dara.Model
	String() string
	GoString() string
	SetIdp(v *CommonConfigIdp) *CommonConfig
	GetIdp() *CommonConfigIdp
}

type CommonConfig struct {
	Idp *CommonConfigIdp `json:"Idp,omitempty" xml:"Idp,omitempty" type:"Struct"`
}

func (s CommonConfig) String() string {
	return dara.Prettify(s)
}

func (s CommonConfig) GoString() string {
	return s.String()
}

func (s *CommonConfig) GetIdp() *CommonConfigIdp {
	return s.Idp
}

func (s *CommonConfig) SetIdp(v *CommonConfigIdp) *CommonConfig {
	s.Idp = v
	return s
}

func (s *CommonConfig) Validate() error {
	return dara.Validate(s)
}

type CommonConfigIdp struct {
	ApOidcCallbackUrl *string                  `json:"ApOidcCallbackUrl,omitempty" xml:"ApOidcCallbackUrl,omitempty"`
	Dingtalk          *CommonConfigIdpDingtalk `json:"Dingtalk,omitempty" xml:"Dingtalk,omitempty" type:"Struct"`
	Feishu            *CommonConfigIdpFeishu   `json:"Feishu,omitempty" xml:"Feishu,omitempty" type:"Struct"`
	Idaas2            *CommonConfigIdpIdaas2   `json:"Idaas2,omitempty" xml:"Idaas2,omitempty" type:"Struct"`
	Saml              *CommonConfigIdpSaml     `json:"Saml,omitempty" xml:"Saml,omitempty" type:"Struct"`
}

func (s CommonConfigIdp) String() string {
	return dara.Prettify(s)
}

func (s CommonConfigIdp) GoString() string {
	return s.String()
}

func (s *CommonConfigIdp) GetApOidcCallbackUrl() *string {
	return s.ApOidcCallbackUrl
}

func (s *CommonConfigIdp) GetDingtalk() *CommonConfigIdpDingtalk {
	return s.Dingtalk
}

func (s *CommonConfigIdp) GetFeishu() *CommonConfigIdpFeishu {
	return s.Feishu
}

func (s *CommonConfigIdp) GetIdaas2() *CommonConfigIdpIdaas2 {
	return s.Idaas2
}

func (s *CommonConfigIdp) GetSaml() *CommonConfigIdpSaml {
	return s.Saml
}

func (s *CommonConfigIdp) SetApOidcCallbackUrl(v string) *CommonConfigIdp {
	s.ApOidcCallbackUrl = &v
	return s
}

func (s *CommonConfigIdp) SetDingtalk(v *CommonConfigIdpDingtalk) *CommonConfigIdp {
	s.Dingtalk = v
	return s
}

func (s *CommonConfigIdp) SetFeishu(v *CommonConfigIdpFeishu) *CommonConfigIdp {
	s.Feishu = v
	return s
}

func (s *CommonConfigIdp) SetIdaas2(v *CommonConfigIdpIdaas2) *CommonConfigIdp {
	s.Idaas2 = v
	return s
}

func (s *CommonConfigIdp) SetSaml(v *CommonConfigIdpSaml) *CommonConfigIdp {
	s.Saml = v
	return s
}

func (s *CommonConfigIdp) Validate() error {
	return dara.Validate(s)
}

type CommonConfigIdpDingtalk struct {
	EventCallbackBase *string `json:"EventCallbackBase,omitempty" xml:"EventCallbackBase,omitempty"`
	EventLabel        *string `json:"EventLabel,omitempty" xml:"EventLabel,omitempty"`
	HomePage          *string `json:"HomePage,omitempty" xml:"HomePage,omitempty"`
	LoginRedirect     *string `json:"LoginRedirect,omitempty" xml:"LoginRedirect,omitempty"`
}

func (s CommonConfigIdpDingtalk) String() string {
	return dara.Prettify(s)
}

func (s CommonConfigIdpDingtalk) GoString() string {
	return s.String()
}

func (s *CommonConfigIdpDingtalk) GetEventCallbackBase() *string {
	return s.EventCallbackBase
}

func (s *CommonConfigIdpDingtalk) GetEventLabel() *string {
	return s.EventLabel
}

func (s *CommonConfigIdpDingtalk) GetHomePage() *string {
	return s.HomePage
}

func (s *CommonConfigIdpDingtalk) GetLoginRedirect() *string {
	return s.LoginRedirect
}

func (s *CommonConfigIdpDingtalk) SetEventCallbackBase(v string) *CommonConfigIdpDingtalk {
	s.EventCallbackBase = &v
	return s
}

func (s *CommonConfigIdpDingtalk) SetEventLabel(v string) *CommonConfigIdpDingtalk {
	s.EventLabel = &v
	return s
}

func (s *CommonConfigIdpDingtalk) SetHomePage(v string) *CommonConfigIdpDingtalk {
	s.HomePage = &v
	return s
}

func (s *CommonConfigIdpDingtalk) SetLoginRedirect(v string) *CommonConfigIdpDingtalk {
	s.LoginRedirect = &v
	return s
}

func (s *CommonConfigIdpDingtalk) Validate() error {
	return dara.Validate(s)
}

type CommonConfigIdpFeishu struct {
	EventCallbackBase *string `json:"EventCallbackBase,omitempty" xml:"EventCallbackBase,omitempty"`
	EventLabel        *string `json:"EventLabel,omitempty" xml:"EventLabel,omitempty"`
	HomePage          *string `json:"HomePage,omitempty" xml:"HomePage,omitempty"`
	LoginRedirect     *string `json:"LoginRedirect,omitempty" xml:"LoginRedirect,omitempty"`
}

func (s CommonConfigIdpFeishu) String() string {
	return dara.Prettify(s)
}

func (s CommonConfigIdpFeishu) GoString() string {
	return s.String()
}

func (s *CommonConfigIdpFeishu) GetEventCallbackBase() *string {
	return s.EventCallbackBase
}

func (s *CommonConfigIdpFeishu) GetEventLabel() *string {
	return s.EventLabel
}

func (s *CommonConfigIdpFeishu) GetHomePage() *string {
	return s.HomePage
}

func (s *CommonConfigIdpFeishu) GetLoginRedirect() *string {
	return s.LoginRedirect
}

func (s *CommonConfigIdpFeishu) SetEventCallbackBase(v string) *CommonConfigIdpFeishu {
	s.EventCallbackBase = &v
	return s
}

func (s *CommonConfigIdpFeishu) SetEventLabel(v string) *CommonConfigIdpFeishu {
	s.EventLabel = &v
	return s
}

func (s *CommonConfigIdpFeishu) SetHomePage(v string) *CommonConfigIdpFeishu {
	s.HomePage = &v
	return s
}

func (s *CommonConfigIdpFeishu) SetLoginRedirect(v string) *CommonConfigIdpFeishu {
	s.LoginRedirect = &v
	return s
}

func (s *CommonConfigIdpFeishu) Validate() error {
	return dara.Validate(s)
}

type CommonConfigIdpIdaas2 struct {
	EventCallbackBase *string `json:"EventCallbackBase,omitempty" xml:"EventCallbackBase,omitempty"`
	EventLabel        *string `json:"EventLabel,omitempty" xml:"EventLabel,omitempty"`
}

func (s CommonConfigIdpIdaas2) String() string {
	return dara.Prettify(s)
}

func (s CommonConfigIdpIdaas2) GoString() string {
	return s.String()
}

func (s *CommonConfigIdpIdaas2) GetEventCallbackBase() *string {
	return s.EventCallbackBase
}

func (s *CommonConfigIdpIdaas2) GetEventLabel() *string {
	return s.EventLabel
}

func (s *CommonConfigIdpIdaas2) SetEventCallbackBase(v string) *CommonConfigIdpIdaas2 {
	s.EventCallbackBase = &v
	return s
}

func (s *CommonConfigIdpIdaas2) SetEventLabel(v string) *CommonConfigIdpIdaas2 {
	s.EventLabel = &v
	return s
}

func (s *CommonConfigIdpIdaas2) Validate() error {
	return dara.Validate(s)
}

type CommonConfigIdpSaml struct {
	Acs      *string `json:"Acs,omitempty" xml:"Acs,omitempty"`
	Metadata *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
}

func (s CommonConfigIdpSaml) String() string {
	return dara.Prettify(s)
}

func (s CommonConfigIdpSaml) GoString() string {
	return s.String()
}

func (s *CommonConfigIdpSaml) GetAcs() *string {
	return s.Acs
}

func (s *CommonConfigIdpSaml) GetMetadata() *string {
	return s.Metadata
}

func (s *CommonConfigIdpSaml) SetAcs(v string) *CommonConfigIdpSaml {
	s.Acs = &v
	return s
}

func (s *CommonConfigIdpSaml) SetMetadata(v string) *CommonConfigIdpSaml {
	s.Metadata = &v
	return s
}

func (s *CommonConfigIdpSaml) Validate() error {
	return dara.Validate(s)
}
