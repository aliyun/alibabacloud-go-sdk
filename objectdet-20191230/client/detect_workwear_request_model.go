// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectWorkwearRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClothes(v *DetectWorkwearRequestClothes) *DetectWorkwearRequest
	GetClothes() *DetectWorkwearRequestClothes
	SetImageUrl(v string) *DetectWorkwearRequest
	GetImageUrl() *string
	SetLabels(v []*string) *DetectWorkwearRequest
	GetLabels() []*string
}

type DetectWorkwearRequest struct {
	Clothes *DetectWorkwearRequestClothes `json:"Clothes,omitempty" xml:"Clothes,omitempty" type:"Struct"`
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/test-team/zhangchaorun/tiyan/xxxx.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 1
	Labels []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s DetectWorkwearRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectWorkwearRequest) GoString() string {
	return s.String()
}

func (s *DetectWorkwearRequest) GetClothes() *DetectWorkwearRequestClothes {
	return s.Clothes
}

func (s *DetectWorkwearRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *DetectWorkwearRequest) GetLabels() []*string {
	return s.Labels
}

func (s *DetectWorkwearRequest) SetClothes(v *DetectWorkwearRequestClothes) *DetectWorkwearRequest {
	s.Clothes = v
	return s
}

func (s *DetectWorkwearRequest) SetImageUrl(v string) *DetectWorkwearRequest {
	s.ImageUrl = &v
	return s
}

func (s *DetectWorkwearRequest) SetLabels(v []*string) *DetectWorkwearRequest {
	s.Labels = v
	return s
}

func (s *DetectWorkwearRequest) Validate() error {
	return dara.Validate(s)
}

type DetectWorkwearRequestClothes struct {
	// example:
	//
	// 1
	MaxNum *int64 `json:"MaxNum,omitempty" xml:"MaxNum,omitempty"`
	// example:
	//
	// 0.4
	Threshold *float64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DetectWorkwearRequestClothes) String() string {
	return dara.Prettify(s)
}

func (s DetectWorkwearRequestClothes) GoString() string {
	return s.String()
}

func (s *DetectWorkwearRequestClothes) GetMaxNum() *int64 {
	return s.MaxNum
}

func (s *DetectWorkwearRequestClothes) GetThreshold() *float64 {
	return s.Threshold
}

func (s *DetectWorkwearRequestClothes) SetMaxNum(v int64) *DetectWorkwearRequestClothes {
	s.MaxNum = &v
	return s
}

func (s *DetectWorkwearRequestClothes) SetThreshold(v float64) *DetectWorkwearRequestClothes {
	s.Threshold = &v
	return s
}

func (s *DetectWorkwearRequestClothes) Validate() error {
	return dara.Validate(s)
}
