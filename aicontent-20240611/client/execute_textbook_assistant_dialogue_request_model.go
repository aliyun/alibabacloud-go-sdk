// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTextbookAssistantDialogueRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAuthToken(v string) *ExecuteTextbookAssistantDialogueRequest
  GetAuthToken() *string 
  SetChatId(v string) *ExecuteTextbookAssistantDialogueRequest
  GetChatId() *string 
  SetScenario(v string) *ExecuteTextbookAssistantDialogueRequest
  GetScenario() *string 
  SetUserMessage(v string) *ExecuteTextbookAssistantDialogueRequest
  GetUserMessage() *string 
}

type ExecuteTextbookAssistantDialogueRequest struct {
  // This parameter is required.
  AuthToken *string `json:"authToken,omitempty" xml:"authToken,omitempty"`
  // This parameter is required.
  ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
  // This parameter is required.
  Scenario *string `json:"scenario,omitempty" xml:"scenario,omitempty"`
  // This parameter is required.
  UserMessage *string `json:"userMessage,omitempty" xml:"userMessage,omitempty"`
}

func (s ExecuteTextbookAssistantDialogueRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantDialogueRequest) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantDialogueRequest) GetAuthToken() *string  {
  return s.AuthToken
}

func (s *ExecuteTextbookAssistantDialogueRequest) GetChatId() *string  {
  return s.ChatId
}

func (s *ExecuteTextbookAssistantDialogueRequest) GetScenario() *string  {
  return s.Scenario
}

func (s *ExecuteTextbookAssistantDialogueRequest) GetUserMessage() *string  {
  return s.UserMessage
}

func (s *ExecuteTextbookAssistantDialogueRequest) SetAuthToken(v string) *ExecuteTextbookAssistantDialogueRequest {
  s.AuthToken = &v
  return s
}

func (s *ExecuteTextbookAssistantDialogueRequest) SetChatId(v string) *ExecuteTextbookAssistantDialogueRequest {
  s.ChatId = &v
  return s
}

func (s *ExecuteTextbookAssistantDialogueRequest) SetScenario(v string) *ExecuteTextbookAssistantDialogueRequest {
  s.Scenario = &v
  return s
}

func (s *ExecuteTextbookAssistantDialogueRequest) SetUserMessage(v string) *ExecuteTextbookAssistantDialogueRequest {
  s.UserMessage = &v
  return s
}

func (s *ExecuteTextbookAssistantDialogueRequest) Validate() error {
  return dara.Validate(s)
}

