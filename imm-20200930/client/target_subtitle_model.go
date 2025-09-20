// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTargetSubtitle interface {
	dara.Model
	String() string
	GoString() string
	SetDisableSubtitle(v bool) *TargetSubtitle
	GetDisableSubtitle() *bool
	SetExtractSubtitle(v *TargetSubtitleExtractSubtitle) *TargetSubtitle
	GetExtractSubtitle() *TargetSubtitleExtractSubtitle
	SetStream(v []*int32) *TargetSubtitle
	GetStream() []*int32
}

type TargetSubtitle struct {
	DisableSubtitle *bool                          `json:"DisableSubtitle,omitempty" xml:"DisableSubtitle,omitempty"`
	ExtractSubtitle *TargetSubtitleExtractSubtitle `json:"ExtractSubtitle,omitempty" xml:"ExtractSubtitle,omitempty" type:"Struct"`
	Stream          []*int32                       `json:"Stream,omitempty" xml:"Stream,omitempty" type:"Repeated"`
}

func (s TargetSubtitle) String() string {
	return dara.Prettify(s)
}

func (s TargetSubtitle) GoString() string {
	return s.String()
}

func (s *TargetSubtitle) GetDisableSubtitle() *bool {
	return s.DisableSubtitle
}

func (s *TargetSubtitle) GetExtractSubtitle() *TargetSubtitleExtractSubtitle {
	return s.ExtractSubtitle
}

func (s *TargetSubtitle) GetStream() []*int32 {
	return s.Stream
}

func (s *TargetSubtitle) SetDisableSubtitle(v bool) *TargetSubtitle {
	s.DisableSubtitle = &v
	return s
}

func (s *TargetSubtitle) SetExtractSubtitle(v *TargetSubtitleExtractSubtitle) *TargetSubtitle {
	s.ExtractSubtitle = v
	return s
}

func (s *TargetSubtitle) SetStream(v []*int32) *TargetSubtitle {
	s.Stream = v
	return s
}

func (s *TargetSubtitle) Validate() error {
	return dara.Validate(s)
}

type TargetSubtitleExtractSubtitle struct {
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	URI    *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s TargetSubtitleExtractSubtitle) String() string {
	return dara.Prettify(s)
}

func (s TargetSubtitleExtractSubtitle) GoString() string {
	return s.String()
}

func (s *TargetSubtitleExtractSubtitle) GetFormat() *string {
	return s.Format
}

func (s *TargetSubtitleExtractSubtitle) GetURI() *string {
	return s.URI
}

func (s *TargetSubtitleExtractSubtitle) SetFormat(v string) *TargetSubtitleExtractSubtitle {
	s.Format = &v
	return s
}

func (s *TargetSubtitleExtractSubtitle) SetURI(v string) *TargetSubtitleExtractSubtitle {
	s.URI = &v
	return s
}

func (s *TargetSubtitleExtractSubtitle) Validate() error {
	return dara.Validate(s)
}
