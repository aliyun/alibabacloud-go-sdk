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
	Content          *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Outline          *string `json:"Outline,omitempty" xml:"Outline,omitempty"`
	PrimaryOutline   *string `json:"PrimaryOutline,omitempty" xml:"PrimaryOutline,omitempty"`
	PubTime          *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	SearchSource     *string `json:"SearchSource,omitempty" xml:"SearchSource,omitempty"`
	SearchSourceName *string `json:"SearchSourceName,omitempty" xml:"SearchSourceName,omitempty"`
	Title            *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Url              *string `json:"Url,omitempty" xml:"Url,omitempty"`
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
