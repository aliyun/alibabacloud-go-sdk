// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWeiboItem interface {
	dara.Model
	String() string
	GoString() string
	SetCardType(v string) *WeiboItem
	GetCardType() *string
	SetHomepageLink(v string) *WeiboItem
	GetHomepageLink() *string
	SetHtmlSnippet(v string) *WeiboItem
	GetHtmlSnippet() *string
	SetImages(v []*string) *WeiboItem
	GetImages() []*string
	SetLink(v string) *WeiboItem
	GetLink() *string
	SetPublishDisplayTime(v string) *WeiboItem
	GetPublishDisplayTime() *string
	SetUsername(v string) *WeiboItem
	GetUsername() *string
}

type WeiboItem struct {
	// This parameter is required.
	//
	// example:
	//
	// weibo_strong
	CardType *string `json:"cardType,omitempty" xml:"cardType,omitempty"`
	// example:
	//
	// https://m.weibo.cn/u/7761793874?luicode=20000061&lfid=5024099350350075
	HomepageLink *string `json:"homepageLink,omitempty" xml:"homepageLink,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 【小调查：你会买<em>小米SU7</em>吗？】#小米SU7路测覆盖300多城市#4月17日，@小米汽车 发文称SU7道路测试覆盖全国300多个城市，涵盖极寒，极热天气，总里程数高达540万公里，目前仍在进行中。  网页链接
	HtmlSnippet *string   `json:"htmlSnippet,omitempty" xml:"htmlSnippet,omitempty"`
	Images      []*string `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// https://m.weibo.cn/detail/5024099350350075?wm=90194_90009
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1小时前
	PublishDisplayTime *string `json:"publishDisplayTime,omitempty" xml:"publishDisplayTime,omitempty"`
	// example:
	//
	// 白鹿科技
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s WeiboItem) String() string {
	return dara.Prettify(s)
}

func (s WeiboItem) GoString() string {
	return s.String()
}

func (s *WeiboItem) GetCardType() *string {
	return s.CardType
}

func (s *WeiboItem) GetHomepageLink() *string {
	return s.HomepageLink
}

func (s *WeiboItem) GetHtmlSnippet() *string {
	return s.HtmlSnippet
}

func (s *WeiboItem) GetImages() []*string {
	return s.Images
}

func (s *WeiboItem) GetLink() *string {
	return s.Link
}

func (s *WeiboItem) GetPublishDisplayTime() *string {
	return s.PublishDisplayTime
}

func (s *WeiboItem) GetUsername() *string {
	return s.Username
}

func (s *WeiboItem) SetCardType(v string) *WeiboItem {
	s.CardType = &v
	return s
}

func (s *WeiboItem) SetHomepageLink(v string) *WeiboItem {
	s.HomepageLink = &v
	return s
}

func (s *WeiboItem) SetHtmlSnippet(v string) *WeiboItem {
	s.HtmlSnippet = &v
	return s
}

func (s *WeiboItem) SetImages(v []*string) *WeiboItem {
	s.Images = v
	return s
}

func (s *WeiboItem) SetLink(v string) *WeiboItem {
	s.Link = &v
	return s
}

func (s *WeiboItem) SetPublishDisplayTime(v string) *WeiboItem {
	s.PublishDisplayTime = &v
	return s
}

func (s *WeiboItem) SetUsername(v string) *WeiboItem {
	s.Username = &v
	return s
}

func (s *WeiboItem) Validate() error {
	return dara.Validate(s)
}
