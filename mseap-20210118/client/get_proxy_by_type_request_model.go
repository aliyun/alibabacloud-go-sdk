// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProxyByTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunKp(v string) *GetProxyByTypeRequest
	GetAliyunKp() *string
	SetApiType(v string) *GetProxyByTypeRequest
	GetApiType() *string
	SetBid(v string) *GetProxyByTypeRequest
	GetBid() *string
	SetLang(v string) *GetProxyByTypeRequest
	GetLang() *string
	SetOriginalRequest(v string) *GetProxyByTypeRequest
	GetOriginalRequest() *string
	SetType(v int32) *GetProxyByTypeRequest
	GetType() *int32
	SetUserAccessKeyId(v string) *GetProxyByTypeRequest
	GetUserAccessKeyId() *string
	SetUserBid(v string) *GetProxyByTypeRequest
	GetUserBid() *string
	SetUserCallerParentId(v int64) *GetProxyByTypeRequest
	GetUserCallerParentId() *int64
	SetUserCallerSecurityTransport(v bool) *GetProxyByTypeRequest
	GetUserCallerSecurityTransport() *bool
	SetUserCallerType(v string) *GetProxyByTypeRequest
	GetUserCallerType() *string
	SetUserClientIp(v string) *GetProxyByTypeRequest
	GetUserClientIp() *string
	SetUserKp(v string) *GetProxyByTypeRequest
	GetUserKp() *string
	SetUserMfaPresent(v bool) *GetProxyByTypeRequest
	GetUserMfaPresent() *bool
	SetUserSecurityToken(v string) *GetProxyByTypeRequest
	GetUserSecurityToken() *string
}

type GetProxyByTypeRequest struct {
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
	// type
	//
	// example:
	//
	// 2
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
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
	// 1
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
	// 1
	UserMfaPresent *bool `json:"UserMfaPresent,omitempty" xml:"UserMfaPresent,omitempty"`
	// userSecurityToken
	//
	// example:
	//
	// 1
	UserSecurityToken *string `json:"UserSecurityToken,omitempty" xml:"UserSecurityToken,omitempty"`
}

func (s GetProxyByTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProxyByTypeRequest) GoString() string {
	return s.String()
}

func (s *GetProxyByTypeRequest) GetAliyunKp() *string {
	return s.AliyunKp
}

func (s *GetProxyByTypeRequest) GetApiType() *string {
	return s.ApiType
}

func (s *GetProxyByTypeRequest) GetBid() *string {
	return s.Bid
}

func (s *GetProxyByTypeRequest) GetLang() *string {
	return s.Lang
}

func (s *GetProxyByTypeRequest) GetOriginalRequest() *string {
	return s.OriginalRequest
}

func (s *GetProxyByTypeRequest) GetType() *int32 {
	return s.Type
}

func (s *GetProxyByTypeRequest) GetUserAccessKeyId() *string {
	return s.UserAccessKeyId
}

func (s *GetProxyByTypeRequest) GetUserBid() *string {
	return s.UserBid
}

func (s *GetProxyByTypeRequest) GetUserCallerParentId() *int64 {
	return s.UserCallerParentId
}

func (s *GetProxyByTypeRequest) GetUserCallerSecurityTransport() *bool {
	return s.UserCallerSecurityTransport
}

func (s *GetProxyByTypeRequest) GetUserCallerType() *string {
	return s.UserCallerType
}

func (s *GetProxyByTypeRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *GetProxyByTypeRequest) GetUserKp() *string {
	return s.UserKp
}

func (s *GetProxyByTypeRequest) GetUserMfaPresent() *bool {
	return s.UserMfaPresent
}

func (s *GetProxyByTypeRequest) GetUserSecurityToken() *string {
	return s.UserSecurityToken
}

func (s *GetProxyByTypeRequest) SetAliyunKp(v string) *GetProxyByTypeRequest {
	s.AliyunKp = &v
	return s
}

func (s *GetProxyByTypeRequest) SetApiType(v string) *GetProxyByTypeRequest {
	s.ApiType = &v
	return s
}

func (s *GetProxyByTypeRequest) SetBid(v string) *GetProxyByTypeRequest {
	s.Bid = &v
	return s
}

func (s *GetProxyByTypeRequest) SetLang(v string) *GetProxyByTypeRequest {
	s.Lang = &v
	return s
}

func (s *GetProxyByTypeRequest) SetOriginalRequest(v string) *GetProxyByTypeRequest {
	s.OriginalRequest = &v
	return s
}

func (s *GetProxyByTypeRequest) SetType(v int32) *GetProxyByTypeRequest {
	s.Type = &v
	return s
}

func (s *GetProxyByTypeRequest) SetUserAccessKeyId(v string) *GetProxyByTypeRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *GetProxyByTypeRequest) SetUserBid(v string) *GetProxyByTypeRequest {
	s.UserBid = &v
	return s
}

func (s *GetProxyByTypeRequest) SetUserCallerParentId(v int64) *GetProxyByTypeRequest {
	s.UserCallerParentId = &v
	return s
}

func (s *GetProxyByTypeRequest) SetUserCallerSecurityTransport(v bool) *GetProxyByTypeRequest {
	s.UserCallerSecurityTransport = &v
	return s
}

func (s *GetProxyByTypeRequest) SetUserCallerType(v string) *GetProxyByTypeRequest {
	s.UserCallerType = &v
	return s
}

func (s *GetProxyByTypeRequest) SetUserClientIp(v string) *GetProxyByTypeRequest {
	s.UserClientIp = &v
	return s
}

func (s *GetProxyByTypeRequest) SetUserKp(v string) *GetProxyByTypeRequest {
	s.UserKp = &v
	return s
}

func (s *GetProxyByTypeRequest) SetUserMfaPresent(v bool) *GetProxyByTypeRequest {
	s.UserMfaPresent = &v
	return s
}

func (s *GetProxyByTypeRequest) SetUserSecurityToken(v string) *GetProxyByTypeRequest {
	s.UserSecurityToken = &v
	return s
}

func (s *GetProxyByTypeRequest) Validate() error {
	return dara.Validate(s)
}
