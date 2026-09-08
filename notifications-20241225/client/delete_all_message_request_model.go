// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAllMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteAllMessageRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *DeleteAllMessageRequest
	GetAppName() *string
	SetBizName(v string) *DeleteAllMessageRequest
	GetBizName() *string
	SetCallerProtocol(v string) *DeleteAllMessageRequest
	GetCallerProtocol() *string
	SetClassId(v int64) *DeleteAllMessageRequest
	GetClassId() *int64
	SetClientSource(v string) *DeleteAllMessageRequest
	GetClientSource() *string
	SetCookies(v string) *DeleteAllMessageRequest
	GetCookies() *string
	SetGroupCode(v string) *DeleteAllMessageRequest
	GetGroupCode() *string
	SetSrcUrl(v string) *DeleteAllMessageRequest
	GetSrcUrl() *string
	SetTenantCode(v string) *DeleteAllMessageRequest
	GetTenantCode() *string
	SetUidType(v string) *DeleteAllMessageRequest
	GetUidType() *string
}

type DeleteAllMessageRequest struct {
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
	// /
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

func (s DeleteAllMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAllMessageRequest) GoString() string {
	return s.String()
}

func (s *DeleteAllMessageRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteAllMessageRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteAllMessageRequest) GetBizName() *string {
	return s.BizName
}

func (s *DeleteAllMessageRequest) GetCallerProtocol() *string {
	return s.CallerProtocol
}

func (s *DeleteAllMessageRequest) GetClassId() *int64 {
	return s.ClassId
}

func (s *DeleteAllMessageRequest) GetClientSource() *string {
	return s.ClientSource
}

func (s *DeleteAllMessageRequest) GetCookies() *string {
	return s.Cookies
}

func (s *DeleteAllMessageRequest) GetGroupCode() *string {
	return s.GroupCode
}

func (s *DeleteAllMessageRequest) GetSrcUrl() *string {
	return s.SrcUrl
}

func (s *DeleteAllMessageRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *DeleteAllMessageRequest) GetUidType() *string {
	return s.UidType
}

func (s *DeleteAllMessageRequest) SetAcceptLanguage(v string) *DeleteAllMessageRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteAllMessageRequest) SetAppName(v string) *DeleteAllMessageRequest {
	s.AppName = &v
	return s
}

func (s *DeleteAllMessageRequest) SetBizName(v string) *DeleteAllMessageRequest {
	s.BizName = &v
	return s
}

func (s *DeleteAllMessageRequest) SetCallerProtocol(v string) *DeleteAllMessageRequest {
	s.CallerProtocol = &v
	return s
}

func (s *DeleteAllMessageRequest) SetClassId(v int64) *DeleteAllMessageRequest {
	s.ClassId = &v
	return s
}

func (s *DeleteAllMessageRequest) SetClientSource(v string) *DeleteAllMessageRequest {
	s.ClientSource = &v
	return s
}

func (s *DeleteAllMessageRequest) SetCookies(v string) *DeleteAllMessageRequest {
	s.Cookies = &v
	return s
}

func (s *DeleteAllMessageRequest) SetGroupCode(v string) *DeleteAllMessageRequest {
	s.GroupCode = &v
	return s
}

func (s *DeleteAllMessageRequest) SetSrcUrl(v string) *DeleteAllMessageRequest {
	s.SrcUrl = &v
	return s
}

func (s *DeleteAllMessageRequest) SetTenantCode(v string) *DeleteAllMessageRequest {
	s.TenantCode = &v
	return s
}

func (s *DeleteAllMessageRequest) SetUidType(v string) *DeleteAllMessageRequest {
	s.UidType = &v
	return s
}

func (s *DeleteAllMessageRequest) Validate() error {
	return dara.Validate(s)
}
