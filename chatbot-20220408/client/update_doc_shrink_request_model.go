// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDocShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *UpdateDocShrinkRequest
	GetAgentKey() *string
	SetCategoryId(v int64) *UpdateDocShrinkRequest
	GetCategoryId() *int64
	SetConfig(v string) *UpdateDocShrinkRequest
	GetConfig() *string
	SetContent(v string) *UpdateDocShrinkRequest
	GetContent() *string
	SetDocMetadataShrink(v string) *UpdateDocShrinkRequest
	GetDocMetadataShrink() *string
	SetDocName(v string) *UpdateDocShrinkRequest
	GetDocName() *string
	SetEndDate(v string) *UpdateDocShrinkRequest
	GetEndDate() *string
	SetKnowledgeId(v int64) *UpdateDocShrinkRequest
	GetKnowledgeId() *int64
	SetMeta(v string) *UpdateDocShrinkRequest
	GetMeta() *string
	SetStartDate(v string) *UpdateDocShrinkRequest
	GetStartDate() *string
	SetTagIdsShrink(v string) *UpdateDocShrinkRequest
	GetTagIdsShrink() *string
	SetTitle(v string) *UpdateDocShrinkRequest
	GetTitle() *string
}

type UpdateDocShrinkRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 231001028593
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// {"Splitter":"treeSplitter","ChunkSize":500,"TreePatterns":["^# .*","^## .*","^### .*","^#### .*"],"TitleSource":""}
	Config            *string `json:"Config,omitempty" xml:"Config,omitempty"`
	Content           *string `json:"Content,omitempty" xml:"Content,omitempty"`
	DocMetadataShrink *string `json:"DocMetadata,omitempty" xml:"DocMetadata,omitempty"`
	DocName           *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
	// example:
	//
	// 2023-03-11T23:59:59Z
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001905617
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// example:
	//
	// {"code":"xxx"}
	Meta *string `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// example:
	//
	// 2022-05-25T16:28:36Z
	StartDate    *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	TagIdsShrink *string `json:"TagIds,omitempty" xml:"TagIds,omitempty"`
	Title        *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateDocShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDocShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDocShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *UpdateDocShrinkRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *UpdateDocShrinkRequest) GetConfig() *string {
	return s.Config
}

func (s *UpdateDocShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateDocShrinkRequest) GetDocMetadataShrink() *string {
	return s.DocMetadataShrink
}

func (s *UpdateDocShrinkRequest) GetDocName() *string {
	return s.DocName
}

func (s *UpdateDocShrinkRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *UpdateDocShrinkRequest) GetKnowledgeId() *int64 {
	return s.KnowledgeId
}

func (s *UpdateDocShrinkRequest) GetMeta() *string {
	return s.Meta
}

func (s *UpdateDocShrinkRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *UpdateDocShrinkRequest) GetTagIdsShrink() *string {
	return s.TagIdsShrink
}

func (s *UpdateDocShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateDocShrinkRequest) SetAgentKey(v string) *UpdateDocShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateDocShrinkRequest) SetCategoryId(v int64) *UpdateDocShrinkRequest {
	s.CategoryId = &v
	return s
}

func (s *UpdateDocShrinkRequest) SetConfig(v string) *UpdateDocShrinkRequest {
	s.Config = &v
	return s
}

func (s *UpdateDocShrinkRequest) SetContent(v string) *UpdateDocShrinkRequest {
	s.Content = &v
	return s
}

func (s *UpdateDocShrinkRequest) SetDocMetadataShrink(v string) *UpdateDocShrinkRequest {
	s.DocMetadataShrink = &v
	return s
}

func (s *UpdateDocShrinkRequest) SetDocName(v string) *UpdateDocShrinkRequest {
	s.DocName = &v
	return s
}

func (s *UpdateDocShrinkRequest) SetEndDate(v string) *UpdateDocShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *UpdateDocShrinkRequest) SetKnowledgeId(v int64) *UpdateDocShrinkRequest {
	s.KnowledgeId = &v
	return s
}

func (s *UpdateDocShrinkRequest) SetMeta(v string) *UpdateDocShrinkRequest {
	s.Meta = &v
	return s
}

func (s *UpdateDocShrinkRequest) SetStartDate(v string) *UpdateDocShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *UpdateDocShrinkRequest) SetTagIdsShrink(v string) *UpdateDocShrinkRequest {
	s.TagIdsShrink = &v
	return s
}

func (s *UpdateDocShrinkRequest) SetTitle(v string) *UpdateDocShrinkRequest {
	s.Title = &v
	return s
}

func (s *UpdateDocShrinkRequest) Validate() error {
	return dara.Validate(s)
}
