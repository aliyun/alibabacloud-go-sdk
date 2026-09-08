// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadAllCommonContactsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ReadAllCommonContactsRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *ReadAllCommonContactsRequest
	GetAppName() *string
	SetBizName(v string) *ReadAllCommonContactsRequest
	GetBizName() *string
	SetCallerProtocol(v string) *ReadAllCommonContactsRequest
	GetCallerProtocol() *string
	SetClientSource(v string) *ReadAllCommonContactsRequest
	GetClientSource() *string
	SetCookies(v string) *ReadAllCommonContactsRequest
	GetCookies() *string
	SetSrcUrl(v string) *ReadAllCommonContactsRequest
	GetSrcUrl() *string
	SetTenantCode(v string) *ReadAllCommonContactsRequest
	GetTenantCode() *string
	SetUidType(v string) *ReadAllCommonContactsRequest
	GetUidType() *string
}

type ReadAllCommonContactsRequest struct {
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

func (s ReadAllCommonContactsRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadAllCommonContactsRequest) GoString() string {
	return s.String()
}

func (s *ReadAllCommonContactsRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ReadAllCommonContactsRequest) GetAppName() *string {
	return s.AppName
}

func (s *ReadAllCommonContactsRequest) GetBizName() *string {
	return s.BizName
}

func (s *ReadAllCommonContactsRequest) GetCallerProtocol() *string {
	return s.CallerProtocol
}

func (s *ReadAllCommonContactsRequest) GetClientSource() *string {
	return s.ClientSource
}

func (s *ReadAllCommonContactsRequest) GetCookies() *string {
	return s.Cookies
}

func (s *ReadAllCommonContactsRequest) GetSrcUrl() *string {
	return s.SrcUrl
}

func (s *ReadAllCommonContactsRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *ReadAllCommonContactsRequest) GetUidType() *string {
	return s.UidType
}

func (s *ReadAllCommonContactsRequest) SetAcceptLanguage(v string) *ReadAllCommonContactsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReadAllCommonContactsRequest) SetAppName(v string) *ReadAllCommonContactsRequest {
	s.AppName = &v
	return s
}

func (s *ReadAllCommonContactsRequest) SetBizName(v string) *ReadAllCommonContactsRequest {
	s.BizName = &v
	return s
}

func (s *ReadAllCommonContactsRequest) SetCallerProtocol(v string) *ReadAllCommonContactsRequest {
	s.CallerProtocol = &v
	return s
}

func (s *ReadAllCommonContactsRequest) SetClientSource(v string) *ReadAllCommonContactsRequest {
	s.ClientSource = &v
	return s
}

func (s *ReadAllCommonContactsRequest) SetCookies(v string) *ReadAllCommonContactsRequest {
	s.Cookies = &v
	return s
}

func (s *ReadAllCommonContactsRequest) SetSrcUrl(v string) *ReadAllCommonContactsRequest {
	s.SrcUrl = &v
	return s
}

func (s *ReadAllCommonContactsRequest) SetTenantCode(v string) *ReadAllCommonContactsRequest {
	s.TenantCode = &v
	return s
}

func (s *ReadAllCommonContactsRequest) SetUidType(v string) *ReadAllCommonContactsRequest {
	s.UidType = &v
	return s
}

func (s *ReadAllCommonContactsRequest) Validate() error {
	return dara.Validate(s)
}
