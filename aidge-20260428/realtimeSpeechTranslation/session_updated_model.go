// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iSessionUpdated interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *SessionUpdated
  GetType() *string 
  SetEventId(v string) *SessionUpdated
  GetEventId() *string 
  SetSessionId(v string) *SessionUpdated
  GetSessionId() *string 
  SetCausedBy(v string) *SessionUpdated
  GetCausedBy() *string 
  SetWarnings(v []*string) *SessionUpdated
  GetWarnings() []*string 
  SetSession(v *SessionUpdatedSession) *SessionUpdated
  GetSession() *SessionUpdatedSession 
}

type SessionUpdated struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  SessionId *string `json:"session_id,omitempty" xml:"session_id,omitempty"`
  CausedBy *string `json:"caused_by,omitempty" xml:"caused_by,omitempty"`
  Warnings []*string `json:"warnings,omitempty" xml:"warnings,omitempty" type:"Repeated"`
  Session *SessionUpdatedSession `json:"session,omitempty" xml:"session,omitempty" type:"Struct"`
}

func (s SessionUpdated) String() string {
  return dara.Prettify(s)
}

func (s SessionUpdated) GoString() string {
  return s.String()
}

func (s *SessionUpdated) GetType() *string  {
  return s.Type
}

func (s *SessionUpdated) GetEventId() *string  {
  return s.EventId
}

func (s *SessionUpdated) GetSessionId() *string  {
  return s.SessionId
}

func (s *SessionUpdated) GetCausedBy() *string  {
  return s.CausedBy
}

func (s *SessionUpdated) GetWarnings() []*string  {
  return s.Warnings
}

func (s *SessionUpdated) GetSession() *SessionUpdatedSession  {
  return s.Session
}

func (s *SessionUpdated) SetType(v string) *SessionUpdated {
  s.Type = &v
  return s
}

func (s *SessionUpdated) SetEventId(v string) *SessionUpdated {
  s.EventId = &v
  return s
}

func (s *SessionUpdated) SetSessionId(v string) *SessionUpdated {
  s.SessionId = &v
  return s
}

func (s *SessionUpdated) SetCausedBy(v string) *SessionUpdated {
  s.CausedBy = &v
  return s
}

func (s *SessionUpdated) SetWarnings(v []*string) *SessionUpdated {
  s.Warnings = v
  return s
}

func (s *SessionUpdated) SetSession(v *SessionUpdatedSession) *SessionUpdated {
  s.Session = v
  return s
}

