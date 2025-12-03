// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCallbackTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunKp(v string) *CallbackTaskRequest
	GetAliyunKp() *string
	SetApiType(v string) *CallbackTaskRequest
	GetApiType() *string
	SetBid(v string) *CallbackTaskRequest
	GetBid() *string
	SetBizCode(v string) *CallbackTaskRequest
	GetBizCode() *string
	SetLang(v string) *CallbackTaskRequest
	GetLang() *string
	SetOrderId(v string) *CallbackTaskRequest
	GetOrderId() *string
	SetOriginalRequest(v string) *CallbackTaskRequest
	GetOriginalRequest() *string
	SetOutTaskId(v string) *CallbackTaskRequest
	GetOutTaskId() *string
	SetPrincipalKey(v string) *CallbackTaskRequest
	GetPrincipalKey() *string
	SetTaskData(v string) *CallbackTaskRequest
	GetTaskData() *string
	SetTaskId(v string) *CallbackTaskRequest
	GetTaskId() *string
	SetTaskType(v string) *CallbackTaskRequest
	GetTaskType() *string
	SetUserAccessKeyId(v string) *CallbackTaskRequest
	GetUserAccessKeyId() *string
	SetUserBid(v string) *CallbackTaskRequest
	GetUserBid() *string
	SetUserCallerParentId(v int64) *CallbackTaskRequest
	GetUserCallerParentId() *int64
	SetUserCallerSecurityTransport(v bool) *CallbackTaskRequest
	GetUserCallerSecurityTransport() *bool
	SetUserCallerType(v string) *CallbackTaskRequest
	GetUserCallerType() *string
	SetUserClientIp(v string) *CallbackTaskRequest
	GetUserClientIp() *string
	SetUserKp(v string) *CallbackTaskRequest
	GetUserKp() *string
	SetUserMfaPresent(v bool) *CallbackTaskRequest
	GetUserMfaPresent() *bool
	SetUserSecurityToken(v string) *CallbackTaskRequest
	GetUserSecurityToken() *string
}

