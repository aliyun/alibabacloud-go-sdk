// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadCommonContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ReadCommonContactRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *ReadCommonContactRequest
	GetAppName() *string
	SetBizName(v string) *ReadCommonContactRequest
	GetBizName() *string
	SetCallerProtocol(v string) *ReadCommonContactRequest
	GetCallerProtocol() *string
	SetClientSource(v string) *ReadCommonContactRequest
	GetClientSource() *string
	SetContactId(v int64) *ReadCommonContactRequest
	GetContactId() *int64
	SetCookies(v string) *ReadCommonContactRequest
	GetCookies() *string
	SetSrcUrl(v string) *ReadCommonContactRequest
	GetSrcUrl() *string
	SetTenantCode(v string) *ReadCommonContactRequest
	GetTenantCode() *string
	SetUidType(v string) *ReadCommonContactRequest
	GetUidType() *string
}

type ReadCommonContactRequest struct {
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
	// The user cookies.
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

func (s ReadCommonContactRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadCommonContactRequest) GoString() string {
	return s.String()
}

func (s *ReadCommonContactRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ReadCommonContactRequest) GetAppName() *string {
	return s.AppName
}

func (s *ReadCommonContactRequest) GetBizName() *string {
	return s.BizName
}

func (s *ReadCommonContactRequest) GetCallerProtocol() *string {
	return s.CallerProtocol
}

func (s *ReadCommonContactRequest) GetClientSource() *string {
	return s.ClientSource
}

func (s *ReadCommonContactRequest) GetContactId() *int64 {
	return s.ContactId
}

func (s *ReadCommonContactRequest) GetCookies() *string {
	return s.Cookies
}

func (s *ReadCommonContactRequest) GetSrcUrl() *string {
	return s.SrcUrl
}

func (s *ReadCommonContactRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *ReadCommonContactRequest) GetUidType() *string {
	return s.UidType
}

func (s *ReadCommonContactRequest) SetAcceptLanguage(v string) *ReadCommonContactRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReadCommonContactRequest) SetAppName(v string) *ReadCommonContactRequest {
	s.AppName = &v
	return s
}

func (s *ReadCommonContactRequest) SetBizName(v string) *ReadCommonContactRequest {
	s.BizName = &v
	return s
}

func (s *ReadCommonContactRequest) SetCallerProtocol(v string) *ReadCommonContactRequest {
	s.CallerProtocol = &v
	return s
}

func (s *ReadCommonContactRequest) SetClientSource(v string) *ReadCommonContactRequest {
	s.ClientSource = &v
	return s
}

func (s *ReadCommonContactRequest) SetContactId(v int64) *ReadCommonContactRequest {
	s.ContactId = &v
	return s
}

func (s *ReadCommonContactRequest) SetCookies(v string) *ReadCommonContactRequest {
	s.Cookies = &v
	return s
}

func (s *ReadCommonContactRequest) SetSrcUrl(v string) *ReadCommonContactRequest {
	s.SrcUrl = &v
	return s
}

func (s *ReadCommonContactRequest) SetTenantCode(v string) *ReadCommonContactRequest {
	s.TenantCode = &v
	return s
}

func (s *ReadCommonContactRequest) SetUidType(v string) *ReadCommonContactRequest {
	s.UidType = &v
	return s
}

func (s *ReadCommonContactRequest) Validate() error {
	return dara.Validate(s)
}