func (s *SessionUpdated) Validate() error {
  if s.Session != nil {
    if err := s.Session.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type SessionUpdatedSession struct {
  Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
  LanguagePair []*string `json:"language_pair,omitempty" xml:"language_pair,omitempty" type:"Repeated"`
  InputAudio *SessionUpdatedSessionInputAudio `json:"input_audio,omitempty" xml:"input_audio,omitempty" type:"Struct"`
  Output *SessionUpdatedSessionOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
  Transcription *SessionUpdatedSessionTranscription `json:"transcription,omitempty" xml:"transcription,omitempty" type:"Struct"`
  TurnDetection *SessionUpdatedSessionTurnDetection `json:"turn_detection,omitempty" xml:"turn_detection,omitempty" type:"Struct"`
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  Object *string `json:"object,omitempty" xml:"object,omitempty"`
  Modalities []*string `json:"modalities,omitempty" xml:"modalities,omitempty" type:"Repeated"`
  Voice *string `json:"voice,omitempty" xml:"voice,omitempty"`
  InputAudioFormat *string `json:"input_audio_format,omitempty" xml:"input_audio_format,omitempty"`
  SampleRate *int32 `json:"sample_rate,omitempty" xml:"sample_rate,omitempty"`
  OutputAudioFormat *string `json:"output_audio_format,omitempty" xml:"output_audio_format,omitempty"`
  Translation *SessionUpdatedSessionTranslation `json:"translation,omitempty" xml:"translation,omitempty" type:"Struct"`
  EnableVoiceClone *bool `json:"enable_voice_clone,omitempty" xml:"enable_voice_clone,omitempty"`
  VoiceCloneOptions *SessionUpdatedSessionVoiceCloneOptions `json:"voice_clone_options,omitempty" xml:"voice_clone_options,omitempty" type:"Struct"`
  // 单向回执固定包含 enabled、language；双向使用 transcription 字段。内部转写 model 不对客回显。
  InputAudioTranscription *SessionUpdatedSessionInputAudioTranscription `json:"input_audio_transcription,omitempty" xml:"input_audio_transcription,omitempty" type:"Struct"`
}

func (s SessionUpdatedSession) String() string {
  return dara.Prettify(s)
}

func (s SessionUpdatedSession) GoString() string {
  return s.String()
}

func (s *SessionUpdatedSession) GetMode() *string  {
  return s.Mode
}

func (s *SessionUpdatedSession) GetLanguagePair() []*string  {
  return s.LanguagePair
}

func (s *SessionUpdatedSession) GetInputAudio() *SessionUpdatedSessionInputAudio  {
  return s.InputAudio
}

func (s *SessionUpdatedSession) GetOutput() *SessionUpdatedSessionOutput  {
  return s.Output
}

func (s *SessionUpdatedSession) GetTranscription() *SessionUpdatedSessionTranscription  {
  return s.Transcription
}

func (s *SessionUpdatedSession) GetTurnDetection() *SessionUpdatedSessionTurnDetection  {
  return s.TurnDetection
}

func (s *SessionUpdatedSession) GetId() *string  {
  return s.Id
}

func (s *SessionUpdatedSession) GetObject() *string  {
  return s.Object
}

func (s *SessionUpdatedSession) GetModalities() []*string  {
  return s.Modalities
}

func (s *SessionUpdatedSession) GetVoice() *string  {
  return s.Voice
}

func (s *SessionUpdatedSession) GetInputAudioFormat() *string  {
  return s.InputAudioFormat
}

func (s *SessionUpdatedSession) GetSampleRate() *int32  {
  return s.SampleRate
}

func (s *SessionUpdatedSession) GetOutputAudioFormat() *string  {
  return s.OutputAudioFormat
}

func (s *SessionUpdatedSession) GetTranslation() *SessionUpdatedSessionTranslation  {
  return s.Translation
}

func (s *SessionUpdatedSession) GetEnableVoiceClone() *bool  {
  return s.EnableVoiceClone
}

func (s *SessionUpdatedSession) GetVoiceCloneOptions() *SessionUpdatedSessionVoiceCloneOptions  {
  return s.VoiceCloneOptions
}

func (s *SessionUpdatedSession) GetInputAudioTranscription() *SessionUpdatedSessionInputAudioTranscription  {
  return s.InputAudioTranscription
}

func (s *SessionUpdatedSession) SetMode(v string) *SessionUpdatedSession {
  s.Mode = &v
  return s
}

func (s *SessionUpdatedSession) SetLanguagePair(v []*string) *SessionUpdatedSession {
  s.LanguagePair = v
  return s
}

func (s *SessionUpdatedSession) SetInputAudio(v *SessionUpdatedSessionInputAudio) *SessionUpdatedSession {
  s.InputAudio = v
  return s
}

func (s *SessionUpdatedSession) SetOutput(v *SessionUpdatedSessionOutput) *SessionUpdatedSession {
  s.Output = v
  return s
}

func (s *SessionUpdatedSession) SetTranscription(v *SessionUpdatedSessionTranscription) *SessionUpdatedSession {
  s.Transcription = v
  return s
}

func (s *SessionUpdatedSession) SetTurnDetection(v *SessionUpdatedSessionTurnDetection) *SessionUpdatedSession {
  s.TurnDetection = v
  return s
}

func (s *SessionUpdatedSession) SetId(v string) *SessionUpdatedSession {
  s.Id = &v
  return s
}

func (s *SessionUpdatedSession) SetObject(v string) *SessionUpdatedSession {
  s.Object = &v
  return s
}

func (s *SessionUpdatedSession) SetModalities(v []*string) *SessionUpdatedSession {
  s.Modalities = v
  return s
}

func (s *SessionUpdatedSession) SetVoice(v string) *SessionUpdatedSession {
  s.Voice = &v
  return s
}

func (s *SessionUpdatedSession) SetInputAudioFormat(v string) *SessionUpdatedSession {
  s.InputAudioFormat = &v
  return s
}

func (s *SessionUpdatedSession) SetSampleRate(v int32) *SessionUpdatedSession {
  s.SampleRate = &v
  return s
}

func (s *SessionUpdatedSession) SetOutputAudioFormat(v string) *SessionUpdatedSession {
  s.OutputAudioFormat = &v
  return s
}

func (s *SessionUpdatedSession) SetTranslation(v *SessionUpdatedSessionTranslation) *SessionUpdatedSession {
  s.Translation = v
  return s
}

func (s *SessionUpdatedSession) SetEnableVoiceClone(v bool) *SessionUpdatedSession {
  s.EnableVoiceClone = &v
  return s
}

func (s *SessionUpdatedSession) SetVoiceCloneOptions(v *SessionUpdatedSessionVoiceCloneOptions) *SessionUpdatedSession {
  s.VoiceCloneOptions = v
  return s
}

func (s *SessionUpdatedSession) SetInputAudioTranscription(v *SessionUpdatedSessionInputAudioTranscription) *SessionUpdatedSession {
  s.InputAudioTranscription = v
  return s
}

func (s *SessionUpdatedSession) Validate() error {
  if s.InputAudio != nil {
    if err := s.InputAudio.Validate(); err != nil {
      return err
    }
  }
  if s.Output != nil {
    if err := s.Output.Validate(); err != nil {
      return err
    }
  }
  if s.Transcription != nil {
    if err := s.Transcription.Validate(); err != nil {
      return err
    }
  }
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
  if s.VoiceCloneOptions != nil {
    if err := s.VoiceCloneOptions.Validate(); err != nil {
      return err
    }
  }
  if s.InputAudioTranscription != nil {
    if err := s.InputAudioTranscription.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type SessionUpdatedSessionInputAudio struct {
  Format *string `json:"format,omitempty" xml:"format,omitempty"`
  SampleRate *int32 `json:"sample_rate,omitempty" xml:"sample_rate,omitempty"`
  Channels *int32 `json:"channels,omitempty" xml:"channels,omitempty"`
}

func (s SessionUpdatedSessionInputAudio) String() string {
  return dara.Prettify(s)
}

func (s SessionUpdatedSessionInputAudio) GoString() string {
  return s.String()
}

func (s *SessionUpdatedSessionInputAudio) GetFormat() *string  {
  return s.Format
}

func (s *SessionUpdatedSessionInputAudio) GetSampleRate() *int32  {
  return s.SampleRate
}

func (s *SessionUpdatedSessionInputAudio) GetChannels() *int32  {
  return s.Channels
}

func (s *SessionUpdatedSessionInputAudio) SetFormat(v string) *SessionUpdatedSessionInputAudio {
  s.Format = &v
  return s
}

func (s *SessionUpdatedSessionInputAudio) SetSampleRate(v int32) *SessionUpdatedSessionInputAudio {
  s.SampleRate = &v
  return s
}

func (s *SessionUpdatedSessionInputAudio) SetChannels(v int32) *SessionUpdatedSessionInputAudio {
  s.Channels = &v
  return s
}

func (s *SessionUpdatedSessionInputAudio) Validate() error {
  return dara.Validate(s)
}

type SessionUpdatedSessionOutput struct {
  Modalities []*string `json:"modalities,omitempty" xml:"modalities,omitempty" type:"Repeated"`
  Audio *SessionUpdatedSessionOutputAudio `json:"audio,omitempty" xml:"audio,omitempty" type:"Struct"`
}

func (s SessionUpdatedSessionOutput) String() string {
  return dara.Prettify(s)
}

func (s SessionUpdatedSessionOutput) GoString() string {
  return s.String()
}

func (s *SessionUpdatedSessionOutput) GetModalities() []*string  {
  return s.Modalities
}

func (s *SessionUpdatedSessionOutput) GetAudio() *SessionUpdatedSessionOutputAudio  {
  return s.Audio
}

func (s *SessionUpdatedSessionOutput) SetModalities(v []*string) *SessionUpdatedSessionOutput {
  s.Modalities = v
  return s
}

func (s *SessionUpdatedSessionOutput) SetAudio(v *SessionUpdatedSessionOutputAudio) *SessionUpdatedSessionOutput {
  s.Audio = v
  return s
}

func (s *SessionUpdatedSessionOutput) Validate() error {
  if s.Audio != nil {
    if err := s.Audio.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type SessionUpdatedSessionOutputAudio struct {
  Format *string `json:"format,omitempty" xml:"format,omitempty"`
  SampleRate *int32 `json:"sample_rate,omitempty" xml:"sample_rate,omitempty"`
  Channels *int32 `json:"channels,omitempty" xml:"channels,omitempty"`
  Voice *SessionUpdatedSessionOutputAudioVoice `json:"voice,omitempty" xml:"voice,omitempty" type:"Struct"`
}

func (s SessionUpdatedSessionOutputAudio) String() string {
  return dara.Prettify(s)
}

func (s SessionUpdatedSessionOutputAudio) GoString() string {
  return s.String()
}

func (s *SessionUpdatedSessionOutputAudio) GetFormat() *string  {
  return s.Format
}

func (s *SessionUpdatedSessionOutputAudio) GetSampleRate() *int32  {
  return s.SampleRate
}

func (s *SessionUpdatedSessionOutputAudio) GetChannels() *int32  {
  return s.Channels
}

func (s *SessionUpdatedSessionOutputAudio) GetVoice() *SessionUpdatedSessionOutputAudioVoice  {
  return s.Voice
}

func (s *SessionUpdatedSessionOutputAudio) SetFormat(v string) *SessionUpdatedSessionOutputAudio {
  s.Format = &v
  return s
}

func (s *SessionUpdatedSessionOutputAudio) SetSampleRate(v int32) *SessionUpdatedSessionOutputAudio {
  s.SampleRate = &v
  return s
}

func (s *SessionUpdatedSessionOutputAudio) SetChannels(v int32) *SessionUpdatedSessionOutputAudio {
  s.Channels = &v
  return s
}

func (s *SessionUpdatedSessionOutputAudio) SetVoice(v *SessionUpdatedSessionOutputAudioVoice) *SessionUpdatedSessionOutputAudio {
  s.Voice = v
  return s
}

func (s *SessionUpdatedSessionOutputAudio) Validate() error {
  if s.Voice != nil {
    if err := s.Voice.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type SessionUpdatedSessionOutputAudioVoice struct {
  Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s SessionUpdatedSessionOutputAudioVoice) String() string {
  return dara.Prettify(s)
}

func (s SessionUpdatedSessionOutputAudioVoice) GoString() string {
  return s.String()
}

func (s *SessionUpdatedSessionOutputAudioVoice) GetMode() *string  {
  return s.Mode
}

func (s *SessionUpdatedSessionOutputAudioVoice) GetId() *string  {
  return s.Id
}

func (s *SessionUpdatedSessionOutputAudioVoice) SetMode(v string) *SessionUpdatedSessionOutputAudioVoice {
  s.Mode = &v
  return s
}

func (s *SessionUpdatedSessionOutputAudioVoice) SetId(v string) *SessionUpdatedSessionOutputAudioVoice {
  s.Id = &v
  return s
}

func (s *SessionUpdatedSessionOutputAudioVoice) Validate() error {
  return dara.Validate(s)
}

type SessionUpdatedSessionTranscription struct {
  Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s SessionUpdatedSessionTranscription) String() string {
  return dara.Prettify(s)
}

func (s SessionUpdatedSessionTranscription) GoString() string {
  return s.String()
}

func (s *SessionUpdatedSessionTranscription) GetEnabled() *bool  {
  return s.Enabled
}

func (s *SessionUpdatedSessionTranscription) SetEnabled(v bool) *SessionUpdatedSessionTranscription {
  s.Enabled = &v
  return s
}

func (s *SessionUpdatedSessionTranscription) Validate() error {
  return dara.Validate(s)
}

type SessionUpdatedSessionTurnDetection struct {
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  Threshold *float32 `json:"threshold,omitempty" xml:"threshold,omitempty"`
  SilenceDurationMs *int32 `json:"silence_duration_ms,omitempty" xml:"silence_duration_ms,omitempty"`
  PrefixPaddingMs *int32 `json:"prefix_padding_ms,omitempty" xml:"prefix_padding_ms,omitempty"`
  CreateResponse *bool `json:"create_response,omitempty" xml:"create_response,omitempty"`
  InterruptResponse *bool `json:"interrupt_response,omitempty" xml:"interrupt_response,omitempty"`
}

func (s SessionUpdatedSessionTurnDetection) String() string {
  return dara.Prettify(s)
}

func (s SessionUpdatedSessionTurnDetection) GoString() string {
  return s.String()
}

func (s *SessionUpdatedSessionTurnDetection) GetType() *string  {
  return s.Type
}

func (s *SessionUpdatedSessionTurnDetection) GetThreshold() *float32  {
  return s.Threshold
}

func (s *SessionUpdatedSessionTurnDetection) GetSilenceDurationMs() *int32  {
  return s.SilenceDurationMs
}

func (s *SessionUpdatedSessionTurnDetection) GetPrefixPaddingMs() *int32  {
  return s.PrefixPaddingMs
}

func (s *SessionUpdatedSessionTurnDetection) GetCreateResponse() *bool  {
  return s.CreateResponse
}

func (s *SessionUpdatedSessionTurnDetection) GetInterruptResponse() *bool  {
  return s.InterruptResponse
}

func (s *SessionUpdatedSessionTurnDetection) SetType(v string) *SessionUpdatedSessionTurnDetection {
  s.Type = &v
  return s
}

func (s *SessionUpdatedSessionTurnDetection) SetThreshold(v float32) *SessionUpdatedSessionTurnDetection {
  s.Threshold = &v
  return s
}

func (s *SessionUpdatedSessionTurnDetection) SetSilenceDurationMs(v int32) *SessionUpdatedSessionTurnDetection {
  s.SilenceDurationMs = &v
  return s
}

func (s *SessionUpdatedSessionTurnDetection) SetPrefixPaddingMs(v int32) *SessionUpdatedSessionTurnDetection {
  s.PrefixPaddingMs = &v
  return s
}

func (s *SessionUpdatedSessionTurnDetection) SetCreateResponse(v bool) *SessionUpdatedSessionTurnDetection {
  s.CreateResponse = &v
  return s
}

func (s *SessionUpdatedSessionTurnDetection) SetInterruptResponse(v bool) *SessionUpdatedSessionTurnDetection {
  s.InterruptResponse = &v
  return s
}

func (s *SessionUpdatedSessionTurnDetection) Validate() error {
  return dara.Validate(s)
}

type SessionUpdatedSessionTranslation struct {
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  Corpus *SessionUpdatedSessionTranslationCorpus `json:"corpus,omitempty" xml:"corpus,omitempty" type:"Struct"`
}

func (s SessionUpdatedSessionTranslation) String() string {
  return dara.Prettify(s)
}

func (s SessionUpdatedSessionTranslation) GoString() string {
  return s.String()
}

func (s *SessionUpdatedSessionTranslation) GetLanguage() *string  {
  return s.Language
}

func (s *SessionUpdatedSessionTranslation) GetCorpus() *SessionUpdatedSessionTranslationCorpus  {
  return s.Corpus
}

func (s *SessionUpdatedSessionTranslation) SetLanguage(v string) *SessionUpdatedSessionTranslation {
  s.Language = &v
  return s
}

func (s *SessionUpdatedSessionTranslation) SetCorpus(v *SessionUpdatedSessionTranslationCorpus) *SessionUpdatedSessionTranslation {
  s.Corpus = v
  return s
}

func (s *SessionUpdatedSessionTranslation) Validate() error {
  if s.Corpus != nil {
    if err := s.Corpus.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type SessionUpdatedSessionTranslationCorpus struct {
  Phrases map[string]*string `json:"phrases,omitempty" xml:"phrases,omitempty"`
}

func (s SessionUpdatedSessionTranslationCorpus) String() string {
  return dara.Prettify(s)
}

func (s SessionUpdatedSessionTranslationCorpus) GoString() string {
  return s.String()
}

func (s *SessionUpdatedSessionTranslationCorpus) GetPhrases() map[string]*string  {
  return s.Phrases
}

func (s *SessionUpdatedSessionTranslationCorpus) SetPhrases(v map[string]*string) *SessionUpdatedSessionTranslationCorpus {
  s.Phrases = v
  return s
}

func (s *SessionUpdatedSessionTranslationCorpus) Validate() error {
  return dara.Validate(s)
}

type SessionUpdatedSessionVoiceCloneOptions struct {
  Frequency *string `json:"frequency,omitempty" xml:"frequency,omitempty"`
}

func (s SessionUpdatedSessionVoiceCloneOptions) String() string {
  return dara.Prettify(s)
}

func (s SessionUpdatedSessionVoiceCloneOptions) GoString() string {
  return s.String()
}

func (s *SessionUpdatedSessionVoiceCloneOptions) GetFrequency() *string  {
  return s.Frequency
}

func (s *SessionUpdatedSessionVoiceCloneOptions) SetFrequency(v string) *SessionUpdatedSessionVoiceCloneOptions {
  s.Frequency = &v
  return s
}

func (s *SessionUpdatedSessionVoiceCloneOptions) Validate() error {
  return dara.Validate(s)
}

type SessionUpdatedSessionInputAudioTranscription struct {
  // This parameter is required.
  Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
  // 开启时回显指定语种，自动识别回显 auto；关闭时统一回显 auto。
  // 
  // This parameter is required.
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s SessionUpdatedSessionInputAudioTranscription) String() string {
  return dara.Prettify(s)
}

func (s SessionUpdatedSessionInputAudioTranscription) GoString() string {
  return s.String()
}

func (s *SessionUpdatedSessionInputAudioTranscription) GetEnabled() *bool  {
  return s.Enabled
}

func (s *SessionUpdatedSessionInputAudioTranscription) GetLanguage() *string  {
  return s.Language
}

func (s *SessionUpdatedSessionInputAudioTranscription) SetEnabled(v bool) *SessionUpdatedSessionInputAudioTranscription {
  s.Enabled = &v
  return s
}

func (s *SessionUpdatedSessionInputAudioTranscription) SetLanguage(v string) *SessionUpdatedSessionInputAudioTranscription {
  s.Language = &v
  return s
}

func (s *SessionUpdatedSessionInputAudioTranscription) Validate() error {
  return dara.Validate(s)
}

