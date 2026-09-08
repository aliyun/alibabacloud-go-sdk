// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWebhookContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateWebhookContactRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *CreateWebhookContactRequest
	GetAppName() *string
	SetBizName(v string) *CreateWebhookContactRequest
	GetBizName() *string
	SetBotSecurityToken(v string) *CreateWebhookContactRequest
	GetBotSecurityToken() *string
	SetCallerProtocol(v string) *CreateWebhookContactRequest
	GetCallerProtocol() *string
	SetClientSource(v string) *CreateWebhookContactRequest
	GetClientSource() *string
	SetContactName(v string) *CreateWebhookContactRequest
	GetContactName() *string
	SetCookies(v string) *CreateWebhookContactRequest
	GetCookies() *string
	SetSecurityToken(v string) *CreateWebhookContactRequest
	GetSecurityToken() *string
	SetServerUrl(v string) *CreateWebhookContactRequest
	GetServerUrl() *string
	SetSrcUrl(v string) *CreateWebhookContactRequest
	GetSrcUrl() *string
	SetTemplateCode(v string) *CreateWebhookContactRequest
	GetTemplateCode() *string
	SetTenantCode(v string) *CreateWebhookContactRequest
	GetTenantCode() *string
	SetUidType(v string) *CreateWebhookContactRequest
	GetUidType() *string
	SetVerificationCode(v string) *CreateWebhookContactRequest
	GetVerificationCode() *string
	SetWebhookType(v string) *CreateWebhookContactRequest
	GetWebhookType() *string
}

type CreateWebhookContactRequest struct {
	// The language.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The application name of the caller.
	//
	// example:
	//
	// yunge-user
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The business line of the caller.
	//
	// example:
	//
	// /
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// The webhook security signature token.
	//
	// example:
	//
	// xxxx
	BotSecurityToken *string `json:"BotSecurityToken,omitempty" xml:"BotSecurityToken,omitempty"`
	// The request protocol type.
	//
	// example:
	//
	// /
	CallerProtocol *string `json:"CallerProtocol,omitempty" xml:"CallerProtocol,omitempty"`
	// The source of the operation terminal.
	//
	// example:
	//
	// /
	ClientSource *string `json:"ClientSource,omitempty" xml:"ClientSource,omitempty"`
	// The name of the webhook contact.
	//
	// example:
	//
	// test
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// The user cookies.
	//
	// example:
	//
	// /
	Cookies *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
	// Deprecated
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The DingTalk group chatbot URL.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=xxxxxxx
	ServerUrl *string `json:"ServerUrl,omitempty" xml:"ServerUrl,omitempty"`
	// The URL of the source page.
	//
	// example:
	//
	// /
	SrcUrl *string `json:"SrcUrl,omitempty" xml:"SrcUrl,omitempty"`
	// The template code.
	//
	// example:
	//
	// lark
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The tenant information.
	//
	// example:
	//
	// /
	TenantCode *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
	// The user type.
	//
	// example:
	//
	// /
	UidType *string `json:"UidType,omitempty" xml:"UidType,omitempty"`
	// The verification code.
	//
	// example:
	//
	// 352036
	VerificationCode *string `json:"VerificationCode,omitempty" xml:"VerificationCode,omitempty"`
	// The webhook type.
	//
	// example:
	//
	// dingtalk
	WebhookType *string `json:"WebhookType,omitempty" xml:"WebhookType,omitempty"`
}

func (s CreateWebhookContactRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWebhookContactRequest) GoString() string {
	return s.String()
}

func (s *CreateWebhookContactRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateWebhookContactRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateWebhookContactRequest) GetBizName() *string {
	return s.BizName
}

func (s *CreateWebhookContactRequest) GetBotSecurityToken() *string {
	return s.BotSecurityToken
}

func (s *CreateWebhookContactRequest) GetCallerProtocol() *string {
	return s.CallerProtocol
}

func (s *CreateWebhookContactRequest) GetClientSource() *string {
	return s.ClientSource
}

func (s *CreateWebhookContactRequest) GetContactName() *string {
	return s.ContactName
}

func (s *CreateWebhookContactRequest) GetCookies() *string {
	return s.Cookies
}

func (s *CreateWebhookContactRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateWebhookContactRequest) GetServerUrl() *string {
	return s.ServerUrl
}

func (s *CreateWebhookContactRequest) GetSrcUrl() *string {
	return s.SrcUrl
}

func (s *CreateWebhookContactRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *CreateWebhookContactRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *CreateWebhookContactRequest) GetUidType() *string {
	return s.UidType
}

func (s *CreateWebhookContactRequest) GetVerificationCode() *string {
	return s.VerificationCode
}

func (s *CreateWebhookContactRequest) GetWebhookType() *string {
	return s.WebhookType
}

func (s *CreateWebhookContactRequest) SetAcceptLanguage(v string) *CreateWebhookContactRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateWebhookContactRequest) SetAppName(v string) *CreateWebhookContactRequest {
	s.AppName = &v
	return s
}

func (s *CreateWebhookContactRequest) SetBizName(v string) *CreateWebhookContactRequest {
	s.BizName = &v
	return s
}

func (s *CreateWebhookContactRequest) SetBotSecurityToken(v string) *CreateWebhookContactRequest {
	s.BotSecurityToken = &v
	return s
}

func (s *CreateWebhookContactRequest) SetCallerProtocol(v string) *CreateWebhookContactRequest {
	s.CallerProtocol = &v
	return s
}

func (s *CreateWebhookContactRequest) SetClientSource(v string) *CreateWebhookContactRequest {
	s.ClientSource = &v
	return s
}

func (s *CreateWebhookContactRequest) SetContactName(v string) *CreateWebhookContactRequest {
	s.ContactName = &v
	return s
}

func (s *CreateWebhookContactRequest) SetCookies(v string) *CreateWebhookContactRequest {
	s.Cookies = &v
	return s
}

func (s *CreateWebhookContactRequest) SetSecurityToken(v string) *CreateWebhookContactRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateWebhookContactRequest) SetServerUrl(v string) *CreateWebhookContactRequest {
	s.ServerUrl = &v
	return s
}

func (s *CreateWebhookContactRequest) SetSrcUrl(v string) *CreateWebhookContactRequest {
	s.SrcUrl = &v
	return s
}

func (s *CreateWebhookContactRequest) SetTemplateCode(v string) *CreateWebhookContactRequest {
	s.TemplateCode = &v
	return s
}

func (s *CreateWebhookContactRequest) SetTenantCode(v string) *CreateWebhookContactRequest {
	s.TenantCode = &v
	return s
}

func (s *CreateWebhookContactRequest) SetUidType(v string) *CreateWebhookContactRequest {
	s.UidType = &v
	return s
}

func (s *CreateWebhookContactRequest) SetVerificationCode(v string) *CreateWebhookContactRequest {
	s.VerificationCode = &v
	return s
}

func (s *CreateWebhookContactRequest) SetWebhookType(v string) *CreateWebhookContactRequest {
	s.WebhookType = &v
	return s
}

func (s *CreateWebhookContactRequest) Validate() error {
	return dara.Validate(s)
}