type CallbackTaskRequest struct {
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
	Bid     *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// lang
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// orderId
	//
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
	// outTaskId
	//
	// example:
	//
	// 1
	OutTaskId    *string `json:"OutTaskId,omitempty" xml:"OutTaskId,omitempty"`
	PrincipalKey *string `json:"PrincipalKey,omitempty" xml:"PrincipalKey,omitempty"`
	// taskData
	//
	// example:
	//
	// {\\"result\\":\\"SUCCESS\\",\\"message\\":\\"null\\",\\"taskId\\":\\"8cbc97d8-9b2b-4c2f-862f-983ea5dbedc2\\"}
	TaskData *string `json:"TaskData,omitempty" xml:"TaskData,omitempty"`
	// taskId
	//
	// example:
	//
	// 2559492
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// taskType
	//
	// example:
	//
	// PATENT_QUERY
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

func (s CallbackTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CallbackTaskRequest) GoString() string {
	return s.String()
}

func (s *CallbackTaskRequest) GetAliyunKp() *string {
	return s.AliyunKp
}

func (s *CallbackTaskRequest) GetApiType() *string {
	return s.ApiType
}

func (s *CallbackTaskRequest) GetBid() *string {
	return s.Bid
}

func (s *CallbackTaskRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *CallbackTaskRequest) GetLang() *string {
	return s.Lang
}

func (s *CallbackTaskRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *CallbackTaskRequest) GetOriginalRequest() *string {
	return s.OriginalRequest
}

func (s *CallbackTaskRequest) GetOutTaskId() *string {
	return s.OutTaskId
}

func (s *CallbackTaskRequest) GetPrincipalKey() *string {
	return s.PrincipalKey
}

func (s *CallbackTaskRequest) GetTaskData() *string {
	return s.TaskData
}

func (s *CallbackTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *CallbackTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *CallbackTaskRequest) GetUserAccessKeyId() *string {
	return s.UserAccessKeyId
}

func (s *CallbackTaskRequest) GetUserBid() *string {
	return s.UserBid
}

func (s *CallbackTaskRequest) GetUserCallerParentId() *int64 {
	return s.UserCallerParentId
}

func (s *CallbackTaskRequest) GetUserCallerSecurityTransport() *bool {
	return s.UserCallerSecurityTransport
}

func (s *CallbackTaskRequest) GetUserCallerType() *string {
	return s.UserCallerType
}

func (s *CallbackTaskRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *CallbackTaskRequest) GetUserKp() *string {
	return s.UserKp
}

func (s *CallbackTaskRequest) GetUserMfaPresent() *bool {
	return s.UserMfaPresent
}

func (s *CallbackTaskRequest) GetUserSecurityToken() *string {
	return s.UserSecurityToken
}

func (s *CallbackTaskRequest) SetAliyunKp(v string) *CallbackTaskRequest {
	s.AliyunKp = &v
	return s
}

func (s *CallbackTaskRequest) SetApiType(v string) *CallbackTaskRequest {
	s.ApiType = &v
	return s
}

func (s *CallbackTaskRequest) SetBid(v string) *CallbackTaskRequest {
	s.Bid = &v
	return s
}

func (s *CallbackTaskRequest) SetBizCode(v string) *CallbackTaskRequest {
	s.BizCode = &v
	return s
}

func (s *CallbackTaskRequest) SetLang(v string) *CallbackTaskRequest {
	s.Lang = &v
	return s
}

func (s *CallbackTaskRequest) SetOrderId(v string) *CallbackTaskRequest {
	s.OrderId = &v
	return s
}

func (s *CallbackTaskRequest) SetOriginalRequest(v string) *CallbackTaskRequest {
	s.OriginalRequest = &v
	return s
}

func (s *CallbackTaskRequest) SetOutTaskId(v string) *CallbackTaskRequest {
	s.OutTaskId = &v
	return s
}

func (s *CallbackTaskRequest) SetPrincipalKey(v string) *CallbackTaskRequest {
	s.PrincipalKey = &v
	return s
}

func (s *CallbackTaskRequest) SetTaskData(v string) *CallbackTaskRequest {
	s.TaskData = &v
	return s
}

func (s *CallbackTaskRequest) SetTaskId(v string) *CallbackTaskRequest {
	s.TaskId = &v
	return s
}

func (s *CallbackTaskRequest) SetTaskType(v string) *CallbackTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CallbackTaskRequest) SetUserAccessKeyId(v string) *CallbackTaskRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *CallbackTaskRequest) SetUserBid(v string) *CallbackTaskRequest {
	s.UserBid = &v
	return s
}

func (s *CallbackTaskRequest) SetUserCallerParentId(v int64) *CallbackTaskRequest {
	s.UserCallerParentId = &v
	return s
}

func (s *CallbackTaskRequest) SetUserCallerSecurityTransport(v bool) *CallbackTaskRequest {
	s.UserCallerSecurityTransport = &v
	return s
}

func (s *CallbackTaskRequest) SetUserCallerType(v string) *CallbackTaskRequest {
	s.UserCallerType = &v
	return s
}

func (s *CallbackTaskRequest) SetUserClientIp(v string) *CallbackTaskRequest {
	s.UserClientIp = &v
	return s
}

func (s *CallbackTaskRequest) SetUserKp(v string) *CallbackTaskRequest {
	s.UserKp = &v
	return s
}

func (s *CallbackTaskRequest) SetUserMfaPresent(v bool) *CallbackTaskRequest {
	s.UserMfaPresent = &v
	return s
}

func (s *CallbackTaskRequest) SetUserSecurityToken(v string) *CallbackTaskRequest {
	s.UserSecurityToken = &v
	return s
}

func (s *CallbackTaskRequest) Validate() error {
	return dara.Validate(s)
}
