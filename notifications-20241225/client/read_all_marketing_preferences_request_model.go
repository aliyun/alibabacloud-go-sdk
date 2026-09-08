// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadAllMarketingPreferencesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ReadAllMarketingPreferencesRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *ReadAllMarketingPreferencesRequest
	GetAppName() *string
	SetBizName(v string) *ReadAllMarketingPreferencesRequest
	GetBizName() *string
	SetCallerProtocol(v string) *ReadAllMarketingPreferencesRequest
	GetCallerProtocol() *string
	SetClientSource(v string) *ReadAllMarketingPreferencesRequest
	GetClientSource() *string
	SetCookies(v string) *ReadAllMarketingPreferencesRequest
	GetCookies() *string
	SetSrcUrl(v string) *ReadAllMarketingPreferencesRequest
	GetSrcUrl() *string
	SetTenantCode(v string) *ReadAllMarketingPreferencesRequest
	GetTenantCode() *string
	SetUidType(v string) *ReadAllMarketingPreferencesRequest
	GetUidType() *string
}

type ReadAllMarketingPreferencesRequest struct {
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
	// yunge-user-aliyun-service
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

func (s ReadAllMarketingPreferencesRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadAllMarketingPreferencesRequest) GoString() string {
	return s.String()
}

func (s *ReadAllMarketingPreferencesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ReadAllMarketingPreferencesRequest) GetAppName() *string {
	return s.AppName
}

func (s *ReadAllMarketingPreferencesRequest) GetBizName() *string {
	return s.BizName
}

func (s *ReadAllMarketingPreferencesRequest) GetCallerProtocol() *string {
	return s.CallerProtocol
}

func (s *ReadAllMarketingPreferencesRequest) GetClientSource() *string {
	return s.ClientSource
}

func (s *ReadAllMarketingPreferencesRequest) GetCookies() *string {
	return s.Cookies
}

func (s *ReadAllMarketingPreferencesRequest) GetSrcUrl() *string {
	return s.SrcUrl
}

func (s *ReadAllMarketingPreferencesRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *ReadAllMarketingPreferencesRequest) GetUidType() *string {
	return s.UidType
}

func (s *ReadAllMarketingPreferencesRequest) SetAcceptLanguage(v string) *ReadAllMarketingPreferencesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReadAllMarketingPreferencesRequest) SetAppName(v string) *ReadAllMarketingPreferencesRequest {
	s.AppName = &v
	return s
}

func (s *ReadAllMarketingPreferencesRequest) SetBizName(v string) *ReadAllMarketingPreferencesRequest {
	s.BizName = &v
	return s
}

func (s *ReadAllMarketingPreferencesRequest) SetCallerProtocol(v string) *ReadAllMarketingPreferencesRequest {
	s.CallerProtocol = &v
	return s
}

func (s *ReadAllMarketingPreferencesRequest) SetClientSource(v string) *ReadAllMarketingPreferencesRequest {
	s.ClientSource = &v
	return s
}

func (s *ReadAllMarketingPreferencesRequest) SetCookies(v string) *ReadAllMarketingPreferencesRequest {
	s.Cookies = &v
	return s
}

func (s *ReadAllMarketingPreferencesRequest) SetSrcUrl(v string) *ReadAllMarketingPreferencesRequest {
	s.SrcUrl = &v
	return s
}

func (s *ReadAllMarketingPreferencesRequest) SetTenantCode(v string) *ReadAllMarketingPreferencesRequest {
	s.TenantCode = &v
	return s
}

func (s *ReadAllMarketingPreferencesRequest) SetUidType(v string) *ReadAllMarketingPreferencesRequest {
	s.UidType = &v
	return s
}

func (s *ReadAllMarketingPreferencesRequest) Validate() error {
	return dara.Validate(s)
}
