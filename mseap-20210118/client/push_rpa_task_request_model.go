// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushRpaTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunKp(v string) *PushRpaTaskRequest
	GetAliyunKp() *string
	SetApiType(v string) *PushRpaTaskRequest
	GetApiType() *string
	SetBid(v string) *PushRpaTaskRequest
	GetBid() *string
	SetLang(v string) *PushRpaTaskRequest
	GetLang() *string
	SetModelId(v int64) *PushRpaTaskRequest
	GetModelId() *int64
	SetName(v string) *PushRpaTaskRequest
	GetName() *string
	SetOriginalRequest(v string) *PushRpaTaskRequest
	GetOriginalRequest() *string
	SetRequest(v string) *PushRpaTaskRequest
	GetRequest() *string
	SetStatus(v int32) *PushRpaTaskRequest
	GetStatus() *int32
	SetTaskId(v int64) *PushRpaTaskRequest
	GetTaskId() *int64
	SetTemplateId(v int64) *PushRpaTaskRequest
	GetTemplateId() *int64
	SetUserAccessKeyId(v string) *PushRpaTaskRequest
	GetUserAccessKeyId() *string
	SetUserBid(v string) *PushRpaTaskRequest
	GetUserBid() *string
	SetUserCallerParentId(v int64) *PushRpaTaskRequest
	GetUserCallerParentId() *int64
	SetUserCallerType(v string) *PushRpaTaskRequest
	GetUserCallerType() *string
	SetUserClientIp(v string) *PushRpaTaskRequest
	GetUserClientIp() *string
	SetUserKp(v string) *PushRpaTaskRequest
	GetUserKp() *string
	SetUserMfaPresent(v bool) *PushRpaTaskRequest
	GetUserMfaPresent() *bool
	SetUserSecurityToken(v string) *PushRpaTaskRequest
	GetUserSecurityToken() *string
}

type PushRpaTaskRequest struct {
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
	// public
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
	// modelId
	//
	// example:
	//
	// 1951
	ModelId *int64 `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// originalRequest
	//
	// example:
	//
	// 1
	OriginalRequest *string `json:"OriginalRequest,omitempty" xml:"OriginalRequest,omitempty"`
	// request
	//
	// example:
	//
	// {}
	Request *string `json:"Request,omitempty" xml:"Request,omitempty"`
	// status
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// taskId
	//
	// example:
	//
	// 833812
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// templateId
	//
	// example:
	//
	// 26
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

func (s PushRpaTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s PushRpaTaskRequest) GoString() string {
	return s.String()
}

func (s *PushRpaTaskRequest) GetAliyunKp() *string {
	return s.AliyunKp
}

func (s *PushRpaTaskRequest) GetApiType() *string {
	return s.ApiType
}

func (s *PushRpaTaskRequest) GetBid() *string {
	return s.Bid
}

func (s *PushRpaTaskRequest) GetLang() *string {
	return s.Lang
}

func (s *PushRpaTaskRequest) GetModelId() *int64 {
	return s.ModelId
}

func (s *PushRpaTaskRequest) GetName() *string {
	return s.Name
}

func (s *PushRpaTaskRequest) GetOriginalRequest() *string {
	return s.OriginalRequest
}

func (s *PushRpaTaskRequest) GetRequest() *string {
	return s.Request
}

func (s *PushRpaTaskRequest) GetStatus() *int32 {
	return s.Status
}

func (s *PushRpaTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *PushRpaTaskRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *PushRpaTaskRequest) GetUserAccessKeyId() *string {
	return s.UserAccessKeyId
}

func (s *PushRpaTaskRequest) GetUserBid() *string {
	return s.UserBid
}

func (s *PushRpaTaskRequest) GetUserCallerParentId() *int64 {
	return s.UserCallerParentId
}

func (s *PushRpaTaskRequest) GetUserCallerType() *string {
	return s.UserCallerType
}

func (s *PushRpaTaskRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *PushRpaTaskRequest) GetUserKp() *string {
	return s.UserKp
}

func (s *PushRpaTaskRequest) GetUserMfaPresent() *bool {
	return s.UserMfaPresent
}

func (s *PushRpaTaskRequest) GetUserSecurityToken() *string {
	return s.UserSecurityToken
}

func (s *PushRpaTaskRequest) SetAliyunKp(v string) *PushRpaTaskRequest {
	s.AliyunKp = &v
	return s
}

func (s *PushRpaTaskRequest) SetApiType(v string) *PushRpaTaskRequest {
	s.ApiType = &v
	return s
}

func (s *PushRpaTaskRequest) SetBid(v string) *PushRpaTaskRequest {
	s.Bid = &v
	return s
}

func (s *PushRpaTaskRequest) SetLang(v string) *PushRpaTaskRequest {
	s.Lang = &v
	return s
}

func (s *PushRpaTaskRequest) SetModelId(v int64) *PushRpaTaskRequest {
	s.ModelId = &v
	return s
}

func (s *PushRpaTaskRequest) SetName(v string) *PushRpaTaskRequest {
	s.Name = &v
	return s
}

func (s *PushRpaTaskRequest) SetOriginalRequest(v string) *PushRpaTaskRequest {
	s.OriginalRequest = &v
	return s
}

func (s *PushRpaTaskRequest) SetRequest(v string) *PushRpaTaskRequest {
	s.Request = &v
	return s
}

func (s *PushRpaTaskRequest) SetStatus(v int32) *PushRpaTaskRequest {
	s.Status = &v
	return s
}

func (s *PushRpaTaskRequest) SetTaskId(v int64) *PushRpaTaskRequest {
	s.TaskId = &v
	return s
}

func (s *PushRpaTaskRequest) SetTemplateId(v int64) *PushRpaTaskRequest {
	s.TemplateId = &v
	return s
}

func (s *PushRpaTaskRequest) SetUserAccessKeyId(v string) *PushRpaTaskRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *PushRpaTaskRequest) SetUserBid(v string) *PushRpaTaskRequest {
	s.UserBid = &v
	return s
}

func (s *PushRpaTaskRequest) SetUserCallerParentId(v int64) *PushRpaTaskRequest {
	s.UserCallerParentId = &v
	return s
}

func (s *PushRpaTaskRequest) SetUserCallerType(v string) *PushRpaTaskRequest {
	s.UserCallerType = &v
	return s
}

func (s *PushRpaTaskRequest) SetUserClientIp(v string) *PushRpaTaskRequest {
	s.UserClientIp = &v
	return s
}

func (s *PushRpaTaskRequest) SetUserKp(v string) *PushRpaTaskRequest {
	s.UserKp = &v
	return s
}

func (s *PushRpaTaskRequest) SetUserMfaPresent(v bool) *PushRpaTaskRequest {
	s.UserMfaPresent = &v
	return s
}

func (s *PushRpaTaskRequest) SetUserSecurityToken(v string) *PushRpaTaskRequest {
	s.UserSecurityToken = &v
	return s
}

func (s *PushRpaTaskRequest) Validate() error {
	return dara.Validate(s)
}
