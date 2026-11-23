// This file is auto-generated, don't edit it. Thanks.
package streamVoice

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iVoiceOutputAudio interface {
  dara.Model
  String() string
  GoString() string
}

type VoiceOutputAudio struct {
}

func (s VoiceOutputAudio) String() string {
  return dara.Prettify(s)
}

func (s VoiceOutputAudio) GoString() string {
  return s.String()
}

func (s *VoiceOutputAudio) Validate() error {
  return dara.Validate(s)
}

