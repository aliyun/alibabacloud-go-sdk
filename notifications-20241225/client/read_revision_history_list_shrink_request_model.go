// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadRevisionHistoryListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ReadRevisionHistoryListShrinkRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *ReadRevisionHistoryListShrinkRequest
	GetAppName() *string
	SetBizName(v string) *ReadRevisionHistoryListShrinkRequest
	GetBizName() *string
	SetCallerProtocol(v string) *ReadRevisionHistoryListShrinkRequest
	GetCallerProtocol() *string
	SetCategoryCode(v string) *ReadRevisionHistoryListShrinkRequest
	GetCategoryCode() *string
	SetChannelGroupCode(v string) *ReadRevisionHistoryListShrinkRequest
	GetChannelGroupCode() *string
	SetClientSource(v string) *ReadRevisionHistoryListShrinkRequest
	GetClientSource() *string
	SetCookies(v string) *ReadRevisionHistoryListShrinkRequest
	GetCookies() *string
	SetPageInfoShrink(v string) *ReadRevisionHistoryListShrinkRequest
	GetPageInfoShrink() *string
	SetSrcUrl(v string) *ReadRevisionHistoryListShrinkRequest
	GetSrcUrl() *string
	SetTenantCode(v string) *ReadRevisionHistoryListShrinkRequest
	GetTenantCode() *string
	SetUidType(v string) *ReadRevisionHistoryListShrinkRequest
	GetUidType() *string
}

type ReadRevisionHistoryListShrinkRequest struct {
	// The language. Automatically passed through by the browser. You can manually override this value.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// Ignored. No need to pass this parameter. The application name of the caller.
	//
	// example:
	//
	// /
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Ignored. No need to pass this parameter. The business line of the caller.
	//
	// example:
	//
	// /
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// Ignored. No need to pass this parameter. The request protocol type.
	//
	// example:
	//
	// /
	CallerProtocol *string `json:"CallerProtocol,omitempty" xml:"CallerProtocol,omitempty"`
	// The category code.
	//
	// example:
	//
	// prod_edu_content
	CategoryCode *string `json:"CategoryCode,omitempty" xml:"CategoryCode,omitempty"`
	// The channel group.
	//
	// example:
	//
	// base
	ChannelGroupCode *string `json:"ChannelGroupCode,omitempty" xml:"ChannelGroupCode,omitempty"`
	// Ignored. No need to pass this parameter. The source of the operation terminal.
	//
	// example:
	//
	// /
	ClientSource *string `json:"ClientSource,omitempty" xml:"ClientSource,omitempty"`
	// Ignored. No need to pass this parameter. The user cookies.
	//
	// example:
	//
	// /
	Cookies *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
	// The pagination information.
	PageInfoShrink *string `json:"PageInfo,omitempty" xml:"PageInfo,omitempty"`
	// Ignored. No need to pass this parameter. The source page URL.
	//
	// example:
	//
	// /
	SrcUrl *string `json:"SrcUrl,omitempty" xml:"SrcUrl,omitempty"`
	// Ignored. No need to pass this parameter. The tenant information.
	//
	// example:
	//
	// /
	TenantCode *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
	// Ignored. No need to pass this parameter. The user type.
	//
	// example:
	//
	// /
	UidType *string `json:"UidType,omitempty" xml:"UidType,omitempty"`
}

func (s ReadRevisionHistoryListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadRevisionHistoryListShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReadRevisionHistoryListShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ReadRevisionHistoryListShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *ReadRevisionHistoryListShrinkRequest) GetBizName() *string {
	return s.BizName
}

func (s *ReadRevisionHistoryListShrinkRequest) GetCallerProtocol() *string {
	return s.CallerProtocol
}

func (s *ReadRevisionHistoryListShrinkRequest) GetCategoryCode() *string {
	return s.CategoryCode
}

func (s *ReadRevisionHistoryListShrinkRequest) GetChannelGroupCode() *string {
	return s.ChannelGroupCode
}

func (s *ReadRevisionHistoryListShrinkRequest) GetClientSource() *string {
	return s.ClientSource
}

func (s *ReadRevisionHistoryListShrinkRequest) GetCookies() *string {
	return s.Cookies
}

func (s *ReadRevisionHistoryListShrinkRequest) GetPageInfoShrink() *string {
	return s.PageInfoShrink
}

func (s *ReadRevisionHistoryListShrinkRequest) GetSrcUrl() *string {
	return s.SrcUrl
}

func (s *ReadRevisionHistoryListShrinkRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *ReadRevisionHistoryListShrinkRequest) GetUidType() *string {
	return s.UidType
}

func (s *ReadRevisionHistoryListShrinkRequest) SetAcceptLanguage(v string) *ReadRevisionHistoryListShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReadRevisionHistoryListShrinkRequest) SetAppName(v string) *ReadRevisionHistoryListShrinkRequest {
	s.AppName = &v
	return s
}

func (s *ReadRevisionHistoryListShrinkRequest) SetBizName(v string) *ReadRevisionHistoryListShrinkRequest {
	s.BizName = &v
	return s
}

func (s *ReadRevisionHistoryListShrinkRequest) SetCallerProtocol(v string) *ReadRevisionHistoryListShrinkRequest {
	s.CallerProtocol = &v
	return s
}

func (s *ReadRevisionHistoryListShrinkRequest) SetCategoryCode(v string) *ReadRevisionHistoryListShrinkRequest {
	s.CategoryCode = &v
	return s
}

func (s *ReadRevisionHistoryListShrinkRequest) SetChannelGroupCode(v string) *ReadRevisionHistoryListShrinkRequest {
	s.ChannelGroupCode = &v
	return s
}

func (s *ReadRevisionHistoryListShrinkRequest) SetClientSource(v string) *ReadRevisionHistoryListShrinkRequest {
	s.ClientSource = &v
	return s
}

func (s *ReadRevisionHistoryListShrinkRequest) SetCookies(v string) *ReadRevisionHistoryListShrinkRequest {
	s.Cookies = &v
	return s
}

func (s *ReadRevisionHistoryListShrinkRequest) SetPageInfoShrink(v string) *ReadRevisionHistoryListShrinkRequest {
	s.PageInfoShrink = &v
	return s
}

func (s *ReadRevisionHistoryListShrinkRequest) SetSrcUrl(v string) *ReadRevisionHistoryListShrinkRequest {
	s.SrcUrl = &v
	return s
}

func (s *ReadRevisionHistoryListShrinkRequest) SetTenantCode(v string) *ReadRevisionHistoryListShrinkRequest {
	s.TenantCode = &v
	return s
}

func (s *ReadRevisionHistoryListShrinkRequest) SetUidType(v string) *ReadRevisionHistoryListShrinkRequest {
	s.UidType = &v
	return s
}

func (s *ReadRevisionHistoryListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
