// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectWorkwearShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClothesShrink(v string) *DetectWorkwearShrinkRequest
	GetClothesShrink() *string
	SetImageUrl(v string) *DetectWorkwearShrinkRequest
	GetImageUrl() *string
	SetLabels(v []*string) *DetectWorkwearShrinkRequest
	GetLabels() []*string
}

type DetectWorkwearShrinkRequest struct {
	ClothesShrink *string `json:"Clothes,omitempty" xml:"Clothes,omitempty"`
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/test-team/zhangchaorun/tiyan/xxxx.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 1
	Labels []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s DetectWorkwearShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectWorkwearShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetectWorkwearShrinkRequest) GetClothesShrink() *string {
	return s.ClothesShrink
}

func (s *DetectWorkwearShrinkRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *DetectWorkwearShrinkRequest) GetLabels() []*string {
	return s.Labels
}

func (s *DetectWorkwearShrinkRequest) SetClothesShrink(v string) *DetectWorkwearShrinkRequest {
	s.ClothesShrink = &v
	return s
}

func (s *DetectWorkwearShrinkRequest) SetImageUrl(v string) *DetectWorkwearShrinkRequest {
	s.ImageUrl = &v
	return s
}

func (s *DetectWorkwearShrinkRequest) SetLabels(v []*string) *DetectWorkwearShrinkRequest {
	s.Labels = v
	return s
}

func (s *DetectWorkwearShrinkRequest) Validate() error {
	return dara.Validate(s)
}
