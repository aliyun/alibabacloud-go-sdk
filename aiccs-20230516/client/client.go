// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddBlacklistRequest struct {
	// 有效天数
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ExpiredDay *string `json:"ExpiredDay,omitempty" xml:"ExpiredDay,omitempty"`
	// 号码列表
	//
	// This parameter is required.
	Numbers []*string `json:"Numbers,omitempty" xml:"Numbers,omitempty" type:"Repeated"`
	OwnerId *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 备注
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AddBlacklistRequest) String() string {
	return tea.Prettify(s)
}

func (s AddBlacklistRequest) GoString() string {
	return s.String()
}

func (s *AddBlacklistRequest) SetExpiredDay(v string) *AddBlacklistRequest {
	s.ExpiredDay = &v
	return s
}

func (s *AddBlacklistRequest) SetNumbers(v []*string) *AddBlacklistRequest {
	s.Numbers = v
	return s
}

func (s *AddBlacklistRequest) SetOwnerId(v int64) *AddBlacklistRequest {
	s.OwnerId = &v
	return s
}

func (s *AddBlacklistRequest) SetRemark(v string) *AddBlacklistRequest {
	s.Remark = &v
	return s
}

func (s *AddBlacklistRequest) SetResourceOwnerAccount(v string) *AddBlacklistRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddBlacklistRequest) SetResourceOwnerId(v int64) *AddBlacklistRequest {
	s.ResourceOwnerId = &v
	return s
}

type AddBlacklistShrinkRequest struct {
	// 有效天数
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ExpiredDay *string `json:"ExpiredDay,omitempty" xml:"ExpiredDay,omitempty"`
	// 号码列表
	//
	// This parameter is required.
	NumbersShrink *string `json:"Numbers,omitempty" xml:"Numbers,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 备注
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AddBlacklistShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AddBlacklistShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddBlacklistShrinkRequest) SetExpiredDay(v string) *AddBlacklistShrinkRequest {
	s.ExpiredDay = &v
	return s
}

func (s *AddBlacklistShrinkRequest) SetNumbersShrink(v string) *AddBlacklistShrinkRequest {
	s.NumbersShrink = &v
	return s
}

func (s *AddBlacklistShrinkRequest) SetOwnerId(v int64) *AddBlacklistShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *AddBlacklistShrinkRequest) SetRemark(v string) *AddBlacklistShrinkRequest {
	s.Remark = &v
	return s
}

func (s *AddBlacklistShrinkRequest) SetResourceOwnerAccount(v string) *AddBlacklistShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddBlacklistShrinkRequest) SetResourceOwnerId(v int64) *AddBlacklistShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

type AddBlacklistResponseBody struct {
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *AddBlacklistResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 8EFC6D10-307B-1ECA-A8C6-7CBDF776AAD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1683440860035
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s AddBlacklistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddBlacklistResponseBody) GoString() string {
	return s.String()
}

func (s *AddBlacklistResponseBody) SetCode(v int64) *AddBlacklistResponseBody {
	s.Code = &v
	return s
}

func (s *AddBlacklistResponseBody) SetMessage(v string) *AddBlacklistResponseBody {
	s.Message = &v
	return s
}

func (s *AddBlacklistResponseBody) SetModel(v *AddBlacklistResponseBodyModel) *AddBlacklistResponseBody {
	s.Model = v
	return s
}

func (s *AddBlacklistResponseBody) SetRequestId(v string) *AddBlacklistResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddBlacklistResponseBody) SetSuccess(v string) *AddBlacklistResponseBody {
	s.Success = &v
	return s
}

func (s *AddBlacklistResponseBody) SetTimestamp(v int64) *AddBlacklistResponseBody {
	s.Timestamp = &v
	return s
}

type AddBlacklistResponseBodyModel struct {
	// 错误手机号
	UnHandleNumbers []*string `json:"UnHandleNumbers,omitempty" xml:"UnHandleNumbers,omitempty" type:"Repeated"`
}

func (s AddBlacklistResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s AddBlacklistResponseBodyModel) GoString() string {
	return s.String()
}

func (s *AddBlacklistResponseBodyModel) SetUnHandleNumbers(v []*string) *AddBlacklistResponseBodyModel {
	s.UnHandleNumbers = v
	return s
}

type AddBlacklistResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddBlacklistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddBlacklistResponse) String() string {
	return tea.Prettify(s)
}

func (s AddBlacklistResponse) GoString() string {
	return s.String()
}

func (s *AddBlacklistResponse) SetHeaders(v map[string]*string) *AddBlacklistResponse {
	s.Headers = v
	return s
}

func (s *AddBlacklistResponse) SetStatusCode(v int32) *AddBlacklistResponse {
	s.StatusCode = &v
	return s
}

func (s *AddBlacklistResponse) SetBody(v *AddBlacklistResponseBody) *AddBlacklistResponse {
	s.Body = v
	return s
}

type AddTaskRequest struct {
	// 外呼时间
	CallTimeList []*AddTaskRequestCallTimeList `json:"CallTimeList,omitempty" xml:"CallTimeList,omitempty" type:"Repeated"`
	// 回调地址
	//
	// example:
	//
	// 示例值示例值示例值
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// 当发送闪信配置为1时，闪信模板ID必填
	//
	// example:
	//
	// 96
	FlashSmsTemplateId *int64 `json:"FlashSmsTemplateId,omitempty" xml:"FlashSmsTemplateId,omitempty"`
	// 发送闪信配置
	//
	// example:
	//
	// 0
	FlashSmsType *int64 `json:"FlashSmsType,omitempty" xml:"FlashSmsType,omitempty"`
	// 并发数
	//
	// example:
	//
	// 99
	MaxConcurrency *int64 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// 任务名称
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 播放间隔时长
	//
	// example:
	//
	// 89
	PlaySleepVal *int64 `json:"PlaySleepVal,omitempty" xml:"PlaySleepVal,omitempty"`
	// 录音播放次数
	//
	// example:
	//
	// 51
	PlayTimes *int64 `json:"PlayTimes,omitempty" xml:"PlayTimes,omitempty"`
	// 重呼配置
	//
	// example:
	//
	// 53
	RecallType *int64 `json:"RecallType,omitempty" xml:"RecallType,omitempty"`
	// 录音地址
	//
	// example:
	//
	// 示例值
	RecordPath *string `json:"RecordPath,omitempty" xml:"RecordPath,omitempty"`
	// 重呼次数
	//
	// example:
	//
	// 37
	RepeatCount *int64 `json:"RepeatCount,omitempty" xml:"RepeatCount,omitempty"`
	// 重呼间隔
	//
	// example:
	//
	// 14
	RepeatInterval *int64 `json:"RepeatInterval,omitempty" xml:"RepeatInterval,omitempty"`
	// 重呼条件
	RepeatReason []*string `json:"RepeatReason,omitempty" xml:"RepeatReason,omitempty" type:"Repeated"`
	// 重呼时间
	RepeatTimes          []*string `json:"RepeatTimes,omitempty" xml:"RepeatTimes,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 短信发送规则
	SendSmsPlan []*AddTaskRequestSendSmsPlan `json:"SendSmsPlan,omitempty" xml:"SendSmsPlan,omitempty" type:"Repeated"`
	// 任务启动日期
	//
	// example:
	//
	// 2022-09-16
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 任务类型
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TaskType *int64 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// 话术模板ID
	//
	// example:
	//
	// 17
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 话术模板类型
	//
	// example:
	//
	// 1
	TemplateType *int64 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s AddTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s AddTaskRequest) GoString() string {
	return s.String()
}

func (s *AddTaskRequest) SetCallTimeList(v []*AddTaskRequestCallTimeList) *AddTaskRequest {
	s.CallTimeList = v
	return s
}

func (s *AddTaskRequest) SetCallbackUrl(v string) *AddTaskRequest {
	s.CallbackUrl = &v
	return s
}

func (s *AddTaskRequest) SetFlashSmsTemplateId(v int64) *AddTaskRequest {
	s.FlashSmsTemplateId = &v
	return s
}

func (s *AddTaskRequest) SetFlashSmsType(v int64) *AddTaskRequest {
	s.FlashSmsType = &v
	return s
}

func (s *AddTaskRequest) SetMaxConcurrency(v int64) *AddTaskRequest {
	s.MaxConcurrency = &v
	return s
}

func (s *AddTaskRequest) SetName(v string) *AddTaskRequest {
	s.Name = &v
	return s
}

func (s *AddTaskRequest) SetOwnerId(v int64) *AddTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *AddTaskRequest) SetPlaySleepVal(v int64) *AddTaskRequest {
	s.PlaySleepVal = &v
	return s
}

func (s *AddTaskRequest) SetPlayTimes(v int64) *AddTaskRequest {
	s.PlayTimes = &v
	return s
}

func (s *AddTaskRequest) SetRecallType(v int64) *AddTaskRequest {
	s.RecallType = &v
	return s
}

func (s *AddTaskRequest) SetRecordPath(v string) *AddTaskRequest {
	s.RecordPath = &v
	return s
}

func (s *AddTaskRequest) SetRepeatCount(v int64) *AddTaskRequest {
	s.RepeatCount = &v
	return s
}

func (s *AddTaskRequest) SetRepeatInterval(v int64) *AddTaskRequest {
	s.RepeatInterval = &v
	return s
}

func (s *AddTaskRequest) SetRepeatReason(v []*string) *AddTaskRequest {
	s.RepeatReason = v
	return s
}

func (s *AddTaskRequest) SetRepeatTimes(v []*string) *AddTaskRequest {
	s.RepeatTimes = v
	return s
}

func (s *AddTaskRequest) SetResourceOwnerAccount(v string) *AddTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddTaskRequest) SetResourceOwnerId(v int64) *AddTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddTaskRequest) SetSendSmsPlan(v []*AddTaskRequestSendSmsPlan) *AddTaskRequest {
	s.SendSmsPlan = v
	return s
}

func (s *AddTaskRequest) SetStartTime(v string) *AddTaskRequest {
	s.StartTime = &v
	return s
}

func (s *AddTaskRequest) SetTaskType(v int64) *AddTaskRequest {
	s.TaskType = &v
	return s
}

func (s *AddTaskRequest) SetTemplateId(v int64) *AddTaskRequest {
	s.TemplateId = &v
	return s
}

func (s *AddTaskRequest) SetTemplateType(v int64) *AddTaskRequest {
	s.TemplateType = &v
	return s
}

type AddTaskRequestCallTimeList struct {
	CallTime []*string `json:"CallTime,omitempty" xml:"CallTime,omitempty" type:"Repeated"`
}

func (s AddTaskRequestCallTimeList) String() string {
	return tea.Prettify(s)
}

func (s AddTaskRequestCallTimeList) GoString() string {
	return s.String()
}

func (s *AddTaskRequestCallTimeList) SetCallTime(v []*string) *AddTaskRequestCallTimeList {
	s.CallTime = v
	return s
}

type AddTaskRequestSendSmsPlan struct {
	// 意向标签
	IntentTags []*string `json:"IntentTags,omitempty" xml:"IntentTags,omitempty" type:"Repeated"`
	// 短信模板ID
	//
	// example:
	//
	// 71
	SmsTemplateId *int64 `json:"SmsTemplateId,omitempty" xml:"SmsTemplateId,omitempty"`
}

func (s AddTaskRequestSendSmsPlan) String() string {
	return tea.Prettify(s)
}

func (s AddTaskRequestSendSmsPlan) GoString() string {
	return s.String()
}

func (s *AddTaskRequestSendSmsPlan) SetIntentTags(v []*string) *AddTaskRequestSendSmsPlan {
	s.IntentTags = v
	return s
}

func (s *AddTaskRequestSendSmsPlan) SetSmsTemplateId(v int64) *AddTaskRequestSendSmsPlan {
	s.SmsTemplateId = &v
	return s
}

type AddTaskShrinkRequest struct {
	// 外呼时间
	CallTimeListShrink *string `json:"CallTimeList,omitempty" xml:"CallTimeList,omitempty"`
	// 回调地址
	//
	// example:
	//
	// 示例值示例值示例值
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// 当发送闪信配置为1时，闪信模板ID必填
	//
	// example:
	//
	// 96
	FlashSmsTemplateId *int64 `json:"FlashSmsTemplateId,omitempty" xml:"FlashSmsTemplateId,omitempty"`
	// 发送闪信配置
	//
	// example:
	//
	// 0
	FlashSmsType *int64 `json:"FlashSmsType,omitempty" xml:"FlashSmsType,omitempty"`
	// 并发数
	//
	// example:
	//
	// 99
	MaxConcurrency *int64 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// 任务名称
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 播放间隔时长
	//
	// example:
	//
	// 89
	PlaySleepVal *int64 `json:"PlaySleepVal,omitempty" xml:"PlaySleepVal,omitempty"`
	// 录音播放次数
	//
	// example:
	//
	// 51
	PlayTimes *int64 `json:"PlayTimes,omitempty" xml:"PlayTimes,omitempty"`
	// 重呼配置
	//
	// example:
	//
	// 53
	RecallType *int64 `json:"RecallType,omitempty" xml:"RecallType,omitempty"`
	// 录音地址
	//
	// example:
	//
	// 示例值
	RecordPath *string `json:"RecordPath,omitempty" xml:"RecordPath,omitempty"`
	// 重呼次数
	//
	// example:
	//
	// 37
	RepeatCount *int64 `json:"RepeatCount,omitempty" xml:"RepeatCount,omitempty"`
	// 重呼间隔
	//
	// example:
	//
	// 14
	RepeatInterval *int64 `json:"RepeatInterval,omitempty" xml:"RepeatInterval,omitempty"`
	// 重呼条件
	RepeatReasonShrink *string `json:"RepeatReason,omitempty" xml:"RepeatReason,omitempty"`
	// 重呼时间
	RepeatTimesShrink    *string `json:"RepeatTimes,omitempty" xml:"RepeatTimes,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 短信发送规则
	SendSmsPlanShrink *string `json:"SendSmsPlan,omitempty" xml:"SendSmsPlan,omitempty"`
	// 任务启动日期
	//
	// example:
	//
	// 2022-09-16
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 任务类型
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TaskType *int64 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// 话术模板ID
	//
	// example:
	//
	// 17
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 话术模板类型
	//
	// example:
	//
	// 1
	TemplateType *int64 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s AddTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AddTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddTaskShrinkRequest) SetCallTimeListShrink(v string) *AddTaskShrinkRequest {
	s.CallTimeListShrink = &v
	return s
}

func (s *AddTaskShrinkRequest) SetCallbackUrl(v string) *AddTaskShrinkRequest {
	s.CallbackUrl = &v
	return s
}

func (s *AddTaskShrinkRequest) SetFlashSmsTemplateId(v int64) *AddTaskShrinkRequest {
	s.FlashSmsTemplateId = &v
	return s
}

func (s *AddTaskShrinkRequest) SetFlashSmsType(v int64) *AddTaskShrinkRequest {
	s.FlashSmsType = &v
	return s
}

func (s *AddTaskShrinkRequest) SetMaxConcurrency(v int64) *AddTaskShrinkRequest {
	s.MaxConcurrency = &v
	return s
}

func (s *AddTaskShrinkRequest) SetName(v string) *AddTaskShrinkRequest {
	s.Name = &v
	return s
}

func (s *AddTaskShrinkRequest) SetOwnerId(v int64) *AddTaskShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *AddTaskShrinkRequest) SetPlaySleepVal(v int64) *AddTaskShrinkRequest {
	s.PlaySleepVal = &v
	return s
}

func (s *AddTaskShrinkRequest) SetPlayTimes(v int64) *AddTaskShrinkRequest {
	s.PlayTimes = &v
	return s
}

func (s *AddTaskShrinkRequest) SetRecallType(v int64) *AddTaskShrinkRequest {
	s.RecallType = &v
	return s
}

func (s *AddTaskShrinkRequest) SetRecordPath(v string) *AddTaskShrinkRequest {
	s.RecordPath = &v
	return s
}

func (s *AddTaskShrinkRequest) SetRepeatCount(v int64) *AddTaskShrinkRequest {
	s.RepeatCount = &v
	return s
}

func (s *AddTaskShrinkRequest) SetRepeatInterval(v int64) *AddTaskShrinkRequest {
	s.RepeatInterval = &v
	return s
}

func (s *AddTaskShrinkRequest) SetRepeatReasonShrink(v string) *AddTaskShrinkRequest {
	s.RepeatReasonShrink = &v
	return s
}

func (s *AddTaskShrinkRequest) SetRepeatTimesShrink(v string) *AddTaskShrinkRequest {
	s.RepeatTimesShrink = &v
	return s
}

func (s *AddTaskShrinkRequest) SetResourceOwnerAccount(v string) *AddTaskShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddTaskShrinkRequest) SetResourceOwnerId(v int64) *AddTaskShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddTaskShrinkRequest) SetSendSmsPlanShrink(v string) *AddTaskShrinkRequest {
	s.SendSmsPlanShrink = &v
	return s
}

func (s *AddTaskShrinkRequest) SetStartTime(v string) *AddTaskShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *AddTaskShrinkRequest) SetTaskType(v int64) *AddTaskShrinkRequest {
	s.TaskType = &v
	return s
}

func (s *AddTaskShrinkRequest) SetTemplateId(v int64) *AddTaskShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *AddTaskShrinkRequest) SetTemplateType(v int64) *AddTaskShrinkRequest {
	s.TemplateType = &v
	return s
}

type AddTaskResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值
	Message *string                   `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *AddTaskResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 5453cc9b-02bc-4dbb-9527-f28a9100b912
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1686225227338
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s AddTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddTaskResponseBody) GoString() string {
	return s.String()
}

func (s *AddTaskResponseBody) SetAccessDeniedDetail(v string) *AddTaskResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AddTaskResponseBody) SetCode(v int64) *AddTaskResponseBody {
	s.Code = &v
	return s
}

func (s *AddTaskResponseBody) SetMessage(v string) *AddTaskResponseBody {
	s.Message = &v
	return s
}

func (s *AddTaskResponseBody) SetModel(v *AddTaskResponseBodyModel) *AddTaskResponseBody {
	s.Model = v
	return s
}

func (s *AddTaskResponseBody) SetRequestId(v string) *AddTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTaskResponseBody) SetSuccess(v bool) *AddTaskResponseBody {
	s.Success = &v
	return s
}

func (s *AddTaskResponseBody) SetTimestamp(v int64) *AddTaskResponseBody {
	s.Timestamp = &v
	return s
}

type AddTaskResponseBodyModel struct {
	// 任务ID
	//
	// example:
	//
	// 47
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AddTaskResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s AddTaskResponseBodyModel) GoString() string {
	return s.String()
}

func (s *AddTaskResponseBodyModel) SetTaskId(v int64) *AddTaskResponseBodyModel {
	s.TaskId = &v
	return s
}

type AddTaskResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s AddTaskResponse) GoString() string {
	return s.String()
}

func (s *AddTaskResponse) SetHeaders(v map[string]*string) *AddTaskResponse {
	s.Headers = v
	return s
}

func (s *AddTaskResponse) SetStatusCode(v int32) *AddTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTaskResponse) SetBody(v *AddTaskResponseBody) *AddTaskResponse {
	s.Body = v
	return s
}

