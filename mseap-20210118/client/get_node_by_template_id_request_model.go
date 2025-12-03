// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeByTemplateIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunKp(v string) *GetNodeByTemplateIdRequest
	GetAliyunKp() *string
	SetApiType(v string) *GetNodeByTemplateIdRequest
	GetApiType() *string
	SetBid(v string) *GetNodeByTemplateIdRequest
	GetBid() *string
	SetLang(v string) *GetNodeByTemplateIdRequest
	GetLang() *string
	SetOriginalRequest(v string) *GetNodeByTemplateIdRequest
	GetOriginalRequest() *string
	SetTemplateId(v int64) *GetNodeByTemplateIdRequest
	GetTemplateId() *int64
	SetUserAccessKeyId(v string) *GetNodeByTemplateIdRequest
	GetUserAccessKeyId() *string
	SetUserBid(v string) *GetNodeByTemplateIdRequest
	GetUserBid() *string
	SetUserCallerParentId(v int64) *GetNodeByTemplateIdRequest
	GetUserCallerParentId() *int64
	SetUserCallerSecurityTransport(v bool) *GetNodeByTemplateIdRequest
	GetUserCallerSecurityTransport() *bool
	SetUserCallerType(v string) *GetNodeByTemplateIdRequest
	GetUserCallerType() *string
	SetUserClientIp(v string) *GetNodeByTemplateIdRequest
	GetUserClientIp() *string
	SetUserKp(v string) *GetNodeByTemplateIdRequest
	GetUserKp() *string
	SetUserMfaPresent(v bool) *GetNodeByTemplateIdRequest
	GetUserMfaPresent() *bool
	SetUserSecurityToken(v string) *GetNodeByTemplateIdRequest
	GetUserSecurityToken() *string
}

type GetNodeByTemplateIdRequest struct {
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
	// 13
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
	// true
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

func (s GetNodeByTemplateIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNodeByTemplateIdRequest) GoString() string {
	return s.String()
}

func (s *GetNodeByTemplateIdRequest) GetAliyunKp() *string {
	return s.AliyunKp
}

func (s *GetNodeByTemplateIdRequest) GetApiType() *string {
	return s.ApiType
}

func (s *GetNodeByTemplateIdRequest) GetBid() *string {
	return s.Bid
}

func (s *GetNodeByTemplateIdRequest) GetLang() *string {
	return s.Lang
}

func (s *GetNodeByTemplateIdRequest) GetOriginalRequest() *string {
	return s.OriginalRequest
}

func (s *GetNodeByTemplateIdRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *GetNodeByTemplateIdRequest) GetUserAccessKeyId() *string {
	return s.UserAccessKeyId
}

func (s *GetNodeByTemplateIdRequest) GetUserBid() *string {
	return s.UserBid
}

func (s *GetNodeByTemplateIdRequest) GetUserCallerParentId() *int64 {
	return s.UserCallerParentId
}

func (s *GetNodeByTemplateIdRequest) GetUserCallerSecurityTransport() *bool {
	return s.UserCallerSecurityTransport
}

func (s *GetNodeByTemplateIdRequest) GetUserCallerType() *string {
	return s.UserCallerType
}

func (s *GetNodeByTemplateIdRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *GetNodeByTemplateIdRequest) GetUserKp() *string {
	return s.UserKp
}

func (s *GetNodeByTemplateIdRequest) GetUserMfaPresent() *bool {
	return s.UserMfaPresent
}

func (s *GetNodeByTemplateIdRequest) GetUserSecurityToken() *string {
	return s.UserSecurityToken
}

func (s *GetNodeByTemplateIdRequest) SetAliyunKp(v string) *GetNodeByTemplateIdRequest {
	s.AliyunKp = &v
	return s
}

func (s *GetNodeByTemplateIdRequest) SetApiType(v string) *GetNodeByTemplateIdRequest {
	s.ApiType = &v
	return s
}

func (s *GetNodeByTemplateIdRequest) SetBid(v string) *GetNodeByTemplateIdRequest {
	s.Bid = &v
	return s
}

func (s *GetNodeByTemplateIdRequest) SetLang(v string) *GetNodeByTemplateIdRequest {
	s.Lang = &v
	return s
}

func (s *GetNodeByTemplateIdRequest) SetOriginalRequest(v string) *GetNodeByTemplateIdRequest {
	s.OriginalRequest = &v
	return s
}

func (s *GetNodeByTemplateIdRequest) SetTemplateId(v int64) *GetNodeByTemplateIdRequest {
	s.TemplateId = &v
	return s
}

func (s *GetNodeByTemplateIdRequest) SetUserAccessKeyId(v string) *GetNodeByTemplateIdRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *GetNodeByTemplateIdRequest) SetUserBid(v string) *GetNodeByTemplateIdRequest {
	s.UserBid = &v
	return s
}

func (s *GetNodeByTemplateIdRequest) SetUserCallerParentId(v int64) *GetNodeByTemplateIdRequest {
	s.UserCallerParentId = &v
	return s
}

func (s *GetNodeByTemplateIdRequest) SetUserCallerSecurityTransport(v bool) *GetNodeByTemplateIdRequest {
	s.UserCallerSecurityTransport = &v
	return s
}

func (s *GetNodeByTemplateIdRequest) SetUserCallerType(v string) *GetNodeByTemplateIdRequest {
	s.UserCallerType = &v
	return s
}

func (s *GetNodeByTemplateIdRequest) SetUserClientIp(v string) *GetNodeByTemplateIdRequest {
	s.UserClientIp = &v
	return s
}

func (s *GetNodeByTemplateIdRequest) SetUserKp(v string) *GetNodeByTemplateIdRequest {
	s.UserKp = &v
	return s
}

func (s *GetNodeByTemplateIdRequest) SetUserMfaPresent(v bool) *GetNodeByTemplateIdRequest {
	s.UserMfaPresent = &v
	return s
}

func (s *GetNodeByTemplateIdRequest) SetUserSecurityToken(v string) *GetNodeByTemplateIdRequest {
	s.UserSecurityToken = &v
	return s
}

func (s *GetNodeByTemplateIdRequest) Validate() error {
	return dara.Validate(s)
}
