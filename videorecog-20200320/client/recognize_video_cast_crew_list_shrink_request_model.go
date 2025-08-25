// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeVideoCastCrewListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParamsShrink(v string) *RecognizeVideoCastCrewListShrinkRequest
	GetParamsShrink() *string
	SetVideoUrl(v string) *RecognizeVideoCastCrewListShrinkRequest
	GetVideoUrl() *string
}

type RecognizeVideoCastCrewListShrinkRequest struct {
	ParamsShrink *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://shanghai.oss-cn-shanghai.aliyuncs.com/download/xxxx.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s RecognizeVideoCastCrewListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVideoCastCrewListShrinkRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListShrinkRequest) GetParamsShrink() *string {
	return s.ParamsShrink
}

func (s *RecognizeVideoCastCrewListShrinkRequest) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *RecognizeVideoCastCrewListShrinkRequest) SetParamsShrink(v string) *RecognizeVideoCastCrewListShrinkRequest {
	s.ParamsShrink = &v
	return s
}

func (s *RecognizeVideoCastCrewListShrinkRequest) SetVideoUrl(v string) *RecognizeVideoCastCrewListShrinkRequest {
	s.VideoUrl = &v
	return s
}

func (s *RecognizeVideoCastCrewListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
