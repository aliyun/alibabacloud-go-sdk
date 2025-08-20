// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentFoodRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *SegmentFoodRequest
	GetImageURL() *string
	SetReturnForm(v string) *SegmentFoodRequest
	GetReturnForm() *string
}

type SegmentFoodRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentFood/SegmentFood5.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// mask
	ReturnForm *string `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentFoodRequest) String() string {
	return dara.Prettify(s)
}

func (s SegmentFoodRequest) GoString() string {
	return s.String()
}

func (s *SegmentFoodRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *SegmentFoodRequest) GetReturnForm() *string {
	return s.ReturnForm
}

func (s *SegmentFoodRequest) SetImageURL(v string) *SegmentFoodRequest {
	s.ImageURL = &v
	return s
}

func (s *SegmentFoodRequest) SetReturnForm(v string) *SegmentFoodRequest {
	s.ReturnForm = &v
	return s
}

func (s *SegmentFoodRequest) Validate() error {
	return dara.Validate(s)
}
