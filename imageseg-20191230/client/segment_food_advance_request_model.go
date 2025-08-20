// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iSegmentFoodAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *SegmentFoodAdvanceRequest
	GetImageURLObject() io.Reader
	SetReturnForm(v string) *SegmentFoodAdvanceRequest
	GetReturnForm() *string
}

type SegmentFoodAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentFood/SegmentFood5.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// mask
	ReturnForm *string `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentFoodAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SegmentFoodAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentFoodAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *SegmentFoodAdvanceRequest) GetReturnForm() *string {
	return s.ReturnForm
}

func (s *SegmentFoodAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentFoodAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *SegmentFoodAdvanceRequest) SetReturnForm(v string) *SegmentFoodAdvanceRequest {
	s.ReturnForm = &v
	return s
}

func (s *SegmentFoodAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
