// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iSegmentHDSkyAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *SegmentHDSkyAdvanceRequest
	GetImageURLObject() io.Reader
}

type SegmentHDSkyAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentHDSky/SegmentHDSky4.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentHDSkyAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SegmentHDSkyAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentHDSkyAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *SegmentHDSkyAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentHDSkyAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *SegmentHDSkyAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
