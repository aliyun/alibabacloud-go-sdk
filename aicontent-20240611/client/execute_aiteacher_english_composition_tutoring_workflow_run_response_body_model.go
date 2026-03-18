// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetContent(v string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody
  GetContent() *string 
  SetEvent(v string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody
  GetEvent() *string 
  SetRequestId(v string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody
  GetRequestId() *string 
}

type ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody struct {
  // example:
  // 
  // hi
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

func (s ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody) GetContent() *string  {
  return s.Content
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody) GetEvent() *string  {
  return s.Event
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody) SetContent(v string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody {
  s.Content = &v
  return s
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody) SetEvent(v string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody {
  s.Event = &v
  return s
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody) SetRequestId(v string) *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteAITeacherEnglishCompositionTutoringWorkflowRunResponseBody) Validate() error {
  return dara.Validate(s)
}

