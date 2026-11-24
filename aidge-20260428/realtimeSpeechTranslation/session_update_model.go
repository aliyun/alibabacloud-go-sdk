// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iSessionUpdate interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *SessionUpdate
  GetType() *string 
  SetEventId(v string) *SessionUpdate
  GetEventId() *string 
  SetSession(v *SessionUpdateSession) *SessionUpdate
  GetSession() *SessionUpdateSession 
}

type SessionUpdate struct {
  // This parameter is required.
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // 可省略；提供时必须为非空、非纯空白字符串。
  // 
  // example:
  // 
  // evt_update_1
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  // 仅填写当前模式适用字段。双向必须提供 session 对象，仅允许输入前更新；单向按 Qwen 3.5 原生契约可省略 session。
  // 
  // 双向不支持 glossary、声音复刻及单向平铺配置，模式、语言对、条件必填及能力校验由后端执行。网关不补默认值。
  // 
  // 单向手动提交、关闭转写：{"modalities":["text"],"sample_rate":16000,"input_audio_format":"pcm","input_audio_transcription":{"enabled":false},"translation":{"language":"en"},"turn_detection":null}
  // 
  // 单向 VAD、译音与转写：{"modalities":["text","audio"],"sample_rate":16000,"input_audio_format":"pcm","input_audio_transcription":{"enabled":true,"language":"auto"},"translation":{"language":"en"},"turn_detection":{"type":"server_vad","threshold":0.2,"silence_duration_ms":1000},"output_audio_format":"pcm","voice":"Tina","enable_voice_clone":false}
  // 
  // 双向手动提交、仅译文：{"input_audio":{"format":"pcm_s16le","sample_rate":16000,"channels":1},"output":{"modalities":["text"]},"transcription":{"enabled":false},"turn_detection":{"type":"manual"}}
  // 
  // 双向 VAD、译文与译音：{"input_audio":{"format":"pcm_s16le","sample_rate":16000,"channels":1},"output":{"modalities":["text","audio"],"audio":{"format":"pcm_s16le","sample_rate":24000,"channels":1,"voice":{"mode":"preset","id":"Tina"}}},"transcription":{"enabled":true},"turn_detection":{"type":"server_vad","threshold":0.2,"silence_duration_ms":1000}}
  // 
  // 单向扩展字段依据百炼 live-translator-client-events；新增原生字段须同步登记。显式 null 和热词动态键的映射保留行为仍需网关实测。
  Session *SessionUpdateSession `json:"session,omitempty" xml:"session,omitempty" type:"Struct"`
}

func (s SessionUpdate) String() string {
  return dara.Prettify(s)
}

func (s SessionUpdate) GoString() string {
  return s.String()
}

func (s *SessionUpdate) GetType() *string  {
  return s.Type
}

func (s *SessionUpdate) GetEventId() *string  {
  return s.EventId
}

func (s *SessionUpdate) GetSession() *SessionUpdateSession  {
  return s.Session
}

func (s *SessionUpdate) SetType(v string) *SessionUpdate {
  s.Type = &v
  return s
}

func (s *SessionUpdate) SetEventId(v string) *SessionUpdate {
  s.EventId = &v
  return s
}

func (s *SessionUpdate) SetSession(v *SessionUpdateSession) *SessionUpdate {
  s.Session = v
  return s
}

