// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushRpaTaskDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunKp(v string) *PushRpaTaskDetailRequest
	GetAliyunKp() *string
	SetApiType(v string) *PushRpaTaskDetailRequest
	GetApiType() *string
	SetBid(v string) *PushRpaTaskDetailRequest
	GetBid() *string
	SetInputData(v string) *PushRpaTaskDetailRequest
	GetInputData() *string
	SetInputHtml(v string) *PushRpaTaskDetailRequest
	GetInputHtml() *string
	SetInputScreenshot(v string) *PushRpaTaskDetailRequest
	GetInputScreenshot() *string
	SetLang(v string) *PushRpaTaskDetailRequest
	GetLang() *string
	SetModelDetailId(v int64) *PushRpaTaskDetailRequest
	GetModelDetailId() *int64
	SetOriginalRequest(v string) *PushRpaTaskDetailRequest
	GetOriginalRequest() *string
	SetOutputData(v string) *PushRpaTaskDetailRequest
	GetOutputData() *string
	SetOutputHtml(v string) *PushRpaTaskDetailRequest
	GetOutputHtml() *string
	SetOutputScreenshot(v string) *PushRpaTaskDetailRequest
	GetOutputScreenshot() *string
	SetStatus(v int32) *PushRpaTaskDetailRequest
	GetStatus() *int32
	SetTaskDetailId(v int64) *PushRpaTaskDetailRequest
	GetTaskDetailId() *int64
	SetTaskId(v int64) *PushRpaTaskDetailRequest
	GetTaskId() *int64
	SetUserAccessKeyId(v string) *PushRpaTaskDetailRequest
	GetUserAccessKeyId() *string
	SetUserBid(v string) *PushRpaTaskDetailRequest
	GetUserBid() *string
	SetUserCallerParentId(v int64) *PushRpaTaskDetailRequest
	GetUserCallerParentId() *int64
	SetUserCallerType(v string) *PushRpaTaskDetailRequest
	GetUserCallerType() *string
	SetUserClientIp(v string) *PushRpaTaskDetailRequest
	GetUserClientIp() *string
	SetUserKp(v string) *PushRpaTaskDetailRequest
	GetUserKp() *string
	SetUserSecurityToken(v string) *PushRpaTaskDetailRequest
	GetUserSecurityToken() *string
}

type PushRpaTaskDetailRequest struct {
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
	// inputData
	//
	// example:
	//
	// http://wssq.sbj.cnipa.gov.cn:9080/tmsve/wssqsy_getCayzDl.xhtml
	InputData *string `json:"InputData,omitempty" xml:"InputData,omitempty"`
	// inputHtml
	//
	// example:
	//
	// 1
	InputHtml *string `json:"InputHtml,omitempty" xml:"InputHtml,omitempty"`
	// inputScreenshot
	//
	// example:
	//
	// 1
	InputScreenshot *string `json:"InputScreenshot,omitempty" xml:"InputScreenshot,omitempty"`
	// lang
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// modelDetailId
	//
	// example:
	//
	// 6
	ModelDetailId *int64 `json:"ModelDetailId,omitempty" xml:"ModelDetailId,omitempty"`
	// originalRequest
	//
	// example:
	//
	// 1
	OriginalRequest *string `json:"OriginalRequest,omitempty" xml:"OriginalRequest,omitempty"`
	// outputData
	OutputData *string `json:"OutputData,omitempty" xml:"OutputData,omitempty"`
	// outputHtml
	//
	// example:
	//
	// <div class=\\"photobox\\" id=\\"Layer3\\" style=\\"visibility: visible
	OutputHtml *string `json:"OutputHtml,omitempty" xml:"OutputHtml,omitempty"`
	// outputScreenshot
	//
	// example:
	//
	// http://dbu-nap-p-test.oss-cn-beijing.aliyuncs.com/202301/20230110/5782089/1673334129101-d111874e-f181-4d1c-8edc-83e789975329.jpg?Expires=1675926129&OSSAccessKeyId=hObpgEXoca42qH3V&Signature=------
	OutputScreenshot *string `json:"OutputScreenshot,omitempty" xml:"OutputScreenshot,omitempty"`
	// status
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// taskDetailId
	//
	// example:
	//
	// 1
	TaskDetailId *int64 `json:"TaskDetailId,omitempty" xml:"TaskDetailId,omitempty"`
	// taskId
	//
	// example:
	//
	// 5596654
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	// userSecurityToken
	//
	// example:
	//
	// 1
	UserSecurityToken *string `json:"UserSecurityToken,omitempty" xml:"UserSecurityToken,omitempty"`
}

