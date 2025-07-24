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
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// example:
	//
	// 100km/h-0制动能力上，仅有33.3m，不黑不吹，单看这个，小米SU7确实表现不错。而续航方面，101kWh电池容量，实现CLTC续航800km，还有现5分钟补能220km，15分钟补能510km的800V高压平台。而在...
	Snippet *string `json:"snippet,omitempty" xml:"snippet,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 小米SU7售价22.99万元起 高管亲自辟谣：发布前不会有价格
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
