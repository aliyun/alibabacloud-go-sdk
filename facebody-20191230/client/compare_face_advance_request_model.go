// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iCompareFaceAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageDataA(v string) *CompareFaceAdvanceRequest
	GetImageDataA() *string
	SetImageDataB(v string) *CompareFaceAdvanceRequest
	GetImageDataB() *string
	SetImageURLAObject(v io.Reader) *CompareFaceAdvanceRequest
	GetImageURLAObject() io.Reader
	SetImageURLBObject(v io.Reader) *CompareFaceAdvanceRequest
	GetImageURLBObject() io.Reader
	SetQualityScoreThreshold(v float32) *CompareFaceAdvanceRequest
	GetQualityScoreThreshold() *float32
}

type CompareFaceAdvanceRequest struct {
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
	ImageURLAObject io.Reader `json:"ImageURLA,omitempty" xml:"ImageURLA,omitempty"`
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/CompareFace/CompareFace-left1.png
	ImageURLBObject io.Reader `json:"ImageURLB,omitempty" xml:"ImageURLB,omitempty"`
	// example:
	//
	// 75.12
	QualityScoreThreshold *float32 `json:"QualityScoreThreshold,omitempty" xml:"QualityScoreThreshold,omitempty"`
}

func (s CompareFaceAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CompareFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *CompareFaceAdvanceRequest) GetImageDataA() *string {
	return s.ImageDataA
}

func (s *CompareFaceAdvanceRequest) GetImageDataB() *string {
	return s.ImageDataB
}

func (s *CompareFaceAdvanceRequest) GetImageURLAObject() io.Reader {
	return s.ImageURLAObject
}

func (s *CompareFaceAdvanceRequest) GetImageURLBObject() io.Reader {
	return s.ImageURLBObject
}

func (s *CompareFaceAdvanceRequest) GetQualityScoreThreshold() *float32 {
	return s.QualityScoreThreshold
}

func (s *CompareFaceAdvanceRequest) SetImageDataA(v string) *CompareFaceAdvanceRequest {
	s.ImageDataA = &v
	return s
}

func (s *CompareFaceAdvanceRequest) SetImageDataB(v string) *CompareFaceAdvanceRequest {
	s.ImageDataB = &v
	return s
}

func (s *CompareFaceAdvanceRequest) SetImageURLAObject(v io.Reader) *CompareFaceAdvanceRequest {
	s.ImageURLAObject = v
	return s
}

func (s *CompareFaceAdvanceRequest) SetImageURLBObject(v io.Reader) *CompareFaceAdvanceRequest {
	s.ImageURLBObject = v
	return s
}

func (s *CompareFaceAdvanceRequest) SetQualityScoreThreshold(v float32) *CompareFaceAdvanceRequest {
	s.QualityScoreThreshold = &v
	return s
}

func (s *CompareFaceAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
