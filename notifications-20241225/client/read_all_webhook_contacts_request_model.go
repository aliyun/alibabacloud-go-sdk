// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadAllWebhookContactsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ReadAllWebhookContactsRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *ReadAllWebhookContactsRequest
	GetAppName() *string
	SetBizName(v string) *ReadAllWebhookContactsRequest
	GetBizName() *string
	SetCallerProtocol(v string) *ReadAllWebhookContactsRequest
	GetCallerProtocol() *string
	SetClientSource(v string) *ReadAllWebhookContactsRequest
	GetClientSource() *string
	SetCookies(v string) *ReadAllWebhookContactsRequest
	GetCookies() *string
	SetSrcUrl(v string) *ReadAllWebhookContactsRequest
	GetSrcUrl() *string
	SetTenantCode(v string) *ReadAllWebhookContactsRequest
	GetTenantCode() *string
	SetUidType(v string) *ReadAllWebhookContactsRequest
	GetUidType() *string
}

type ReadAllWebhookContactsRequest struct {
	// The language. Automatically passed through by the browser. You can forcefully override this value.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// Ignore. This parameter does not need to be specified. The application name of the caller.
	//
	// example:
	//
	// /
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Ignore. This parameter does not need to be specified. The business line of the caller.
	//
	// example:
	//
	// /
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// Ignore. This parameter does not need to be specified. The request protocol type.
	//
	// example:
	//
	// /
	CallerProtocol *string `json:"CallerProtocol,omitempty" xml:"CallerProtocol,omitempty"`
	// Ignore. This parameter does not need to be specified. The client source of the operation.
	//
	// example:
	//
	// /
	ClientSource *string `json:"ClientSource,omitempty" xml:"ClientSource,omitempty"`
	// Ignore. This parameter does not need to be specified. The user cookies.
	//
	// example:
	//
	// /
	Cookies *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
	// Ignore. This parameter does not need to be specified. The URL of the source page.
	//
	// example:
	//
	// /
	SrcUrl *string `json:"SrcUrl,omitempty" xml:"SrcUrl,omitempty"`
	// Ignore. This parameter does not need to be specified. The tenant information.
	//
	// example:
	//
	// /
	TenantCode *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
	// Ignore. This parameter does not need to be specified. The user type.
	//
	// example:
	//
	// /
	UidType *string `json:"UidType,omitempty" xml:"UidType,omitempty"`
}

func (s ReadAllWebhookContactsRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadAllWebhookContactsRequest) GoString() string {
	return s.String()
}

func (s *ReadAllWebhookContactsRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ReadAllWebhookContactsRequest) GetAppName() *string {
	return s.AppName
}

func (s *ReadAllWebhookContactsRequest) GetBizName() *string {
	return s.BizName
}

func (s *ReadAllWebhookContactsRequest) GetCallerProtocol() *string {
	return s.CallerProtocol
}

func (s *ReadAllWebhookContactsRequest) GetClientSource() *string {
	return s.ClientSource
}

func (s *ReadAllWebhookContactsRequest) GetCookies() *string {
	return s.Cookies
}

func (s *ReadAllWebhookContactsRequest) GetSrcUrl() *string {
	return s.SrcUrl
}

func (s *ReadAllWebhookContactsRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *ReadAllWebhookContactsRequest) GetUidType() *string {
	return s.UidType
}

func (s *ReadAllWebhookContactsRequest) SetAcceptLanguage(v string) *ReadAllWebhookContactsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReadAllWebhookContactsRequest) SetAppName(v string) *ReadAllWebhookContactsRequest {
	s.AppName = &v
	return s
}

func (s *ReadAllWebhookContactsRequest) SetBizName(v string) *ReadAllWebhookContactsRequest {
	s.BizName = &v
	return s
}

func (s *ReadAllWebhookContactsRequest) SetCallerProtocol(v string) *ReadAllWebhookContactsRequest {
	s.CallerProtocol = &v
	return s
}

func (s *ReadAllWebhookContactsRequest) SetClientSource(v string) *ReadAllWebhookContactsRequest {
	s.ClientSource = &v
	return s
}

func (s *ReadAllWebhookContactsRequest) SetCookies(v string) *ReadAllWebhookContactsRequest {
	s.Cookies = &v
	return s
}

func (s *ReadAllWebhookContactsRequest) SetSrcUrl(v string) *ReadAllWebhookContactsRequest {
	s.SrcUrl = &v
	return s
}

func (s *ReadAllWebhookContactsRequest) SetTenantCode(v string) *ReadAllWebhookContactsRequest {
	s.TenantCode = &v
	return s
}

func (s *ReadAllWebhookContactsRequest) SetUidType(v string) *ReadAllWebhookContactsRequest {
	s.UidType = &v
	return s
}

func (s *ReadAllWebhookContactsRequest) Validate() error {
	return dara.Validate(s)
}
