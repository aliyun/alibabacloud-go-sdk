// This file is auto-generated, don't edit it. Thanks.
package endToEndRealTimeDialog

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iTranscriptionStarted interface {
  dara.Model
  String() string
  GoString() string
  SetSessionId(v string) *TranscriptionStarted
  GetSessionId() *string 
  SetOpeningRemarks(v string) *TranscriptionStarted
  GetOpeningRemarks() *string 
}

type TranscriptionStarted struct {
  SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
  OpeningRemarks *string `json:"openingRemarks,omitempty" xml:"openingRemarks,omitempty"`
}

func (s TranscriptionStarted) String() string {
  return dara.Prettify(s)
}

func (s TranscriptionStarted) GoString() string {
  return s.String()
}

func (s *TranscriptionStarted) GetSessionId() *string  {
  return s.SessionId
}

func (s *TranscriptionStarted) GetOpeningRemarks() *string  {
  return s.OpeningRemarks
}

func (s *TranscriptionStarted) SetSessionId(v string) *TranscriptionStarted {
  s.SessionId = &v
  return s
}

func (s *TranscriptionStarted) SetOpeningRemarks(v string) *TranscriptionStarted {
  s.OpeningRemarks = &v
  return s
}

func (s *TranscriptionStarted) Validate() error {
  return dara.Validate(s)
}

