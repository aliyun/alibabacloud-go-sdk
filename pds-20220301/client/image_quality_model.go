// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageQuality interface {
	dara.Model
	String() string
	GoString() string
	SetOverallScore(v float64) *ImageQuality
	GetOverallScore() *float64
}

type ImageQuality struct {
	// The overall quality score of the image. The image is automatically evaluated by AI. The evaluation is mainly based on subjective aesthetics and is affected by various factors, such as composition, brightness, contrast, color, and definition. Valid values: 0 to 1. The higher the score, the better the quality.
	//
	// example:
	//
	// 0.736
	OverallScore *float64 `json:"overall_score,omitempty" xml:"overall_score,omitempty"`
}

func (s ImageQuality) String() string {
	return dara.Prettify(s)
}

func (s ImageQuality) GoString() string {
	return s.String()
}

func (s *ImageQuality) GetOverallScore() *float64 {
	return s.OverallScore
}

func (s *ImageQuality) SetOverallScore(v float64) *ImageQuality {
	s.OverallScore = &v
	return s
}

func (s *ImageQuality) Validate() error {
	return dara.Validate(s)
}
