// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRedisValueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunKp(v string) *SetRedisValueRequest
	GetAliyunKp() *string
	SetApiType(v string) *SetRedisValueRequest
	GetApiType() *string
	SetBid(v string) *SetRedisValueRequest
	GetBid() *string
	SetKey(v string) *SetRedisValueRequest
	GetKey() *string
	SetLang(v string) *SetRedisValueRequest
	GetLang() *string
	SetOriginalRequest(v string) *SetRedisValueRequest
	GetOriginalRequest() *string
	SetRequestId(v string) *SetRedisValueRequest
	GetRequestId() *string
	SetTimeout(v int64) *SetRedisValueRequest
	GetTimeout() *int64
	SetUserAccessKeyId(v string) *SetRedisValueRequest
	GetUserAccessKeyId() *string
	SetUserBid(v string) *SetRedisValueRequest
	GetUserBid() *string
	SetUserCallerParentId(v int64) *SetRedisValueRequest
	GetUserCallerParentId() *int64
	SetUserCallerType(v string) *SetRedisValueRequest
	GetUserCallerType() *string
	SetUserClientIp(v string) *SetRedisValueRequest
	GetUserClientIp() *string
	SetUserKp(v string) *SetRedisValueRequest
	GetUserKp() *string
	SetUserMfaPresent(v bool) *SetRedisValueRequest
	GetUserMfaPresent() *bool
	SetUserSecurityToken(v string) *SetRedisValueRequest
	GetUserSecurityToken() *string
	SetValue(v string) *SetRedisValueRequest
	GetValue() *string
}

type SetRedisValueRequest struct {
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
	// part_config_data
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// bid
	//
	// example:
	//
	// 26842
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// example:
	//
	// 1684967696495902
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
	// requestId
	//
	// example:
	//
	// F864A883-AD76-53D5-9A24-A6DAD5177697
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// timeout
	//
	// example:
	//
	// 5000
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
	// example:
	//
	// 259200000
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SetRedisValueRequest) String() string {
	return dara.Prettify(s)
}

func (s SetRedisValueRequest) GoString() string {
	return s.String()
}

func (s *SetRedisValueRequest) GetAliyunKp() *string {
	return s.AliyunKp
}

func (s *SetRedisValueRequest) GetApiType() *string {
	return s.ApiType
}

func (s *SetRedisValueRequest) GetBid() *string {
	return s.Bid
}

func (s *SetRedisValueRequest) GetKey() *string {
	return s.Key
}

func (s *SetRedisValueRequest) GetLang() *string {
	return s.Lang
}

func (s *SetRedisValueRequest) GetOriginalRequest() *string {
	return s.OriginalRequest
}

func (s *SetRedisValueRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *SetRedisValueRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *SetRedisValueRequest) GetUserAccessKeyId() *string {
	return s.UserAccessKeyId
}

func (s *SetRedisValueRequest) GetUserBid() *string {
	return s.UserBid
}

func (s *SetRedisValueRequest) GetUserCallerParentId() *int64 {
	return s.UserCallerParentId
}

func (s *SetRedisValueRequest) GetUserCallerType() *string {
	return s.UserCallerType
}

func (s *SetRedisValueRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SetRedisValueRequest) GetUserKp() *string {
	return s.UserKp
}

func (s *SetRedisValueRequest) GetUserMfaPresent() *bool {
	return s.UserMfaPresent
}

func (s *SetRedisValueRequest) GetUserSecurityToken() *string {
	return s.UserSecurityToken
}

func (s *SetRedisValueRequest) GetValue() *string {
	return s.Value
}

func (s *SetRedisValueRequest) SetAliyunKp(v string) *SetRedisValueRequest {
	s.AliyunKp = &v
	return s
}

func (s *SetRedisValueRequest) SetApiType(v string) *SetRedisValueRequest {
	s.ApiType = &v
	return s
}

func (s *SetRedisValueRequest) SetBid(v string) *SetRedisValueRequest {
	s.Bid = &v
	return s
}

func (s *SetRedisValueRequest) SetKey(v string) *SetRedisValueRequest {
	s.Key = &v
	return s
}

func (s *SetRedisValueRequest) SetLang(v string) *SetRedisValueRequest {
	s.Lang = &v
	return s
}

func (s *SetRedisValueRequest) SetOriginalRequest(v string) *SetRedisValueRequest {
	s.OriginalRequest = &v
	return s
}

func (s *SetRedisValueRequest) SetRequestId(v string) *SetRedisValueRequest {
	s.RequestId = &v
	return s
}

func (s *SetRedisValueRequest) SetTimeout(v int64) *SetRedisValueRequest {
	s.Timeout = &v
	return s
}

func (s *SetRedisValueRequest) SetUserAccessKeyId(v string) *SetRedisValueRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *SetRedisValueRequest) SetUserBid(v string) *SetRedisValueRequest {
	s.UserBid = &v
	return s
}

func (s *SetRedisValueRequest) SetUserCallerParentId(v int64) *SetRedisValueRequest {
	s.UserCallerParentId = &v
	return s
}

func (s *SetRedisValueRequest) SetUserCallerType(v string) *SetRedisValueRequest {
	s.UserCallerType = &v
	return s
}

func (s *SetRedisValueRequest) SetUserClientIp(v string) *SetRedisValueRequest {
	s.UserClientIp = &v
	return s
}

func (s *SetRedisValueRequest) SetUserKp(v string) *SetRedisValueRequest {
	s.UserKp = &v
	return s
}

func (s *SetRedisValueRequest) SetUserMfaPresent(v bool) *SetRedisValueRequest {
	s.UserMfaPresent = &v
	return s
}

func (s *SetRedisValueRequest) SetUserSecurityToken(v string) *SetRedisValueRequest {
	s.UserSecurityToken = &v
	return s
}

func (s *SetRedisValueRequest) SetValue(v string) *SetRedisValueRequest {
	s.Value = &v
	return s
}

func (s *SetRedisValueRequest) Validate() error {
	return dara.Validate(s)
}
