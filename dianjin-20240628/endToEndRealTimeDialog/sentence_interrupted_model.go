// This file is auto-generated, don't edit it. Thanks.
package endToEndRealTimeDialog

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iSentenceInterrupted interface {
  dara.Model
  String() string
  GoString() string
}

type SentenceInterrupted struct {
}

func (s SentenceInterrupted) String() string {
  return dara.Prettify(s)
}

func (s SentenceInterrupted) GoString() string {
  return s.String()
}

func (s *SentenceInterrupted) Validate() error {
  return dara.Validate(s)
}

