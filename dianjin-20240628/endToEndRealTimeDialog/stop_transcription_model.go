// This file is auto-generated, don't edit it. Thanks.
package endToEndRealTimeDialog

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iStopTranscription interface {
  dara.Model
  String() string
  GoString() string
  SetSessionId(v string) *StopTranscription
  GetSessionId() *string 
}

type StopTranscription struct {
  SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s StopTranscription) String() string {
  return dara.Prettify(s)
}

func (s StopTranscription) GoString() string {
  return s.String()
}

func (s *StopTranscription) GetSessionId() *string  {
  return s.SessionId
}

func (s *StopTranscription) SetSessionId(v string) *StopTranscription {
  s.SessionId = &v
  return s
}

func (s *StopTranscription) Validate() error {
  return dara.Validate(s)
}

