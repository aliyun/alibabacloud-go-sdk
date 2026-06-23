// This file is auto-generated, don't edit it. Thanks.
package endToEndRealTimeDialog

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iTranscriptionCompleted interface {
  dara.Model
  String() string
  GoString() string
}

type TranscriptionCompleted struct {
}

func (s TranscriptionCompleted) String() string {
  return dara.Prettify(s)
}

func (s TranscriptionCompleted) GoString() string {
  return s.String()
}

func (s *TranscriptionCompleted) Validate() error {
  return dara.Validate(s)
}

