// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentSkyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *SegmentSkyRequest
	GetImageURL() *string
}

type SegmentSkyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentSky/SegmentSky5.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentSkyRequest) String() string {
	return dara.Prettify(s)
}

func (s SegmentSkyRequest) GoString() string {
	return s.String()
}

func (s *SegmentSkyRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *SegmentSkyRequest) SetImageURL(v string) *SegmentSkyRequest {
	s.ImageURL = &v
	return s
}

func (s *SegmentSkyRequest) Validate() error {
	return dara.Validate(s)
}
