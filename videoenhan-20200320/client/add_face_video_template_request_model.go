// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFaceVideoTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVideoScene(v string) *AddFaceVideoTemplateRequest
	GetVideoScene() *string
	SetVideoURL(v string) *AddFaceVideoTemplateRequest
	GetVideoURL() *string
}

type AddFaceVideoTemplateRequest struct {
	VideoScene *string `json:"VideoScene,omitempty" xml:"VideoScene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://invi-label.oss-cn-shanghai.aliyuncs.com/labl/temp/faceswap/test_for_api/xxxx.mp4
	VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s AddFaceVideoTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFaceVideoTemplateRequest) GoString() string {
	return s.String()
}

func (s *AddFaceVideoTemplateRequest) GetVideoScene() *string {
	return s.VideoScene
}

func (s *AddFaceVideoTemplateRequest) GetVideoURL() *string {
	return s.VideoURL
}

func (s *AddFaceVideoTemplateRequest) SetVideoScene(v string) *AddFaceVideoTemplateRequest {
	s.VideoScene = &v
	return s
}

func (s *AddFaceVideoTemplateRequest) SetVideoURL(v string) *AddFaceVideoTemplateRequest {
	s.VideoURL = &v
	return s
}

func (s *AddFaceVideoTemplateRequest) Validate() error {
	return dara.Validate(s)
}
