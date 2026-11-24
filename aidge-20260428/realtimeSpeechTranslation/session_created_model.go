// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iSessionCreated interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *SessionCreated
  GetType() *string 
  SetEventId(v string) *SessionCreated
  GetEventId() *string 
  SetSessionId(v string) *SessionCreated
  GetSessionId() *string 
  SetCausedBy(v string) *SessionCreated
  GetCausedBy() *string 
  SetCreatedAt(v int64) *SessionCreated
  GetCreatedAt() *int64 
  SetExpiresAt(v int64) *SessionCreated
  GetExpiresAt() *int64 
  SetSession(v *SessionCreatedSession) *SessionCreated
  GetSession() *SessionCreatedSession 
}

type SessionCreated struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  SessionId *string `json:"session_id,omitempty" xml:"session_id,omitempty"`
  CausedBy *string `json:"caused_by,omitempty" xml:"caused_by,omitempty"`
  CreatedAt *int64 `json:"created_at,omitempty" xml:"created_at,omitempty"`
  ExpiresAt *int64 `json:"expires_at,omitempty" xml:"expires_at,omitempty"`
  Session *SessionCreatedSession `json:"session,omitempty" xml:"session,omitempty" type:"Struct"`
}

func (s SessionCreated) String() string {
  return dara.Prettify(s)
}

func (s SessionCreated) GoString() string {
  return s.String()
}

func (s *SessionCreated) GetType() *string  {
  return s.Type
}

func (s *SessionCreated) GetEventId() *string  {
  return s.EventId
}

func (s *SessionCreated) GetSessionId() *string  {
  return s.SessionId
}

func (s *SessionCreated) GetCausedBy() *string  {
  return s.CausedBy
}

func (s *SessionCreated) GetCreatedAt() *int64  {
  return s.CreatedAt
}

func (s *SessionCreated) GetExpiresAt() *int64  {
  return s.ExpiresAt
}

func (s *SessionCreated) GetSession() *SessionCreatedSession  {
  return s.Session
}

func (s *SessionCreated) SetType(v string) *SessionCreated {
  s.Type = &v
  return s
}

func (s *SessionCreated) SetEventId(v string) *SessionCreated {
  s.EventId = &v
  return s
}

func (s *SessionCreated) SetSessionId(v string) *SessionCreated {
  s.SessionId = &v
  return s
}

func (s *SessionCreated) SetCausedBy(v string) *SessionCreated {
  s.CausedBy = &v
  return s
}

func (s *SessionCreated) SetCreatedAt(v int64) *SessionCreated {
  s.CreatedAt = &v
  return s
}

func (s *SessionCreated) SetExpiresAt(v int64) *SessionCreated {
  s.ExpiresAt = &v
  return s
}

func (s *SessionCreated) SetSession(v *SessionCreatedSession) *SessionCreated {
  s.Session = v
  return s
}

