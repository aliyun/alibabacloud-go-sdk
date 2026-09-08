// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetUserSubscriptionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ResetUserSubscriptionShrinkRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *ResetUserSubscriptionShrinkRequest
	GetAppName() *string
	SetBizName(v string) *ResetUserSubscriptionShrinkRequest
	GetBizName() *string
	SetCallerProtocol(v string) *ResetUserSubscriptionShrinkRequest
	GetCallerProtocol() *string
	SetCategoryCodesShrink(v string) *ResetUserSubscriptionShrinkRequest
	GetCategoryCodesShrink() *string
	SetChannelGroupCode(v string) *ResetUserSubscriptionShrinkRequest
	GetChannelGroupCode() *string
	SetClientSource(v string) *ResetUserSubscriptionShrinkRequest
	GetClientSource() *string
	SetCookies(v string) *ResetUserSubscriptionShrinkRequest
	GetCookies() *string
	SetRemarks(v string) *ResetUserSubscriptionShrinkRequest
	GetRemarks() *string
	SetSrcUrl(v string) *ResetUserSubscriptionShrinkRequest
	GetSrcUrl() *string
	SetTenantCode(v string) *ResetUserSubscriptionShrinkRequest
	GetTenantCode() *string
	SetUidType(v string) *ResetUserSubscriptionShrinkRequest
	GetUidType() *string
}

type ResetUserSubscriptionShrinkRequest struct {
	// The language.
	//
	// example:
	//
	// /
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The application name of the caller.
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
	// The list of category codes.
	CategoryCodesShrink *string `json:"CategoryCodes,omitempty" xml:"CategoryCodes,omitempty"`
	// The channel group. Valid values:
	//
	// - tts: Voice reception management.
	//
	// - webhook: Bot reception management.
	//
	// - base: Basic reception management.
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
	// The user cookies.
	//
	// example:
	//
	// /
	Cookies *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
	// The remarks.
	//
	// example:
	//
	// /
	Remarks *string `json:"Remarks,omitempty" xml:"Remarks,omitempty"`
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

func (s ResetUserSubscriptionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetUserSubscriptionShrinkRequest) GoString() string {
	return s.String()
}

func (s *ResetUserSubscriptionShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ResetUserSubscriptionShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *ResetUserSubscriptionShrinkRequest) GetBizName() *string {
	return s.BizName
}

func (s *ResetUserSubscriptionShrinkRequest) GetCallerProtocol() *string {
	return s.CallerProtocol
}

func (s *ResetUserSubscriptionShrinkRequest) GetCategoryCodesShrink() *string {
	return s.CategoryCodesShrink
}

func (s *ResetUserSubscriptionShrinkRequest) GetChannelGroupCode() *string {
	return s.ChannelGroupCode
}

func (s *ResetUserSubscriptionShrinkRequest) GetClientSource() *string {
	return s.ClientSource
}

func (s *ResetUserSubscriptionShrinkRequest) GetCookies() *string {
	return s.Cookies
}

func (s *ResetUserSubscriptionShrinkRequest) GetRemarks() *string {
	return s.Remarks
}

func (s *ResetUserSubscriptionShrinkRequest) GetSrcUrl() *string {
	return s.SrcUrl
}

func (s *ResetUserSubscriptionShrinkRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *ResetUserSubscriptionShrinkRequest) GetUidType() *string {
	return s.UidType
}

func (s *ResetUserSubscriptionShrinkRequest) SetAcceptLanguage(v string) *ResetUserSubscriptionShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ResetUserSubscriptionShrinkRequest) SetAppName(v string) *ResetUserSubscriptionShrinkRequest {
	s.AppName = &v
	return s
}

func (s *ResetUserSubscriptionShrinkRequest) SetBizName(v string) *ResetUserSubscriptionShrinkRequest {
	s.BizName = &v
	return s
}

func (s *ResetUserSubscriptionShrinkRequest) SetCallerProtocol(v string) *ResetUserSubscriptionShrinkRequest {
	s.CallerProtocol = &v
	return s
}

func (s *ResetUserSubscriptionShrinkRequest) SetCategoryCodesShrink(v string) *ResetUserSubscriptionShrinkRequest {
	s.CategoryCodesShrink = &v
	return s
}

func (s *ResetUserSubscriptionShrinkRequest) SetChannelGroupCode(v string) *ResetUserSubscriptionShrinkRequest {
	s.ChannelGroupCode = &v
	return s
}

func (s *ResetUserSubscriptionShrinkRequest) SetClientSource(v string) *ResetUserSubscriptionShrinkRequest {
	s.ClientSource = &v
	return s
}

func (s *ResetUserSubscriptionShrinkRequest) SetCookies(v string) *ResetUserSubscriptionShrinkRequest {
	s.Cookies = &v
	return s
}

func (s *ResetUserSubscriptionShrinkRequest) SetRemarks(v string) *ResetUserSubscriptionShrinkRequest {
	s.Remarks = &v
	return s
}

func (s *ResetUserSubscriptionShrinkRequest) SetSrcUrl(v string) *ResetUserSubscriptionShrinkRequest {
	s.SrcUrl = &v
	return s
}

func (s *ResetUserSubscriptionShrinkRequest) SetTenantCode(v string) *ResetUserSubscriptionShrinkRequest {
	s.TenantCode = &v
	return s
}

func (s *ResetUserSubscriptionShrinkRequest) SetUidType(v string) *ResetUserSubscriptionShrinkRequest {
	s.UidType = &v
	return s
}

func (s *ResetUserSubscriptionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
