// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iCompareFaceWithMaskAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLAObject(v io.Reader) *CompareFaceWithMaskAdvanceRequest
	GetImageURLAObject() io.Reader
	SetImageURLBObject(v io.Reader) *CompareFaceWithMaskAdvanceRequest
	GetImageURLBObject() io.Reader
	SetQualityScoreThreshold(v float32) *CompareFaceWithMaskAdvanceRequest
	GetQualityScoreThreshold() *float32
}

type CompareFaceWithMaskAdvanceRequest struct {
	// This parameter is required.
	ImageURLAObject io.Reader `json:"ImageURLA,omitempty" xml:"ImageURLA,omitempty"`
	// This parameter is required.
	ImageURLBObject io.Reader `json:"ImageURLB,omitempty" xml:"ImageURLB,omitempty"`
	// example:
	//
	// 97.0
	QualityScoreThreshold *float32 `json:"QualityScoreThreshold,omitempty" xml:"QualityScoreThreshold,omitempty"`
}

func (s CompareFaceWithMaskAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CompareFaceWithMaskAdvanceRequest) GoString() string {
	return s.String()
}

func (s *CompareFaceWithMaskAdvanceRequest) GetImageURLAObject() io.Reader {
	return s.ImageURLAObject
}

func (s *CompareFaceWithMaskAdvanceRequest) GetImageURLBObject() io.Reader {
	return s.ImageURLBObject
}

func (s *CompareFaceWithMaskAdvanceRequest) GetQualityScoreThreshold() *float32 {
	return s.QualityScoreThreshold
}

func (s *CompareFaceWithMaskAdvanceRequest) SetImageURLAObject(v io.Reader) *CompareFaceWithMaskAdvanceRequest {
	s.ImageURLAObject = v
	return s
}

func (s *CompareFaceWithMaskAdvanceRequest) SetImageURLBObject(v io.Reader) *CompareFaceWithMaskAdvanceRequest {
	s.ImageURLBObject = v
	return s
}

func (s *CompareFaceWithMaskAdvanceRequest) SetQualityScoreThreshold(v float32) *CompareFaceWithMaskAdvanceRequest {
	s.QualityScoreThreshold = &v
	return s
}

func (s *CompareFaceWithMaskAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
