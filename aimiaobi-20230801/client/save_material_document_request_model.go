// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveMaterialDocumentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *SaveMaterialDocumentRequest
	GetAgentKey() *string
	SetAuthor(v string) *SaveMaterialDocumentRequest
	GetAuthor() *string
	SetBothSavePrivateAndShare(v bool) *SaveMaterialDocumentRequest
	GetBothSavePrivateAndShare() *bool
	SetDocKeywords(v []*string) *SaveMaterialDocumentRequest
	GetDocKeywords() []*string
	SetDocType(v string) *SaveMaterialDocumentRequest
	GetDocType() *string
	SetExternalUrl(v string) *SaveMaterialDocumentRequest
	GetExternalUrl() *string
	SetHtmlContent(v string) *SaveMaterialDocumentRequest
	GetHtmlContent() *string
	SetPubTime(v string) *SaveMaterialDocumentRequest
	GetPubTime() *string
	SetShareAttr(v int32) *SaveMaterialDocumentRequest
	GetShareAttr() *int32
	SetSrcFrom(v string) *SaveMaterialDocumentRequest
	GetSrcFrom() *string
	SetSummary(v string) *SaveMaterialDocumentRequest
	GetSummary() *string
	SetTextContent(v string) *SaveMaterialDocumentRequest
	GetTextContent() *string
	SetTitle(v string) *SaveMaterialDocumentRequest
	GetTitle() *string
	SetUrl(v string) *SaveMaterialDocumentRequest
	GetUrl() *string
}

type SaveMaterialDocumentRequest struct {
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
	BothSavePrivateAndShare *bool     `json:"BothSavePrivateAndShare,omitempty" xml:"BothSavePrivateAndShare,omitempty"`
	DocKeywords             []*string `json:"DocKeywords,omitempty" xml:"DocKeywords,omitempty" type:"Repeated"`
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

func (s SaveMaterialDocumentRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveMaterialDocumentRequest) GoString() string {
	return s.String()
}

func (s *SaveMaterialDocumentRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *SaveMaterialDocumentRequest) GetAuthor() *string {
	return s.Author
}

func (s *SaveMaterialDocumentRequest) GetBothSavePrivateAndShare() *bool {
	return s.BothSavePrivateAndShare
}

func (s *SaveMaterialDocumentRequest) GetDocKeywords() []*string {
	return s.DocKeywords
}

func (s *SaveMaterialDocumentRequest) GetDocType() *string {
	return s.DocType
}

func (s *SaveMaterialDocumentRequest) GetExternalUrl() *string {
	return s.ExternalUrl
}

func (s *SaveMaterialDocumentRequest) GetHtmlContent() *string {
	return s.HtmlContent
}

func (s *SaveMaterialDocumentRequest) GetPubTime() *string {
	return s.PubTime
}

func (s *SaveMaterialDocumentRequest) GetShareAttr() *int32 {
	return s.ShareAttr
}

func (s *SaveMaterialDocumentRequest) GetSrcFrom() *string {
	return s.SrcFrom
}

func (s *SaveMaterialDocumentRequest) GetSummary() *string {
	return s.Summary
}

func (s *SaveMaterialDocumentRequest) GetTextContent() *string {
	return s.TextContent
}

func (s *SaveMaterialDocumentRequest) GetTitle() *string {
	return s.Title
}

func (s *SaveMaterialDocumentRequest) GetUrl() *string {
	return s.Url
}

func (s *SaveMaterialDocumentRequest) SetAgentKey(v string) *SaveMaterialDocumentRequest {
	s.AgentKey = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetAuthor(v string) *SaveMaterialDocumentRequest {
	s.Author = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetBothSavePrivateAndShare(v bool) *SaveMaterialDocumentRequest {
	s.BothSavePrivateAndShare = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetDocKeywords(v []*string) *SaveMaterialDocumentRequest {
	s.DocKeywords = v
	return s
}

func (s *SaveMaterialDocumentRequest) SetDocType(v string) *SaveMaterialDocumentRequest {
	s.DocType = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetExternalUrl(v string) *SaveMaterialDocumentRequest {
	s.ExternalUrl = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetHtmlContent(v string) *SaveMaterialDocumentRequest {
	s.HtmlContent = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetPubTime(v string) *SaveMaterialDocumentRequest {
	s.PubTime = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetShareAttr(v int32) *SaveMaterialDocumentRequest {
	s.ShareAttr = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetSrcFrom(v string) *SaveMaterialDocumentRequest {
	s.SrcFrom = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetSummary(v string) *SaveMaterialDocumentRequest {
	s.Summary = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetTextContent(v string) *SaveMaterialDocumentRequest {
	s.TextContent = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetTitle(v string) *SaveMaterialDocumentRequest {
	s.Title = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetUrl(v string) *SaveMaterialDocumentRequest {
	s.Url = &v
	return s
}

func (s *SaveMaterialDocumentRequest) Validate() error {
	return dara.Validate(s)
}
