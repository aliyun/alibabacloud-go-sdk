// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iUsageReported interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *UsageReported
  GetType() *string 
  SetEventId(v string) *UsageReported
  GetEventId() *string 
  SetSessionId(v string) *UsageReported
  GetSessionId() *string 
  SetUsageId(v string) *UsageReported
  GetUsageId() *string 
  SetStatus(v string) *UsageReported
  GetStatus() *string 
  SetUsageMap(v *UsageReportedUsageMap) *UsageReported
  GetUsageMap() *UsageReportedUsageMap 
}

type UsageReported struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  SessionId *string `json:"session_id,omitempty" xml:"session_id,omitempty"`
  // 与计量上报使用相同的稳定标识；不以 event_id 或翻译 segment_id 代替。
  UsageId *string `json:"usage_id,omitempty" xml:"usage_id,omitempty"`
  // 仅在真实 MQ 接收用量后发送；不表示已完成扣费或出账。
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
  // 单次用量记录，不是会话累计值；字段名称及字符串值与计量上报保持一致。
  UsageMap *UsageReportedUsageMap `json:"usageMap,omitempty" xml:"usageMap,omitempty" type:"Struct"`
}

func (s UsageReported) String() string {
  return dara.Prettify(s)
}

func (s UsageReported) GoString() string {
  return s.String()
}

func (s *UsageReported) GetType() *string  {
  return s.Type
}

func (s *UsageReported) GetEventId() *string  {
  return s.EventId
}

func (s *UsageReported) GetSessionId() *string  {
  return s.SessionId
}

func (s *UsageReported) GetUsageId() *string  {
  return s.UsageId
}

func (s *UsageReported) GetStatus() *string  {
  return s.Status
}

func (s *UsageReported) GetUsageMap() *UsageReportedUsageMap  {
  return s.UsageMap
}

func (s *UsageReported) SetType(v string) *UsageReported {
  s.Type = &v
  return s
}

func (s *UsageReported) SetEventId(v string) *UsageReported {
  s.EventId = &v
  return s
}

func (s *UsageReported) SetSessionId(v string) *UsageReported {
  s.SessionId = &v
  return s
}

func (s *UsageReported) SetUsageId(v string) *UsageReported {
  s.UsageId = &v
  return s
}

func (s *UsageReported) SetStatus(v string) *UsageReported {
  s.Status = &v
  return s
}

func (s *UsageReported) SetUsageMap(v *UsageReportedUsageMap) *UsageReported {
  s.UsageMap = v
  return s
}

func (s *UsageReported) Validate() error {
  if s.UsageMap != nil {
    if err := s.UsageMap.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type UsageReportedUsageMap struct {
  // 每段按采样时长独立四舍五入为整数秒，以字符串返回，允许为 "0"。
  InputAudioSeconds *string `json:"InputAudioSeconds,omitempty" xml:"InputAudioSeconds,omitempty"`
  VoiceTranslationType *string `json:"VoiceTranslationType,omitempty" xml:"VoiceTranslationType,omitempty"`
}

func (s UsageReportedUsageMap) String() string {
  return dara.Prettify(s)
}

func (s UsageReportedUsageMap) GoString() string {
  return s.String()
}

func (s *UsageReportedUsageMap) GetInputAudioSeconds() *string  {
  return s.InputAudioSeconds
}

func (s *UsageReportedUsageMap) GetVoiceTranslationType() *string  {
  return s.VoiceTranslationType
}

func (s *UsageReportedUsageMap) SetInputAudioSeconds(v string) *UsageReportedUsageMap {
  s.InputAudioSeconds = &v
  return s
}

func (s *UsageReportedUsageMap) SetVoiceTranslationType(v string) *UsageReportedUsageMap {
  s.VoiceTranslationType = &v
  return s
}

func (s *UsageReportedUsageMap) Validate() error {
  return dara.Validate(s)
}

