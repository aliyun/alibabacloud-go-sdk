// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iToneSdrVideoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBitrate(v int32) *ToneSdrVideoRequest
	GetBitrate() *int32
	SetRecolorModel(v string) *ToneSdrVideoRequest
	GetRecolorModel() *string
	SetVideoURL(v string) *ToneSdrVideoRequest
	GetVideoURL() *string
}

type ToneSdrVideoRequest struct {
	// example:
	//
	// 30
	Bitrate *int32 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// example:
	//
	// auto_l2
	RecolorModel *string `json:"RecolorModel,omitempty" xml:"RecolorModel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://invi-label.oss-cn-shanghai.aliyuncs.com/label/temp/faceswap/test_for_api/xxxx.mp4
	VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s ToneSdrVideoRequest) String() string {
	return dara.Prettify(s)
}

func (s ToneSdrVideoRequest) GoString() string {
	return s.String()
}

func (s *ToneSdrVideoRequest) GetBitrate() *int32 {
	return s.Bitrate
}

func (s *ToneSdrVideoRequest) GetRecolorModel() *string {
	return s.RecolorModel
}

func (s *ToneSdrVideoRequest) GetVideoURL() *string {
	return s.VideoURL
}

func (s *ToneSdrVideoRequest) SetBitrate(v int32) *ToneSdrVideoRequest {
	s.Bitrate = &v
	return s
}

func (s *ToneSdrVideoRequest) SetRecolorModel(v string) *ToneSdrVideoRequest {
	s.RecolorModel = &v
	return s
}

func (s *ToneSdrVideoRequest) SetVideoURL(v string) *ToneSdrVideoRequest {
	s.VideoURL = &v
	return s
}

func (s *ToneSdrVideoRequest) Validate() error {
	return dara.Validate(s)
}
