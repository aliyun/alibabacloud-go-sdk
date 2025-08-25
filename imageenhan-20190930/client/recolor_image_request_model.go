// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecolorImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColorCount(v int32) *RecolorImageRequest
	GetColorCount() *int32
	SetColorTemplate(v []*RecolorImageRequestColorTemplate) *RecolorImageRequest
	GetColorTemplate() []*RecolorImageRequestColorTemplate
	SetMode(v string) *RecolorImageRequest
	GetMode() *string
	SetRefUrl(v string) *RecolorImageRequest
	GetRefUrl() *string
	SetUrl(v string) *RecolorImageRequest
	GetUrl() *string
}

type RecolorImageRequest struct {
	// example:
	//
	// 3
	ColorCount *int32 `json:"ColorCount,omitempty" xml:"ColorCount,omitempty"`
	// 1
	ColorTemplate []*RecolorImageRequestColorTemplate `json:"ColorTemplate,omitempty" xml:"ColorTemplate,omitempty" type:"Repeated"`
	// example:
	//
	// TEMPLATE
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/RecolorImage/RecolorImage-REF_PIC7.jpg
	RefUrl *string `json:"RefUrl,omitempty" xml:"RefUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/RecolorImage/RecolorImage-REF_PIC1.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RecolorImageRequest) String() string {
	return dara.Prettify(s)
}

func (s RecolorImageRequest) GoString() string {
	return s.String()
}

func (s *RecolorImageRequest) GetColorCount() *int32 {
	return s.ColorCount
}

func (s *RecolorImageRequest) GetColorTemplate() []*RecolorImageRequestColorTemplate {
	return s.ColorTemplate
}

func (s *RecolorImageRequest) GetMode() *string {
	return s.Mode
}

func (s *RecolorImageRequest) GetRefUrl() *string {
	return s.RefUrl
}

func (s *RecolorImageRequest) GetUrl() *string {
	return s.Url
}

func (s *RecolorImageRequest) SetColorCount(v int32) *RecolorImageRequest {
	s.ColorCount = &v
	return s
}

func (s *RecolorImageRequest) SetColorTemplate(v []*RecolorImageRequestColorTemplate) *RecolorImageRequest {
	s.ColorTemplate = v
	return s
}

func (s *RecolorImageRequest) SetMode(v string) *RecolorImageRequest {
	s.Mode = &v
	return s
}

func (s *RecolorImageRequest) SetRefUrl(v string) *RecolorImageRequest {
	s.RefUrl = &v
	return s
}

func (s *RecolorImageRequest) SetUrl(v string) *RecolorImageRequest {
	s.Url = &v
	return s
}

func (s *RecolorImageRequest) Validate() error {
	return dara.Validate(s)
}

type RecolorImageRequestColorTemplate struct {
	// example:
	//
	// 056A6B
	Color *string `json:"Color,omitempty" xml:"Color,omitempty"`
}

func (s RecolorImageRequestColorTemplate) String() string {
	return dara.Prettify(s)
}

func (s RecolorImageRequestColorTemplate) GoString() string {
	return s.String()
}

func (s *RecolorImageRequestColorTemplate) GetColor() *string {
	return s.Color
}

func (s *RecolorImageRequestColorTemplate) SetColor(v string) *RecolorImageRequestColorTemplate {
	s.Color = &v
	return s
}

func (s *RecolorImageRequestColorTemplate) Validate() error {
	return dara.Validate(s)
}
