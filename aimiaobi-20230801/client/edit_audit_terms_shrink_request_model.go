// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditAuditTermsShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetExceptionWordShrink(v string) *EditAuditTermsShrinkRequest
  GetExceptionWordShrink() *string 
  SetId(v string) *EditAuditTermsShrinkRequest
  GetId() *string 
  SetKeyword(v string) *EditAuditTermsShrinkRequest
  GetKeyword() *string 
  SetSuggestWord(v string) *EditAuditTermsShrinkRequest
  GetSuggestWord() *string 
  SetTermsDesc(v string) *EditAuditTermsShrinkRequest
  GetTermsDesc() *string 
  SetWorkspaceId(v string) *EditAuditTermsShrinkRequest
  GetWorkspaceId() *string 
}

type EditAuditTermsShrinkRequest struct {
  ExceptionWordShrink *string `json:"ExceptionWord,omitempty" xml:"ExceptionWord,omitempty"`
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

func (s EditAuditTermsShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s EditAuditTermsShrinkRequest) GoString() string {
  return s.String()
}

func (s *EditAuditTermsShrinkRequest) GetExceptionWordShrink() *string  {
  return s.ExceptionWordShrink
}

func (s *EditAuditTermsShrinkRequest) GetId() *string  {
  return s.Id
}

func (s *EditAuditTermsShrinkRequest) GetKeyword() *string  {
  return s.Keyword
}

func (s *EditAuditTermsShrinkRequest) GetSuggestWord() *string  {
  return s.SuggestWord
}

func (s *EditAuditTermsShrinkRequest) GetTermsDesc() *string  {
  return s.TermsDesc
}

func (s *EditAuditTermsShrinkRequest) GetWorkspaceId() *string  {
  return s.WorkspaceId
}

func (s *EditAuditTermsShrinkRequest) SetExceptionWordShrink(v string) *EditAuditTermsShrinkRequest {
  s.ExceptionWordShrink = &v
  return s
}

func (s *EditAuditTermsShrinkRequest) SetId(v string) *EditAuditTermsShrinkRequest {
  s.Id = &v
  return s
}

func (s *EditAuditTermsShrinkRequest) SetKeyword(v string) *EditAuditTermsShrinkRequest {
  s.Keyword = &v
  return s
}

func (s *EditAuditTermsShrinkRequest) SetSuggestWord(v string) *EditAuditTermsShrinkRequest {
  s.SuggestWord = &v
  return s
}

func (s *EditAuditTermsShrinkRequest) SetTermsDesc(v string) *EditAuditTermsShrinkRequest {
  s.TermsDesc = &v
  return s
}

func (s *EditAuditTermsShrinkRequest) SetWorkspaceId(v string) *EditAuditTermsShrinkRequest {
  s.WorkspaceId = &v
  return s
}

func (s *EditAuditTermsShrinkRequest) Validate() error {
  return dara.Validate(s)
}

