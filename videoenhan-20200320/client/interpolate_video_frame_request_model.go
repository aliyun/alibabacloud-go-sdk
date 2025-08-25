// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInterpolateVideoFrameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBitrate(v int32) *InterpolateVideoFrameRequest
	GetBitrate() *int32
	SetFrameRate(v int32) *InterpolateVideoFrameRequest
	GetFrameRate() *int32
	SetVideoURL(v string) *InterpolateVideoFrameRequest
	GetVideoURL() *string
}

type InterpolateVideoFrameRequest struct {
	// example:
	//
	// 30
	Bitrate *int32 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// example:
	//
	// 70
	FrameRate *int32 `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/InterpolateVideoFrame/InterpolateVideoFrame3.mp4
	VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s InterpolateVideoFrameRequest) String() string {
	return dara.Prettify(s)
}

func (s InterpolateVideoFrameRequest) GoString() string {
	return s.String()
}

func (s *InterpolateVideoFrameRequest) GetBitrate() *int32 {
	return s.Bitrate
}

func (s *InterpolateVideoFrameRequest) GetFrameRate() *int32 {
	return s.FrameRate
}

func (s *InterpolateVideoFrameRequest) GetVideoURL() *string {
	return s.VideoURL
}

func (s *InterpolateVideoFrameRequest) SetBitrate(v int32) *InterpolateVideoFrameRequest {
	s.Bitrate = &v
	return s
}

func (s *InterpolateVideoFrameRequest) SetFrameRate(v int32) *InterpolateVideoFrameRequest {
	s.FrameRate = &v
	return s
}

func (s *InterpolateVideoFrameRequest) SetVideoURL(v string) *InterpolateVideoFrameRequest {
	s.VideoURL = &v
	return s
}

func (s *InterpolateVideoFrameRequest) Validate() error {
	return dara.Validate(s)
}
