// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChapterSummary interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ChapterSummary
	GetPageNumber() *int32
	SetSummary(v []*Summary) *ChapterSummary
	GetSummary() []*Summary
	SetTimeRange(v []*int64) *ChapterSummary
	GetTimeRange() []*int64
	SetTitle(v string) *ChapterSummary
	GetTitle() *string
	SetTitleID(v string) *ChapterSummary
	GetTitleID() *string
}

type ChapterSummary struct {
	PageNumber *int32     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Summary    []*Summary `json:"Summary,omitempty" xml:"Summary,omitempty" type:"Repeated"`
	TimeRange  []*int64   `json:"TimeRange,omitempty" xml:"TimeRange,omitempty" type:"Repeated"`
	Title      *string    `json:"Title,omitempty" xml:"Title,omitempty"`
	TitleID    *string    `json:"TitleID,omitempty" xml:"TitleID,omitempty"`
}

func (s ChapterSummary) String() string {
	return dara.Prettify(s)
}

func (s ChapterSummary) GoString() string {
	return s.String()
}

func (s *ChapterSummary) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ChapterSummary) GetSummary() []*Summary {
	return s.Summary
}

func (s *ChapterSummary) GetTimeRange() []*int64 {
	return s.TimeRange
}

func (s *ChapterSummary) GetTitle() *string {
	return s.Title
}

func (s *ChapterSummary) GetTitleID() *string {
	return s.TitleID
}

func (s *ChapterSummary) SetPageNumber(v int32) *ChapterSummary {
	s.PageNumber = &v
	return s
}

func (s *ChapterSummary) SetSummary(v []*Summary) *ChapterSummary {
	s.Summary = v
	return s
}

func (s *ChapterSummary) SetTimeRange(v []*int64) *ChapterSummary {
	s.TimeRange = v
	return s
}

func (s *ChapterSummary) SetTitle(v string) *ChapterSummary {
	s.Title = &v
	return s
}

func (s *ChapterSummary) SetTitleID(v string) *ChapterSummary {
	s.TitleID = &v
	return s
}

func (s *ChapterSummary) Validate() error {
	if s.Summary != nil {
		for _, item := range s.Summary {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
