// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iRealtimeError interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *RealtimeError
  GetType() *string 
  SetEventId(v string) *RealtimeError
  GetEventId() *string 
  SetSessionId(v string) *RealtimeError
  GetSessionId() *string 
  SetError(v *RealtimeErrorError) *RealtimeError
  GetError() *RealtimeErrorError 
}

type RealtimeError struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  SessionId *string `json:"session_id,omitempty" xml:"session_id,omitempty"`
  Error *RealtimeErrorError `json:"error,omitempty" xml:"error,omitempty" type:"Struct"`
}

func (s RealtimeError) String() string {
  return dara.Prettify(s)
}

func (s RealtimeError) GoString() string {
  return s.String()
}

func (s *RealtimeError) GetType() *string  {
  return s.Type
}

func (s *RealtimeError) GetEventId() *string  {
  return s.EventId
}

func (s *RealtimeError) GetSessionId() *string  {
  return s.SessionId
}

func (s *RealtimeError) GetError() *RealtimeErrorError  {
  return s.Error
}

func (s *RealtimeError) SetType(v string) *RealtimeError {
  s.Type = &v
  return s
}

func (s *RealtimeError) SetEventId(v string) *RealtimeError {
  s.EventId = &v
  return s
}

func (s *RealtimeError) SetSessionId(v string) *RealtimeError {
  s.SessionId = &v
  return s
}

func (s *RealtimeError) SetError(v *RealtimeErrorError) *RealtimeError {
  s.Error = v
  return s
}

func (s *RealtimeError) Validate() error {
  if s.Error != nil {
    if err := s.Error.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type RealtimeErrorError struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  Code *string `json:"code,omitempty" xml:"code,omitempty"`
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  Param *string `json:"param,omitempty" xml:"param,omitempty"`
  Fatal *bool `json:"fatal,omitempty" xml:"fatal,omitempty"`
  Retryable *bool `json:"retryable,omitempty" xml:"retryable,omitempty"`
  CausedBy *string `json:"caused_by,omitempty" xml:"caused_by,omitempty"`
}

func (s RealtimeErrorError) String() string {
  return dara.Prettify(s)
}

func (s RealtimeErrorError) GoString() string {
  return s.String()
}

func (s *RealtimeErrorError) GetType() *string  {
  return s.Type
}

func (s *RealtimeErrorError) GetCode() *string  {
  return s.Code
}

func (s *RealtimeErrorError) GetMessage() *string  {
  return s.Message
}

func (s *RealtimeErrorError) GetParam() *string  {
  return s.Param
}

func (s *RealtimeErrorError) GetFatal() *bool  {
  return s.Fatal
}

func (s *RealtimeErrorError) GetRetryable() *bool  {
  return s.Retryable
}

func (s *RealtimeErrorError) GetCausedBy() *string  {
  return s.CausedBy
}

func (s *RealtimeErrorError) SetType(v string) *RealtimeErrorError {
  s.Type = &v
  return s
}

func (s *RealtimeErrorError) SetCode(v string) *RealtimeErrorError {
  s.Code = &v
  return s
}

func (s *RealtimeErrorError) SetMessage(v string) *RealtimeErrorError {
  s.Message = &v
  return s
}

func (s *RealtimeErrorError) SetParam(v string) *RealtimeErrorError {
  s.Param = &v
  return s
}

func (s *RealtimeErrorError) SetFatal(v bool) *RealtimeErrorError {
  s.Fatal = &v
  return s
}

func (s *RealtimeErrorError) SetRetryable(v bool) *RealtimeErrorError {
  s.Retryable = &v
  return s
}

func (s *RealtimeErrorError) SetCausedBy(v string) *RealtimeErrorError {
  s.CausedBy = &v
  return s
}

func (s *RealtimeErrorError) Validate() error {
  return dara.Validate(s)
}

