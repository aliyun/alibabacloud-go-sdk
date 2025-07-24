// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRequestContents interface {
	dara.Model
	String() string
	GoString() string
	SetMainText(v bool) *RequestContents
	GetMainText() *bool
	SetMarkdownText(v bool) *RequestContents
	GetMarkdownText() *bool
	SetRerankScore(v bool) *RequestContents
	GetRerankScore() *bool
	SetSummary(v bool) *RequestContents
	GetSummary() *bool
}

type RequestContents struct {
	MainText     *bool `json:"mainText,omitempty" xml:"mainText,omitempty"`
	MarkdownText *bool `json:"markdownText,omitempty" xml:"markdownText,omitempty"`
	RerankScore  *bool `json:"rerankScore,omitempty" xml:"rerankScore,omitempty"`
	Summary      *bool `json:"summary,omitempty" xml:"summary,omitempty"`
}

func (s RequestContents) String() string {
	return dara.Prettify(s)
}

func (s RequestContents) GoString() string {
	return s.String()
}

func (s *RequestContents) GetMainText() *bool {
	return s.MainText
}

func (s *RequestContents) GetMarkdownText() *bool {
	return s.MarkdownText
}

func (s *RequestContents) GetRerankScore() *bool {
	return s.RerankScore
}

func (s *RequestContents) GetSummary() *bool {
	return s.Summary
}

func (s *RequestContents) SetMainText(v bool) *RequestContents {
	s.MainText = &v
	return s
}

func (s *RequestContents) SetMarkdownText(v bool) *RequestContents {
	s.MarkdownText = &v
	return s
}

func (s *RequestContents) SetRerankScore(v bool) *RequestContents {
	s.RerankScore = &v
	return s
}

func (s *RequestContents) SetSummary(v bool) *RequestContents {
	s.Summary = &v
	return s
}

func (s *RequestContents) Validate() error {
	return dara.Validate(s)
}
