// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLabel interface {
	dara.Model
	String() string
	GoString() string
	SetCentricScore(v float32) *Label
	GetCentricScore() *float32
	SetLabelConfidence(v float32) *Label
	GetLabelConfidence() *float32
	SetLabelLevel(v int64) *Label
	GetLabelLevel() *int64
	SetLabelName(v string) *Label
	GetLabelName() *string
	SetLanguage(v string) *Label
	GetLanguage() *string
	SetParentLabelName(v string) *Label
	GetParentLabelName() *string
}

type Label struct {
	CentricScore    *float32 `json:"CentricScore,omitempty" xml:"CentricScore,omitempty"`
	LabelConfidence *float32 `json:"LabelConfidence,omitempty" xml:"LabelConfidence,omitempty"`
	LabelLevel      *int64   `json:"LabelLevel,omitempty" xml:"LabelLevel,omitempty"`
	LabelName       *string  `json:"LabelName,omitempty" xml:"LabelName,omitempty"`
	Language        *string  `json:"Language,omitempty" xml:"Language,omitempty"`
	ParentLabelName *string  `json:"ParentLabelName,omitempty" xml:"ParentLabelName,omitempty"`
}

func (s Label) String() string {
	return dara.Prettify(s)
}

func (s Label) GoString() string {
	return s.String()
}

func (s *Label) GetCentricScore() *float32 {
	return s.CentricScore
}

func (s *Label) GetLabelConfidence() *float32 {
	return s.LabelConfidence
}

func (s *Label) GetLabelLevel() *int64 {
	return s.LabelLevel
}

func (s *Label) GetLabelName() *string {
	return s.LabelName
}

func (s *Label) GetLanguage() *string {
	return s.Language
}

func (s *Label) GetParentLabelName() *string {
	return s.ParentLabelName
}

func (s *Label) SetCentricScore(v float32) *Label {
	s.CentricScore = &v
	return s
}

func (s *Label) SetLabelConfidence(v float32) *Label {
	s.LabelConfidence = &v
	return s
}

func (s *Label) SetLabelLevel(v int64) *Label {
	s.LabelLevel = &v
	return s
}

func (s *Label) SetLabelName(v string) *Label {
	s.LabelName = &v
	return s
}

func (s *Label) SetLanguage(v string) *Label {
	s.Language = &v
	return s
}

func (s *Label) SetParentLabelName(v string) *Label {
	s.ParentLabelName = &v
	return s
}

func (s *Label) Validate() error {
	return dara.Validate(s)
}
