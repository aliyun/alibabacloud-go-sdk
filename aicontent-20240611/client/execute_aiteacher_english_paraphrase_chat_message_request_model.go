// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAITeacherEnglishParaphraseChatMessageRequest interface {
  dara.Model
  String() string
  GoString() string
  SetChatId(v string) *ExecuteAITeacherEnglishParaphraseChatMessageRequest
  GetChatId() *string 
  SetContent(v string) *ExecuteAITeacherEnglishParaphraseChatMessageRequest
  GetContent() *string 
  SetGrade(v int64) *ExecuteAITeacherEnglishParaphraseChatMessageRequest
  GetGrade() *int64 
  SetQuestionId(v string) *ExecuteAITeacherEnglishParaphraseChatMessageRequest
  GetQuestionId() *string 
  SetQuestionInfo(v string) *ExecuteAITeacherEnglishParaphraseChatMessageRequest
  GetQuestionInfo() *string 
  SetResponseMode(v string) *ExecuteAITeacherEnglishParaphraseChatMessageRequest
  GetResponseMode() *string 
  SetUserAnswer(v string) *ExecuteAITeacherEnglishParaphraseChatMessageRequest
  GetUserAnswer() *string 
  SetUserId(v string) *ExecuteAITeacherEnglishParaphraseChatMessageRequest
  GetUserId() *string 
}

type ExecuteAITeacherEnglishParaphraseChatMessageRequest struct {
  // example:
  // 
  // 6788e0b475a4631ffc626722
  ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // How much is this?
  Content *string `json:"content,omitempty" xml:"content,omitempty"`
  // example:
  // 
  // 3
  Grade *int64 `json:"grade,omitempty" xml:"grade,omitempty"`
  // example:
  // 
  // xxxxxxxxx
  QuestionId *string `json:"questionId,omitempty" xml:"questionId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // How to inquire about the price
  QuestionInfo *string `json:"questionInfo,omitempty" xml:"questionInfo,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // sreaming
  ResponseMode *string `json:"responseMode,omitempty" xml:"responseMode,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // How much is this?
  UserAnswer *string `json:"userAnswer,omitempty" xml:"userAnswer,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // xxxxxxx
  UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ExecuteAITeacherEnglishParaphraseChatMessageRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherEnglishParaphraseChatMessageRequest) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageRequest) GetChatId() *string  {
  return s.ChatId
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageRequest) GetContent() *string  {
  return s.Content
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageRequest) GetGrade() *int64  {
  return s.Grade
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageRequest) GetQuestionId() *string  {
  return s.QuestionId
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageRequest) GetQuestionInfo() *string  {
  return s.QuestionInfo
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageRequest) GetResponseMode() *string  {
  return s.ResponseMode
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageRequest) GetUserAnswer() *string  {
  return s.UserAnswer
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageRequest) GetUserId() *string  {
  return s.UserId
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageRequest) SetChatId(v string) *ExecuteAITeacherEnglishParaphraseChatMessageRequest {
  s.ChatId = &v
  return s
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageRequest) SetContent(v string) *ExecuteAITeacherEnglishParaphraseChatMessageRequest {
  s.Content = &v
  return s
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageRequest) SetGrade(v int64) *ExecuteAITeacherEnglishParaphraseChatMessageRequest {
  s.Grade = &v
  return s
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageRequest) SetQuestionId(v string) *ExecuteAITeacherEnglishParaphraseChatMessageRequest {
  s.QuestionId = &v
  return s
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageRequest) SetQuestionInfo(v string) *ExecuteAITeacherEnglishParaphraseChatMessageRequest {
  s.QuestionInfo = &v
  return s
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageRequest) SetResponseMode(v string) *ExecuteAITeacherEnglishParaphraseChatMessageRequest {
  s.ResponseMode = &v
  return s
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageRequest) SetUserAnswer(v string) *ExecuteAITeacherEnglishParaphraseChatMessageRequest {
  s.UserAnswer = &v
  return s
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageRequest) SetUserId(v string) *ExecuteAITeacherEnglishParaphraseChatMessageRequest {
  s.UserId = &v
  return s
}

func (s *ExecuteAITeacherEnglishParaphraseChatMessageRequest) Validate() error {
  return dara.Validate(s)
}

