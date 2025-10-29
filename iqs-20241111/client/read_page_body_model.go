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
	Formats     []*string                `json:"formats,omitempty" xml:"formats,omitempty" type:"Repeated"`
	Location    *string                  `json:"location,omitempty" xml:"location,omitempty"`
	MaxAge      *int32                   `json:"maxAge,omitempty" xml:"maxAge,omitempty"`
	PageTimeout *int32                   `json:"pageTimeout,omitempty" xml:"pageTimeout,omitempty"`
	Readability *ReadPageBodyReadability `json:"readability,omitempty" xml:"readability,omitempty" type:"Struct"`
	Timeout     *int32                   `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// This parameter is required.
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
	ExcludeAllImages *bool     `json:"excludeAllImages,omitempty" xml:"excludeAllImages,omitempty"`
	ExcludeAllLinks  *bool     `json:"excludeAllLinks,omitempty" xml:"excludeAllLinks,omitempty"`
	ExcludedTags     []*string `json:"excludedTags,omitempty" xml:"excludedTags,omitempty" type:"Repeated"`
	ReadabilityMode  *string   `json:"readabilityMode,omitempty" xml:"readabilityMode,omitempty"`
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
