// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentCommodityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *SegmentCommodityRequest
	GetImageURL() *string
	SetReturnForm(v string) *SegmentCommodityRequest
	GetReturnForm() *string
}

type SegmentCommodityRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentCommodity/SegmentCommodity1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// mask
	ReturnForm *string `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentCommodityRequest) String() string {
	return dara.Prettify(s)
}

func (s SegmentCommodityRequest) GoString() string {
	return s.String()
}

func (s *SegmentCommodityRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *SegmentCommodityRequest) GetReturnForm() *string {
	return s.ReturnForm
}

func (s *SegmentCommodityRequest) SetImageURL(v string) *SegmentCommodityRequest {
	s.ImageURL = &v
	return s
}

func (s *SegmentCommodityRequest) SetReturnForm(v string) *SegmentCommodityRequest {
	s.ReturnForm = &v
	return s
}

func (s *SegmentCommodityRequest) Validate() error {
	return dara.Validate(s)
}
