// This file is auto-generated, don't edit it. Thanks.
package streamVoice

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iVoiceInputAudio interface {
  dara.Model
  String() string
  GoString() string
}

type VoiceInputAudio struct {
}

func (s VoiceInputAudio) String() string {
  return dara.Prettify(s)
}

func (s VoiceInputAudio) GoString() string {
  return s.String()
}

func (s *VoiceInputAudio) Validate() error {
  return dara.Validate(s)
}

