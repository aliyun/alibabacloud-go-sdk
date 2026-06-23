// This file is auto-generated, don't edit it. Thanks.
package endToEndRealTimeDialog

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iTranscriptionEnd interface {
  dara.Model
  String() string
  GoString() string
}

type TranscriptionEnd struct {
}

func (s TranscriptionEnd) String() string {
  return dara.Prettify(s)
}

func (s TranscriptionEnd) GoString() string {
  return s.String()
}

func (s *TranscriptionEnd) Validate() error {
  return dara.Validate(s)
}

