// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTextbookAssistantSuggestionRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAssistant(v string) *ExecuteTextbookAssistantSuggestionRequest
  GetAssistant() *string 
  SetAuthToken(v string) *ExecuteTextbookAssistantSuggestionRequest
  GetAuthToken() *string 
  SetChatId(v string) *ExecuteTextbookAssistantSuggestionRequest
  GetChatId() *string 
  SetScenario(v string) *ExecuteTextbookAssistantSuggestionRequest
  GetScenario() *string 
}

type ExecuteTextbookAssistantSuggestionRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // 6788e0b4b54c5268c1b78638
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
  // 6788e0b4b54c5268c1b78638
  ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // SYNC
  Scenario *string `json:"scenario,omitempty" xml:"scenario,omitempty"`
}

func (s ExecuteTextbookAssistantSuggestionRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantSuggestionRequest) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantSuggestionRequest) GetAssistant() *string  {
  return s.Assistant
}

func (s *ExecuteTextbookAssistantSuggestionRequest) GetAuthToken() *string  {
  return s.AuthToken
}

func (s *ExecuteTextbookAssistantSuggestionRequest) GetChatId() *string  {
  return s.ChatId
}

func (s *ExecuteTextbookAssistantSuggestionRequest) GetScenario() *string  {
  return s.Scenario
}

func (s *ExecuteTextbookAssistantSuggestionRequest) SetAssistant(v string) *ExecuteTextbookAssistantSuggestionRequest {
  s.Assistant = &v
  return s
}

func (s *ExecuteTextbookAssistantSuggestionRequest) SetAuthToken(v string) *ExecuteTextbookAssistantSuggestionRequest {
  s.AuthToken = &v
  return s
}

func (s *ExecuteTextbookAssistantSuggestionRequest) SetChatId(v string) *ExecuteTextbookAssistantSuggestionRequest {
  s.ChatId = &v
  return s
}

func (s *ExecuteTextbookAssistantSuggestionRequest) SetScenario(v string) *ExecuteTextbookAssistantSuggestionRequest {
  s.Scenario = &v
  return s
}

func (s *ExecuteTextbookAssistantSuggestionRequest) Validate() error {
  return dara.Validate(s)
}

