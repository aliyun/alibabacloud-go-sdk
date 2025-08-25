// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iAdjustVideoColorAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMode(v string) *AdjustVideoColorAdvanceRequest
	GetMode() *string
	SetVideoBitrate(v int64) *AdjustVideoColorAdvanceRequest
	GetVideoBitrate() *int64
	SetVideoCodec(v string) *AdjustVideoColorAdvanceRequest
	GetVideoCodec() *string
	SetVideoFormat(v string) *AdjustVideoColorAdvanceRequest
	GetVideoFormat() *string
	SetVideoUrlObject(v io.Reader) *AdjustVideoColorAdvanceRequest
	GetVideoUrlObject() io.Reader
}

type AdjustVideoColorAdvanceRequest struct {
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
	VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s AdjustVideoColorAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s AdjustVideoColorAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AdjustVideoColorAdvanceRequest) GetMode() *string {
	return s.Mode
}

func (s *AdjustVideoColorAdvanceRequest) GetVideoBitrate() *int64 {
	return s.VideoBitrate
}

func (s *AdjustVideoColorAdvanceRequest) GetVideoCodec() *string {
	return s.VideoCodec
}

func (s *AdjustVideoColorAdvanceRequest) GetVideoFormat() *string {
	return s.VideoFormat
}

func (s *AdjustVideoColorAdvanceRequest) GetVideoUrlObject() io.Reader {
	return s.VideoUrlObject
}

func (s *AdjustVideoColorAdvanceRequest) SetMode(v string) *AdjustVideoColorAdvanceRequest {
	s.Mode = &v
	return s
}

func (s *AdjustVideoColorAdvanceRequest) SetVideoBitrate(v int64) *AdjustVideoColorAdvanceRequest {
	s.VideoBitrate = &v
	return s
}

func (s *AdjustVideoColorAdvanceRequest) SetVideoCodec(v string) *AdjustVideoColorAdvanceRequest {
	s.VideoCodec = &v
	return s
}

func (s *AdjustVideoColorAdvanceRequest) SetVideoFormat(v string) *AdjustVideoColorAdvanceRequest {
	s.VideoFormat = &v
	return s
}

func (s *AdjustVideoColorAdvanceRequest) SetVideoUrlObject(v io.Reader) *AdjustVideoColorAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

func (s *AdjustVideoColorAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
