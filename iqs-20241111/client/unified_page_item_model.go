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
	SetTitle(v string) *UnifiedPageItem
	GetTitle() *string
	SetWebsiteAuthorityScore(v int32) *UnifiedPageItem
	GetWebsiteAuthorityScore() *int32
}

type UnifiedPageItem struct {
	CorrelationTag     *int32    `json:"correlationTag,omitempty" xml:"correlationTag,omitempty"`
	HostAuthorityScore *float64  `json:"hostAuthorityScore,omitempty" xml:"hostAuthorityScore,omitempty"`
	HostLogo           *string   `json:"hostLogo,omitempty" xml:"hostLogo,omitempty"`
	Hostname           *string   `json:"hostname,omitempty" xml:"hostname,omitempty"`
	Images             []*string `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	Link               *string   `json:"link,omitempty" xml:"link,omitempty"`
	MainText           *string   `json:"mainText,omitempty" xml:"mainText,omitempty"`
	MarkdownText       *string   `json:"markdownText,omitempty" xml:"markdownText,omitempty"`
	// example:
	//
	// 2025-04-07T10:15:30.123+08:00
	PublishedTime         *string  `json:"publishedTime,omitempty" xml:"publishedTime,omitempty"`
	RerankScore           *float64 `json:"rerankScore,omitempty" xml:"rerankScore,omitempty"`
	RichMainBody          *string  `json:"richMainBody,omitempty" xml:"richMainBody,omitempty"`
	Snippet               *string  `json:"snippet,omitempty" xml:"snippet,omitempty"`
	Summary               *string  `json:"summary,omitempty" xml:"summary,omitempty"`
	Title                 *string  `json:"title,omitempty" xml:"title,omitempty"`
	WebsiteAuthorityScore *int32   `json:"websiteAuthorityScore,omitempty" xml:"websiteAuthorityScore,omitempty"`
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
