// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iSegmentClothAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClothClass(v []*string) *SegmentClothAdvanceRequest
	GetClothClass() []*string
	SetImageURLObject(v io.Reader) *SegmentClothAdvanceRequest
	GetImageURLObject() io.Reader
	SetOutMode(v int64) *SegmentClothAdvanceRequest
	GetOutMode() *int64
	SetReturnForm(v string) *SegmentClothAdvanceRequest
	GetReturnForm() *string
}

type SegmentClothAdvanceRequest struct {
	ClothClass []*string `json:"ClothClass,omitempty" xml:"ClothClass,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentCloth/SegmentCloth1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	OutMode        *int64    `json:"OutMode,omitempty" xml:"OutMode,omitempty"`
	ReturnForm     *string   `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentClothAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SegmentClothAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentClothAdvanceRequest) GetClothClass() []*string {
	return s.ClothClass
}

func (s *SegmentClothAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *SegmentClothAdvanceRequest) GetOutMode() *int64 {
	return s.OutMode
}

func (s *SegmentClothAdvanceRequest) GetReturnForm() *string {
	return s.ReturnForm
}

func (s *SegmentClothAdvanceRequest) SetClothClass(v []*string) *SegmentClothAdvanceRequest {
	s.ClothClass = v
	return s
}

func (s *SegmentClothAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentClothAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *SegmentClothAdvanceRequest) SetOutMode(v int64) *SegmentClothAdvanceRequest {
	s.OutMode = &v
	return s
}

func (s *SegmentClothAdvanceRequest) SetReturnForm(v string) *SegmentClothAdvanceRequest {
	s.ReturnForm = &v
	return s
}

func (s *SegmentClothAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