func (s *SessionCreated) Validate() error {
  if s.Session != nil {
    if err := s.Session.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type SessionCreatedSession struct {
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  Object *string `json:"object,omitempty" xml:"object,omitempty"`
  Modalities []*string `json:"modalities,omitempty" xml:"modalities,omitempty" type:"Repeated"`
  Voice *string `json:"voice,omitempty" xml:"voice,omitempty"`
  InputAudioFormat *string `json:"input_audio_format,omitempty" xml:"input_audio_format,omitempty"`
  SampleRate *int32 `json:"sample_rate,omitempty" xml:"sample_rate,omitempty"`
  OutputAudioFormat *string `json:"output_audio_format,omitempty" xml:"output_audio_format,omitempty"`
  TurnDetection *SessionCreatedSessionTurnDetection `json:"turn_detection,omitempty" xml:"turn_detection,omitempty" type:"Struct"`
  Translation *SessionCreatedSessionTranslation `json:"translation,omitempty" xml:"translation,omitempty" type:"Struct"`
}

func (s SessionCreatedSession) String() string {
  return dara.Prettify(s)
}

func (s SessionCreatedSession) GoString() string {
  return s.String()
}

func (s *SessionCreatedSession) GetId() *string  {
  return s.Id
}

func (s *SessionCreatedSession) GetObject() *string  {
  return s.Object
}

func (s *SessionCreatedSession) GetModalities() []*string  {
  return s.Modalities
}

func (s *SessionCreatedSession) GetVoice() *string  {
  return s.Voice
}

func (s *SessionCreatedSession) GetInputAudioFormat() *string  {
  return s.InputAudioFormat
}

func (s *SessionCreatedSession) GetSampleRate() *int32  {
  return s.SampleRate
}

func (s *SessionCreatedSession) GetOutputAudioFormat() *string  {
  return s.OutputAudioFormat
}

func (s *SessionCreatedSession) GetTurnDetection() *SessionCreatedSessionTurnDetection  {
  return s.TurnDetection
}

func (s *SessionCreatedSession) GetTranslation() *SessionCreatedSessionTranslation  {
  return s.Translation
}

func (s *SessionCreatedSession) SetId(v string) *SessionCreatedSession {
  s.Id = &v
  return s
}

func (s *SessionCreatedSession) SetObject(v string) *SessionCreatedSession {
  s.Object = &v
  return s
}

func (s *SessionCreatedSession) SetModalities(v []*string) *SessionCreatedSession {
  s.Modalities = v
  return s
}

func (s *SessionCreatedSession) SetVoice(v string) *SessionCreatedSession {
  s.Voice = &v
  return s
}

func (s *SessionCreatedSession) SetInputAudioFormat(v string) *SessionCreatedSession {
  s.InputAudioFormat = &v
  return s
}

func (s *SessionCreatedSession) SetSampleRate(v int32) *SessionCreatedSession {
  s.SampleRate = &v
  return s
}

func (s *SessionCreatedSession) SetOutputAudioFormat(v string) *SessionCreatedSession {
  s.OutputAudioFormat = &v
  return s
}

func (s *SessionCreatedSession) SetTurnDetection(v *SessionCreatedSessionTurnDetection) *SessionCreatedSession {
  s.TurnDetection = v
  return s
}

func (s *SessionCreatedSession) SetTranslation(v *SessionCreatedSessionTranslation) *SessionCreatedSession {
  s.Translation = v
  return s
}

func (s *SessionCreatedSession) Validate() error {
  if s.TurnDetection != nil {
    if err := s.TurnDetection.Validate(); err != nil {
      return err
    }
  }
  if s.Translation != nil {
    if err := s.Translation.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type SessionCreatedSessionTurnDetection struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  Threshold *float32 `json:"threshold,omitempty" xml:"threshold,omitempty"`
  SilenceDurationMs *int32 `json:"silence_duration_ms,omitempty" xml:"silence_duration_ms,omitempty"`
}

func (s SessionCreatedSessionTurnDetection) String() string {
  return dara.Prettify(s)
}

func (s SessionCreatedSessionTurnDetection) GoString() string {
  return s.String()
}

func (s *SessionCreatedSessionTurnDetection) GetType() *string  {
  return s.Type
}

func (s *SessionCreatedSessionTurnDetection) GetThreshold() *float32  {
  return s.Threshold
}

func (s *SessionCreatedSessionTurnDetection) GetSilenceDurationMs() *int32  {
  return s.SilenceDurationMs
}

func (s *SessionCreatedSessionTurnDetection) SetType(v string) *SessionCreatedSessionTurnDetection {
  s.Type = &v
  return s
}

func (s *SessionCreatedSessionTurnDetection) SetThreshold(v float32) *SessionCreatedSessionTurnDetection {
  s.Threshold = &v
  return s
}

func (s *SessionCreatedSessionTurnDetection) SetSilenceDurationMs(v int32) *SessionCreatedSessionTurnDetection {
  s.SilenceDurationMs = &v
  return s
}

func (s *SessionCreatedSessionTurnDetection) Validate() error {
  return dara.Validate(s)
}

type SessionCreatedSessionTranslation struct {
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  Corpus *SessionCreatedSessionTranslationCorpus `json:"corpus,omitempty" xml:"corpus,omitempty" type:"Struct"`
}

func (s SessionCreatedSessionTranslation) String() string {
  return dara.Prettify(s)
}

func (s SessionCreatedSessionTranslation) GoString() string {
  return s.String()
}

func (s *SessionCreatedSessionTranslation) GetLanguage() *string  {
  return s.Language
}

func (s *SessionCreatedSessionTranslation) GetCorpus() *SessionCreatedSessionTranslationCorpus  {
  return s.Corpus
}

func (s *SessionCreatedSessionTranslation) SetLanguage(v string) *SessionCreatedSessionTranslation {
  s.Language = &v
  return s
}

func (s *SessionCreatedSessionTranslation) SetCorpus(v *SessionCreatedSessionTranslationCorpus) *SessionCreatedSessionTranslation {
  s.Corpus = v
  return s
}

func (s *SessionCreatedSessionTranslation) Validate() error {
  if s.Corpus != nil {
    if err := s.Corpus.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type SessionCreatedSessionTranslationCorpus struct {
  Phrases map[string]*string `json:"phrases,omitempty" xml:"phrases,omitempty"`
}

func (s SessionCreatedSessionTranslationCorpus) String() string {
  return dara.Prettify(s)
}

func (s SessionCreatedSessionTranslationCorpus) GoString() string {
  return s.String()
}

func (s *SessionCreatedSessionTranslationCorpus) GetPhrases() map[string]*string  {
  return s.Phrases
}

func (s *SessionCreatedSessionTranslationCorpus) SetPhrases(v map[string]*string) *SessionCreatedSessionTranslationCorpus {
  s.Phrases = v
  return s
}

func (s *SessionCreatedSessionTranslationCorpus) Validate() error {
  return dara.Validate(s)
}