type AgentCancelCallRequest struct {
	// 坐席ID
	//
	// example:
	//
	// 64
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// 坐席标签
	//
	// example:
	//
	// abc
	AgentTag *string `json:"AgentTag,omitempty" xml:"AgentTag,omitempty"`
	// 号码列表
	Numbers              []*string `json:"Numbers,omitempty" xml:"Numbers,omitempty" type:"Repeated"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 用户自定义标签列表
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s AgentCancelCallRequest) String() string {
	return tea.Prettify(s)
}

func (s AgentCancelCallRequest) GoString() string {
	return s.String()
}

func (s *AgentCancelCallRequest) SetAgentId(v int64) *AgentCancelCallRequest {
	s.AgentId = &v
	return s
}

func (s *AgentCancelCallRequest) SetAgentTag(v string) *AgentCancelCallRequest {
	s.AgentTag = &v
	return s
}

func (s *AgentCancelCallRequest) SetNumbers(v []*string) *AgentCancelCallRequest {
	s.Numbers = v
	return s
}

func (s *AgentCancelCallRequest) SetOwnerId(v int64) *AgentCancelCallRequest {
	s.OwnerId = &v
	return s
}

func (s *AgentCancelCallRequest) SetResourceOwnerAccount(v string) *AgentCancelCallRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AgentCancelCallRequest) SetResourceOwnerId(v int64) *AgentCancelCallRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AgentCancelCallRequest) SetTags(v []*string) *AgentCancelCallRequest {
	s.Tags = v
	return s
}

type AgentCancelCallShrinkRequest struct {
	// 坐席ID
	//
	// example:
	//
	// 64
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// 坐席标签
	//
	// example:
	//
	// abc
	AgentTag *string `json:"AgentTag,omitempty" xml:"AgentTag,omitempty"`
	// 号码列表
	NumbersShrink        *string `json:"Numbers,omitempty" xml:"Numbers,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 用户自定义标签列表
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s AgentCancelCallShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AgentCancelCallShrinkRequest) GoString() string {
	return s.String()
}

func (s *AgentCancelCallShrinkRequest) SetAgentId(v int64) *AgentCancelCallShrinkRequest {
	s.AgentId = &v
	return s
}

func (s *AgentCancelCallShrinkRequest) SetAgentTag(v string) *AgentCancelCallShrinkRequest {
	s.AgentTag = &v
	return s
}

func (s *AgentCancelCallShrinkRequest) SetNumbersShrink(v string) *AgentCancelCallShrinkRequest {
	s.NumbersShrink = &v
	return s
}

func (s *AgentCancelCallShrinkRequest) SetOwnerId(v int64) *AgentCancelCallShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *AgentCancelCallShrinkRequest) SetResourceOwnerAccount(v string) *AgentCancelCallShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AgentCancelCallShrinkRequest) SetResourceOwnerId(v int64) *AgentCancelCallShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AgentCancelCallShrinkRequest) SetTagsShrink(v string) *AgentCancelCallShrinkRequest {
	s.TagsShrink = &v
	return s
}

type AgentCancelCallResponseBody struct {
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值
	Message *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *AgentCancelCallResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 8EFC6D10-307B-1ECA-A8C6-7CBDF776AAD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1683440860035
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s AgentCancelCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AgentCancelCallResponseBody) GoString() string {
	return s.String()
}

func (s *AgentCancelCallResponseBody) SetCode(v int64) *AgentCancelCallResponseBody {
	s.Code = &v
	return s
}

func (s *AgentCancelCallResponseBody) SetMessage(v string) *AgentCancelCallResponseBody {
	s.Message = &v
	return s
}

func (s *AgentCancelCallResponseBody) SetModel(v *AgentCancelCallResponseBodyModel) *AgentCancelCallResponseBody {
	s.Model = v
	return s
}

func (s *AgentCancelCallResponseBody) SetRequestId(v string) *AgentCancelCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *AgentCancelCallResponseBody) SetSuccess(v string) *AgentCancelCallResponseBody {
	s.Success = &v
	return s
}

func (s *AgentCancelCallResponseBody) SetTimestamp(v int64) *AgentCancelCallResponseBody {
	s.Timestamp = &v
	return s
}

type AgentCancelCallResponseBodyModel struct {
	// 错误手机列表
	UnHandleNumbers []*string `json:"UnHandleNumbers,omitempty" xml:"UnHandleNumbers,omitempty" type:"Repeated"`
}

func (s AgentCancelCallResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s AgentCancelCallResponseBodyModel) GoString() string {
	return s.String()
}

func (s *AgentCancelCallResponseBodyModel) SetUnHandleNumbers(v []*string) *AgentCancelCallResponseBodyModel {
	s.UnHandleNumbers = v
	return s
}

type AgentCancelCallResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgentCancelCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AgentCancelCallResponse) String() string {
	return tea.Prettify(s)
}

func (s AgentCancelCallResponse) GoString() string {
	return s.String()
}

func (s *AgentCancelCallResponse) SetHeaders(v map[string]*string) *AgentCancelCallResponse {
	s.Headers = v
	return s
}

func (s *AgentCancelCallResponse) SetStatusCode(v int32) *AgentCancelCallResponse {
	s.StatusCode = &v
	return s
}

func (s *AgentCancelCallResponse) SetBody(v *AgentCancelCallResponseBody) *AgentCancelCallResponse {
	s.Body = v
	return s
}

type AgentRecoverCallRequest struct {
	// 坐席ID
	//
	// example:
	//
	// 5
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// 坐席标签
	//
	// example:
	//
	// abc
	AgentTag *string `json:"AgentTag,omitempty" xml:"AgentTag,omitempty"`
	// 查询开始导入时间
	//
	// example:
	//
	// 2020-03-06 10:10:10
	BeginImportTime *string `json:"BeginImportTime,omitempty" xml:"BeginImportTime,omitempty"`
	// 查询结束导入时间
	//
	// example:
	//
	// 2021-03-06 10:10:10
	EndImportTime *string `json:"EndImportTime,omitempty" xml:"EndImportTime,omitempty"`
	// 号码列表
	Numbers              []*string `json:"Numbers,omitempty" xml:"Numbers,omitempty" type:"Repeated"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 用户自定义标签列表
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s AgentRecoverCallRequest) String() string {
	return tea.Prettify(s)
}

func (s AgentRecoverCallRequest) GoString() string {
	return s.String()
}

func (s *AgentRecoverCallRequest) SetAgentId(v int64) *AgentRecoverCallRequest {
	s.AgentId = &v
	return s
}

func (s *AgentRecoverCallRequest) SetAgentTag(v string) *AgentRecoverCallRequest {
	s.AgentTag = &v
	return s
}

func (s *AgentRecoverCallRequest) SetBeginImportTime(v string) *AgentRecoverCallRequest {
	s.BeginImportTime = &v
	return s
}

func (s *AgentRecoverCallRequest) SetEndImportTime(v string) *AgentRecoverCallRequest {
	s.EndImportTime = &v
	return s
}

func (s *AgentRecoverCallRequest) SetNumbers(v []*string) *AgentRecoverCallRequest {
	s.Numbers = v
	return s
}

func (s *AgentRecoverCallRequest) SetOwnerId(v int64) *AgentRecoverCallRequest {
	s.OwnerId = &v
	return s
}

func (s *AgentRecoverCallRequest) SetResourceOwnerAccount(v string) *AgentRecoverCallRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AgentRecoverCallRequest) SetResourceOwnerId(v int64) *AgentRecoverCallRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AgentRecoverCallRequest) SetTags(v []*string) *AgentRecoverCallRequest {
	s.Tags = v
	return s
}

type AgentRecoverCallShrinkRequest struct {
	// 坐席ID
	//
	// example:
	//
	// 5
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// 坐席标签
	//
	// example:
	//
	// abc
	AgentTag *string `json:"AgentTag,omitempty" xml:"AgentTag,omitempty"`
	// 查询开始导入时间
	//
	// example:
	//
	// 2020-03-06 10:10:10
	BeginImportTime *string `json:"BeginImportTime,omitempty" xml:"BeginImportTime,omitempty"`
	// 查询结束导入时间
	//
	// example:
	//
	// 2021-03-06 10:10:10
	EndImportTime *string `json:"EndImportTime,omitempty" xml:"EndImportTime,omitempty"`
	// 号码列表
	NumbersShrink        *string `json:"Numbers,omitempty" xml:"Numbers,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 用户自定义标签列表
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s AgentRecoverCallShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AgentRecoverCallShrinkRequest) GoString() string {
	return s.String()
}

func (s *AgentRecoverCallShrinkRequest) SetAgentId(v int64) *AgentRecoverCallShrinkRequest {
	s.AgentId = &v
	return s
}

func (s *AgentRecoverCallShrinkRequest) SetAgentTag(v string) *AgentRecoverCallShrinkRequest {
	s.AgentTag = &v
	return s
}

func (s *AgentRecoverCallShrinkRequest) SetBeginImportTime(v string) *AgentRecoverCallShrinkRequest {
	s.BeginImportTime = &v
	return s
}

func (s *AgentRecoverCallShrinkRequest) SetEndImportTime(v string) *AgentRecoverCallShrinkRequest {
	s.EndImportTime = &v
	return s
}

func (s *AgentRecoverCallShrinkRequest) SetNumbersShrink(v string) *AgentRecoverCallShrinkRequest {
	s.NumbersShrink = &v
	return s
}

func (s *AgentRecoverCallShrinkRequest) SetOwnerId(v int64) *AgentRecoverCallShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *AgentRecoverCallShrinkRequest) SetResourceOwnerAccount(v string) *AgentRecoverCallShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AgentRecoverCallShrinkRequest) SetResourceOwnerId(v int64) *AgentRecoverCallShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AgentRecoverCallShrinkRequest) SetTagsShrink(v string) *AgentRecoverCallShrinkRequest {
	s.TagsShrink = &v
	return s
}

type AgentRecoverCallResponseBody struct {
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *AgentRecoverCallResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 8EFC6D10-307B-1ECA-A8C6-7CBDF776AAD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 坐席标签
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1683440860035
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s AgentRecoverCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AgentRecoverCallResponseBody) GoString() string {
	return s.String()
}

func (s *AgentRecoverCallResponseBody) SetCode(v int64) *AgentRecoverCallResponseBody {
	s.Code = &v
	return s
}

func (s *AgentRecoverCallResponseBody) SetMessage(v string) *AgentRecoverCallResponseBody {
	s.Message = &v
	return s
}

func (s *AgentRecoverCallResponseBody) SetModel(v *AgentRecoverCallResponseBodyModel) *AgentRecoverCallResponseBody {
	s.Model = v
	return s
}

func (s *AgentRecoverCallResponseBody) SetRequestId(v string) *AgentRecoverCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *AgentRecoverCallResponseBody) SetSuccess(v string) *AgentRecoverCallResponseBody {
	s.Success = &v
	return s
}

func (s *AgentRecoverCallResponseBody) SetTimestamp(v int64) *AgentRecoverCallResponseBody {
	s.Timestamp = &v
	return s
}

type AgentRecoverCallResponseBodyModel struct {
	// 错误手机列表
	UnHandleNumbers []*string `json:"UnHandleNumbers,omitempty" xml:"UnHandleNumbers,omitempty" type:"Repeated"`
}

func (s AgentRecoverCallResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s AgentRecoverCallResponseBodyModel) GoString() string {
	return s.String()
}

func (s *AgentRecoverCallResponseBodyModel) SetUnHandleNumbers(v []*string) *AgentRecoverCallResponseBodyModel {
	s.UnHandleNumbers = v
	return s
}

type AgentRecoverCallResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgentRecoverCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AgentRecoverCallResponse) String() string {
	return tea.Prettify(s)
}

func (s AgentRecoverCallResponse) GoString() string {
	return s.String()
}

func (s *AgentRecoverCallResponse) SetHeaders(v map[string]*string) *AgentRecoverCallResponse {
	s.Headers = v
	return s
}

func (s *AgentRecoverCallResponse) SetStatusCode(v int32) *AgentRecoverCallResponse {
	s.StatusCode = &v
	return s
}

func (s *AgentRecoverCallResponse) SetBody(v *AgentRecoverCallResponseBody) *AgentRecoverCallResponse {
	s.Body = v
	return s
}

type DetailsRequest struct {
	// 批次id
	//
	// example:
	//
	// 75
	BatchId *int64 `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// 结束导入时间
	//
	// example:
	//
	// 2023-04-25 15:19:02
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 号码状态
	//
	// example:
	//
	// 1
	NumberStatus *int64 `json:"NumberStatus,omitempty" xml:"NumberStatus,omitempty"`
	// 号码列表
	Numbers []*string `json:"Numbers,omitempty" xml:"Numbers,omitempty" type:"Repeated"`
	OwnerId *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 页数
	//
	// example:
	//
	// 77
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// 每页条数
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 开始导入时间
	//
	// example:
	//
	// 2023-04-25 15:19:02
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 任务id
	//
	// This parameter is required.
	//
	// example:
	//
	// 90
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s DetailsRequest) GoString() string {
	return s.String()
}

func (s *DetailsRequest) SetBatchId(v int64) *DetailsRequest {
	s.BatchId = &v
	return s
}

func (s *DetailsRequest) SetEndTime(v string) *DetailsRequest {
	s.EndTime = &v
	return s
}

func (s *DetailsRequest) SetNumberStatus(v int64) *DetailsRequest {
	s.NumberStatus = &v
	return s
}

func (s *DetailsRequest) SetNumbers(v []*string) *DetailsRequest {
	s.Numbers = v
	return s
}

func (s *DetailsRequest) SetOwnerId(v int64) *DetailsRequest {
	s.OwnerId = &v
	return s
}

func (s *DetailsRequest) SetPageNo(v int64) *DetailsRequest {
	s.PageNo = &v
	return s
}

func (s *DetailsRequest) SetPageSize(v int64) *DetailsRequest {
	s.PageSize = &v
	return s
}

func (s *DetailsRequest) SetResourceOwnerAccount(v string) *DetailsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DetailsRequest) SetResourceOwnerId(v int64) *DetailsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DetailsRequest) SetStartTime(v string) *DetailsRequest {
	s.StartTime = &v
	return s
}

func (s *DetailsRequest) SetTaskId(v int64) *DetailsRequest {
	s.TaskId = &v
	return s
}

type DetailsShrinkRequest struct {
	// 批次id
	//
	// example:
	//
	// 75
	BatchId *int64 `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// 结束导入时间
	//
	// example:
	//
	// 2023-04-25 15:19:02
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 号码状态
	//
	// example:
	//
	// 1
	NumberStatus *int64 `json:"NumberStatus,omitempty" xml:"NumberStatus,omitempty"`
	// 号码列表
	NumbersShrink *string `json:"Numbers,omitempty" xml:"Numbers,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 页数
	//
	// example:
	//
	// 77
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// 每页条数
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 开始导入时间
	//
	// example:
	//
	// 2023-04-25 15:19:02
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 任务id
	//
	// This parameter is required.
	//
	// example:
	//
	// 90
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DetailsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DetailsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetailsShrinkRequest) SetBatchId(v int64) *DetailsShrinkRequest {
	s.BatchId = &v
	return s
}

func (s *DetailsShrinkRequest) SetEndTime(v string) *DetailsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *DetailsShrinkRequest) SetNumberStatus(v int64) *DetailsShrinkRequest {
	s.NumberStatus = &v
	return s
}

func (s *DetailsShrinkRequest) SetNumbersShrink(v string) *DetailsShrinkRequest {
	s.NumbersShrink = &v
	return s
}

func (s *DetailsShrinkRequest) SetOwnerId(v int64) *DetailsShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *DetailsShrinkRequest) SetPageNo(v int64) *DetailsShrinkRequest {
	s.PageNo = &v
	return s
}

func (s *DetailsShrinkRequest) SetPageSize(v int64) *DetailsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DetailsShrinkRequest) SetResourceOwnerAccount(v string) *DetailsShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DetailsShrinkRequest) SetResourceOwnerId(v int64) *DetailsShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DetailsShrinkRequest) SetStartTime(v string) *DetailsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *DetailsShrinkRequest) SetTaskId(v int64) *DetailsShrinkRequest {
	s.TaskId = &v
	return s
}

type DetailsResponseBody struct {
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值
	Message *string                   `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *DetailsResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 389b2d0a-37e2-406d-b576-1fc0827be46a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1686279332221
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DetailsResponseBody) SetCode(v int64) *DetailsResponseBody {
	s.Code = &v
	return s
}

func (s *DetailsResponseBody) SetMessage(v string) *DetailsResponseBody {
	s.Message = &v
	return s
}

func (s *DetailsResponseBody) SetModel(v *DetailsResponseBodyModel) *DetailsResponseBody {
	s.Model = v
	return s
}

func (s *DetailsResponseBody) SetRequestId(v string) *DetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetailsResponseBody) SetSuccess(v bool) *DetailsResponseBody {
	s.Success = &v
	return s
}

func (s *DetailsResponseBody) SetTimestamp(v int64) *DetailsResponseBody {
	s.Timestamp = &v
	return s
}

type DetailsResponseBodyModel struct {
	List []*DetailsResponseBodyModelList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 94
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 79
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 71.8132
	TotalPage *float32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DetailsResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s DetailsResponseBodyModel) GoString() string {
	return s.String()
}

func (s *DetailsResponseBodyModel) SetList(v []*DetailsResponseBodyModelList) *DetailsResponseBodyModel {
	s.List = v
	return s
}

func (s *DetailsResponseBodyModel) SetPageNo(v int64) *DetailsResponseBodyModel {
	s.PageNo = &v
	return s
}

func (s *DetailsResponseBodyModel) SetPageSize(v int64) *DetailsResponseBodyModel {
	s.PageSize = &v
	return s
}

func (s *DetailsResponseBodyModel) SetTotalCount(v int64) *DetailsResponseBodyModel {
	s.TotalCount = &v
	return s
}

func (s *DetailsResponseBodyModel) SetTotalPage(v float32) *DetailsResponseBodyModel {
	s.TotalPage = &v
	return s
}

type DetailsResponseBodyModelList struct {
	// 批次号
	//
	// example:
	//
	// 27
	BatchId *int64 `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// 呼叫状态描述
	//
	// example:
	//
	// 示例值示例值
	CallDesc *string `json:"CallDesc,omitempty" xml:"CallDesc,omitempty"`
	// 外呼ID
	//
	// example:
	//
	// 28dd74a4-30ec-43c0-9bec-272924c51eeb
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// 呼叫状态
	//
	// example:
	//
	// 1
	CallStatus *int64 `json:"CallStatus,omitempty" xml:"CallStatus,omitempty"`
	// 外呼类型
	//
	// example:
	//
	// 2001
	CallType *int64 `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// 导入时间
	//
	// example:
	//
	// 2023-04-25 15:19:02
	ImportTime *string `json:"ImportTime,omitempty" xml:"ImportTime,omitempty"`
	// 拦截原因
	//
	// example:
	//
	// 示例值示例值示例值
	InterceptReason *string `json:"InterceptReason,omitempty" xml:"InterceptReason,omitempty"`
	// 外呼号码
	//
	// example:
	//
	// 188******454
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// 号码状态描述
	//
	// example:
	//
	// 示例值
	NumberDesc *string `json:"NumberDesc,omitempty" xml:"NumberDesc,omitempty"`
	// 外呼号码MD5
	//
	// example:
	//
	// cbe1b40f76f2cca4735954886b706ffa
	NumberMD5 *string `json:"NumberMD5,omitempty" xml:"NumberMD5,omitempty"`
	// 号码状态
	//
	// example:
	//
	// 1
	NumberStatus *int64 `json:"NumberStatus,omitempty" xml:"NumberStatus,omitempty"`
	// 用户自定义标签
	//
	// example:
	//
	// A
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// 任务ID
	//
	// example:
	//
	// 28
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DetailsResponseBodyModelList) String() string {
	return tea.Prettify(s)
}

func (s DetailsResponseBodyModelList) GoString() string {
	return s.String()
}