func (s PushRpaTaskDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s PushRpaTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *PushRpaTaskDetailRequest) GetAliyunKp() *string {
	return s.AliyunKp
}

func (s *PushRpaTaskDetailRequest) GetApiType() *string {
	return s.ApiType
}

func (s *PushRpaTaskDetailRequest) GetBid() *string {
	return s.Bid
}

func (s *PushRpaTaskDetailRequest) GetInputData() *string {
	return s.InputData
}

func (s *PushRpaTaskDetailRequest) GetInputHtml() *string {
	return s.InputHtml
}

func (s *PushRpaTaskDetailRequest) GetInputScreenshot() *string {
	return s.InputScreenshot
}

func (s *PushRpaTaskDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *PushRpaTaskDetailRequest) GetModelDetailId() *int64 {
	return s.ModelDetailId
}

func (s *PushRpaTaskDetailRequest) GetOriginalRequest() *string {
	return s.OriginalRequest
}

func (s *PushRpaTaskDetailRequest) GetOutputData() *string {
	return s.OutputData
}

func (s *PushRpaTaskDetailRequest) GetOutputHtml() *string {
	return s.OutputHtml
}

func (s *PushRpaTaskDetailRequest) GetOutputScreenshot() *string {
	return s.OutputScreenshot
}

func (s *PushRpaTaskDetailRequest) GetStatus() *int32 {
	return s.Status
}

func (s *PushRpaTaskDetailRequest) GetTaskDetailId() *int64 {
	return s.TaskDetailId
}

func (s *PushRpaTaskDetailRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *PushRpaTaskDetailRequest) GetUserAccessKeyId() *string {
	return s.UserAccessKeyId
}

func (s *PushRpaTaskDetailRequest) GetUserBid() *string {
	return s.UserBid
}

func (s *PushRpaTaskDetailRequest) GetUserCallerParentId() *int64 {
	return s.UserCallerParentId
}

func (s *PushRpaTaskDetailRequest) GetUserCallerType() *string {
	return s.UserCallerType
}

func (s *PushRpaTaskDetailRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *PushRpaTaskDetailRequest) GetUserKp() *string {
	return s.UserKp
}

func (s *PushRpaTaskDetailRequest) GetUserSecurityToken() *string {
	return s.UserSecurityToken
}

func (s *PushRpaTaskDetailRequest) SetAliyunKp(v string) *PushRpaTaskDetailRequest {
	s.AliyunKp = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetApiType(v string) *PushRpaTaskDetailRequest {
	s.ApiType = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetBid(v string) *PushRpaTaskDetailRequest {
	s.Bid = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetInputData(v string) *PushRpaTaskDetailRequest {
	s.InputData = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetInputHtml(v string) *PushRpaTaskDetailRequest {
	s.InputHtml = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetInputScreenshot(v string) *PushRpaTaskDetailRequest {
	s.InputScreenshot = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetLang(v string) *PushRpaTaskDetailRequest {
	s.Lang = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetModelDetailId(v int64) *PushRpaTaskDetailRequest {
	s.ModelDetailId = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetOriginalRequest(v string) *PushRpaTaskDetailRequest {
	s.OriginalRequest = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetOutputData(v string) *PushRpaTaskDetailRequest {
	s.OutputData = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetOutputHtml(v string) *PushRpaTaskDetailRequest {
	s.OutputHtml = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetOutputScreenshot(v string) *PushRpaTaskDetailRequest {
	s.OutputScreenshot = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetStatus(v int32) *PushRpaTaskDetailRequest {
	s.Status = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetTaskDetailId(v int64) *PushRpaTaskDetailRequest {
	s.TaskDetailId = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetTaskId(v int64) *PushRpaTaskDetailRequest {
	s.TaskId = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetUserAccessKeyId(v string) *PushRpaTaskDetailRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetUserBid(v string) *PushRpaTaskDetailRequest {
	s.UserBid = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetUserCallerParentId(v int64) *PushRpaTaskDetailRequest {
	s.UserCallerParentId = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetUserCallerType(v string) *PushRpaTaskDetailRequest {
	s.UserCallerType = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetUserClientIp(v string) *PushRpaTaskDetailRequest {
	s.UserClientIp = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetUserKp(v string) *PushRpaTaskDetailRequest {
	s.UserKp = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetUserSecurityToken(v string) *PushRpaTaskDetailRequest {
	s.UserSecurityToken = &v
	return s
}

func (s *PushRpaTaskDetailRequest) Validate() error {
	return dara.Validate(s)
}
