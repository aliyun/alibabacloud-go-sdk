// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAuditTermsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExceptionWordShrink(v string) *AddAuditTermsShrinkRequest
	GetExceptionWordShrink() *string
	SetKeyword(v string) *AddAuditTermsShrinkRequest
	GetKeyword() *string
	SetSuggestWord(v string) *AddAuditTermsShrinkRequest
	GetSuggestWord() *string
	SetTermsDesc(v string) *AddAuditTermsShrinkRequest
	GetTermsDesc() *string
	SetWorkspaceId(v string) *AddAuditTermsShrinkRequest
	GetWorkspaceId() *string
}

type AddAuditTermsShrinkRequest struct {
	ExceptionWordShrink *string `json:"ExceptionWord,omitempty" xml:"ExceptionWord,omitempty"`
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

func (s AddAuditTermsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAuditTermsShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddAuditTermsShrinkRequest) GetExceptionWordShrink() *string {
	return s.ExceptionWordShrink
}

func (s *AddAuditTermsShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *AddAuditTermsShrinkRequest) GetSuggestWord() *string {
	return s.SuggestWord
}

func (s *AddAuditTermsShrinkRequest) GetTermsDesc() *string {
	return s.TermsDesc
}

func (s *AddAuditTermsShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AddAuditTermsShrinkRequest) SetExceptionWordShrink(v string) *AddAuditTermsShrinkRequest {
	s.ExceptionWordShrink = &v
	return s
}

func (s *AddAuditTermsShrinkRequest) SetKeyword(v string) *AddAuditTermsShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *AddAuditTermsShrinkRequest) SetSuggestWord(v string) *AddAuditTermsShrinkRequest {
	s.SuggestWord = &v
	return s
}

func (s *AddAuditTermsShrinkRequest) SetTermsDesc(v string) *AddAuditTermsShrinkRequest {
	s.TermsDesc = &v
	return s
}

func (s *AddAuditTermsShrinkRequest) SetWorkspaceId(v string) *AddAuditTermsShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AddAuditTermsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
