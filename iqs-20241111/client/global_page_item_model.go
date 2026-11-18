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
	// The full URL that the search result points to.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://baijiahao.baidu.com/s?id=1787881554557805096
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// The truncated description of the search result in text format.
	//
	// example:
	//
	// In terms of 100 km/h to 0 braking performance, the distance is only 33.3 m. Objectively speaking, the Xiaomi SU7 performs really well in this regard. As for range, it features a 101 kWh battery capacity, achieving a CLTC range of 800 km, and an 800 V high-voltage platform that supports 220 km of range in 5 minutes and 510 km in 15 minutes. And in..
	Snippet *string `json:"snippet,omitempty" xml:"snippet,omitempty"`
	// The title of the search result in plain text.
	//
	// This parameter is required.
	//
	// example:
	//
	// Xiaomi SU7 priced from 229,900 CNY, executive personally dispels rumors: no price will be announced before launch
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
