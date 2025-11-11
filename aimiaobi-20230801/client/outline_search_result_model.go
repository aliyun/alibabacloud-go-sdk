// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOutlineSearchResult interface {
	dara.Model
	String() string
	GoString() string
	SetArticles(v []*OutlineWritingArticle) *OutlineSearchResult
	GetArticles() []*OutlineWritingArticle
	SetOutline(v string) *OutlineSearchResult
	GetOutline() *string
	SetOutlineId(v string) *OutlineSearchResult
	GetOutlineId() *string
	SetPrimaryOutline(v string) *OutlineSearchResult
	GetPrimaryOutline() *string
	SetQuery(v string) *OutlineSearchResult
	GetQuery() *string
}

type OutlineSearchResult struct {
	Articles       []*OutlineWritingArticle `json:"Articles,omitempty" xml:"Articles,omitempty" type:"Repeated"`
	Outline        *string                  `json:"Outline,omitempty" xml:"Outline,omitempty"`
	OutlineId      *string                  `json:"OutlineId,omitempty" xml:"OutlineId,omitempty"`
	PrimaryOutline *string                  `json:"PrimaryOutline,omitempty" xml:"PrimaryOutline,omitempty"`
	Query          *string                  `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s OutlineSearchResult) String() string {
	return dara.Prettify(s)
}

func (s OutlineSearchResult) GoString() string {
	return s.String()
}

func (s *OutlineSearchResult) GetArticles() []*OutlineWritingArticle {
	return s.Articles
}

func (s *OutlineSearchResult) GetOutline() *string {
	return s.Outline
}

func (s *OutlineSearchResult) GetOutlineId() *string {
	return s.OutlineId
}

func (s *OutlineSearchResult) GetPrimaryOutline() *string {
	return s.PrimaryOutline
}

func (s *OutlineSearchResult) GetQuery() *string {
	return s.Query
}

func (s *OutlineSearchResult) SetArticles(v []*OutlineWritingArticle) *OutlineSearchResult {
	s.Articles = v
	return s
}

func (s *OutlineSearchResult) SetOutline(v string) *OutlineSearchResult {
	s.Outline = &v
	return s
}

func (s *OutlineSearchResult) SetOutlineId(v string) *OutlineSearchResult {
	s.OutlineId = &v
	return s
}

func (s *OutlineSearchResult) SetPrimaryOutline(v string) *OutlineSearchResult {
	s.PrimaryOutline = &v
	return s
}

func (s *OutlineSearchResult) SetQuery(v string) *OutlineSearchResult {
	s.Query = &v
	return s
}

func (s *OutlineSearchResult) Validate() error {
	if s.Articles != nil {
		for _, item := range s.Articles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
