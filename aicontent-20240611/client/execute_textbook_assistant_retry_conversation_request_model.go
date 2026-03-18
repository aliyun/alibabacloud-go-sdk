// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTextbookAssistantRetryConversationRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAssistant(v string) *ExecuteTextbookAssistantRetryConversationRequest
  GetAssistant() *string 
  SetAuthToken(v string) *ExecuteTextbookAssistantRetryConversationRequest
  GetAuthToken() *string 
  SetChatId(v string) *ExecuteTextbookAssistantRetryConversationRequest
  GetChatId() *string 
  SetScenario(v string) *ExecuteTextbookAssistantRetryConversationRequest
  GetScenario() *string 
}

type ExecuteTextbookAssistantRetryConversationRequest struct {
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
  // 6788e0b475a4631ffc626722
  ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // SYNC
  Scenario *string `json:"scenario,omitempty" xml:"scenario,omitempty"`
}

func (s ExecuteTextbookAssistantRetryConversationRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantRetryConversationRequest) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantRetryConversationRequest) GetAssistant() *string  {
  return s.Assistant
}

func (s *ExecuteTextbookAssistantRetryConversationRequest) GetAuthToken() *string  {
  return s.AuthToken
}

func (s *ExecuteTextbookAssistantRetryConversationRequest) GetChatId() *string  {
  return s.ChatId
}

func (s *ExecuteTextbookAssistantRetryConversationRequest) GetScenario() *string  {
  return s.Scenario
}

func (s *ExecuteTextbookAssistantRetryConversationRequest) SetAssistant(v string) *ExecuteTextbookAssistantRetryConversationRequest {
  s.Assistant = &v
  return s
}

func (s *ExecuteTextbookAssistantRetryConversationRequest) SetAuthToken(v string) *ExecuteTextbookAssistantRetryConversationRequest {
  s.AuthToken = &v
  return s
}

func (s *ExecuteTextbookAssistantRetryConversationRequest) SetChatId(v string) *ExecuteTextbookAssistantRetryConversationRequest {
  s.ChatId = &v
  return s
}

func (s *ExecuteTextbookAssistantRetryConversationRequest) SetScenario(v string) *ExecuteTextbookAssistantRetryConversationRequest {
  s.Scenario = &v
  return s
}

func (s *ExecuteTextbookAssistantRetryConversationRequest) Validate() error {
  return dara.Validate(s)
}

