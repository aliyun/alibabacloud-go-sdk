// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentHDCommonImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v string) *SegmentHDCommonImageRequest
	GetImageUrl() *string
}

type SegmentHDCommonImageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentHDCommonImage/SegmentHDCommonImage1.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s SegmentHDCommonImageRequest) String() string {
	return dara.Prettify(s)
}

func (s SegmentHDCommonImageRequest) GoString() string {
	return s.String()
}

func (s *SegmentHDCommonImageRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *SegmentHDCommonImageRequest) SetImageUrl(v string) *SegmentHDCommonImageRequest {
	s.ImageUrl = &v
	return s
}

func (s *SegmentHDCommonImageRequest) Validate() error {
	return dara.Validate(s)
}
