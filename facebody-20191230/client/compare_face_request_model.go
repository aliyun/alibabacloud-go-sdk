// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompareFaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageDataA(v string) *CompareFaceRequest
	GetImageDataA() *string
	SetImageDataB(v string) *CompareFaceRequest
	GetImageDataB() *string
	SetImageURLA(v string) *CompareFaceRequest
	GetImageURLA() *string
	SetImageURLB(v string) *CompareFaceRequest
	GetImageURLB() *string
	SetQualityScoreThreshold(v float32) *CompareFaceRequest
	GetQualityScoreThreshold() *float32
}

type CompareFaceRequest struct {
	// example:
	//
	// /9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAgGBgcGBQgHBwcJCQgK****
	ImageDataA *string `json:"ImageDataA,omitempty" xml:"ImageDataA,omitempty"`
	// example:
	//
	// /9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAgGBgcGBQgHBwcJCQgQ****
	ImageDataB *string `json:"ImageDataB,omitempty" xml:"ImageDataB,omitempty"`
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/CompareFace/CompareFace-right1.png
	ImageURLA *string `json:"ImageURLA,omitempty" xml:"ImageURLA,omitempty"`
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/CompareFace/CompareFace-left1.png
	ImageURLB *string `json:"ImageURLB,omitempty" xml:"ImageURLB,omitempty"`
	// example:
	//
	// 75.12
	QualityScoreThreshold *float32 `json:"QualityScoreThreshold,omitempty" xml:"QualityScoreThreshold,omitempty"`
}

func (s CompareFaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CompareFaceRequest) GoString() string {
	return s.String()
}

func (s *CompareFaceRequest) GetImageDataA() *string {
	return s.ImageDataA
}

func (s *CompareFaceRequest) GetImageDataB() *string {
	return s.ImageDataB
}

func (s *CompareFaceRequest) GetImageURLA() *string {
	return s.ImageURLA
}

func (s *CompareFaceRequest) GetImageURLB() *string {
	return s.ImageURLB
}

func (s *CompareFaceRequest) GetQualityScoreThreshold() *float32 {
	return s.QualityScoreThreshold
}

func (s *CompareFaceRequest) SetImageDataA(v string) *CompareFaceRequest {
	s.ImageDataA = &v
	return s
}

func (s *CompareFaceRequest) SetImageDataB(v string) *CompareFaceRequest {
	s.ImageDataB = &v
	return s
}

func (s *CompareFaceRequest) SetImageURLA(v string) *CompareFaceRequest {
	s.ImageURLA = &v
	return s
}

func (s *CompareFaceRequest) SetImageURLB(v string) *CompareFaceRequest {
	s.ImageURLB = &v
	return s
}

func (s *CompareFaceRequest) SetQualityScoreThreshold(v float32) *CompareFaceRequest {
	s.QualityScoreThreshold = &v
	return s
}

func (s *CompareFaceRequest) Validate() error {
	return dara.Validate(s)
}
