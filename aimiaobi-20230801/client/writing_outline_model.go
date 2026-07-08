// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWritingOutline interface {
	dara.Model
	String() string
	GoString() string
	SetArticles(v []*OutlineWritingArticle) *WritingOutline
	GetArticles() []*OutlineWritingArticle
	SetChildren(v []*WritingOutline) *WritingOutline
	GetChildren() []*WritingOutline
	SetOutline(v string) *WritingOutline
	GetOutline() *string
	SetOutlineId(v string) *WritingOutline
	GetOutlineId() *string
	SetSearchKeyWordList(v []*string) *WritingOutline
	GetSearchKeyWordList() []*string
	SetWordCount(v string) *WritingOutline
	GetWordCount() *string
	SetWritingTips(v string) *WritingOutline
	GetWritingTips() *string
}

type WritingOutline struct {
	// List of articles referenced in the outline
	Articles []*OutlineWritingArticle `json:"Articles,omitempty" xml:"Articles,omitempty" type:"Repeated"`
	// List of sub-outlines
	Children []*WritingOutline `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	// Outline
	//
	// example:
	//
	// 大纲名称
	Outline *string `json:"Outline,omitempty" xml:"Outline,omitempty"`
	// Unique outline identifier
	//
	// example:
	//
	// xxx
	OutlineId *string `json:"OutlineId,omitempty" xml:"OutlineId,omitempty"`
	// Outline writing search keyword list
	SearchKeyWordList []*string `json:"SearchKeyWordList,omitempty" xml:"SearchKeyWordList,omitempty" type:"Repeated"`
	// Description of word count requirements, such as 1K words, 2K words, or no less than 500 words.
	//
	// example:
	//
	// 写作字数要求的描述
	WordCount *string `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
	// Outline writing requirements, writing tips
	//
	// example:
	//
	// 大纲写作要求、写作提示
	WritingTips *string `json:"WritingTips,omitempty" xml:"WritingTips,omitempty"`
}

func (s WritingOutline) String() string {
	return dara.Prettify(s)
}

func (s WritingOutline) GoString() string {
	return s.String()
}

func (s *WritingOutline) GetArticles() []*OutlineWritingArticle {
	return s.Articles
}

func (s *WritingOutline) GetChildren() []*WritingOutline {
	return s.Children
}

func (s *WritingOutline) GetOutline() *string {
	return s.Outline
}

func (s *WritingOutline) GetOutlineId() *string {
	return s.OutlineId
}

func (s *WritingOutline) GetSearchKeyWordList() []*string {
	return s.SearchKeyWordList
}

func (s *WritingOutline) GetWordCount() *string {
	return s.WordCount
}

func (s *WritingOutline) GetWritingTips() *string {
	return s.WritingTips
}

func (s *WritingOutline) SetArticles(v []*OutlineWritingArticle) *WritingOutline {
	s.Articles = v
	return s
}

func (s *WritingOutline) SetChildren(v []*WritingOutline) *WritingOutline {
	s.Children = v
	return s
}

func (s *WritingOutline) SetOutline(v string) *WritingOutline {
	s.Outline = &v
	return s
}

func (s *WritingOutline) SetOutlineId(v string) *WritingOutline {
	s.OutlineId = &v
	return s
}

func (s *WritingOutline) SetSearchKeyWordList(v []*string) *WritingOutline {
	s.SearchKeyWordList = v
	return s
}

func (s *WritingOutline) SetWordCount(v string) *WritingOutline {
	s.WordCount = &v
	return s
}

func (s *WritingOutline) SetWritingTips(v string) *WritingOutline {
	s.WritingTips = &v
	return s
}

func (s *WritingOutline) Validate() error {
	if s.Articles != nil {
		for _, item := range s.Articles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Children != nil {
		for _, item := range s.Children {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
