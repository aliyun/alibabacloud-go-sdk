// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOutlineWritingArticle interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *OutlineWritingArticle
	GetContent() *string
	SetOutline(v string) *OutlineWritingArticle
	GetOutline() *string
	SetPrimaryOutline(v string) *OutlineWritingArticle
	GetPrimaryOutline() *string
	SetPubTime(v string) *OutlineWritingArticle
	GetPubTime() *string
	SetSearchSource(v string) *OutlineWritingArticle
	GetSearchSource() *string
	SetSearchSourceName(v string) *OutlineWritingArticle
	GetSearchSourceName() *string
	SetTitle(v string) *OutlineWritingArticle
	GetTitle() *string
	SetUrl(v string) *OutlineWritingArticle
	GetUrl() *string
}

type OutlineWritingArticle struct {
	// Content
	//
	// example:
	//
	// 新闻内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The leaf outline this material belongs to
	//
	// example:
	//
	// 大纲名称
	Outline *string `json:"Outline,omitempty" xml:"Outline,omitempty"`
	// The primary outline this material belongs to
	//
	// example:
	//
	// 一级大纲名称
	PrimaryOutline *string `json:"PrimaryOutline,omitempty" xml:"PrimaryOutline,omitempty"`
	// Publish time
	//
	// example:
	//
	// 2023-04-11 06:14:07
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// Retrieval source encoding
	//
	// example:
	//
	// 检索源编码
	SearchSource *string `json:"SearchSource,omitempty" xml:"SearchSource,omitempty"`
	// Retrieval name
	//
	// example:
	//
	// 检索源名称
	SearchSourceName *string `json:"SearchSourceName,omitempty" xml:"SearchSourceName,omitempty"`
	// News title
	//
	// example:
	//
	// 新闻标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// News URL
	//
	// example:
	//
	// http://www.example.com/xxxx.html
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s OutlineWritingArticle) String() string {
	return dara.Prettify(s)
}

func (s OutlineWritingArticle) GoString() string {
	return s.String()
}

func (s *OutlineWritingArticle) GetContent() *string {
	return s.Content
}

func (s *OutlineWritingArticle) GetOutline() *string {
	return s.Outline
}

func (s *OutlineWritingArticle) GetPrimaryOutline() *string {
	return s.PrimaryOutline
}

func (s *OutlineWritingArticle) GetPubTime() *string {
	return s.PubTime
}

func (s *OutlineWritingArticle) GetSearchSource() *string {
	return s.SearchSource
}

func (s *OutlineWritingArticle) GetSearchSourceName() *string {
	return s.SearchSourceName
}

func (s *OutlineWritingArticle) GetTitle() *string {
	return s.Title
}

func (s *OutlineWritingArticle) GetUrl() *string {
	return s.Url
}

func (s *OutlineWritingArticle) SetContent(v string) *OutlineWritingArticle {
	s.Content = &v
	return s
}

func (s *OutlineWritingArticle) SetOutline(v string) *OutlineWritingArticle {
	s.Outline = &v
	return s
}

func (s *OutlineWritingArticle) SetPrimaryOutline(v string) *OutlineWritingArticle {
	s.PrimaryOutline = &v
	return s
}

func (s *OutlineWritingArticle) SetPubTime(v string) *OutlineWritingArticle {
	s.PubTime = &v
	return s
}

func (s *OutlineWritingArticle) SetSearchSource(v string) *OutlineWritingArticle {
	s.SearchSource = &v
	return s
}

func (s *OutlineWritingArticle) SetSearchSourceName(v string) *OutlineWritingArticle {
	s.SearchSourceName = &v
	return s
}

func (s *OutlineWritingArticle) SetTitle(v string) *OutlineWritingArticle {
	s.Title = &v
	return s
}

func (s *OutlineWritingArticle) SetUrl(v string) *OutlineWritingArticle {
	s.Url = &v
	return s
}

func (s *OutlineWritingArticle) Validate() error {
	return dara.Validate(s)
}