func (s *DetailsResponseBodyModelList) SetBatchId(v int64) *DetailsResponseBodyModelList {
	s.BatchId = &v
	return s
}

func (s *DetailsResponseBodyModelList) SetCallDesc(v string) *DetailsResponseBodyModelList {
	s.CallDesc = &v
	return s
}

func (s *DetailsResponseBodyModelList) SetCallId(v string) *DetailsResponseBodyModelList {
	s.CallId = &v
	return s
}

func (s *DetailsResponseBodyModelList) SetCallStatus(v int64) *DetailsResponseBodyModelList {
	s.CallStatus = &v
	return s
}

func (s *DetailsResponseBodyModelList) SetCallType(v int64) *DetailsResponseBodyModelList {
	s.CallType = &v
	return s
}

func (s *DetailsResponseBodyModelList) SetImportTime(v string) *DetailsResponseBodyModelList {
	s.ImportTime = &v
	return s
}

func (s *DetailsResponseBodyModelList) SetInterceptReason(v string) *DetailsResponseBodyModelList {
	s.InterceptReason = &v
	return s
}

func (s *DetailsResponseBodyModelList) SetNumber(v string) *DetailsResponseBodyModelList {
	s.Number = &v
	return s
}

func (s *DetailsResponseBodyModelList) SetNumberDesc(v string) *DetailsResponseBodyModelList {
	s.NumberDesc = &v
	return s
}

func (s *DetailsResponseBodyModelList) SetNumberMD5(v string) *DetailsResponseBodyModelList {
	s.NumberMD5 = &v
	return s
}

func (s *DetailsResponseBodyModelList) SetNumberStatus(v int64) *DetailsResponseBodyModelList {
	s.NumberStatus = &v
	return s
}

func (s *DetailsResponseBodyModelList) SetTag(v string) *DetailsResponseBodyModelList {
	s.Tag = &v
	return s
}

func (s *DetailsResponseBodyModelList) SetTaskId(v int64) *DetailsResponseBodyModelList {
	s.TaskId = &v
	return s
}

type DetailsResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s DetailsResponse) GoString() string {
	return s.String()
}

func (s *DetailsResponse) SetHeaders(v map[string]*string) *DetailsResponse {
	s.Headers = v
	return s
}

func (s *DetailsResponse) SetStatusCode(v int32) *DetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DetailsResponse) SetBody(v *DetailsResponseBody) *DetailsResponse {
	s.Body = v
	return s
}

type EditTaskRequest struct {
	// 外呼时间
	CallTimeList []*EditTaskRequestCallTimeList `json:"CallTimeList,omitempty" xml:"CallTimeList,omitempty" type:"Repeated"`
	// 回调地址
	//
	// example:
	//
	// 示例值示例值示例值
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// 当发送闪信配置为1时，闪信模板ID必填
	//
	// example:
	//
	// 60
	FlashSmsTemplateId *int64 `json:"FlashSmsTemplateId,omitempty" xml:"FlashSmsTemplateId,omitempty"`
	// 发送闪信配置,默认为0,0不发送闪信.1发送闪信
	//
	// example:
	//
	// 0
	FlashSmsType *int64 `json:"FlashSmsType,omitempty" xml:"FlashSmsType,omitempty"`
	// 并发数
	//
	// example:
	//
	// 83
	MaxConcurrency *int64 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// 任务名称
	//
	// example:
	//
	// 示例值示例值
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 播放间隔时长
	//
	// example:
	//
	// 29
	PlaySleepVal *int64 `json:"PlaySleepVal,omitempty" xml:"PlaySleepVal,omitempty"`
	// 录音播放次数
	//
	// example:
	//
	// 60
	PlayTimes *int64 `json:"PlayTimes,omitempty" xml:"PlayTimes,omitempty"`
	// 重呼配置
	//
	// example:
	//
	// 1
	RecallType *int64 `json:"RecallType,omitempty" xml:"RecallType,omitempty"`
	// 录音地址
	//
	// example:
	//
	// 示例值示例值示例值
	RecordPath *string `json:"RecordPath,omitempty" xml:"RecordPath,omitempty"`
	// 重呼次数
	//
	// example:
	//
	// 51
	RepeatCount *int64 `json:"RepeatCount,omitempty" xml:"RepeatCount,omitempty"`
	// 重呼间隔
	//
	// example:
	//
	// 91
	RepeatInterval *int64 `json:"RepeatInterval,omitempty" xml:"RepeatInterval,omitempty"`
	// 重呼条件
	RepeatReason []*int64 `json:"RepeatReason,omitempty" xml:"RepeatReason,omitempty" type:"Repeated"`
	// 重呼时间
	RepeatTimes          []*string `json:"RepeatTimes,omitempty" xml:"RepeatTimes,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 短信发送规则
	SendSmsPlan []*EditTaskRequestSendSmsPlan `json:"SendSmsPlan,omitempty" xml:"SendSmsPlan,omitempty" type:"Repeated"`
	// 任务状态
	//
	// example:
	//
	// 2
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 任务id
	//
	// This parameter is required.
	//
	// example:
	//
	// 29
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 话术模板ID
	//
	// example:
	//
	// 24
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 话术模板类型
	//
	// example:
	//
	// 1
	TemplateType *int64 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s EditTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s EditTaskRequest) GoString() string {
	return s.String()
}

func (s *EditTaskRequest) SetCallTimeList(v []*EditTaskRequestCallTimeList) *EditTaskRequest {
	s.CallTimeList = v
	return s
}

func (s *EditTaskRequest) SetCallbackUrl(v string) *EditTaskRequest {
	s.CallbackUrl = &v
	return s
}

func (s *EditTaskRequest) SetFlashSmsTemplateId(v int64) *EditTaskRequest {
	s.FlashSmsTemplateId = &v
	return s
}

func (s *EditTaskRequest) SetFlashSmsType(v int64) *EditTaskRequest {
	s.FlashSmsType = &v
	return s
}

func (s *EditTaskRequest) SetMaxConcurrency(v int64) *EditTaskRequest {
	s.MaxConcurrency = &v
	return s
}

func (s *EditTaskRequest) SetName(v string) *EditTaskRequest {
	s.Name = &v
	return s
}

func (s *EditTaskRequest) SetOwnerId(v int64) *EditTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *EditTaskRequest) SetPlaySleepVal(v int64) *EditTaskRequest {
	s.PlaySleepVal = &v
	return s
}

func (s *EditTaskRequest) SetPlayTimes(v int64) *EditTaskRequest {
	s.PlayTimes = &v
	return s
}

func (s *EditTaskRequest) SetRecallType(v int64) *EditTaskRequest {
	s.RecallType = &v
	return s
}

func (s *EditTaskRequest) SetRecordPath(v string) *EditTaskRequest {
	s.RecordPath = &v
	return s
}

func (s *EditTaskRequest) SetRepeatCount(v int64) *EditTaskRequest {
	s.RepeatCount = &v
	return s
}

func (s *EditTaskRequest) SetRepeatInterval(v int64) *EditTaskRequest {
	s.RepeatInterval = &v
	return s
}

func (s *EditTaskRequest) SetRepeatReason(v []*int64) *EditTaskRequest {
	s.RepeatReason = v
	return s
}

func (s *EditTaskRequest) SetRepeatTimes(v []*string) *EditTaskRequest {
	s.RepeatTimes = v
	return s
}

func (s *EditTaskRequest) SetResourceOwnerAccount(v string) *EditTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *EditTaskRequest) SetResourceOwnerId(v int64) *EditTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *EditTaskRequest) SetSendSmsPlan(v []*EditTaskRequestSendSmsPlan) *EditTaskRequest {
	s.SendSmsPlan = v
	return s
}

func (s *EditTaskRequest) SetStatus(v int64) *EditTaskRequest {
	s.Status = &v
	return s
}

func (s *EditTaskRequest) SetTaskId(v int64) *EditTaskRequest {
	s.TaskId = &v
	return s
}

func (s *EditTaskRequest) SetTemplateId(v int64) *EditTaskRequest {
	s.TemplateId = &v
	return s
}

func (s *EditTaskRequest) SetTemplateType(v int64) *EditTaskRequest {
	s.TemplateType = &v
	return s
}

type EditTaskRequestCallTimeList struct {
	CallTime []*string `json:"CallTime,omitempty" xml:"CallTime,omitempty" type:"Repeated"`
}

func (s EditTaskRequestCallTimeList) String() string {
	return tea.Prettify(s)
}

func (s EditTaskRequestCallTimeList) GoString() string {
	return s.String()
}

func (s *EditTaskRequestCallTimeList) SetCallTime(v []*string) *EditTaskRequestCallTimeList {
	s.CallTime = v
	return s
}

type EditTaskRequestSendSmsPlan struct {
	// 意向标签
	IntentTags []*string `json:"IntentTags,omitempty" xml:"IntentTags,omitempty" type:"Repeated"`
	// 短信模板ID
	//
	// example:
	//
	// 1
	SmsTemplateId *int64 `json:"SmsTemplateId,omitempty" xml:"SmsTemplateId,omitempty"`
}

func (s EditTaskRequestSendSmsPlan) String() string {
	return tea.Prettify(s)
}

func (s EditTaskRequestSendSmsPlan) GoString() string {
	return s.String()
}

func (s *EditTaskRequestSendSmsPlan) SetIntentTags(v []*string) *EditTaskRequestSendSmsPlan {
	s.IntentTags = v
	return s
}

func (s *EditTaskRequestSendSmsPlan) SetSmsTemplateId(v int64) *EditTaskRequestSendSmsPlan {
	s.SmsTemplateId = &v
	return s
}

type EditTaskShrinkRequest struct {
	// 外呼时间
	CallTimeListShrink *string `json:"CallTimeList,omitempty" xml:"CallTimeList,omitempty"`
	// 回调地址
	//
	// example:
	//
	// 示例值示例值示例值
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// 当发送闪信配置为1时，闪信模板ID必填
	//
	// example:
	//
	// 60
	FlashSmsTemplateId *int64 `json:"FlashSmsTemplateId,omitempty" xml:"FlashSmsTemplateId,omitempty"`
	// 发送闪信配置,默认为0,0不发送闪信.1发送闪信
	//
	// example:
	//
	// 0
	FlashSmsType *int64 `json:"FlashSmsType,omitempty" xml:"FlashSmsType,omitempty"`
	// 并发数
	//
	// example:
	//
	// 83
	MaxConcurrency *int64 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// 任务名称
	//
	// example:
	//
	// 示例值示例值
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 播放间隔时长
	//
	// example:
	//
	// 29
	PlaySleepVal *int64 `json:"PlaySleepVal,omitempty" xml:"PlaySleepVal,omitempty"`
	// 录音播放次数
	//
	// example:
	//
	// 60
	PlayTimes *int64 `json:"PlayTimes,omitempty" xml:"PlayTimes,omitempty"`
	// 重呼配置
	//
	// example:
	//
	// 1
	RecallType *int64 `json:"RecallType,omitempty" xml:"RecallType,omitempty"`
	// 录音地址
	//
	// example:
	//
	// 示例值示例值示例值
	RecordPath *string `json:"RecordPath,omitempty" xml:"RecordPath,omitempty"`
	// 重呼次数
	//
	// example:
	//
	// 51
	RepeatCount *int64 `json:"RepeatCount,omitempty" xml:"RepeatCount,omitempty"`
	// 重呼间隔
	//
	// example:
	//
	// 91
	RepeatInterval *int64 `json:"RepeatInterval,omitempty" xml:"RepeatInterval,omitempty"`
	// 重呼条件
	RepeatReasonShrink *string `json:"RepeatReason,omitempty" xml:"RepeatReason,omitempty"`
	// 重呼时间
	RepeatTimesShrink    *string `json:"RepeatTimes,omitempty" xml:"RepeatTimes,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 短信发送规则
	SendSmsPlanShrink *string `json:"SendSmsPlan,omitempty" xml:"SendSmsPlan,omitempty"`
	// 任务状态
	//
	// example:
	//
	// 2
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 任务id
	//
	// This parameter is required.
	//
	// example:
	//
	// 29
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 话术模板ID
	//
	// example:
	//
	// 24
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 话术模板类型
	//
	// example:
	//
	// 1
	TemplateType *int64 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s EditTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s EditTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *EditTaskShrinkRequest) SetCallTimeListShrink(v string) *EditTaskShrinkRequest {
	s.CallTimeListShrink = &v
	return s
}

func (s *EditTaskShrinkRequest) SetCallbackUrl(v string) *EditTaskShrinkRequest {
	s.CallbackUrl = &v
	return s
}

func (s *EditTaskShrinkRequest) SetFlashSmsTemplateId(v int64) *EditTaskShrinkRequest {
	s.FlashSmsTemplateId = &v
	return s
}

func (s *EditTaskShrinkRequest) SetFlashSmsType(v int64) *EditTaskShrinkRequest {
	s.FlashSmsType = &v
	return s
}

func (s *EditTaskShrinkRequest) SetMaxConcurrency(v int64) *EditTaskShrinkRequest {
	s.MaxConcurrency = &v
	return s
}

func (s *EditTaskShrinkRequest) SetName(v string) *EditTaskShrinkRequest {
	s.Name = &v
	return s
}

func (s *EditTaskShrinkRequest) SetOwnerId(v int64) *EditTaskShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *EditTaskShrinkRequest) SetPlaySleepVal(v int64) *EditTaskShrinkRequest {
	s.PlaySleepVal = &v
	return s
}

func (s *EditTaskShrinkRequest) SetPlayTimes(v int64) *EditTaskShrinkRequest {
	s.PlayTimes = &v
	return s
}

func (s *EditTaskShrinkRequest) SetRecallType(v int64) *EditTaskShrinkRequest {
	s.RecallType = &v
	return s
}

func (s *EditTaskShrinkRequest) SetRecordPath(v string) *EditTaskShrinkRequest {
	s.RecordPath = &v
	return s
}

func (s *EditTaskShrinkRequest) SetRepeatCount(v int64) *EditTaskShrinkRequest {
	s.RepeatCount = &v
	return s
}

func (s *EditTaskShrinkRequest) SetRepeatInterval(v int64) *EditTaskShrinkRequest {
	s.RepeatInterval = &v
	return s
}

func (s *EditTaskShrinkRequest) SetRepeatReasonShrink(v string) *EditTaskShrinkRequest {
	s.RepeatReasonShrink = &v
	return s
}

func (s *EditTaskShrinkRequest) SetRepeatTimesShrink(v string) *EditTaskShrinkRequest {
	s.RepeatTimesShrink = &v
	return s
}

func (s *EditTaskShrinkRequest) SetResourceOwnerAccount(v string) *EditTaskShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *EditTaskShrinkRequest) SetResourceOwnerId(v int64) *EditTaskShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *EditTaskShrinkRequest) SetSendSmsPlanShrink(v string) *EditTaskShrinkRequest {
	s.SendSmsPlanShrink = &v
	return s
}

func (s *EditTaskShrinkRequest) SetStatus(v int64) *EditTaskShrinkRequest {
	s.Status = &v
	return s
}

func (s *EditTaskShrinkRequest) SetTaskId(v int64) *EditTaskShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *EditTaskShrinkRequest) SetTemplateId(v int64) *EditTaskShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *EditTaskShrinkRequest) SetTemplateType(v int64) *EditTaskShrinkRequest {
	s.TemplateType = &v
	return s
}

type EditTaskResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *EditTaskResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 5453cc9b-02bc-4dbb-9527-f28a9100b912
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1686225227338
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s EditTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditTaskResponseBody) GoString() string {
	return s.String()
}

func (s *EditTaskResponseBody) SetAccessDeniedDetail(v string) *EditTaskResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *EditTaskResponseBody) SetCode(v int64) *EditTaskResponseBody {
	s.Code = &v
	return s
}

func (s *EditTaskResponseBody) SetMessage(v string) *EditTaskResponseBody {
	s.Message = &v
	return s
}

func (s *EditTaskResponseBody) SetModel(v *EditTaskResponseBodyModel) *EditTaskResponseBody {
	s.Model = v
	return s
}

func (s *EditTaskResponseBody) SetRequestId(v string) *EditTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *EditTaskResponseBody) SetSuccess(v bool) *EditTaskResponseBody {
	s.Success = &v
	return s
}

func (s *EditTaskResponseBody) SetTimestamp(v int64) *EditTaskResponseBody {
	s.Timestamp = &v
	return s
}

type EditTaskResponseBodyModel struct {
	// 任务ID
	//
	// example:
	//
	// 92
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s EditTaskResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s EditTaskResponseBodyModel) GoString() string {
	return s.String()
}

func (s *EditTaskResponseBodyModel) SetTaskId(v int64) *EditTaskResponseBodyModel {
	s.TaskId = &v
	return s
}

type EditTaskResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EditTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s EditTaskResponse) GoString() string {
	return s.String()
}

func (s *EditTaskResponse) SetHeaders(v map[string]*string) *EditTaskResponse {
	s.Headers = v
	return s
}

func (s *EditTaskResponse) SetStatusCode(v int32) *EditTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *EditTaskResponse) SetBody(v *EditTaskResponseBody) *EditTaskResponse {
	s.Body = v
	return s
}

type ImportNumberRequest struct {
	// This parameter is required.
	Customers []*ImportNumberRequestCustomers `json:"Customers,omitempty" xml:"Customers,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	FailReturn *int64 `json:"FailReturn,omitempty" xml:"FailReturn,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 92
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ImportNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportNumberRequest) GoString() string {
	return s.String()
}

func (s *ImportNumberRequest) SetCustomers(v []*ImportNumberRequestCustomers) *ImportNumberRequest {
	s.Customers = v
	return s
}

func (s *ImportNumberRequest) SetFailReturn(v int64) *ImportNumberRequest {
	s.FailReturn = &v
	return s
}

func (s *ImportNumberRequest) SetOutId(v string) *ImportNumberRequest {
	s.OutId = &v
	return s
}

func (s *ImportNumberRequest) SetOwnerId(v int64) *ImportNumberRequest {
	s.OwnerId = &v
	return s
}

func (s *ImportNumberRequest) SetResourceOwnerAccount(v string) *ImportNumberRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ImportNumberRequest) SetResourceOwnerId(v int64) *ImportNumberRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ImportNumberRequest) SetTaskId(v int64) *ImportNumberRequest {
	s.TaskId = &v
	return s
}

type ImportNumberRequestCustomers struct {
	// example:
	//
	// http://test.com
	ClientUrl *string `json:"ClientUrl,omitempty" xml:"ClientUrl,omitempty"`
	// example:
	//
	// 13541251222,18665214444
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// example:
	//
	// 示例值
	NumberMD5 *string `json:"NumberMD5,omitempty" xml:"NumberMD5,omitempty"`
	// example:
	//
	// {"testt":"123"}
	Properties map[string]interface{} `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// example:
	//
	// tag1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ImportNumberRequestCustomers) String() string {
	return tea.Prettify(s)
}

func (s ImportNumberRequestCustomers) GoString() string {
	return s.String()
}

func (s *ImportNumberRequestCustomers) SetClientUrl(v string) *ImportNumberRequestCustomers {
	s.ClientUrl = &v
	return s
}

func (s *ImportNumberRequestCustomers) SetNumber(v string) *ImportNumberRequestCustomers {
	s.Number = &v
	return s
}

func (s *ImportNumberRequestCustomers) SetNumberMD5(v string) *ImportNumberRequestCustomers {
	s.NumberMD5 = &v
	return s
}

func (s *ImportNumberRequestCustomers) SetProperties(v map[string]interface{}) *ImportNumberRequestCustomers {
	s.Properties = v
	return s
}

func (s *ImportNumberRequestCustomers) SetTag(v string) *ImportNumberRequestCustomers {
	s.Tag = &v
	return s
}

type ImportNumberShrinkRequest struct {
	// This parameter is required.
	CustomersShrink *string `json:"Customers,omitempty" xml:"Customers,omitempty"`
	// example:
	//
	// 0
	FailReturn *int64 `json:"FailReturn,omitempty" xml:"FailReturn,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 92
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ImportNumberShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportNumberShrinkRequest) GoString() string {
	return s.String()
}

