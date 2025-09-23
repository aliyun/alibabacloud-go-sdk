// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentVideoBodyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVideoUrl(v string) *SegmentVideoBodyRequest
	GetVideoUrl() *string
}

type SegmentVideoBodyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoseg/SegmentVideoBody/SegmentVideoBody1.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s SegmentVideoBodyRequest) String() string {
	return dara.Prettify(s)
}

func (s SegmentVideoBodyRequest) GoString() string {
	return s.String()
}

func (s *SegmentVideoBodyRequest) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *SegmentVideoBodyRequest) SetVideoUrl(v string) *SegmentVideoBodyRequest {
	s.VideoUrl = &v
	return s
}

func (s *SegmentVideoBodyRequest) Validate() error {
	return dara.Validate(s)
}
