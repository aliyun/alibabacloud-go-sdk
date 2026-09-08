// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadMessageLanguageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ReadMessageLanguageRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *ReadMessageLanguageRequest
	GetAppName() *string
	SetBizName(v string) *ReadMessageLanguageRequest
	GetBizName() *string
	SetCallerProtocol(v string) *ReadMessageLanguageRequest
	GetCallerProtocol() *string
	SetClientSource(v string) *ReadMessageLanguageRequest
	GetClientSource() *string
	SetCookies(v string) *ReadMessageLanguageRequest
	GetCookies() *string
	SetReturnDefaultLang(v string) *ReadMessageLanguageRequest
	GetReturnDefaultLang() *string
	SetSrcUrl(v string) *ReadMessageLanguageRequest
	GetSrcUrl() *string
	SetTenantCode(v string) *ReadMessageLanguageRequest
	GetTenantCode() *string
	SetUidType(v string) *ReadMessageLanguageRequest
	GetUidType() *string
}

type ReadMessageLanguageRequest struct {
	// Ignored. This parameter does not need to be specified. The page language.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// Ignored. This parameter does not need to be specified. The application name of the requester.
	//
	// example:
	//
	// /
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Ignored. This parameter does not need to be specified. The business line of the requester.
	//
	// example:
	//
	// /
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// Ignored. This parameter does not need to be specified. The request protocol type.
	//
	// example:
	//
	// /
	CallerProtocol *string `json:"CallerProtocol,omitempty" xml:"CallerProtocol,omitempty"`
	// Ignored. This parameter does not need to be specified. The operation terminal source.
	//
	// example:
	//
	// /
	ClientSource *string `json:"ClientSource,omitempty" xml:"ClientSource,omitempty"`
	// Ignored. This parameter does not need to be specified. The user cookies.
	//
	// example:
	//
	// /
	Cookies *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
	// Specifies whether to return the default language. If the value is YES, the default language is returned.
	//
	// example:
	//
	// YES
	ReturnDefaultLang *string `json:"ReturnDefaultLang,omitempty" xml:"ReturnDefaultLang,omitempty"`
	// Ignored. This parameter does not need to be specified. The source page URL.
	//
	// example:
	//
	// /
	SrcUrl *string `json:"SrcUrl,omitempty" xml:"SrcUrl,omitempty"`
	// Ignored. This parameter does not need to be specified. The tenant information.
	//
	// example:
	//
	// /
	TenantCode *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
	// Ignored. This parameter does not need to be specified. The user type.
	//
	// example:
	//
	// /
	UidType *string `json:"UidType,omitempty" xml:"UidType,omitempty"`
}

func (s ReadMessageLanguageRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadMessageLanguageRequest) GoString() string {
	return s.String()
}

func (s *ReadMessageLanguageRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ReadMessageLanguageRequest) GetAppName() *string {
	return s.AppName
}

func (s *ReadMessageLanguageRequest) GetBizName() *string {
	return s.BizName
}

func (s *ReadMessageLanguageRequest) GetCallerProtocol() *string {
	return s.CallerProtocol
}

func (s *ReadMessageLanguageRequest) GetClientSource() *string {
	return s.ClientSource
}

func (s *ReadMessageLanguageRequest) GetCookies() *string {
	return s.Cookies
}

func (s *ReadMessageLanguageRequest) GetReturnDefaultLang() *string {
	return s.ReturnDefaultLang
}

func (s *ReadMessageLanguageRequest) GetSrcUrl() *string {
	return s.SrcUrl
}

func (s *ReadMessageLanguageRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *ReadMessageLanguageRequest) GetUidType() *string {
	return s.UidType
}

func (s *ReadMessageLanguageRequest) SetAcceptLanguage(v string) *ReadMessageLanguageRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReadMessageLanguageRequest) SetAppName(v string) *ReadMessageLanguageRequest {
	s.AppName = &v
	return s
}

func (s *ReadMessageLanguageRequest) SetBizName(v string) *ReadMessageLanguageRequest {
	s.BizName = &v
	return s
}

func (s *ReadMessageLanguageRequest) SetCallerProtocol(v string) *ReadMessageLanguageRequest {
	s.CallerProtocol = &v
	return s
}

func (s *ReadMessageLanguageRequest) SetClientSource(v string) *ReadMessageLanguageRequest {
	s.ClientSource = &v
	return s
}

func (s *ReadMessageLanguageRequest) SetCookies(v string) *ReadMessageLanguageRequest {
	s.Cookies = &v
	return s
}

func (s *ReadMessageLanguageRequest) SetReturnDefaultLang(v string) *ReadMessageLanguageRequest {
	s.ReturnDefaultLang = &v
	return s
}

func (s *ReadMessageLanguageRequest) SetSrcUrl(v string) *ReadMessageLanguageRequest {
	s.SrcUrl = &v
	return s
}

func (s *ReadMessageLanguageRequest) SetTenantCode(v string) *ReadMessageLanguageRequest {
	s.TenantCode = &v
	return s
}

func (s *ReadMessageLanguageRequest) SetUidType(v string) *ReadMessageLanguageRequest {
	s.UidType = &v
	return s
}

func (s *ReadMessageLanguageRequest) Validate() error {
	return dara.Validate(s)
}
