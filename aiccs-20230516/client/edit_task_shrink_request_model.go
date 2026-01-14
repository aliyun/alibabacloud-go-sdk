// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditTaskShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetCallTimeListShrink(v string) *EditTaskShrinkRequest
  GetCallTimeListShrink() *string 
  SetCallbackUrl(v string) *EditTaskShrinkRequest
  GetCallbackUrl() *string 
  SetFlashSmsTemplateId(v int64) *EditTaskShrinkRequest
  GetFlashSmsTemplateId() *int64 
  SetFlashSmsType(v int64) *EditTaskShrinkRequest
  GetFlashSmsType() *int64 
  SetMaxConcurrency(v int64) *EditTaskShrinkRequest
  GetMaxConcurrency() *int64 
  SetName(v string) *EditTaskShrinkRequest
  GetName() *string 
  SetOwnerId(v int64) *EditTaskShrinkRequest
  GetOwnerId() *int64 
  SetPlaySleepVal(v int64) *EditTaskShrinkRequest
  GetPlaySleepVal() *int64 
  SetPlayTimes(v int64) *EditTaskShrinkRequest
  GetPlayTimes() *int64 
  SetRecallType(v int64) *EditTaskShrinkRequest
  GetRecallType() *int64 
  SetRecordPath(v string) *EditTaskShrinkRequest
  GetRecordPath() *string 
  SetRepeatCount(v int64) *EditTaskShrinkRequest
  GetRepeatCount() *int64 
  SetRepeatInterval(v int64) *EditTaskShrinkRequest
  GetRepeatInterval() *int64 
  SetRepeatReasonShrink(v string) *EditTaskShrinkRequest
  GetRepeatReasonShrink() *string 
  SetRepeatTimesShrink(v string) *EditTaskShrinkRequest
  GetRepeatTimesShrink() *string 
  SetResourceOwnerAccount(v string) *EditTaskShrinkRequest
  GetResourceOwnerAccount() *string 
  SetResourceOwnerId(v int64) *EditTaskShrinkRequest
  GetResourceOwnerId() *int64 
  SetSendSmsPlanShrink(v string) *EditTaskShrinkRequest
  GetSendSmsPlanShrink() *string 
  SetStatus(v int64) *EditTaskShrinkRequest
  GetStatus() *int64 
  SetTaskId(v int64) *EditTaskShrinkRequest
  GetTaskId() *int64 
  SetTemplateId(v int64) *EditTaskShrinkRequest
  GetTemplateId() *int64 
  SetTemplateType(v int64) *EditTaskShrinkRequest
  GetTemplateType() *int64 
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
  Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
  RepeatTimesShrink *string `json:"RepeatTimes,omitempty" xml:"RepeatTimes,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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
  return dara.Prettify(s)
}

func (s EditTaskShrinkRequest) GoString() string {
  return s.String()
}

func (s *EditTaskShrinkRequest) GetCallTimeListShrink() *string  {
  return s.CallTimeListShrink
}

func (s *EditTaskShrinkRequest) GetCallbackUrl() *string  {
  return s.CallbackUrl
}

func (s *EditTaskShrinkRequest) GetFlashSmsTemplateId() *int64  {
  return s.FlashSmsTemplateId
}

func (s *EditTaskShrinkRequest) GetFlashSmsType() *int64  {
  return s.FlashSmsType
}

func (s *EditTaskShrinkRequest) GetMaxConcurrency() *int64  {
  return s.MaxConcurrency
}

func (s *EditTaskShrinkRequest) GetName() *string  {
  return s.Name
}

func (s *EditTaskShrinkRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EditTaskShrinkRequest) GetPlaySleepVal() *int64  {
  return s.PlaySleepVal
}

func (s *EditTaskShrinkRequest) GetPlayTimes() *int64  {
  return s.PlayTimes
}

func (s *EditTaskShrinkRequest) GetRecallType() *int64  {
  return s.RecallType
}

func (s *EditTaskShrinkRequest) GetRecordPath() *string  {
  return s.RecordPath
}

func (s *EditTaskShrinkRequest) GetRepeatCount() *int64  {
  return s.RepeatCount
}

func (s *EditTaskShrinkRequest) GetRepeatInterval() *int64  {
  return s.RepeatInterval
}

func (s *EditTaskShrinkRequest) GetRepeatReasonShrink() *string  {
  return s.RepeatReasonShrink
}

func (s *EditTaskShrinkRequest) GetRepeatTimesShrink() *string  {
  return s.RepeatTimesShrink
}

func (s *EditTaskShrinkRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *EditTaskShrinkRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *EditTaskShrinkRequest) GetSendSmsPlanShrink() *string  {
  return s.SendSmsPlanShrink
}

func (s *EditTaskShrinkRequest) GetStatus() *int64  {
  return s.Status
}

func (s *EditTaskShrinkRequest) GetTaskId() *int64  {
  return s.TaskId
}

func (s *EditTaskShrinkRequest) GetTemplateId() *int64  {
  return s.TemplateId
}

func (s *EditTaskShrinkRequest) GetTemplateType() *int64  {
  return s.TemplateType
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

func (s *EditTaskShrinkRequest) Validate() error {
  return dara.Validate(s)
}

