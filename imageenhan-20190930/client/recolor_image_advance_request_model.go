// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iRecolorImageAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColorCount(v int32) *RecolorImageAdvanceRequest
	GetColorCount() *int32
	SetColorTemplate(v []*RecolorImageAdvanceRequestColorTemplate) *RecolorImageAdvanceRequest
	GetColorTemplate() []*RecolorImageAdvanceRequestColorTemplate
	SetMode(v string) *RecolorImageAdvanceRequest
	GetMode() *string
	SetRefUrlObject(v io.Reader) *RecolorImageAdvanceRequest
	GetRefUrlObject() io.Reader
	SetUrlObject(v io.Reader) *RecolorImageAdvanceRequest
	GetUrlObject() io.Reader
}

type RecolorImageAdvanceRequest struct {
	// example:
	//
	// 3
	ColorCount *int32 `json:"ColorCount,omitempty" xml:"ColorCount,omitempty"`
	// 1
	ColorTemplate []*RecolorImageAdvanceRequestColorTemplate `json:"ColorTemplate,omitempty" xml:"ColorTemplate,omitempty" type:"Repeated"`
	// example:
	//
	// TEMPLATE
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/RecolorImage/RecolorImage-REF_PIC7.jpg
	RefUrlObject io.Reader `json:"RefUrl,omitempty" xml:"RefUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/RecolorImage/RecolorImage-REF_PIC1.jpg
	UrlObject io.Reader `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RecolorImageAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecolorImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecolorImageAdvanceRequest) GetColorCount() *int32 {
	return s.ColorCount
}

func (s *RecolorImageAdvanceRequest) GetColorTemplate() []*RecolorImageAdvanceRequestColorTemplate {
	return s.ColorTemplate
}

func (s *RecolorImageAdvanceRequest) GetMode() *string {
	return s.Mode
}

func (s *RecolorImageAdvanceRequest) GetRefUrlObject() io.Reader {
	return s.RefUrlObject
}

func (s *RecolorImageAdvanceRequest) GetUrlObject() io.Reader {
	return s.UrlObject
}

func (s *RecolorImageAdvanceRequest) SetColorCount(v int32) *RecolorImageAdvanceRequest {
	s.ColorCount = &v
	return s
}

func (s *RecolorImageAdvanceRequest) SetColorTemplate(v []*RecolorImageAdvanceRequestColorTemplate) *RecolorImageAdvanceRequest {
	s.ColorTemplate = v
	return s
}

func (s *RecolorImageAdvanceRequest) SetMode(v string) *RecolorImageAdvanceRequest {
	s.Mode = &v
	return s
}

func (s *RecolorImageAdvanceRequest) SetRefUrlObject(v io.Reader) *RecolorImageAdvanceRequest {
	s.RefUrlObject = v
	return s
}

func (s *RecolorImageAdvanceRequest) SetUrlObject(v io.Reader) *RecolorImageAdvanceRequest {
	s.UrlObject = v
	return s
}

func (s *RecolorImageAdvanceRequest) Validate() error {
	if s.ColorTemplate != nil {
		for _, item := range s.ColorTemplate {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RecolorImageAdvanceRequestColorTemplate struct {
	// example:
	//
	// 056A6B
	Color *string `json:"Color,omitempty" xml:"Color,omitempty"`
}

func (s RecolorImageAdvanceRequestColorTemplate) String() string {
	return dara.Prettify(s)
}

func (s RecolorImageAdvanceRequestColorTemplate) GoString() string {
	return s.String()
}

func (s *RecolorImageAdvanceRequestColorTemplate) GetColor() *string {
	return s.Color
}

func (s *RecolorImageAdvanceRequestColorTemplate) SetColor(v string) *RecolorImageAdvanceRequestColorTemplate {
	s.Color = &v
	return s
}

func (s *RecolorImageAdvanceRequestColorTemplate) Validate() error {
	return dara.Validate(s)
}
