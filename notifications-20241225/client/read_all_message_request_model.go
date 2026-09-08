// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadAllMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ReadAllMessageRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *ReadAllMessageRequest
	GetAppName() *string
	SetBizName(v string) *ReadAllMessageRequest
	GetBizName() *string
	SetCallerProtocol(v string) *ReadAllMessageRequest
	GetCallerProtocol() *string
	SetClassId(v int64) *ReadAllMessageRequest
	GetClassId() *int64
	SetClientSource(v string) *ReadAllMessageRequest
	GetClientSource() *string
	SetCookies(v string) *ReadAllMessageRequest
	GetCookies() *string
	SetGroupCode(v string) *ReadAllMessageRequest
	GetGroupCode() *string
	SetSrcUrl(v string) *ReadAllMessageRequest
	GetSrcUrl() *string
	SetTenantCode(v string) *ReadAllMessageRequest
	GetTenantCode() *string
	SetUidType(v string) *ReadAllMessageRequest
	GetUidType() *string
}

type ReadAllMessageRequest struct {
	// The language. Default value: Simplified Chinese.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// A system parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// /
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// A system parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// /
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// A system parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// /
	CallerProtocol *string `json:"CallerProtocol,omitempty" xml:"CallerProtocol,omitempty"`
	// The message category ID.
	//
	// example:
	//
	// 1
	ClassId *int64 `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	// A system parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// /
	ClientSource *string `json:"ClientSource,omitempty" xml:"ClientSource,omitempty"`
	// A system parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// []
	Cookies *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
	// The group code.
	//
	// example:
	//
	// test
	GroupCode *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
	// A system parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// /
	SrcUrl *string `json:"SrcUrl,omitempty" xml:"SrcUrl,omitempty"`
	// A system parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// /
	TenantCode *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
	// A system parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// /
	UidType *string `json:"UidType,omitempty" xml:"UidType,omitempty"`
}

func (s ReadAllMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadAllMessageRequest) GoString() string {
	return s.String()
}

func (s *ReadAllMessageRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ReadAllMessageRequest) GetAppName() *string {
	return s.AppName
}

func (s *ReadAllMessageRequest) GetBizName() *string {
	return s.BizName
}

func (s *ReadAllMessageRequest) GetCallerProtocol() *string {
	return s.CallerProtocol
}

func (s *ReadAllMessageRequest) GetClassId() *int64 {
	return s.ClassId
}

func (s *ReadAllMessageRequest) GetClientSource() *string {
	return s.ClientSource
}

func (s *ReadAllMessageRequest) GetCookies() *string {
	return s.Cookies
}

func (s *ReadAllMessageRequest) GetGroupCode() *string {
	return s.GroupCode
}

func (s *ReadAllMessageRequest) GetSrcUrl() *string {
	return s.SrcUrl
}

func (s *ReadAllMessageRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *ReadAllMessageRequest) GetUidType() *string {
	return s.UidType
}

func (s *ReadAllMessageRequest) SetAcceptLanguage(v string) *ReadAllMessageRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReadAllMessageRequest) SetAppName(v string) *ReadAllMessageRequest {
	s.AppName = &v
	return s
}

func (s *ReadAllMessageRequest) SetBizName(v string) *ReadAllMessageRequest {
	s.BizName = &v
	return s
}

func (s *ReadAllMessageRequest) SetCallerProtocol(v string) *ReadAllMessageRequest {
	s.CallerProtocol = &v
	return s
}

func (s *ReadAllMessageRequest) SetClassId(v int64) *ReadAllMessageRequest {
	s.ClassId = &v
	return s
}

func (s *ReadAllMessageRequest) SetClientSource(v string) *ReadAllMessageRequest {
	s.ClientSource = &v
	return s
}

func (s *ReadAllMessageRequest) SetCookies(v string) *ReadAllMessageRequest {
	s.Cookies = &v
	return s
}

func (s *ReadAllMessageRequest) SetGroupCode(v string) *ReadAllMessageRequest {
	s.GroupCode = &v
	return s
}

func (s *ReadAllMessageRequest) SetSrcUrl(v string) *ReadAllMessageRequest {
	s.SrcUrl = &v
	return s
}

func (s *ReadAllMessageRequest) SetTenantCode(v string) *ReadAllMessageRequest {
	s.TenantCode = &v
	return s
}

func (s *ReadAllMessageRequest) SetUidType(v string) *ReadAllMessageRequest {
	s.UidType = &v
	return s
}

func (s *ReadAllMessageRequest) Validate() error {
	return dara.Validate(s)
}
