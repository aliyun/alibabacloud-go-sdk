// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTopicSelection interface {
	dara.Model
	String() string
	GoString() string
	SetOutlines(v []*TopicSelectionOutlines) *TopicSelection
	GetOutlines() []*TopicSelectionOutlines
	SetPoint(v string) *TopicSelection
	GetPoint() *string
	SetSummary(v string) *TopicSelection
	GetSummary() *string
}

type TopicSelection struct {
	Outlines []*TopicSelectionOutlines `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	Point    *string                   `json:"Point,omitempty" xml:"Point,omitempty"`
	Summary  *string                   `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s TopicSelection) String() string {
	return dara.Prettify(s)
}

func (s TopicSelection) GoString() string {
	return s.String()
}

func (s *TopicSelection) GetOutlines() []*TopicSelectionOutlines {
	return s.Outlines
}

func (s *TopicSelection) GetPoint() *string {
	return s.Point
}

func (s *TopicSelection) GetSummary() *string {
	return s.Summary
}

func (s *TopicSelection) SetOutlines(v []*TopicSelectionOutlines) *TopicSelection {
	s.Outlines = v
	return s
}

func (s *TopicSelection) SetPoint(v string) *TopicSelection {
	s.Point = &v
	return s
}

func (s *TopicSelection) SetSummary(v string) *TopicSelection {
	s.Summary = &v
	return s
}

func (s *TopicSelection) Validate() error {
	if s.Outlines != nil {
		for _, item := range s.Outlines {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TopicSelectionOutlines struct {
	Outline *string `json:"Outline,omitempty" xml:"Outline,omitempty"`
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s TopicSelectionOutlines) String() string {
	return dara.Prettify(s)
}

func (s TopicSelectionOutlines) GoString() string {
	return s.String()
}

func (s *TopicSelectionOutlines) GetOutline() *string {
	return s.Outline
}

func (s *TopicSelectionOutlines) GetSummary() *string {
	return s.Summary
}

func (s *TopicSelectionOutlines) SetOutline(v string) *TopicSelectionOutlines {
	s.Outline = &v
	return s
}

func (s *TopicSelectionOutlines) SetSummary(v string) *TopicSelectionOutlines {
	s.Summary = &v
	return s
}

func (s *TopicSelectionOutlines) Validate() error {
	return dara.Validate(s)
}
