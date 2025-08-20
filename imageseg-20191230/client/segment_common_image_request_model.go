// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentCommonImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *SegmentCommonImageRequest
	GetImageURL() *string
	SetReturnForm(v string) *SegmentCommonImageRequest
	GetReturnForm() *string
}

type SegmentCommonImageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentCommonImage/SegmentCommonImage1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// mask
	ReturnForm *string `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentCommonImageRequest) String() string {
	return dara.Prettify(s)
}

func (s SegmentCommonImageRequest) GoString() string {
	return s.String()
}

func (s *SegmentCommonImageRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *SegmentCommonImageRequest) GetReturnForm() *string {
	return s.ReturnForm
}

func (s *SegmentCommonImageRequest) SetImageURL(v string) *SegmentCommonImageRequest {
	s.ImageURL = &v
	return s
}

func (s *SegmentCommonImageRequest) SetReturnForm(v string) *SegmentCommonImageRequest {
	s.ReturnForm = &v
	return s
}

func (s *SegmentCommonImageRequest) Validate() error {
	return dara.Validate(s)
}
