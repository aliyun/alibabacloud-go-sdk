// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadPageScrapeBody interface {
	dara.Model
	String() string
	GoString() string
	SetFormats(v []*string) *ReadPageScrapeBody
	GetFormats() []*string
	SetLocation(v string) *ReadPageScrapeBody
	GetLocation() *string
	SetMaxAge(v int32) *ReadPageScrapeBody
	GetMaxAge() *int32
	SetPageTimeout(v int32) *ReadPageScrapeBody
	GetPageTimeout() *int32
	SetReadability(v *ReadPageScrapeBodyReadability) *ReadPageScrapeBody
	GetReadability() *ReadPageScrapeBodyReadability
	SetTimeout(v int32) *ReadPageScrapeBody
	GetTimeout() *int32
	SetUrl(v string) *ReadPageScrapeBody
	GetUrl() *string
}

type ReadPageScrapeBody struct {
	// The format of the parsing result.
	//
	// - rawHtml: the HTML of the target site.
	//
	// - html: the page content processed based on readabilityMode.
	//
	// - markdown: the Markdown content converted from the HTML.
	//
	// - text: the text content extracted from the HTML.
	//
	// - screenshot: a screenshot of the target site.
	Formats []*string `json:"formats,omitempty" xml:"formats,omitempty" type:"Repeated"`
	// This parameter does not need to be specified.
	//
	// example:
	//
	// null
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// The maximum cache validity period. Unit: seconds. Default value: 1296000.
	//
	// 1. If the cache duration is less than the value of maxAge, cached content is returned.
	//
	// 2. If maxAge is set to 0, caching is not used.
	//
	// example:
	//
	// 1296000
	MaxAge *int32 `json:"maxAge,omitempty" xml:"maxAge,omitempty"`
	// The timeout period for waiting for the target site resources to fully load. The value of pageTimeout must be less than the value of timeout.
	//
	// Default value: 15000.
	//
	// example:
	//
	// 15000
	PageTimeout *int32 `json:"pageTimeout,omitempty" xml:"pageTimeout,omitempty"`
	// The readability configuration for the parsing result.
	Readability *ReadPageScrapeBodyReadability `json:"readability,omitempty" xml:"readability,omitempty" type:"Struct"`
	// The end-to-end processing timeout period. Unit: ms.
	//
	// Valid values: [0, 180000].
	//
	// Default value: 60000.
	//
	// example:
	//
	// 60000
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// The target URL to parse. The URL must start with http:// or https://.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://www.example.com
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ReadPageScrapeBody) String() string {
	return dara.Prettify(s)
}

func (s ReadPageScrapeBody) GoString() string {
	return s.String()
}

func (s *ReadPageScrapeBody) GetFormats() []*string {
	return s.Formats
}

func (s *ReadPageScrapeBody) GetLocation() *string {
	return s.Location
}

func (s *ReadPageScrapeBody) GetMaxAge() *int32 {
	return s.MaxAge
}

func (s *ReadPageScrapeBody) GetPageTimeout() *int32 {
	return s.PageTimeout
}

func (s *ReadPageScrapeBody) GetReadability() *ReadPageScrapeBodyReadability {
	return s.Readability
}

func (s *ReadPageScrapeBody) GetTimeout() *int32 {
	return s.Timeout
}

func (s *ReadPageScrapeBody) GetUrl() *string {
	return s.Url
}

func (s *ReadPageScrapeBody) SetFormats(v []*string) *ReadPageScrapeBody {
	s.Formats = v
	return s
}

func (s *ReadPageScrapeBody) SetLocation(v string) *ReadPageScrapeBody {
	s.Location = &v
	return s
}

func (s *ReadPageScrapeBody) SetMaxAge(v int32) *ReadPageScrapeBody {
	s.MaxAge = &v
	return s
}

func (s *ReadPageScrapeBody) SetPageTimeout(v int32) *ReadPageScrapeBody {
	s.PageTimeout = &v
	return s
}

func (s *ReadPageScrapeBody) SetReadability(v *ReadPageScrapeBodyReadability) *ReadPageScrapeBody {
	s.Readability = v
	return s
}

func (s *ReadPageScrapeBody) SetTimeout(v int32) *ReadPageScrapeBody {
	s.Timeout = &v
	return s
}

func (s *ReadPageScrapeBody) SetUrl(v string) *ReadPageScrapeBody {
	s.Url = &v
	return s
}

func (s *ReadPageScrapeBody) Validate() error {
	if s.Readability != nil {
		if err := s.Readability.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReadPageScrapeBodyReadability struct {
	// 是否剔除所有图片
	//
	// 默认值：false
	//
	// example:
	//
	// false
	ExcludeAllImages *bool `json:"excludeAllImages,omitempty" xml:"excludeAllImages,omitempty"`
	// 是否剔除所有链接
	//
	// 默认值：false
	//
	// example:
	//
	// false
	ExcludeAllLinks *bool `json:"excludeAllLinks,omitempty" xml:"excludeAllLinks,omitempty"`
	// 指定排除的标签
	ExcludedTags []*string `json:"excludedTags,omitempty" xml:"excludedTags,omitempty" type:"Repeated"`
	// none：不删除信息，默认为 none
	//
	// normal: 基于自研算法，剔除目标页面无关信息（页头/页脚，导航等）
	//
	// article: 基于自研算法，获取站点主要正文内容(适用于博客、新闻站点，不适用于目录页、导航页)
	//
	// example:
	//
	// none
	ReadabilityMode *string `json:"readabilityMode,omitempty" xml:"readabilityMode,omitempty"`
}

func (s ReadPageScrapeBodyReadability) String() string {
	return dara.Prettify(s)
}

func (s ReadPageScrapeBodyReadability) GoString() string {
	return s.String()
}

func (s *ReadPageScrapeBodyReadability) GetExcludeAllImages() *bool {
	return s.ExcludeAllImages
}

func (s *ReadPageScrapeBodyReadability) GetExcludeAllLinks() *bool {
	return s.ExcludeAllLinks
}

func (s *ReadPageScrapeBodyReadability) GetExcludedTags() []*string {
	return s.ExcludedTags
}

func (s *ReadPageScrapeBodyReadability) GetReadabilityMode() *string {
	return s.ReadabilityMode
}

func (s *ReadPageScrapeBodyReadability) SetExcludeAllImages(v bool) *ReadPageScrapeBodyReadability {
	s.ExcludeAllImages = &v
	return s
}

func (s *ReadPageScrapeBodyReadability) SetExcludeAllLinks(v bool) *ReadPageScrapeBodyReadability {
	s.ExcludeAllLinks = &v
	return s
}

func (s *ReadPageScrapeBodyReadability) SetExcludedTags(v []*string) *ReadPageScrapeBodyReadability {
	s.ExcludedTags = v
	return s
}

func (s *ReadPageScrapeBodyReadability) SetReadabilityMode(v string) *ReadPageScrapeBodyReadability {
	s.ReadabilityMode = &v
	return s
}

func (s *ReadPageScrapeBodyReadability) Validate() error {
	return dara.Validate(s)
}
