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
	SetClips(v []*Clip) *Label
	GetClips() []*Clip
	SetLabelAlias(v string) *Label
	GetLabelAlias() *string
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
	// The central value of the label. This value indicates the confidence that the label is the majority component of the image. Valid values: 0 to 1. A higher value indicates greater confidence.
	//
	// example:
	//
	// 0.7319999933242798
	CentricScore *float32 `json:"CentricScore,omitempty" xml:"CentricScore,omitempty"`
	Clips        []*Clip  `json:"Clips,omitempty" xml:"Clips,omitempty" type:"Repeated"`
	LabelAlias   *string  `json:"LabelAlias,omitempty" xml:"LabelAlias,omitempty"`
	// The confidence level of the label. Valid values: 0 to 1. A higher value indicates greater confidence.
	//
	// example:
	//
	// 0.9891784601980591
	LabelConfidence *float32 `json:"LabelConfidence,omitempty" xml:"LabelConfidence,omitempty"`
	// The label level. Valid values: 1, 2, and 3.
	//
	// example:
	//
	// 1
	LabelLevel *int64 `json:"LabelLevel,omitempty" xml:"LabelLevel,omitempty"`
	// The label name.
	LabelName *string `json:"LabelName,omitempty" xml:"LabelName,omitempty"`
	// The label language, which is represented as a BCP 47 language tag.
	//
	// example:
	//
	// zh-Hans
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The name of the parent label.
	ParentLabelName *string `json:"ParentLabelName,omitempty" xml:"ParentLabelName,omitempty"`
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

func (s *Label) GetClips() []*Clip {
	return s.Clips
}

func (s *Label) GetLabelAlias() *string {
	return s.LabelAlias
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

func (s *Label) SetClips(v []*Clip) *Label {
	s.Clips = v
	return s
}

func (s *Label) SetLabelAlias(v string) *Label {
	s.LabelAlias = &v
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
	if s.Clips != nil {
		for _, item := range s.Clips {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
