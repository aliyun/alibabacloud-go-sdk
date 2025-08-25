// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertHdrVideoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBitrate(v int32) *ConvertHdrVideoRequest
	GetBitrate() *int32
	SetHDRFormat(v string) *ConvertHdrVideoRequest
	GetHDRFormat() *string
	SetMaxIlluminance(v int32) *ConvertHdrVideoRequest
	GetMaxIlluminance() *int32
	SetVideoURL(v string) *ConvertHdrVideoRequest
	GetVideoURL() *string
}

type ConvertHdrVideoRequest struct {
	// example:
	//
	// 30
	Bitrate *int32 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// example:
	//
	// HLG
	HDRFormat *string `json:"HDRFormat,omitempty" xml:"HDRFormat,omitempty"`
	// example:
	//
	// 1000
	MaxIlluminance *int32 `json:"MaxIlluminance,omitempty" xml:"MaxIlluminance,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://invi-label.oss-cn-shanghai.aliyuncs.com/label/temp/faceswap/test_for_api/xxxx.mp4
	VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s ConvertHdrVideoRequest) String() string {
	return dara.Prettify(s)
}

func (s ConvertHdrVideoRequest) GoString() string {
	return s.String()
}

func (s *ConvertHdrVideoRequest) GetBitrate() *int32 {
	return s.Bitrate
}

func (s *ConvertHdrVideoRequest) GetHDRFormat() *string {
	return s.HDRFormat
}

func (s *ConvertHdrVideoRequest) GetMaxIlluminance() *int32 {
	return s.MaxIlluminance
}

func (s *ConvertHdrVideoRequest) GetVideoURL() *string {
	return s.VideoURL
}

func (s *ConvertHdrVideoRequest) SetBitrate(v int32) *ConvertHdrVideoRequest {
	s.Bitrate = &v
	return s
}

func (s *ConvertHdrVideoRequest) SetHDRFormat(v string) *ConvertHdrVideoRequest {
	s.HDRFormat = &v
	return s
}

func (s *ConvertHdrVideoRequest) SetMaxIlluminance(v int32) *ConvertHdrVideoRequest {
	s.MaxIlluminance = &v
	return s
}

func (s *ConvertHdrVideoRequest) SetVideoURL(v string) *ConvertHdrVideoRequest {
	s.VideoURL = &v
	return s
}

func (s *ConvertHdrVideoRequest) Validate() error {
	return dara.Validate(s)
}
