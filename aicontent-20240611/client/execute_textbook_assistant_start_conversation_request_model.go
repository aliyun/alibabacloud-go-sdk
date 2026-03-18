// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTextbookAssistantStartConversationRequest interface {
  dara.Model
  String() string
  GoString() string
  SetArticleId(v string) *ExecuteTextbookAssistantStartConversationRequest
  GetArticleId() *string 
  SetAuthToken(v string) *ExecuteTextbookAssistantStartConversationRequest
  GetAuthToken() *string 
  SetScenario(v string) *ExecuteTextbookAssistantStartConversationRequest
  GetScenario() *string 
}

type ExecuteTextbookAssistantStartConversationRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // 0c05700d4d9411efbe6e0c42a106bb02
  ArticleId *string `json:"articleId,omitempty" xml:"articleId,omitempty"`
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
  // SYNC
  Scenario *string `json:"scenario,omitempty" xml:"scenario,omitempty"`
}

func (s ExecuteTextbookAssistantStartConversationRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTextbookAssistantStartConversationRequest) GoString() string {
  return s.String()
}

func (s *ExecuteTextbookAssistantStartConversationRequest) GetArticleId() *string  {
  return s.ArticleId
}

func (s *ExecuteTextbookAssistantStartConversationRequest) GetAuthToken() *string  {
  return s.AuthToken
}

func (s *ExecuteTextbookAssistantStartConversationRequest) GetScenario() *string  {
  return s.Scenario
}

func (s *ExecuteTextbookAssistantStartConversationRequest) SetArticleId(v string) *ExecuteTextbookAssistantStartConversationRequest {
  s.ArticleId = &v
  return s
}

func (s *ExecuteTextbookAssistantStartConversationRequest) SetAuthToken(v string) *ExecuteTextbookAssistantStartConversationRequest {
  s.AuthToken = &v
  return s
}

func (s *ExecuteTextbookAssistantStartConversationRequest) SetScenario(v string) *ExecuteTextbookAssistantStartConversationRequest {
  s.Scenario = &v
  return s
}

func (s *ExecuteTextbookAssistantStartConversationRequest) Validate() error {
  return dara.Validate(s)
}

