// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTextbookAssistantRefineByContextRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAuthToken(v string) *ExecuteTextbookAssistantRefineByContextRequest
  GetAuthToken() *string 
  SetChatId(v string) *ExecuteTextbookAssistantRefineByContextRequest
  GetChatId() *string 
  SetScenario(v string) *ExecuteTextbookAssistantRefineByContextRequest
  GetScenario() *string 
  SetUser(v string) *ExecuteTextbookAssistantRefineByContextRequest
  GetUser() *string 
}

type ExecuteTextbookAssistantRefineByContextRequest struct {
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

func (s ExecuteTextbookAssistantRefineByContextRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantRefineByContextRequest) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantRefineByContextRequest) GetAuthToken() *string  {
  return s.AuthToken
}

func (s *ExecuteTextbookAssistantRefineByContextRequest) GetChatId() *string  {
  return s.ChatId
}

func (s *ExecuteTextbookAssistantRefineByContextRequest) GetScenario() *string  {
  return s.Scenario
}

func (s *ExecuteTextbookAssistantRefineByContextRequest) GetUser() *string  {
  return s.User
}

func (s *ExecuteTextbookAssistantRefineByContextRequest) SetAuthToken(v string) *ExecuteTextbookAssistantRefineByContextRequest {
  s.AuthToken = &v
  return s
}

func (s *ExecuteTextbookAssistantRefineByContextRequest) SetChatId(v string) *ExecuteTextbookAssistantRefineByContextRequest {
  s.ChatId = &v
  return s
}

func (s *ExecuteTextbookAssistantRefineByContextRequest) SetScenario(v string) *ExecuteTextbookAssistantRefineByContextRequest {
  s.Scenario = &v
  return s
}

func (s *ExecuteTextbookAssistantRefineByContextRequest) SetUser(v string) *ExecuteTextbookAssistantRefineByContextRequest {
  s.User = &v
  return s
}

func (s *ExecuteTextbookAssistantRefineByContextRequest) Validate() error {
  return dara.Validate(s)
}

