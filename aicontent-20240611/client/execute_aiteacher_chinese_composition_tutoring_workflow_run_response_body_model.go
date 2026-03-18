// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetContent(v string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody
  GetContent() *string 
  SetEvent(v string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody
  GetEvent() *string 
  SetRequestId(v string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody
  GetRequestId() *string 
}

type ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody struct {
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

func (s ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody) GetContent() *string  {
  return s.Content
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody) GetEvent() *string  {
  return s.Event
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody) SetContent(v string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody {
  s.Content = &v
  return s
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody) SetEvent(v string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody {
  s.Event = &v
  return s
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody) SetRequestId(v string) *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteAITeacherChineseCompositionTutoringWorkflowRunResponseBody) Validate() error {
  return dara.Validate(s)
}

