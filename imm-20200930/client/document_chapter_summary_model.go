// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocumentChapterSummary interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ChapterSummary) *DocumentChapterSummary
	GetData() []*ChapterSummary
	SetNextMarker(v int32) *DocumentChapterSummary
	GetNextMarker() *int32
	SetVersion(v string) *DocumentChapterSummary
	GetVersion() *string
}

type DocumentChapterSummary struct {
	Data       []*ChapterSummary `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	NextMarker *int32            `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	Version    *string           `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DocumentChapterSummary) String() string {
	return dara.Prettify(s)
}

func (s DocumentChapterSummary) GoString() string {
	return s.String()
}

func (s *DocumentChapterSummary) GetData() []*ChapterSummary {
	return s.Data
}

func (s *DocumentChapterSummary) GetNextMarker() *int32 {
	return s.NextMarker
}

func (s *DocumentChapterSummary) GetVersion() *string {
	return s.Version
}

func (s *DocumentChapterSummary) SetData(v []*ChapterSummary) *DocumentChapterSummary {
	s.Data = v
	return s
}

func (s *DocumentChapterSummary) SetNextMarker(v int32) *DocumentChapterSummary {
	s.NextMarker = &v
	return s
}

func (s *DocumentChapterSummary) SetVersion(v string) *DocumentChapterSummary {
	s.Version = &v
	return s
}

func (s *DocumentChapterSummary) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
