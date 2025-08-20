// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iSegmentCommodityAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *SegmentCommodityAdvanceRequest
	GetImageURLObject() io.Reader
	SetReturnForm(v string) *SegmentCommodityAdvanceRequest
	GetReturnForm() *string
}

type SegmentCommodityAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentCommodity/SegmentCommodity1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// mask
	ReturnForm *string `json:"ReturnForm,omitempty" xml:"ReturnForm,omitempty"`
}

func (s SegmentCommodityAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SegmentCommodityAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentCommodityAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *SegmentCommodityAdvanceRequest) GetReturnForm() *string {
	return s.ReturnForm
}

func (s *SegmentCommodityAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentCommodityAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *SegmentCommodityAdvanceRequest) SetReturnForm(v string) *SegmentCommodityAdvanceRequest {
	s.ReturnForm = &v
	return s
}

func (s *SegmentCommodityAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
