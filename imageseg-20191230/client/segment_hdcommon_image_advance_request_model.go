// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iSegmentHDCommonImageAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrlObject(v io.Reader) *SegmentHDCommonImageAdvanceRequest
	GetImageUrlObject() io.Reader
}

type SegmentHDCommonImageAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentHDCommonImage/SegmentHDCommonImage1.jpg
	ImageUrlObject io.Reader `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s SegmentHDCommonImageAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SegmentHDCommonImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentHDCommonImageAdvanceRequest) GetImageUrlObject() io.Reader {
	return s.ImageUrlObject
}

func (s *SegmentHDCommonImageAdvanceRequest) SetImageUrlObject(v io.Reader) *SegmentHDCommonImageAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

func (s *SegmentHDCommonImageAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
