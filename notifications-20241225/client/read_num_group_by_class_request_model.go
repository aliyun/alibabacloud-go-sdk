// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadNumGroupByClassRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ReadNumGroupByClassRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *ReadNumGroupByClassRequest
	GetAppName() *string
	SetBizName(v string) *ReadNumGroupByClassRequest
	GetBizName() *string
	SetCallerProtocol(v string) *ReadNumGroupByClassRequest
	GetCallerProtocol() *string
	SetClientSource(v string) *ReadNumGroupByClassRequest
	GetClientSource() *string
	SetCookies(v string) *ReadNumGroupByClassRequest
	GetCookies() *string
	SetSrcUrl(v string) *ReadNumGroupByClassRequest
	GetSrcUrl() *string
	SetTenantCode(v string) *ReadNumGroupByClassRequest
	GetTenantCode() *string
	SetUidType(v string) *ReadNumGroupByClassRequest
	GetUidType() *string
}

type ReadNumGroupByClassRequest struct {
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
	UidType        *string `json:"UidType,omitempty" xml:"UidType,omitempty"`
}

func (s ReadNumGroupByClassRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadNumGroupByClassRequest) GoString() string {
	return s.String()
}

func (s *ReadNumGroupByClassRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ReadNumGroupByClassRequest) GetAppName() *string {
	return s.AppName
}

func (s *ReadNumGroupByClassRequest) GetBizName() *string {
	return s.BizName
}

func (s *ReadNumGroupByClassRequest) GetCallerProtocol() *string {
	return s.CallerProtocol
}

func (s *ReadNumGroupByClassRequest) GetClientSource() *string {
	return s.ClientSource
}

func (s *ReadNumGroupByClassRequest) GetCookies() *string {
	return s.Cookies
}

func (s *ReadNumGroupByClassRequest) GetSrcUrl() *string {
	return s.SrcUrl
}

func (s *ReadNumGroupByClassRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *ReadNumGroupByClassRequest) GetUidType() *string {
	return s.UidType
}

func (s *ReadNumGroupByClassRequest) SetAcceptLanguage(v string) *ReadNumGroupByClassRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReadNumGroupByClassRequest) SetAppName(v string) *ReadNumGroupByClassRequest {
	s.AppName = &v
	return s
}

func (s *ReadNumGroupByClassRequest) SetBizName(v string) *ReadNumGroupByClassRequest {
	s.BizName = &v
	return s
}

func (s *ReadNumGroupByClassRequest) SetCallerProtocol(v string) *ReadNumGroupByClassRequest {
	s.CallerProtocol = &v
	return s
}

func (s *ReadNumGroupByClassRequest) SetClientSource(v string) *ReadNumGroupByClassRequest {
	s.ClientSource = &v
	return s
}

func (s *ReadNumGroupByClassRequest) SetCookies(v string) *ReadNumGroupByClassRequest {
	s.Cookies = &v
	return s
}

func (s *ReadNumGroupByClassRequest) SetSrcUrl(v string) *ReadNumGroupByClassRequest {
	s.SrcUrl = &v
	return s
}

func (s *ReadNumGroupByClassRequest) SetTenantCode(v string) *ReadNumGroupByClassRequest {
	s.TenantCode = &v
	return s
}

func (s *ReadNumGroupByClassRequest) SetUidType(v string) *ReadNumGroupByClassRequest {
	s.UidType = &v
	return s
}

func (s *ReadNumGroupByClassRequest) Validate() error {
	return dara.Validate(s)
}
