// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iInterpolateVideoFrameAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBitrate(v int32) *InterpolateVideoFrameAdvanceRequest
	GetBitrate() *int32
	SetFrameRate(v int32) *InterpolateVideoFrameAdvanceRequest
	GetFrameRate() *int32
	SetVideoURLObject(v io.Reader) *InterpolateVideoFrameAdvanceRequest
	GetVideoURLObject() io.Reader
}

type InterpolateVideoFrameAdvanceRequest struct {
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
	VideoURLObject io.Reader `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s InterpolateVideoFrameAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s InterpolateVideoFrameAdvanceRequest) GoString() string {
	return s.String()
}

func (s *InterpolateVideoFrameAdvanceRequest) GetBitrate() *int32 {
	return s.Bitrate
}

func (s *InterpolateVideoFrameAdvanceRequest) GetFrameRate() *int32 {
	return s.FrameRate
}

func (s *InterpolateVideoFrameAdvanceRequest) GetVideoURLObject() io.Reader {
	return s.VideoURLObject
}

func (s *InterpolateVideoFrameAdvanceRequest) SetBitrate(v int32) *InterpolateVideoFrameAdvanceRequest {
	s.Bitrate = &v
	return s
}

func (s *InterpolateVideoFrameAdvanceRequest) SetFrameRate(v int32) *InterpolateVideoFrameAdvanceRequest {
	s.FrameRate = &v
	return s
}

func (s *InterpolateVideoFrameAdvanceRequest) SetVideoURLObject(v io.Reader) *InterpolateVideoFrameAdvanceRequest {
	s.VideoURLObject = v
	return s
}

func (s *InterpolateVideoFrameAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
