// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVariableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunKp(v string) *GetVariableRequest
	GetAliyunKp() *string
	SetApiType(v string) *GetVariableRequest
	GetApiType() *string
	SetBid(v string) *GetVariableRequest
	GetBid() *string
	SetLang(v string) *GetVariableRequest
	GetLang() *string
	SetOriginalRequest(v string) *GetVariableRequest
	GetOriginalRequest() *string
	SetTemplateId(v int64) *GetVariableRequest
	GetTemplateId() *int64
	SetUserAccessKeyId(v string) *GetVariableRequest
	GetUserAccessKeyId() *string
	SetUserBid(v string) *GetVariableRequest
	GetUserBid() *string
	SetUserCallerParentId(v int64) *GetVariableRequest
	GetUserCallerParentId() *int64
	SetUserCallerSecurityTransport(v bool) *GetVariableRequest
	GetUserCallerSecurityTransport() *bool
	SetUserCallerType(v string) *GetVariableRequest
	GetUserCallerType() *string
	SetUserClientIp(v string) *GetVariableRequest
	GetUserClientIp() *string
	SetUserKp(v string) *GetVariableRequest
	GetUserKp() *string
	SetUserMfaPresent(v bool) *GetVariableRequest
	GetUserMfaPresent() *bool
	SetUserSecurityToken(v string) *GetVariableRequest
	GetUserSecurityToken() *string
}

type GetVariableRequest struct {
	// aliyunKp
	//
	// example:
	//
	// 1
	AliyunKp *string `json:"AliyunKp,omitempty" xml:"AliyunKp,omitempty"`
	// apiType
	//
	// example:
	//
	// openAPI
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// bid
	//
	// example:
	//
	// 26842
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// lang
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// originalRequest
	//
	// example:
	//
	// 1
	OriginalRequest *string `json:"OriginalRequest,omitempty" xml:"OriginalRequest,omitempty"`
	// example:
	//
	// 17
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// userAccessKeyId
	//
	// example:
	//
	// 1
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	// userBid
	//
	// example:
	//
	// 1
	UserBid *string `json:"UserBid,omitempty" xml:"UserBid,omitempty"`
	// userCallerParentId
	//
	// example:
	//
	// 1
	UserCallerParentId *int64 `json:"UserCallerParentId,omitempty" xml:"UserCallerParentId,omitempty"`
	// userCallerSecurityTransport
	//
	// example:
	//
	// true
	UserCallerSecurityTransport *bool `json:"UserCallerSecurityTransport,omitempty" xml:"UserCallerSecurityTransport,omitempty"`
	// userCallerType
	//
	// example:
	//
	// 1
	UserCallerType *string `json:"UserCallerType,omitempty" xml:"UserCallerType,omitempty"`
	// userClientIp
	//
	// example:
	//
	// 1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// userKp
	//
	// example:
	//
	// 1
	UserKp *string `json:"UserKp,omitempty" xml:"UserKp,omitempty"`
	// userMfaPresent
	//
	// example:
	//
	// true
	UserMfaPresent *bool `json:"UserMfaPresent,omitempty" xml:"UserMfaPresent,omitempty"`
	// userSecurityToken
	//
	// example:
	//
	// 1
	UserSecurityToken *string `json:"UserSecurityToken,omitempty" xml:"UserSecurityToken,omitempty"`
}

func (s GetVariableRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVariableRequest) GoString() string {
	return s.String()
}

func (s *GetVariableRequest) GetAliyunKp() *string {
	return s.AliyunKp
}

func (s *GetVariableRequest) GetApiType() *string {
	return s.ApiType
}

func (s *GetVariableRequest) GetBid() *string {
	return s.Bid
}

func (s *GetVariableRequest) GetLang() *string {
	return s.Lang
}

func (s *GetVariableRequest) GetOriginalRequest() *string {
	return s.OriginalRequest
}

func (s *GetVariableRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *GetVariableRequest) GetUserAccessKeyId() *string {
	return s.UserAccessKeyId
}

func (s *GetVariableRequest) GetUserBid() *string {
	return s.UserBid
}

func (s *GetVariableRequest) GetUserCallerParentId() *int64 {
	return s.UserCallerParentId
}

func (s *GetVariableRequest) GetUserCallerSecurityTransport() *bool {
	return s.UserCallerSecurityTransport
}

func (s *GetVariableRequest) GetUserCallerType() *string {
	return s.UserCallerType
}

func (s *GetVariableRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *GetVariableRequest) GetUserKp() *string {
	return s.UserKp
}

func (s *GetVariableRequest) GetUserMfaPresent() *bool {
	return s.UserMfaPresent
}

func (s *GetVariableRequest) GetUserSecurityToken() *string {
	return s.UserSecurityToken
}

func (s *GetVariableRequest) SetAliyunKp(v string) *GetVariableRequest {
	s.AliyunKp = &v
	return s
}

func (s *GetVariableRequest) SetApiType(v string) *GetVariableRequest {
	s.ApiType = &v
	return s
}

func (s *GetVariableRequest) SetBid(v string) *GetVariableRequest {
	s.Bid = &v
	return s
}

func (s *GetVariableRequest) SetLang(v string) *GetVariableRequest {
	s.Lang = &v
	return s
}

func (s *GetVariableRequest) SetOriginalRequest(v string) *GetVariableRequest {
	s.OriginalRequest = &v
	return s
}

func (s *GetVariableRequest) SetTemplateId(v int64) *GetVariableRequest {
	s.TemplateId = &v
	return s
}

func (s *GetVariableRequest) SetUserAccessKeyId(v string) *GetVariableRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *GetVariableRequest) SetUserBid(v string) *GetVariableRequest {
	s.UserBid = &v
	return s
}

func (s *GetVariableRequest) SetUserCallerParentId(v int64) *GetVariableRequest {
	s.UserCallerParentId = &v
	return s
}

func (s *GetVariableRequest) SetUserCallerSecurityTransport(v bool) *GetVariableRequest {
	s.UserCallerSecurityTransport = &v
	return s
}

func (s *GetVariableRequest) SetUserCallerType(v string) *GetVariableRequest {
	s.UserCallerType = &v
	return s
}

func (s *GetVariableRequest) SetUserClientIp(v string) *GetVariableRequest {
	s.UserClientIp = &v
	return s
}

func (s *GetVariableRequest) SetUserKp(v string) *GetVariableRequest {
	s.UserKp = &v
	return s
}

func (s *GetVariableRequest) SetUserMfaPresent(v bool) *GetVariableRequest {
	s.UserMfaPresent = &v
	return s
}

func (s *GetVariableRequest) SetUserSecurityToken(v string) *GetVariableRequest {
	s.UserSecurityToken = &v
	return s
}

func (s *GetVariableRequest) Validate() error {
	return dara.Validate(s)
}
