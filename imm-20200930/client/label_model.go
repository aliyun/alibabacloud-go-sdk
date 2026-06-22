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
	// The centric score of the tag. This indicates whether the tag is the main subject in the image. The value ranges from 0 to 1. A higher value indicates higher confidence that the tag is the main subject of the image.
	//
	// example:
	//
	// 0.877
	CentricScore *float32 `json:"CentricScore,omitempty" xml:"CentricScore,omitempty"`
	// Event clips.
	Clips []*Clip `json:"Clips,omitempty" xml:"Clips,omitempty" type:"Repeated"`
	// The tag alias.
	//
	// example:
	//
	// 座椅
	LabelAlias *string `json:"LabelAlias,omitempty" xml:"LabelAlias,omitempty"`
	// The tag confidence level. The value ranges from 0 (lowest confidence) to 1 (highest confidence).
	//
	// example:
	//
	// 0.95
	LabelConfidence *float32 `json:"LabelConfidence,omitempty" xml:"LabelConfidence,omitempty"`
	// The tag level. Valid values are 1, 2, and 3, representing first-level, second-level, and third-level tags, respectively.
	//
	// example:
	//
	// 2
	LabelLevel *int64 `json:"LabelLevel,omitempty" xml:"LabelLevel,omitempty"`
	// The tag name.
	//
	// example:
	//
	// 椅子
	LabelName *string `json:"LabelName,omitempty" xml:"LabelName,omitempty"`
	// The tag language, in BCP 47 format.
	//
	// example:
	//
	// zh-Hans
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The parent tag name.
	//
	// example:
	//
	// 家具
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
