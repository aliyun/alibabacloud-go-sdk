// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMessageLanguageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateMessageLanguageRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *UpdateMessageLanguageRequest
	GetAppName() *string
	SetBizName(v string) *UpdateMessageLanguageRequest
	GetBizName() *string
	SetCallerProtocol(v string) *UpdateMessageLanguageRequest
	GetCallerProtocol() *string
	SetClientSource(v string) *UpdateMessageLanguageRequest
	GetClientSource() *string
	SetCookies(v string) *UpdateMessageLanguageRequest
	GetCookies() *string
	SetPreferLang(v string) *UpdateMessageLanguageRequest
	GetPreferLang() *string
	SetSrcUrl(v string) *UpdateMessageLanguageRequest
	GetSrcUrl() *string
	SetTenantCode(v string) *UpdateMessageLanguageRequest
	GetTenantCode() *string
	SetUidType(v string) *UpdateMessageLanguageRequest
	GetUidType() *string
}

type UpdateMessageLanguageRequest struct {
	// Ignore. This parameter does not need to be specified. The page language.
	//
	// example:
	//
	// /
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// Ignore. This parameter does not need to be specified. The application name of the requester.
	//
	// example:
	//
	// /
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Ignore. This parameter does not need to be specified. The business line of the requester.
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
	// Ignore. This parameter does not need to be specified. The operation terminal source.
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
	// The message language. This parameter is required.
	//
	// example:
	//
	// zh-CN
	PreferLang *string `json:"PreferLang,omitempty" xml:"PreferLang,omitempty"`
	// Ignore. This parameter does not need to be specified. The source page URL.
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

func (s UpdateMessageLanguageRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMessageLanguageRequest) GoString() string {
	return s.String()
}

func (s *UpdateMessageLanguageRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateMessageLanguageRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateMessageLanguageRequest) GetBizName() *string {
	return s.BizName
}

func (s *UpdateMessageLanguageRequest) GetCallerProtocol() *string {
	return s.CallerProtocol
}

func (s *UpdateMessageLanguageRequest) GetClientSource() *string {
	return s.ClientSource
}

func (s *UpdateMessageLanguageRequest) GetCookies() *string {
	return s.Cookies
}

func (s *UpdateMessageLanguageRequest) GetPreferLang() *string {
	return s.PreferLang
}

func (s *UpdateMessageLanguageRequest) GetSrcUrl() *string {
	return s.SrcUrl
}

func (s *UpdateMessageLanguageRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *UpdateMessageLanguageRequest) GetUidType() *string {
	return s.UidType
}

func (s *UpdateMessageLanguageRequest) SetAcceptLanguage(v string) *UpdateMessageLanguageRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateMessageLanguageRequest) SetAppName(v string) *UpdateMessageLanguageRequest {
	s.AppName = &v
	return s
}

func (s *UpdateMessageLanguageRequest) SetBizName(v string) *UpdateMessageLanguageRequest {
	s.BizName = &v
	return s
}

func (s *UpdateMessageLanguageRequest) SetCallerProtocol(v string) *UpdateMessageLanguageRequest {
	s.CallerProtocol = &v
	return s
}

func (s *UpdateMessageLanguageRequest) SetClientSource(v string) *UpdateMessageLanguageRequest {
	s.ClientSource = &v
	return s
}

func (s *UpdateMessageLanguageRequest) SetCookies(v string) *UpdateMessageLanguageRequest {
	s.Cookies = &v
	return s
}

func (s *UpdateMessageLanguageRequest) SetPreferLang(v string) *UpdateMessageLanguageRequest {
	s.PreferLang = &v
	return s
}

func (s *UpdateMessageLanguageRequest) SetSrcUrl(v string) *UpdateMessageLanguageRequest {
	s.SrcUrl = &v
	return s
}

func (s *UpdateMessageLanguageRequest) SetTenantCode(v string) *UpdateMessageLanguageRequest {
	s.TenantCode = &v
	return s
}

func (s *UpdateMessageLanguageRequest) SetUidType(v string) *UpdateMessageLanguageRequest {
	s.UidType = &v
	return s
}

func (s *UpdateMessageLanguageRequest) Validate() error {
	return dara.Validate(s)
}
