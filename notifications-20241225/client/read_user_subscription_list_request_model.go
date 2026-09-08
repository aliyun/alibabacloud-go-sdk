// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadUserSubscriptionListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ReadUserSubscriptionListRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *ReadUserSubscriptionListRequest
	GetAppName() *string
	SetBizName(v string) *ReadUserSubscriptionListRequest
	GetBizName() *string
	SetCallerProtocol(v string) *ReadUserSubscriptionListRequest
	GetCallerProtocol() *string
	SetCategoryGroupCode(v string) *ReadUserSubscriptionListRequest
	GetCategoryGroupCode() *string
	SetChannelGroupCode(v string) *ReadUserSubscriptionListRequest
	GetChannelGroupCode() *string
	SetClientSource(v string) *ReadUserSubscriptionListRequest
	GetClientSource() *string
	SetCookies(v string) *ReadUserSubscriptionListRequest
	GetCookies() *string
	SetSrcUrl(v string) *ReadUserSubscriptionListRequest
	GetSrcUrl() *string
	SetTenantCode(v string) *ReadUserSubscriptionListRequest
	GetTenantCode() *string
	SetUidType(v string) *ReadUserSubscriptionListRequest
	GetUidType() *string
}

type ReadUserSubscriptionListRequest struct {
	// The language.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The application project name of the caller.
	//
	// example:
	//
	// /
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
	// The category group code.
	//
	// example:
	//
	// prod_msg
	CategoryGroupCode *string `json:"CategoryGroupCode,omitempty" xml:"CategoryGroupCode,omitempty"`
	// The channel group.
	//
	// example:
	//
	// base
	ChannelGroupCode *string `json:"ChannelGroupCode,omitempty" xml:"ChannelGroupCode,omitempty"`
	// The source of the operation terminal.
	//
	// example:
	//
	// /
	ClientSource *string `json:"ClientSource,omitempty" xml:"ClientSource,omitempty"`
	// The cookies of the user.
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

func (s ReadUserSubscriptionListRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadUserSubscriptionListRequest) GoString() string {
	return s.String()
}

func (s *ReadUserSubscriptionListRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ReadUserSubscriptionListRequest) GetAppName() *string {
	return s.AppName
}

func (s *ReadUserSubscriptionListRequest) GetBizName() *string {
	return s.BizName
}

func (s *ReadUserSubscriptionListRequest) GetCallerProtocol() *string {
	return s.CallerProtocol
}

func (s *ReadUserSubscriptionListRequest) GetCategoryGroupCode() *string {
	return s.CategoryGroupCode
}

func (s *ReadUserSubscriptionListRequest) GetChannelGroupCode() *string {
	return s.ChannelGroupCode
}

func (s *ReadUserSubscriptionListRequest) GetClientSource() *string {
	return s.ClientSource
}

func (s *ReadUserSubscriptionListRequest) GetCookies() *string {
	return s.Cookies
}

func (s *ReadUserSubscriptionListRequest) GetSrcUrl() *string {
	return s.SrcUrl
}

func (s *ReadUserSubscriptionListRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *ReadUserSubscriptionListRequest) GetUidType() *string {
	return s.UidType
}

func (s *ReadUserSubscriptionListRequest) SetAcceptLanguage(v string) *ReadUserSubscriptionListRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReadUserSubscriptionListRequest) SetAppName(v string) *ReadUserSubscriptionListRequest {
	s.AppName = &v
	return s
}

func (s *ReadUserSubscriptionListRequest) SetBizName(v string) *ReadUserSubscriptionListRequest {
	s.BizName = &v
	return s
}

func (s *ReadUserSubscriptionListRequest) SetCallerProtocol(v string) *ReadUserSubscriptionListRequest {
	s.CallerProtocol = &v
	return s
}

func (s *ReadUserSubscriptionListRequest) SetCategoryGroupCode(v string) *ReadUserSubscriptionListRequest {
	s.CategoryGroupCode = &v
	return s
}

func (s *ReadUserSubscriptionListRequest) SetChannelGroupCode(v string) *ReadUserSubscriptionListRequest {
	s.ChannelGroupCode = &v
	return s
}

func (s *ReadUserSubscriptionListRequest) SetClientSource(v string) *ReadUserSubscriptionListRequest {
	s.ClientSource = &v
	return s
}

func (s *ReadUserSubscriptionListRequest) SetCookies(v string) *ReadUserSubscriptionListRequest {
	s.Cookies = &v
	return s
}

func (s *ReadUserSubscriptionListRequest) SetSrcUrl(v string) *ReadUserSubscriptionListRequest {
	s.SrcUrl = &v
	return s
}

func (s *ReadUserSubscriptionListRequest) SetTenantCode(v string) *ReadUserSubscriptionListRequest {
	s.TenantCode = &v
	return s
}

func (s *ReadUserSubscriptionListRequest) SetUidType(v string) *ReadUserSubscriptionListRequest {
	s.UidType = &v
	return s
}

func (s *ReadUserSubscriptionListRequest) Validate() error {
	return dara.Validate(s)
}
