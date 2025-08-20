// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentClothRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClothClass(v []*string) *SegmentClothRequest
	GetClothClass() []*string
	SetImageURL(v string) *SegmentClothRequest
	GetImageURL() *string
	SetOutMode(v int64) *SegmentClothRequest
	GetOutMode() *int64
	SetReturnForm(v string) *SegmentClothRequest
	GetReturnForm() *string
}

type SegmentClothRequest struct {
	ClothClass []*string `json:"ClothClass,omitempty" xml:"ClothClass,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentCloth/SegmentCloth1.jpg
	ImageURL   *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	OutMode    *int64  `json:"OutMode,omitempty" xml:"OutMode,omitempty"`
	ReturnForm *string `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentClothRequest) String() string {
	return dara.Prettify(s)
}

func (s SegmentClothRequest) GoString() string {
	return s.String()
}

func (s *SegmentClothRequest) GetClothClass() []*string {
	return s.ClothClass
}

func (s *SegmentClothRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *SegmentClothRequest) GetOutMode() *int64 {
	return s.OutMode
}

func (s *SegmentClothRequest) GetReturnForm() *string {
	return s.ReturnForm
}

func (s *SegmentClothRequest) SetClothClass(v []*string) *SegmentClothRequest {
	s.ClothClass = v
	return s
}

func (s *SegmentClothRequest) SetImageURL(v string) *SegmentClothRequest {
	s.ImageURL = &v
	return s
}

func (s *SegmentClothRequest) SetOutMode(v int64) *SegmentClothRequest {
	s.OutMode = &v
	return s
}

func (s *SegmentClothRequest) SetReturnForm(v string) *SegmentClothRequest {
	s.ReturnForm = &v
	return s
}

func (s *SegmentClothRequest) Validate() error {
	return dara.Validate(s)
}