func (s *ImportNumberShrinkRequest) SetCustomersShrink(v string) *ImportNumberShrinkRequest {
	s.CustomersShrink = &v
	return s
}

func (s *ImportNumberShrinkRequest) SetFailReturn(v int64) *ImportNumberShrinkRequest {
	s.FailReturn = &v
	return s
}

func (s *ImportNumberShrinkRequest) SetOutId(v string) *ImportNumberShrinkRequest {
	s.OutId = &v
	return s
}

func (s *ImportNumberShrinkRequest) SetOwnerId(v int64) *ImportNumberShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ImportNumberShrinkRequest) SetResourceOwnerAccount(v string) *ImportNumberShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ImportNumberShrinkRequest) SetResourceOwnerId(v int64) *ImportNumberShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ImportNumberShrinkRequest) SetTaskId(v int64) *ImportNumberShrinkRequest {
	s.TaskId = &v
	return s
}

type ImportNumberResponseBody struct {
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *ImportNumberResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 8EFC6D10-307B-1ECA-A8C6-7CBDF776AAD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1683440860035
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s ImportNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportNumberResponseBody) GoString() string {
	return s.String()
}

func (s *ImportNumberResponseBody) SetCode(v int64) *ImportNumberResponseBody {
	s.Code = &v
	return s
}

func (s *ImportNumberResponseBody) SetMessage(v string) *ImportNumberResponseBody {
	s.Message = &v
	return s
}

func (s *ImportNumberResponseBody) SetModel(v *ImportNumberResponseBodyModel) *ImportNumberResponseBody {
	s.Model = v
	return s
}

func (s *ImportNumberResponseBody) SetRequestId(v string) *ImportNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportNumberResponseBody) SetSuccess(v string) *ImportNumberResponseBody {
	s.Success = &v
	return s
}

func (s *ImportNumberResponseBody) SetTimestamp(v int64) *ImportNumberResponseBody {
	s.Timestamp = &v
	return s
}

type ImportNumberResponseBodyModel struct {
	// example:
	//
	// 54
	BatchId *int64 `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// example:
	//
	// 94
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 26
	ImportNum *int64 `json:"ImportNum,omitempty" xml:"ImportNum,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ImportNumberResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s ImportNumberResponseBodyModel) GoString() string {
	return s.String()
}

func (s *ImportNumberResponseBodyModel) SetBatchId(v int64) *ImportNumberResponseBodyModel {
	s.BatchId = &v
	return s
}

func (s *ImportNumberResponseBodyModel) SetCode(v int64) *ImportNumberResponseBodyModel {
	s.Code = &v
	return s
}

func (s *ImportNumberResponseBodyModel) SetData(v string) *ImportNumberResponseBodyModel {
	s.Data = &v
	return s
}

func (s *ImportNumberResponseBodyModel) SetImportNum(v int64) *ImportNumberResponseBodyModel {
	s.ImportNum = &v
	return s
}

func (s *ImportNumberResponseBodyModel) SetMessage(v string) *ImportNumberResponseBodyModel {
	s.Message = &v
	return s
}

type ImportNumberResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportNumberResponse) GoString() string {
	return s.String()
}

func (s *ImportNumberResponse) SetHeaders(v map[string]*string) *ImportNumberResponse {
	s.Headers = v
	return s
}

func (s *ImportNumberResponse) SetStatusCode(v int32) *ImportNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportNumberResponse) SetBody(v *ImportNumberResponseBody) *ImportNumberResponse {
	s.Body = v
	return s
}

type PageRequest struct {
	Numbers []*string `json:"Numbers,omitempty" xml:"Numbers,omitempty" type:"Repeated"`
	OwnerId *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 20
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s PageRequest) String() string {
	return tea.Prettify(s)
}

func (s PageRequest) GoString() string {
	return s.String()
}

func (s *PageRequest) SetNumbers(v []*string) *PageRequest {
	s.Numbers = v
	return s
}

func (s *PageRequest) SetOwnerId(v int64) *PageRequest {
	s.OwnerId = &v
	return s
}

func (s *PageRequest) SetPageNo(v int64) *PageRequest {
	s.PageNo = &v
	return s
}

func (s *PageRequest) SetPageSize(v int64) *PageRequest {
	s.PageSize = &v
	return s
}

func (s *PageRequest) SetResourceOwnerAccount(v string) *PageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PageRequest) SetResourceOwnerId(v int64) *PageRequest {
	s.ResourceOwnerId = &v
	return s
}

type PageShrinkRequest struct {
	NumbersShrink *string `json:"Numbers,omitempty" xml:"Numbers,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 20
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s PageShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s PageShrinkRequest) GoString() string {
	return s.String()
}

func (s *PageShrinkRequest) SetNumbersShrink(v string) *PageShrinkRequest {
	s.NumbersShrink = &v
	return s
}

func (s *PageShrinkRequest) SetOwnerId(v int64) *PageShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *PageShrinkRequest) SetPageNo(v int64) *PageShrinkRequest {
	s.PageNo = &v
	return s
}

func (s *PageShrinkRequest) SetPageSize(v int64) *PageShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *PageShrinkRequest) SetResourceOwnerAccount(v string) *PageShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PageShrinkRequest) SetResourceOwnerId(v int64) *PageShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

type PageResponseBody struct {
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *PageResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 8EFC6D10-307B-1ECA-A8C6-7CBDF776AAD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1683440860035
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s PageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PageResponseBody) GoString() string {
	return s.String()
}

func (s *PageResponseBody) SetCode(v int64) *PageResponseBody {
	s.Code = &v
	return s
}

func (s *PageResponseBody) SetMessage(v string) *PageResponseBody {
	s.Message = &v
	return s
}

func (s *PageResponseBody) SetModel(v *PageResponseBodyModel) *PageResponseBody {
	s.Model = v
	return s
}

func (s *PageResponseBody) SetRequestId(v string) *PageResponseBody {
	s.RequestId = &v
	return s
}

func (s *PageResponseBody) SetSuccess(v string) *PageResponseBody {
	s.Success = &v
	return s
}

func (s *PageResponseBody) SetTimestamp(v int64) *PageResponseBody {
	s.Timestamp = &v
	return s
}

type PageResponseBodyModel struct {
	List []*PageResponseBodyModelList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 97
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 5
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s PageResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s PageResponseBodyModel) GoString() string {
	return s.String()
}

func (s *PageResponseBodyModel) SetList(v []*PageResponseBodyModelList) *PageResponseBodyModel {
	s.List = v
	return s
}

func (s *PageResponseBodyModel) SetPageNo(v int64) *PageResponseBodyModel {
	s.PageNo = &v
	return s
}

func (s *PageResponseBodyModel) SetPageSize(v int64) *PageResponseBodyModel {
	s.PageSize = &v
	return s
}

func (s *PageResponseBodyModel) SetTotalCount(v int64) *PageResponseBodyModel {
	s.TotalCount = &v
	return s
}

func (s *PageResponseBodyModel) SetTotalPage(v int64) *PageResponseBodyModel {
	s.TotalPage = &v
	return s
}

type PageResponseBodyModelList struct {
	// 添加时间
	//
	// example:
	//
	// 2020-03-06 10:10:10
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 过期时间
	//
	// example:
	//
	// 1
	ExpirationTime *string `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// 手机号码
	//
	// example:
	//
	// 13314206082
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// 手机号MD5
	//
	// example:
	//
	// e10adc3949ba59abbe56e057f20f883e
	NumberMD5 *string `json:"NumberMD5,omitempty" xml:"NumberMD5,omitempty"`
	// 备注
	//
	// example:
	//
	// 示例值
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s PageResponseBodyModelList) String() string {
	return tea.Prettify(s)
}

func (s PageResponseBodyModelList) GoString() string {
	return s.String()
}

func (s *PageResponseBodyModelList) SetCreateTime(v string) *PageResponseBodyModelList {
	s.CreateTime = &v
	return s
}

func (s *PageResponseBodyModelList) SetExpirationTime(v string) *PageResponseBodyModelList {
	s.ExpirationTime = &v
	return s
}

func (s *PageResponseBodyModelList) SetNumber(v string) *PageResponseBodyModelList {
	s.Number = &v
	return s
}

func (s *PageResponseBodyModelList) SetNumberMD5(v string) *PageResponseBodyModelList {
	s.NumberMD5 = &v
	return s
}

func (s *PageResponseBodyModelList) SetRemark(v string) *PageResponseBodyModelList {
	s.Remark = &v
	return s
}

type PageResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PageResponseBody  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PageResponse) String() string {
	return tea.Prettify(s)
}

func (s PageResponse) GoString() string {
	return s.String()
}

func (s *PageResponse) SetHeaders(v map[string]*string) *PageResponse {
	s.Headers = v
	return s
}

func (s *PageResponse) SetStatusCode(v int32) *PageResponse {
	s.StatusCode = &v
	return s
}

func (s *PageResponse) SetBody(v *PageResponseBody) *PageResponse {
	s.Body = v
	return s
}

type SmsTemplateCreateRequest struct {
	// 短信内容
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	Content              *string `json:"Content,omitempty" xml:"Content,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 短信签名
	//
	// This parameter is required.
	//
	// example:
	//
	// ef2i29fsljf
	Sign *string `json:"Sign,omitempty" xml:"Sign,omitempty"`
	// 短信类型
	//
	// example:
	//
	// 73
	SmsType *int64 `json:"SmsType,omitempty" xml:"SmsType,omitempty"`
	// 模板名称
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// 模板类型
	//
	// example:
	//
	// 56
	TemplateType *int64 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s SmsTemplateCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s SmsTemplateCreateRequest) GoString() string {
	return s.String()
}

func (s *SmsTemplateCreateRequest) SetContent(v string) *SmsTemplateCreateRequest {
	s.Content = &v
	return s
}

func (s *SmsTemplateCreateRequest) SetOwnerId(v int64) *SmsTemplateCreateRequest {
	s.OwnerId = &v
	return s
}

func (s *SmsTemplateCreateRequest) SetResourceOwnerAccount(v string) *SmsTemplateCreateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SmsTemplateCreateRequest) SetResourceOwnerId(v int64) *SmsTemplateCreateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SmsTemplateCreateRequest) SetSign(v string) *SmsTemplateCreateRequest {
	s.Sign = &v
	return s
}

func (s *SmsTemplateCreateRequest) SetSmsType(v int64) *SmsTemplateCreateRequest {
	s.SmsType = &v
	return s
}

func (s *SmsTemplateCreateRequest) SetTemplateName(v string) *SmsTemplateCreateRequest {
	s.TemplateName = &v
	return s
}

func (s *SmsTemplateCreateRequest) SetTemplateType(v int64) *SmsTemplateCreateRequest {
	s.TemplateType = &v
	return s
}

type SmsTemplateCreateResponseBody struct {
	// example:
	//
	// 23
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值示例值
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// 8EFC6D10-307B-1ECA-A8C6-7CBDF776AAD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1683440860035
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s SmsTemplateCreateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SmsTemplateCreateResponseBody) GoString() string {
	return s.String()
}

func (s *SmsTemplateCreateResponseBody) SetCode(v int64) *SmsTemplateCreateResponseBody {
	s.Code = &v
	return s
}

func (s *SmsTemplateCreateResponseBody) SetMessage(v string) *SmsTemplateCreateResponseBody {
	s.Message = &v
	return s
}

func (s *SmsTemplateCreateResponseBody) SetModel(v string) *SmsTemplateCreateResponseBody {
	s.Model = &v
	return s
}

func (s *SmsTemplateCreateResponseBody) SetRequestId(v string) *SmsTemplateCreateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SmsTemplateCreateResponseBody) SetSuccess(v bool) *SmsTemplateCreateResponseBody {
	s.Success = &v
	return s
}

func (s *SmsTemplateCreateResponseBody) SetTimestamp(v int64) *SmsTemplateCreateResponseBody {
	s.Timestamp = &v
	return s
}

type SmsTemplateCreateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SmsTemplateCreateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SmsTemplateCreateResponse) String() string {
	return tea.Prettify(s)
}

func (s SmsTemplateCreateResponse) GoString() string {
	return s.String()
}

func (s *SmsTemplateCreateResponse) SetHeaders(v map[string]*string) *SmsTemplateCreateResponse {
	s.Headers = v
	return s
}

func (s *SmsTemplateCreateResponse) SetStatusCode(v int32) *SmsTemplateCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *SmsTemplateCreateResponse) SetBody(v *SmsTemplateCreateResponseBody) *SmsTemplateCreateResponse {
	s.Body = v
	return s
}

type SmsTemplatePageListRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 页码
	//
	// example:
	//
	// 24
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// 每页条数
	//
	// example:
	//
	// 97
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 短信签名
	//
	// example:
	//
	// 114ah23m
	Sign *string `json:"Sign,omitempty" xml:"Sign,omitempty"`
	// 短信类型
	//
	// example:
	//
	// 42
	SmsType *int64 `json:"SmsType,omitempty" xml:"SmsType,omitempty"`
	// 模板状态
	//
	// example:
	//
	// 92
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 模板ID
	//
	// example:
	//
	// 83
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 模板类型
	//
	// example:
	//
	// 19
	TemplateType *int64 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s SmsTemplatePageListRequest) String() string {
	return tea.Prettify(s)
}

func (s SmsTemplatePageListRequest) GoString() string {
	return s.String()
}

func (s *SmsTemplatePageListRequest) SetOwnerId(v int64) *SmsTemplatePageListRequest {
	s.OwnerId = &v
	return s
}

func (s *SmsTemplatePageListRequest) SetPageNo(v int64) *SmsTemplatePageListRequest {
	s.PageNo = &v
	return s
}

func (s *SmsTemplatePageListRequest) SetPageSize(v int64) *SmsTemplatePageListRequest {
	s.PageSize = &v
	return s
}

func (s *SmsTemplatePageListRequest) SetResourceOwnerAccount(v string) *SmsTemplatePageListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SmsTemplatePageListRequest) SetResourceOwnerId(v int64) *SmsTemplatePageListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SmsTemplatePageListRequest) SetSign(v string) *SmsTemplatePageListRequest {
	s.Sign = &v
	return s
}

func (s *SmsTemplatePageListRequest) SetSmsType(v int64) *SmsTemplatePageListRequest {
	s.SmsType = &v
	return s
}

func (s *SmsTemplatePageListRequest) SetStatus(v int64) *SmsTemplatePageListRequest {
	s.Status = &v
	return s
}

func (s *SmsTemplatePageListRequest) SetTemplateId(v int64) *SmsTemplatePageListRequest {
	s.TemplateId = &v
	return s
}

func (s *SmsTemplatePageListRequest) SetTemplateType(v int64) *SmsTemplatePageListRequest {
	s.TemplateType = &v
	return s
}

type SmsTemplatePageListResponseBody struct {
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值
	Message *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *SmsTemplatePageListResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 8EFC6D10-307B-1ECA-A8C6-7CBDF776AAD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1683440860035
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s SmsTemplatePageListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SmsTemplatePageListResponseBody) GoString() string {
	return s.String()
}

func (s *SmsTemplatePageListResponseBody) SetCode(v int64) *SmsTemplatePageListResponseBody {
	s.Code = &v
	return s
}

func (s *SmsTemplatePageListResponseBody) SetMessage(v string) *SmsTemplatePageListResponseBody {
	s.Message = &v
	return s
}

func (s *SmsTemplatePageListResponseBody) SetModel(v *SmsTemplatePageListResponseBodyModel) *SmsTemplatePageListResponseBody {
	s.Model = v
	return s
}

func (s *SmsTemplatePageListResponseBody) SetRequestId(v string) *SmsTemplatePageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *SmsTemplatePageListResponseBody) SetSuccess(v bool) *SmsTemplatePageListResponseBody {
	s.Success = &v
	return s
}

func (s *SmsTemplatePageListResponseBody) SetTimestamp(v int64) *SmsTemplatePageListResponseBody {
	s.Timestamp = &v
	return s
}

type SmsTemplatePageListResponseBodyModel struct {
	List []*SmsTemplatePageListResponseBodyModelList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 53
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 42
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 31
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s SmsTemplatePageListResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s SmsTemplatePageListResponseBodyModel) GoString() string {
	return s.String()
}

func (s *SmsTemplatePageListResponseBodyModel) SetList(v []*SmsTemplatePageListResponseBodyModelList) *SmsTemplatePageListResponseBodyModel {
	s.List = v
	return s
}

func (s *SmsTemplatePageListResponseBodyModel) SetPageNo(v int64) *SmsTemplatePageListResponseBodyModel {
	s.PageNo = &v
	return s
}

func (s *SmsTemplatePageListResponseBodyModel) SetPageSize(v int64) *SmsTemplatePageListResponseBodyModel {
	s.PageSize = &v
	return s
}

func (s *SmsTemplatePageListResponseBodyModel) SetTotalCount(v int64) *SmsTemplatePageListResponseBodyModel {
	s.TotalCount = &v
	return s
}

func (s *SmsTemplatePageListResponseBodyModel) SetTotalPage(v int64) *SmsTemplatePageListResponseBodyModel {
	s.TotalPage = &v
	return s
}

type SmsTemplatePageListResponseBodyModelList struct {
	// 短信内容
	//
	// example:
	//
	// 示例值示例值示例值
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 创建时间
	//
	// example:
	//
	// 2021-09-26 11:34:59
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 模板所需参数
	//
	// example:
	//
	// 示例值示例值
	Properties *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// 智能短链ID
	//
	// example:
	//
	// 46
	ShortUrlTaskId *int64 `json:"ShortUrlTaskId,omitempty" xml:"ShortUrlTaskId,omitempty"`
	// 短信签名
	//
	// example:
	//
	// a234n
	Sign *string `json:"Sign,omitempty" xml:"Sign,omitempty"`
	// 短信类型
	//
	// example:
	//
	// 示例值示例值
	SmsType *string `json:"SmsType,omitempty" xml:"SmsType,omitempty"`
	// 模板状态
	//
	// example:
	//
	// 18
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 模板ID
	//
	// example:
	//
	// 67
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 模板名称
	//
	// example:
	//
	// 示例值示例值示例值
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// 模板类型
	//
	// example:
	//
	// 56
	TemplateType *int64 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s SmsTemplatePageListResponseBodyModelList) String() string {
	return tea.Prettify(s)
}

func (s SmsTemplatePageListResponseBodyModelList) GoString() string {
	return s.String()
}

func (s *SmsTemplatePageListResponseBodyModelList) SetContent(v string) *SmsTemplatePageListResponseBodyModelList {
	s.Content = &v
	return s
}

func (s *SmsTemplatePageListResponseBodyModelList) SetCreateTime(v string) *SmsTemplatePageListResponseBodyModelList {
	s.CreateTime = &v
	return s
}

func (s *SmsTemplatePageListResponseBodyModelList) SetProperties(v string) *SmsTemplatePageListResponseBodyModelList {
	s.Properties = &v
	return s
}

func (s *SmsTemplatePageListResponseBodyModelList) SetShortUrlTaskId(v int64) *SmsTemplatePageListResponseBodyModelList {
	s.ShortUrlTaskId = &v
	return s
}

func (s *SmsTemplatePageListResponseBodyModelList) SetSign(v string) *SmsTemplatePageListResponseBodyModelList {
	s.Sign = &v
	return s
}

func (s *SmsTemplatePageListResponseBodyModelList) SetSmsType(v string) *SmsTemplatePageListResponseBodyModelList {
	s.SmsType = &v
	return s
}

func (s *SmsTemplatePageListResponseBodyModelList) SetStatus(v int64) *SmsTemplatePageListResponseBodyModelList {
	s.Status = &v
	return s
}

func (s *SmsTemplatePageListResponseBodyModelList) SetTemplateId(v int64) *SmsTemplatePageListResponseBodyModelList {
	s.TemplateId = &v
	return s
}

func (s *SmsTemplatePageListResponseBodyModelList) SetTemplateName(v string) *SmsTemplatePageListResponseBodyModelList {
	s.TemplateName = &v
	return s
}

func (s *SmsTemplatePageListResponseBodyModelList) SetTemplateType(v int64) *SmsTemplatePageListResponseBodyModelList {
	s.TemplateType = &v
	return s
}

type SmsTemplatePageListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SmsTemplatePageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SmsTemplatePageListResponse) String() string {
	return tea.Prettify(s)
}

func (s SmsTemplatePageListResponse) GoString() string {
	return s.String()
}

func (s *SmsTemplatePageListResponse) SetHeaders(v map[string]*string) *SmsTemplatePageListResponse {
	s.Headers = v
	return s
}

func (s *SmsTemplatePageListResponse) SetStatusCode(v int32) *SmsTemplatePageListResponse {
	s.StatusCode = &v
	return s
}

func (s *SmsTemplatePageListResponse) SetBody(v *SmsTemplatePageListResponseBody) *SmsTemplatePageListResponse {
	s.Body = v
	return s
}

type TaskCallChatsRequest struct {
	// 坐席ID
	//
	// example:
	//
	// 72
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// 坐席标签
	//
	// example:
	//
	// AA
	AgentTag *string `json:"AgentTag,omitempty" xml:"AgentTag,omitempty"`
	// 外呼ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 9b2eb6b8-7a27-4357-b5ec-104450086e24
	CallId               *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 任务ID
	//
	// example:
	//
	// 26
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s TaskCallChatsRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskCallChatsRequest) GoString() string {
	return s.String()
}

func (s *TaskCallChatsRequest) SetAgentId(v int64) *TaskCallChatsRequest {
	s.AgentId = &v
	return s
}

func (s *TaskCallChatsRequest) SetAgentTag(v string) *TaskCallChatsRequest {
	s.AgentTag = &v
	return s
}

func (s *TaskCallChatsRequest) SetCallId(v string) *TaskCallChatsRequest {
	s.CallId = &v
	return s
}

func (s *TaskCallChatsRequest) SetOwnerId(v int64) *TaskCallChatsRequest {
	s.OwnerId = &v
	return s
}

func (s *TaskCallChatsRequest) SetResourceOwnerAccount(v string) *TaskCallChatsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TaskCallChatsRequest) SetResourceOwnerId(v int64) *TaskCallChatsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TaskCallChatsRequest) SetTaskId(v int64) *TaskCallChatsRequest {
	s.TaskId = &v
	return s
}

type TaskCallChatsResponseBody struct {
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   []*TaskCallChatsResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Repeated"`
	// example:
	//
	// 8EFC6D10-307B-1ECA-A8C6-7CBDF776AAD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1683440860035
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s TaskCallChatsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TaskCallChatsResponseBody) GoString() string {
	return s.String()
}

