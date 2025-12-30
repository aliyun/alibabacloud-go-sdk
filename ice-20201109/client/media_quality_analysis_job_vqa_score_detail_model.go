// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMediaQualityAnalysisJobVqaScoreDetail interface {
	dara.Model
	String() string
	GoString() string
	SetIntensityValue(v float64) *MediaQualityAnalysisJobVqaScoreDetail
	GetIntensityValue() *float64
	SetPerceptualScore(v float64) *MediaQualityAnalysisJobVqaScoreDetail
	GetPerceptualScore() *float64
}

type MediaQualityAnalysisJobVqaScoreDetail struct {
	IntensityValue  *float64 `json:"IntensityValue,omitempty" xml:"IntensityValue,omitempty"`
	PerceptualScore *float64 `json:"PerceptualScore,omitempty" xml:"PerceptualScore,omitempty"`
}

func (s MediaQualityAnalysisJobVqaScoreDetail) String() string {
	return dara.Prettify(s)
}

func (s MediaQualityAnalysisJobVqaScoreDetail) GoString() string {
	return s.String()
}

func (s *MediaQualityAnalysisJobVqaScoreDetail) GetIntensityValue() *float64 {
	return s.IntensityValue
}

func (s *MediaQualityAnalysisJobVqaScoreDetail) GetPerceptualScore() *float64 {
	return s.PerceptualScore
}

func (s *MediaQualityAnalysisJobVqaScoreDetail) SetIntensityValue(v float64) *MediaQualityAnalysisJobVqaScoreDetail {
	s.IntensityValue = &v
	return s
}

func (s *MediaQualityAnalysisJobVqaScoreDetail) SetPerceptualScore(v float64) *MediaQualityAnalysisJobVqaScoreDetail {
	s.PerceptualScore = &v
	return s
}

func (s *MediaQualityAnalysisJobVqaScoreDetail) Validate() error {
	return dara.Validate(s)
}
