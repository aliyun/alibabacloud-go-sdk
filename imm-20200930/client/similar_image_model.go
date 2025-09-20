// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSimilarImage interface {
	dara.Model
	String() string
	GoString() string
	SetImageScore(v float64) *SimilarImage
	GetImageScore() *float64
	SetURI(v string) *SimilarImage
	GetURI() *string
}

type SimilarImage struct {
	ImageScore *float64 `json:"ImageScore,omitempty" xml:"ImageScore,omitempty"`
	URI        *string  `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s SimilarImage) String() string {
	return dara.Prettify(s)
}

func (s SimilarImage) GoString() string {
	return s.String()
}

func (s *SimilarImage) GetImageScore() *float64 {
	return s.ImageScore
}

func (s *SimilarImage) GetURI() *string {
	return s.URI
}

func (s *SimilarImage) SetImageScore(v float64) *SimilarImage {
	s.ImageScore = &v
	return s
}

func (s *SimilarImage) SetURI(v string) *SimilarImage {
	s.URI = &v
	return s
}

func (s *SimilarImage) Validate() error {
	return dara.Validate(s)
}
