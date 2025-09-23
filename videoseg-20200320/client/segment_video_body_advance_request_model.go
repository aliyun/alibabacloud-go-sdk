// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iSegmentVideoBodyAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVideoUrlObject(v io.Reader) *SegmentVideoBodyAdvanceRequest
	GetVideoUrlObject() io.Reader
}

type SegmentVideoBodyAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoseg/SegmentVideoBody/SegmentVideoBody1.mp4
	VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s SegmentVideoBodyAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SegmentVideoBodyAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentVideoBodyAdvanceRequest) GetVideoUrlObject() io.Reader {
	return s.VideoUrlObject
}

func (s *SegmentVideoBodyAdvanceRequest) SetVideoUrlObject(v io.Reader) *SegmentVideoBodyAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

func (s *SegmentVideoBodyAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
