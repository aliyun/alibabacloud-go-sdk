// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAuditTermsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExceptionWord(v []*string) *AddAuditTermsRequest
	GetExceptionWord() []*string
	SetKeyword(v string) *AddAuditTermsRequest
	GetKeyword() *string
	SetSuggestWord(v string) *AddAuditTermsRequest
	GetSuggestWord() *string
	SetTermsDesc(v string) *AddAuditTermsRequest
	GetTermsDesc() *string
	SetWorkspaceId(v string) *AddAuditTermsRequest
	GetWorkspaceId() *string
}

type AddAuditTermsRequest struct {
	ExceptionWord []*string `json:"ExceptionWord,omitempty" xml:"ExceptionWord,omitempty" type:"Repeated"`
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

func (s AddAuditTermsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAuditTermsRequest) GoString() string {
	return s.String()
}

func (s *AddAuditTermsRequest) GetExceptionWord() []*string {
	return s.ExceptionWord
}

func (s *AddAuditTermsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *AddAuditTermsRequest) GetSuggestWord() *string {
	return s.SuggestWord
}

func (s *AddAuditTermsRequest) GetTermsDesc() *string {
	return s.TermsDesc
}

func (s *AddAuditTermsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AddAuditTermsRequest) SetExceptionWord(v []*string) *AddAuditTermsRequest {
	s.ExceptionWord = v
	return s
}

func (s *AddAuditTermsRequest) SetKeyword(v string) *AddAuditTermsRequest {
	s.Keyword = &v
	return s
}

func (s *AddAuditTermsRequest) SetSuggestWord(v string) *AddAuditTermsRequest {
	s.SuggestWord = &v
	return s
}

func (s *AddAuditTermsRequest) SetTermsDesc(v string) *AddAuditTermsRequest {
	s.TermsDesc = &v
	return s
}

func (s *AddAuditTermsRequest) SetWorkspaceId(v string) *AddAuditTermsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AddAuditTermsRequest) Validate() error {
	return dara.Validate(s)
}
