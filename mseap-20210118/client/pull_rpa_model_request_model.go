// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPullRpaModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunKp(v string) *PullRpaModelRequest
	GetAliyunKp() *string
	SetApiType(v string) *PullRpaModelRequest
	GetApiType() *string
	SetBid(v string) *PullRpaModelRequest
	GetBid() *string
	SetLang(v string) *PullRpaModelRequest
	GetLang() *string
	SetOriginalRequest(v string) *PullRpaModelRequest
	GetOriginalRequest() *string
	SetTemplateId(v int64) *PullRpaModelRequest
	GetTemplateId() *int64
	SetUserAccessKeyId(v string) *PullRpaModelRequest
	GetUserAccessKeyId() *string
	SetUserBid(v string) *PullRpaModelRequest
	GetUserBid() *string
	SetUserCallerParentId(v int64) *PullRpaModelRequest
	GetUserCallerParentId() *int64
	SetUserCallerType(v string) *PullRpaModelRequest
	GetUserCallerType() *string
	SetUserClientIp(v string) *PullRpaModelRequest
	GetUserClientIp() *string
	SetUserKp(v string) *PullRpaModelRequest
	GetUserKp() *string
	SetUserMfaPresent(v bool) *PullRpaModelRequest
	GetUserMfaPresent() *bool
	SetUserSecurityToken(v string) *PullRpaModelRequest
	GetUserSecurityToken() *string
}

type PullRpaModelRequest struct {
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
	// templateId
	//
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

func (s PullRpaModelRequest) String() string {
	return dara.Prettify(s)
}

func (s PullRpaModelRequest) GoString() string {
	return s.String()
}

func (s *PullRpaModelRequest) GetAliyunKp() *string {
	return s.AliyunKp
}

func (s *PullRpaModelRequest) GetApiType() *string {
	return s.ApiType
}

func (s *PullRpaModelRequest) GetBid() *string {
	return s.Bid
}

func (s *PullRpaModelRequest) GetLang() *string {
	return s.Lang
}

func (s *PullRpaModelRequest) GetOriginalRequest() *string {
	return s.OriginalRequest
}

func (s *PullRpaModelRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *PullRpaModelRequest) GetUserAccessKeyId() *string {
	return s.UserAccessKeyId
}

func (s *PullRpaModelRequest) GetUserBid() *string {
	return s.UserBid
}

func (s *PullRpaModelRequest) GetUserCallerParentId() *int64 {
	return s.UserCallerParentId
}

func (s *PullRpaModelRequest) GetUserCallerType() *string {
	return s.UserCallerType
}

func (s *PullRpaModelRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *PullRpaModelRequest) GetUserKp() *string {
	return s.UserKp
}

func (s *PullRpaModelRequest) GetUserMfaPresent() *bool {
	return s.UserMfaPresent
}

func (s *PullRpaModelRequest) GetUserSecurityToken() *string {
	return s.UserSecurityToken
}

func (s *PullRpaModelRequest) SetAliyunKp(v string) *PullRpaModelRequest {
	s.AliyunKp = &v
	return s
}

func (s *PullRpaModelRequest) SetApiType(v string) *PullRpaModelRequest {
	s.ApiType = &v
	return s
}

func (s *PullRpaModelRequest) SetBid(v string) *PullRpaModelRequest {
	s.Bid = &v
	return s
}

func (s *PullRpaModelRequest) SetLang(v string) *PullRpaModelRequest {
	s.Lang = &v
	return s
}

func (s *PullRpaModelRequest) SetOriginalRequest(v string) *PullRpaModelRequest {
	s.OriginalRequest = &v
	return s
}

func (s *PullRpaModelRequest) SetTemplateId(v int64) *PullRpaModelRequest {
	s.TemplateId = &v
	return s
}

func (s *PullRpaModelRequest) SetUserAccessKeyId(v string) *PullRpaModelRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *PullRpaModelRequest) SetUserBid(v string) *PullRpaModelRequest {
	s.UserBid = &v
	return s
}

func (s *PullRpaModelRequest) SetUserCallerParentId(v int64) *PullRpaModelRequest {
	s.UserCallerParentId = &v
	return s
}

func (s *PullRpaModelRequest) SetUserCallerType(v string) *PullRpaModelRequest {
	s.UserCallerType = &v
	return s
}

func (s *PullRpaModelRequest) SetUserClientIp(v string) *PullRpaModelRequest {
	s.UserClientIp = &v
	return s
}

func (s *PullRpaModelRequest) SetUserKp(v string) *PullRpaModelRequest {
	s.UserKp = &v
	return s
}

func (s *PullRpaModelRequest) SetUserMfaPresent(v bool) *PullRpaModelRequest {
	s.UserMfaPresent = &v
	return s
}

func (s *PullRpaModelRequest) SetUserSecurityToken(v string) *PullRpaModelRequest {
	s.UserSecurityToken = &v
	return s
}

func (s *PullRpaModelRequest) Validate() error {
	return dara.Validate(s)
}
