// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTextbookAssistantDifficultyRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAction(v string) *ExecuteTextbookAssistantDifficultyRequest
  GetAction() *string 
  SetAssistant(v string) *ExecuteTextbookAssistantDifficultyRequest
  GetAssistant() *string 
  SetAuthToken(v string) *ExecuteTextbookAssistantDifficultyRequest
  GetAuthToken() *string 
  SetChatId(v string) *ExecuteTextbookAssistantDifficultyRequest
  GetChatId() *string 
  SetScenario(v string) *ExecuteTextbookAssistantDifficultyRequest
  GetScenario() *string 
}

type ExecuteTextbookAssistantDifficultyRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // UP
  Action *string `json:"action,omitempty" xml:"action,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 6788f4a6b54c5268c1b78a25
  Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // tc_e6dc70c890866f4028ca685b6fa29874
  AuthToken *string `json:"authToken,omitempty" xml:"authToken,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 6788e0b475a4631ffc626722
  ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // SYNC
  Scenario *string `json:"scenario,omitempty" xml:"scenario,omitempty"`
}

func (s ExecuteTextbookAssistantDifficultyRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantDifficultyRequest) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantDifficultyRequest) GetAction() *string  {
  return s.Action
}

func (s *ExecuteTextbookAssistantDifficultyRequest) GetAssistant() *string  {
  return s.Assistant
}

func (s *ExecuteTextbookAssistantDifficultyRequest) GetAuthToken() *string  {
  return s.AuthToken
}

func (s *ExecuteTextbookAssistantDifficultyRequest) GetChatId() *string  {
  return s.ChatId
}

func (s *ExecuteTextbookAssistantDifficultyRequest) GetScenario() *string  {
  return s.Scenario
}

func (s *ExecuteTextbookAssistantDifficultyRequest) SetAction(v string) *ExecuteTextbookAssistantDifficultyRequest {
  s.Action = &v
  return s
}

func (s *ExecuteTextbookAssistantDifficultyRequest) SetAssistant(v string) *ExecuteTextbookAssistantDifficultyRequest {
  s.Assistant = &v
  return s
}

func (s *ExecuteTextbookAssistantDifficultyRequest) SetAuthToken(v string) *ExecuteTextbookAssistantDifficultyRequest {
  s.AuthToken = &v
  return s
}

func (s *ExecuteTextbookAssistantDifficultyRequest) SetChatId(v string) *ExecuteTextbookAssistantDifficultyRequest {
  s.ChatId = &v
  return s
}

func (s *ExecuteTextbookAssistantDifficultyRequest) SetScenario(v string) *ExecuteTextbookAssistantDifficultyRequest {
  s.Scenario = &v
  return s
}

func (s *ExecuteTextbookAssistantDifficultyRequest) Validate() error {
  return dara.Validate(s)
}

