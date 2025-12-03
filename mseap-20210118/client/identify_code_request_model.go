// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIdentifyCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunKp(v string) *IdentifyCodeRequest
	GetAliyunKp() *string
	SetApiType(v string) *IdentifyCodeRequest
	GetApiType() *string
	SetBid(v string) *IdentifyCodeRequest
	GetBid() *string
	SetData(v string) *IdentifyCodeRequest
	GetData() *string
	SetLabel(v string) *IdentifyCodeRequest
	GetLabel() *string
	SetLang(v string) *IdentifyCodeRequest
	GetLang() *string
	SetOriginalRequest(v string) *IdentifyCodeRequest
	GetOriginalRequest() *string
	SetType(v string) *IdentifyCodeRequest
	GetType() *string
	SetUserAccessKeyId(v string) *IdentifyCodeRequest
	GetUserAccessKeyId() *string
	SetUserBid(v string) *IdentifyCodeRequest
	GetUserBid() *string
	SetUserCallerParentId(v int64) *IdentifyCodeRequest
	GetUserCallerParentId() *int64
	SetUserCallerSecurityTransport(v bool) *IdentifyCodeRequest
	GetUserCallerSecurityTransport() *bool
	SetUserCallerType(v string) *IdentifyCodeRequest
	GetUserCallerType() *string
	SetUserClientIp(v string) *IdentifyCodeRequest
	GetUserClientIp() *string
	SetUserKp(v string) *IdentifyCodeRequest
	GetUserKp() *string
	SetUserMfaPresent(v bool) *IdentifyCodeRequest
	GetUserMfaPresent() *bool
	SetUserSecurityToken(v string) *IdentifyCodeRequest
	GetUserSecurityToken() *string
}

type IdentifyCodeRequest struct {
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
	// example:
	//
	// {\\"engine\\": \\"MySQL\\", \\"instanceId\\": \\"rm-2zes07949gc0febg6\\", \\"userId\\": \\"1204765431532768\\", \\"previousExistConfig\\": False, \\"engineVersion\\": \\"8.0\\", \\"autoResourceOptimize\\": 0, \\"dasProOn\\": False}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 1551278
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
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
	// CBWP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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

func (s IdentifyCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s IdentifyCodeRequest) GoString() string {
	return s.String()
}

func (s *IdentifyCodeRequest) GetAliyunKp() *string {
	return s.AliyunKp
}

func (s *IdentifyCodeRequest) GetApiType() *string {
	return s.ApiType
}

func (s *IdentifyCodeRequest) GetBid() *string {
	return s.Bid
}

func (s *IdentifyCodeRequest) GetData() *string {
	return s.Data
}

func (s *IdentifyCodeRequest) GetLabel() *string {
	return s.Label
}

func (s *IdentifyCodeRequest) GetLang() *string {
	return s.Lang
}

func (s *IdentifyCodeRequest) GetOriginalRequest() *string {
	return s.OriginalRequest
}

func (s *IdentifyCodeRequest) GetType() *string {
	return s.Type
}

func (s *IdentifyCodeRequest) GetUserAccessKeyId() *string {
	return s.UserAccessKeyId
}

func (s *IdentifyCodeRequest) GetUserBid() *string {
	return s.UserBid
}

func (s *IdentifyCodeRequest) GetUserCallerParentId() *int64 {
	return s.UserCallerParentId
}

func (s *IdentifyCodeRequest) GetUserCallerSecurityTransport() *bool {
	return s.UserCallerSecurityTransport
}

func (s *IdentifyCodeRequest) GetUserCallerType() *string {
	return s.UserCallerType
}

func (s *IdentifyCodeRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *IdentifyCodeRequest) GetUserKp() *string {
	return s.UserKp
}

func (s *IdentifyCodeRequest) GetUserMfaPresent() *bool {
	return s.UserMfaPresent
}

func (s *IdentifyCodeRequest) GetUserSecurityToken() *string {
	return s.UserSecurityToken
}

func (s *IdentifyCodeRequest) SetAliyunKp(v string) *IdentifyCodeRequest {
	s.AliyunKp = &v
	return s
}

func (s *IdentifyCodeRequest) SetApiType(v string) *IdentifyCodeRequest {
	s.ApiType = &v
	return s
}

func (s *IdentifyCodeRequest) SetBid(v string) *IdentifyCodeRequest {
	s.Bid = &v
	return s
}

func (s *IdentifyCodeRequest) SetData(v string) *IdentifyCodeRequest {
	s.Data = &v
	return s
}

func (s *IdentifyCodeRequest) SetLabel(v string) *IdentifyCodeRequest {
	s.Label = &v
	return s
}

func (s *IdentifyCodeRequest) SetLang(v string) *IdentifyCodeRequest {
	s.Lang = &v
	return s
}

func (s *IdentifyCodeRequest) SetOriginalRequest(v string) *IdentifyCodeRequest {
	s.OriginalRequest = &v
	return s
}

func (s *IdentifyCodeRequest) SetType(v string) *IdentifyCodeRequest {
	s.Type = &v
	return s
}

func (s *IdentifyCodeRequest) SetUserAccessKeyId(v string) *IdentifyCodeRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *IdentifyCodeRequest) SetUserBid(v string) *IdentifyCodeRequest {
	s.UserBid = &v
	return s
}

func (s *IdentifyCodeRequest) SetUserCallerParentId(v int64) *IdentifyCodeRequest {
	s.UserCallerParentId = &v
	return s
}

func (s *IdentifyCodeRequest) SetUserCallerSecurityTransport(v bool) *IdentifyCodeRequest {
	s.UserCallerSecurityTransport = &v
	return s
}

func (s *IdentifyCodeRequest) SetUserCallerType(v string) *IdentifyCodeRequest {
	s.UserCallerType = &v
	return s
}

func (s *IdentifyCodeRequest) SetUserClientIp(v string) *IdentifyCodeRequest {
	s.UserClientIp = &v
	return s
}

func (s *IdentifyCodeRequest) SetUserKp(v string) *IdentifyCodeRequest {
	s.UserKp = &v
	return s
}

func (s *IdentifyCodeRequest) SetUserMfaPresent(v bool) *IdentifyCodeRequest {
	s.UserMfaPresent = &v
	return s
}

func (s *IdentifyCodeRequest) SetUserSecurityToken(v string) *IdentifyCodeRequest {
	s.UserSecurityToken = &v
	return s
}

func (s *IdentifyCodeRequest) Validate() error {
	return dara.Validate(s)
}
