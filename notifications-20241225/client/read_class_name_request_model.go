// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadClassNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ReadClassNameRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *ReadClassNameRequest
	GetAppName() *string
	SetBizName(v string) *ReadClassNameRequest
	GetBizName() *string
	SetCallerProtocol(v string) *ReadClassNameRequest
	GetCallerProtocol() *string
	SetClientSource(v string) *ReadClassNameRequest
	GetClientSource() *string
	SetCookies(v string) *ReadClassNameRequest
	GetCookies() *string
	SetSrcUrl(v string) *ReadClassNameRequest
	GetSrcUrl() *string
	SetTenantCode(v string) *ReadClassNameRequest
	GetTenantCode() *string
	SetUidType(v string) *ReadClassNameRequest
	GetUidType() *string
}

type ReadClassNameRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BizName        *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	CallerProtocol *string `json:"CallerProtocol,omitempty" xml:"CallerProtocol,omitempty"`
	ClientSource   *string `json:"ClientSource,omitempty" xml:"ClientSource,omitempty"`
	Cookies        *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
	SrcUrl         *string `json:"SrcUrl,omitempty" xml:"SrcUrl,omitempty"`
	TenantCode     *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
	UidType        *string `json:"UidType,omitempty" xml:"UidType,omitempty"`
}

func (s ReadClassNameRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadClassNameRequest) GoString() string {
	return s.String()
}

func (s *ReadClassNameRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ReadClassNameRequest) GetAppName() *string {
	return s.AppName
}

func (s *ReadClassNameRequest) GetBizName() *string {
	return s.BizName
}

func (s *ReadClassNameRequest) GetCallerProtocol() *string {
	return s.CallerProtocol
}

func (s *ReadClassNameRequest) GetClientSource() *string {
	return s.ClientSource
}

func (s *ReadClassNameRequest) GetCookies() *string {
	return s.Cookies
}

func (s *ReadClassNameRequest) GetSrcUrl() *string {
	return s.SrcUrl
}

func (s *ReadClassNameRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *ReadClassNameRequest) GetUidType() *string {
	return s.UidType
}

func (s *ReadClassNameRequest) SetAcceptLanguage(v string) *ReadClassNameRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReadClassNameRequest) SetAppName(v string) *ReadClassNameRequest {
	s.AppName = &v
	return s
}

func (s *ReadClassNameRequest) SetBizName(v string) *ReadClassNameRequest {
	s.BizName = &v
	return s
}

func (s *ReadClassNameRequest) SetCallerProtocol(v string) *ReadClassNameRequest {
	s.CallerProtocol = &v
	return s
}

func (s *ReadClassNameRequest) SetClientSource(v string) *ReadClassNameRequest {
	s.ClientSource = &v
	return s
}

func (s *ReadClassNameRequest) SetCookies(v string) *ReadClassNameRequest {
	s.Cookies = &v
	return s
}

func (s *ReadClassNameRequest) SetSrcUrl(v string) *ReadClassNameRequest {
	s.SrcUrl = &v
	return s
}

func (s *ReadClassNameRequest) SetTenantCode(v string) *ReadClassNameRequest {
	s.TenantCode = &v
	return s
}

func (s *ReadClassNameRequest) SetUidType(v string) *ReadClassNameRequest {
	s.UidType = &v
	return s
}

func (s *ReadClassNameRequest) Validate() error {
	return dara.Validate(s)
}
