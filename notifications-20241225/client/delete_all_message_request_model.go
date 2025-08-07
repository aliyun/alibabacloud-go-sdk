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
	SetSrcUrl(v string) *DeleteAllMessageRequest
	GetSrcUrl() *string
	SetTenantCode(v string) *DeleteAllMessageRequest
	GetTenantCode() *string
	SetUidType(v string) *DeleteAllMessageRequest
	GetUidType() *string
}

type DeleteAllMessageRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BizName        *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	CallerProtocol *string `json:"CallerProtocol,omitempty" xml:"CallerProtocol,omitempty"`
	ClassId        *int64  `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	ClientSource   *string `json:"ClientSource,omitempty" xml:"ClientSource,omitempty"`
	Cookies        *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
	SrcUrl         *string `json:"SrcUrl,omitempty" xml:"SrcUrl,omitempty"`
	TenantCode     *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
	UidType        *string `json:"UidType,omitempty" xml:"UidType,omitempty"`
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
