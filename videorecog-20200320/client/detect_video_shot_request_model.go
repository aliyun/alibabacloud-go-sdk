// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectVideoShotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVideoUrl(v string) *DetectVideoShotRequest
	GetVideoUrl() *string
}

type DetectVideoShotRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videorecog/DetectVideoShot/DetectVideoShot2.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s DetectVideoShotRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectVideoShotRequest) GoString() string {
	return s.String()
}

func (s *DetectVideoShotRequest) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *DetectVideoShotRequest) SetVideoUrl(v string) *DetectVideoShotRequest {
	s.VideoUrl = &v
	return s
}

func (s *DetectVideoShotRequest) Validate() error {
	return dara.Validate(s)
}
