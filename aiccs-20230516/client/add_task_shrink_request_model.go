// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallTimeListShrink(v string) *AddTaskShrinkRequest
	GetCallTimeListShrink() *string
	SetCallbackUrl(v string) *AddTaskShrinkRequest
	GetCallbackUrl() *string
	SetFlashSmsTemplateId(v int64) *AddTaskShrinkRequest
	GetFlashSmsTemplateId() *int64
	SetFlashSmsType(v int64) *AddTaskShrinkRequest
	GetFlashSmsType() *int64
	SetMaxConcurrency(v int64) *AddTaskShrinkRequest
	GetMaxConcurrency() *int64
	SetName(v string) *AddTaskShrinkRequest
	GetName() *string
	SetOwnerId(v int64) *AddTaskShrinkRequest
	GetOwnerId() *int64
	SetPlaySleepVal(v int64) *AddTaskShrinkRequest
	GetPlaySleepVal() *int64
	SetPlayTimes(v int64) *AddTaskShrinkRequest
	GetPlayTimes() *int64
	SetRecallType(v int64) *AddTaskShrinkRequest
	GetRecallType() *int64
	SetRecordPath(v string) *AddTaskShrinkRequest
	GetRecordPath() *string
	SetRepeatCount(v int64) *AddTaskShrinkRequest
	GetRepeatCount() *int64
	SetRepeatInterval(v int64) *AddTaskShrinkRequest
	GetRepeatInterval() *int64
	SetRepeatReasonShrink(v string) *AddTaskShrinkRequest
	GetRepeatReasonShrink() *string
	SetRepeatTimesShrink(v string) *AddTaskShrinkRequest
	GetRepeatTimesShrink() *string
	SetResourceOwnerAccount(v string) *AddTaskShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddTaskShrinkRequest
	GetResourceOwnerId() *int64
	SetSendSmsPlanShrink(v string) *AddTaskShrinkRequest
	GetSendSmsPlanShrink() *string
	SetStartTime(v string) *AddTaskShrinkRequest
	GetStartTime() *string
	SetTaskType(v int64) *AddTaskShrinkRequest
	GetTaskType() *int64
	SetTemplateId(v int64) *AddTaskShrinkRequest
	GetTemplateId() *int64
	SetTemplateType(v int64) *AddTaskShrinkRequest
	GetTemplateType() *int64
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
	return dara.Prettify(s)
}

func (s AddTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddTaskShrinkRequest) GetCallTimeListShrink() *string {
	return s.CallTimeListShrink
}

func (s *AddTaskShrinkRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *AddTaskShrinkRequest) GetFlashSmsTemplateId() *int64 {
	return s.FlashSmsTemplateId
}

func (s *AddTaskShrinkRequest) GetFlashSmsType() *int64 {
	return s.FlashSmsType
}

func (s *AddTaskShrinkRequest) GetMaxConcurrency() *int64 {
	return s.MaxConcurrency
}

func (s *AddTaskShrinkRequest) GetName() *string {
	return s.Name
}

func (s *AddTaskShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddTaskShrinkRequest) GetPlaySleepVal() *int64 {
	return s.PlaySleepVal
}

func (s *AddTaskShrinkRequest) GetPlayTimes() *int64 {
	return s.PlayTimes
}

func (s *AddTaskShrinkRequest) GetRecallType() *int64 {
	return s.RecallType
}

func (s *AddTaskShrinkRequest) GetRecordPath() *string {
	return s.RecordPath
}

func (s *AddTaskShrinkRequest) GetRepeatCount() *int64 {
	return s.RepeatCount
}

func (s *AddTaskShrinkRequest) GetRepeatInterval() *int64 {
	return s.RepeatInterval
}

func (s *AddTaskShrinkRequest) GetRepeatReasonShrink() *string {
	return s.RepeatReasonShrink
}

func (s *AddTaskShrinkRequest) GetRepeatTimesShrink() *string {
	return s.RepeatTimesShrink
}

func (s *AddTaskShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddTaskShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddTaskShrinkRequest) GetSendSmsPlanShrink() *string {
	return s.SendSmsPlanShrink
}

func (s *AddTaskShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *AddTaskShrinkRequest) GetTaskType() *int64 {
	return s.TaskType
}

func (s *AddTaskShrinkRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *AddTaskShrinkRequest) GetTemplateType() *int64 {
	return s.TemplateType
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

func (s *AddTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
