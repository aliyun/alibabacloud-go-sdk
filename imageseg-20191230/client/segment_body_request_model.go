// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentBodyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *SegmentBodyRequest
	GetImageURL() *string
	SetReturnForm(v string) *SegmentBodyRequest
	GetReturnForm() *string
}

type SegmentBodyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentBody/SegmentBody1.png
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// mask
	ReturnForm *string `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentBodyRequest) String() string {
	return dara.Prettify(s)
}

func (s SegmentBodyRequest) GoString() string {
	return s.String()
}

func (s *SegmentBodyRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *SegmentBodyRequest) GetReturnForm() *string {
	return s.ReturnForm
}

func (s *SegmentBodyRequest) SetImageURL(v string) *SegmentBodyRequest {
	s.ImageURL = &v
	return s
}

func (s *SegmentBodyRequest) SetReturnForm(v string) *SegmentBodyRequest {
	s.ReturnForm = &v
	return s
}

func (s *SegmentBodyRequest) Validate() error {
	return dara.Validate(s)
}
