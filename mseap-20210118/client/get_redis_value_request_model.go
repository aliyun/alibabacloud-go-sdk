// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRedisValueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunKp(v string) *GetRedisValueRequest
	GetAliyunKp() *string
	SetApiType(v string) *GetRedisValueRequest
	GetApiType() *string
	SetBid(v string) *GetRedisValueRequest
	GetBid() *string
	SetKey(v string) *GetRedisValueRequest
	GetKey() *string
	SetLang(v string) *GetRedisValueRequest
	GetLang() *string
	SetOriginalRequest(v string) *GetRedisValueRequest
	GetOriginalRequest() *string
	SetTimeout(v int64) *GetRedisValueRequest
	GetTimeout() *int64
	SetUserAccessKeyId(v string) *GetRedisValueRequest
	GetUserAccessKeyId() *string
	SetUserBid(v string) *GetRedisValueRequest
	GetUserBid() *string
	SetUserCallerParentId(v int64) *GetRedisValueRequest
	GetUserCallerParentId() *int64
	SetUserCallerSecurityTransport(v bool) *GetRedisValueRequest
	GetUserCallerSecurityTransport() *bool
	SetUserCallerType(v string) *GetRedisValueRequest
	GetUserCallerType() *string
	SetUserClientIp(v string) *GetRedisValueRequest
	GetUserClientIp() *string
	SetUserKp(v string) *GetRedisValueRequest
	GetUserKp() *string
	SetUserMfaPresent(v bool) *GetRedisValueRequest
	GetUserMfaPresent() *bool
	SetUserSecurityToken(v string) *GetRedisValueRequest
	GetUserSecurityToken() *string
	SetValue(v string) *GetRedisValueRequest
	GetValue() *string
}

type GetRedisValueRequest struct {
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
	// MPC
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// bid
	//
	// example:
	//
	// 26842
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// key
	//
	// example:
	//
	// 106.14.34.208
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	// timeout
	//
	// example:
	//
	// 60
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
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
	// example:
	//
	// {   \\"cust_id\\":\\"1111111\\",   \\"cust_name\\":\\"aa\\" }
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetRedisValueRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRedisValueRequest) GoString() string {
	return s.String()
}

func (s *GetRedisValueRequest) GetAliyunKp() *string {
	return s.AliyunKp
}

func (s *GetRedisValueRequest) GetApiType() *string {
	return s.ApiType
}

func (s *GetRedisValueRequest) GetBid() *string {
	return s.Bid
}

func (s *GetRedisValueRequest) GetKey() *string {
	return s.Key
}

func (s *GetRedisValueRequest) GetLang() *string {
	return s.Lang
}

func (s *GetRedisValueRequest) GetOriginalRequest() *string {
	return s.OriginalRequest
}

func (s *GetRedisValueRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *GetRedisValueRequest) GetUserAccessKeyId() *string {
	return s.UserAccessKeyId
}

func (s *GetRedisValueRequest) GetUserBid() *string {
	return s.UserBid
}

func (s *GetRedisValueRequest) GetUserCallerParentId() *int64 {
	return s.UserCallerParentId
}

func (s *GetRedisValueRequest) GetUserCallerSecurityTransport() *bool {
	return s.UserCallerSecurityTransport
}

func (s *GetRedisValueRequest) GetUserCallerType() *string {
	return s.UserCallerType
}

func (s *GetRedisValueRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *GetRedisValueRequest) GetUserKp() *string {
	return s.UserKp
}

func (s *GetRedisValueRequest) GetUserMfaPresent() *bool {
	return s.UserMfaPresent
}

func (s *GetRedisValueRequest) GetUserSecurityToken() *string {
	return s.UserSecurityToken
}

func (s *GetRedisValueRequest) GetValue() *string {
	return s.Value
}

func (s *GetRedisValueRequest) SetAliyunKp(v string) *GetRedisValueRequest {
	s.AliyunKp = &v
	return s
}

func (s *GetRedisValueRequest) SetApiType(v string) *GetRedisValueRequest {
	s.ApiType = &v
	return s
}

func (s *GetRedisValueRequest) SetBid(v string) *GetRedisValueRequest {
	s.Bid = &v
	return s
}

func (s *GetRedisValueRequest) SetKey(v string) *GetRedisValueRequest {
	s.Key = &v
	return s
}

func (s *GetRedisValueRequest) SetLang(v string) *GetRedisValueRequest {
	s.Lang = &v
	return s
}

func (s *GetRedisValueRequest) SetOriginalRequest(v string) *GetRedisValueRequest {
	s.OriginalRequest = &v
	return s
}

func (s *GetRedisValueRequest) SetTimeout(v int64) *GetRedisValueRequest {
	s.Timeout = &v
	return s
}

func (s *GetRedisValueRequest) SetUserAccessKeyId(v string) *GetRedisValueRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *GetRedisValueRequest) SetUserBid(v string) *GetRedisValueRequest {
	s.UserBid = &v
	return s
}

func (s *GetRedisValueRequest) SetUserCallerParentId(v int64) *GetRedisValueRequest {
	s.UserCallerParentId = &v
	return s
}

func (s *GetRedisValueRequest) SetUserCallerSecurityTransport(v bool) *GetRedisValueRequest {
	s.UserCallerSecurityTransport = &v
	return s
}

func (s *GetRedisValueRequest) SetUserCallerType(v string) *GetRedisValueRequest {
	s.UserCallerType = &v
	return s
}

func (s *GetRedisValueRequest) SetUserClientIp(v string) *GetRedisValueRequest {
	s.UserClientIp = &v
	return s
}

func (s *GetRedisValueRequest) SetUserKp(v string) *GetRedisValueRequest {
	s.UserKp = &v
	return s
}

func (s *GetRedisValueRequest) SetUserMfaPresent(v bool) *GetRedisValueRequest {
	s.UserMfaPresent = &v
	return s
}

func (s *GetRedisValueRequest) SetUserSecurityToken(v string) *GetRedisValueRequest {
	s.UserSecurityToken = &v
	return s
}

func (s *GetRedisValueRequest) SetValue(v string) *GetRedisValueRequest {
	s.Value = &v
	return s
}

func (s *GetRedisValueRequest) Validate() error {
	return dara.Validate(s)
}
