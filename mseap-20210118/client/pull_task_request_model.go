// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPullTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunKp(v string) *PullTaskRequest
	GetAliyunKp() *string
	SetApiType(v string) *PullTaskRequest
	GetApiType() *string
	SetBid(v string) *PullTaskRequest
	GetBid() *string
	SetBizCode(v string) *PullTaskRequest
	GetBizCode() *string
	SetLang(v string) *PullTaskRequest
	GetLang() *string
	SetOrderId(v string) *PullTaskRequest
	GetOrderId() *string
	SetOriginalRequest(v string) *PullTaskRequest
	GetOriginalRequest() *string
	SetPrincipalKey(v string) *PullTaskRequest
	GetPrincipalKey() *string
	SetTaskType(v string) *PullTaskRequest
	GetTaskType() *string
	SetUserAccessKeyId(v string) *PullTaskRequest
	GetUserAccessKeyId() *string
	SetUserBid(v string) *PullTaskRequest
	GetUserBid() *string
	SetUserCallerParentId(v int64) *PullTaskRequest
	GetUserCallerParentId() *int64
	SetUserCallerSecurityTransport(v bool) *PullTaskRequest
	GetUserCallerSecurityTransport() *bool
	SetUserCallerType(v string) *PullTaskRequest
	GetUserCallerType() *string
	SetUserClientIp(v string) *PullTaskRequest
	GetUserClientIp() *string
	SetUserKp(v string) *PullTaskRequest
	GetUserKp() *string
	SetUserMfaPresent(v bool) *PullTaskRequest
	GetUserMfaPresent() *bool
	SetUserSecurityToken(v string) *PullTaskRequest
	GetUserSecurityToken() *string
}

type PullTaskRequest struct {
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
	Bid     *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// lang
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 1672369049358
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// originalRequest
	//
	// example:
	//
	// 1
	OriginalRequest *string `json:"OriginalRequest,omitempty" xml:"OriginalRequest,omitempty"`
	PrincipalKey    *string `json:"PrincipalKey,omitempty" xml:"PrincipalKey,omitempty"`
	// taskType
	//
	// example:
	//
	// PATENT_CHECK
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
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

func (s PullTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s PullTaskRequest) GoString() string {
	return s.String()
}

func (s *PullTaskRequest) GetAliyunKp() *string {
	return s.AliyunKp
}

func (s *PullTaskRequest) GetApiType() *string {
	return s.ApiType
}

func (s *PullTaskRequest) GetBid() *string {
	return s.Bid
}

func (s *PullTaskRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *PullTaskRequest) GetLang() *string {
	return s.Lang
}

func (s *PullTaskRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *PullTaskRequest) GetOriginalRequest() *string {
	return s.OriginalRequest
}

func (s *PullTaskRequest) GetPrincipalKey() *string {
	return s.PrincipalKey
}

func (s *PullTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *PullTaskRequest) GetUserAccessKeyId() *string {
	return s.UserAccessKeyId
}

func (s *PullTaskRequest) GetUserBid() *string {
	return s.UserBid
}

func (s *PullTaskRequest) GetUserCallerParentId() *int64 {
	return s.UserCallerParentId
}

func (s *PullTaskRequest) GetUserCallerSecurityTransport() *bool {
	return s.UserCallerSecurityTransport
}

func (s *PullTaskRequest) GetUserCallerType() *string {
	return s.UserCallerType
}

func (s *PullTaskRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *PullTaskRequest) GetUserKp() *string {
	return s.UserKp
}

func (s *PullTaskRequest) GetUserMfaPresent() *bool {
	return s.UserMfaPresent
}

func (s *PullTaskRequest) GetUserSecurityToken() *string {
	return s.UserSecurityToken
}

func (s *PullTaskRequest) SetAliyunKp(v string) *PullTaskRequest {
	s.AliyunKp = &v
	return s
}

func (s *PullTaskRequest) SetApiType(v string) *PullTaskRequest {
	s.ApiType = &v
	return s
}

func (s *PullTaskRequest) SetBid(v string) *PullTaskRequest {
	s.Bid = &v
	return s
}

func (s *PullTaskRequest) SetBizCode(v string) *PullTaskRequest {
	s.BizCode = &v
	return s
}

func (s *PullTaskRequest) SetLang(v string) *PullTaskRequest {
	s.Lang = &v
	return s
}

func (s *PullTaskRequest) SetOrderId(v string) *PullTaskRequest {
	s.OrderId = &v
	return s
}

func (s *PullTaskRequest) SetOriginalRequest(v string) *PullTaskRequest {
	s.OriginalRequest = &v
	return s
}

func (s *PullTaskRequest) SetPrincipalKey(v string) *PullTaskRequest {
	s.PrincipalKey = &v
	return s
}

func (s *PullTaskRequest) SetTaskType(v string) *PullTaskRequest {
	s.TaskType = &v
	return s
}

func (s *PullTaskRequest) SetUserAccessKeyId(v string) *PullTaskRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *PullTaskRequest) SetUserBid(v string) *PullTaskRequest {
	s.UserBid = &v
	return s
}

func (s *PullTaskRequest) SetUserCallerParentId(v int64) *PullTaskRequest {
	s.UserCallerParentId = &v
	return s
}

func (s *PullTaskRequest) SetUserCallerSecurityTransport(v bool) *PullTaskRequest {
	s.UserCallerSecurityTransport = &v
	return s
}

func (s *PullTaskRequest) SetUserCallerType(v string) *PullTaskRequest {
	s.UserCallerType = &v
	return s
}

func (s *PullTaskRequest) SetUserClientIp(v string) *PullTaskRequest {
	s.UserClientIp = &v
	return s
}

func (s *PullTaskRequest) SetUserKp(v string) *PullTaskRequest {
	s.UserKp = &v
	return s
}

func (s *PullTaskRequest) SetUserMfaPresent(v bool) *PullTaskRequest {
	s.UserMfaPresent = &v
	return s
}

func (s *PullTaskRequest) SetUserSecurityToken(v string) *PullTaskRequest {
	s.UserSecurityToken = &v
	return s
}

func (s *PullTaskRequest) Validate() error {
	return dara.Validate(s)
}
