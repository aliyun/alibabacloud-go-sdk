// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadMessageContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ReadMessageContentRequest
	GetAcceptLanguage() *string
	SetAppName(v string) *ReadMessageContentRequest
	GetAppName() *string
	SetBizName(v string) *ReadMessageContentRequest
	GetBizName() *string
	SetCallerProtocol(v string) *ReadMessageContentRequest
	GetCallerProtocol() *string
	SetClassId(v int64) *ReadMessageContentRequest
	GetClassId() *int64
	SetClientSource(v string) *ReadMessageContentRequest
	GetClientSource() *string
	SetCookies(v string) *ReadMessageContentRequest
	GetCookies() *string
	SetGroupCode(v string) *ReadMessageContentRequest
	GetGroupCode() *string
	SetHistory(v bool) *ReadMessageContentRequest
	GetHistory() *bool
	SetMsgId(v string) *ReadMessageContentRequest
	GetMsgId() *string
	SetSrcUrl(v string) *ReadMessageContentRequest
	GetSrcUrl() *string
	SetStatus(v int32) *ReadMessageContentRequest
	GetStatus() *int32
	SetTenantCode(v string) *ReadMessageContentRequest
	GetTenantCode() *string
	SetUidType(v string) *ReadMessageContentRequest
	GetUidType() *string
}

type ReadMessageContentRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BizName        *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	CallerProtocol *string `json:"CallerProtocol,omitempty" xml:"CallerProtocol,omitempty"`
	ClassId        *int64  `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	ClientSource   *string `json:"ClientSource,omitempty" xml:"ClientSource,omitempty"`
	Cookies        *string `json:"Cookies,omitempty" xml:"Cookies,omitempty"`
	GroupCode      *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
	History        *bool   `json:"History,omitempty" xml:"History,omitempty"`
	MsgId          *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	SrcUrl         *string `json:"SrcUrl,omitempty" xml:"SrcUrl,omitempty"`
	Status         *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantCode     *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
	UidType        *string `json:"UidType,omitempty" xml:"UidType,omitempty"`
}

func (s ReadMessageContentRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadMessageContentRequest) GoString() string {
	return s.String()
}

func (s *ReadMessageContentRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ReadMessageContentRequest) GetAppName() *string {
	return s.AppName
}

func (s *ReadMessageContentRequest) GetBizName() *string {
	return s.BizName
}

func (s *ReadMessageContentRequest) GetCallerProtocol() *string {
	return s.CallerProtocol
}

func (s *ReadMessageContentRequest) GetClassId() *int64 {
	return s.ClassId
}

func (s *ReadMessageContentRequest) GetClientSource() *string {
	return s.ClientSource
}

func (s *ReadMessageContentRequest) GetCookies() *string {
	return s.Cookies
}

func (s *ReadMessageContentRequest) GetGroupCode() *string {
	return s.GroupCode
}

func (s *ReadMessageContentRequest) GetHistory() *bool {
	return s.History
}

func (s *ReadMessageContentRequest) GetMsgId() *string {
	return s.MsgId
}

func (s *ReadMessageContentRequest) GetSrcUrl() *string {
	return s.SrcUrl
}

func (s *ReadMessageContentRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ReadMessageContentRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *ReadMessageContentRequest) GetUidType() *string {
	return s.UidType
}

func (s *ReadMessageContentRequest) SetAcceptLanguage(v string) *ReadMessageContentRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReadMessageContentRequest) SetAppName(v string) *ReadMessageContentRequest {
	s.AppName = &v
	return s
}

func (s *ReadMessageContentRequest) SetBizName(v string) *ReadMessageContentRequest {
	s.BizName = &v
	return s
}

func (s *ReadMessageContentRequest) SetCallerProtocol(v string) *ReadMessageContentRequest {
	s.CallerProtocol = &v
	return s
}

func (s *ReadMessageContentRequest) SetClassId(v int64) *ReadMessageContentRequest {
	s.ClassId = &v
	return s
}

func (s *ReadMessageContentRequest) SetClientSource(v string) *ReadMessageContentRequest {
	s.ClientSource = &v
	return s
}

func (s *ReadMessageContentRequest) SetCookies(v string) *ReadMessageContentRequest {
	s.Cookies = &v
	return s
}

func (s *ReadMessageContentRequest) SetGroupCode(v string) *ReadMessageContentRequest {
	s.GroupCode = &v
	return s
}

func (s *ReadMessageContentRequest) SetHistory(v bool) *ReadMessageContentRequest {
	s.History = &v
	return s
}

func (s *ReadMessageContentRequest) SetMsgId(v string) *ReadMessageContentRequest {
	s.MsgId = &v
	return s
}

func (s *ReadMessageContentRequest) SetSrcUrl(v string) *ReadMessageContentRequest {
	s.SrcUrl = &v
	return s
}

func (s *ReadMessageContentRequest) SetStatus(v int32) *ReadMessageContentRequest {
	s.Status = &v
	return s
}

func (s *ReadMessageContentRequest) SetTenantCode(v string) *ReadMessageContentRequest {
	s.TenantCode = &v
	return s
}

func (s *ReadMessageContentRequest) SetUidType(v string) *ReadMessageContentRequest {
	s.UidType = &v
	return s
}

func (s *ReadMessageContentRequest) Validate() error {
	return dara.Validate(s)
}
