// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageInsight interface {
	dara.Model
	String() string
	GoString() string
	SetCaption(v string) *ImageInsight
	GetCaption() *string
	SetDescription(v string) *ImageInsight
	GetDescription() *string
	SetMultilingualContent(v map[string]*MultilingualContentEntry) *ImageInsight
	GetMultilingualContent() map[string]*MultilingualContentEntry
}

type ImageInsight struct {
	// if can be null:
	// true
	Caption *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	// if can be null:
	// true
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The multilingual image content.
	MultilingualContent map[string]*MultilingualContentEntry `json:"MultilingualContent,omitempty" xml:"MultilingualContent,omitempty"`
}

func (s ImageInsight) String() string {
	return dara.Prettify(s)
}

func (s ImageInsight) GoString() string {
	return s.String()
}

func (s *ImageInsight) GetCaption() *string {
	return s.Caption
}

func (s *ImageInsight) GetDescription() *string {
	return s.Description
}

func (s *ImageInsight) GetMultilingualContent() map[string]*MultilingualContentEntry {
	return s.MultilingualContent
}

func (s *ImageInsight) SetCaption(v string) *ImageInsight {
	s.Caption = &v
	return s
}

func (s *ImageInsight) SetDescription(v string) *ImageInsight {
	s.Description = &v
	return s
}

func (s *ImageInsight) SetMultilingualContent(v map[string]*MultilingualContentEntry) *ImageInsight {
	s.MultilingualContent = v
	return s
}

func (s *ImageInsight) Validate() error {
	return dara.Validate(s)
}
