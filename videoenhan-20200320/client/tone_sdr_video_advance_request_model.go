// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iToneSdrVideoAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBitrate(v int32) *ToneSdrVideoAdvanceRequest
	GetBitrate() *int32
	SetRecolorModel(v string) *ToneSdrVideoAdvanceRequest
	GetRecolorModel() *string
	SetVideoURLObject(v io.Reader) *ToneSdrVideoAdvanceRequest
	GetVideoURLObject() io.Reader
}

type ToneSdrVideoAdvanceRequest struct {
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
	VideoURLObject io.Reader `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s ToneSdrVideoAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ToneSdrVideoAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ToneSdrVideoAdvanceRequest) GetBitrate() *int32 {
	return s.Bitrate
}

func (s *ToneSdrVideoAdvanceRequest) GetRecolorModel() *string {
	return s.RecolorModel
}

func (s *ToneSdrVideoAdvanceRequest) GetVideoURLObject() io.Reader {
	return s.VideoURLObject
}

func (s *ToneSdrVideoAdvanceRequest) SetBitrate(v int32) *ToneSdrVideoAdvanceRequest {
	s.Bitrate = &v
	return s
}

func (s *ToneSdrVideoAdvanceRequest) SetRecolorModel(v string) *ToneSdrVideoAdvanceRequest {
	s.RecolorModel = &v
	return s
}

func (s *ToneSdrVideoAdvanceRequest) SetVideoURLObject(v io.Reader) *ToneSdrVideoAdvanceRequest {
	s.VideoURLObject = v
	return s
}

func (s *ToneSdrVideoAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
