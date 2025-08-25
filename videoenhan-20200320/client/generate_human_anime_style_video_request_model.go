// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateHumanAnimeStyleVideoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCartoonStyle(v string) *GenerateHumanAnimeStyleVideoRequest
	GetCartoonStyle() *string
	SetVideoUrl(v string) *GenerateHumanAnimeStyleVideoRequest
	GetVideoUrl() *string
}

type GenerateHumanAnimeStyleVideoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// anime
	CartoonStyle *string `json:"CartoonStyle,omitempty" xml:"CartoonStyle,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/test/xxx/eas/EvaluateVideoQuality/123.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s GenerateHumanAnimeStyleVideoRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateHumanAnimeStyleVideoRequest) GoString() string {
	return s.String()
}

func (s *GenerateHumanAnimeStyleVideoRequest) GetCartoonStyle() *string {
	return s.CartoonStyle
}

func (s *GenerateHumanAnimeStyleVideoRequest) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *GenerateHumanAnimeStyleVideoRequest) SetCartoonStyle(v string) *GenerateHumanAnimeStyleVideoRequest {
	s.CartoonStyle = &v
	return s
}

func (s *GenerateHumanAnimeStyleVideoRequest) SetVideoUrl(v string) *GenerateHumanAnimeStyleVideoRequest {
	s.VideoUrl = &v
	return s
}

func (s *GenerateHumanAnimeStyleVideoRequest) Validate() error {
	return dara.Validate(s)
}
