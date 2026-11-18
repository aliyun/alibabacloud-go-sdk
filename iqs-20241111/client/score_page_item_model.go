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
	SetCorrelationTag(v int32) *ScorePageItem
	GetCorrelationTag() *int32
	SetDisplayLink(v string) *ScorePageItem
	GetDisplayLink() *string
	SetHostAuthorityScore(v float64) *ScorePageItem
	GetHostAuthorityScore() *float64
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
	SetRichMainBody(v string) *ScorePageItem
	GetRichMainBody() *string
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
	SetWebsiteAuthorityScore(v int32) *ScorePageItem
	GetWebsiteAuthorityScore() *int32
}

type ScorePageItem struct {
	// The search result template. Different templates return different fields.
	//
	// This parameter is required.
	//
	// example:
	//
	// structure_web_info
	CardType *string `json:"cardType,omitempty" xml:"cardType,omitempty"`
	// The relevance tag. A value of 0 indicates low relevance, and 1 indicates non-low relevance.
	//
	// example:
	//
	// 1
	CorrelationTag *int32 `json:"correlationTag,omitempty" xml:"correlationTag,omitempty"`
	// The abbreviated URL of this search result.
	//
	// This parameter is required.
	//
	// example:
	//
	// baijiahao.baidu.com
	DisplayLink *string `json:"displayLink,omitempty" xml:"displayLink,omitempty"`
	// The dynamic authority score, which ranges from 0 to 2 and is calculated based on the relevance of the query content.
	//
	// example:
	//
	// 1.923339
	HostAuthorityScore *float64 `json:"hostAuthorityScore,omitempty" xml:"hostAuthorityScore,omitempty"`
	// The site logo.
	//
	// example:
	//
	// https://s2.zimgs.cn/ims?kt=url&at=smstruct&key=aHR0cHM6Ly9ndy5hbGljZG4uY29tL0wxLzcyMy8xNTY1MjU2NjAwLzJhL2YwL2I0LzJhZjBiNDQxMGI5YmVlMDVjOGVlNGJmODk3MTNkNTFjLnBuZw==&sign=yx:CUlNNQVJQjFrk3Kxt2F3KWhTOFU=&tv=400_400
	HostLogo *string `json:"hostLogo,omitempty" xml:"hostLogo,omitempty"`
	// The site name.
	//
	// example:
	//
	// Xinhua News
	Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty"`
	// The description snippet of the search result in HTML format.
	//
	// This parameter is required.
	//
	// example:
	//
	// In terms of 100 km/h-0 braking capability, it is only 33.3m. Without bias, just looking at this, the <em>Xiaomi SU7</em> indeed performs well. As for range, with a 101 kWh battery capacity, it achieves a CLTC range of 800 km, and features an 800V high-voltage platform that can add 220 km of range in 5 minutes and 510 km in 15 minutes. And..
	HtmlSnippet *string `json:"htmlSnippet,omitempty" xml:"htmlSnippet,omitempty"`
	// The search result title in HTML format.
	//
	// This parameter is required.
	//
	// example:
	//
	// <em>Xiaomi</em> SU7 priced from 229,900 yuan, executive personally dispels rumors: no <em>price</em> before launch - Baijiahao
	HtmlTitle *string `json:"htmlTitle,omitempty" xml:"htmlTitle,omitempty"`
	// The image information.
	Images []*IncludeImage `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	// The full URL that the search result points to.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://baijiahao.baidu.com/s?id=1787881554557805096
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// The full text of the search result, which is available in some scenarios.
	//
	// example:
	//
	// Yesterday, Xiaomi Auto was not launched, but related information has been mostly revealed. At the launch event, Lei Jun directly stated that Xiaomi SU7 benchmarks Tesla and Porsche, with 100 industry-leading features. It seems that "far ahead" is no longer a Huawei patent? When introducing the technology, Lei Jun also expanded on five major aspects: motor, battery, mega casting, autonomous driving, and smart cockpit, fully demonstrating Xiaomi Auto\\"s technology and technical reserves. All the materials that can be stacked are stacked... Regarding the performance that everyone is interested in, 0-100 km/h acceleration in 2.78s, and a top speed of 265 km/h
	MainText *string `json:"mainText,omitempty" xml:"mainText,omitempty"`
	// The web page body in Markdown format.
	//
	// example:
	//
	// # 2022 Apsara Conference Homepage
	MarkdownText *string `json:"markdownText,omitempty" xml:"markdownText,omitempty"`
	// The MIME type of the search result.
	//
	// example:
	//
	// text/html
	Mime *string `json:"mime,omitempty" xml:"mime,omitempty"`
	// The supplementary page information.
	//
	// example:
	//
	// {}
	PageMap map[string]*string `json:"pageMap,omitempty" xml:"pageMap,omitempty"`
	// The publish time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1704426524000
	PublishTime *int64 `json:"publishTime,omitempty" xml:"publishTime,omitempty"`
	// The full rich text body, which preserves the order of text, images, audio, video, tables, and code in the web page.
	//
	// example:
	//
	// <quark-ocr> 24 hx 3 </quark-ocr>
	RichMainBody *string `json:"richMainBody,omitempty" xml:"richMainBody,omitempty"`
	// The score.
	//
	// example:
	//
	// 0.234325235
	Score *float64 `json:"score,omitempty" xml:"score,omitempty"`
	// The site label. Two types are currently supported:
	//
	// 	- Authoritative media
	//
	//  	- Party and government organs
	//
	// example:
	//
	// Authoritative media
	SiteLabel *string `json:"siteLabel,omitempty" xml:"siteLabel,omitempty"`
	// The description snippet of the search result in text format.
	//
	// example:
	//
	// In terms of 100 km/h-0 braking capability, it is only 33.3m. Without bias, just looking at this, the Xiaomi SU7 indeed performs well. As for range, with a 101 kWh battery capacity, it achieves a CLTC range of 800 km, and features an 800V high-voltage platform that can add 220 km of range in 5 minutes and 510 km in 15 minutes. And..
	Snippet *string `json:"snippet,omitempty" xml:"snippet,omitempty"`
	// The enhanced summary.
	//
	// example:
	//
	// Tomorrow
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// The search result title in plain text.
	//
	// This parameter is required.
	//
	// example:
	//
	// Xiaomi SU7 priced from 229,900 yuan, executive personally dispels rumors: no price before launch
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// The static authority score, which is determined solely by the authority of the website itself and is independent of the query content.
	//
	// example:
	//
	// 3
	WebsiteAuthorityScore *int32 `json:"websiteAuthorityScore,omitempty" xml:"websiteAuthorityScore,omitempty"`
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

func (s *ScorePageItem) GetCorrelationTag() *int32 {
	return s.CorrelationTag
}

func (s *ScorePageItem) GetDisplayLink() *string {
	return s.DisplayLink
}

func (s *ScorePageItem) GetHostAuthorityScore() *float64 {
	return s.HostAuthorityScore
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

func (s *ScorePageItem) GetRichMainBody() *string {
	return s.RichMainBody
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

func (s *ScorePageItem) GetWebsiteAuthorityScore() *int32 {
	return s.WebsiteAuthorityScore
}

func (s *ScorePageItem) SetCardType(v string) *ScorePageItem {
	s.CardType = &v
	return s
}

func (s *ScorePageItem) SetCorrelationTag(v int32) *ScorePageItem {
	s.CorrelationTag = &v
	return s
}

func (s *ScorePageItem) SetDisplayLink(v string) *ScorePageItem {
	s.DisplayLink = &v
	return s
}

func (s *ScorePageItem) SetHostAuthorityScore(v float64) *ScorePageItem {
	s.HostAuthorityScore = &v
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

func (s *ScorePageItem) SetRichMainBody(v string) *ScorePageItem {
	s.RichMainBody = &v
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

func (s *ScorePageItem) SetWebsiteAuthorityScore(v int32) *ScorePageItem {
	s.WebsiteAuthorityScore = &v
	return s
}

func (s *ScorePageItem) Validate() error {
	if s.Images != nil {
		for _, item := range s.Images {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
