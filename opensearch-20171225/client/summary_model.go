// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSummary interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v bool) *Summary
	GetActive() *bool
	SetMeta(v *SummaryMeta) *Summary
	GetMeta() *SummaryMeta
	SetName(v string) *Summary
	GetName() *string
}

type Summary struct {
	Active *bool        `json:"active,omitempty" xml:"active,omitempty"`
	Meta   *SummaryMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
	Name   *string      `json:"name,omitempty" xml:"name,omitempty"`
}

func (s Summary) String() string {
	return dara.Prettify(s)
}

func (s Summary) GoString() string {
	return s.String()
}

func (s *Summary) GetActive() *bool {
	return s.Active
}

func (s *Summary) GetMeta() *SummaryMeta {
	return s.Meta
}

func (s *Summary) GetName() *string {
	return s.Name
}

func (s *Summary) SetActive(v bool) *Summary {
	s.Active = &v
	return s
}

func (s *Summary) SetMeta(v *SummaryMeta) *Summary {
	s.Meta = v
	return s
}

func (s *Summary) SetName(v string) *Summary {
	s.Name = &v
	return s
}

func (s *Summary) Validate() error {
	if s.Meta != nil {
		if err := s.Meta.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SummaryMeta struct {
	Element  *string `json:"element,omitempty" xml:"element,omitempty"`
	Ellipsis *string `json:"ellipsis,omitempty" xml:"ellipsis,omitempty"`
	Field    *string `json:"field,omitempty" xml:"field,omitempty"`
	Len      *int32  `json:"len,omitempty" xml:"len,omitempty"`
	Snippet  *string `json:"snippet,omitempty" xml:"snippet,omitempty"`
}

func (s SummaryMeta) String() string {
	return dara.Prettify(s)
}

func (s SummaryMeta) GoString() string {
	return s.String()
}

func (s *SummaryMeta) GetElement() *string {
	return s.Element
}

func (s *SummaryMeta) GetEllipsis() *string {
	return s.Ellipsis
}

func (s *SummaryMeta) GetField() *string {
	return s.Field
}

func (s *SummaryMeta) GetLen() *int32 {
	return s.Len
}

func (s *SummaryMeta) GetSnippet() *string {
	return s.Snippet
}

func (s *SummaryMeta) SetElement(v string) *SummaryMeta {
	s.Element = &v
	return s
}

func (s *SummaryMeta) SetEllipsis(v string) *SummaryMeta {
	s.Ellipsis = &v
	return s
}

func (s *SummaryMeta) SetField(v string) *SummaryMeta {
	s.Field = &v
	return s
}

func (s *SummaryMeta) SetLen(v int32) *SummaryMeta {
	s.Len = &v
	return s
}

func (s *SummaryMeta) SetSnippet(v string) *SummaryMeta {
	s.Snippet = &v
	return s
}

func (s *SummaryMeta) Validate() error {
	return dara.Validate(s)
}
