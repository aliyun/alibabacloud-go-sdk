// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTextbookAssistantGrammarCheckRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAuthToken(v string) *ExecuteTextbookAssistantGrammarCheckRequest
  GetAuthToken() *string 
  SetChatId(v string) *ExecuteTextbookAssistantGrammarCheckRequest
  GetChatId() *string 
  SetScenario(v string) *ExecuteTextbookAssistantGrammarCheckRequest
  GetScenario() *string 
  SetUser(v string) *ExecuteTextbookAssistantGrammarCheckRequest
  GetUser() *string 
}

type ExecuteTextbookAssistantGrammarCheckRequest struct {
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
  // This parameter is required.
  // 
  // example:
  // 
  // 6788e0b45bdfc807f077a5a1
  User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ExecuteTextbookAssistantGrammarCheckRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantGrammarCheckRequest) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantGrammarCheckRequest) GetAuthToken() *string  {
  return s.AuthToken
}

func (s *ExecuteTextbookAssistantGrammarCheckRequest) GetChatId() *string  {
  return s.ChatId
}

func (s *ExecuteTextbookAssistantGrammarCheckRequest) GetScenario() *string  {
  return s.Scenario
}

func (s *ExecuteTextbookAssistantGrammarCheckRequest) GetUser() *string  {
  return s.User
}

func (s *ExecuteTextbookAssistantGrammarCheckRequest) SetAuthToken(v string) *ExecuteTextbookAssistantGrammarCheckRequest {
  s.AuthToken = &v
  return s
}

func (s *ExecuteTextbookAssistantGrammarCheckRequest) SetChatId(v string) *ExecuteTextbookAssistantGrammarCheckRequest {
  s.ChatId = &v
  return s
}

func (s *ExecuteTextbookAssistantGrammarCheckRequest) SetScenario(v string) *ExecuteTextbookAssistantGrammarCheckRequest {
  s.Scenario = &v
  return s
}

func (s *ExecuteTextbookAssistantGrammarCheckRequest) SetUser(v string) *ExecuteTextbookAssistantGrammarCheckRequest {
  s.User = &v
  return s
}

func (s *ExecuteTextbookAssistantGrammarCheckRequest) Validate() error {
  return dara.Validate(s)
}

