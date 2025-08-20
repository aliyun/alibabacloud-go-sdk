// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentHairRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *SegmentHairRequest
	GetImageURL() *string
}

type SegmentHairRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentHair/SegmentHair1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentHairRequest) String() string {
	return dara.Prettify(s)
}

func (s SegmentHairRequest) GoString() string {
	return s.String()
}

func (s *SegmentHairRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *SegmentHairRequest) SetImageURL(v string) *SegmentHairRequest {
	s.ImageURL = &v
	return s
}

func (s *SegmentHairRequest) Validate() error {
	return dara.Validate(s)
}
