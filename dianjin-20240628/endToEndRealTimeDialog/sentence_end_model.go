// This file is auto-generated, don't edit it. Thanks.
package endToEndRealTimeDialog

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iSentenceEnd interface {
  dara.Model
  String() string
  GoString() string
  SetData(v []*int64) *SentenceEnd
  GetData() []*int64 
  SetMessageId(v string) *SentenceEnd
  GetMessageId() *string 
}

type SentenceEnd struct {
  Data []*int64 `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
  MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
}

func (s SentenceEnd) String() string {
  return dara.Prettify(s)
}

func (s SentenceEnd) GoString() string {
  return s.String()
}

func (s *SentenceEnd) GetData() []*int64  {
  return s.Data
}

func (s *SentenceEnd) GetMessageId() *string  {
  return s.MessageId
}

func (s *SentenceEnd) SetData(v []*int64) *SentenceEnd {
  s.Data = v
  return s
}

func (s *SentenceEnd) SetMessageId(v string) *SentenceEnd {
  s.MessageId = &v
  return s
}

func (s *SentenceEnd) Validate() error {
  return dara.Validate(s)
}

