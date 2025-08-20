// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentHeadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *SegmentHeadRequest
	GetImageURL() *string
	SetReturnForm(v string) *SegmentHeadRequest
	GetReturnForm() *string
}

type SegmentHeadRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentHead/SegmentHead1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// mask
	ReturnForm *string `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentHeadRequest) String() string {
	return dara.Prettify(s)
}

func (s SegmentHeadRequest) GoString() string {
	return s.String()
}

func (s *SegmentHeadRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *SegmentHeadRequest) GetReturnForm() *string {
	return s.ReturnForm
}

func (s *SegmentHeadRequest) SetImageURL(v string) *SegmentHeadRequest {
	s.ImageURL = &v
	return s
}

func (s *SegmentHeadRequest) SetReturnForm(v string) *SegmentHeadRequest {
	s.ReturnForm = &v
	return s
}

func (s *SegmentHeadRequest) Validate() error {
	return dara.Validate(s)
}
