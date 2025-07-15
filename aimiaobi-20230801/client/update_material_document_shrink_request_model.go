// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMaterialDocumentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *UpdateMaterialDocumentShrinkRequest
	GetAgentKey() *string
	SetAuthor(v string) *UpdateMaterialDocumentShrinkRequest
	GetAuthor() *string
	SetDocKeywordsShrink(v string) *UpdateMaterialDocumentShrinkRequest
	GetDocKeywordsShrink() *string
	SetDocType(v string) *UpdateMaterialDocumentShrinkRequest
	GetDocType() *string
	SetExternalUrl(v string) *UpdateMaterialDocumentShrinkRequest
	GetExternalUrl() *string
	SetHtmlContent(v string) *UpdateMaterialDocumentShrinkRequest
	GetHtmlContent() *string
	SetId(v int64) *UpdateMaterialDocumentShrinkRequest
	GetId() *int64
	SetPubTime(v string) *UpdateMaterialDocumentShrinkRequest
	GetPubTime() *string
	SetRegionId(v string) *UpdateMaterialDocumentShrinkRequest
	GetRegionId() *string
	SetShareAttr(v int32) *UpdateMaterialDocumentShrinkRequest
	GetShareAttr() *int32
	SetSrcFrom(v string) *UpdateMaterialDocumentShrinkRequest
	GetSrcFrom() *string
	SetSummary(v string) *UpdateMaterialDocumentShrinkRequest
	GetSummary() *string
	SetTextContent(v string) *UpdateMaterialDocumentShrinkRequest
	GetTextContent() *string
	SetTitle(v string) *UpdateMaterialDocumentShrinkRequest
	GetTitle() *string
	SetUrl(v string) *UpdateMaterialDocumentShrinkRequest
	GetUrl() *string
}

type UpdateMaterialDocumentShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 67c520d1fa43455ea44fb69fa402d54d_p_beebot_public
	AgentKey          *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Author            *string `json:"Author,omitempty" xml:"Author,omitempty"`
	DocKeywordsShrink *string `json:"DocKeywords,omitempty" xml:"DocKeywords,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// image
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// example:
	//
	// http://xxxxx/xxx
	ExternalUrl *string `json:"ExternalUrl,omitempty" xml:"ExternalUrl,omitempty"`
	HtmlContent *string `json:"HtmlContent,omitempty" xml:"HtmlContent,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 44
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2023-04-11 06:14:07
	PubTime  *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 1
	ShareAttr *int32 `json:"ShareAttr,omitempty" xml:"ShareAttr,omitempty"`
	// example:
	//
	// UserUpload
	SrcFrom     *string `json:"SrcFrom,omitempty" xml:"SrcFrom,omitempty"`
	Summary     *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	TextContent *string `json:"TextContent,omitempty" xml:"TextContent,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// http://xxxxx/xxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s UpdateMaterialDocumentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMaterialDocumentShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMaterialDocumentShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *UpdateMaterialDocumentShrinkRequest) GetAuthor() *string {
	return s.Author
}

func (s *UpdateMaterialDocumentShrinkRequest) GetDocKeywordsShrink() *string {
	return s.DocKeywordsShrink
}

func (s *UpdateMaterialDocumentShrinkRequest) GetDocType() *string {
	return s.DocType
}

func (s *UpdateMaterialDocumentShrinkRequest) GetExternalUrl() *string {
	return s.ExternalUrl
}

func (s *UpdateMaterialDocumentShrinkRequest) GetHtmlContent() *string {
	return s.HtmlContent
}

func (s *UpdateMaterialDocumentShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateMaterialDocumentShrinkRequest) GetPubTime() *string {
	return s.PubTime
}

func (s *UpdateMaterialDocumentShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateMaterialDocumentShrinkRequest) GetShareAttr() *int32 {
	return s.ShareAttr
}

func (s *UpdateMaterialDocumentShrinkRequest) GetSrcFrom() *string {
	return s.SrcFrom
}

func (s *UpdateMaterialDocumentShrinkRequest) GetSummary() *string {
	return s.Summary
}

func (s *UpdateMaterialDocumentShrinkRequest) GetTextContent() *string {
	return s.TextContent
}

func (s *UpdateMaterialDocumentShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateMaterialDocumentShrinkRequest) GetUrl() *string {
	return s.Url
}

func (s *UpdateMaterialDocumentShrinkRequest) SetAgentKey(v string) *UpdateMaterialDocumentShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetAuthor(v string) *UpdateMaterialDocumentShrinkRequest {
	s.Author = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetDocKeywordsShrink(v string) *UpdateMaterialDocumentShrinkRequest {
	s.DocKeywordsShrink = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetDocType(v string) *UpdateMaterialDocumentShrinkRequest {
	s.DocType = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetExternalUrl(v string) *UpdateMaterialDocumentShrinkRequest {
	s.ExternalUrl = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetHtmlContent(v string) *UpdateMaterialDocumentShrinkRequest {
	s.HtmlContent = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetId(v int64) *UpdateMaterialDocumentShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetPubTime(v string) *UpdateMaterialDocumentShrinkRequest {
	s.PubTime = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetRegionId(v string) *UpdateMaterialDocumentShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetShareAttr(v int32) *UpdateMaterialDocumentShrinkRequest {
	s.ShareAttr = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetSrcFrom(v string) *UpdateMaterialDocumentShrinkRequest {
	s.SrcFrom = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetSummary(v string) *UpdateMaterialDocumentShrinkRequest {
	s.Summary = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetTextContent(v string) *UpdateMaterialDocumentShrinkRequest {
	s.TextContent = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetTitle(v string) *UpdateMaterialDocumentShrinkRequest {
	s.Title = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetUrl(v string) *UpdateMaterialDocumentShrinkRequest {
	s.Url = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