func (s *SessionUpdate) Validate() error {
  if s.Session != nil {
    if err := s.Session.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type SessionUpdateSession struct {
  // 仅双向使用；可省略，提供时必须为 bidirectional，不能切换模式。
  Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
  // 仅双向使用；可省略，提供时必须与创建时内容及顺序一致。
  LanguagePair []*string `json:"language_pair,omitempty" xml:"language_pair,omitempty" type:"Repeated"`
  // 仅双向使用；省略时后端使用 pcm_s16le、16000Hz、单声道，不接受 null。
  InputAudio *SessionUpdateSessionInputAudio `json:"input_audio,omitempty" xml:"input_audio,omitempty" type:"Struct"`
  // 仅双向使用；省略时后端输出 text 和 audio。
  Output *SessionUpdateSessionOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
  Transcription *SessionUpdateSessionTranscription `json:"transcription,omitempty" xml:"transcription,omitempty" type:"Struct"`
  // 单向 manual 必须显式传 null，VAD 使用 server_vad 对象。双向不接受 null，manual 使用对象；省略时后端为 server_vad。nullable 仅兼容单向，模式约束交后端校验。
  // 
  // if can be null:
  // true
  TurnDetection *SessionUpdateSessionTurnDetection `json:"turn_detection,omitempty" xml:"turn_detection,omitempty" type:"Struct"`
  // 仅单向使用；支持 text 或 text 与 audio，省略时模型默认 text 与 audio。双向使用 output.modalities。
  Modalities []*string `json:"modalities,omitempty" xml:"modalities,omitempty" type:"Repeated"`
  // 仅单向使用；须与输入音频一致，支持8000或16000。
  // 
  // example:
  // 
  // 16000
  SampleRate *int32 `json:"sample_rate,omitempty" xml:"sample_rate,omitempty"`
  // 仅单向原生协议使用；现有冒烟客户端使用 pcm，opus 尚未在本次实测。双向仅支持 pcm_s16le。
  // 
  // example:
  // 
  // pcm
  InputAudioFormat *string `json:"input_audio_format,omitempty" xml:"input_audio_format,omitempty"`
  // 仅单向使用，仅支持 enabled、language。对象省略或为空时默认开启、自动识别。enabled=false 时忽略 language，关闭原文转写；model 不再接受。默认值由后端补齐，网关不注入。
  InputAudioTranscription *SessionUpdateSessionInputAudioTranscription `json:"input_audio_transcription,omitempty" xml:"input_audio_transcription,omitempty" type:"Struct"`
  // 仅单向原生协议使用；双向目标语言来自创建时的语言对，不接受本字段。
  Translation *SessionUpdateSessionTranslation `json:"translation,omitempty" xml:"translation,omitempty" type:"Struct"`
  // 仅单向译音使用，当前原生协议只支持 pcm。
  // 
  // example:
  // 
  // pcm
  OutputAudioFormat *string `json:"output_audio_format,omitempty" xml:"output_audio_format,omitempty"`
  // 未启用复刻时使用预设音色，模型缺省为 Tina。启用复刻且 frequency 为 once/always 时必须为 default；never 时须为已有复刻音色ID。双向使用 output.audio.voice.id。
  // 
  // example:
  // 
  // Tina
  Voice *string `json:"voice,omitempty" xml:"voice,omitempty"`
  // 仅单向原生协议使用，模型默认 false；开启时联动 voice 和 voice_clone_options，实际可用性由上游校验。本次未实测复刻，双向不支持。
  EnableVoiceClone *bool `json:"enable_voice_clone,omitempty" xml:"enable_voice_clone,omitempty"`
  // 仅 enable_voice_clone=true 时生效；双向不支持本字段。
  VoiceCloneOptions *SessionUpdateSessionVoiceCloneOptions `json:"voice_clone_options,omitempty" xml:"voice_clone_options,omitempty" type:"Struct"`
}

func (s SessionUpdateSession) String() string {
  return dara.Prettify(s)
}

func (s SessionUpdateSession) GoString() string {
  return s.String()
}

func (s *SessionUpdateSession) GetMode() *string  {
  return s.Mode
}

func (s *SessionUpdateSession) GetLanguagePair() []*string  {
  return s.LanguagePair
}

func (s *SessionUpdateSession) GetInputAudio() *SessionUpdateSessionInputAudio  {
  return s.InputAudio
}

func (s *SessionUpdateSession) GetOutput() *SessionUpdateSessionOutput  {
  return s.Output
}

func (s *SessionUpdateSession) GetTranscription() *SessionUpdateSessionTranscription  {
  return s.Transcription
}

func (s *SessionUpdateSession) GetTurnDetection() *SessionUpdateSessionTurnDetection  {
  return s.TurnDetection
}

func (s *SessionUpdateSession) GetModalities() []*string  {
  return s.Modalities
}

func (s *SessionUpdateSession) GetSampleRate() *int32  {
  return s.SampleRate
}

func (s *SessionUpdateSession) GetInputAudioFormat() *string  {
  return s.InputAudioFormat
}

func (s *SessionUpdateSession) GetInputAudioTranscription() *SessionUpdateSessionInputAudioTranscription  {
  return s.InputAudioTranscription
}

func (s *SessionUpdateSession) GetTranslation() *SessionUpdateSessionTranslation  {
  return s.Translation
}

func (s *SessionUpdateSession) GetOutputAudioFormat() *string  {
  return s.OutputAudioFormat
}

func (s *SessionUpdateSession) GetVoice() *string  {
  return s.Voice
}

func (s *SessionUpdateSession) GetEnableVoiceClone() *bool  {
  return s.EnableVoiceClone
}

func (s *SessionUpdateSession) GetVoiceCloneOptions() *SessionUpdateSessionVoiceCloneOptions  {
  return s.VoiceCloneOptions
}

func (s *SessionUpdateSession) SetMode(v string) *SessionUpdateSession {
  s.Mode = &v
  return s
}

func (s *SessionUpdateSession) SetLanguagePair(v []*string) *SessionUpdateSession {
  s.LanguagePair = v
  return s
}

func (s *SessionUpdateSession) SetInputAudio(v *SessionUpdateSessionInputAudio) *SessionUpdateSession {
  s.InputAudio = v
  return s
}

func (s *SessionUpdateSession) SetOutput(v *SessionUpdateSessionOutput) *SessionUpdateSession {
  s.Output = v
  return s
}

func (s *SessionUpdateSession) SetTranscription(v *SessionUpdateSessionTranscription) *SessionUpdateSession {
  s.Transcription = v
  return s
}

func (s *SessionUpdateSession) SetTurnDetection(v *SessionUpdateSessionTurnDetection) *SessionUpdateSession {
  s.TurnDetection = v
  return s
}

func (s *SessionUpdateSession) SetModalities(v []*string) *SessionUpdateSession {
  s.Modalities = v
  return s
}

func (s *SessionUpdateSession) SetSampleRate(v int32) *SessionUpdateSession {
  s.SampleRate = &v
  return s
}

func (s *SessionUpdateSession) SetInputAudioFormat(v string) *SessionUpdateSession {
  s.InputAudioFormat = &v
  return s
}

func (s *SessionUpdateSession) SetInputAudioTranscription(v *SessionUpdateSessionInputAudioTranscription) *SessionUpdateSession {
  s.InputAudioTranscription = v
  return s
}

func (s *SessionUpdateSession) SetTranslation(v *SessionUpdateSessionTranslation) *SessionUpdateSession {
  s.Translation = v
  return s
}

func (s *SessionUpdateSession) SetOutputAudioFormat(v string) *SessionUpdateSession {
  s.OutputAudioFormat = &v
  return s
}

func (s *SessionUpdateSession) SetVoice(v string) *SessionUpdateSession {
  s.Voice = &v
  return s
}

func (s *SessionUpdateSession) SetEnableVoiceClone(v bool) *SessionUpdateSession {
  s.EnableVoiceClone = &v
  return s
}

func (s *SessionUpdateSession) SetVoiceCloneOptions(v *SessionUpdateSessionVoiceCloneOptions) *SessionUpdateSession {
  s.VoiceCloneOptions = v
  return s
}

func (s *SessionUpdateSession) Validate() error {
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
  if s.InputAudioTranscription != nil {
    if err := s.InputAudioTranscription.Validate(); err != nil {
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
  return nil
}

type SessionUpdateSessionInputAudio struct {
  Format *string `json:"format,omitempty" xml:"format,omitempty"`
  // example:
  // 
  // 16000
  SampleRate *int32 `json:"sample_rate,omitempty" xml:"sample_rate,omitempty"`
  Channels *int32 `json:"channels,omitempty" xml:"channels,omitempty"`
}

func (s SessionUpdateSessionInputAudio) String() string {
  return dara.Prettify(s)
}

func (s SessionUpdateSessionInputAudio) GoString() string {
  return s.String()
}

func (s *SessionUpdateSessionInputAudio) GetFormat() *string  {
  return s.Format
}

func (s *SessionUpdateSessionInputAudio) GetSampleRate() *int32  {
  return s.SampleRate
}

func (s *SessionUpdateSessionInputAudio) GetChannels() *int32  {
  return s.Channels
}

func (s *SessionUpdateSessionInputAudio) SetFormat(v string) *SessionUpdateSessionInputAudio {
  s.Format = &v
  return s
}

func (s *SessionUpdateSessionInputAudio) SetSampleRate(v int32) *SessionUpdateSessionInputAudio {
  s.SampleRate = &v
  return s
}

func (s *SessionUpdateSessionInputAudio) SetChannels(v int32) *SessionUpdateSessionInputAudio {
  s.Channels = &v
  return s
}

func (s *SessionUpdateSessionInputAudio) Validate() error {
  return dara.Validate(s)
}

type SessionUpdateSessionOutput struct {
  // 必须包含 text，可加 audio；不可重复，不能只有 audio。允许顺序反转，后端规范化为 text、audio。
  Modalities []*string `json:"modalities,omitempty" xml:"modalities,omitempty" type:"Repeated"`
  Audio *SessionUpdateSessionOutputAudio `json:"audio,omitempty" xml:"audio,omitempty" type:"Struct"`
}

func (s SessionUpdateSessionOutput) String() string {
  return dara.Prettify(s)
}

func (s SessionUpdateSessionOutput) GoString() string {
  return s.String()
}

func (s *SessionUpdateSessionOutput) GetModalities() []*string  {
  return s.Modalities
}

func (s *SessionUpdateSessionOutput) GetAudio() *SessionUpdateSessionOutputAudio  {
  return s.Audio
}

func (s *SessionUpdateSessionOutput) SetModalities(v []*string) *SessionUpdateSessionOutput {
  s.Modalities = v
  return s
}

func (s *SessionUpdateSessionOutput) SetAudio(v *SessionUpdateSessionOutputAudio) *SessionUpdateSessionOutput {
  s.Audio = v
  return s
}

func (s *SessionUpdateSessionOutput) Validate() error {
  if s.Audio != nil {
    if err := s.Audio.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type SessionUpdateSessionOutputAudio struct {
  Format *string `json:"format,omitempty" xml:"format,omitempty"`
  SampleRate *int32 `json:"sample_rate,omitempty" xml:"sample_rate,omitempty"`
  Channels *int32 `json:"channels,omitempty" xml:"channels,omitempty"`
  // 仅支持 preset；不支持 clone_frequency。id 省略或为 default 时后端使用 Tina。
  Voice *SessionUpdateSessionOutputAudioVoice `json:"voice,omitempty" xml:"voice,omitempty" type:"Struct"`
}

func (s SessionUpdateSessionOutputAudio) String() string {
  return dara.Prettify(s)
}

func (s SessionUpdateSessionOutputAudio) GoString() string {
  return s.String()
}

func (s *SessionUpdateSessionOutputAudio) GetFormat() *string  {
  return s.Format
}

func (s *SessionUpdateSessionOutputAudio) GetSampleRate() *int32  {
  return s.SampleRate
}

func (s *SessionUpdateSessionOutputAudio) GetChannels() *int32  {
  return s.Channels
}

func (s *SessionUpdateSessionOutputAudio) GetVoice() *SessionUpdateSessionOutputAudioVoice  {
  return s.Voice
}

func (s *SessionUpdateSessionOutputAudio) SetFormat(v string) *SessionUpdateSessionOutputAudio {
  s.Format = &v
  return s
}

func (s *SessionUpdateSessionOutputAudio) SetSampleRate(v int32) *SessionUpdateSessionOutputAudio {
  s.SampleRate = &v
  return s
}

func (s *SessionUpdateSessionOutputAudio) SetChannels(v int32) *SessionUpdateSessionOutputAudio {
  s.Channels = &v
  return s
}

func (s *SessionUpdateSessionOutputAudio) SetVoice(v *SessionUpdateSessionOutputAudioVoice) *SessionUpdateSessionOutputAudio {
  s.Voice = v
  return s
}

func (s *SessionUpdateSessionOutputAudio) Validate() error {
  if s.Voice != nil {
    if err := s.Voice.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type SessionUpdateSessionOutputAudioVoice struct {
  Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
  // 非空且不能含控制字符，最长128字符；具体音色支持范围由模型校验。
  // 
  // example:
  // 
  // Tina
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s SessionUpdateSessionOutputAudioVoice) String() string {
  return dara.Prettify(s)
}

func (s SessionUpdateSessionOutputAudioVoice) GoString() string {
  return s.String()
}

func (s *SessionUpdateSessionOutputAudioVoice) GetMode() *string  {
  return s.Mode
}

func (s *SessionUpdateSessionOutputAudioVoice) GetId() *string  {
  return s.Id
}

func (s *SessionUpdateSessionOutputAudioVoice) SetMode(v string) *SessionUpdateSessionOutputAudioVoice {
  s.Mode = &v
  return s
}

func (s *SessionUpdateSessionOutputAudioVoice) SetId(v string) *SessionUpdateSessionOutputAudioVoice {
  s.Id = &v
  return s
}

func (s *SessionUpdateSessionOutputAudioVoice) Validate() error {
  return dara.Validate(s)
}

type SessionUpdateSessionTranscription struct {
  // 省略时后端为 true；关闭输出不会关闭内部语种识别。
  Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s SessionUpdateSessionTranscription) String() string {
  return dara.Prettify(s)
}

func (s SessionUpdateSessionTranscription) GoString() string {
  return s.String()
}

func (s *SessionUpdateSessionTranscription) GetEnabled() *bool  {
  return s.Enabled
}

func (s *SessionUpdateSessionTranscription) SetEnabled(v bool) *SessionUpdateSessionTranscription {
  s.Enabled = &v
  return s
}

func (s *SessionUpdateSessionTranscription) Validate() error {
  return dara.Validate(s)
}

type SessionUpdateSessionTurnDetection struct {
  // manual 对象仅用于双向；单向手动模式使用整个 turn_detection 为 null。
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // 范围 -1 到 1，后端缺省0.2。
  // 
  // example:
  // 
  // 0.2
  Threshold *float32 `json:"threshold,omitempty" xml:"threshold,omitempty"`
  // 范围200到6000，后端缺省1000。
  // 
  // example:
  // 
  // 1000
  SilenceDurationMs *int32 `json:"silence_duration_ms,omitempty" xml:"silence_duration_ms,omitempty"`
}

func (s SessionUpdateSessionTurnDetection) String() string {
  return dara.Prettify(s)
}

func (s SessionUpdateSessionTurnDetection) GoString() string {
  return s.String()
}

func (s *SessionUpdateSessionTurnDetection) GetType() *string  {
  return s.Type
}

func (s *SessionUpdateSessionTurnDetection) GetThreshold() *float32  {
  return s.Threshold
}

func (s *SessionUpdateSessionTurnDetection) GetSilenceDurationMs() *int32  {
  return s.SilenceDurationMs
}

func (s *SessionUpdateSessionTurnDetection) SetType(v string) *SessionUpdateSessionTurnDetection {
  s.Type = &v
  return s
}

func (s *SessionUpdateSessionTurnDetection) SetThreshold(v float32) *SessionUpdateSessionTurnDetection {
  s.Threshold = &v
  return s
}

func (s *SessionUpdateSessionTurnDetection) SetSilenceDurationMs(v int32) *SessionUpdateSessionTurnDetection {
  s.SilenceDurationMs = &v
  return s
}

func (s *SessionUpdateSessionTurnDetection) Validate() error {
  return dara.Validate(s)
}

type SessionUpdateSessionInputAudioTranscription struct {
  // 默认 auto，省略或 auto 表示自动识别。开启时保留指定的受支持语种；关闭时忽略本字段。此字段不是翻译目标语言。
  // 
  // example:
  // 
  // auto
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  // 默认 true。false 时忽略 language，不影响译文和译音。
  // 
  // example:
  // 
  // true
  Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s SessionUpdateSessionInputAudioTranscription) String() string {
  return dara.Prettify(s)
}

func (s SessionUpdateSessionInputAudioTranscription) GoString() string {
  return s.String()
}

func (s *SessionUpdateSessionInputAudioTranscription) GetLanguage() *string  {
  return s.Language
}

func (s *SessionUpdateSessionInputAudioTranscription) GetEnabled() *bool  {
  return s.Enabled
}

func (s *SessionUpdateSessionInputAudioTranscription) SetLanguage(v string) *SessionUpdateSessionInputAudioTranscription {
  s.Language = &v
  return s
}

func (s *SessionUpdateSessionInputAudioTranscription) SetEnabled(v bool) *SessionUpdateSessionInputAudioTranscription {
  s.Enabled = &v
  return s
}

func (s *SessionUpdateSessionInputAudioTranscription) Validate() error {
  return dara.Validate(s)
}

type SessionUpdateSessionTranslation struct {
  // 例如 en，缺省时模型默认 en；具体支持范围由模型校验。
  // 
  // example:
  // 
  // en
  Language *string `json:"language,omitempty" xml:"language,omitempty"`
  // 仅单向 Qwen 3.5 使用，目标语种为 zh 或 en 时生效；源语种等于目标语种时按设置跳过输出。
  SameLanguageSkipOptions *SessionUpdateSessionTranslationSameLanguageSkipOptions `json:"same_language_skip_options,omitempty" xml:"same_language_skip_options,omitempty" type:"Struct"`
  // 仅单向原生协议使用；与双向不支持的 glossary 字段不同。
  Corpus *SessionUpdateSessionTranslationCorpus `json:"corpus,omitempty" xml:"corpus,omitempty" type:"Struct"`
}

func (s SessionUpdateSessionTranslation) String() string {
  return dara.Prettify(s)
}

func (s SessionUpdateSessionTranslation) GoString() string {
  return s.String()
}

func (s *SessionUpdateSessionTranslation) GetLanguage() *string  {
  return s.Language
}

func (s *SessionUpdateSessionTranslation) GetSameLanguageSkipOptions() *SessionUpdateSessionTranslationSameLanguageSkipOptions  {
  return s.SameLanguageSkipOptions
}

func (s *SessionUpdateSessionTranslation) GetCorpus() *SessionUpdateSessionTranslationCorpus  {
  return s.Corpus
}

func (s *SessionUpdateSessionTranslation) SetLanguage(v string) *SessionUpdateSessionTranslation {
  s.Language = &v
  return s
}

func (s *SessionUpdateSessionTranslation) SetSameLanguageSkipOptions(v *SessionUpdateSessionTranslationSameLanguageSkipOptions) *SessionUpdateSessionTranslation {
  s.SameLanguageSkipOptions = v
  return s
}

func (s *SessionUpdateSessionTranslation) SetCorpus(v *SessionUpdateSessionTranslationCorpus) *SessionUpdateSessionTranslation {
  s.Corpus = v
  return s
}

func (s *SessionUpdateSessionTranslation) Validate() error {
  if s.SameLanguageSkipOptions != nil {
    if err := s.SameLanguageSkipOptions.Validate(); err != nil {
      return err
    }
  }
  if s.Corpus != nil {
    if err := s.Corpus.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type SessionUpdateSessionTranslationSameLanguageSkipOptions struct {
  SkipText *bool `json:"skip_text,omitempty" xml:"skip_text,omitempty"`
  SkipAudio *bool `json:"skip_audio,omitempty" xml:"skip_audio,omitempty"`
}

func (s SessionUpdateSessionTranslationSameLanguageSkipOptions) String() string {
  return dara.Prettify(s)
}

func (s SessionUpdateSessionTranslationSameLanguageSkipOptions) GoString() string {
  return s.String()
}

func (s *SessionUpdateSessionTranslationSameLanguageSkipOptions) GetSkipText() *bool  {
  return s.SkipText
}

func (s *SessionUpdateSessionTranslationSameLanguageSkipOptions) GetSkipAudio() *bool  {
  return s.SkipAudio
}

func (s *SessionUpdateSessionTranslationSameLanguageSkipOptions) SetSkipText(v bool) *SessionUpdateSessionTranslationSameLanguageSkipOptions {
  s.SkipText = &v
  return s
}

func (s *SessionUpdateSessionTranslationSameLanguageSkipOptions) SetSkipAudio(v bool) *SessionUpdateSessionTranslationSameLanguageSkipOptions {
  s.SkipAudio = &v
  return s
}

func (s *SessionUpdateSessionTranslationSameLanguageSkipOptions) Validate() error {
  return dara.Validate(s)
}

type SessionUpdateSessionTranslationCorpus struct {
  // 动态 key 为源语言词汇，value 为目标语言译词；不固定热词名称，动态键映射保留需网关实测。
  // 
  // example:
  // 
  // {"人工智能":"Artificial Intelligence","机器学习":"Machine Learning"}
  Phrases map[string]*string `json:"phrases,omitempty" xml:"phrases,omitempty"`
}

func (s SessionUpdateSessionTranslationCorpus) String() string {
  return dara.Prettify(s)
}

func (s SessionUpdateSessionTranslationCorpus) GoString() string {
  return s.String()
}

func (s *SessionUpdateSessionTranslationCorpus) GetPhrases() map[string]*string  {
  return s.Phrases
}

func (s *SessionUpdateSessionTranslationCorpus) SetPhrases(v map[string]*string) *SessionUpdateSessionTranslationCorpus {
  s.Phrases = v
  return s
}

func (s *SessionUpdateSessionTranslationCorpus) Validate() error {
  return dara.Validate(s)
}

type SessionUpdateSessionVoiceCloneOptions struct {
  // never 使用已有复刻音色；once 会话开始时复刻一次；always 每次输出前复刻。once/always 必须设置 voice=default。
  Frequency *string `json:"frequency,omitempty" xml:"frequency,omitempty"`
}

func (s SessionUpdateSessionVoiceCloneOptions) String() string {
  return dara.Prettify(s)
}

func (s SessionUpdateSessionVoiceCloneOptions) GoString() string {
  return s.String()
}

func (s *SessionUpdateSessionVoiceCloneOptions) GetFrequency() *string  {
  return s.Frequency
}

func (s *SessionUpdateSessionVoiceCloneOptions) SetFrequency(v string) *SessionUpdateSessionVoiceCloneOptions {
  s.Frequency = &v
  return s
}

func (s *SessionUpdateSessionVoiceCloneOptions) Validate() error {
  return dara.Validate(s)
}

