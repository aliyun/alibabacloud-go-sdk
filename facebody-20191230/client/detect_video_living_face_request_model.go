// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectVideoLivingFaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVideoUrl(v string) *DetectVideoLivingFaceRequest
	GetVideoUrl() *string
}

type DetectVideoLivingFaceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/DetectVideoLivingFace/DetectVideoLivingFace1.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s DetectVideoLivingFaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectVideoLivingFaceRequest) GoString() string {
	return s.String()
}

func (s *DetectVideoLivingFaceRequest) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *DetectVideoLivingFaceRequest) SetVideoUrl(v string) *DetectVideoLivingFaceRequest {
	s.VideoUrl = &v
	return s
}

func (s *DetectVideoLivingFaceRequest) Validate() error {
	return dara.Validate(s)
}
