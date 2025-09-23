// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iDetectWorkwearAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClothes(v *DetectWorkwearAdvanceRequestClothes) *DetectWorkwearAdvanceRequest
	GetClothes() *DetectWorkwearAdvanceRequestClothes
	SetImageUrlObject(v io.Reader) *DetectWorkwearAdvanceRequest
	GetImageUrlObject() io.Reader
	SetLabels(v []*string) *DetectWorkwearAdvanceRequest
	GetLabels() []*string
}

type DetectWorkwearAdvanceRequest struct {
	Clothes *DetectWorkwearAdvanceRequestClothes `json:"Clothes,omitempty" xml:"Clothes,omitempty" type:"Struct"`
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/test-team/zhangchaorun/tiyan/xxxx.jpg
	ImageUrlObject io.Reader `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 1
	Labels []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s DetectWorkwearAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectWorkwearAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectWorkwearAdvanceRequest) GetClothes() *DetectWorkwearAdvanceRequestClothes {
	return s.Clothes
}

func (s *DetectWorkwearAdvanceRequest) GetImageUrlObject() io.Reader {
	return s.ImageUrlObject
}

func (s *DetectWorkwearAdvanceRequest) GetLabels() []*string {
	return s.Labels
}

func (s *DetectWorkwearAdvanceRequest) SetClothes(v *DetectWorkwearAdvanceRequestClothes) *DetectWorkwearAdvanceRequest {
	s.Clothes = v
	return s
}

func (s *DetectWorkwearAdvanceRequest) SetImageUrlObject(v io.Reader) *DetectWorkwearAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

func (s *DetectWorkwearAdvanceRequest) SetLabels(v []*string) *DetectWorkwearAdvanceRequest {
	s.Labels = v
	return s
}

func (s *DetectWorkwearAdvanceRequest) Validate() error {
	return dara.Validate(s)
}

type DetectWorkwearAdvanceRequestClothes struct {
	// example:
	//
	// 1
	MaxNum *int64 `json:"MaxNum,omitempty" xml:"MaxNum,omitempty"`
	// example:
	//
	// 0.4
	Threshold *float64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DetectWorkwearAdvanceRequestClothes) String() string {
	return dara.Prettify(s)
}

func (s DetectWorkwearAdvanceRequestClothes) GoString() string {
	return s.String()
}

func (s *DetectWorkwearAdvanceRequestClothes) GetMaxNum() *int64 {
	return s.MaxNum
}

func (s *DetectWorkwearAdvanceRequestClothes) GetThreshold() *float64 {
	return s.Threshold
}

func (s *DetectWorkwearAdvanceRequestClothes) SetMaxNum(v int64) *DetectWorkwearAdvanceRequestClothes {
	s.MaxNum = &v
	return s
}

func (s *DetectWorkwearAdvanceRequestClothes) SetThreshold(v float64) *DetectWorkwearAdvanceRequestClothes {
	s.Threshold = &v
	return s
}

func (s *DetectWorkwearAdvanceRequestClothes) Validate() error {
	return dara.Validate(s)
}
