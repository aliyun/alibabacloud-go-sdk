// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTextbookAssistantSseDialogueRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAuthToken(v string) *ExecuteTextbookAssistantSseDialogueRequest
  GetAuthToken() *string 
  SetChatId(v string) *ExecuteTextbookAssistantSseDialogueRequest
  GetChatId() *string 
  SetScenario(v string) *ExecuteTextbookAssistantSseDialogueRequest
  GetScenario() *string 
  SetUserMessage(v string) *ExecuteTextbookAssistantSseDialogueRequest
  GetUserMessage() *string 
}

type ExecuteTextbookAssistantSseDialogueRequest struct {
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
  // 67e374acb54c526c95c4fbd4
  ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // EXPAND
  Scenario *string `json:"scenario,omitempty" xml:"scenario,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // hello
  UserMessage *string `json:"userMessage,omitempty" xml:"userMessage,omitempty"`
}

func (s ExecuteTextbookAssistantSseDialogueRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantSseDialogueRequest) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantSseDialogueRequest) GetAuthToken() *string  {
  return s.AuthToken
}

func (s *ExecuteTextbookAssistantSseDialogueRequest) GetChatId() *string  {
  return s.ChatId
}

func (s *ExecuteTextbookAssistantSseDialogueRequest) GetScenario() *string  {
  return s.Scenario
}

func (s *ExecuteTextbookAssistantSseDialogueRequest) GetUserMessage() *string  {
  return s.UserMessage
}

func (s *ExecuteTextbookAssistantSseDialogueRequest) SetAuthToken(v string) *ExecuteTextbookAssistantSseDialogueRequest {
  s.AuthToken = &v
  return s
}

func (s *ExecuteTextbookAssistantSseDialogueRequest) SetChatId(v string) *ExecuteTextbookAssistantSseDialogueRequest {
  s.ChatId = &v
  return s
}

func (s *ExecuteTextbookAssistantSseDialogueRequest) SetScenario(v string) *ExecuteTextbookAssistantSseDialogueRequest {
  s.Scenario = &v
  return s
}

func (s *ExecuteTextbookAssistantSseDialogueRequest) SetUserMessage(v string) *ExecuteTextbookAssistantSseDialogueRequest {
  s.UserMessage = &v
  return s
}

func (s *ExecuteTextbookAssistantSseDialogueRequest) Validate() error {
  return dara.Validate(s)
}

