// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScorePageItem interface {
	dara.Model
	String() string
	GoString() string
	SetCardType(v string) *ScorePageItem
	GetCardType() *string
	SetDisplayLink(v string) *ScorePageItem
	GetDisplayLink() *string
	SetHostLogo(v string) *ScorePageItem
	GetHostLogo() *string
	SetHostname(v string) *ScorePageItem
	GetHostname() *string
	SetHtmlSnippet(v string) *ScorePageItem
	GetHtmlSnippet() *string
	SetHtmlTitle(v string) *ScorePageItem
	GetHtmlTitle() *string
	SetImages(v []*IncludeImage) *ScorePageItem
	GetImages() []*IncludeImage
	SetLink(v string) *ScorePageItem
	GetLink() *string
	SetMainText(v string) *ScorePageItem
	GetMainText() *string
	SetMarkdownText(v string) *ScorePageItem
	GetMarkdownText() *string
	SetMime(v string) *ScorePageItem
	GetMime() *string
	SetPageMap(v map[string]*string) *ScorePageItem
	GetPageMap() map[string]*string
	SetPublishTime(v int64) *ScorePageItem
	GetPublishTime() *int64
	SetScore(v float64) *ScorePageItem
	GetScore() *float64
	SetSiteLabel(v string) *ScorePageItem
	GetSiteLabel() *string
	SetSnippet(v string) *ScorePageItem
	GetSnippet() *string
	SetSummary(v string) *ScorePageItem
	GetSummary() *string
	SetTitle(v string) *ScorePageItem
	GetTitle() *string
}

type ScorePageItem struct {
	// This parameter is required.
	//
	// example:
	//
	// structure_web_info
	CardType *string `json:"cardType,omitempty" xml:"cardType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// baijiahao.baidu.com
	DisplayLink *string `json:"displayLink,omitempty" xml:"displayLink,omitempty"`
	// example:
	//
	// https://s2.zimgs.cn/ims?kt=url&at=smstruct&key=aHR0cHM6Ly9ndy5hbGljZG4uY29tL0wxLzcyMy8xNTY1MjU2NjAwLzJhL2YwL2I0LzJhZjBiNDQxMGI5YmVlMDVjOGVlNGJmODk3MTNkNTFjLnBuZw==&sign=yx:CUlNNQVJQjFrk3Kxt2F3KWhTOFU=&tv=400_400
	HostLogo *string `json:"hostLogo,omitempty" xml:"hostLogo,omitempty"`
	// example:
	//
	// 新华网
	Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100km/h-0制动能力上，仅有33.3m，不黑不吹，单看这个，<em>小米SU7</em>确实表现不错。而续航方面，101kWh电池容量，实现CLTC续航800km，还有现5分钟补能220km，15分钟补能510km的800V高压平台。而在...
	HtmlSnippet *string `json:"htmlSnippet,omitempty" xml:"htmlSnippet,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// <em>小米</em>SU7售价22.99万元起 高管亲自辟谣：发布前不会有<em>价格</em>-百家号
	HtmlTitle *string         `json:"htmlTitle,omitempty" xml:"htmlTitle,omitempty"`
	Images    []*IncludeImage `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// https://baijiahao.baidu.com/s?id=1787881554557805096
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// example:
	//
	// 昨天	，	小米	汽车	没有	发布	，	但	相关	的	信息	透露	的	差	不	多	了	。
	//
	// 		在	发布	会	现场	，	雷军	直接	称	小米	S	U	7	对	标	特斯拉	保时捷	，	有	100	项	行业	领先	，	可见	“	遥遥	领先	”	已经	不再	是	华为	专利	了	？
	//
	//
	//
	// 		而	在	介绍	技术	时	，	雷军	也	从	电机	、	电池	、	大	压铸	、	自动	驾驶	、	智能	座舱	等	五	大	方面	展开	，	充分	展示	了	小米	汽车	的	技术	以及	技术	储备	，		 		能	堆	的	料	，	全都	堆	上去	了	…	…
	//
	// 		大家	比较	感	兴趣	的	性能	方面	，	2	.	78	s	的	0	-	100	km	/	h	加速	，	265	km	/	h	的	最高	时速
	MainText     *string `json:"mainText,omitempty" xml:"mainText,omitempty"`
	MarkdownText *string `json:"markdownText,omitempty" xml:"markdownText,omitempty"`
	// example:
	//
	// text/html
	Mime    *string            `json:"mime,omitempty" xml:"mime,omitempty"`
	PageMap map[string]*string `json:"pageMap,omitempty" xml:"pageMap,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1704426524000
	PublishTime *int64 `json:"publishTime,omitempty" xml:"publishTime,omitempty"`
	// example:
	//
	// 0.234325235
	Score *float64 `json:"score,omitempty" xml:"score,omitempty"`
	// example:
	//
	// 权威媒体
	SiteLabel *string `json:"siteLabel,omitempty" xml:"siteLabel,omitempty"`
	// example:
	//
	// 100km/h-0制动能力上，仅有33.3m，不黑不吹，单看这个，小米SU7确实表现不错。而续航方面，101kWh电池容量，实现CLTC续航800km，还有现5分钟补能220km，15分钟补能510km的800V高压平台。而在...
	Snippet *string `json:"snippet,omitempty" xml:"snippet,omitempty"`
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 小米SU7售价22.99万元起 高管亲自辟谣：发布前不会有价格
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ScorePageItem) String() string {
	return dara.Prettify(s)
}

