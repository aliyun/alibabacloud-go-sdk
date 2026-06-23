// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnifiedPageItem interface {
	dara.Model
	String() string
	GoString() string
	SetCorrelationTag(v int32) *UnifiedPageItem
	GetCorrelationTag() *int32
	SetHostAuthorityScore(v float64) *UnifiedPageItem
	GetHostAuthorityScore() *float64
	SetHostLogo(v string) *UnifiedPageItem
	GetHostLogo() *string
	SetHostname(v string) *UnifiedPageItem
	GetHostname() *string
	SetImages(v []*string) *UnifiedPageItem
	GetImages() []*string
	SetLink(v string) *UnifiedPageItem
	GetLink() *string
	SetMainText(v string) *UnifiedPageItem
	GetMainText() *string
	SetMarkdownText(v string) *UnifiedPageItem
	GetMarkdownText() *string
	SetPublishedTime(v string) *UnifiedPageItem
	GetPublishedTime() *string
	SetRerankScore(v float64) *UnifiedPageItem
	GetRerankScore() *float64
	SetRichMainBody(v string) *UnifiedPageItem
	GetRichMainBody() *string
	SetSnippet(v string) *UnifiedPageItem
	GetSnippet() *string
	SetSummary(v string) *UnifiedPageItem
	GetSummary() *string
	SetTags(v map[string]interface{}) *UnifiedPageItem
	GetTags() map[string]interface{}
	SetTitle(v string) *UnifiedPageItem
	GetTitle() *string
	SetWebsiteAuthorityScore(v int32) *UnifiedPageItem
	GetWebsiteAuthorityScore() *int32
}

