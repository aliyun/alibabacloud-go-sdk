// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadCategoryGroupListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ReadCategoryGroupListRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *ReadCategoryGroupListRequest
	GetAppName() *string
	SetBizName(v string) *ReadCategoryGroupListRequest
	GetBizName() *string
	SetCallerProtocol(v string) *ReadCategoryGroupListRequest
	GetCallerProtocol() *string
	SetChannelGroupCode(v string) *ReadCategoryGroupListRequest
	GetChannelGroupCode() *string
	SetClientSource(v string) *ReadCategoryGroupListRequest
	GetClientSource() *string
	SetCookies(v string) *ReadCategoryGroupListRequest
	GetCookies() *string
	SetSrcUrl(v string) *ReadCategoryGroupListRequest
	GetSrcUrl() *string
	SetTenantCode(v string) *ReadCategoryGroupListRequest
	GetTenantCode() *string
	SetUidType(v string) *ReadCategoryGroupListRequest
	GetUidType() *string
}

type ReadCategoryGroupListRequest struct {
	// The language. Automatically passed through by the browser and can be manually overridden.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// Ignore. No need to pass for now. Application name of the requester.
	//
	// example:
	//
	// /
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Ignore. No need to pass for now. Business line of the requester.
	//
	// example:
	//
	// /
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// Ignore. No need to pass for now. Request protocol type.
	//
	// example:
	//
	// /
	CallerProtocol *string `json:"CallerProtocol,omitempty" xml:"CallerProtocol,omitempty"`
	// The channel group.
	//
	// example:
	//
	// base
	ChannelGroupCode *string `json:"ChannelGroupCode,omitempty" xml:"ChannelGroupCode,omitempty"`
	// Ignore. No need to pass for now. Operation terminal source.
	//
	// example:
	//
	// /
	ClientSource *string `json:"ClientSource,omitempty" xml:"ClientSource,omitempty"`
	// Ignore. No need to pass for now. User cookies.
	//
	// example:
	//
	// /
	Cookies *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
	// Ignore. No need to pass for now. Source page URL.
	//
	// example:
	//
	// /
	SrcUrl *string `json:"SrcUrl,omitempty" xml:"SrcUrl,omitempty"`
	// Ignore. No need to pass for now. Tenant information.
	//
	// example:
	//
	// /
	TenantCode *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
	// Ignore. No need to pass for now. User type.
	//
	// example:
	//
	// /
	UidType *string `json:"UidType,omitempty" xml:"UidType,omitempty"`
}

func (s ReadCategoryGroupListRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadCategoryGroupListRequest) GoString() string {
	return s.String()
}

func (s *ReadCategoryGroupListRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ReadCategoryGroupListRequest) GetAppName() *string {
	return s.AppName
}

func (s *ReadCategoryGroupListRequest) GetBizName() *string {
	return s.BizName
}

func (s *ReadCategoryGroupListRequest) GetCallerProtocol() *string {
	return s.CallerProtocol
}

func (s *ReadCategoryGroupListRequest) GetChannelGroupCode() *string {
	return s.ChannelGroupCode
}

func (s *ReadCategoryGroupListRequest) GetClientSource() *string {
	return s.ClientSource
}

func (s *ReadCategoryGroupListRequest) GetCookies() *string {
	return s.Cookies
}

func (s *ReadCategoryGroupListRequest) GetSrcUrl() *string {
	return s.SrcUrl
}

func (s *ReadCategoryGroupListRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *ReadCategoryGroupListRequest) GetUidType() *string {
	return s.UidType
}

func (s *ReadCategoryGroupListRequest) SetAcceptLanguage(v string) *ReadCategoryGroupListRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReadCategoryGroupListRequest) SetAppName(v string) *ReadCategoryGroupListRequest {
	s.AppName = &v
	return s
}

func (s *ReadCategoryGroupListRequest) SetBizName(v string) *ReadCategoryGroupListRequest {
	s.BizName = &v
	return s
}

func (s *ReadCategoryGroupListRequest) SetCallerProtocol(v string) *ReadCategoryGroupListRequest {
	s.CallerProtocol = &v
	return s
}

func (s *ReadCategoryGroupListRequest) SetChannelGroupCode(v string) *ReadCategoryGroupListRequest {
	s.ChannelGroupCode = &v
	return s
}

func (s *ReadCategoryGroupListRequest) SetClientSource(v string) *ReadCategoryGroupListRequest {
	s.ClientSource = &v
	return s
}

func (s *ReadCategoryGroupListRequest) SetCookies(v string) *ReadCategoryGroupListRequest {
	s.Cookies = &v
	return s
}

func (s *ReadCategoryGroupListRequest) SetSrcUrl(v string) *ReadCategoryGroupListRequest {
	s.SrcUrl = &v
	return s
}

func (s *ReadCategoryGroupListRequest) SetTenantCode(v string) *ReadCategoryGroupListRequest {
	s.TenantCode = &v
	return s
}

func (s *ReadCategoryGroupListRequest) SetUidType(v string) *ReadCategoryGroupListRequest {
	s.UidType = &v
	return s
}

func (s *ReadCategoryGroupListRequest) Validate() error {
	return dara.Validate(s)
}
