// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iGenerateVideoCoverAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsGif(v bool) *GenerateVideoCoverAdvanceRequest
	GetIsGif() *bool
	SetVideoUrlObject(v io.Reader) *GenerateVideoCoverAdvanceRequest
	GetVideoUrlObject() io.Reader
}

type GenerateVideoCoverAdvanceRequest struct {
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
	VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s GenerateVideoCoverAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateVideoCoverAdvanceRequest) GoString() string {
	return s.String()
}

func (s *GenerateVideoCoverAdvanceRequest) GetIsGif() *bool {
	return s.IsGif
}

func (s *GenerateVideoCoverAdvanceRequest) GetVideoUrlObject() io.Reader {
	return s.VideoUrlObject
}

func (s *GenerateVideoCoverAdvanceRequest) SetIsGif(v bool) *GenerateVideoCoverAdvanceRequest {
	s.IsGif = &v
	return s
}

func (s *GenerateVideoCoverAdvanceRequest) SetVideoUrlObject(v io.Reader) *GenerateVideoCoverAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

func (s *GenerateVideoCoverAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
