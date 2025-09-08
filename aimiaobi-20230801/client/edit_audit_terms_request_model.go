// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditAuditTermsRequest interface {
  dara.Model
  String() string
  GoString() string
  SetExceptionWord(v []*string) *EditAuditTermsRequest
  GetExceptionWord() []*string 
  SetId(v string) *EditAuditTermsRequest
  GetId() *string 
  SetKeyword(v string) *EditAuditTermsRequest
  GetKeyword() *string 
  SetSuggestWord(v string) *EditAuditTermsRequest
  GetSuggestWord() *string 
  SetTermsDesc(v string) *EditAuditTermsRequest
  GetTermsDesc() *string 
  SetWorkspaceId(v string) *EditAuditTermsRequest
  GetWorkspaceId() *string 
}

type EditAuditTermsRequest struct {
  ExceptionWord []*string `json:"ExceptionWord,omitempty" xml:"ExceptionWord,omitempty" type:"Repeated"`
  // example:
  // 
  // 20103
  Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
  // example:
  // 
  // 龘
  Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
  // example:
  // 
  // 龘(dá)
  SuggestWord *string `json:"SuggestWord,omitempty" xml:"SuggestWord,omitempty"`
  // example:
  // 
  // 龙行龘龘出自四库本《玉篇》23龙部第8字，文字释义为群龙腾飞的样子，昂扬而热烈。
  TermsDesc *string `json:"TermsDesc,omitempty" xml:"TermsDesc,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // llm-xx
  WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s EditAuditTermsRequest) String() string {
  return dara.Prettify(s)
}

func (s EditAuditTermsRequest) GoString() string {
  return s.String()
}

func (s *EditAuditTermsRequest) GetExceptionWord() []*string  {
  return s.ExceptionWord
}

func (s *EditAuditTermsRequest) GetId() *string  {
  return s.Id
}

func (s *EditAuditTermsRequest) GetKeyword() *string  {
  return s.Keyword
}

func (s *EditAuditTermsRequest) GetSuggestWord() *string  {
  return s.SuggestWord
}

func (s *EditAuditTermsRequest) GetTermsDesc() *string  {
  return s.TermsDesc
}

func (s *EditAuditTermsRequest) GetWorkspaceId() *string  {
  return s.WorkspaceId
}

func (s *EditAuditTermsRequest) SetExceptionWord(v []*string) *EditAuditTermsRequest {
  s.ExceptionWord = v
  return s
}

func (s *EditAuditTermsRequest) SetId(v string) *EditAuditTermsRequest {
  s.Id = &v
  return s
}

func (s *EditAuditTermsRequest) SetKeyword(v string) *EditAuditTermsRequest {
  s.Keyword = &v
  return s
}

func (s *EditAuditTermsRequest) SetSuggestWord(v string) *EditAuditTermsRequest {
  s.SuggestWord = &v
  return s
}

func (s *EditAuditTermsRequest) SetTermsDesc(v string) *EditAuditTermsRequest {
  s.TermsDesc = &v
  return s
}

func (s *EditAuditTermsRequest) SetWorkspaceId(v string) *EditAuditTermsRequest {
  s.WorkspaceId = &v
  return s
}

func (s *EditAuditTermsRequest) Validate() error {
  return dara.Validate(s)
}

