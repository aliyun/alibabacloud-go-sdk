// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iSegmentHDBodyAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *SegmentHDBodyAdvanceRequest
	GetImageURLObject() io.Reader
}

type SegmentHDBodyAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/SegmentHDBody/SegmentHDBody1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentHDBodyAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SegmentHDBodyAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentHDBodyAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *SegmentHDBodyAdvanceRequest) SetImageURLObject(v io.Reader) *SegmentHDBodyAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *SegmentHDBodyAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
