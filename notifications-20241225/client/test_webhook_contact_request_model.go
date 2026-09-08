// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestWebhookContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *TestWebhookContactRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *TestWebhookContactRequest
	GetAppName() *string
	SetBizName(v string) *TestWebhookContactRequest
	GetBizName() *string
	SetBotSecurityToken(v string) *TestWebhookContactRequest
	GetBotSecurityToken() *string
	SetCallerProtocol(v string) *TestWebhookContactRequest
	GetCallerProtocol() *string
	SetClientSource(v string) *TestWebhookContactRequest
	GetClientSource() *string
	SetContactId(v int64) *TestWebhookContactRequest
	GetContactId() *int64
	SetContactName(v string) *TestWebhookContactRequest
	GetContactName() *string
	SetCookies(v string) *TestWebhookContactRequest
	GetCookies() *string
	SetServerUrl(v string) *TestWebhookContactRequest
	GetServerUrl() *string
	SetSrcUrl(v string) *TestWebhookContactRequest
	GetSrcUrl() *string
	SetTemplateCode(v string) *TestWebhookContactRequest
	GetTemplateCode() *string
	SetTenantCode(v string) *TestWebhookContactRequest
	GetTenantCode() *string
	SetUidType(v string) *TestWebhookContactRequest
	GetUidType() *string
	SetWebhookType(v string) *TestWebhookContactRequest
	GetWebhookType() *string
}

type TestWebhookContactRequest struct {
	// The language.
	//
	// example:
	//
	// /
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The name of the caller application.
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
	// The security token.
	//
	// example:
	//
	// /
	BotSecurityToken *string `json:"BotSecurityToken,omitempty" xml:"BotSecurityToken,omitempty"`
	// The type of the request protocol.
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
	// 0
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
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
	// The webhook server URL.
	//
	// example:
	//
	// /
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
	// The webhook type.
	//
	// example:
	//
	// dingtalk
	WebhookType *string `json:"WebhookType,omitempty" xml:"WebhookType,omitempty"`
}

func (s TestWebhookContactRequest) String() string {
	return dara.Prettify(s)
}

func (s TestWebhookContactRequest) GoString() string {
	return s.String()
}

func (s *TestWebhookContactRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *TestWebhookContactRequest) GetAppName() *string {
	return s.AppName
}

func (s *TestWebhookContactRequest) GetBizName() *string {
	return s.BizName
}

func (s *TestWebhookContactRequest) GetBotSecurityToken() *string {
	return s.BotSecurityToken
}

func (s *TestWebhookContactRequest) GetCallerProtocol() *string {
	return s.CallerProtocol
}

func (s *TestWebhookContactRequest) GetClientSource() *string {
	return s.ClientSource
}

func (s *TestWebhookContactRequest) GetContactId() *int64 {
	return s.ContactId
}

func (s *TestWebhookContactRequest) GetContactName() *string {
	return s.ContactName
}

func (s *TestWebhookContactRequest) GetCookies() *string {
	return s.Cookies
}

func (s *TestWebhookContactRequest) GetServerUrl() *string {
	return s.ServerUrl
}

func (s *TestWebhookContactRequest) GetSrcUrl() *string {
	return s.SrcUrl
}

func (s *TestWebhookContactRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *TestWebhookContactRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *TestWebhookContactRequest) GetUidType() *string {
	return s.UidType
}

func (s *TestWebhookContactRequest) GetWebhookType() *string {
	return s.WebhookType
}

func (s *TestWebhookContactRequest) SetAcceptLanguage(v string) *TestWebhookContactRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *TestWebhookContactRequest) SetAppName(v string) *TestWebhookContactRequest {
	s.AppName = &v
	return s
}

func (s *TestWebhookContactRequest) SetBizName(v string) *TestWebhookContactRequest {
	s.BizName = &v
	return s
}

func (s *TestWebhookContactRequest) SetBotSecurityToken(v string) *TestWebhookContactRequest {
	s.BotSecurityToken = &v
	return s
}

func (s *TestWebhookContactRequest) SetCallerProtocol(v string) *TestWebhookContactRequest {
	s.CallerProtocol = &v
	return s
}

func (s *TestWebhookContactRequest) SetClientSource(v string) *TestWebhookContactRequest {
	s.ClientSource = &v
	return s
}

func (s *TestWebhookContactRequest) SetContactId(v int64) *TestWebhookContactRequest {
	s.ContactId = &v
	return s
}

func (s *TestWebhookContactRequest) SetContactName(v string) *TestWebhookContactRequest {
	s.ContactName = &v
	return s
}

func (s *TestWebhookContactRequest) SetCookies(v string) *TestWebhookContactRequest {
	s.Cookies = &v
	return s
}

func (s *TestWebhookContactRequest) SetServerUrl(v string) *TestWebhookContactRequest {
	s.ServerUrl = &v
	return s
}

func (s *TestWebhookContactRequest) SetSrcUrl(v string) *TestWebhookContactRequest {
	s.SrcUrl = &v
	return s
}

func (s *TestWebhookContactRequest) SetTemplateCode(v string) *TestWebhookContactRequest {
	s.TemplateCode = &v
	return s
}

func (s *TestWebhookContactRequest) SetTenantCode(v string) *TestWebhookContactRequest {
	s.TenantCode = &v
	return s
}

func (s *TestWebhookContactRequest) SetUidType(v string) *TestWebhookContactRequest {
	s.UidType = &v
	return s
}

func (s *TestWebhookContactRequest) SetWebhookType(v string) *TestWebhookContactRequest {
	s.WebhookType = &v
	return s
}

func (s *TestWebhookContactRequest) Validate() error {
	return dara.Validate(s)
}
