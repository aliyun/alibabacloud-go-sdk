// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeByFlowIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunKp(v string) *GetNodeByFlowIdRequest
	GetAliyunKp() *string
	SetApiType(v string) *GetNodeByFlowIdRequest
	GetApiType() *string
	SetBid(v string) *GetNodeByFlowIdRequest
	GetBid() *string
	SetFlowId(v int64) *GetNodeByFlowIdRequest
	GetFlowId() *int64
	SetLang(v string) *GetNodeByFlowIdRequest
	GetLang() *string
	SetOriginalRequest(v string) *GetNodeByFlowIdRequest
	GetOriginalRequest() *string
	SetUserAccessKeyId(v string) *GetNodeByFlowIdRequest
	GetUserAccessKeyId() *string
	SetUserBid(v string) *GetNodeByFlowIdRequest
	GetUserBid() *string
	SetUserCallerParentId(v int64) *GetNodeByFlowIdRequest
	GetUserCallerParentId() *int64
	SetUserCallerSecurityTransport(v bool) *GetNodeByFlowIdRequest
	GetUserCallerSecurityTransport() *bool
	SetUserCallerType(v string) *GetNodeByFlowIdRequest
	GetUserCallerType() *string
	SetUserClientIp(v string) *GetNodeByFlowIdRequest
	GetUserClientIp() *string
	SetUserKp(v string) *GetNodeByFlowIdRequest
	GetUserKp() *string
	SetUserMfaPresent(v bool) *GetNodeByFlowIdRequest
	GetUserMfaPresent() *bool
	SetUserSecurityToken(v string) *GetNodeByFlowIdRequest
	GetUserSecurityToken() *string
}

type GetNodeByFlowIdRequest struct {
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
	// example:
	//
	// 180
	FlowId *int64 `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// lang
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// originalRequest
	//
	// example:
	//
	// 1
	OriginalRequest *string `json:"OriginalRequest,omitempty" xml:"OriginalRequest,omitempty"`
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

func (s GetNodeByFlowIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNodeByFlowIdRequest) GoString() string {
	return s.String()
}

func (s *GetNodeByFlowIdRequest) GetAliyunKp() *string {
	return s.AliyunKp
}

func (s *GetNodeByFlowIdRequest) GetApiType() *string {
	return s.ApiType
}

func (s *GetNodeByFlowIdRequest) GetBid() *string {
	return s.Bid
}

func (s *GetNodeByFlowIdRequest) GetFlowId() *int64 {
	return s.FlowId
}

func (s *GetNodeByFlowIdRequest) GetLang() *string {
	return s.Lang
}

func (s *GetNodeByFlowIdRequest) GetOriginalRequest() *string {
	return s.OriginalRequest
}

func (s *GetNodeByFlowIdRequest) GetUserAccessKeyId() *string {
	return s.UserAccessKeyId
}

func (s *GetNodeByFlowIdRequest) GetUserBid() *string {
	return s.UserBid
}

func (s *GetNodeByFlowIdRequest) GetUserCallerParentId() *int64 {
	return s.UserCallerParentId
}

func (s *GetNodeByFlowIdRequest) GetUserCallerSecurityTransport() *bool {
	return s.UserCallerSecurityTransport
}

func (s *GetNodeByFlowIdRequest) GetUserCallerType() *string {
	return s.UserCallerType
}

func (s *GetNodeByFlowIdRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *GetNodeByFlowIdRequest) GetUserKp() *string {
	return s.UserKp
}

func (s *GetNodeByFlowIdRequest) GetUserMfaPresent() *bool {
	return s.UserMfaPresent
}

func (s *GetNodeByFlowIdRequest) GetUserSecurityToken() *string {
	return s.UserSecurityToken
}

func (s *GetNodeByFlowIdRequest) SetAliyunKp(v string) *GetNodeByFlowIdRequest {
	s.AliyunKp = &v
	return s
}

func (s *GetNodeByFlowIdRequest) SetApiType(v string) *GetNodeByFlowIdRequest {
	s.ApiType = &v
	return s
}

func (s *GetNodeByFlowIdRequest) SetBid(v string) *GetNodeByFlowIdRequest {
	s.Bid = &v
	return s
}

func (s *GetNodeByFlowIdRequest) SetFlowId(v int64) *GetNodeByFlowIdRequest {
	s.FlowId = &v
	return s
}

func (s *GetNodeByFlowIdRequest) SetLang(v string) *GetNodeByFlowIdRequest {
	s.Lang = &v
	return s
}

func (s *GetNodeByFlowIdRequest) SetOriginalRequest(v string) *GetNodeByFlowIdRequest {
	s.OriginalRequest = &v
	return s
}

func (s *GetNodeByFlowIdRequest) SetUserAccessKeyId(v string) *GetNodeByFlowIdRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *GetNodeByFlowIdRequest) SetUserBid(v string) *GetNodeByFlowIdRequest {
	s.UserBid = &v
	return s
}

func (s *GetNodeByFlowIdRequest) SetUserCallerParentId(v int64) *GetNodeByFlowIdRequest {
	s.UserCallerParentId = &v
	return s
}

func (s *GetNodeByFlowIdRequest) SetUserCallerSecurityTransport(v bool) *GetNodeByFlowIdRequest {
	s.UserCallerSecurityTransport = &v
	return s
}

func (s *GetNodeByFlowIdRequest) SetUserCallerType(v string) *GetNodeByFlowIdRequest {
	s.UserCallerType = &v
	return s
}

func (s *GetNodeByFlowIdRequest) SetUserClientIp(v string) *GetNodeByFlowIdRequest {
	s.UserClientIp = &v
	return s
}

func (s *GetNodeByFlowIdRequest) SetUserKp(v string) *GetNodeByFlowIdRequest {
	s.UserKp = &v
	return s
}

func (s *GetNodeByFlowIdRequest) SetUserMfaPresent(v bool) *GetNodeByFlowIdRequest {
	s.UserMfaPresent = &v
	return s
}

func (s *GetNodeByFlowIdRequest) SetUserSecurityToken(v string) *GetNodeByFlowIdRequest {
	s.UserSecurityToken = &v
	return s
}

func (s *GetNodeByFlowIdRequest) Validate() error {
	return dara.Validate(s)
}
