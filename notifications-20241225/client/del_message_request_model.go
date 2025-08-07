// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDelMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DelMessageRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *DelMessageRequest
	GetAppName() *string
	SetBizName(v string) *DelMessageRequest
	GetBizName() *string
	SetCallerProtocol(v string) *DelMessageRequest
	GetCallerProtocol() *string
	SetClientSource(v string) *DelMessageRequest
	GetClientSource() *string
	SetCookies(v string) *DelMessageRequest
	GetCookies() *string
	SetMsgId(v string) *DelMessageRequest
	GetMsgId() *string
	SetSrcUrl(v string) *DelMessageRequest
	GetSrcUrl() *string
	SetTenantCode(v string) *DelMessageRequest
	GetTenantCode() *string
	SetUidType(v string) *DelMessageRequest
	GetUidType() *string
}

type DelMessageRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BizName        *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	CallerProtocol *string `json:"CallerProtocol,omitempty" xml:"CallerProtocol,omitempty"`
	ClientSource   *string `json:"ClientSource,omitempty" xml:"ClientSource,omitempty"`
	Cookies        *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
	MsgId          *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	SrcUrl         *string `json:"SrcUrl,omitempty" xml:"SrcUrl,omitempty"`
	TenantCode     *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
	UidType        *string `json:"UidType,omitempty" xml:"UidType,omitempty"`
}

func (s DelMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s DelMessageRequest) GoString() string {
	return s.String()
}

func (s *DelMessageRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DelMessageRequest) GetAppName() *string {
	return s.AppName
}

func (s *DelMessageRequest) GetBizName() *string {
	return s.BizName
}

func (s *DelMessageRequest) GetCallerProtocol() *string {
	return s.CallerProtocol
}

func (s *DelMessageRequest) GetClientSource() *string {
	return s.ClientSource
}

func (s *DelMessageRequest) GetCookies() *string {
	return s.Cookies
}

func (s *DelMessageRequest) GetMsgId() *string {
	return s.MsgId
}

func (s *DelMessageRequest) GetSrcUrl() *string {
	return s.SrcUrl
}

func (s *DelMessageRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *DelMessageRequest) GetUidType() *string {
	return s.UidType
}

func (s *DelMessageRequest) SetAcceptLanguage(v string) *DelMessageRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DelMessageRequest) SetAppName(v string) *DelMessageRequest {
	s.AppName = &v
	return s
}

func (s *DelMessageRequest) SetBizName(v string) *DelMessageRequest {
	s.BizName = &v
	return s
}

func (s *DelMessageRequest) SetCallerProtocol(v string) *DelMessageRequest {
	s.CallerProtocol = &v
	return s
}

func (s *DelMessageRequest) SetClientSource(v string) *DelMessageRequest {
	s.ClientSource = &v
	return s
}

func (s *DelMessageRequest) SetCookies(v string) *DelMessageRequest {
	s.Cookies = &v
	return s
}

func (s *DelMessageRequest) SetMsgId(v string) *DelMessageRequest {
	s.MsgId = &v
	return s
}

func (s *DelMessageRequest) SetSrcUrl(v string) *DelMessageRequest {
	s.SrcUrl = &v
	return s
}

func (s *DelMessageRequest) SetTenantCode(v string) *DelMessageRequest {
	s.TenantCode = &v
	return s
}

func (s *DelMessageRequest) SetUidType(v string) *DelMessageRequest {
	s.UidType = &v
	return s
}

func (s *DelMessageRequest) Validate() error {
	return dara.Validate(s)
}
