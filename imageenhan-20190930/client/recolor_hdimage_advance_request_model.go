// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iRecolorHDImageAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColorCount(v int32) *RecolorHDImageAdvanceRequest
	GetColorCount() *int32
	SetColorTemplate(v []*RecolorHDImageAdvanceRequestColorTemplate) *RecolorHDImageAdvanceRequest
	GetColorTemplate() []*RecolorHDImageAdvanceRequestColorTemplate
	SetDegree(v string) *RecolorHDImageAdvanceRequest
	GetDegree() *string
	SetMode(v string) *RecolorHDImageAdvanceRequest
	GetMode() *string
	SetRefUrlObject(v io.Reader) *RecolorHDImageAdvanceRequest
	GetRefUrlObject() io.Reader
	SetUrlObject(v io.Reader) *RecolorHDImageAdvanceRequest
	GetUrlObject() io.Reader
}

type RecolorHDImageAdvanceRequest struct {
	// example:
	//
	// 2
	ColorCount *int32 `json:"ColorCount,omitempty" xml:"ColorCount,omitempty"`
	// 1
	ColorTemplate []*RecolorHDImageAdvanceRequestColorTemplate `json:"ColorTemplate,omitempty" xml:"ColorTemplate,omitempty" type:"Repeated"`
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
	RefUrlObject io.Reader `json:"RefUrl,omitempty" xml:"RefUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/RecolorHDImage/RecolorHDImage-auto1.jpg
	UrlObject io.Reader `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RecolorHDImageAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecolorHDImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecolorHDImageAdvanceRequest) GetColorCount() *int32 {
	return s.ColorCount
}

func (s *RecolorHDImageAdvanceRequest) GetColorTemplate() []*RecolorHDImageAdvanceRequestColorTemplate {
	return s.ColorTemplate
}

func (s *RecolorHDImageAdvanceRequest) GetDegree() *string {
	return s.Degree
}

func (s *RecolorHDImageAdvanceRequest) GetMode() *string {
	return s.Mode
}

func (s *RecolorHDImageAdvanceRequest) GetRefUrlObject() io.Reader {
	return s.RefUrlObject
}

func (s *RecolorHDImageAdvanceRequest) GetUrlObject() io.Reader {
	return s.UrlObject
}

func (s *RecolorHDImageAdvanceRequest) SetColorCount(v int32) *RecolorHDImageAdvanceRequest {
	s.ColorCount = &v
	return s
}

func (s *RecolorHDImageAdvanceRequest) SetColorTemplate(v []*RecolorHDImageAdvanceRequestColorTemplate) *RecolorHDImageAdvanceRequest {
	s.ColorTemplate = v
	return s
}

func (s *RecolorHDImageAdvanceRequest) SetDegree(v string) *RecolorHDImageAdvanceRequest {
	s.Degree = &v
	return s
}

func (s *RecolorHDImageAdvanceRequest) SetMode(v string) *RecolorHDImageAdvanceRequest {
	s.Mode = &v
	return s
}

func (s *RecolorHDImageAdvanceRequest) SetRefUrlObject(v io.Reader) *RecolorHDImageAdvanceRequest {
	s.RefUrlObject = v
	return s
}

func (s *RecolorHDImageAdvanceRequest) SetUrlObject(v io.Reader) *RecolorHDImageAdvanceRequest {
	s.UrlObject = v
	return s
}

func (s *RecolorHDImageAdvanceRequest) Validate() error {
	return dara.Validate(s)
}

type RecolorHDImageAdvanceRequestColorTemplate struct {
	// example:
	//
	// [3F6A6B,0A0A6F]
	Color *string `json:"Color,omitempty" xml:"Color,omitempty"`
}

func (s RecolorHDImageAdvanceRequestColorTemplate) String() string {
	return dara.Prettify(s)
}

func (s RecolorHDImageAdvanceRequestColorTemplate) GoString() string {
	return s.String()
}

func (s *RecolorHDImageAdvanceRequestColorTemplate) GetColor() *string {
	return s.Color
}

func (s *RecolorHDImageAdvanceRequestColorTemplate) SetColor(v string) *RecolorHDImageAdvanceRequestColorTemplate {
	s.Color = &v
	return s
}

func (s *RecolorHDImageAdvanceRequestColorTemplate) Validate() error {
	return dara.Validate(s)
}
