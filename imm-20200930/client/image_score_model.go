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
