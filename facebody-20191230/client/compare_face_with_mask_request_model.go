// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompareFaceWithMaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLA(v string) *CompareFaceWithMaskRequest
	GetImageURLA() *string
	SetImageURLB(v string) *CompareFaceWithMaskRequest
	GetImageURLB() *string
	SetQualityScoreThreshold(v float32) *CompareFaceWithMaskRequest
	GetQualityScoreThreshold() *float32
}

type CompareFaceWithMaskRequest struct {
	// This parameter is required.
	ImageURLA *string `json:"ImageURLA,omitempty" xml:"ImageURLA,omitempty"`
	// This parameter is required.
	ImageURLB *string `json:"ImageURLB,omitempty" xml:"ImageURLB,omitempty"`
	// example:
	//
	// 97.0
	QualityScoreThreshold *float32 `json:"QualityScoreThreshold,omitempty" xml:"QualityScoreThreshold,omitempty"`
}

func (s CompareFaceWithMaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CompareFaceWithMaskRequest) GoString() string {
	return s.String()
}

func (s *CompareFaceWithMaskRequest) GetImageURLA() *string {
	return s.ImageURLA
}

func (s *CompareFaceWithMaskRequest) GetImageURLB() *string {
	return s.ImageURLB
}

func (s *CompareFaceWithMaskRequest) GetQualityScoreThreshold() *float32 {
	return s.QualityScoreThreshold
}

func (s *CompareFaceWithMaskRequest) SetImageURLA(v string) *CompareFaceWithMaskRequest {
	s.ImageURLA = &v
	return s
}

func (s *CompareFaceWithMaskRequest) SetImageURLB(v string) *CompareFaceWithMaskRequest {
	s.ImageURLB = &v
	return s
}

func (s *CompareFaceWithMaskRequest) SetQualityScoreThreshold(v float32) *CompareFaceWithMaskRequest {
	s.QualityScoreThreshold = &v
	return s
}

func (s *CompareFaceWithMaskRequest) Validate() error {
	return dara.Validate(s)
}
