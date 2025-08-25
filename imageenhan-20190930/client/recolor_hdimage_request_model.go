// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecolorHDImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColorCount(v int32) *RecolorHDImageRequest
	GetColorCount() *int32
	SetColorTemplate(v []*RecolorHDImageRequestColorTemplate) *RecolorHDImageRequest
	GetColorTemplate() []*RecolorHDImageRequestColorTemplate
	SetDegree(v string) *RecolorHDImageRequest
	GetDegree() *string
	SetMode(v string) *RecolorHDImageRequest
	GetMode() *string
	SetRefUrl(v string) *RecolorHDImageRequest
	GetRefUrl() *string
	SetUrl(v string) *RecolorHDImageRequest
	GetUrl() *string
}

type RecolorHDImageRequest struct {
	// example:
	//
	// 2
	ColorCount *int32 `json:"ColorCount,omitempty" xml:"ColorCount,omitempty"`
	// 1
	ColorTemplate []*RecolorHDImageRequestColorTemplate `json:"ColorTemplate,omitempty" xml:"ColorTemplate,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 0.4
	Degree *string `json:"Degree,omitempty" xml:"Degree,omitempty"`
	// example:
	//
	// REF_PIC
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/RecolorHDImage/RecolorHDImage-REF_PIC6.jpg
	RefUrl *string `json:"RefUrl,omitempty" xml:"RefUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/RecolorHDImage/RecolorHDImage-auto1.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RecolorHDImageRequest) String() string {
	return dara.Prettify(s)
}

func (s RecolorHDImageRequest) GoString() string {
	return s.String()
}

func (s *RecolorHDImageRequest) GetColorCount() *int32 {
	return s.ColorCount
}

func (s *RecolorHDImageRequest) GetColorTemplate() []*RecolorHDImageRequestColorTemplate {
	return s.ColorTemplate
}

func (s *RecolorHDImageRequest) GetDegree() *string {
	return s.Degree
}

func (s *RecolorHDImageRequest) GetMode() *string {
	return s.Mode
}

func (s *RecolorHDImageRequest) GetRefUrl() *string {
	return s.RefUrl
}

func (s *RecolorHDImageRequest) GetUrl() *string {
	return s.Url
}

func (s *RecolorHDImageRequest) SetColorCount(v int32) *RecolorHDImageRequest {
	s.ColorCount = &v
	return s
}

func (s *RecolorHDImageRequest) SetColorTemplate(v []*RecolorHDImageRequestColorTemplate) *RecolorHDImageRequest {
	s.ColorTemplate = v
	return s
}

func (s *RecolorHDImageRequest) SetDegree(v string) *RecolorHDImageRequest {
	s.Degree = &v
	return s
}

func (s *RecolorHDImageRequest) SetMode(v string) *RecolorHDImageRequest {
	s.Mode = &v
	return s
}

func (s *RecolorHDImageRequest) SetRefUrl(v string) *RecolorHDImageRequest {
	s.RefUrl = &v
	return s
}

func (s *RecolorHDImageRequest) SetUrl(v string) *RecolorHDImageRequest {
	s.Url = &v
	return s
}

func (s *RecolorHDImageRequest) Validate() error {
	return dara.Validate(s)
}

type RecolorHDImageRequestColorTemplate struct {
	// example:
	//
	// [3F6A6B,0A0A6F]
	Color *string `json:"Color,omitempty" xml:"Color,omitempty"`
}

func (s RecolorHDImageRequestColorTemplate) String() string {
	return dara.Prettify(s)
}

func (s RecolorHDImageRequestColorTemplate) GoString() string {
	return s.String()
}

func (s *RecolorHDImageRequestColorTemplate) GetColor() *string {
	return s.Color
}

func (s *RecolorHDImageRequestColorTemplate) SetColor(v string) *RecolorHDImageRequestColorTemplate {
	s.Color = &v
	return s
}

func (s *RecolorHDImageRequestColorTemplate) Validate() error {
	return dara.Validate(s)
}
