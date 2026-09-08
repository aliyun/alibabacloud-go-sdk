// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ReadMessageRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *ReadMessageRequest
	GetAppName() *string
	SetBizName(v string) *ReadMessageRequest
	GetBizName() *string
	SetCallerProtocol(v string) *ReadMessageRequest
	GetCallerProtocol() *string
	SetClientSource(v string) *ReadMessageRequest
	GetClientSource() *string
	SetCookies(v string) *ReadMessageRequest
	GetCookies() *string
	SetMsgId(v string) *ReadMessageRequest
	GetMsgId() *string
	SetSrcUrl(v string) *ReadMessageRequest
	GetSrcUrl() *string
	SetTenantCode(v string) *ReadMessageRequest
	GetTenantCode() *string
	SetUidType(v string) *ReadMessageRequest
	GetUidType() *string
}

type ReadMessageRequest struct {
	// The language. Default value: zh-CN (Simplified Chinese).
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// System parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// /
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// System parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// /
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// System parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// /
	CallerProtocol *string `json:"CallerProtocol,omitempty" xml:"CallerProtocol,omitempty"`
	// System parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// /
	ClientSource *string `json:"ClientSource,omitempty" xml:"ClientSource,omitempty"`
	// System parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// /
	Cookies *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
	// The message ID.
	//
	// example:
	//
	// A9FEF04A004A383534AA6400092A904F
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The URL of the source page.
	//
	// example:
	//
	// /
	SrcUrl *string `json:"SrcUrl,omitempty" xml:"SrcUrl,omitempty"`
	// System parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// /
	TenantCode *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
	// System parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// /
	UidType *string `json:"UidType,omitempty" xml:"UidType,omitempty"`
}

func (s ReadMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadMessageRequest) GoString() string {
	return s.String()
}

func (s *ReadMessageRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ReadMessageRequest) GetAppName() *string {
	return s.AppName
}

func (s *ReadMessageRequest) GetBizName() *string {
	return s.BizName
}

func (s *ReadMessageRequest) GetCallerProtocol() *string {
	return s.CallerProtocol
}

func (s *ReadMessageRequest) GetClientSource() *string {
	return s.ClientSource
}

func (s *ReadMessageRequest) GetCookies() *string {
	return s.Cookies
}

func (s *ReadMessageRequest) GetMsgId() *string {
	return s.MsgId
}

func (s *ReadMessageRequest) GetSrcUrl() *string {
	return s.SrcUrl
}

func (s *ReadMessageRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *ReadMessageRequest) GetUidType() *string {
	return s.UidType
}

func (s *ReadMessageRequest) SetAcceptLanguage(v string) *ReadMessageRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReadMessageRequest) SetAppName(v string) *ReadMessageRequest {
	s.AppName = &v
	return s
}

func (s *ReadMessageRequest) SetBizName(v string) *ReadMessageRequest {
	s.BizName = &v
	return s
}

func (s *ReadMessageRequest) SetCallerProtocol(v string) *ReadMessageRequest {
	s.CallerProtocol = &v
	return s
}

func (s *ReadMessageRequest) SetClientSource(v string) *ReadMessageRequest {
	s.ClientSource = &v
	return s
}

func (s *ReadMessageRequest) SetCookies(v string) *ReadMessageRequest {
	s.Cookies = &v
	return s
}

func (s *ReadMessageRequest) SetMsgId(v string) *ReadMessageRequest {
	s.MsgId = &v
	return s
}

func (s *ReadMessageRequest) SetSrcUrl(v string) *ReadMessageRequest {
	s.SrcUrl = &v
	return s
}

func (s *ReadMessageRequest) SetTenantCode(v string) *ReadMessageRequest {
	s.TenantCode = &v
	return s
}

func (s *ReadMessageRequest) SetUidType(v string) *ReadMessageRequest {
	s.UidType = &v
	return s
}

func (s *ReadMessageRequest) Validate() error {
	return dara.Validate(s)
}
