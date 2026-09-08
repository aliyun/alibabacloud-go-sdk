// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebhookContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteWebhookContactRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *DeleteWebhookContactRequest
	GetAppName() *string
	SetBizName(v string) *DeleteWebhookContactRequest
	GetBizName() *string
	SetCallerProtocol(v string) *DeleteWebhookContactRequest
	GetCallerProtocol() *string
	SetClientSource(v string) *DeleteWebhookContactRequest
	GetClientSource() *string
	SetContactId(v int64) *DeleteWebhookContactRequest
	GetContactId() *int64
	SetCookies(v string) *DeleteWebhookContactRequest
	GetCookies() *string
	SetSrcUrl(v string) *DeleteWebhookContactRequest
	GetSrcUrl() *string
	SetTenantCode(v string) *DeleteWebhookContactRequest
	GetTenantCode() *string
	SetUidType(v string) *DeleteWebhookContactRequest
	GetUidType() *string
}

type DeleteWebhookContactRequest struct {
	// The language.
	//
	// example:
	//
	// /
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The application name of the caller.
	//
	// example:
	//
	// CallerName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The business line of the caller.
	//
	// example:
	//
	// /
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// The Protocol Type of the request.
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
	// 123
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The cookies of the user.
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

func (s DeleteWebhookContactRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebhookContactRequest) GoString() string {
	return s.String()
}

func (s *DeleteWebhookContactRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteWebhookContactRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteWebhookContactRequest) GetBizName() *string {
	return s.BizName
}

func (s *DeleteWebhookContactRequest) GetCallerProtocol() *string {
	return s.CallerProtocol
}

func (s *DeleteWebhookContactRequest) GetClientSource() *string {
	return s.ClientSource
}

func (s *DeleteWebhookContactRequest) GetContactId() *int64 {
	return s.ContactId
}

func (s *DeleteWebhookContactRequest) GetCookies() *string {
	return s.Cookies
}

func (s *DeleteWebhookContactRequest) GetSrcUrl() *string {
	return s.SrcUrl
}

func (s *DeleteWebhookContactRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *DeleteWebhookContactRequest) GetUidType() *string {
	return s.UidType
}

func (s *DeleteWebhookContactRequest) SetAcceptLanguage(v string) *DeleteWebhookContactRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteWebhookContactRequest) SetAppName(v string) *DeleteWebhookContactRequest {
	s.AppName = &v
	return s
}

func (s *DeleteWebhookContactRequest) SetBizName(v string) *DeleteWebhookContactRequest {
	s.BizName = &v
	return s
}

func (s *DeleteWebhookContactRequest) SetCallerProtocol(v string) *DeleteWebhookContactRequest {
	s.CallerProtocol = &v
	return s
}

func (s *DeleteWebhookContactRequest) SetClientSource(v string) *DeleteWebhookContactRequest {
	s.ClientSource = &v
	return s
}

func (s *DeleteWebhookContactRequest) SetContactId(v int64) *DeleteWebhookContactRequest {
	s.ContactId = &v
	return s
}

func (s *DeleteWebhookContactRequest) SetCookies(v string) *DeleteWebhookContactRequest {
	s.Cookies = &v
	return s
}

func (s *DeleteWebhookContactRequest) SetSrcUrl(v string) *DeleteWebhookContactRequest {
	s.SrcUrl = &v
	return s
}

func (s *DeleteWebhookContactRequest) SetTenantCode(v string) *DeleteWebhookContactRequest {
	s.TenantCode = &v
	return s
}

func (s *DeleteWebhookContactRequest) SetUidType(v string) *DeleteWebhookContactRequest {
	s.UidType = &v
	return s
}

func (s *DeleteWebhookContactRequest) Validate() error {
	return dara.Validate(s)
}
