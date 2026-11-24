// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iTranscriptionTextDelta interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *TranscriptionTextDelta
  GetType() *string 
  SetEventId(v string) *TranscriptionTextDelta
  GetEventId() *string 
  SetSessionId(v string) *TranscriptionTextDelta
  GetSessionId() *string 
  SetText(v string) *TranscriptionTextDelta
  GetText() *string 
  SetStartMs(v int64) *TranscriptionTextDelta
  GetStartMs() *int64 
  SetEndMs(v int64) *TranscriptionTextDelta
  GetEndMs() *int64 
  SetTranscriptionId(v string) *TranscriptionTextDelta
  GetTranscriptionId() *string 
  SetRevision(v int64) *TranscriptionTextDelta
  GetRevision() *int64 
  SetStash(v string) *TranscriptionTextDelta
  GetStash() *string 
}

type TranscriptionTextDelta struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  SessionId *string `json:"session_id,omitempty" xml:"session_id,omitempty"`
  // 可回改的累计全文，按最新 revision 覆盖，不作增量拼接。
  Text *string `json:"text,omitempty" xml:"text,omitempty"`
  // 仅已知原文起始边界时返回。
  StartMs *int64 `json:"start_ms,omitempty" xml:"start_ms,omitempty"`
  // 仅已知原文结束边界时返回。
  EndMs *int64 `json:"end_ms,omitempty" xml:"end_ms,omitempty"`
  // 同一原文的快照与完成事件共享此标识，不关联翻译 segment_id。
  TranscriptionId *string `json:"transcription_id,omitempty" xml:"transcription_id,omitempty"`
  // 同一 transcription_id 内单调递增；中间快照可能合并，版本号不保证连续。
  Revision *int64 `json:"revision,omitempty" xml:"revision,omitempty"`
  // 可被修正的暂存文本，可为空；仅 stash 变化也会产生新版本。
  Stash *string `json:"stash,omitempty" xml:"stash,omitempty"`
}

func (s TranscriptionTextDelta) String() string {
  return dara.Prettify(s)
}

func (s TranscriptionTextDelta) GoString() string {
  return s.String()
}

func (s *TranscriptionTextDelta) GetType() *string  {
  return s.Type
}

func (s *TranscriptionTextDelta) GetEventId() *string  {
  return s.EventId
}

func (s *TranscriptionTextDelta) GetSessionId() *string  {
  return s.SessionId
}

func (s *TranscriptionTextDelta) GetText() *string  {
  return s.Text
}

func (s *TranscriptionTextDelta) GetStartMs() *int64  {
  return s.StartMs
}

func (s *TranscriptionTextDelta) GetEndMs() *int64  {
  return s.EndMs
}

func (s *TranscriptionTextDelta) GetTranscriptionId() *string  {
  return s.TranscriptionId
}

func (s *TranscriptionTextDelta) GetRevision() *int64  {
  return s.Revision
}

func (s *TranscriptionTextDelta) GetStash() *string  {
  return s.Stash
}

func (s *TranscriptionTextDelta) SetType(v string) *TranscriptionTextDelta {
  s.Type = &v
  return s
}

func (s *TranscriptionTextDelta) SetEventId(v string) *TranscriptionTextDelta {
  s.EventId = &v
  return s
}

func (s *TranscriptionTextDelta) SetSessionId(v string) *TranscriptionTextDelta {
  s.SessionId = &v
  return s
}

func (s *TranscriptionTextDelta) SetText(v string) *TranscriptionTextDelta {
  s.Text = &v
  return s
}

func (s *TranscriptionTextDelta) SetStartMs(v int64) *TranscriptionTextDelta {
  s.StartMs = &v
  return s
}

func (s *TranscriptionTextDelta) SetEndMs(v int64) *TranscriptionTextDelta {
  s.EndMs = &v
  return s
}

func (s *TranscriptionTextDelta) SetTranscriptionId(v string) *TranscriptionTextDelta {
  s.TranscriptionId = &v
  return s
}

func (s *TranscriptionTextDelta) SetRevision(v int64) *TranscriptionTextDelta {
  s.Revision = &v
  return s
}

func (s *TranscriptionTextDelta) SetStash(v string) *TranscriptionTextDelta {
  s.Stash = &v
  return s
}

func (s *TranscriptionTextDelta) Validate() error {
  return dara.Validate(s)
}

