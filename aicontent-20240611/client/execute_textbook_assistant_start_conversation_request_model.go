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
  // How you obtain this ID depends on the value of `scenario`.
  // 
  // **When the `scenario` input parameter is `SYNC`:**
  // 
  // 1. From the `Get Article List` response, use the top-level `articleId` field.
  // 
  // 2. From the `Get Article Details` response, use the top-level `articleId` field.
  // 
  // **When the `scenario` input parameter is `EXPAND`:**
  // 
  // 1. From the `Get Article Details` response, use the `sceneid` field from an element in the `sceneList` array.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 0c05700d4d9411efbe6e0c42a106bb02
  ArticleId *string `json:"articleId,omitempty" xml:"articleId,omitempty"`
  // The authorization token for the API call. Obtain this token by calling the operation that provides the authorization token for the textbook-style AI teacher feature.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // tc_e6dc70c890866f4028ca685b6fa29874
  AuthToken *string `json:"authToken,omitempty" xml:"authToken,omitempty"`
  // The practice scenario. Valid values:
  // 
  // `SYNC`: synchronous practice
  // 
  // `EXPAND`: expansion practice
  // 
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

