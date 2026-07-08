// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMaterialDocumentsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListMaterialDocumentsShrinkRequest
	GetAgentKey() *string
	SetContent(v string) *ListMaterialDocumentsShrinkRequest
	GetContent() *string
	SetCreateTimeEnd(v string) *ListMaterialDocumentsShrinkRequest
	GetCreateTimeEnd() *string
	SetCreateTimeStart(v string) *ListMaterialDocumentsShrinkRequest
	GetCreateTimeStart() *string
	SetCurrent(v int32) *ListMaterialDocumentsShrinkRequest
	GetCurrent() *int32
	SetDocType(v string) *ListMaterialDocumentsShrinkRequest
	GetDocType() *string
	SetDocTypeListShrink(v string) *ListMaterialDocumentsShrinkRequest
	GetDocTypeListShrink() *string
	SetGeneratePublicUrl(v bool) *ListMaterialDocumentsShrinkRequest
	GetGeneratePublicUrl() *bool
	SetId(v int64) *ListMaterialDocumentsShrinkRequest
	GetId() *int64
	SetKeywordsShrink(v string) *ListMaterialDocumentsShrinkRequest
	GetKeywordsShrink() *string
	SetQuery(v string) *ListMaterialDocumentsShrinkRequest
	GetQuery() *string
	SetShareAttr(v int32) *ListMaterialDocumentsShrinkRequest
	GetShareAttr() *int32
	SetSize(v int32) *ListMaterialDocumentsShrinkRequest
	GetSize() *int32
	SetTitle(v string) *ListMaterialDocumentsShrinkRequest
	GetTitle() *string
	SetUpdateTimeEnd(v string) *ListMaterialDocumentsShrinkRequest
	GetUpdateTimeEnd() *string
	SetUpdateTimeStart(v string) *ListMaterialDocumentsShrinkRequest
	GetUpdateTimeStart() *string
}

type ListMaterialDocumentsShrinkRequest struct {
	// Unique identifier for the workspace: [AgentKey](https://help.aliyun.com/document_detail/2587494.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// 33a2658aaabf4c24b45d50e575125311_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// Full-text search of document content
	//
	// example:
	//
	// 新闻内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Creation time - end range. Format: yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2023-03-18 02:00:00
	CreateTimeEnd *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// Creation time - start range. Format: yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2023-02-19 07:28:11
	CreateTimeStart *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	Current *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	// Document type (html: web page, plainText: plain text, image: image, pdf: pdf, word: word, excel: excel, csv: csv, jsonLine: jsonLine)
	//
	// example:
	//
	// jsonLine
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// Document type list (html: web page, plainText: plain text, image: image, pdf: pdf, word: word, excel: excel, csv: csv, jsonLine: jsonLine)
	//
	// example:
	//
	// excel
	DocTypeListShrink *string `json:"DocTypeList,omitempty" xml:"DocTypeList,omitempty"`
	// Generate a public URL for the file.
	//
	// example:
	//
	// true
	GeneratePublicUrl *bool `json:"GeneratePublicUrl,omitempty" xml:"GeneratePublicUrl,omitempty"`
	// Material primary key ID.
	//
	// example:
	//
	// 69
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Document keywords.
	KeywordsShrink *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// Supports comprehensive queries for title, content, and content.
	//
	// example:
	//
	// 小猫
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// Share property: 0: personal private, 1: shared within the workspace.
	//
	// example:
	//
	// 1
	ShareAttr *int32 `json:"ShareAttr,omitempty" xml:"ShareAttr,omitempty"`
	// Number of items per page. Default is 10.
	//
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// Performs a full-text search on material titles.
	//
	// example:
	//
	// 新闻标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// Update time - end range. Format: yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2023-03-18 03:00:00
	UpdateTimeEnd *string `json:"UpdateTimeEnd,omitempty" xml:"UpdateTimeEnd,omitempty"`
	// Update time - start range. Format: yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2023-03-18 02:00:00
	UpdateTimeStart *string `json:"UpdateTimeStart,omitempty" xml:"UpdateTimeStart,omitempty"`
}

func (s ListMaterialDocumentsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMaterialDocumentsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListMaterialDocumentsShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListMaterialDocumentsShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *ListMaterialDocumentsShrinkRequest) GetCreateTimeEnd() *string {
	return s.CreateTimeEnd
}

func (s *ListMaterialDocumentsShrinkRequest) GetCreateTimeStart() *string {
	return s.CreateTimeStart
}

func (s *ListMaterialDocumentsShrinkRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *ListMaterialDocumentsShrinkRequest) GetDocType() *string {
	return s.DocType
}

func (s *ListMaterialDocumentsShrinkRequest) GetDocTypeListShrink() *string {
	return s.DocTypeListShrink
}

func (s *ListMaterialDocumentsShrinkRequest) GetGeneratePublicUrl() *bool {
	return s.GeneratePublicUrl
}

func (s *ListMaterialDocumentsShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *ListMaterialDocumentsShrinkRequest) GetKeywordsShrink() *string {
	return s.KeywordsShrink
}

func (s *ListMaterialDocumentsShrinkRequest) GetQuery() *string {
	return s.Query
}

func (s *ListMaterialDocumentsShrinkRequest) GetShareAttr() *int32 {
	return s.ShareAttr
}

func (s *ListMaterialDocumentsShrinkRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListMaterialDocumentsShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *ListMaterialDocumentsShrinkRequest) GetUpdateTimeEnd() *string {
	return s.UpdateTimeEnd
}

func (s *ListMaterialDocumentsShrinkRequest) GetUpdateTimeStart() *string {
	return s.UpdateTimeStart
}

func (s *ListMaterialDocumentsShrinkRequest) SetAgentKey(v string) *ListMaterialDocumentsShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetContent(v string) *ListMaterialDocumentsShrinkRequest {
	s.Content = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetCreateTimeEnd(v string) *ListMaterialDocumentsShrinkRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetCreateTimeStart(v string) *ListMaterialDocumentsShrinkRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetCurrent(v int32) *ListMaterialDocumentsShrinkRequest {
	s.Current = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetDocType(v string) *ListMaterialDocumentsShrinkRequest {
	s.DocType = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetDocTypeListShrink(v string) *ListMaterialDocumentsShrinkRequest {
	s.DocTypeListShrink = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetGeneratePublicUrl(v bool) *ListMaterialDocumentsShrinkRequest {
	s.GeneratePublicUrl = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetId(v int64) *ListMaterialDocumentsShrinkRequest {
	s.Id = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetKeywordsShrink(v string) *ListMaterialDocumentsShrinkRequest {
	s.KeywordsShrink = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetQuery(v string) *ListMaterialDocumentsShrinkRequest {
	s.Query = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetShareAttr(v int32) *ListMaterialDocumentsShrinkRequest {
	s.ShareAttr = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetSize(v int32) *ListMaterialDocumentsShrinkRequest {
	s.Size = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetTitle(v string) *ListMaterialDocumentsShrinkRequest {
	s.Title = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetUpdateTimeEnd(v string) *ListMaterialDocumentsShrinkRequest {
	s.UpdateTimeEnd = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetUpdateTimeStart(v string) *ListMaterialDocumentsShrinkRequest {
	s.UpdateTimeStart = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
