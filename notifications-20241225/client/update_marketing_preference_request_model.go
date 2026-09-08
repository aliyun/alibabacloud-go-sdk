// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMarketingPreferenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateMarketingPreferenceRequest
	GetAcceptLanguage() *string
	SetAllowMarketing(v bool) *UpdateMarketingPreferenceRequest
	GetAllowMarketing() *bool
	SetAppName(v string) *UpdateMarketingPreferenceRequest
	GetAppName() *string
	SetBizName(v string) *UpdateMarketingPreferenceRequest
	GetBizName() *string
	SetCallerProtocol(v string) *UpdateMarketingPreferenceRequest
	GetCallerProtocol() *string
	SetClientSource(v string) *UpdateMarketingPreferenceRequest
	GetClientSource() *string
	SetContactId(v int64) *UpdateMarketingPreferenceRequest
	GetContactId() *int64
	SetCookies(v string) *UpdateMarketingPreferenceRequest
	GetCookies() *string
	SetSrcUrl(v string) *UpdateMarketingPreferenceRequest
	GetSrcUrl() *string
	SetTenantCode(v string) *UpdateMarketingPreferenceRequest
	GetTenantCode() *string
	SetUidType(v string) *UpdateMarketingPreferenceRequest
	GetUidType() *string
}

type UpdateMarketingPreferenceRequest struct {
	// The language.
	//
	// example:
	//
	// /
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// Specifies whether to allow notifications.
	//
	// example:
	//
	// true
	AllowMarketing *bool `json:"AllowMarketing,omitempty" xml:"AllowMarketing,omitempty"`
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
	// The source of the operation terminal.
	//
	// example:
	//
	// /
	ClientSource *string `json:"ClientSource,omitempty" xml:"ClientSource,omitempty"`
	// webhook id
	//
	// example:
	//
	// 0
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
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

func (s UpdateMarketingPreferenceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMarketingPreferenceRequest) GoString() string {
	return s.String()
}

func (s *UpdateMarketingPreferenceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateMarketingPreferenceRequest) GetAllowMarketing() *bool {
	return s.AllowMarketing
}

func (s *UpdateMarketingPreferenceRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateMarketingPreferenceRequest) GetBizName() *string {
	return s.BizName
}

func (s *UpdateMarketingPreferenceRequest) GetCallerProtocol() *string {
	return s.CallerProtocol
}

func (s *UpdateMarketingPreferenceRequest) GetClientSource() *string {
	return s.ClientSource
}

func (s *UpdateMarketingPreferenceRequest) GetContactId() *int64 {
	return s.ContactId
}

func (s *UpdateMarketingPreferenceRequest) GetCookies() *string {
	return s.Cookies
}

func (s *UpdateMarketingPreferenceRequest) GetSrcUrl() *string {
	return s.SrcUrl
}

func (s *UpdateMarketingPreferenceRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *UpdateMarketingPreferenceRequest) GetUidType() *string {
	return s.UidType
}

func (s *UpdateMarketingPreferenceRequest) SetAcceptLanguage(v string) *UpdateMarketingPreferenceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateMarketingPreferenceRequest) SetAllowMarketing(v bool) *UpdateMarketingPreferenceRequest {
	s.AllowMarketing = &v
	return s
}

func (s *UpdateMarketingPreferenceRequest) SetAppName(v string) *UpdateMarketingPreferenceRequest {
	s.AppName = &v
	return s
}

func (s *UpdateMarketingPreferenceRequest) SetBizName(v string) *UpdateMarketingPreferenceRequest {
	s.BizName = &v
	return s
}

func (s *UpdateMarketingPreferenceRequest) SetCallerProtocol(v string) *UpdateMarketingPreferenceRequest {
	s.CallerProtocol = &v
	return s
}

func (s *UpdateMarketingPreferenceRequest) SetClientSource(v string) *UpdateMarketingPreferenceRequest {
	s.ClientSource = &v
	return s
}

func (s *UpdateMarketingPreferenceRequest) SetContactId(v int64) *UpdateMarketingPreferenceRequest {
	s.ContactId = &v
	return s
}

func (s *UpdateMarketingPreferenceRequest) SetCookies(v string) *UpdateMarketingPreferenceRequest {
	s.Cookies = &v
	return s
}

func (s *UpdateMarketingPreferenceRequest) SetSrcUrl(v string) *UpdateMarketingPreferenceRequest {
	s.SrcUrl = &v
	return s
}

func (s *UpdateMarketingPreferenceRequest) SetTenantCode(v string) *UpdateMarketingPreferenceRequest {
	s.TenantCode = &v
	return s
}

func (s *UpdateMarketingPreferenceRequest) SetUidType(v string) *UpdateMarketingPreferenceRequest {
	s.UidType = &v
	return s
}

func (s *UpdateMarketingPreferenceRequest) Validate() error {
	return dara.Validate(s)
}
