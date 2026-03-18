// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAITeacherGrammarCheckRequest interface {
  dara.Model
  String() string
  GoString() string
  SetContent(v string) *ExecuteAITeacherGrammarCheckRequest
  GetContent() *string 
  SetUserId(v string) *ExecuteAITeacherGrammarCheckRequest
  GetUserId() *string 
}

type ExecuteAITeacherGrammarCheckRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // i is good
  Content *string `json:"content,omitempty" xml:"content,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 886eba3702xxxxxxxxx4ba52a87a525
  UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ExecuteAITeacherGrammarCheckRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherGrammarCheckRequest) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherGrammarCheckRequest) GetContent() *string  {
  return s.Content
}

func (s *ExecuteAITeacherGrammarCheckRequest) GetUserId() *string  {
  return s.UserId
}

func (s *ExecuteAITeacherGrammarCheckRequest) SetContent(v string) *ExecuteAITeacherGrammarCheckRequest {
  s.Content = &v
  return s
}

func (s *ExecuteAITeacherGrammarCheckRequest) SetUserId(v string) *ExecuteAITeacherGrammarCheckRequest {
  s.UserId = &v
  return s
}

func (s *ExecuteAITeacherGrammarCheckRequest) Validate() error {
  return dara.Validate(s)
}

