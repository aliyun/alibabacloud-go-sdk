// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDocShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *CreateDocShrinkRequest
	GetAgentKey() *string
	SetCategoryId(v int64) *CreateDocShrinkRequest
	GetCategoryId() *int64
	SetConfig(v string) *CreateDocShrinkRequest
	GetConfig() *string
	SetContent(v string) *CreateDocShrinkRequest
	GetContent() *string
	SetDocMetadataShrink(v string) *CreateDocShrinkRequest
	GetDocMetadataShrink() *string
	SetEndDate(v string) *CreateDocShrinkRequest
	GetEndDate() *string
	SetMeta(v string) *CreateDocShrinkRequest
	GetMeta() *string
	SetStartDate(v string) *CreateDocShrinkRequest
	GetStartDate() *string
	SetTagIdsShrink(v string) *CreateDocShrinkRequest
	GetTagIdsShrink() *string
	SetTitle(v string) *CreateDocShrinkRequest
	GetTitle() *string
	SetUrl(v string) *CreateDocShrinkRequest
	GetUrl() *string
}

type CreateDocShrinkRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30000049006
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// {"Splitter":"treeSplitter","ChunkSize":500,"TreePatterns":["^# .*","^## .*","^### .*","^#### .*"],"TitleSource":""}
	Config            *string `json:"Config,omitempty" xml:"Config,omitempty"`
	Content           *string `json:"Content,omitempty" xml:"Content,omitempty"`
	DocMetadataShrink *string `json:"DocMetadata,omitempty" xml:"DocMetadata,omitempty"`
	// example:
	//
	// 2032-05-25T16:28:36Z
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// {"code":"xxx"}
	Meta *string `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// example:
	//
	// 2022-05-25T16:28:36Z
	StartDate    *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	TagIdsShrink *string `json:"TagIds,omitempty" xml:"TagIds,omitempty"`
	// This parameter is required.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// https://example.com/example.pdf
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateDocShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDocShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDocShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CreateDocShrinkRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *CreateDocShrinkRequest) GetConfig() *string {
	return s.Config
}

func (s *CreateDocShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *CreateDocShrinkRequest) GetDocMetadataShrink() *string {
	return s.DocMetadataShrink
}

func (s *CreateDocShrinkRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *CreateDocShrinkRequest) GetMeta() *string {
	return s.Meta
}

func (s *CreateDocShrinkRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *CreateDocShrinkRequest) GetTagIdsShrink() *string {
	return s.TagIdsShrink
}

func (s *CreateDocShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateDocShrinkRequest) GetUrl() *string {
	return s.Url
}

func (s *CreateDocShrinkRequest) SetAgentKey(v string) *CreateDocShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateDocShrinkRequest) SetCategoryId(v int64) *CreateDocShrinkRequest {
	s.CategoryId = &v
	return s
}

func (s *CreateDocShrinkRequest) SetConfig(v string) *CreateDocShrinkRequest {
	s.Config = &v
	return s
}

func (s *CreateDocShrinkRequest) SetContent(v string) *CreateDocShrinkRequest {
	s.Content = &v
	return s
}

func (s *CreateDocShrinkRequest) SetDocMetadataShrink(v string) *CreateDocShrinkRequest {
	s.DocMetadataShrink = &v
	return s
}

func (s *CreateDocShrinkRequest) SetEndDate(v string) *CreateDocShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *CreateDocShrinkRequest) SetMeta(v string) *CreateDocShrinkRequest {
	s.Meta = &v
	return s
}

func (s *CreateDocShrinkRequest) SetStartDate(v string) *CreateDocShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *CreateDocShrinkRequest) SetTagIdsShrink(v string) *CreateDocShrinkRequest {
	s.TagIdsShrink = &v
	return s
}

func (s *CreateDocShrinkRequest) SetTitle(v string) *CreateDocShrinkRequest {
	s.Title = &v
	return s
}

func (s *CreateDocShrinkRequest) SetUrl(v string) *CreateDocShrinkRequest {
	s.Url = &v
	return s
}

func (s *CreateDocShrinkRequest) Validate() error {
	return dara.Validate(s)
}
