// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadMessageNewTotalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ReadMessageNewTotalRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *ReadMessageNewTotalRequest
	GetAppName() *string
	SetBizName(v string) *ReadMessageNewTotalRequest
	GetBizName() *string
	SetCallerProtocol(v string) *ReadMessageNewTotalRequest
	GetCallerProtocol() *string
	SetClientSource(v string) *ReadMessageNewTotalRequest
	GetClientSource() *string
	SetCookies(v string) *ReadMessageNewTotalRequest
	GetCookies() *string
	SetSrcUrl(v string) *ReadMessageNewTotalRequest
	GetSrcUrl() *string
	SetTenantCode(v string) *ReadMessageNewTotalRequest
	GetTenantCode() *string
	SetUidType(v string) *ReadMessageNewTotalRequest
	GetUidType() *string
}

type ReadMessageNewTotalRequest struct {
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

func (s ReadMessageNewTotalRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadMessageNewTotalRequest) GoString() string {
	return s.String()
}

func (s *ReadMessageNewTotalRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ReadMessageNewTotalRequest) GetAppName() *string {
	return s.AppName
}

func (s *ReadMessageNewTotalRequest) GetBizName() *string {
	return s.BizName
}

func (s *ReadMessageNewTotalRequest) GetCallerProtocol() *string {
	return s.CallerProtocol
}

func (s *ReadMessageNewTotalRequest) GetClientSource() *string {
	return s.ClientSource
}

func (s *ReadMessageNewTotalRequest) GetCookies() *string {
	return s.Cookies
}

func (s *ReadMessageNewTotalRequest) GetSrcUrl() *string {
	return s.SrcUrl
}

func (s *ReadMessageNewTotalRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *ReadMessageNewTotalRequest) GetUidType() *string {
	return s.UidType
}

func (s *ReadMessageNewTotalRequest) SetAcceptLanguage(v string) *ReadMessageNewTotalRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReadMessageNewTotalRequest) SetAppName(v string) *ReadMessageNewTotalRequest {
	s.AppName = &v
	return s
}

func (s *ReadMessageNewTotalRequest) SetBizName(v string) *ReadMessageNewTotalRequest {
	s.BizName = &v
	return s
}

func (s *ReadMessageNewTotalRequest) SetCallerProtocol(v string) *ReadMessageNewTotalRequest {
	s.CallerProtocol = &v
	return s
}

func (s *ReadMessageNewTotalRequest) SetClientSource(v string) *ReadMessageNewTotalRequest {
	s.ClientSource = &v
	return s
}

func (s *ReadMessageNewTotalRequest) SetCookies(v string) *ReadMessageNewTotalRequest {
	s.Cookies = &v
	return s
}

func (s *ReadMessageNewTotalRequest) SetSrcUrl(v string) *ReadMessageNewTotalRequest {
	s.SrcUrl = &v
	return s
}

func (s *ReadMessageNewTotalRequest) SetTenantCode(v string) *ReadMessageNewTotalRequest {
	s.TenantCode = &v
	return s
}

func (s *ReadMessageNewTotalRequest) SetUidType(v string) *ReadMessageNewTotalRequest {
	s.UidType = &v
	return s
}

func (s *ReadMessageNewTotalRequest) Validate() error {
	return dara.Validate(s)
}
