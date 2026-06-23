// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEndToEndRealTimeDialogRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAsrModelId(v string) *EndToEndRealTimeDialogRequest
  GetAsrModelId() *string 
  SetInputFormat(v string) *EndToEndRealTimeDialogRequest
  GetInputFormat() *string 
  SetOutputFormat(v string) *EndToEndRealTimeDialogRequest
  GetOutputFormat() *string 
  SetPitchRate(v int32) *EndToEndRealTimeDialogRequest
  GetPitchRate() *int32 
  SetSampleRate(v string) *EndToEndRealTimeDialogRequest
  GetSampleRate() *string 
  SetSpeechRate(v int32) *EndToEndRealTimeDialogRequest
  GetSpeechRate() *int32 
  SetTtsModelId(v string) *EndToEndRealTimeDialogRequest
  GetTtsModelId() *string 
  SetVoiceCode(v string) *EndToEndRealTimeDialogRequest
  GetVoiceCode() *string 
  SetVolume(v int32) *EndToEndRealTimeDialogRequest
  GetVolume() *int32 
}

type EndToEndRealTimeDialogRequest struct {
  // The ID of the speech recognition model. The default value is nls-base. Supported models include paraformer-realtime-v2 and paraformer-realtime-8k-v2.
  // 
  // example:
  // 
  // nls-base
  AsrModelId *string `json:"asrModelId,omitempty" xml:"asrModelId,omitempty"`
  // The input audio format. Supported formats are pcm, wav, and mp3.
  // 
  // example:
  // 
  // pcm
  InputFormat *string `json:"inputFormat,omitempty" xml:"inputFormat,omitempty"`
  // The output audio format.
  // 
  // example:
  // 
  // wav
  OutputFormat *string `json:"outputFormat,omitempty" xml:"outputFormat,omitempty"`
  // The pitch rate.
  // 
  // ● If \\`ttsModelId\\` is \\`nls-base\\`:
  // 
  // The value ranges from -500 to 500. The default is 0.
  // 
  // ● If \\`ttsModelId\\` is \\`cosyvoice-v2\\`: Specifies the pitch of the synthesized audio. The value ranges from 0.5 to 2.
  // 
  // example:
  // 
  // 0
  PitchRate *int32 `json:"pitchRate,omitempty" xml:"pitchRate,omitempty"`
  // The sample rate.
  // 
  // example:
  // 
  // SAMPLE_RATE_16K
  SampleRate *string `json:"sampleRate,omitempty" xml:"sampleRate,omitempty"`
  // The speech rate.
  // 
  // ● If \\`ttsModelId\\` is \\`nls-base\\`: The value ranges from -500 to 500. The default is 0.
  // 
  // ● If \\`ttsModelId\\` is \\`cosyvoice-v2\\`:
  // 
  // Specifies the speech rate of the synthesized audio. The value ranges from 0.5 to 2.
  // 
  // ○ 0.5: Half the default speed.
  // 
  // ○ 1: The default speed. The default speed is the model\\"s standard output speed and may vary slightly by speaker. It is about four characters per second.
  // 
  // ○ 2: Twice the default speed.
  // 
  // example:
  // 
  // 0
  SpeechRate *int32 `json:"speechRate,omitempty" xml:"speechRate,omitempty"`
  // The ID of the speech synthesis model. The default value is nls-base. The cosyvoice-v2 model is supported.
  // 
  // example:
  // 
  // nls-base
  TtsModelId *string `json:"ttsModelId,omitempty" xml:"ttsModelId,omitempty"`
  // The voice parameter. This is available only for models that support word-level or sentence-level timestamps.
  // 
  // example:
  // 
  // longxiaochun_v2
  VoiceCode *string `json:"voiceCode,omitempty" xml:"voiceCode,omitempty"`
  // The volume. The value ranges from 0 to 100. This parameter is optional. The default value is 50.
  // 
  // example:
  // 
  // 50
  Volume *int32 `json:"volume,omitempty" xml:"volume,omitempty"`
}

func (s EndToEndRealTimeDialogRequest) String() string {
  return dara.Prettify(s)
}

func (s EndToEndRealTimeDialogRequest) GoString() string {
  return s.String()
}

func (s *EndToEndRealTimeDialogRequest) GetAsrModelId() *string  {
  return s.AsrModelId
}

func (s *EndToEndRealTimeDialogRequest) GetInputFormat() *string  {
  return s.InputFormat
}

func (s *EndToEndRealTimeDialogRequest) GetOutputFormat() *string  {
  return s.OutputFormat
}

func (s *EndToEndRealTimeDialogRequest) GetPitchRate() *int32  {
  return s.PitchRate
}

func (s *EndToEndRealTimeDialogRequest) GetSampleRate() *string  {
  return s.SampleRate
}

func (s *EndToEndRealTimeDialogRequest) GetSpeechRate() *int32  {
  return s.SpeechRate
}

func (s *EndToEndRealTimeDialogRequest) GetTtsModelId() *string  {
  return s.TtsModelId
}

func (s *EndToEndRealTimeDialogRequest) GetVoiceCode() *string  {
  return s.VoiceCode
}

func (s *EndToEndRealTimeDialogRequest) GetVolume() *int32  {
  return s.Volume
}

func (s *EndToEndRealTimeDialogRequest) SetAsrModelId(v string) *EndToEndRealTimeDialogRequest {
  s.AsrModelId = &v
  return s
}

func (s *EndToEndRealTimeDialogRequest) SetInputFormat(v string) *EndToEndRealTimeDialogRequest {
  s.InputFormat = &v
  return s
}

func (s *EndToEndRealTimeDialogRequest) SetOutputFormat(v string) *EndToEndRealTimeDialogRequest {
  s.OutputFormat = &v
  return s
}

func (s *EndToEndRealTimeDialogRequest) SetPitchRate(v int32) *EndToEndRealTimeDialogRequest {
  s.PitchRate = &v
  return s
}

func (s *EndToEndRealTimeDialogRequest) SetSampleRate(v string) *EndToEndRealTimeDialogRequest {
  s.SampleRate = &v
  return s
}

func (s *EndToEndRealTimeDialogRequest) SetSpeechRate(v int32) *EndToEndRealTimeDialogRequest {
  s.SpeechRate = &v
  return s
}

func (s *EndToEndRealTimeDialogRequest) SetTtsModelId(v string) *EndToEndRealTimeDialogRequest {
  s.TtsModelId = &v
  return s
}

func (s *EndToEndRealTimeDialogRequest) SetVoiceCode(v string) *EndToEndRealTimeDialogRequest {
  s.VoiceCode = &v
  return s
}

func (s *EndToEndRealTimeDialogRequest) SetVolume(v int32) *EndToEndRealTimeDialogRequest {
  s.Volume = &v
  return s
}

func (s *EndToEndRealTimeDialogRequest) Validate() error {
  return dara.Validate(s)
}

