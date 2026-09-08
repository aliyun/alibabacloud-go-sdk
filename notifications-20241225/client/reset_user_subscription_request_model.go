// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetUserSubscriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ResetUserSubscriptionRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *ResetUserSubscriptionRequest
	GetAppName() *string
	SetBizName(v string) *ResetUserSubscriptionRequest
	GetBizName() *string
	SetCallerProtocol(v string) *ResetUserSubscriptionRequest
	GetCallerProtocol() *string
	SetCategoryCodes(v []*string) *ResetUserSubscriptionRequest
	GetCategoryCodes() []*string
	SetChannelGroupCode(v string) *ResetUserSubscriptionRequest
	GetChannelGroupCode() *string
	SetClientSource(v string) *ResetUserSubscriptionRequest
	GetClientSource() *string
	SetCookies(v string) *ResetUserSubscriptionRequest
	GetCookies() *string
	SetRemarks(v string) *ResetUserSubscriptionRequest
	GetRemarks() *string
	SetSrcUrl(v string) *ResetUserSubscriptionRequest
	GetSrcUrl() *string
	SetTenantCode(v string) *ResetUserSubscriptionRequest
	GetTenantCode() *string
	SetUidType(v string) *ResetUserSubscriptionRequest
	GetUidType() *string
}

type ResetUserSubscriptionRequest struct {
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
	CategoryCodes []*string `json:"CategoryCodes,omitempty" xml:"CategoryCodes,omitempty" type:"Repeated"`
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

func (s ResetUserSubscriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetUserSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *ResetUserSubscriptionRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ResetUserSubscriptionRequest) GetAppName() *string {
	return s.AppName
}

func (s *ResetUserSubscriptionRequest) GetBizName() *string {
	return s.BizName
}

func (s *ResetUserSubscriptionRequest) GetCallerProtocol() *string {
	return s.CallerProtocol
}

func (s *ResetUserSubscriptionRequest) GetCategoryCodes() []*string {
	return s.CategoryCodes
}

func (s *ResetUserSubscriptionRequest) GetChannelGroupCode() *string {
	return s.ChannelGroupCode
}

func (s *ResetUserSubscriptionRequest) GetClientSource() *string {
	return s.ClientSource
}

func (s *ResetUserSubscriptionRequest) GetCookies() *string {
	return s.Cookies
}

func (s *ResetUserSubscriptionRequest) GetRemarks() *string {
	return s.Remarks
}

func (s *ResetUserSubscriptionRequest) GetSrcUrl() *string {
	return s.SrcUrl
}

func (s *ResetUserSubscriptionRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *ResetUserSubscriptionRequest) GetUidType() *string {
	return s.UidType
}

func (s *ResetUserSubscriptionRequest) SetAcceptLanguage(v string) *ResetUserSubscriptionRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ResetUserSubscriptionRequest) SetAppName(v string) *ResetUserSubscriptionRequest {
	s.AppName = &v
	return s
}

func (s *ResetUserSubscriptionRequest) SetBizName(v string) *ResetUserSubscriptionRequest {
	s.BizName = &v
	return s
}

func (s *ResetUserSubscriptionRequest) SetCallerProtocol(v string) *ResetUserSubscriptionRequest {
	s.CallerProtocol = &v
	return s
}

func (s *ResetUserSubscriptionRequest) SetCategoryCodes(v []*string) *ResetUserSubscriptionRequest {
	s.CategoryCodes = v
	return s
}

func (s *ResetUserSubscriptionRequest) SetChannelGroupCode(v string) *ResetUserSubscriptionRequest {
	s.ChannelGroupCode = &v
	return s
}

func (s *ResetUserSubscriptionRequest) SetClientSource(v string) *ResetUserSubscriptionRequest {
	s.ClientSource = &v
	return s
}

func (s *ResetUserSubscriptionRequest) SetCookies(v string) *ResetUserSubscriptionRequest {
	s.Cookies = &v
	return s
}

func (s *ResetUserSubscriptionRequest) SetRemarks(v string) *ResetUserSubscriptionRequest {
	s.Remarks = &v
	return s
}

func (s *ResetUserSubscriptionRequest) SetSrcUrl(v string) *ResetUserSubscriptionRequest {
	s.SrcUrl = &v
	return s
}

func (s *ResetUserSubscriptionRequest) SetTenantCode(v string) *ResetUserSubscriptionRequest {
	s.TenantCode = &v
	return s
}

func (s *ResetUserSubscriptionRequest) SetUidType(v string) *ResetUserSubscriptionRequest {
	s.UidType = &v
	return s
}

func (s *ResetUserSubscriptionRequest) Validate() error {
	return dara.Validate(s)
}
