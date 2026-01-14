// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditTaskRequest interface {
  dara.Model
  String() string
  GoString() string
  SetCallTimeList(v []*EditTaskRequestCallTimeList) *EditTaskRequest
  GetCallTimeList() []*EditTaskRequestCallTimeList 
  SetCallbackUrl(v string) *EditTaskRequest
  GetCallbackUrl() *string 
  SetFlashSmsTemplateId(v int64) *EditTaskRequest
  GetFlashSmsTemplateId() *int64 
  SetFlashSmsType(v int64) *EditTaskRequest
  GetFlashSmsType() *int64 
  SetMaxConcurrency(v int64) *EditTaskRequest
  GetMaxConcurrency() *int64 
  SetName(v string) *EditTaskRequest
  GetName() *string 
  SetOwnerId(v int64) *EditTaskRequest
  GetOwnerId() *int64 
  SetPlaySleepVal(v int64) *EditTaskRequest
  GetPlaySleepVal() *int64 
  SetPlayTimes(v int64) *EditTaskRequest
  GetPlayTimes() *int64 
  SetRecallType(v int64) *EditTaskRequest
  GetRecallType() *int64 
  SetRecordPath(v string) *EditTaskRequest
  GetRecordPath() *string 
  SetRepeatCount(v int64) *EditTaskRequest
  GetRepeatCount() *int64 
  SetRepeatInterval(v int64) *EditTaskRequest
  GetRepeatInterval() *int64 
  SetRepeatReason(v []*int64) *EditTaskRequest
  GetRepeatReason() []*int64 
  SetRepeatTimes(v []*string) *EditTaskRequest
  GetRepeatTimes() []*string 
  SetResourceOwnerAccount(v string) *EditTaskRequest
  GetResourceOwnerAccount() *string 
  SetResourceOwnerId(v int64) *EditTaskRequest
  GetResourceOwnerId() *int64 
  SetSendSmsPlan(v []*EditTaskRequestSendSmsPlan) *EditTaskRequest
  GetSendSmsPlan() []*EditTaskRequestSendSmsPlan 
  SetStatus(v int64) *EditTaskRequest
  GetStatus() *int64 
  SetTaskId(v int64) *EditTaskRequest
  GetTaskId() *int64 
  SetTemplateId(v int64) *EditTaskRequest
  GetTemplateId() *int64 
  SetTemplateType(v int64) *EditTaskRequest
  GetTemplateType() *int64 
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
  RepeatReason []*int64 `json:"RepeatReason,omitempty" xml:"RepeatReason,omitempty" type:"Repeated"`
  // 重呼时间
  RepeatTimes []*string `json:"RepeatTimes,omitempty" xml:"RepeatTimes,omitempty" type:"Repeated"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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
  return dara.Prettify(s)
}

func (s EditTaskRequest) GoString() string {
  return s.String()
}

func (s *EditTaskRequest) GetCallTimeList() []*EditTaskRequestCallTimeList  {
  return s.CallTimeList
}

func (s *EditTaskRequest) GetCallbackUrl() *string  {
  return s.CallbackUrl
}

func (s *EditTaskRequest) GetFlashSmsTemplateId() *int64  {
  return s.FlashSmsTemplateId
}

func (s *EditTaskRequest) GetFlashSmsType() *int64  {
  return s.FlashSmsType
}

func (s *EditTaskRequest) GetMaxConcurrency() *int64  {
  return s.MaxConcurrency
}

func (s *EditTaskRequest) GetName() *string  {
  return s.Name
}

func (s *EditTaskRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EditTaskRequest) GetPlaySleepVal() *int64  {
  return s.PlaySleepVal
}

func (s *EditTaskRequest) GetPlayTimes() *int64  {
  return s.PlayTimes
}

func (s *EditTaskRequest) GetRecallType() *int64  {
  return s.RecallType
}

func (s *EditTaskRequest) GetRecordPath() *string  {
  return s.RecordPath
}

func (s *EditTaskRequest) GetRepeatCount() *int64  {
  return s.RepeatCount
}

func (s *EditTaskRequest) GetRepeatInterval() *int64  {
  return s.RepeatInterval
}

func (s *EditTaskRequest) GetRepeatReason() []*int64  {
  return s.RepeatReason
}

func (s *EditTaskRequest) GetRepeatTimes() []*string  {
  return s.RepeatTimes
}

func (s *EditTaskRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *EditTaskRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *EditTaskRequest) GetSendSmsPlan() []*EditTaskRequestSendSmsPlan  {
  return s.SendSmsPlan
}

func (s *EditTaskRequest) GetStatus() *int64  {
  return s.Status
}

func (s *EditTaskRequest) GetTaskId() *int64  {
  return s.TaskId
}

func (s *EditTaskRequest) GetTemplateId() *int64  {
  return s.TemplateId
}

func (s *EditTaskRequest) GetTemplateType() *int64  {
  return s.TemplateType
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

func (s *EditTaskRequest) Validate() error {
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

type EditTaskRequestCallTimeList struct {
  CallTime []*string `json:"CallTime,omitempty" xml:"CallTime,omitempty" type:"Repeated"`
}

func (s EditTaskRequestCallTimeList) String() string {
  return dara.Prettify(s)
}

func (s EditTaskRequestCallTimeList) GoString() string {
  return s.String()
}

func (s *EditTaskRequestCallTimeList) GetCallTime() []*string  {
  return s.CallTime
}

func (s *EditTaskRequestCallTimeList) SetCallTime(v []*string) *EditTaskRequestCallTimeList {
  s.CallTime = v
  return s
}

func (s *EditTaskRequestCallTimeList) Validate() error {
  return dara.Validate(s)
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
  return dara.Prettify(s)
}

func (s EditTaskRequestSendSmsPlan) GoString() string {
  return s.String()
}

func (s *EditTaskRequestSendSmsPlan) GetIntentTags() []*string  {
  return s.IntentTags
}

func (s *EditTaskRequestSendSmsPlan) GetSmsTemplateId() *int64  {
  return s.SmsTemplateId
}

func (s *EditTaskRequestSendSmsPlan) SetIntentTags(v []*string) *EditTaskRequestSendSmsPlan {
  s.IntentTags = v
  return s
}

func (s *EditTaskRequestSendSmsPlan) SetSmsTemplateId(v int64) *EditTaskRequestSendSmsPlan {
  s.SmsTemplateId = &v
  return s
}

func (s *EditTaskRequestSendSmsPlan) Validate() error {
  return dara.Validate(s)
}

