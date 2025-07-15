// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveMaterialDocumentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *SaveMaterialDocumentShrinkRequest
	GetAgentKey() *string
	SetAuthor(v string) *SaveMaterialDocumentShrinkRequest
	GetAuthor() *string
	SetBothSavePrivateAndShare(v bool) *SaveMaterialDocumentShrinkRequest
	GetBothSavePrivateAndShare() *bool
	SetDocKeywordsShrink(v string) *SaveMaterialDocumentShrinkRequest
	GetDocKeywordsShrink() *string
	SetDocType(v string) *SaveMaterialDocumentShrinkRequest
	GetDocType() *string
	SetExternalUrl(v string) *SaveMaterialDocumentShrinkRequest
	GetExternalUrl() *string
	SetHtmlContent(v string) *SaveMaterialDocumentShrinkRequest
	GetHtmlContent() *string
	SetPubTime(v string) *SaveMaterialDocumentShrinkRequest
	GetPubTime() *string
	SetShareAttr(v int32) *SaveMaterialDocumentShrinkRequest
	GetShareAttr() *int32
	SetSrcFrom(v string) *SaveMaterialDocumentShrinkRequest
	GetSrcFrom() *string
	SetSummary(v string) *SaveMaterialDocumentShrinkRequest
	GetSummary() *string
	SetTextContent(v string) *SaveMaterialDocumentShrinkRequest
	GetTextContent() *string
	SetTitle(v string) *SaveMaterialDocumentShrinkRequest
	GetTitle() *string
	SetUrl(v string) *SaveMaterialDocumentShrinkRequest
	GetUrl() *string
}

type SaveMaterialDocumentShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c160c841c8e54295bf2f441432785944_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Author   *string `json:"Author,omitempty" xml:"Author,omitempty"`
	// example:
	//
	// false
	BothSavePrivateAndShare *bool   `json:"BothSavePrivateAndShare,omitempty" xml:"BothSavePrivateAndShare,omitempty"`
	DocKeywordsShrink       *string `json:"DocKeywords,omitempty" xml:"DocKeywords,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// excel
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// example:
	//
	// http://xxxxx/xxx
	ExternalUrl *string `json:"ExternalUrl,omitempty" xml:"ExternalUrl,omitempty"`
	HtmlContent *string `json:"HtmlContent,omitempty" xml:"HtmlContent,omitempty"`
	// example:
	//
	// 2023-04-11 06:14:07
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// 1
	ShareAttr *int32 `json:"ShareAttr,omitempty" xml:"ShareAttr,omitempty"`
	// example:
	//
	// IntellijSearch
	SrcFrom     *string `json:"SrcFrom,omitempty" xml:"SrcFrom,omitempty"`
	Summary     *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	TextContent *string `json:"TextContent,omitempty" xml:"TextContent,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// http://xxxxx/xxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SaveMaterialDocumentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveMaterialDocumentShrinkRequest) GoString() string {
	return s.String()
}

func (s *SaveMaterialDocumentShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *SaveMaterialDocumentShrinkRequest) GetAuthor() *string {
	return s.Author
}

func (s *SaveMaterialDocumentShrinkRequest) GetBothSavePrivateAndShare() *bool {
	return s.BothSavePrivateAndShare
}

func (s *SaveMaterialDocumentShrinkRequest) GetDocKeywordsShrink() *string {
	return s.DocKeywordsShrink
}

func (s *SaveMaterialDocumentShrinkRequest) GetDocType() *string {
	return s.DocType
}

func (s *SaveMaterialDocumentShrinkRequest) GetExternalUrl() *string {
	return s.ExternalUrl
}

func (s *SaveMaterialDocumentShrinkRequest) GetHtmlContent() *string {
	return s.HtmlContent
}

func (s *SaveMaterialDocumentShrinkRequest) GetPubTime() *string {
	return s.PubTime
}

func (s *SaveMaterialDocumentShrinkRequest) GetShareAttr() *int32 {
	return s.ShareAttr
}

func (s *SaveMaterialDocumentShrinkRequest) GetSrcFrom() *string {
	return s.SrcFrom
}

func (s *SaveMaterialDocumentShrinkRequest) GetSummary() *string {
	return s.Summary
}

func (s *SaveMaterialDocumentShrinkRequest) GetTextContent() *string {
	return s.TextContent
}

func (s *SaveMaterialDocumentShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *SaveMaterialDocumentShrinkRequest) GetUrl() *string {
	return s.Url
}

func (s *SaveMaterialDocumentShrinkRequest) SetAgentKey(v string) *SaveMaterialDocumentShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetAuthor(v string) *SaveMaterialDocumentShrinkRequest {
	s.Author = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetBothSavePrivateAndShare(v bool) *SaveMaterialDocumentShrinkRequest {
	s.BothSavePrivateAndShare = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetDocKeywordsShrink(v string) *SaveMaterialDocumentShrinkRequest {
	s.DocKeywordsShrink = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetDocType(v string) *SaveMaterialDocumentShrinkRequest {
	s.DocType = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetExternalUrl(v string) *SaveMaterialDocumentShrinkRequest {
	s.ExternalUrl = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetHtmlContent(v string) *SaveMaterialDocumentShrinkRequest {
	s.HtmlContent = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetPubTime(v string) *SaveMaterialDocumentShrinkRequest {
	s.PubTime = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetShareAttr(v int32) *SaveMaterialDocumentShrinkRequest {
	s.ShareAttr = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetSrcFrom(v string) *SaveMaterialDocumentShrinkRequest {
	s.SrcFrom = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetSummary(v string) *SaveMaterialDocumentShrinkRequest {
	s.Summary = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetTextContent(v string) *SaveMaterialDocumentShrinkRequest {
	s.TextContent = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetTitle(v string) *SaveMaterialDocumentShrinkRequest {
	s.Title = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetUrl(v string) *SaveMaterialDocumentShrinkRequest {
	s.Url = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
