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
	Formats []*string `json:"formats,omitempty" xml:"formats,omitempty" type:"Repeated"`
	// example:
	//
	// null
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// example:
	//
	// 1296000
	MaxAge *int32 `json:"maxAge,omitempty" xml:"maxAge,omitempty"`
	// example:
	//
	// 15000
	PageTimeout *int32                         `json:"pageTimeout,omitempty" xml:"pageTimeout,omitempty"`
	Readability *ReadPageScrapeBodyReadability `json:"readability,omitempty" xml:"readability,omitempty" type:"Struct"`
	// example:
	//
	// 60000
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
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
	// example:
	//
	// false
	ExcludeAllImages *bool `json:"excludeAllImages,omitempty" xml:"excludeAllImages,omitempty"`
	// example:
	//
	// false
	ExcludeAllLinks *bool     `json:"excludeAllLinks,omitempty" xml:"excludeAllLinks,omitempty"`
	ExcludedTags    []*string `json:"excludedTags,omitempty" xml:"excludedTags,omitempty" type:"Repeated"`
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
