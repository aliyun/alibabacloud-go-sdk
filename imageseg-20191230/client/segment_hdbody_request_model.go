// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentHDBodyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *SegmentHDBodyRequest
	GetImageURL() *string
}

type SegmentHDBodyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentHDBody/SegmentHDBody1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentHDBodyRequest) String() string {
	return dara.Prettify(s)
}

func (s SegmentHDBodyRequest) GoString() string {
	return s.String()
}

func (s *SegmentHDBodyRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *SegmentHDBodyRequest) SetImageURL(v string) *SegmentHDBodyRequest {
	s.ImageURL = &v
	return s
}

func (s *SegmentHDBodyRequest) Validate() error {
	return dara.Validate(s)
}
