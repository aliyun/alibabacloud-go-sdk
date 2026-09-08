// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadMarketingPreferenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ReadMarketingPreferenceRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *ReadMarketingPreferenceRequest
	GetAppName() *string
	SetBizName(v string) *ReadMarketingPreferenceRequest
	GetBizName() *string
	SetCallerProtocol(v string) *ReadMarketingPreferenceRequest
	GetCallerProtocol() *string
	SetClientSource(v string) *ReadMarketingPreferenceRequest
	GetClientSource() *string
	SetContactId(v int64) *ReadMarketingPreferenceRequest
	GetContactId() *int64
	SetCookies(v string) *ReadMarketingPreferenceRequest
	GetCookies() *string
	SetSrcUrl(v string) *ReadMarketingPreferenceRequest
	GetSrcUrl() *string
	SetTenantCode(v string) *ReadMarketingPreferenceRequest
	GetTenantCode() *string
	SetUidType(v string) *ReadMarketingPreferenceRequest
	GetUidType() *string
}

type ReadMarketingPreferenceRequest struct {
	// The language.
	//
	// example:
	//
	// /
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The application name of the requester.
	//
	// example:
	//
	// /
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The business line of the requester.
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
	// The contact ID.
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

func (s ReadMarketingPreferenceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadMarketingPreferenceRequest) GoString() string {
	return s.String()
}

func (s *ReadMarketingPreferenceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ReadMarketingPreferenceRequest) GetAppName() *string {
	return s.AppName
}

func (s *ReadMarketingPreferenceRequest) GetBizName() *string {
	return s.BizName
}

func (s *ReadMarketingPreferenceRequest) GetCallerProtocol() *string {
	return s.CallerProtocol
}

func (s *ReadMarketingPreferenceRequest) GetClientSource() *string {
	return s.ClientSource
}

func (s *ReadMarketingPreferenceRequest) GetContactId() *int64 {
	return s.ContactId
}

func (s *ReadMarketingPreferenceRequest) GetCookies() *string {
	return s.Cookies
}

func (s *ReadMarketingPreferenceRequest) GetSrcUrl() *string {
	return s.SrcUrl
}

func (s *ReadMarketingPreferenceRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *ReadMarketingPreferenceRequest) GetUidType() *string {
	return s.UidType
}

func (s *ReadMarketingPreferenceRequest) SetAcceptLanguage(v string) *ReadMarketingPreferenceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReadMarketingPreferenceRequest) SetAppName(v string) *ReadMarketingPreferenceRequest {
	s.AppName = &v
	return s
}

func (s *ReadMarketingPreferenceRequest) SetBizName(v string) *ReadMarketingPreferenceRequest {
	s.BizName = &v
	return s
}

func (s *ReadMarketingPreferenceRequest) SetCallerProtocol(v string) *ReadMarketingPreferenceRequest {
	s.CallerProtocol = &v
	return s
}

func (s *ReadMarketingPreferenceRequest) SetClientSource(v string) *ReadMarketingPreferenceRequest {
	s.ClientSource = &v
	return s
}

func (s *ReadMarketingPreferenceRequest) SetContactId(v int64) *ReadMarketingPreferenceRequest {
	s.ContactId = &v
	return s
}

func (s *ReadMarketingPreferenceRequest) SetCookies(v string) *ReadMarketingPreferenceRequest {
	s.Cookies = &v
	return s
}

func (s *ReadMarketingPreferenceRequest) SetSrcUrl(v string) *ReadMarketingPreferenceRequest {
	s.SrcUrl = &v
	return s
}

func (s *ReadMarketingPreferenceRequest) SetTenantCode(v string) *ReadMarketingPreferenceRequest {
	s.TenantCode = &v
	return s
}

func (s *ReadMarketingPreferenceRequest) SetUidType(v string) *ReadMarketingPreferenceRequest {
	s.UidType = &v
	return s
}

func (s *ReadMarketingPreferenceRequest) Validate() error {
	return dara.Validate(s)
}
