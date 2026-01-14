// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallTimeList(v []*AddTaskRequestCallTimeList) *AddTaskRequest
	GetCallTimeList() []*AddTaskRequestCallTimeList
	SetCallbackUrl(v string) *AddTaskRequest
	GetCallbackUrl() *string
	SetFlashSmsTemplateId(v int64) *AddTaskRequest
	GetFlashSmsTemplateId() *int64
	SetFlashSmsType(v int64) *AddTaskRequest
	GetFlashSmsType() *int64
	SetMaxConcurrency(v int64) *AddTaskRequest
	GetMaxConcurrency() *int64
	SetName(v string) *AddTaskRequest
	GetName() *string
	SetOwnerId(v int64) *AddTaskRequest
	GetOwnerId() *int64
	SetPlaySleepVal(v int64) *AddTaskRequest
	GetPlaySleepVal() *int64
	SetPlayTimes(v int64) *AddTaskRequest
	GetPlayTimes() *int64
	SetRecallType(v int64) *AddTaskRequest
	GetRecallType() *int64
	SetRecordPath(v string) *AddTaskRequest
	GetRecordPath() *string
	SetRepeatCount(v int64) *AddTaskRequest
	GetRepeatCount() *int64
	SetRepeatInterval(v int64) *AddTaskRequest
	GetRepeatInterval() *int64
	SetRepeatReason(v []*string) *AddTaskRequest
	GetRepeatReason() []*string
	SetRepeatTimes(v []*string) *AddTaskRequest
	GetRepeatTimes() []*string
	SetResourceOwnerAccount(v string) *AddTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddTaskRequest
	GetResourceOwnerId() *int64
	SetSendSmsPlan(v []*AddTaskRequestSendSmsPlan) *AddTaskRequest
	GetSendSmsPlan() []*AddTaskRequestSendSmsPlan
	SetStartTime(v string) *AddTaskRequest
	GetStartTime() *string
	SetTaskType(v int64) *AddTaskRequest
	GetTaskType() *int64
	SetTemplateId(v int64) *AddTaskRequest
	GetTemplateId() *int64
	SetTemplateType(v int64) *AddTaskRequest
	GetTemplateType() *int64
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
	return dara.Prettify(s)
}

func (s AddTaskRequest) GoString() string {
	return s.String()
}

func (s *AddTaskRequest) GetCallTimeList() []*AddTaskRequestCallTimeList {
	return s.CallTimeList
}

func (s *AddTaskRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *AddTaskRequest) GetFlashSmsTemplateId() *int64 {
	return s.FlashSmsTemplateId
}

func (s *AddTaskRequest) GetFlashSmsType() *int64 {
	return s.FlashSmsType
}

func (s *AddTaskRequest) GetMaxConcurrency() *int64 {
	return s.MaxConcurrency
}

func (s *AddTaskRequest) GetName() *string {
	return s.Name
}

func (s *AddTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddTaskRequest) GetPlaySleepVal() *int64 {
	return s.PlaySleepVal
}

func (s *AddTaskRequest) GetPlayTimes() *int64 {
	return s.PlayTimes
}

func (s *AddTaskRequest) GetRecallType() *int64 {
	return s.RecallType
}

func (s *AddTaskRequest) GetRecordPath() *string {
	return s.RecordPath
}

func (s *AddTaskRequest) GetRepeatCount() *int64 {
	return s.RepeatCount
}

func (s *AddTaskRequest) GetRepeatInterval() *int64 {
	return s.RepeatInterval
}

func (s *AddTaskRequest) GetRepeatReason() []*string {
	return s.RepeatReason
}

func (s *AddTaskRequest) GetRepeatTimes() []*string {
	return s.RepeatTimes
}

func (s *AddTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddTaskRequest) GetSendSmsPlan() []*AddTaskRequestSendSmsPlan {
	return s.SendSmsPlan
}

func (s *AddTaskRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *AddTaskRequest) GetTaskType() *int64 {
	return s.TaskType
}

func (s *AddTaskRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *AddTaskRequest) GetTemplateType() *int64 {
	return s.TemplateType
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

func (s *AddTaskRequest) Validate() error {
	if s.CallTimeList != nil {
		for _, item := range s.CallTimeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SendSmsPlan != nil {
		for _, item := range s.SendSmsPlan {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddTaskRequestCallTimeList struct {
	CallTime []*string `json:"CallTime,omitempty" xml:"CallTime,omitempty" type:"Repeated"`
}

func (s AddTaskRequestCallTimeList) String() string {
	return dara.Prettify(s)
}

func (s AddTaskRequestCallTimeList) GoString() string {
	return s.String()
}

func (s *AddTaskRequestCallTimeList) GetCallTime() []*string {
	return s.CallTime
}

func (s *AddTaskRequestCallTimeList) SetCallTime(v []*string) *AddTaskRequestCallTimeList {
	s.CallTime = v
	return s
}

func (s *AddTaskRequestCallTimeList) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s AddTaskRequestSendSmsPlan) GoString() string {
	return s.String()
}

func (s *AddTaskRequestSendSmsPlan) GetIntentTags() []*string {
	return s.IntentTags
}

func (s *AddTaskRequestSendSmsPlan) GetSmsTemplateId() *int64 {
	return s.SmsTemplateId
}

func (s *AddTaskRequestSendSmsPlan) SetIntentTags(v []*string) *AddTaskRequestSendSmsPlan {
	s.IntentTags = v
	return s
}

func (s *AddTaskRequestSendSmsPlan) SetSmsTemplateId(v int64) *AddTaskRequestSendSmsPlan {
	s.SmsTemplateId = &v
	return s
}

func (s *AddTaskRequestSendSmsPlan) Validate() error {
	return dara.Validate(s)
}
