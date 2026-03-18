// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTextbookAssistantTranslateRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAssistant(v string) *ExecuteTextbookAssistantTranslateRequest
  GetAssistant() *string 
  SetAuthToken(v string) *ExecuteTextbookAssistantTranslateRequest
  GetAuthToken() *string 
  SetChatId(v string) *ExecuteTextbookAssistantTranslateRequest
  GetChatId() *string 
  SetScenario(v string) *ExecuteTextbookAssistantTranslateRequest
  GetScenario() *string 
}

type ExecuteTextbookAssistantTranslateRequest struct {
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

func (s ExecuteTextbookAssistantTranslateRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantTranslateRequest) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantTranslateRequest) GetAssistant() *string  {
  return s.Assistant
}

func (s *ExecuteTextbookAssistantTranslateRequest) GetAuthToken() *string  {
  return s.AuthToken
}

func (s *ExecuteTextbookAssistantTranslateRequest) GetChatId() *string  {
  return s.ChatId
}

func (s *ExecuteTextbookAssistantTranslateRequest) GetScenario() *string  {
  return s.Scenario
}

func (s *ExecuteTextbookAssistantTranslateRequest) SetAssistant(v string) *ExecuteTextbookAssistantTranslateRequest {
  s.Assistant = &v
  return s
}

func (s *ExecuteTextbookAssistantTranslateRequest) SetAuthToken(v string) *ExecuteTextbookAssistantTranslateRequest {
  s.AuthToken = &v
  return s
}

func (s *ExecuteTextbookAssistantTranslateRequest) SetChatId(v string) *ExecuteTextbookAssistantTranslateRequest {
  s.ChatId = &v
  return s
}

func (s *ExecuteTextbookAssistantTranslateRequest) SetScenario(v string) *ExecuteTextbookAssistantTranslateRequest {
  s.Scenario = &v
  return s
}

func (s *ExecuteTextbookAssistantTranslateRequest) Validate() error {
  return dara.Validate(s)
}