func (s ScorePageItem) GoString() string {
	return s.String()
}

func (s *ScorePageItem) GetCardType() *string {
	return s.CardType
}

func (s *ScorePageItem) GetDisplayLink() *string {
	return s.DisplayLink
}

func (s *ScorePageItem) GetHostLogo() *string {
	return s.HostLogo
}

func (s *ScorePageItem) GetHostname() *string {
	return s.Hostname
}

func (s *ScorePageItem) GetHtmlSnippet() *string {
	return s.HtmlSnippet
}

func (s *ScorePageItem) GetHtmlTitle() *string {
	return s.HtmlTitle
}

func (s *ScorePageItem) GetImages() []*IncludeImage {
	return s.Images
}

func (s *ScorePageItem) GetLink() *string {
	return s.Link
}

func (s *ScorePageItem) GetMainText() *string {
	return s.MainText
}

func (s *ScorePageItem) GetMarkdownText() *string {
	return s.MarkdownText
}

func (s *ScorePageItem) GetMime() *string {
	return s.Mime
}

func (s *ScorePageItem) GetPageMap() map[string]*string {
	return s.PageMap
}

func (s *ScorePageItem) GetPublishTime() *int64 {
	return s.PublishTime
}

func (s *ScorePageItem) GetScore() *float64 {
	return s.Score
}

func (s *ScorePageItem) GetSiteLabel() *string {
	return s.SiteLabel
}

func (s *ScorePageItem) GetSnippet() *string {
	return s.Snippet
}

func (s *ScorePageItem) GetSummary() *string {
	return s.Summary
}

func (s *ScorePageItem) GetTitle() *string {
	return s.Title
}

func (s *ScorePageItem) SetCardType(v string) *ScorePageItem {
	s.CardType = &v
	return s
}

func (s *ScorePageItem) SetDisplayLink(v string) *ScorePageItem {
	s.DisplayLink = &v
	return s
}

func (s *ScorePageItem) SetHostLogo(v string) *ScorePageItem {
	s.HostLogo = &v
	return s
}

func (s *ScorePageItem) SetHostname(v string) *ScorePageItem {
	s.Hostname = &v
	return s
}

func (s *ScorePageItem) SetHtmlSnippet(v string) *ScorePageItem {
	s.HtmlSnippet = &v
	return s
}

func (s *ScorePageItem) SetHtmlTitle(v string) *ScorePageItem {
	s.HtmlTitle = &v
	return s
}

func (s *ScorePageItem) SetImages(v []*IncludeImage) *ScorePageItem {
	s.Images = v
	return s
}

func (s *ScorePageItem) SetLink(v string) *ScorePageItem {
	s.Link = &v
	return s
}

func (s *ScorePageItem) SetMainText(v string) *ScorePageItem {
	s.MainText = &v
	return s
}

func (s *ScorePageItem) SetMarkdownText(v string) *ScorePageItem {
	s.MarkdownText = &v
	return s
}

func (s *ScorePageItem) SetMime(v string) *ScorePageItem {
	s.Mime = &v
	return s
}

func (s *ScorePageItem) SetPageMap(v map[string]*string) *ScorePageItem {
	s.PageMap = v
	return s
}

func (s *ScorePageItem) SetPublishTime(v int64) *ScorePageItem {
	s.PublishTime = &v
	return s
}

func (s *ScorePageItem) SetScore(v float64) *ScorePageItem {
	s.Score = &v
	return s
}

func (s *ScorePageItem) SetSiteLabel(v string) *ScorePageItem {
	s.SiteLabel = &v
	return s
}

func (s *ScorePageItem) SetSnippet(v string) *ScorePageItem {
	s.Snippet = &v
	return s
}

func (s *ScorePageItem) SetSummary(v string) *ScorePageItem {
	s.Summary = &v
	return s
}

func (s *ScorePageItem) SetTitle(v string) *ScorePageItem {
	s.Title = &v
	return s
}

func (s *ScorePageItem) Validate() error {
	return dara.Validate(s)
}
