// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iGenerateHumanAnimeStyleVideoAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCartoonStyle(v string) *GenerateHumanAnimeStyleVideoAdvanceRequest
	GetCartoonStyle() *string
	SetVideoUrlObject(v io.Reader) *GenerateHumanAnimeStyleVideoAdvanceRequest
	GetVideoUrlObject() io.Reader
}

type GenerateHumanAnimeStyleVideoAdvanceRequest struct {
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
	VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s GenerateHumanAnimeStyleVideoAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateHumanAnimeStyleVideoAdvanceRequest) GoString() string {
	return s.String()
}

func (s *GenerateHumanAnimeStyleVideoAdvanceRequest) GetCartoonStyle() *string {
	return s.CartoonStyle
}

func (s *GenerateHumanAnimeStyleVideoAdvanceRequest) GetVideoUrlObject() io.Reader {
	return s.VideoUrlObject
}

func (s *GenerateHumanAnimeStyleVideoAdvanceRequest) SetCartoonStyle(v string) *GenerateHumanAnimeStyleVideoAdvanceRequest {
	s.CartoonStyle = &v
	return s
}

func (s *GenerateHumanAnimeStyleVideoAdvanceRequest) SetVideoUrlObject(v io.Reader) *GenerateHumanAnimeStyleVideoAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

func (s *GenerateHumanAnimeStyleVideoAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
