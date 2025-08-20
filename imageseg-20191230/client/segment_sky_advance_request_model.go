// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iSegmentSkyAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *SegmentSkyAdvanceRequest
	GetImageURLObject() io.Reader
}

type SegmentSkyAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentSky/SegmentSky5.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentSkyAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SegmentSkyAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentSkyAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *SegmentSkyAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentSkyAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *SegmentSkyAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
