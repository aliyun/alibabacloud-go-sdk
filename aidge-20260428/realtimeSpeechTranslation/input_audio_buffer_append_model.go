// This file is auto-generated, don't edit it. Thanks.
package realtimeSpeechTranslation

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iInputAudioBufferAppend interface {
  dara.Model
  String() string
  GoString() string
  SetType(v string) *InputAudioBufferAppend
  GetType() *string 
  SetEventId(v string) *InputAudioBufferAppend
  GetEventId() *string 
  SetAudio(v []byte) *InputAudioBufferAppend
  GetAudio() []byte 
  SetSequence(v int64) *InputAudioBufferAppend
  GetSequence() *int64 
}

type InputAudioBufferAppend struct {
  // This parameter is required.
  Type *string `json:"type,omitempty" xml:"type,omitempty"`
  // 可省略；提供时必须为非空、非纯空白字符串。
  // 
  // example:
  // 
  // evt_audio_1
  EventId *string `json:"event_id,omitempty" xml:"event_id,omitempty"`
  // 等待 SessionUpdated 后发送；必须为 JSON 文本消息中的 Base64 字符串，不能发送 binary 帧。解码后每块最多10240字节；双向必须为非空 PCM16 完整采样，字节数为偶数。解码与大小检查由后端执行。
  // 
  // This parameter is required.
  Audio []byte `json:"audio,omitempty" xml:"audio,omitempty"`
  // 仅双向使用；可省略，提供时从0开始递增。只有携带 sequence 且成功接受的块推进计数；clear 不重置。网关不补默认值，连续性由后端校验。
  // 
  // example:
  // 
  // 0
  Sequence *int64 `json:"sequence,omitempty" xml:"sequence,omitempty"`
}

func (s InputAudioBufferAppend) String() string {
  return dara.Prettify(s)
}

func (s InputAudioBufferAppend) GoString() string {
  return s.String()
}

func (s *InputAudioBufferAppend) GetType() *string  {
  return s.Type
}

func (s *InputAudioBufferAppend) GetEventId() *string  {
  return s.EventId
}

func (s *InputAudioBufferAppend) GetAudio() []byte  {
  return s.Audio
}

func (s *InputAudioBufferAppend) GetSequence() *int64  {
  return s.Sequence
}

func (s *InputAudioBufferAppend) SetType(v string) *InputAudioBufferAppend {
  s.Type = &v
  return s
}

func (s *InputAudioBufferAppend) SetEventId(v string) *InputAudioBufferAppend {
  s.EventId = &v
  return s
}

func (s *InputAudioBufferAppend) SetAudio(v []byte) *InputAudioBufferAppend {
  s.Audio = v
  return s
}

func (s *InputAudioBufferAppend) SetSequence(v int64) *InputAudioBufferAppend {
  s.Sequence = &v
  return s
}

func (s *InputAudioBufferAppend) Validate() error {
  return dara.Validate(s)
}