func (s *TaskCallChatsResponseBody) SetCode(v int64) *TaskCallChatsResponseBody {
	s.Code = &v
	return s
}

func (s *TaskCallChatsResponseBody) SetMessage(v string) *TaskCallChatsResponseBody {
	s.Message = &v
	return s
}

func (s *TaskCallChatsResponseBody) SetModel(v []*TaskCallChatsResponseBodyModel) *TaskCallChatsResponseBody {
	s.Model = v
	return s
}

func (s *TaskCallChatsResponseBody) SetRequestId(v string) *TaskCallChatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *TaskCallChatsResponseBody) SetSuccess(v bool) *TaskCallChatsResponseBody {
	s.Success = &v
	return s
}

func (s *TaskCallChatsResponseBody) SetTimestamp(v int64) *TaskCallChatsResponseBody {
	s.Timestamp = &v
	return s
}

type TaskCallChatsResponseBodyModel struct {
	// 说话内容
	//
	// example:
	//
	// 示例值示例值
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 说话时间
	//
	// example:
	//
	// 2022-01-13 14:56:46.604
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 说话号码
	//
	// example:
	//
	// 138*****265
	FromNumber *string `json:"FromNumber,omitempty" xml:"FromNumber,omitempty"`
}

func (s TaskCallChatsResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s TaskCallChatsResponseBodyModel) GoString() string {
	return s.String()
}

func (s *TaskCallChatsResponseBodyModel) SetContent(v string) *TaskCallChatsResponseBodyModel {
	s.Content = &v
	return s
}

func (s *TaskCallChatsResponseBodyModel) SetCreateTime(v string) *TaskCallChatsResponseBodyModel {
	s.CreateTime = &v
	return s
}

func (s *TaskCallChatsResponseBodyModel) SetFromNumber(v string) *TaskCallChatsResponseBodyModel {
	s.FromNumber = &v
	return s
}

type TaskCallChatsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TaskCallChatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TaskCallChatsResponse) String() string {
	return tea.Prettify(s)
}

func (s TaskCallChatsResponse) GoString() string {
	return s.String()
}

func (s *TaskCallChatsResponse) SetHeaders(v map[string]*string) *TaskCallChatsResponse {
	s.Headers = v
	return s
}

func (s *TaskCallChatsResponse) SetStatusCode(v int32) *TaskCallChatsResponse {
	s.StatusCode = &v
	return s
}

func (s *TaskCallChatsResponse) SetBody(v *TaskCallChatsResponseBody) *TaskCallChatsResponse {
	s.Body = v
	return s
}

type TaskCallInfoRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s TaskCallInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskCallInfoRequest) GoString() string {
	return s.String()
}

func (s *TaskCallInfoRequest) SetOwnerId(v int64) *TaskCallInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *TaskCallInfoRequest) SetResourceOwnerAccount(v string) *TaskCallInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TaskCallInfoRequest) SetResourceOwnerId(v int64) *TaskCallInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TaskCallInfoRequest) SetTaskId(v int64) *TaskCallInfoRequest {
	s.TaskId = &v
	return s
}

type TaskCallInfoResponseBody struct {
	// example:
	//
	// 15
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *TaskCallInfoResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 示例值示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 62
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s TaskCallInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TaskCallInfoResponseBody) GoString() string {
	return s.String()
}

func (s *TaskCallInfoResponseBody) SetCode(v int64) *TaskCallInfoResponseBody {
	s.Code = &v
	return s
}

func (s *TaskCallInfoResponseBody) SetMessage(v string) *TaskCallInfoResponseBody {
	s.Message = &v
	return s
}

func (s *TaskCallInfoResponseBody) SetModel(v *TaskCallInfoResponseBodyModel) *TaskCallInfoResponseBody {
	s.Model = v
	return s
}

func (s *TaskCallInfoResponseBody) SetRequestId(v string) *TaskCallInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *TaskCallInfoResponseBody) SetSuccess(v bool) *TaskCallInfoResponseBody {
	s.Success = &v
	return s
}

func (s *TaskCallInfoResponseBody) SetTimestamp(v int64) *TaskCallInfoResponseBody {
	s.Timestamp = &v
	return s
}

type TaskCallInfoResponseBodyModel struct {
	// example:
	//
	// 75
	FinishTotal *int64 `json:"FinishTotal,omitempty" xml:"FinishTotal,omitempty"`
	// example:
	//
	// 59
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
	// example:
	//
	// 3
	UnCallTotal *int64 `json:"UnCallTotal,omitempty" xml:"UnCallTotal,omitempty"`
}

func (s TaskCallInfoResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s TaskCallInfoResponseBodyModel) GoString() string {
	return s.String()
}

func (s *TaskCallInfoResponseBodyModel) SetFinishTotal(v int64) *TaskCallInfoResponseBodyModel {
	s.FinishTotal = &v
	return s
}

func (s *TaskCallInfoResponseBodyModel) SetTotal(v int64) *TaskCallInfoResponseBodyModel {
	s.Total = &v
	return s
}

func (s *TaskCallInfoResponseBodyModel) SetUnCallTotal(v int64) *TaskCallInfoResponseBodyModel {
	s.UnCallTotal = &v
	return s
}

type TaskCallInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TaskCallInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TaskCallInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s TaskCallInfoResponse) GoString() string {
	return s.String()
}

func (s *TaskCallInfoResponse) SetHeaders(v map[string]*string) *TaskCallInfoResponse {
	s.Headers = v
	return s
}

func (s *TaskCallInfoResponse) SetStatusCode(v int32) *TaskCallInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *TaskCallInfoResponse) SetBody(v *TaskCallInfoResponseBody) *TaskCallInfoResponse {
	s.Body = v
	return s
}

type TaskCallListRequest struct {
	// 导入号码时返回的批次号
	//
	// example:
	//
	// 1
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// 开始外呼时间
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-01-25 00:00:00
	CallDate *string `json:"CallDate,omitempty" xml:"CallDate,omitempty"`
	// 结束外呼时间
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-01-25 00:00:00
	EndCallDate *string `json:"EndCallDate,omitempty" xml:"EndCallDate,omitempty"`
	// 意向标签
	IntentTags []*string `json:"IntentTags,omitempty" xml:"IntentTags,omitempty" type:"Repeated"`
	// 号码列表
	Numbers []*string `json:"Numbers,omitempty" xml:"Numbers,omitempty" type:"Repeated"`
	OwnerId *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 页数
	//
	// This parameter is required.
	//
	// example:
	//
	// 39
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// 每页外呼记录数,正整数，默认10000
	//
	// example:
	//
	// 97
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 任务ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 93
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s TaskCallListRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskCallListRequest) GoString() string {
	return s.String()
}

func (s *TaskCallListRequest) SetBatchId(v string) *TaskCallListRequest {
	s.BatchId = &v
	return s
}

func (s *TaskCallListRequest) SetCallDate(v string) *TaskCallListRequest {
	s.CallDate = &v
	return s
}

func (s *TaskCallListRequest) SetEndCallDate(v string) *TaskCallListRequest {
	s.EndCallDate = &v
	return s
}

func (s *TaskCallListRequest) SetIntentTags(v []*string) *TaskCallListRequest {
	s.IntentTags = v
	return s
}

func (s *TaskCallListRequest) SetNumbers(v []*string) *TaskCallListRequest {
	s.Numbers = v
	return s
}

func (s *TaskCallListRequest) SetOwnerId(v int64) *TaskCallListRequest {
	s.OwnerId = &v
	return s
}

func (s *TaskCallListRequest) SetPage(v int64) *TaskCallListRequest {
	s.Page = &v
	return s
}

func (s *TaskCallListRequest) SetPageSize(v int64) *TaskCallListRequest {
	s.PageSize = &v
	return s
}

func (s *TaskCallListRequest) SetResourceOwnerAccount(v string) *TaskCallListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TaskCallListRequest) SetResourceOwnerId(v int64) *TaskCallListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TaskCallListRequest) SetTaskId(v int64) *TaskCallListRequest {
	s.TaskId = &v
	return s
}

type TaskCallListShrinkRequest struct {
	// 导入号码时返回的批次号
	//
	// example:
	//
	// 1
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// 开始外呼时间
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-01-25 00:00:00
	CallDate *string `json:"CallDate,omitempty" xml:"CallDate,omitempty"`
	// 结束外呼时间
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-01-25 00:00:00
	EndCallDate *string `json:"EndCallDate,omitempty" xml:"EndCallDate,omitempty"`
	// 意向标签
	IntentTagsShrink *string `json:"IntentTags,omitempty" xml:"IntentTags,omitempty"`
	// 号码列表
	NumbersShrink *string `json:"Numbers,omitempty" xml:"Numbers,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 页数
	//
	// This parameter is required.
	//
	// example:
	//
	// 39
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// 每页外呼记录数,正整数，默认10000
	//
	// example:
	//
	// 97
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 任务ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 93
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s TaskCallListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskCallListShrinkRequest) GoString() string {
	return s.String()
}

func (s *TaskCallListShrinkRequest) SetBatchId(v string) *TaskCallListShrinkRequest {
	s.BatchId = &v
	return s
}

func (s *TaskCallListShrinkRequest) SetCallDate(v string) *TaskCallListShrinkRequest {
	s.CallDate = &v
	return s
}

func (s *TaskCallListShrinkRequest) SetEndCallDate(v string) *TaskCallListShrinkRequest {
	s.EndCallDate = &v
	return s
}

func (s *TaskCallListShrinkRequest) SetIntentTagsShrink(v string) *TaskCallListShrinkRequest {
	s.IntentTagsShrink = &v
	return s
}

func (s *TaskCallListShrinkRequest) SetNumbersShrink(v string) *TaskCallListShrinkRequest {
	s.NumbersShrink = &v
	return s
}

func (s *TaskCallListShrinkRequest) SetOwnerId(v int64) *TaskCallListShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *TaskCallListShrinkRequest) SetPage(v int64) *TaskCallListShrinkRequest {
	s.Page = &v
	return s
}

func (s *TaskCallListShrinkRequest) SetPageSize(v int64) *TaskCallListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *TaskCallListShrinkRequest) SetResourceOwnerAccount(v string) *TaskCallListShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TaskCallListShrinkRequest) SetResourceOwnerId(v int64) *TaskCallListShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TaskCallListShrinkRequest) SetTaskId(v int64) *TaskCallListShrinkRequest {
	s.TaskId = &v
	return s
}

type TaskCallListResponseBody struct {
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *TaskCallListResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 8EFC6D10-307B-1ECA-A8C6-7CBDF776AAD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1683440860035
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s TaskCallListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TaskCallListResponseBody) GoString() string {
	return s.String()
}

func (s *TaskCallListResponseBody) SetCode(v int64) *TaskCallListResponseBody {
	s.Code = &v
	return s
}

func (s *TaskCallListResponseBody) SetMessage(v string) *TaskCallListResponseBody {
	s.Message = &v
	return s
}

func (s *TaskCallListResponseBody) SetModel(v *TaskCallListResponseBodyModel) *TaskCallListResponseBody {
	s.Model = v
	return s
}

func (s *TaskCallListResponseBody) SetRequestId(v string) *TaskCallListResponseBody {
	s.RequestId = &v
	return s
}

func (s *TaskCallListResponseBody) SetSuccess(v string) *TaskCallListResponseBody {
	s.Success = &v
	return s
}

func (s *TaskCallListResponseBody) SetTimestamp(v int64) *TaskCallListResponseBody {
	s.Timestamp = &v
	return s
}

type TaskCallListResponseBodyModel struct {
	List []*TaskCallListResponseBodyModelList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 62
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 95
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 80
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 39
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s TaskCallListResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s TaskCallListResponseBodyModel) GoString() string {
	return s.String()
}

func (s *TaskCallListResponseBodyModel) SetList(v []*TaskCallListResponseBodyModelList) *TaskCallListResponseBodyModel {
	s.List = v
	return s
}

func (s *TaskCallListResponseBodyModel) SetPageNo(v int64) *TaskCallListResponseBodyModel {
	s.PageNo = &v
	return s
}

func (s *TaskCallListResponseBodyModel) SetPageSize(v int64) *TaskCallListResponseBodyModel {
	s.PageSize = &v
	return s
}

func (s *TaskCallListResponseBodyModel) SetTotalCount(v int64) *TaskCallListResponseBodyModel {
	s.TotalCount = &v
	return s
}

func (s *TaskCallListResponseBodyModel) SetTotalPage(v int64) *TaskCallListResponseBodyModel {
	s.TotalPage = &v
	return s
}

type TaskCallListResponseBodyModelList struct {
	// 加微
	//
	// example:
	//
	// 0
	AddWx *int64 `json:"AddWx,omitempty" xml:"AddWx,omitempty"`
	// 加微进度
	//
	// example:
	//
	// 示例值示例值
	AddWxStatus *string `json:"AddWxStatus,omitempty" xml:"AddWxStatus,omitempty"`
	// 坐席分机
	//
	// example:
	//
	// 112
	AgentExtension *string `json:"AgentExtension,omitempty" xml:"AgentExtension,omitempty"`
	// 坐席ID
	//
	// example:
	//
	// 87
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// 人工通话时长
	//
	// example:
	//
	// 98
	AgentSpeakingDuration *int64 `json:"AgentSpeakingDuration,omitempty" xml:"AgentSpeakingDuration,omitempty"`
	// 人工通话时长
	//
	// example:
	//
	// 示例值示例值
	AgentSpeakingTime *string `json:"AgentSpeakingTime,omitempty" xml:"AgentSpeakingTime,omitempty"`
	// 坐席标签
	//
	// example:
	//
	// A
	AgentTag *string `json:"AgentTag,omitempty" xml:"AgentTag,omitempty"`
	// 是否接通重呼
	//
	// example:
	//
	// 24
	AnswerRecall *int64 `json:"AnswerRecall,omitempty" xml:"AnswerRecall,omitempty"`
	// 接通时间
	//
	// example:
	//
	// 2022-01-26 18:58:25
	AnswerTime *string `json:"AnswerTime,omitempty" xml:"AnswerTime,omitempty"`
	// 批次ID
	//
	// example:
	//
	// 1
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// 开始通话时间
	//
	// example:
	//
	// 2022-01-26 18:58:25
	CallBeginTime *string `json:"CallBeginTime,omitempty" xml:"CallBeginTime,omitempty"`
	// 外呼ID
	//
	// example:
	//
	// 9197ed9e-ceda-42a5-b683-823b23ef208e
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// 呼叫次数
	//
	// example:
	//
	// 1
	CallTimes *string `json:"CallTimes,omitempty" xml:"CallTimes,omitempty"`
	// 外呼类型
	//
	// example:
	//
	// 1001
	CallType *int64 `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// 对话录音
	//
	// example:
	//
	// 示例值示例值
	ChatRecord *string `json:"ChatRecord,omitempty" xml:"ChatRecord,omitempty"`
	// 外呼网关
	//
	// example:
	//
	// 123
	Gateway *string `json:"Gateway,omitempty" xml:"Gateway,omitempty"`
	// 挂断时间
	//
	// example:
	//
	// 2022-01-26 18:58:25
	HangupTime *string `json:"HangupTime,omitempty" xml:"HangupTime,omitempty"`
	// 挂机方式
	//
	// example:
	//
	// 1
	HangupType *int64 `json:"HangupType,omitempty" xml:"HangupType,omitempty"`
	// 导入时间
	//
	// example:
	//
	// 2022-01-26 18:58:25
	ImportTime *string `json:"ImportTime,omitempty" xml:"ImportTime,omitempty"`
	// 个性标签
	//
	// example:
	//
	// A
	IndividualTag *string `json:"IndividualTag,omitempty" xml:"IndividualTag,omitempty"`
	// 意向说明
	//
	// example:
	//
	// 示例值示例值示例值
	IntentDescription *string `json:"IntentDescription,omitempty" xml:"IntentDescription,omitempty"`
	// 意向标签
	//
	// example:
	//
	// “C”
	IntentTag *string `json:"IntentTag,omitempty" xml:"IntentTag,omitempty"`
	// 拦截原因
	//
	// example:
	//
	// 示例值
	InterceptReason *string `json:"InterceptReason,omitempty" xml:"InterceptReason,omitempty"`
	// 回复关键词
	//
	// example:
	//
	// 示例值示例值
	Keywords *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// 外呼号码
	//
	// example:
	//
	// 138*****123
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// 外呼号码MD5
	//
	// example:
	//
	// 75916b635568954583783d
	NumberMD5 *string `json:"NumberMD5,omitempty" xml:"NumberMD5,omitempty"`
	// 参数
	//
	// example:
	//
	// 示例值示例值
	Properties *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// 备注
	//
	// example:
	//
	// 示例值示例值
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 振铃时长
	//
	// example:
	//
	// 66
	RingTime *int64 `json:"RingTime,omitempty" xml:"RingTime,omitempty"`
	// 挂机短信
	//
	// example:
	//
	// 示例值
	Sms *string `json:"Sms,omitempty" xml:"Sms,omitempty"`
	// AI通话时长
	//
	// example:
	//
	// 45
	SpeakingDuration *int64 `json:"SpeakingDuration,omitempty" xml:"SpeakingDuration,omitempty"`
	// AI通话时长
	//
	// example:
	//
	// 示例值示例值示例值
	SpeakingTime *string `json:"SpeakingTime,omitempty" xml:"SpeakingTime,omitempty"`
	// 对话轮次
	//
	// example:
	//
	// 0
	SpeakingTurns *string `json:"SpeakingTurns,omitempty" xml:"SpeakingTurns,omitempty"`
	// 外呼状态
	//
	// example:
	//
	// 示例值示例值示例值
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 外呼状态编码
	//
	// example:
	//
	// 1
	StatusCode *int64 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// 外呼状态描述
	//
	// example:
	//
	// 示例值示例值
	StatusDescription *string `json:"StatusDescription,omitempty" xml:"StatusDescription,omitempty"`
	// 用户自定义标签
	//
	// example:
	//
	// 示例值
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// 外呼任务ID
	//
	// example:
	//
	// 70
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// AI话术ID
	//
	// example:
	//
	// 66
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 转人工状态
	//
	// example:
	//
	// 示例值
	TransferStatus *string `json:"TransferStatus,omitempty" xml:"TransferStatus,omitempty"`
	// 转人工状态编码
	//
	// example:
	//
	// 1
	TransferStatusCode *int64 `json:"TransferStatusCode,omitempty" xml:"TransferStatusCode,omitempty"`
}

