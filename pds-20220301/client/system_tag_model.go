// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSystemTag interface {
	dara.Model
	String() string
	GoString() string
	SetCentricScore(v float32) *SystemTag
	GetCentricScore() *float32
	SetConfidence(v float32) *SystemTag
	GetConfidence() *float32
	SetName(v string) *SystemTag
	GetName() *string
	SetParentName(v string) *SystemTag
	GetParentName() *string
	SetTagLevel(v int32) *SystemTag
	GetTagLevel() *int32
}

type SystemTag struct {
	// The center value of the tag, which specifies whether the tag is the subject in the image. Valid values: 0 to 1. A value of 0 specifies the lowest proportion. A value of 1 specifies the highest proportion.
	//
	// example:
	//
	// 0.877
	CentricScore *float32 `json:"centric_score,omitempty" xml:"centric_score,omitempty"`
	// The confidence level of the tag. Valid values: 0 to 1. A value of 0 specifies the lowest confidence level. A value of 1 specifies the highest confidence level.
	//
	// example:
	//
	// 0.98
	Confidence *float32 `json:"confidence,omitempty" xml:"confidence,omitempty"`
	// The name of the tag.
	//
	// example:
	//
	// basketball
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of the parent tag of the tag.
	//
	// example:
	//
	// sport
	ParentName *string `json:"parent_name,omitempty" xml:"parent_name,omitempty"`
	// The level of the tag. The value must be greater than or equal to 1.
	//
	// example:
	//
	// 3
	TagLevel *int32 `json:"tag_level,omitempty" xml:"tag_level,omitempty"`
}

func (s SystemTag) String() string {
	return dara.Prettify(s)
}

func (s SystemTag) GoString() string {
	return s.String()
}

func (s *SystemTag) GetCentricScore() *float32 {
	return s.CentricScore
}

func (s *SystemTag) GetConfidence() *float32 {
	return s.Confidence
}

func (s *SystemTag) GetName() *string {
	return s.Name
}

func (s *SystemTag) GetParentName() *string {
	return s.ParentName
}

func (s *SystemTag) GetTagLevel() *int32 {
	return s.TagLevel
}

func (s *SystemTag) SetCentricScore(v float32) *SystemTag {
	s.CentricScore = &v
	return s
}

func (s *SystemTag) SetConfidence(v float32) *SystemTag {
	s.Confidence = &v
	return s
}

func (s *SystemTag) SetName(v string) *SystemTag {
	s.Name = &v
	return s
}

func (s *SystemTag) SetParentName(v string) *SystemTag {
	s.ParentName = &v
	return s
}

func (s *SystemTag) SetTagLevel(v int32) *SystemTag {
	s.TagLevel = &v
	return s
}

func (s *SystemTag) Validate() error {
	return dara.Validate(s)
}
