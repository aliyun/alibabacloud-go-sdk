// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWebhookContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateWebhookContactRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *UpdateWebhookContactRequest
	GetAppName() *string
	SetBizName(v string) *UpdateWebhookContactRequest
	GetBizName() *string
	SetBotSecurityToken(v string) *UpdateWebhookContactRequest
	GetBotSecurityToken() *string
	SetCallerProtocol(v string) *UpdateWebhookContactRequest
	GetCallerProtocol() *string
	SetClientSource(v string) *UpdateWebhookContactRequest
	GetClientSource() *string
	SetContactId(v int64) *UpdateWebhookContactRequest
	GetContactId() *int64
	SetContactName(v string) *UpdateWebhookContactRequest
	GetContactName() *string
	SetCookies(v string) *UpdateWebhookContactRequest
	GetCookies() *string
	SetSecurityToken(v string) *UpdateWebhookContactRequest
	GetSecurityToken() *string
	SetServerUrl(v string) *UpdateWebhookContactRequest
	GetServerUrl() *string
	SetSrcUrl(v string) *UpdateWebhookContactRequest
	GetSrcUrl() *string
	SetTemplateCode(v string) *UpdateWebhookContactRequest
	GetTemplateCode() *string
	SetTenantCode(v string) *UpdateWebhookContactRequest
	GetTenantCode() *string
	SetUidType(v string) *UpdateWebhookContactRequest
	GetUidType() *string
	SetVerificationCode(v string) *UpdateWebhookContactRequest
	GetVerificationCode() *string
	SetWebhookType(v string) *UpdateWebhookContactRequest
	GetWebhookType() *string
}

type UpdateWebhookContactRequest struct {
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
	// /
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The business line of the caller.
	//
	// example:
	//
	// /
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// The security signature.
	//
	// example:
	//
	// ***
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
	// webhook id
	//
	// example:
	//
	// 3
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The webhook name.
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
	// The webhook server URL.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=xxxxx
	ServerUrl *string `json:"ServerUrl,omitempty" xml:"ServerUrl,omitempty"`
	// The URL of the source page.
	//
	// example:
	//
	// /
	SrcUrl *string `json:"SrcUrl,omitempty" xml:"SrcUrl,omitempty"`
	// The template code. This parameter is required only for custom webhooks.
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
	// 123456
	VerificationCode *string `json:"VerificationCode,omitempty" xml:"VerificationCode,omitempty"`
	// The webhook type.
	//
	// example:
	//
	// dingtalk
	WebhookType *string `json:"WebhookType,omitempty" xml:"WebhookType,omitempty"`
}

func (s UpdateWebhookContactRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWebhookContactRequest) GoString() string {
	return s.String()
}

func (s *UpdateWebhookContactRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateWebhookContactRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateWebhookContactRequest) GetBizName() *string {
	return s.BizName
}

func (s *UpdateWebhookContactRequest) GetBotSecurityToken() *string {
	return s.BotSecurityToken
}

func (s *UpdateWebhookContactRequest) GetCallerProtocol() *string {
	return s.CallerProtocol
}

func (s *UpdateWebhookContactRequest) GetClientSource() *string {
	return s.ClientSource
}

func (s *UpdateWebhookContactRequest) GetContactId() *int64 {
	return s.ContactId
}

func (s *UpdateWebhookContactRequest) GetContactName() *string {
	return s.ContactName
}

func (s *UpdateWebhookContactRequest) GetCookies() *string {
	return s.Cookies
}

func (s *UpdateWebhookContactRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *UpdateWebhookContactRequest) GetServerUrl() *string {
	return s.ServerUrl
}

func (s *UpdateWebhookContactRequest) GetSrcUrl() *string {
	return s.SrcUrl
}

func (s *UpdateWebhookContactRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *UpdateWebhookContactRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *UpdateWebhookContactRequest) GetUidType() *string {
	return s.UidType
}

func (s *UpdateWebhookContactRequest) GetVerificationCode() *string {
	return s.VerificationCode
}

func (s *UpdateWebhookContactRequest) GetWebhookType() *string {
	return s.WebhookType
}

func (s *UpdateWebhookContactRequest) SetAcceptLanguage(v string) *UpdateWebhookContactRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateWebhookContactRequest) SetAppName(v string) *UpdateWebhookContactRequest {
	s.AppName = &v
	return s
}

func (s *UpdateWebhookContactRequest) SetBizName(v string) *UpdateWebhookContactRequest {
	s.BizName = &v
	return s
}

func (s *UpdateWebhookContactRequest) SetBotSecurityToken(v string) *UpdateWebhookContactRequest {
	s.BotSecurityToken = &v
	return s
}

func (s *UpdateWebhookContactRequest) SetCallerProtocol(v string) *UpdateWebhookContactRequest {
	s.CallerProtocol = &v
	return s
}

func (s *UpdateWebhookContactRequest) SetClientSource(v string) *UpdateWebhookContactRequest {
	s.ClientSource = &v
	return s
}

func (s *UpdateWebhookContactRequest) SetContactId(v int64) *UpdateWebhookContactRequest {
	s.ContactId = &v
	return s
}

func (s *UpdateWebhookContactRequest) SetContactName(v string) *UpdateWebhookContactRequest {
	s.ContactName = &v
	return s
}

func (s *UpdateWebhookContactRequest) SetCookies(v string) *UpdateWebhookContactRequest {
	s.Cookies = &v
	return s
}

func (s *UpdateWebhookContactRequest) SetSecurityToken(v string) *UpdateWebhookContactRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpdateWebhookContactRequest) SetServerUrl(v string) *UpdateWebhookContactRequest {
	s.ServerUrl = &v
	return s
}

func (s *UpdateWebhookContactRequest) SetSrcUrl(v string) *UpdateWebhookContactRequest {
	s.SrcUrl = &v
	return s
}

func (s *UpdateWebhookContactRequest) SetTemplateCode(v string) *UpdateWebhookContactRequest {
	s.TemplateCode = &v
	return s
}

func (s *UpdateWebhookContactRequest) SetTenantCode(v string) *UpdateWebhookContactRequest {
	s.TenantCode = &v
	return s
}

func (s *UpdateWebhookContactRequest) SetUidType(v string) *UpdateWebhookContactRequest {
	s.UidType = &v
	return s
}

func (s *UpdateWebhookContactRequest) SetVerificationCode(v string) *UpdateWebhookContactRequest {
	s.VerificationCode = &v
	return s
}

func (s *UpdateWebhookContactRequest) SetWebhookType(v string) *UpdateWebhookContactRequest {
	s.WebhookType = &v
	return s
}

func (s *UpdateWebhookContactRequest) Validate() error {
	return dara.Validate(s)
}
