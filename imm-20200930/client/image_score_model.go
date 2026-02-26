// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageScore interface {
	dara.Model
	String() string
	GoString() string
	SetOverallQualityScore(v float32) *ImageScore
	GetOverallQualityScore() *float32
}

type ImageScore struct {
	// The score for the overall image quality. The image is automatically evaluated by AI. The evaluation is mainly based on subjective aesthetics and is affected by various factors, such as composition, brightness, contrast, color, and definition.
	//
	// Valid values: 0 to 1. A higher value indicates better quality.
	//
	// example:
	//
	// 0.736
	OverallQualityScore *float32 `json:"OverallQualityScore,omitempty" xml:"OverallQualityScore,omitempty"`
}

func (s ImageScore) String() string {
	return dara.Prettify(s)
}

func (s ImageScore) GoString() string {
	return s.String()
}

func (s *ImageScore) GetOverallQualityScore() *float32 {
	return s.OverallQualityScore
}

func (s *ImageScore) SetOverallQualityScore(v float32) *ImageScore {
	s.OverallQualityScore = &v
	return s
}

func (s *ImageScore) Validate() error {
	return dara.Validate(s)
}
