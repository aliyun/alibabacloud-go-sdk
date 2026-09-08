// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadWebhookContactSendTemplateListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ReadWebhookContactSendTemplateListRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *ReadWebhookContactSendTemplateListRequest
	GetAppName() *string
	SetBizName(v string) *ReadWebhookContactSendTemplateListRequest
	GetBizName() *string
	SetCallerProtocol(v string) *ReadWebhookContactSendTemplateListRequest
	GetCallerProtocol() *string
	SetClientSource(v string) *ReadWebhookContactSendTemplateListRequest
	GetClientSource() *string
	SetCookies(v string) *ReadWebhookContactSendTemplateListRequest
	GetCookies() *string
	SetSrcUrl(v string) *ReadWebhookContactSendTemplateListRequest
	GetSrcUrl() *string
	SetTemplateCode(v string) *ReadWebhookContactSendTemplateListRequest
	GetTemplateCode() *string
	SetTenantCode(v string) *ReadWebhookContactSendTemplateListRequest
	GetTenantCode() *string
	SetUidType(v string) *ReadWebhookContactSendTemplateListRequest
	GetUidType() *string
}

type ReadWebhookContactSendTemplateListRequest struct {
	// The language.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The application name of the requester.
	//
	// example:
	//
	// /
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The business line of the requester.
	//
	// example:
	//
	// /
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
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
	// The user cookies.
	//
	// example:
	//
	// /
	Cookies *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
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
}

func (s ReadWebhookContactSendTemplateListRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadWebhookContactSendTemplateListRequest) GoString() string {
	return s.String()
}

func (s *ReadWebhookContactSendTemplateListRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ReadWebhookContactSendTemplateListRequest) GetAppName() *string {
	return s.AppName
}

func (s *ReadWebhookContactSendTemplateListRequest) GetBizName() *string {
	return s.BizName
}

func (s *ReadWebhookContactSendTemplateListRequest) GetCallerProtocol() *string {
	return s.CallerProtocol
}

func (s *ReadWebhookContactSendTemplateListRequest) GetClientSource() *string {
	return s.ClientSource
}

func (s *ReadWebhookContactSendTemplateListRequest) GetCookies() *string {
	return s.Cookies
}

func (s *ReadWebhookContactSendTemplateListRequest) GetSrcUrl() *string {
	return s.SrcUrl
}

func (s *ReadWebhookContactSendTemplateListRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *ReadWebhookContactSendTemplateListRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *ReadWebhookContactSendTemplateListRequest) GetUidType() *string {
	return s.UidType
}

func (s *ReadWebhookContactSendTemplateListRequest) SetAcceptLanguage(v string) *ReadWebhookContactSendTemplateListRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReadWebhookContactSendTemplateListRequest) SetAppName(v string) *ReadWebhookContactSendTemplateListRequest {
	s.AppName = &v
	return s
}

func (s *ReadWebhookContactSendTemplateListRequest) SetBizName(v string) *ReadWebhookContactSendTemplateListRequest {
	s.BizName = &v
	return s
}

func (s *ReadWebhookContactSendTemplateListRequest) SetCallerProtocol(v string) *ReadWebhookContactSendTemplateListRequest {
	s.CallerProtocol = &v
	return s
}

func (s *ReadWebhookContactSendTemplateListRequest) SetClientSource(v string) *ReadWebhookContactSendTemplateListRequest {
	s.ClientSource = &v
	return s
}

func (s *ReadWebhookContactSendTemplateListRequest) SetCookies(v string) *ReadWebhookContactSendTemplateListRequest {
	s.Cookies = &v
	return s
}

func (s *ReadWebhookContactSendTemplateListRequest) SetSrcUrl(v string) *ReadWebhookContactSendTemplateListRequest {
	s.SrcUrl = &v
	return s
}

func (s *ReadWebhookContactSendTemplateListRequest) SetTemplateCode(v string) *ReadWebhookContactSendTemplateListRequest {
	s.TemplateCode = &v
	return s
}

func (s *ReadWebhookContactSendTemplateListRequest) SetTenantCode(v string) *ReadWebhookContactSendTemplateListRequest {
	s.TenantCode = &v
	return s
}

func (s *ReadWebhookContactSendTemplateListRequest) SetUidType(v string) *ReadWebhookContactSendTemplateListRequest {
	s.UidType = &v
	return s
}

func (s *ReadWebhookContactSendTemplateListRequest) Validate() error {
	return dara.Validate(s)
}