type UnifiedPageItem struct {
	CorrelationTag     *int32   `json:"correlationTag,omitempty" xml:"correlationTag,omitempty"`
	HostAuthorityScore *float64 `json:"hostAuthorityScore,omitempty" xml:"hostAuthorityScore,omitempty"`
	// The site logo.
	//
	// example:
	//
	// https://www.china.com/zh_cn/img1905/2023/logo.png
	HostLogo *string `json:"hostLogo,omitempty" xml:"hostLogo,omitempty"`
	// The site name.
	//
	// example:
	//
	// 中华网
	Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty"`
	// The images in the search result URL. A maximum of three images are returned.
	Images []*string `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	// The full URL that the search result points to.
	//
	// example:
	//
	// https://hea.china.com/articles/20250427/202504271666343.html
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// The full text of the searched web page.
	//
	// example:
	//
	// 2025上海车展期间，最火爆的无疑是问界品牌，成为众多大佬的话题中心。赛力斯集团董事长（创始人）张兴海、华为常务董事、终端BG董事长余承东、著名演员陈道明、宁德时代董事长曾毓群都分享了对问界的使用体验。\\n“问界M9、M8正重构 中国 豪华汽车市场格局。”张兴海说， (此处省略....)。\\n责任编辑：kj005\\n 文章投诉热线:157 3889 8464 投诉邮箱:7983347 16@qq.com\\n关键词：
	MainText *string `json:"mainText,omitempty" xml:"mainText,omitempty"`
	// The Markdown content.
	//
	// example:
	//
	// # 上海车展大佬齐聚 问界成为话题中心**来源**：财讯网
	//
	// **时间**：2025-04-27 20:36:04
	//
	// 2025上海车展期间，最火爆的无疑是问界品牌，成为众多大佬的话题中心。赛力斯集团董事长（创始人）张兴海、华为常务董事、终端BG董事长余承东、著名演员陈道明、宁德时代董事长曾毓群都分享了对问界的使用体验。
	//
	// ![图片](http://objectnsg.oss-cn-beijing.aliyuncs.com/yhdoc/202504/27/202504272025531927023138.png)余承东表示，问界M9、M8、M7和M5，都深受消费者喜爱！问界M9连续3个月中国纯电车型保值率第一！纯电、增程车型包揽新能源大型SUV保值率前两名！
	//
	// ![图片](http://objectnsg.oss-cn-beijing.aliyuncs.com/yhdoc/202504/27/202504272025531312025791.jpeg)“自己是问界M9的用户
	//
	// *责任编辑：kj005*文章投诉热线: 157 3889 8464
	//
	// 投诉邮箱: 798334716@qq.com
	MarkdownText *string `json:"markdownText,omitempty" xml:"markdownText,omitempty"`
	// The time when the web page was published, in ISO time format.
	//
	// example:
	//
	// 2025-04-27T20:36:04+08:00
	PublishedTime *string `json:"publishedTime,omitempty" xml:"publishedTime,omitempty"`
	// The rerank score.
	//
	// example:
	//
	// 0.7786493301391602
	RerankScore  *float64 `json:"rerankScore,omitempty" xml:"rerankScore,omitempty"`
	RichMainBody *string  `json:"richMainBody,omitempty" xml:"richMainBody,omitempty"`
	// The text summary.
	//
	// example:
	//
	// 2025上海车展期间，最火爆的无疑是问界品牌，成为众多大佬的话题中心。赛力斯集团董事长（创始人）张兴海、华为常务董事、终端BG董事长余承东、著名演员陈道明、宁德时代董事长曾毓群都分享了对问界的使用体验。 ...
	Snippet *string `json:"snippet,omitempty" xml:"snippet,omitempty"`
	// The enhanced summary, which contains 400 or more characters.
	//
	// example:
	//
	// 2025上海车展期间，最火爆的无疑是问界品牌，成为众多大佬的话题中心。赛力斯集团董事长（创始人）张兴海、华为常务董事、终端BG董事长余承东、著名演员陈道明、宁德时代董事长曾毓群都分享了对问界的使用体验。“自己是问界M9的用户，用车感受非常好。”陈道明表示，体验后才敢现场给大家推荐。未来，希望能实现问界M9的特别定制，让外观、内饰、座椅布局等，更符合自己的使用需求。据悉，2025款问界M9上市25天就交付超1万，已连续11个月蝉联50万元级豪车车销冠；问界M8上市4天，大定就突破了五万 台 ，深受市场任何和用户喜爱...“问界M9、M8正重构 中国 豪华汽车市场格局。”张兴海说， 近 期曾极限试驾问界M8，从重庆出发，历时55小时、行驶超3500公里抵达冈仁波齐。整个行程中，车辆的安全 性 和体验感都表现出色。
	Summary *string                `json:"summary,omitempty" xml:"summary,omitempty"`
	Tags    map[string]interface{} `json:"tags,omitempty" xml:"tags,omitempty"`
	// The full URL that the search result points to.
	//
	// example:
	//
	// 上海车展大佬齐聚 问界成为话题中心
	Title                 *string `json:"title,omitempty" xml:"title,omitempty"`
	WebsiteAuthorityScore *int32  `json:"websiteAuthorityScore,omitempty" xml:"websiteAuthorityScore,omitempty"`
}

func (s UnifiedPageItem) String() string {
	return dara.Prettify(s)
}

func (s UnifiedPageItem) GoString() string {
	return s.String()
}

func (s *UnifiedPageItem) GetCorrelationTag() *int32 {
	return s.CorrelationTag
}

func (s *UnifiedPageItem) GetHostAuthorityScore() *float64 {
	return s.HostAuthorityScore
}

func (s *UnifiedPageItem) GetHostLogo() *string {
	return s.HostLogo
}

func (s *UnifiedPageItem) GetHostname() *string {
	return s.Hostname
}

func (s *UnifiedPageItem) GetImages() []*string {
	return s.Images
}

func (s *UnifiedPageItem) GetLink() *string {
	return s.Link
}

func (s *UnifiedPageItem) GetMainText() *string {
	return s.MainText
}

func (s *UnifiedPageItem) GetMarkdownText() *string {
	return s.MarkdownText
}

func (s *UnifiedPageItem) GetPublishedTime() *string {
	return s.PublishedTime
}

func (s *UnifiedPageItem) GetRerankScore() *float64 {
	return s.RerankScore
}

func (s *UnifiedPageItem) GetRichMainBody() *string {
	return s.RichMainBody
}

func (s *UnifiedPageItem) GetSnippet() *string {
	return s.Snippet
}

func (s *UnifiedPageItem) GetSummary() *string {
	return s.Summary
}

func (s *UnifiedPageItem) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *UnifiedPageItem) GetTitle() *string {
	return s.Title
}

func (s *UnifiedPageItem) GetWebsiteAuthorityScore() *int32 {
	return s.WebsiteAuthorityScore
}

func (s *UnifiedPageItem) SetCorrelationTag(v int32) *UnifiedPageItem {
	s.CorrelationTag = &v
	return s
}

func (s *UnifiedPageItem) SetHostAuthorityScore(v float64) *UnifiedPageItem {
	s.HostAuthorityScore = &v
	return s
}

func (s *UnifiedPageItem) SetHostLogo(v string) *UnifiedPageItem {
	s.HostLogo = &v
	return s
}

func (s *UnifiedPageItem) SetHostname(v string) *UnifiedPageItem {
	s.Hostname = &v
	return s
}

func (s *UnifiedPageItem) SetImages(v []*string) *UnifiedPageItem {
	s.Images = v
	return s
}

func (s *UnifiedPageItem) SetLink(v string) *UnifiedPageItem {
	s.Link = &v
	return s
}

func (s *UnifiedPageItem) SetMainText(v string) *UnifiedPageItem {
	s.MainText = &v
	return s
}

func (s *UnifiedPageItem) SetMarkdownText(v string) *UnifiedPageItem {
	s.MarkdownText = &v
	return s
}

func (s *UnifiedPageItem) SetPublishedTime(v string) *UnifiedPageItem {
	s.PublishedTime = &v
	return s
}

func (s *UnifiedPageItem) SetRerankScore(v float64) *UnifiedPageItem {
	s.RerankScore = &v
	return s
}

func (s *UnifiedPageItem) SetRichMainBody(v string) *UnifiedPageItem {
	s.RichMainBody = &v
	return s
}

func (s *UnifiedPageItem) SetSnippet(v string) *UnifiedPageItem {
	s.Snippet = &v
	return s
}

func (s *UnifiedPageItem) SetSummary(v string) *UnifiedPageItem {
	s.Summary = &v
	return s
}

func (s *UnifiedPageItem) SetTags(v map[string]interface{}) *UnifiedPageItem {
	s.Tags = v
	return s
}

func (s *UnifiedPageItem) SetTitle(v string) *UnifiedPageItem {
	s.Title = &v
	return s
}

func (s *UnifiedPageItem) SetWebsiteAuthorityScore(v int32) *UnifiedPageItem {
	s.WebsiteAuthorityScore = &v
	return s
}

func (s *UnifiedPageItem) Validate() error {
	return dara.Validate(s)
}
