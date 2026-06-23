// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadPageBody interface {
	dara.Model
	String() string
	GoString() string
	SetFormats(v []*string) *ReadPageBody
	GetFormats() []*string
	SetLocation(v string) *ReadPageBody
	GetLocation() *string
	SetMaxAge(v int32) *ReadPageBody
	GetMaxAge() *int32
	SetPageTimeout(v int32) *ReadPageBody
	GetPageTimeout() *int32
	SetReadability(v *ReadPageBodyReadability) *ReadPageBody
	GetReadability() *ReadPageBodyReadability
	SetTimeout(v int32) *ReadPageBody
	GetTimeout() *int32
	SetUrl(v string) *ReadPageBody
	GetUrl() *string
}

type ReadPageBody struct {
	// The format of the parsing result.
	//
	// - rawHtml: the HTML of the target website.
	//
	// - html: the page content processed based on readabilityMode.
	//
	// - markdown: the Markdown content converted from HTML.
	//
	// - text: the text content extracted from HTML.
	Formats []*string `json:"formats,omitempty" xml:"formats,omitempty" type:"Repeated"`
	// You do not need to specify this parameter.
	//
	// example:
	//
	// null
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// The maximum cache validity period. Unit: seconds. Default value: 1296000.
	//
	// - If the cache duration is less than the value of maxAge, cached content is returned.
	//
	// - If the value of maxAge is 0, caching is not used.
	//
	// example:
	//
	// 1296000
	MaxAge *int32 `json:"maxAge,omitempty" xml:"maxAge,omitempty"`
	// The URL read timeout period. The value of pageTimeout must be less than the value of timeout.
	//
	// Default value: 10000.
	//
	// example:
	//
	// 10000
	PageTimeout *int32 `json:"pageTimeout,omitempty" xml:"pageTimeout,omitempty"`
	// The readability configuration for the parsing result.
	Readability *ReadPageBodyReadability `json:"readability,omitempty" xml:"readability,omitempty" type:"Struct"`
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
	// The target URL to parse. The value must start with http:// or https://.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://help.aliyun.com/document_detail/2837301.html?spm=a2c4g.11186623.help-menu-2837261.d_0_0_0.59ed3e95CppOt2&scm=20140722.H_2837301._.OR_help-T_cn~zh-V_1
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ReadPageBody) String() string {
	return dara.Prettify(s)
}

func (s ReadPageBody) GoString() string {
	return s.String()
}

func (s *ReadPageBody) GetFormats() []*string {
	return s.Formats
}

func (s *ReadPageBody) GetLocation() *string {
	return s.Location
}

func (s *ReadPageBody) GetMaxAge() *int32 {
	return s.MaxAge
}

func (s *ReadPageBody) GetPageTimeout() *int32 {
	return s.PageTimeout
}

func (s *ReadPageBody) GetReadability() *ReadPageBodyReadability {
	return s.Readability
}

func (s *ReadPageBody) GetTimeout() *int32 {
	return s.Timeout
}

func (s *ReadPageBody) GetUrl() *string {
	return s.Url
}

func (s *ReadPageBody) SetFormats(v []*string) *ReadPageBody {
	s.Formats = v
	return s
}

func (s *ReadPageBody) SetLocation(v string) *ReadPageBody {
	s.Location = &v
	return s
}

func (s *ReadPageBody) SetMaxAge(v int32) *ReadPageBody {
	s.MaxAge = &v
	return s
}

func (s *ReadPageBody) SetPageTimeout(v int32) *ReadPageBody {
	s.PageTimeout = &v
	return s
}

func (s *ReadPageBody) SetReadability(v *ReadPageBodyReadability) *ReadPageBody {
	s.Readability = v
	return s
}

func (s *ReadPageBody) SetTimeout(v int32) *ReadPageBody {
	s.Timeout = &v
	return s
}

func (s *ReadPageBody) SetUrl(v string) *ReadPageBody {
	s.Url = &v
	return s
}

func (s *ReadPageBody) Validate() error {
	if s.Readability != nil {
		if err := s.Readability.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReadPageBodyReadability struct {
	// Specifies whether to exclude all images.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	ExcludeAllImages *bool `json:"excludeAllImages,omitempty" xml:"excludeAllImages,omitempty"`
	// Specifies whether to exclude all links.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	ExcludeAllLinks *bool `json:"excludeAllLinks,omitempty" xml:"excludeAllLinks,omitempty"`
	// The tags to exclude.
	ExcludedTags []*string `json:"excludedTags,omitempty" xml:"excludedTags,omitempty" type:"Repeated"`
	// Valid values:
	//
	// - none: does not remove any information. Default value: none.
	//
	// - normal: removes irrelevant information from the target page, such as headers, footers, and navigation elements, based on a proprietary algorithm.
	//
	// - article: extracts the main body content of the website based on a proprietary algorithm. This mode is suitable for blogs and news websites, but not for directory or navigation pages.
	//
	// example:
	//
	// none
	ReadabilityMode *string `json:"readabilityMode,omitempty" xml:"readabilityMode,omitempty"`
}

func (s ReadPageBodyReadability) String() string {
	return dara.Prettify(s)
}

func (s ReadPageBodyReadability) GoString() string {
	return s.String()
}

func (s *ReadPageBodyReadability) GetExcludeAllImages() *bool {
	return s.ExcludeAllImages
}

func (s *ReadPageBodyReadability) GetExcludeAllLinks() *bool {
	return s.ExcludeAllLinks
}

func (s *ReadPageBodyReadability) GetExcludedTags() []*string {
	return s.ExcludedTags
}

func (s *ReadPageBodyReadability) GetReadabilityMode() *string {
	return s.ReadabilityMode
}

func (s *ReadPageBodyReadability) SetExcludeAllImages(v bool) *ReadPageBodyReadability {
	s.ExcludeAllImages = &v
	return s
}

func (s *ReadPageBodyReadability) SetExcludeAllLinks(v bool) *ReadPageBodyReadability {
	s.ExcludeAllLinks = &v
	return s
}

func (s *ReadPageBodyReadability) SetExcludedTags(v []*string) *ReadPageBodyReadability {
	s.ExcludedTags = v
	return s
}

func (s *ReadPageBodyReadability) SetReadabilityMode(v string) *ReadPageBodyReadability {
	s.ReadabilityMode = &v
	return s
}

func (s *ReadPageBodyReadability) Validate() error {
	return dara.Validate(s)
}
