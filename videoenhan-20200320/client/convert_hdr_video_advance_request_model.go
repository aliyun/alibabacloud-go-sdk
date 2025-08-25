// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iConvertHdrVideoAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBitrate(v int32) *ConvertHdrVideoAdvanceRequest
	GetBitrate() *int32
	SetHDRFormat(v string) *ConvertHdrVideoAdvanceRequest
	GetHDRFormat() *string
	SetMaxIlluminance(v int32) *ConvertHdrVideoAdvanceRequest
	GetMaxIlluminance() *int32
	SetVideoURLObject(v io.Reader) *ConvertHdrVideoAdvanceRequest
	GetVideoURLObject() io.Reader
}

type ConvertHdrVideoAdvanceRequest struct {
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
	VideoURLObject io.Reader `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s ConvertHdrVideoAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ConvertHdrVideoAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ConvertHdrVideoAdvanceRequest) GetBitrate() *int32 {
	return s.Bitrate
}

func (s *ConvertHdrVideoAdvanceRequest) GetHDRFormat() *string {
	return s.HDRFormat
}

func (s *ConvertHdrVideoAdvanceRequest) GetMaxIlluminance() *int32 {
	return s.MaxIlluminance
}

func (s *ConvertHdrVideoAdvanceRequest) GetVideoURLObject() io.Reader {
	return s.VideoURLObject
}

func (s *ConvertHdrVideoAdvanceRequest) SetBitrate(v int32) *ConvertHdrVideoAdvanceRequest {
	s.Bitrate = &v
	return s
}

func (s *ConvertHdrVideoAdvanceRequest) SetHDRFormat(v string) *ConvertHdrVideoAdvanceRequest {
	s.HDRFormat = &v
	return s
}

func (s *ConvertHdrVideoAdvanceRequest) SetMaxIlluminance(v int32) *ConvertHdrVideoAdvanceRequest {
	s.MaxIlluminance = &v
	return s
}

func (s *ConvertHdrVideoAdvanceRequest) SetVideoURLObject(v io.Reader) *ConvertHdrVideoAdvanceRequest {
	s.VideoURLObject = v
	return s
}

func (s *ConvertHdrVideoAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
