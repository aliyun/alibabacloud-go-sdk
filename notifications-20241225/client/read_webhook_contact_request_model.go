// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadWebhookContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ReadWebhookContactRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *ReadWebhookContactRequest
	GetAppName() *string
	SetBizName(v string) *ReadWebhookContactRequest
	GetBizName() *string
	SetCallerProtocol(v string) *ReadWebhookContactRequest
	GetCallerProtocol() *string
	SetClientSource(v string) *ReadWebhookContactRequest
	GetClientSource() *string
	SetContactId(v int64) *ReadWebhookContactRequest
	GetContactId() *int64
	SetCookies(v string) *ReadWebhookContactRequest
	GetCookies() *string
	SetSrcUrl(v string) *ReadWebhookContactRequest
	GetSrcUrl() *string
	SetTenantCode(v string) *ReadWebhookContactRequest
	GetTenantCode() *string
	SetUidType(v string) *ReadWebhookContactRequest
	GetUidType() *string
}

type ReadWebhookContactRequest struct {
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
	// 0
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
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

func (s ReadWebhookContactRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadWebhookContactRequest) GoString() string {
	return s.String()
}

func (s *ReadWebhookContactRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ReadWebhookContactRequest) GetAppName() *string {
	return s.AppName
}

func (s *ReadWebhookContactRequest) GetBizName() *string {
	return s.BizName
}

func (s *ReadWebhookContactRequest) GetCallerProtocol() *string {
	return s.CallerProtocol
}

func (s *ReadWebhookContactRequest) GetClientSource() *string {
	return s.ClientSource
}

func (s *ReadWebhookContactRequest) GetContactId() *int64 {
	return s.ContactId
}

func (s *ReadWebhookContactRequest) GetCookies() *string {
	return s.Cookies
}

func (s *ReadWebhookContactRequest) GetSrcUrl() *string {
	return s.SrcUrl
}

func (s *ReadWebhookContactRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *ReadWebhookContactRequest) GetUidType() *string {
	return s.UidType
}

func (s *ReadWebhookContactRequest) SetAcceptLanguage(v string) *ReadWebhookContactRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReadWebhookContactRequest) SetAppName(v string) *ReadWebhookContactRequest {
	s.AppName = &v
	return s
}

func (s *ReadWebhookContactRequest) SetBizName(v string) *ReadWebhookContactRequest {
	s.BizName = &v
	return s
}

func (s *ReadWebhookContactRequest) SetCallerProtocol(v string) *ReadWebhookContactRequest {
	s.CallerProtocol = &v
	return s
}

func (s *ReadWebhookContactRequest) SetClientSource(v string) *ReadWebhookContactRequest {
	s.ClientSource = &v
	return s
}

func (s *ReadWebhookContactRequest) SetContactId(v int64) *ReadWebhookContactRequest {
	s.ContactId = &v
	return s
}

func (s *ReadWebhookContactRequest) SetCookies(v string) *ReadWebhookContactRequest {
	s.Cookies = &v
	return s
}

func (s *ReadWebhookContactRequest) SetSrcUrl(v string) *ReadWebhookContactRequest {
	s.SrcUrl = &v
	return s
}

func (s *ReadWebhookContactRequest) SetTenantCode(v string) *ReadWebhookContactRequest {
	s.TenantCode = &v
	return s
}

func (s *ReadWebhookContactRequest) SetUidType(v string) *ReadWebhookContactRequest {
	s.UidType = &v
	return s
}

func (s *ReadWebhookContactRequest) Validate() error {
	return dara.Validate(s)
}
