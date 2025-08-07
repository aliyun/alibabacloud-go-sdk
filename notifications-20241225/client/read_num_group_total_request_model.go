// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadNumGroupTotalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ReadNumGroupTotalRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *ReadNumGroupTotalRequest
	GetAppName() *string
	SetBizName(v string) *ReadNumGroupTotalRequest
	GetBizName() *string
	SetCallerProtocol(v string) *ReadNumGroupTotalRequest
	GetCallerProtocol() *string
	SetClientSource(v string) *ReadNumGroupTotalRequest
	GetClientSource() *string
	SetCookies(v string) *ReadNumGroupTotalRequest
	GetCookies() *string
	SetSrcUrl(v string) *ReadNumGroupTotalRequest
	GetSrcUrl() *string
	SetTenantCode(v string) *ReadNumGroupTotalRequest
	GetTenantCode() *string
	SetTitle(v string) *ReadNumGroupTotalRequest
	GetTitle() *string
	SetUidType(v string) *ReadNumGroupTotalRequest
	GetUidType() *string
}

type ReadNumGroupTotalRequest struct {
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BizName        *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	CallerProtocol *string `json:"CallerProtocol,omitempty" xml:"CallerProtocol,omitempty"`
	ClientSource   *string `json:"ClientSource,omitempty" xml:"ClientSource,omitempty"`
	Cookies        *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
	SrcUrl         *string `json:"SrcUrl,omitempty" xml:"SrcUrl,omitempty"`
	TenantCode     *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
	Title          *string `json:"Title,omitempty" xml:"Title,omitempty"`
	UidType        *string `json:"UidType,omitempty" xml:"UidType,omitempty"`
}

func (s ReadNumGroupTotalRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadNumGroupTotalRequest) GoString() string {
	return s.String()
}

func (s *ReadNumGroupTotalRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ReadNumGroupTotalRequest) GetAppName() *string {
	return s.AppName
}

func (s *ReadNumGroupTotalRequest) GetBizName() *string {
	return s.BizName
}

func (s *ReadNumGroupTotalRequest) GetCallerProtocol() *string {
	return s.CallerProtocol
}

func (s *ReadNumGroupTotalRequest) GetClientSource() *string {
	return s.ClientSource
}

func (s *ReadNumGroupTotalRequest) GetCookies() *string {
	return s.Cookies
}

func (s *ReadNumGroupTotalRequest) GetSrcUrl() *string {
	return s.SrcUrl
}

func (s *ReadNumGroupTotalRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *ReadNumGroupTotalRequest) GetTitle() *string {
	return s.Title
}

func (s *ReadNumGroupTotalRequest) GetUidType() *string {
	return s.UidType
}

func (s *ReadNumGroupTotalRequest) SetAcceptLanguage(v string) *ReadNumGroupTotalRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReadNumGroupTotalRequest) SetAppName(v string) *ReadNumGroupTotalRequest {
	s.AppName = &v
	return s
}

func (s *ReadNumGroupTotalRequest) SetBizName(v string) *ReadNumGroupTotalRequest {
	s.BizName = &v
	return s
}

func (s *ReadNumGroupTotalRequest) SetCallerProtocol(v string) *ReadNumGroupTotalRequest {
	s.CallerProtocol = &v
	return s
}

func (s *ReadNumGroupTotalRequest) SetClientSource(v string) *ReadNumGroupTotalRequest {
	s.ClientSource = &v
	return s
}

func (s *ReadNumGroupTotalRequest) SetCookies(v string) *ReadNumGroupTotalRequest {
	s.Cookies = &v
	return s
}

func (s *ReadNumGroupTotalRequest) SetSrcUrl(v string) *ReadNumGroupTotalRequest {
	s.SrcUrl = &v
	return s
}

func (s *ReadNumGroupTotalRequest) SetTenantCode(v string) *ReadNumGroupTotalRequest {
	s.TenantCode = &v
	return s
}

func (s *ReadNumGroupTotalRequest) SetTitle(v string) *ReadNumGroupTotalRequest {
	s.Title = &v
	return s
}

func (s *ReadNumGroupTotalRequest) SetUidType(v string) *ReadNumGroupTotalRequest {
	s.UidType = &v
	return s
}

func (s *ReadNumGroupTotalRequest) Validate() error {
	return dara.Validate(s)
}
