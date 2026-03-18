// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAITeacherEnglishParaphraseChatMessageResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetContent(v string) *ExecuteAITeacherEnglishParaphraseChatMessageResponseBody
  GetContent() *string 
  SetEvent(v string) *ExecuteAITeacherEnglishParaphraseChatMessageResponseBody
  GetEvent() *string 
  SetRequestId(v string) *ExecuteAITeacherEnglishParaphraseChatMessageResponseBody
  GetRequestId() *string 
}

type ExecuteAITeacherEnglishParaphraseChatMessageResponseBody struct {
  // example:
  // 
  // how
  Content *string `json:"content,omitempty" xml:"content,omitempty"`
  // example:
  // 
  // message
  Event *string `json:"event,omitempty" xml:"event,omitempty"`
  // example:
  // 
  // xxxx-xxxx-xxxx-xxxxxxxx
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ExecuteAITeacherEnglishParaphraseChatMessageResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherEnglishParaphraseChatMessageResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageResponseBody) GetContent() *string  {
  return s.Content
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageResponseBody) GetEvent() *string  {
  return s.Event
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageResponseBody) SetContent(v string) *ExecuteAITeacherEnglishParaphraseChatMessageResponseBody {
  s.Content = &v
  return s
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageResponseBody) SetEvent(v string) *ExecuteAITeacherEnglishParaphraseChatMessageResponseBody {
  s.Event = &v
  return s
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageResponseBody) SetRequestId(v string) *ExecuteAITeacherEnglishParaphraseChatMessageResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageResponseBody) Validate() error {
  return dara.Validate(s)
}

