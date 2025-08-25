// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iAddFaceVideoTemplateAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVideoScene(v string) *AddFaceVideoTemplateAdvanceRequest
	GetVideoScene() *string
	SetVideoURLObject(v io.Reader) *AddFaceVideoTemplateAdvanceRequest
	GetVideoURLObject() io.Reader
}

type AddFaceVideoTemplateAdvanceRequest struct {
	VideoScene *string `json:"VideoScene,omitempty" xml:"VideoScene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://invi-label.oss-cn-shanghai.aliyuncs.com/labl/temp/faceswap/test_for_api/xxxx.mp4
	VideoURLObject io.Reader `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s AddFaceVideoTemplateAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFaceVideoTemplateAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AddFaceVideoTemplateAdvanceRequest) GetVideoScene() *string {
	return s.VideoScene
}

func (s *AddFaceVideoTemplateAdvanceRequest) GetVideoURLObject() io.Reader {
	return s.VideoURLObject
}

func (s *AddFaceVideoTemplateAdvanceRequest) SetVideoScene(v string) *AddFaceVideoTemplateAdvanceRequest {
	s.VideoScene = &v
	return s
}

func (s *AddFaceVideoTemplateAdvanceRequest) SetVideoURLObject(v io.Reader) *AddFaceVideoTemplateAdvanceRequest {
	s.VideoURLObject = v
	return s
}

func (s *AddFaceVideoTemplateAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
