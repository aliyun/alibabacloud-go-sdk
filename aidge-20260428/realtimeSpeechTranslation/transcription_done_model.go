// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iTranscriptionDone interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *TranscriptionDone
  GetType() *string 
  SetEventId(v string) *TranscriptionDone
  GetEventId() *string 
  SetSessionId(v string) *TranscriptionDone
  GetSessionId() *string 
  SetTranscriptionId(v string) *TranscriptionDone
  GetTranscriptionId() *string 
  SetRevision(v int64) *TranscriptionDone
  GetRevision() *int64 
  SetText(v string) *TranscriptionDone
  GetText() *string 
  SetStartMs(v int64) *TranscriptionDone
  GetStartMs() *int64 
  SetEndMs(v int64) *TranscriptionDone
  GetEndMs() *int64 
  SetError(v *TranscriptionDoneError) *TranscriptionDone
  GetError() *TranscriptionDoneError 
}

type TranscriptionDone struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  SessionId *string `json:"session_id,omitempty" xml:"session_id,omitempty"`
  // 与该原文的 TranscriptionTextDelta 共享标识，不关联翻译 segment_id。
  TranscriptionId *string `json:"transcription_id,omitempty" xml:"transcription_id,omitempty"`
  // 同一 transcription_id 内单调递增，完成后不再更新。
  Revision *int64 `json:"revision,omitempty" xml:"revision,omitempty"`
  // 最终原文，可以为空；转写失败时保留已收到的 text 与 stash 合并结果，并附带 error。
  Text *string `json:"text,omitempty" xml:"text,omitempty"`
  // 仅已知原文起始边界时返回。
  StartMs *int64 `json:"start_ms,omitempty" xml:"start_ms,omitempty"`
  // 仅已知原文结束边界时返回。
  EndMs *int64 `json:"end_ms,omitempty" xml:"end_ms,omitempty"`
  // 仅原文转写失败时返回；正常完成时省略。
  Error *TranscriptionDoneError `json:"error,omitempty" xml:"error,omitempty" type:"Struct"`
}

func (s TranscriptionDone) String() string {
  return dara.Prettify(s)
}

func (s TranscriptionDone) GoString() string {
  return s.String()
}

func (s *TranscriptionDone) GetType() *string  {
  return s.Type
}

func (s *TranscriptionDone) GetEventId() *string  {
  return s.EventId
}

func (s *TranscriptionDone) GetSessionId() *string  {
  return s.SessionId
}

func (s *TranscriptionDone) GetTranscriptionId() *string  {
  return s.TranscriptionId
}

func (s *TranscriptionDone) GetRevision() *int64  {
  return s.Revision
}

func (s *TranscriptionDone) GetText() *string  {
  return s.Text
}

func (s *TranscriptionDone) GetStartMs() *int64  {
  return s.StartMs
}

func (s *TranscriptionDone) GetEndMs() *int64  {
  return s.EndMs
}

func (s *TranscriptionDone) GetError() *TranscriptionDoneError  {
  return s.Error
}

func (s *TranscriptionDone) SetType(v string) *TranscriptionDone {
  s.Type = &v
  return s
}

func (s *TranscriptionDone) SetEventId(v string) *TranscriptionDone {
  s.EventId = &v
  return s
}

func (s *TranscriptionDone) SetSessionId(v string) *TranscriptionDone {
  s.SessionId = &v
  return s
}

func (s *TranscriptionDone) SetTranscriptionId(v string) *TranscriptionDone {
  s.TranscriptionId = &v
  return s
}

func (s *TranscriptionDone) SetRevision(v int64) *TranscriptionDone {
  s.Revision = &v
  return s
}

func (s *TranscriptionDone) SetText(v string) *TranscriptionDone {
  s.Text = &v
  return s
}

func (s *TranscriptionDone) SetStartMs(v int64) *TranscriptionDone {
  s.StartMs = &v
  return s
}

func (s *TranscriptionDone) SetEndMs(v int64) *TranscriptionDone {
  s.EndMs = &v
  return s
}

func (s *TranscriptionDone) SetError(v *TranscriptionDoneError) *TranscriptionDone {
  s.Error = v
  return s
}

func (s *TranscriptionDone) Validate() error {
  if s.Error != nil {
    if err := s.Error.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type TranscriptionDoneError struct {
  Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s TranscriptionDoneError) String() string {
  return dara.Prettify(s)
}

func (s TranscriptionDoneError) GoString() string {
  return s.String()
}

func (s *TranscriptionDoneError) GetCode() *string  {
  return s.Code
}

func (s *TranscriptionDoneError) SetCode(v string) *TranscriptionDoneError {
  s.Code = &v
  return s
}

func (s *TranscriptionDoneError) Validate() error {
  return dara.Validate(s)
}

