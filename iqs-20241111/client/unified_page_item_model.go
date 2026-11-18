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
	// The relevance tag. A value of 0 indicates low relevance, and 1 indicates non-low relevance.
	//
	// example:
	//
	// 1
	CorrelationTag *int32 `json:"correlationTag,omitempty" xml:"correlationTag,omitempty"`
	// The dynamic authority score, which ranges from 0 to 2 and is calculated based on the relevance to the query content.
	//
	// example:
	//
	// 1.923339
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
	// China.com
	Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty"`
	// The images from the search result web page, up to three.
	Images []*string `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	// The full URL pointed to by the search result.
	//
	// example:
	//
	// https://hea.china.com/articles/20250427/202504271666343.html
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// The full text of the searched web page.
	//
	// example:
	//
	// During the 2025 Shanghai Auto Show, the AITO brand was undoubtedly the most popular, becoming the center of attention for many industry leaders. Zhang Xinghai, Chairman and Founder of SERES Group, Yu Chengdong, Managing Director of Huawei and Chairman of Terminal BG, Chen Daoming, a famous actor, and Zeng Yuqun, Chairman of CATL, all shared their experiences using AITO.\\n"AITO M9 and M8 are reshaping the landscape of China\\"s luxury car market," Zhang Xinghai said, (omitted here....).\\nEditor: kj005\\n Article complaint hotline: 157 3889 8464 Complaint email: 7983347 16@qq.com\\nKeywords:
	MainText *string `json:"mainText,omitempty" xml:"mainText,omitempty"`
	// The Markdown content.
	//
	// example:
	//
	// # Industry leaders gather at the Shanghai Auto Show as AITO becomes the center of attention**Source**: Caixun.com
	//
	// **Time**: 2025-04-27 20:36:04
	//
	// During the 2025 Shanghai Auto Show, the AITO brand was undoubtedly the most popular, becoming the center of attention for many industry leaders. Zhang Xinghai, Chairman and Founder of SERES Group, Yu Chengdong, Managing Director of Huawei and Chairman of Terminal BG, Chen Daoming, a famous actor, and Zeng Yuqun, Chairman of CATL, all shared their experiences using AITO.
	//
	// ![Image](http://objectnsg.oss-cn-beijing.aliyuncs.com/yhdoc/202504/27/202504272025531927023138.png)Yu Chengdong stated that AITO M9, M8, M7, and M5 are all deeply loved by consumers! AITO M9 has ranked first in the residual value rate of pure electric models in China for three consecutive months! Pure electric and extended-range models took the top two spots in the residual value rate of new energy large SUVs!
	//
	// ![Image](http://objectnsg.oss-cn-beijing.aliyuncs.com/yhdoc/202504/27/202504272025531312025791.jpeg)"I am an AITO M9 user
	//
	// *Editor: kj005*Article complaint hotline: 157 3889 8464
	//
	// Complaint email: 798334716@qq.com
	MarkdownText *string `json:"markdownText,omitempty" xml:"markdownText,omitempty"`
	// The publish time of the web page in ISO format.
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
	RerankScore *float64 `json:"rerankScore,omitempty" xml:"rerankScore,omitempty"`
	// The full rich text body, which preserves the order of text, images, audio, video, tables, and code on the web page.
	//
	// example:
	//
	// Salute	to	the	workers	!	\\n	[	Building	dreams	for	modernization	 	and	drawing	a	new	picture	together	·	Song	of	the	workers	]	Make	new	contributions	to	promoting	Chinese	modernization	\\n	<quark-image available_status="3" caption="CCTV news studio" fetch_status="0" is_broken="0" is_high_quality_chart="0" is_high_quality_text="0" is_watermark="0" ori_height="1080" ori_width="1920" phash="13061254655082448525" relevant_score="0.5" render_height="137.046875" render_width="243.65625" tag="CCTV logo, studio, architecture, scenery" url="http://s2.zimgs.cn/ims?kt=url&amp;key=aHR0cHM6Ly92Mi5jcmkuY24vTTAwLzg2LzAyL3JCQUJDbVkzRzhDQVhFR0pBQUFBQUFBQUFBQTQ0LjE5MjB4MTA4MC5qcGVn&amp;sign=yx:G3TGlZZMQhnxU2bJSDcb_g8xPQ0=&amp;tv=0_0&amp;p=" value_score="0.5"><quark-ocr>	CCTV	 	13	 	News	 	CTVR	 	Morning	World	</quark-ocr></quark-image>	\\n	More	>	>	>	 	Most	beautiful	workers	\\n
	RichMainBody *string `json:"richMainBody,omitempty" xml:"richMainBody,omitempty"`
	// The summary.
	//
	// example:
	//
	// During the 2025 Shanghai Auto Show, the AITO brand was undoubtedly the most popular, becoming the center of attention for many industry leaders. Zhang Xinghai, Chairman and Founder of SERES Group, Yu Chengdong, Managing Director of Huawei and Chairman of Terminal BG, Chen Daoming, a famous actor, and Zeng Yuqun, Chairman of CATL, all shared their experiences using AITO. ..
	Snippet *string `json:"snippet,omitempty" xml:"snippet,omitempty"`
	// The enhanced summary, 400 words or more.
	//
	// example:
	//
	// During the 2025 Shanghai Auto Show, the AITO brand was undoubtedly the most popular, becoming the center of attention for many industry leaders. Zhang Xinghai, Chairman and Founder of SERES Group, Yu Chengdong, Managing Director of Huawei and Chairman of Terminal BG, Chen Daoming, a famous actor, and Zeng Yuqun, Chairman of CATL, all shared their experiences using AITO. "I am an AITO M9 user, and the driving experience is excellent," Chen Daoming said, adding that he only dared to recommend it on the spot after experiencing it himself. In the future, he hopes to achieve special customization of the AITO M9, making the exterior, interior, and seat layout more in line with his own needs. It is reported that the 2025 AITO M9 delivered over 10,000 units within 25 days of its launch, and has been the sales champion of 500,000 RMB luxury cars for 11 consecutive months; the AITO M8 received over 50,000 firm orders within 4 days of its launch, and is deeply loved by the market and users... "AITO M9 and M8 are reshaping the landscape of China\\"s luxury car market," Zhang Xinghai said. He recently test-drove the AITO M8 to the limit, departing from Chongqing and arriving at Mount Kailash after 55 hours and over 3,500 kilometers. Throughout the journey, the safety and experience of the vehicle performed excellently
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// The web page tags.
	//
	// example:
	//
	// {"genre":"NewsPortal"}
	Tags map[string]interface{} `json:"tags,omitempty" xml:"tags,omitempty"`
	// The full URL pointed to by the search result.
	//
	// example:
	//
	// Industry leaders gather at the Shanghai Auto Show as AITO becomes the center of attention
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// The static authority score, which is determined solely by the authority of the website itself and is independent of the query content.
	//
	// example:
	//
	// 3
	WebsiteAuthorityScore *int32 `json:"websiteAuthorityScore,omitempty" xml:"websiteAuthorityScore,omitempty"`
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
