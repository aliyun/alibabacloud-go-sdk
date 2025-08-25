// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateVideoCoverRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsGif(v bool) *GenerateVideoCoverRequest
	GetIsGif() *bool
	SetVideoUrl(v string) *GenerateVideoCoverRequest
	GetVideoUrl() *string
}

type GenerateVideoCoverRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// false
	IsGif *bool `json:"IsGif,omitempty" xml:"IsGif,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videorecog/videorecog/videorecog1.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s GenerateVideoCoverRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateVideoCoverRequest) GoString() string {
	return s.String()
}

func (s *GenerateVideoCoverRequest) GetIsGif() *bool {
	return s.IsGif
}

func (s *GenerateVideoCoverRequest) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *GenerateVideoCoverRequest) SetIsGif(v bool) *GenerateVideoCoverRequest {
	s.IsGif = &v
	return s
}

func (s *GenerateVideoCoverRequest) SetVideoUrl(v string) *GenerateVideoCoverRequest {
	s.VideoUrl = &v
	return s
}

func (s *GenerateVideoCoverRequest) Validate() error {
	return dara.Validate(s)
}
