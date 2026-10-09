// This file is auto-generated, don't edit it. Thanks.
package endToEndRealTimeDialog

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iTranscriptionResultChanged interface {
  dara.Model
  String() string
  GoString() string
  SetContent(v string) *TranscriptionResultChanged
  GetContent() *string 
  SetMessageId(v string) *TranscriptionResultChanged
  GetMessageId() *string 
}

type TranscriptionResultChanged struct {
  Content *string `json:"content,omitempty" xml:"content,omitempty"`
  MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
}

func (s TranscriptionResultChanged) String() string {
  return dara.Prettify(s)
}

func (s TranscriptionResultChanged) GoString() string {
  return s.String()
}

func (s *TranscriptionResultChanged) GetContent() *string  {
  return s.Content
}

func (s *TranscriptionResultChanged) GetMessageId() *string  {
  return s.MessageId
}

func (s *TranscriptionResultChanged) SetContent(v string) *TranscriptionResultChanged {
  s.Content = &v
  return s
}

func (s *TranscriptionResultChanged) SetMessageId(v string) *TranscriptionResultChanged {
  s.MessageId = &v
  return s
}

func (s *TranscriptionResultChanged) Validate() error {
  return dara.Validate(s)
}

