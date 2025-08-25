// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAdjustVideoColorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMode(v string) *AdjustVideoColorRequest
	GetMode() *string
	SetVideoBitrate(v int64) *AdjustVideoColorRequest
	GetVideoBitrate() *int64
	SetVideoCodec(v string) *AdjustVideoColorRequest
	GetVideoCodec() *string
	SetVideoFormat(v string) *AdjustVideoColorRequest
	GetVideoFormat() *string
	SetVideoUrl(v string) *AdjustVideoColorRequest
	GetVideoUrl() *string
}

type AdjustVideoColorRequest struct {
	// example:
	//
	// LogC
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// 20
	VideoBitrate *int64 `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	// example:
	//
	// h264
	VideoCodec *string `json:"VideoCodec,omitempty" xml:"VideoCodec,omitempty"`
	// example:
	//
	// mp4
	VideoFormat *string `json:"VideoFormat,omitempty" xml:"VideoFormat,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/AdjustVideoColor/AdjustVideoColor1.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s AdjustVideoColorRequest) String() string {
	return dara.Prettify(s)
}

func (s AdjustVideoColorRequest) GoString() string {
	return s.String()
}

func (s *AdjustVideoColorRequest) GetMode() *string {
	return s.Mode
}

func (s *AdjustVideoColorRequest) GetVideoBitrate() *int64 {
	return s.VideoBitrate
}

func (s *AdjustVideoColorRequest) GetVideoCodec() *string {
	return s.VideoCodec
}

func (s *AdjustVideoColorRequest) GetVideoFormat() *string {
	return s.VideoFormat
}

func (s *AdjustVideoColorRequest) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *AdjustVideoColorRequest) SetMode(v string) *AdjustVideoColorRequest {
	s.Mode = &v
	return s
}

func (s *AdjustVideoColorRequest) SetVideoBitrate(v int64) *AdjustVideoColorRequest {
	s.VideoBitrate = &v
	return s
}

func (s *AdjustVideoColorRequest) SetVideoCodec(v string) *AdjustVideoColorRequest {
	s.VideoCodec = &v
	return s
}

func (s *AdjustVideoColorRequest) SetVideoFormat(v string) *AdjustVideoColorRequest {
	s.VideoFormat = &v
	return s
}

func (s *AdjustVideoColorRequest) SetVideoUrl(v string) *AdjustVideoColorRequest {
	s.VideoUrl = &v
	return s
}

func (s *AdjustVideoColorRequest) Validate() error {
	return dara.Validate(s)
}