func (s TaskCallListResponseBodyModelList) String() string {
	return tea.Prettify(s)
}

func (s TaskCallListResponseBodyModelList) GoString() string {
	return s.String()
}

func (s *TaskCallListResponseBodyModelList) SetAddWx(v int64) *TaskCallListResponseBodyModelList {
	s.AddWx = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetAddWxStatus(v string) *TaskCallListResponseBodyModelList {
	s.AddWxStatus = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetAgentExtension(v string) *TaskCallListResponseBodyModelList {
	s.AgentExtension = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetAgentId(v int64) *TaskCallListResponseBodyModelList {
	s.AgentId = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetAgentSpeakingDuration(v int64) *TaskCallListResponseBodyModelList {
	s.AgentSpeakingDuration = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetAgentSpeakingTime(v string) *TaskCallListResponseBodyModelList {
	s.AgentSpeakingTime = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetAgentTag(v string) *TaskCallListResponseBodyModelList {
	s.AgentTag = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetAnswerRecall(v int64) *TaskCallListResponseBodyModelList {
	s.AnswerRecall = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetAnswerTime(v string) *TaskCallListResponseBodyModelList {
	s.AnswerTime = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetBatchId(v string) *TaskCallListResponseBodyModelList {
	s.BatchId = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetCallBeginTime(v string) *TaskCallListResponseBodyModelList {
	s.CallBeginTime = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetCallId(v string) *TaskCallListResponseBodyModelList {
	s.CallId = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetCallTimes(v string) *TaskCallListResponseBodyModelList {
	s.CallTimes = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetCallType(v int64) *TaskCallListResponseBodyModelList {
	s.CallType = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetChatRecord(v string) *TaskCallListResponseBodyModelList {
	s.ChatRecord = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetGateway(v string) *TaskCallListResponseBodyModelList {
	s.Gateway = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetHangupTime(v string) *TaskCallListResponseBodyModelList {
	s.HangupTime = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetHangupType(v int64) *TaskCallListResponseBodyModelList {
	s.HangupType = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetImportTime(v string) *TaskCallListResponseBodyModelList {
	s.ImportTime = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetIndividualTag(v string) *TaskCallListResponseBodyModelList {
	s.IndividualTag = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetIntentDescription(v string) *TaskCallListResponseBodyModelList {
	s.IntentDescription = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetIntentTag(v string) *TaskCallListResponseBodyModelList {
	s.IntentTag = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetInterceptReason(v string) *TaskCallListResponseBodyModelList {
	s.InterceptReason = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetKeywords(v string) *TaskCallListResponseBodyModelList {
	s.Keywords = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetNumber(v string) *TaskCallListResponseBodyModelList {
	s.Number = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetNumberMD5(v string) *TaskCallListResponseBodyModelList {
	s.NumberMD5 = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetProperties(v string) *TaskCallListResponseBodyModelList {
	s.Properties = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetRemark(v string) *TaskCallListResponseBodyModelList {
	s.Remark = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetRingTime(v int64) *TaskCallListResponseBodyModelList {
	s.RingTime = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetSms(v string) *TaskCallListResponseBodyModelList {
	s.Sms = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetSpeakingDuration(v int64) *TaskCallListResponseBodyModelList {
	s.SpeakingDuration = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetSpeakingTime(v string) *TaskCallListResponseBodyModelList {
	s.SpeakingTime = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetSpeakingTurns(v string) *TaskCallListResponseBodyModelList {
	s.SpeakingTurns = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetStatus(v string) *TaskCallListResponseBodyModelList {
	s.Status = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetStatusCode(v int64) *TaskCallListResponseBodyModelList {
	s.StatusCode = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetStatusDescription(v string) *TaskCallListResponseBodyModelList {
	s.StatusDescription = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetTag(v string) *TaskCallListResponseBodyModelList {
	s.Tag = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetTaskId(v int64) *TaskCallListResponseBodyModelList {
	s.TaskId = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetTemplateId(v int64) *TaskCallListResponseBodyModelList {
	s.TemplateId = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetTransferStatus(v string) *TaskCallListResponseBodyModelList {
	s.TransferStatus = &v
	return s
}

func (s *TaskCallListResponseBodyModelList) SetTransferStatusCode(v int64) *TaskCallListResponseBodyModelList {
	s.TransferStatusCode = &v
	return s
}

type TaskCallListResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TaskCallListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TaskCallListResponse) String() string {
	return tea.Prettify(s)
}

func (s TaskCallListResponse) GoString() string {
	return s.String()
}

func (s *TaskCallListResponse) SetHeaders(v map[string]*string) *TaskCallListResponse {
	s.Headers = v
	return s
}

func (s *TaskCallListResponse) SetStatusCode(v int32) *TaskCallListResponse {
	s.StatusCode = &v
	return s
}

func (s *TaskCallListResponse) SetBody(v *TaskCallListResponseBody) *TaskCallListResponse {
	s.Body = v
	return s
}

type TaskCancelCallRequest struct {
	Numbers              []*string `json:"Numbers,omitempty" xml:"Numbers,omitempty" type:"Repeated"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Tags                 []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// 任务ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s TaskCancelCallRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskCancelCallRequest) GoString() string {
	return s.String()
}

func (s *TaskCancelCallRequest) SetNumbers(v []*string) *TaskCancelCallRequest {
	s.Numbers = v
	return s
}

func (s *TaskCancelCallRequest) SetOwnerId(v int64) *TaskCancelCallRequest {
	s.OwnerId = &v
	return s
}

func (s *TaskCancelCallRequest) SetResourceOwnerAccount(v string) *TaskCancelCallRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TaskCancelCallRequest) SetResourceOwnerId(v int64) *TaskCancelCallRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TaskCancelCallRequest) SetTags(v []*string) *TaskCancelCallRequest {
	s.Tags = v
	return s
}

func (s *TaskCancelCallRequest) SetTaskId(v int64) *TaskCancelCallRequest {
	s.TaskId = &v
	return s
}

type TaskCancelCallShrinkRequest struct {
	NumbersShrink        *string `json:"Numbers,omitempty" xml:"Numbers,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TagsShrink           *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// 任务ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s TaskCancelCallShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskCancelCallShrinkRequest) GoString() string {
	return s.String()
}

func (s *TaskCancelCallShrinkRequest) SetNumbersShrink(v string) *TaskCancelCallShrinkRequest {
	s.NumbersShrink = &v
	return s
}

func (s *TaskCancelCallShrinkRequest) SetOwnerId(v int64) *TaskCancelCallShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *TaskCancelCallShrinkRequest) SetResourceOwnerAccount(v string) *TaskCancelCallShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TaskCancelCallShrinkRequest) SetResourceOwnerId(v int64) *TaskCancelCallShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TaskCancelCallShrinkRequest) SetTagsShrink(v string) *TaskCancelCallShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *TaskCancelCallShrinkRequest) SetTaskId(v int64) *TaskCancelCallShrinkRequest {
	s.TaskId = &v
	return s
}

type TaskCancelCallResponseBody struct {
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *TaskCancelCallResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 8EFC6D10-307B-1ECA-A8C6-7CBDF776AAD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1683440860035
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s TaskCancelCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TaskCancelCallResponseBody) GoString() string {
	return s.String()
}

func (s *TaskCancelCallResponseBody) SetCode(v string) *TaskCancelCallResponseBody {
	s.Code = &v
	return s
}

func (s *TaskCancelCallResponseBody) SetMessage(v string) *TaskCancelCallResponseBody {
	s.Message = &v
	return s
}

func (s *TaskCancelCallResponseBody) SetModel(v *TaskCancelCallResponseBodyModel) *TaskCancelCallResponseBody {
	s.Model = v
	return s
}

func (s *TaskCancelCallResponseBody) SetRequestId(v string) *TaskCancelCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *TaskCancelCallResponseBody) SetSuccess(v string) *TaskCancelCallResponseBody {
	s.Success = &v
	return s
}

func (s *TaskCancelCallResponseBody) SetTimestamp(v int64) *TaskCancelCallResponseBody {
	s.Timestamp = &v
	return s
}

type TaskCancelCallResponseBodyModel struct {
	// 错误手机号
	UnHandleNumbers []*string `json:"UnHandleNumbers,omitempty" xml:"UnHandleNumbers,omitempty" type:"Repeated"`
}

func (s TaskCancelCallResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s TaskCancelCallResponseBodyModel) GoString() string {
	return s.String()
}

func (s *TaskCancelCallResponseBodyModel) SetUnHandleNumbers(v []*string) *TaskCancelCallResponseBodyModel {
	s.UnHandleNumbers = v
	return s
}

type TaskCancelCallResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TaskCancelCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TaskCancelCallResponse) String() string {
	return tea.Prettify(s)
}

func (s TaskCancelCallResponse) GoString() string {
	return s.String()
}

func (s *TaskCancelCallResponse) SetHeaders(v map[string]*string) *TaskCancelCallResponse {
	s.Headers = v
	return s
}

func (s *TaskCancelCallResponse) SetStatusCode(v int32) *TaskCancelCallResponse {
	s.StatusCode = &v
	return s
}

func (s *TaskCancelCallResponse) SetBody(v *TaskCancelCallResponseBody) *TaskCancelCallResponse {
	s.Body = v
	return s
}

type TaskListRequest struct {
	// 创建时间
	//
	// example:
	//
	// 2023-04-05 12:11:11
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 最后外呼时间
	//
	// example:
	//
	// 2023-04-05 12:11:11
	LastCallTime         *string `json:"LastCallTime,omitempty" xml:"LastCallTime,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 任务状态。1 未启用，2 启用中，4 已停止
	//
	// example:
	//
	// 2
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 任务ID
	//
	// example:
	//
	// 2
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s TaskListRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskListRequest) GoString() string {
	return s.String()
}

func (s *TaskListRequest) SetCreateTime(v string) *TaskListRequest {
	s.CreateTime = &v
	return s
}

func (s *TaskListRequest) SetLastCallTime(v string) *TaskListRequest {
	s.LastCallTime = &v
	return s
}

func (s *TaskListRequest) SetOwnerId(v int64) *TaskListRequest {
	s.OwnerId = &v
	return s
}

func (s *TaskListRequest) SetResourceOwnerAccount(v string) *TaskListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TaskListRequest) SetResourceOwnerId(v int64) *TaskListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TaskListRequest) SetStatus(v int64) *TaskListRequest {
	s.Status = &v
	return s
}

func (s *TaskListRequest) SetTaskId(v int64) *TaskListRequest {
	s.TaskId = &v
	return s
}

type TaskListResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 30
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   []*TaskListResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Repeated"`
	// example:
	//
	// 8EFC6D10-307B-1ECA-A8C6-7CBDF776AAD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 80
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s TaskListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TaskListResponseBody) GoString() string {
	return s.String()
}

func (s *TaskListResponseBody) SetAccessDeniedDetail(v string) *TaskListResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *TaskListResponseBody) SetCode(v int64) *TaskListResponseBody {
	s.Code = &v
	return s
}

func (s *TaskListResponseBody) SetMessage(v string) *TaskListResponseBody {
	s.Message = &v
	return s
}

func (s *TaskListResponseBody) SetModel(v []*TaskListResponseBodyModel) *TaskListResponseBody {
	s.Model = v
	return s
}

func (s *TaskListResponseBody) SetRequestId(v string) *TaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *TaskListResponseBody) SetSuccess(v bool) *TaskListResponseBody {
	s.Success = &v
	return s
}

func (s *TaskListResponseBody) SetTimestamp(v int64) *TaskListResponseBody {
	s.Timestamp = &v
	return s
}

type TaskListResponseBodyModel struct {
	// 外呼时间段
	//
	// example:
	//
	// “8:00~20:30”
	AllowCallTime *string `json:"AllowCallTime,omitempty" xml:"AllowCallTime,omitempty"`
	// 外呼时间段格式化
	//
	// example:
	//
	// “8:00 ~ 20:00”
	AllowCallTimeFormat *string `json:"AllowCallTimeFormat,omitempty" xml:"AllowCallTimeFormat,omitempty"`
	// 外呼时间
	//
	// example:
	//
	// “1,2,3”
	AllowDayOfWeek *string `json:"AllowDayOfWeek,omitempty" xml:"AllowDayOfWeek,omitempty"`
	// 外呼类型
	//
	// example:
	//
	// 95
	CallType *int64 `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// 创建时间
	//
	// example:
	//
	// 2023-04-05 12:11:11
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 闪信模板id
	//
	// example:
	//
	// 99
	FlashSmsTemplateId *int64 `json:"FlashSmsTemplateId,omitempty" xml:"FlashSmsTemplateId,omitempty"`
	// 闪信模板名称
	//
	// example:
	//
	// 示例值示例值示例值
	FlashSmsTemplateName *string `json:"FlashSmsTemplateName,omitempty" xml:"FlashSmsTemplateName,omitempty"`
	// 发送闪信配置，可选0，1；0表示否，1表示是
	//
	// example:
	//
	// 1
	FlashSmsType *int64 `json:"FlashSmsType,omitempty" xml:"FlashSmsType,omitempty"`
	// 最近导入时间
	//
	// example:
	//
	// 2023-04-05 12:11:11
	ImportTime *string `json:"ImportTime,omitempty" xml:"ImportTime,omitempty"`
	// 意向标签列表
	IntentTags []*TaskListResponseBodyModelIntentTags `json:"IntentTags,omitempty" xml:"IntentTags,omitempty" type:"Repeated"`
	// 接通重呼次数
	//
	// example:
	//
	// 31
	InvalidReCall *int64 `json:"InvalidReCall,omitempty" xml:"InvalidReCall,omitempty"`
	// 最后外呼时间
	//
	// example:
	//
	// 2023-04-05 12:11:11
	LastCallTime *string `json:"LastCallTime,omitempty" xml:"LastCallTime,omitempty"`
	// 最大并发数
	//
	// example:
	//
	// 95
	MaxConcurrency *int64 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// 个性标签列表
	PersonalityTags []*string `json:"PersonalityTags,omitempty" xml:"PersonalityTags,omitempty" type:"Repeated"`
	// 优先任务
	//
	// example:
	//
	// 66
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// 任务所需参数
	//
	// example:
	//
	// 示例值示例值
	Properties *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// 自动重呼
	//
	// example:
	//
	// 20
	RecallType *int64 `json:"RecallType,omitempty" xml:"RecallType,omitempty"`
	// 发送短信
	//
	// example:
	//
	// 39
	SendSms *int64 `json:"SendSms,omitempty" xml:"SendSms,omitempty"`
	// 短信模板
	//
	// example:
	//
	// 示例值示例值
	SmsName *string `json:"SmsName,omitempty" xml:"SmsName,omitempty"`
	// 任务状态
	//
	// example:
	//
	// 79
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 任务ID
	//
	// example:
	//
	// 68
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 任务名称
	//
	// example:
	//
	// 示例值示例值
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// 话术模板Id
	//
	// example:
	//
	// 示例值示例值示例值
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 话术模板名称
	//
	// example:
	//
	// 示例值示例值
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s TaskListResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s TaskListResponseBodyModel) GoString() string {
	return s.String()
}

func (s *TaskListResponseBodyModel) SetAllowCallTime(v string) *TaskListResponseBodyModel {
	s.AllowCallTime = &v
	return s
}

func (s *TaskListResponseBodyModel) SetAllowCallTimeFormat(v string) *TaskListResponseBodyModel {
	s.AllowCallTimeFormat = &v
	return s
}

func (s *TaskListResponseBodyModel) SetAllowDayOfWeek(v string) *TaskListResponseBodyModel {
	s.AllowDayOfWeek = &v
	return s
}

func (s *TaskListResponseBodyModel) SetCallType(v int64) *TaskListResponseBodyModel {
	s.CallType = &v
	return s
}

func (s *TaskListResponseBodyModel) SetCreateTime(v string) *TaskListResponseBodyModel {
	s.CreateTime = &v
	return s
}

func (s *TaskListResponseBodyModel) SetFlashSmsTemplateId(v int64) *TaskListResponseBodyModel {
	s.FlashSmsTemplateId = &v
	return s
}

func (s *TaskListResponseBodyModel) SetFlashSmsTemplateName(v string) *TaskListResponseBodyModel {
	s.FlashSmsTemplateName = &v
	return s
}

func (s *TaskListResponseBodyModel) SetFlashSmsType(v int64) *TaskListResponseBodyModel {
	s.FlashSmsType = &v
	return s
}

func (s *TaskListResponseBodyModel) SetImportTime(v string) *TaskListResponseBodyModel {
	s.ImportTime = &v
	return s
}

func (s *TaskListResponseBodyModel) SetIntentTags(v []*TaskListResponseBodyModelIntentTags) *TaskListResponseBodyModel {
	s.IntentTags = v
	return s
}

func (s *TaskListResponseBodyModel) SetInvalidReCall(v int64) *TaskListResponseBodyModel {
	s.InvalidReCall = &v
	return s
}

func (s *TaskListResponseBodyModel) SetLastCallTime(v string) *TaskListResponseBodyModel {
	s.LastCallTime = &v
	return s
}

func (s *TaskListResponseBodyModel) SetMaxConcurrency(v int64) *TaskListResponseBodyModel {
	s.MaxConcurrency = &v
	return s
}

func (s *TaskListResponseBodyModel) SetPersonalityTags(v []*string) *TaskListResponseBodyModel {
	s.PersonalityTags = v
	return s
}

func (s *TaskListResponseBodyModel) SetPriority(v int64) *TaskListResponseBodyModel {
	s.Priority = &v
	return s
}

func (s *TaskListResponseBodyModel) SetProperties(v string) *TaskListResponseBodyModel {
	s.Properties = &v
	return s
}

func (s *TaskListResponseBodyModel) SetRecallType(v int64) *TaskListResponseBodyModel {
	s.RecallType = &v
	return s
}

func (s *TaskListResponseBodyModel) SetSendSms(v int64) *TaskListResponseBodyModel {
	s.SendSms = &v
	return s
}

func (s *TaskListResponseBodyModel) SetSmsName(v string) *TaskListResponseBodyModel {
	s.SmsName = &v
	return s
}

func (s *TaskListResponseBodyModel) SetStatus(v int64) *TaskListResponseBodyModel {
	s.Status = &v
	return s
}

func (s *TaskListResponseBodyModel) SetTaskId(v int64) *TaskListResponseBodyModel {
	s.TaskId = &v
	return s
}

func (s *TaskListResponseBodyModel) SetTaskName(v string) *TaskListResponseBodyModel {
	s.TaskName = &v
	return s
}

func (s *TaskListResponseBodyModel) SetTemplateId(v string) *TaskListResponseBodyModel {
	s.TemplateId = &v
	return s
}

func (s *TaskListResponseBodyModel) SetTemplateName(v string) *TaskListResponseBodyModel {
	s.TemplateName = &v
	return s
}

type TaskListResponseBodyModelIntentTags struct {
	// 意向标签描述
	//
	// example:
	//
	// 示例值示例值
	IntentDescription *string `json:"IntentDescription,omitempty" xml:"IntentDescription,omitempty"`
	// 意向标签ID
	//
	// example:
	//
	// 示例值示例值
	IntentTag *string `json:"IntentTag,omitempty" xml:"IntentTag,omitempty"`
}

func (s TaskListResponseBodyModelIntentTags) String() string {
	return tea.Prettify(s)
}

func (s TaskListResponseBodyModelIntentTags) GoString() string {
	return s.String()
}

func (s *TaskListResponseBodyModelIntentTags) SetIntentDescription(v string) *TaskListResponseBodyModelIntentTags {
	s.IntentDescription = &v
	return s
}

func (s *TaskListResponseBodyModelIntentTags) SetIntentTag(v string) *TaskListResponseBodyModelIntentTags {
	s.IntentTag = &v
	return s
}

type TaskListResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TaskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TaskListResponse) String() string {
	return tea.Prettify(s)
}

func (s TaskListResponse) GoString() string {
	return s.String()
}

func (s *TaskListResponse) SetHeaders(v map[string]*string) *TaskListResponse {
	s.Headers = v
	return s
}

func (s *TaskListResponse) SetStatusCode(v int32) *TaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *TaskListResponse) SetBody(v *TaskListResponseBody) *TaskListResponse {
	s.Body = v
	return s
}

type TaskRecoverCallRequest struct {
	// 查询开始导入时间
	//
	// example:
	//
	// "2023-01-09 18:58:19"
	BeginImportTime *string `json:"BeginImportTime,omitempty" xml:"BeginImportTime,omitempty"`
	// 查询结束导入时间
	//
	// example:
	//
	// "2023-01-09 18:58:19"
	EndImportTime *string `json:"EndImportTime,omitempty" xml:"EndImportTime,omitempty"`
	// 号码列表
	Numbers              []*string `json:"Numbers,omitempty" xml:"Numbers,omitempty" type:"Repeated"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 用户自定义标签列表
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// 任务ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 93
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s TaskRecoverCallRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskRecoverCallRequest) GoString() string {
	return s.String()
}

func (s *TaskRecoverCallRequest) SetBeginImportTime(v string) *TaskRecoverCallRequest {
	s.BeginImportTime = &v
	return s
}

func (s *TaskRecoverCallRequest) SetEndImportTime(v string) *TaskRecoverCallRequest {
	s.EndImportTime = &v
	return s
}

func (s *TaskRecoverCallRequest) SetNumbers(v []*string) *TaskRecoverCallRequest {
	s.Numbers = v
	return s
}

func (s *TaskRecoverCallRequest) SetOwnerId(v int64) *TaskRecoverCallRequest {
	s.OwnerId = &v
	return s
}

func (s *TaskRecoverCallRequest) SetResourceOwnerAccount(v string) *TaskRecoverCallRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TaskRecoverCallRequest) SetResourceOwnerId(v int64) *TaskRecoverCallRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TaskRecoverCallRequest) SetTags(v []*string) *TaskRecoverCallRequest {
	s.Tags = v
	return s
}

func (s *TaskRecoverCallRequest) SetTaskId(v int64) *TaskRecoverCallRequest {
	s.TaskId = &v
	return s
}

type TaskRecoverCallShrinkRequest struct {
	// 查询开始导入时间
	//
	// example:
	//
	// "2023-01-09 18:58:19"
	BeginImportTime *string `json:"BeginImportTime,omitempty" xml:"BeginImportTime,omitempty"`
	// 查询结束导入时间
	//
	// example:
	//
	// "2023-01-09 18:58:19"
	EndImportTime *string `json:"EndImportTime,omitempty" xml:"EndImportTime,omitempty"`
	// 号码列表
	NumbersShrink        *string `json:"Numbers,omitempty" xml:"Numbers,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 用户自定义标签列表
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// 任务ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 93
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s TaskRecoverCallShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskRecoverCallShrinkRequest) GoString() string {
	return s.String()
}

func (s *TaskRecoverCallShrinkRequest) SetBeginImportTime(v string) *TaskRecoverCallShrinkRequest {
	s.BeginImportTime = &v
	return s
}

func (s *TaskRecoverCallShrinkRequest) SetEndImportTime(v string) *TaskRecoverCallShrinkRequest {
	s.EndImportTime = &v
	return s
}

func (s *TaskRecoverCallShrinkRequest) SetNumbersShrink(v string) *TaskRecoverCallShrinkRequest {
	s.NumbersShrink = &v
	return s
}

func (s *TaskRecoverCallShrinkRequest) SetOwnerId(v int64) *TaskRecoverCallShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *TaskRecoverCallShrinkRequest) SetResourceOwnerAccount(v string) *TaskRecoverCallShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TaskRecoverCallShrinkRequest) SetResourceOwnerId(v int64) *TaskRecoverCallShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TaskRecoverCallShrinkRequest) SetTagsShrink(v string) *TaskRecoverCallShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *TaskRecoverCallShrinkRequest) SetTaskId(v int64) *TaskRecoverCallShrinkRequest {
	s.TaskId = &v
	return s
}

type TaskRecoverCallResponseBody struct {
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值
	Message *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   map[string]interface{} `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// 8EFC6D10-307B-1ECA-A8C6-7CBDF776AAD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1683440860035
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s TaskRecoverCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TaskRecoverCallResponseBody) GoString() string {
	return s.String()
}

func (s *TaskRecoverCallResponseBody) SetCode(v int64) *TaskRecoverCallResponseBody {
	s.Code = &v
	return s
}

func (s *TaskRecoverCallResponseBody) SetMessage(v string) *TaskRecoverCallResponseBody {
	s.Message = &v
	return s
}

func (s *TaskRecoverCallResponseBody) SetModel(v map[string]interface{}) *TaskRecoverCallResponseBody {
	s.Model = v
	return s
}

func (s *TaskRecoverCallResponseBody) SetRequestId(v string) *TaskRecoverCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *TaskRecoverCallResponseBody) SetSuccess(v string) *TaskRecoverCallResponseBody {
	s.Success = &v
	return s
}

func (s *TaskRecoverCallResponseBody) SetTimestamp(v int64) *TaskRecoverCallResponseBody {
	s.Timestamp = &v
	return s
}

type TaskRecoverCallResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TaskRecoverCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TaskRecoverCallResponse) String() string {
	return tea.Prettify(s)
}

func (s TaskRecoverCallResponse) GoString() string {
	return s.String()
}

func (s *TaskRecoverCallResponse) SetHeaders(v map[string]*string) *TaskRecoverCallResponse {
	s.Headers = v
	return s
}

func (s *TaskRecoverCallResponse) SetStatusCode(v int32) *TaskRecoverCallResponse {
	s.StatusCode = &v
	return s
}

func (s *TaskRecoverCallResponse) SetBody(v *TaskRecoverCallResponseBody) *TaskRecoverCallResponse {
	s.Body = v
	return s
}

type TemplateListRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 必须空参
	//
	// example:
	//
	// 9
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s TemplateListRequest) String() string {
	return tea.Prettify(s)
}

func (s TemplateListRequest) GoString() string {
	return s.String()
}

func (s *TemplateListRequest) SetOwnerId(v int64) *TemplateListRequest {
	s.OwnerId = &v
	return s
}

func (s *TemplateListRequest) SetResourceOwnerAccount(v string) *TemplateListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TemplateListRequest) SetResourceOwnerId(v int64) *TemplateListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TemplateListRequest) SetTemplateId(v int64) *TemplateListRequest {
	s.TemplateId = &v
	return s
}

type TemplateListResponseBody struct {
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   []*TemplateListResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Repeated"`
	// example:
	//
	// 8EFC6D10-307B-1ECA-A8C6-7CBDF776AAD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1683440860035
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s TemplateListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TemplateListResponseBody) GoString() string {
	return s.String()
}

func (s *TemplateListResponseBody) SetCode(v int64) *TemplateListResponseBody {
	s.Code = &v
	return s
}

func (s *TemplateListResponseBody) SetMessage(v string) *TemplateListResponseBody {
	s.Message = &v
	return s
}

func (s *TemplateListResponseBody) SetModel(v []*TemplateListResponseBodyModel) *TemplateListResponseBody {
	s.Model = v
	return s
}

func (s *TemplateListResponseBody) SetRequestId(v string) *TemplateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *TemplateListResponseBody) SetSuccess(v bool) *TemplateListResponseBody {
	s.Success = &v
	return s
}

func (s *TemplateListResponseBody) SetTimestamp(v int64) *TemplateListResponseBody {
	s.Timestamp = &v
	return s
}

type TemplateListResponseBodyModel struct {
	// 意向标签
	IntentTags []map[string]interface{} `json:"IntentTags,omitempty" xml:"IntentTags,omitempty" type:"Repeated"`
	// 个性标签
	PersonalityTags []*string `json:"PersonalityTags,omitempty" xml:"PersonalityTags,omitempty" type:"Repeated"`
	// 话术所需参数
	//
	// example:
	//
	// 示例值示例值
	Properties *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// AI话术ID
	//
	// example:
	//
	// 59
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 话术模板名称
	//
	// example:
	//
	// 示例值示例值
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// 模板类型
	//
	// example:
	//
	// 55
	TemplateType *int64 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s TemplateListResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s TemplateListResponseBodyModel) GoString() string {
	return s.String()
}

func (s *TemplateListResponseBodyModel) SetIntentTags(v []map[string]interface{}) *TemplateListResponseBodyModel {
	s.IntentTags = v
	return s
}

func (s *TemplateListResponseBodyModel) SetPersonalityTags(v []*string) *TemplateListResponseBodyModel {
	s.PersonalityTags = v
	return s
}

func (s *TemplateListResponseBodyModel) SetProperties(v string) *TemplateListResponseBodyModel {
	s.Properties = &v
	return s
}

func (s *TemplateListResponseBodyModel) SetTemplateId(v int64) *TemplateListResponseBodyModel {
	s.TemplateId = &v
	return s
}

func (s *TemplateListResponseBodyModel) SetTemplateName(v string) *TemplateListResponseBodyModel {
	s.TemplateName = &v
	return s
}

func (s *TemplateListResponseBodyModel) SetTemplateType(v int64) *TemplateListResponseBodyModel {
	s.TemplateType = &v
	return s
}

type TemplateListResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TemplateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TemplateListResponse) String() string {
	return tea.Prettify(s)
}

func (s TemplateListResponse) GoString() string {
	return s.String()
}

func (s *TemplateListResponse) SetHeaders(v map[string]*string) *TemplateListResponse {
	s.Headers = v
	return s
}

func (s *TemplateListResponse) SetStatusCode(v int32) *TemplateListResponse {
	s.StatusCode = &v
	return s
}

func (s *TemplateListResponse) SetBody(v *TemplateListResponseBody) *TemplateListResponse {
	s.Body = v
	return s
}

type UpdateAgentStatusRequest struct {
	// 坐席ID
	//
	// example:
	//
	// 58
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// 坐席状态 1:在线；2:忙碌；3:小休；4:离线
	//
	// example:
	//
	// 1
	AgentStatus *int64 `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	// 坐席标签
	//
	// example:
	//
	// abac
	AgentTag             *string `json:"AgentTag,omitempty" xml:"AgentTag,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateAgentStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAgentStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgentStatusRequest) SetAgentId(v int64) *UpdateAgentStatusRequest {
	s.AgentId = &v
	return s
}

func (s *UpdateAgentStatusRequest) SetAgentStatus(v int64) *UpdateAgentStatusRequest {
	s.AgentStatus = &v
	return s
}

func (s *UpdateAgentStatusRequest) SetAgentTag(v string) *UpdateAgentStatusRequest {
	s.AgentTag = &v
	return s
}

func (s *UpdateAgentStatusRequest) SetOwnerId(v int64) *UpdateAgentStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateAgentStatusRequest) SetResourceOwnerAccount(v string) *UpdateAgentStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateAgentStatusRequest) SetResourceOwnerId(v int64) *UpdateAgentStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

type UpdateAgentStatusResponseBody struct {
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   map[string]interface{} `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// 8EFC6D10-307B-1ECA-A8C6-7CBDF776AAD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1683440860035
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s UpdateAgentStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAgentStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAgentStatusResponseBody) SetCode(v int64) *UpdateAgentStatusResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAgentStatusResponseBody) SetMessage(v string) *UpdateAgentStatusResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAgentStatusResponseBody) SetModel(v map[string]interface{}) *UpdateAgentStatusResponseBody {
	s.Model = v
	return s
}

func (s *UpdateAgentStatusResponseBody) SetRequestId(v string) *UpdateAgentStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAgentStatusResponseBody) SetSuccess(v string) *UpdateAgentStatusResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAgentStatusResponseBody) SetTimestamp(v int64) *UpdateAgentStatusResponseBody {
	s.Timestamp = &v
	return s
}

type UpdateAgentStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAgentStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAgentStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAgentStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateAgentStatusResponse) SetHeaders(v map[string]*string) *UpdateAgentStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateAgentStatusResponse) SetStatusCode(v int32) *UpdateAgentStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAgentStatusResponse) SetBody(v *UpdateAgentStatusResponseBody) *UpdateAgentStatusResponse {
	s.Body = v
	return s
}

type UpdateTaskCustomerRequest struct {
	// 外呼客户
	//
	// This parameter is required.
	Customers            []*UpdateTaskCustomerRequestCustomers `json:"Customers,omitempty" xml:"Customers,omitempty" type:"Repeated"`
	OwnerId              *int64                                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                               `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                                `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 任务ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 59
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateTaskCustomerRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskCustomerRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskCustomerRequest) SetCustomers(v []*UpdateTaskCustomerRequestCustomers) *UpdateTaskCustomerRequest {
	s.Customers = v
	return s
}

func (s *UpdateTaskCustomerRequest) SetOwnerId(v int64) *UpdateTaskCustomerRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateTaskCustomerRequest) SetResourceOwnerAccount(v string) *UpdateTaskCustomerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateTaskCustomerRequest) SetResourceOwnerId(v int64) *UpdateTaskCustomerRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateTaskCustomerRequest) SetTaskId(v int64) *UpdateTaskCustomerRequest {
	s.TaskId = &v
	return s
}

type UpdateTaskCustomerRequestCustomers struct {
	// 是否取消外呼 0否，1是
	//
	// example:
	//
	// 0
	Cancel *int64 `json:"Cancel,omitempty" xml:"Cancel,omitempty"`
	// 电话号码
	//
	// example:
	//
	// 13661109390
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// 需根据具体任务参数要求传输
	//
	// example:
	//
	// {"test":"234"}
	Properties map[string]interface{} `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// 用户自定义标签
	//
	// example:
	//
	// tag1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s UpdateTaskCustomerRequestCustomers) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskCustomerRequestCustomers) GoString() string {
	return s.String()
}

func (s *UpdateTaskCustomerRequestCustomers) SetCancel(v int64) *UpdateTaskCustomerRequestCustomers {
	s.Cancel = &v
	return s
}

func (s *UpdateTaskCustomerRequestCustomers) SetNumber(v string) *UpdateTaskCustomerRequestCustomers {
	s.Number = &v
	return s
}

func (s *UpdateTaskCustomerRequestCustomers) SetProperties(v map[string]interface{}) *UpdateTaskCustomerRequestCustomers {
	s.Properties = v
	return s
}

func (s *UpdateTaskCustomerRequestCustomers) SetTag(v string) *UpdateTaskCustomerRequestCustomers {
	s.Tag = &v
	return s
}

type UpdateTaskCustomerShrinkRequest struct {
	// 外呼客户
	//
	// This parameter is required.
	CustomersShrink      *string `json:"Customers,omitempty" xml:"Customers,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 任务ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 59
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateTaskCustomerShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskCustomerShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskCustomerShrinkRequest) SetCustomersShrink(v string) *UpdateTaskCustomerShrinkRequest {
	s.CustomersShrink = &v
	return s
}

func (s *UpdateTaskCustomerShrinkRequest) SetOwnerId(v int64) *UpdateTaskCustomerShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateTaskCustomerShrinkRequest) SetResourceOwnerAccount(v string) *UpdateTaskCustomerShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateTaskCustomerShrinkRequest) SetResourceOwnerId(v int64) *UpdateTaskCustomerShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateTaskCustomerShrinkRequest) SetTaskId(v int64) *UpdateTaskCustomerShrinkRequest {
	s.TaskId = &v
	return s
}

type UpdateTaskCustomerResponseBody struct {
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *UpdateTaskCustomerResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 8EFC6D10-307B-1ECA-A8C6-7CBDF776AAD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1683440860035
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s UpdateTaskCustomerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskCustomerResponseBody) SetCode(v int64) *UpdateTaskCustomerResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTaskCustomerResponseBody) SetMessage(v string) *UpdateTaskCustomerResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTaskCustomerResponseBody) SetModel(v *UpdateTaskCustomerResponseBodyModel) *UpdateTaskCustomerResponseBody {
	s.Model = v
	return s
}

func (s *UpdateTaskCustomerResponseBody) SetRequestId(v string) *UpdateTaskCustomerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTaskCustomerResponseBody) SetSuccess(v string) *UpdateTaskCustomerResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTaskCustomerResponseBody) SetTimestamp(v int64) *UpdateTaskCustomerResponseBody {
	s.Timestamp = &v
	return s
}

type UpdateTaskCustomerResponseBodyModel struct {
	// 错误手机列表
	UnHandleNumbers []*string `json:"UnHandleNumbers,omitempty" xml:"UnHandleNumbers,omitempty" type:"Repeated"`
}

func (s UpdateTaskCustomerResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskCustomerResponseBodyModel) GoString() string {
	return s.String()
}

func (s *UpdateTaskCustomerResponseBodyModel) SetUnHandleNumbers(v []*string) *UpdateTaskCustomerResponseBodyModel {
	s.UnHandleNumbers = v
	return s
}

type UpdateTaskCustomerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTaskCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTaskCustomerResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskCustomerResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskCustomerResponse) SetHeaders(v map[string]*string) *UpdateTaskCustomerResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskCustomerResponse) SetStatusCode(v int32) *UpdateTaskCustomerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskCustomerResponse) SetBody(v *UpdateTaskCustomerResponseBody) *UpdateTaskCustomerResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("aiccs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加黑名单接口
//
// @param tmpReq - AddBlacklistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddBlacklistResponse
func (client *Client) AddBlacklistWithOptions(tmpReq *AddBlacklistRequest, runtime *util.RuntimeOptions) (_result *AddBlacklistResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AddBlacklistShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Numbers)) {
		request.NumbersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Numbers, tea.String("Numbers"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExpiredDay)) {
		query["ExpiredDay"] = request.ExpiredDay
	}

	if !tea.BoolValue(util.IsUnset(request.NumbersShrink)) {
		query["Numbers"] = request.NumbersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddBlacklist"),
		Version:     tea.String("2023-05-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddBlacklistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加黑名单接口
//
// @param request - AddBlacklistRequest
//
// @return AddBlacklistResponse
func (client *Client) AddBlacklist(request *AddBlacklistRequest) (_result *AddBlacklistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddBlacklistResponse{}
	_body, _err := client.AddBlacklistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建任务接口
//
// @param tmpReq - AddTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTaskResponse
func (client *Client) AddTaskWithOptions(tmpReq *AddTaskRequest, runtime *util.RuntimeOptions) (_result *AddTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AddTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CallTimeList)) {
		request.CallTimeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CallTimeList, tea.String("CallTimeList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.RepeatReason)) {
		request.RepeatReasonShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RepeatReason, tea.String("RepeatReason"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.RepeatTimes)) {
		request.RepeatTimesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RepeatTimes, tea.String("RepeatTimes"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SendSmsPlan)) {
		request.SendSmsPlanShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SendSmsPlan, tea.String("SendSmsPlan"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallTimeListShrink)) {
		query["CallTimeList"] = request.CallTimeListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackUrl)) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !tea.BoolValue(util.IsUnset(request.FlashSmsTemplateId)) {
		query["FlashSmsTemplateId"] = request.FlashSmsTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.FlashSmsType)) {
		query["FlashSmsType"] = request.FlashSmsType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxConcurrency)) {
		query["MaxConcurrency"] = request.MaxConcurrency
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlaySleepVal)) {
		query["PlaySleepVal"] = request.PlaySleepVal
	}

	if !tea.BoolValue(util.IsUnset(request.PlayTimes)) {
		query["PlayTimes"] = request.PlayTimes
	}

	if !tea.BoolValue(util.IsUnset(request.RecallType)) {
		query["RecallType"] = request.RecallType
	}

	if !tea.BoolValue(util.IsUnset(request.RecordPath)) {
		query["RecordPath"] = request.RecordPath
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatCount)) {
		query["RepeatCount"] = request.RepeatCount
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatInterval)) {
		query["RepeatInterval"] = request.RepeatInterval
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatReasonShrink)) {
		query["RepeatReason"] = request.RepeatReasonShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatTimesShrink)) {
		query["RepeatTimes"] = request.RepeatTimesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SendSmsPlanShrink)) {
		query["SendSmsPlan"] = request.SendSmsPlanShrink
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddTask"),
		Version:     tea.String("2023-05-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建任务接口
//
// @param request - AddTaskRequest
//
// @return AddTaskResponse
func (client *Client) AddTask(request *AddTaskRequest) (_result *AddTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddTaskResponse{}
	_body, _err := client.AddTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 坐席取消号码外呼
//
// @param tmpReq - AgentCancelCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AgentCancelCallResponse
func (client *Client) AgentCancelCallWithOptions(tmpReq *AgentCancelCallRequest, runtime *util.RuntimeOptions) (_result *AgentCancelCallResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AgentCancelCallShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Numbers)) {
		request.NumbersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Numbers, tea.String("Numbers"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		query["AgentId"] = request.AgentId
	}

	if !tea.BoolValue(util.IsUnset(request.AgentTag)) {
		query["AgentTag"] = request.AgentTag
	}

	if !tea.BoolValue(util.IsUnset(request.NumbersShrink)) {
		query["Numbers"] = request.NumbersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AgentCancelCall"),
		Version:     tea.String("2023-05-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AgentCancelCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 坐席取消号码外呼
//
// @param request - AgentCancelCallRequest
//
// @return AgentCancelCallResponse
func (client *Client) AgentCancelCall(request *AgentCancelCallRequest) (_result *AgentCancelCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AgentCancelCallResponse{}
	_body, _err := client.AgentCancelCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 坐席任务恢复号码
//
// @param tmpReq - AgentRecoverCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AgentRecoverCallResponse
func (client *Client) AgentRecoverCallWithOptions(tmpReq *AgentRecoverCallRequest, runtime *util.RuntimeOptions) (_result *AgentRecoverCallResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AgentRecoverCallShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Numbers)) {
		request.NumbersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Numbers, tea.String("Numbers"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		query["AgentId"] = request.AgentId
	}

	if !tea.BoolValue(util.IsUnset(request.AgentTag)) {
		query["AgentTag"] = request.AgentTag
	}

	if !tea.BoolValue(util.IsUnset(request.BeginImportTime)) {
		query["BeginImportTime"] = request.BeginImportTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndImportTime)) {
		query["EndImportTime"] = request.EndImportTime
	}

	if !tea.BoolValue(util.IsUnset(request.NumbersShrink)) {
		query["Numbers"] = request.NumbersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AgentRecoverCall"),
		Version:     tea.String("2023-05-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AgentRecoverCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 坐席任务恢复号码
//
// @param request - AgentRecoverCallRequest
//
// @return AgentRecoverCallResponse
func (client *Client) AgentRecoverCall(request *AgentRecoverCallRequest) (_result *AgentRecoverCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AgentRecoverCallResponse{}
	_body, _err := client.AgentRecoverCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// AI批量任务查询号码状态接口
//
// @param tmpReq - DetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetailsResponse
func (client *Client) DetailsWithOptions(tmpReq *DetailsRequest, runtime *util.RuntimeOptions) (_result *DetailsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DetailsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Numbers)) {
		request.NumbersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Numbers, tea.String("Numbers"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BatchId)) {
		query["BatchId"] = request.BatchId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.NumberStatus)) {
		query["NumberStatus"] = request.NumberStatus
	}

	if !tea.BoolValue(util.IsUnset(request.NumbersShrink)) {
		query["Numbers"] = request.NumbersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Details"),
		Version:     tea.String("2023-05-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// AI批量任务查询号码状态接口
//
// @param request - DetailsRequest
//
// @return DetailsResponse
func (client *Client) Details(request *DetailsRequest) (_result *DetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetailsResponse{}
	_body, _err := client.DetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 编辑任务接口
//
// @param tmpReq - EditTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditTaskResponse
func (client *Client) EditTaskWithOptions(tmpReq *EditTaskRequest, runtime *util.RuntimeOptions) (_result *EditTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &EditTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CallTimeList)) {
		request.CallTimeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CallTimeList, tea.String("CallTimeList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.RepeatReason)) {
		request.RepeatReasonShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RepeatReason, tea.String("RepeatReason"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.RepeatTimes)) {
		request.RepeatTimesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RepeatTimes, tea.String("RepeatTimes"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SendSmsPlan)) {
		request.SendSmsPlanShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SendSmsPlan, tea.String("SendSmsPlan"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallTimeListShrink)) {
		query["CallTimeList"] = request.CallTimeListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackUrl)) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !tea.BoolValue(util.IsUnset(request.FlashSmsTemplateId)) {
		query["FlashSmsTemplateId"] = request.FlashSmsTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.FlashSmsType)) {
		query["FlashSmsType"] = request.FlashSmsType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxConcurrency)) {
		query["MaxConcurrency"] = request.MaxConcurrency
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlaySleepVal)) {
		query["PlaySleepVal"] = request.PlaySleepVal
	}

	if !tea.BoolValue(util.IsUnset(request.PlayTimes)) {
		query["PlayTimes"] = request.PlayTimes
	}

	if !tea.BoolValue(util.IsUnset(request.RecallType)) {
		query["RecallType"] = request.RecallType
	}

	if !tea.BoolValue(util.IsUnset(request.RecordPath)) {
		query["RecordPath"] = request.RecordPath
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatCount)) {
		query["RepeatCount"] = request.RepeatCount
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatInterval)) {
		query["RepeatInterval"] = request.RepeatInterval
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatReasonShrink)) {
		query["RepeatReason"] = request.RepeatReasonShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatTimesShrink)) {
		query["RepeatTimes"] = request.RepeatTimesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SendSmsPlanShrink)) {
		query["SendSmsPlan"] = request.SendSmsPlanShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EditTask"),
		Version:     tea.String("2023-05-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EditTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑任务接口
//
// @param request - EditTaskRequest
//
// @return EditTaskResponse
func (client *Client) EditTask(request *EditTaskRequest) (_result *EditTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EditTaskResponse{}
	_body, _err := client.EditTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导入号码
//
// @param tmpReq - ImportNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportNumberResponse
func (client *Client) ImportNumberWithOptions(tmpReq *ImportNumberRequest, runtime *util.RuntimeOptions) (_result *ImportNumberResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ImportNumberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Customers)) {
		request.CustomersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Customers, tea.String("Customers"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomersShrink)) {
		query["Customers"] = request.CustomersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.FailReturn)) {
		query["FailReturn"] = request.FailReturn
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportNumber"),
		Version:     tea.String("2023-05-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImportNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导入号码
//
// @param request - ImportNumberRequest
//
// @return ImportNumberResponse
func (client *Client) ImportNumber(request *ImportNumberRequest) (_result *ImportNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportNumberResponse{}
	_body, _err := client.ImportNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询企业黑名单
//
// @param tmpReq - PageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PageResponse
func (client *Client) PageWithOptions(tmpReq *PageRequest, runtime *util.RuntimeOptions) (_result *PageResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &PageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Numbers)) {
		request.NumbersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Numbers, tea.String("Numbers"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NumbersShrink)) {
		query["Numbers"] = request.NumbersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Page"),
		Version:     tea.String("2023-05-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询企业黑名单
//
// @param request - PageRequest
//
// @return PageResponse
func (client *Client) Page(request *PageRequest) (_result *PageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PageResponse{}
	_body, _err := client.PageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 短信模板创建
//
// @param request - SmsTemplateCreateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SmsTemplateCreateResponse
func (client *Client) SmsTemplateCreateWithOptions(request *SmsTemplateCreateRequest, runtime *util.RuntimeOptions) (_result *SmsTemplateCreateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Sign)) {
		query["Sign"] = request.Sign
	}

	if !tea.BoolValue(util.IsUnset(request.SmsType)) {
		query["SmsType"] = request.SmsType
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SmsTemplateCreate"),
		Version:     tea.String("2023-05-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SmsTemplateCreateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 短信模板创建
//
// @param request - SmsTemplateCreateRequest
//
// @return SmsTemplateCreateResponse
func (client *Client) SmsTemplateCreate(request *SmsTemplateCreateRequest) (_result *SmsTemplateCreateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SmsTemplateCreateResponse{}
	_body, _err := client.SmsTemplateCreateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 短信模板列表查询
//
// @param request - SmsTemplatePageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SmsTemplatePageListResponse
func (client *Client) SmsTemplatePageListWithOptions(request *SmsTemplatePageListRequest, runtime *util.RuntimeOptions) (_result *SmsTemplatePageListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Sign)) {
		query["Sign"] = request.Sign
	}

	if !tea.BoolValue(util.IsUnset(request.SmsType)) {
		query["SmsType"] = request.SmsType
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SmsTemplatePageList"),
		Version:     tea.String("2023-05-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SmsTemplatePageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 短信模板列表查询
//
// @param request - SmsTemplatePageListRequest
//
// @return SmsTemplatePageListResponse
func (client *Client) SmsTemplatePageList(request *SmsTemplatePageListRequest) (_result *SmsTemplatePageListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SmsTemplatePageListResponse{}
	_body, _err := client.SmsTemplatePageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询聊天记录接口
//
// @param request - TaskCallChatsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TaskCallChatsResponse
func (client *Client) TaskCallChatsWithOptions(request *TaskCallChatsRequest, runtime *util.RuntimeOptions) (_result *TaskCallChatsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		query["AgentId"] = request.AgentId
	}

	if !tea.BoolValue(util.IsUnset(request.AgentTag)) {
		query["AgentTag"] = request.AgentTag
	}

	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		query["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TaskCallChats"),
		Version:     tea.String("2023-05-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TaskCallChatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询聊天记录接口
//
// @param request - TaskCallChatsRequest
//
// @return TaskCallChatsResponse
func (client *Client) TaskCallChats(request *TaskCallChatsRequest) (_result *TaskCallChatsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TaskCallChatsResponse{}
	_body, _err := client.TaskCallChatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取任务外呼情况接口
//
// @param request - TaskCallInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TaskCallInfoResponse
func (client *Client) TaskCallInfoWithOptions(request *TaskCallInfoRequest, runtime *util.RuntimeOptions) (_result *TaskCallInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TaskCallInfo"),
		Version:     tea.String("2023-05-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TaskCallInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务外呼情况接口
//
// @param request - TaskCallInfoRequest
//
// @return TaskCallInfoResponse
func (client *Client) TaskCallInfo(request *TaskCallInfoRequest) (_result *TaskCallInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TaskCallInfoResponse{}
	_body, _err := client.TaskCallInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// AI批量任务查询外呼记录接口
//
// @param tmpReq - TaskCallListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TaskCallListResponse
func (client *Client) TaskCallListWithOptions(tmpReq *TaskCallListRequest, runtime *util.RuntimeOptions) (_result *TaskCallListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &TaskCallListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.IntentTags)) {
		request.IntentTagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IntentTags, tea.String("IntentTags"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Numbers)) {
		request.NumbersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Numbers, tea.String("Numbers"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BatchId)) {
		query["BatchId"] = request.BatchId
	}

	if !tea.BoolValue(util.IsUnset(request.CallDate)) {
		query["CallDate"] = request.CallDate
	}

	if !tea.BoolValue(util.IsUnset(request.EndCallDate)) {
		query["EndCallDate"] = request.EndCallDate
	}

	if !tea.BoolValue(util.IsUnset(request.IntentTagsShrink)) {
		query["IntentTags"] = request.IntentTagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NumbersShrink)) {
		query["Numbers"] = request.NumbersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TaskCallList"),
		Version:     tea.String("2023-05-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TaskCallListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// AI批量任务查询外呼记录接口
//
// @param request - TaskCallListRequest
//
// @return TaskCallListResponse
func (client *Client) TaskCallList(request *TaskCallListRequest) (_result *TaskCallListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TaskCallListResponse{}
	_body, _err := client.TaskCallListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量任务取消号码外呼
//
// @param tmpReq - TaskCancelCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TaskCancelCallResponse
func (client *Client) TaskCancelCallWithOptions(tmpReq *TaskCancelCallRequest, runtime *util.RuntimeOptions) (_result *TaskCancelCallResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &TaskCancelCallShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Numbers)) {
		request.NumbersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Numbers, tea.String("Numbers"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NumbersShrink)) {
		query["Numbers"] = request.NumbersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TaskCancelCall"),
		Version:     tea.String("2023-05-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TaskCancelCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量任务取消号码外呼
//
// @param request - TaskCancelCallRequest
//
// @return TaskCancelCallResponse
func (client *Client) TaskCancelCall(request *TaskCancelCallRequest) (_result *TaskCancelCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TaskCancelCallResponse{}
	_body, _err := client.TaskCancelCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询任务列表接口
//
// @param request - TaskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TaskListResponse
func (client *Client) TaskListWithOptions(request *TaskListRequest, runtime *util.RuntimeOptions) (_result *TaskListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateTime)) {
		query["CreateTime"] = request.CreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.LastCallTime)) {
		query["LastCallTime"] = request.LastCallTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TaskList"),
		Version:     tea.String("2023-05-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TaskListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询任务列表接口
//
// @param request - TaskListRequest
//
// @return TaskListResponse
func (client *Client) TaskList(request *TaskListRequest) (_result *TaskListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TaskListResponse{}
	_body, _err := client.TaskListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量任务恢复号码
//
// @param tmpReq - TaskRecoverCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TaskRecoverCallResponse
func (client *Client) TaskRecoverCallWithOptions(tmpReq *TaskRecoverCallRequest, runtime *util.RuntimeOptions) (_result *TaskRecoverCallResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &TaskRecoverCallShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Numbers)) {
		request.NumbersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Numbers, tea.String("Numbers"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginImportTime)) {
		query["BeginImportTime"] = request.BeginImportTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndImportTime)) {
		query["EndImportTime"] = request.EndImportTime
	}

	if !tea.BoolValue(util.IsUnset(request.NumbersShrink)) {
		query["Numbers"] = request.NumbersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TaskRecoverCall"),
		Version:     tea.String("2023-05-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TaskRecoverCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量任务恢复号码
//
// @param request - TaskRecoverCallRequest
//
// @return TaskRecoverCallResponse
func (client *Client) TaskRecoverCall(request *TaskRecoverCallRequest) (_result *TaskRecoverCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TaskRecoverCallResponse{}
	_body, _err := client.TaskRecoverCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 话术模板列表查询接口
//
// @param request - TemplateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TemplateListResponse
func (client *Client) TemplateListWithOptions(request *TemplateListRequest, runtime *util.RuntimeOptions) (_result *TemplateListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TemplateList"),
		Version:     tea.String("2023-05-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TemplateListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 话术模板列表查询接口
//
// @param request - TemplateListRequest
//
// @return TemplateListResponse
func (client *Client) TemplateList(request *TemplateListRequest) (_result *TemplateListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TemplateListResponse{}
	_body, _err := client.TemplateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改坐席状态
//
// @param request - UpdateAgentStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAgentStatusResponse
func (client *Client) UpdateAgentStatusWithOptions(request *UpdateAgentStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateAgentStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		query["AgentId"] = request.AgentId
	}

	if !tea.BoolValue(util.IsUnset(request.AgentStatus)) {
		query["AgentStatus"] = request.AgentStatus
	}

	if !tea.BoolValue(util.IsUnset(request.AgentTag)) {
		query["AgentTag"] = request.AgentTag
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAgentStatus"),
		Version:     tea.String("2023-05-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAgentStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改坐席状态
//
// @param request - UpdateAgentStatusRequest
//
// @return UpdateAgentStatusResponse
func (client *Client) UpdateAgentStatus(request *UpdateAgentStatusRequest) (_result *UpdateAgentStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAgentStatusResponse{}
	_body, _err := client.UpdateAgentStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新当天导入的号码
//
// @param tmpReq - UpdateTaskCustomerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskCustomerResponse
func (client *Client) UpdateTaskCustomerWithOptions(tmpReq *UpdateTaskCustomerRequest, runtime *util.RuntimeOptions) (_result *UpdateTaskCustomerResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateTaskCustomerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Customers)) {
		request.CustomersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Customers, tea.String("Customers"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomersShrink)) {
		query["Customers"] = request.CustomersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTaskCustomer"),
		Version:     tea.String("2023-05-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTaskCustomerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新当天导入的号码
//
// @param request - UpdateTaskCustomerRequest
//
// @return UpdateTaskCustomerResponse
func (client *Client) UpdateTaskCustomer(request *UpdateTaskCustomerRequest) (_result *UpdateTaskCustomerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTaskCustomerResponse{}
	_body, _err := client.UpdateTaskCustomerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
