// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadMetaConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ReadMetaConfigRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *ReadMetaConfigRequest
	GetAppName() *string
	SetBizName(v string) *ReadMetaConfigRequest
	GetBizName() *string
	SetCallerProtocol(v string) *ReadMetaConfigRequest
	GetCallerProtocol() *string
	SetClientSource(v string) *ReadMetaConfigRequest
	GetClientSource() *string
	SetCookies(v string) *ReadMetaConfigRequest
	GetCookies() *string
	SetSrcUrl(v string) *ReadMetaConfigRequest
	GetSrcUrl() *string
	SetTenantCode(v string) *ReadMetaConfigRequest
	GetTenantCode() *string
	SetUidType(v string) *ReadMetaConfigRequest
	GetUidType() *string
}

type ReadMetaConfigRequest struct {
	// The language type of the returned information. Valid values:
	//
	// - zh: Chinese.
	//
	// - en: English.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The application project name of the requester.
	//
	// example:
	//
	// xxxApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The business line of the requester.
	//
	// example:
	//
	// SystemAlerts
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// The request protocol type.
	//
	// example:
	//
	// https
	CallerProtocol *string `json:"CallerProtocol,omitempty" xml:"CallerProtocol,omitempty"`
	// The source of the operation terminal.
	//
	// example:
	//
	// h5
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
	// https://example.com/notify,0
	SrcUrl *string `json:"SrcUrl,omitempty" xml:"SrcUrl,omitempty"`
	// The tenant information.
	//
	// example:
	//
	// T002
	TenantCode *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
	// The user type.
	//
	// example:
	//
	// aliyunPk
	UidType *string `json:"UidType,omitempty" xml:"UidType,omitempty"`
}

func (s ReadMetaConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadMetaConfigRequest) GoString() string {
	return s.String()
}

func (s *ReadMetaConfigRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ReadMetaConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *ReadMetaConfigRequest) GetBizName() *string {
	return s.BizName
}

func (s *ReadMetaConfigRequest) GetCallerProtocol() *string {
	return s.CallerProtocol
}

func (s *ReadMetaConfigRequest) GetClientSource() *string {
	return s.ClientSource
}

func (s *ReadMetaConfigRequest) GetCookies() *string {
	return s.Cookies
}

func (s *ReadMetaConfigRequest) GetSrcUrl() *string {
	return s.SrcUrl
}

func (s *ReadMetaConfigRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *ReadMetaConfigRequest) GetUidType() *string {
	return s.UidType
}

func (s *ReadMetaConfigRequest) SetAcceptLanguage(v string) *ReadMetaConfigRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReadMetaConfigRequest) SetAppName(v string) *ReadMetaConfigRequest {
	s.AppName = &v
	return s
}

func (s *ReadMetaConfigRequest) SetBizName(v string) *ReadMetaConfigRequest {
	s.BizName = &v
	return s
}

func (s *ReadMetaConfigRequest) SetCallerProtocol(v string) *ReadMetaConfigRequest {
	s.CallerProtocol = &v
	return s
}

func (s *ReadMetaConfigRequest) SetClientSource(v string) *ReadMetaConfigRequest {
	s.ClientSource = &v
	return s
}

func (s *ReadMetaConfigRequest) SetCookies(v string) *ReadMetaConfigRequest {
	s.Cookies = &v
	return s
}

func (s *ReadMetaConfigRequest) SetSrcUrl(v string) *ReadMetaConfigRequest {
	s.SrcUrl = &v
	return s
}

func (s *ReadMetaConfigRequest) SetTenantCode(v string) *ReadMetaConfigRequest {
	s.TenantCode = &v
	return s
}

func (s *ReadMetaConfigRequest) SetUidType(v string) *ReadMetaConfigRequest {
	s.UidType = &v
	return s
}

func (s *ReadMetaConfigRequest) Validate() error {
	return dara.Validate(s)
}
