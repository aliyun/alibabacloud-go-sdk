// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalPageItem interface {
	dara.Model
	String() string
	GoString() string
	SetLink(v string) *GlobalPageItem
	GetLink() *string
	SetSnippet(v string) *GlobalPageItem
	GetSnippet() *string
	SetTitle(v string) *GlobalPageItem
	GetTitle() *string
}

type GlobalPageItem struct {
	// This parameter is required.
	//
	// example:
	//
	// https://baijiahao.baidu.com/s?id=1787881554557805096
	Link    *string `json:"link,omitempty" xml:"link,omitempty"`
	Snippet *string `json:"snippet,omitempty" xml:"snippet,omitempty"`
	// This parameter is required.
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GlobalPageItem) String() string {
	return dara.Prettify(s)
}

func (s GlobalPageItem) GoString() string {
	return s.String()
}

func (s *GlobalPageItem) GetLink() *string {
	return s.Link
}

func (s *GlobalPageItem) GetSnippet() *string {
	return s.Snippet
}

func (s *GlobalPageItem) GetTitle() *string {
	return s.Title
}

func (s *GlobalPageItem) SetLink(v string) *GlobalPageItem {
	s.Link = &v
	return s
}

func (s *GlobalPageItem) SetSnippet(v string) *GlobalPageItem {
	s.Snippet = &v
	return s
}

func (s *GlobalPageItem) SetTitle(v string) *GlobalPageItem {
	s.Title = &v
	return s
}

func (s *GlobalPageItem) Validate() error {
	return dara.Validate(s)
}
