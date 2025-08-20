// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentHDSkyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *SegmentHDSkyRequest
	GetImageURL() *string
}

type SegmentHDSkyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentHDSky/SegmentHDSky4.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentHDSkyRequest) String() string {
	return dara.Prettify(s)
}

func (s SegmentHDSkyRequest) GoString() string {
	return s.String()
}

func (s *SegmentHDSkyRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *SegmentHDSkyRequest) SetImageURL(v string) *SegmentHDSkyRequest {
	s.ImageURL = &v
	return s
}

func (s *SegmentHDSkyRequest) Validate() error {
	return dara.Validate(s)
}
