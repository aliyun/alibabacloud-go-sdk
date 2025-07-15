// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMaterialDocumentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *UpdateMaterialDocumentRequest
	GetAgentKey() *string
	SetAuthor(v string) *UpdateMaterialDocumentRequest
	GetAuthor() *string
	SetDocKeywords(v []*string) *UpdateMaterialDocumentRequest
	GetDocKeywords() []*string
	SetDocType(v string) *UpdateMaterialDocumentRequest
	GetDocType() *string
	SetExternalUrl(v string) *UpdateMaterialDocumentRequest
	GetExternalUrl() *string
	SetHtmlContent(v string) *UpdateMaterialDocumentRequest
	GetHtmlContent() *string
	SetId(v int64) *UpdateMaterialDocumentRequest
	GetId() *int64
	SetPubTime(v string) *UpdateMaterialDocumentRequest
	GetPubTime() *string
	SetRegionId(v string) *UpdateMaterialDocumentRequest
	GetRegionId() *string
	SetShareAttr(v int32) *UpdateMaterialDocumentRequest
	GetShareAttr() *int32
	SetSrcFrom(v string) *UpdateMaterialDocumentRequest
	GetSrcFrom() *string
	SetSummary(v string) *UpdateMaterialDocumentRequest
	GetSummary() *string
	SetTextContent(v string) *UpdateMaterialDocumentRequest
	GetTextContent() *string
	SetTitle(v string) *UpdateMaterialDocumentRequest
	GetTitle() *string
	SetUrl(v string) *UpdateMaterialDocumentRequest
	GetUrl() *string
}

type UpdateMaterialDocumentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 67c520d1fa43455ea44fb69fa402d54d_p_beebot_public
	AgentKey    *string   `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Author      *string   `json:"Author,omitempty" xml:"Author,omitempty"`
	DocKeywords []*string `json:"DocKeywords,omitempty" xml:"DocKeywords,omitempty" type:"Repeated"`
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

func (s UpdateMaterialDocumentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMaterialDocumentRequest) GoString() string {
	return s.String()
}

func (s *UpdateMaterialDocumentRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *UpdateMaterialDocumentRequest) GetAuthor() *string {
	return s.Author
}

func (s *UpdateMaterialDocumentRequest) GetDocKeywords() []*string {
	return s.DocKeywords
}

func (s *UpdateMaterialDocumentRequest) GetDocType() *string {
	return s.DocType
}

func (s *UpdateMaterialDocumentRequest) GetExternalUrl() *string {
	return s.ExternalUrl
}

func (s *UpdateMaterialDocumentRequest) GetHtmlContent() *string {
	return s.HtmlContent
}

func (s *UpdateMaterialDocumentRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateMaterialDocumentRequest) GetPubTime() *string {
	return s.PubTime
}

func (s *UpdateMaterialDocumentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateMaterialDocumentRequest) GetShareAttr() *int32 {
	return s.ShareAttr
}

func (s *UpdateMaterialDocumentRequest) GetSrcFrom() *string {
	return s.SrcFrom
}

func (s *UpdateMaterialDocumentRequest) GetSummary() *string {
	return s.Summary
}

func (s *UpdateMaterialDocumentRequest) GetTextContent() *string {
	return s.TextContent
}

func (s *UpdateMaterialDocumentRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateMaterialDocumentRequest) GetUrl() *string {
	return s.Url
}

func (s *UpdateMaterialDocumentRequest) SetAgentKey(v string) *UpdateMaterialDocumentRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetAuthor(v string) *UpdateMaterialDocumentRequest {
	s.Author = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetDocKeywords(v []*string) *UpdateMaterialDocumentRequest {
	s.DocKeywords = v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetDocType(v string) *UpdateMaterialDocumentRequest {
	s.DocType = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetExternalUrl(v string) *UpdateMaterialDocumentRequest {
	s.ExternalUrl = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetHtmlContent(v string) *UpdateMaterialDocumentRequest {
	s.HtmlContent = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetId(v int64) *UpdateMaterialDocumentRequest {
	s.Id = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetPubTime(v string) *UpdateMaterialDocumentRequest {
	s.PubTime = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetRegionId(v string) *UpdateMaterialDocumentRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetShareAttr(v int32) *UpdateMaterialDocumentRequest {
	s.ShareAttr = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetSrcFrom(v string) *UpdateMaterialDocumentRequest {
	s.SrcFrom = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetSummary(v string) *UpdateMaterialDocumentRequest {
	s.Summary = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetTextContent(v string) *UpdateMaterialDocumentRequest {
	s.TextContent = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetTitle(v string) *UpdateMaterialDocumentRequest {
	s.Title = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetUrl(v string) *UpdateMaterialDocumentRequest {
	s.Url = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) Validate() error {
	return dara.Validate(s)
}
