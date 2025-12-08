// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iDetectVideoLivingFaceAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVideoUrlObject(v io.Reader) *DetectVideoLivingFaceAdvanceRequest
	GetVideoUrlObject() io.Reader
}

type DetectVideoLivingFaceAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/DetectVideoLivingFace/DetectVideoLivingFace1.mp4
	VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s DetectVideoLivingFaceAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectVideoLivingFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectVideoLivingFaceAdvanceRequest) GetVideoUrlObject() io.Reader {
	return s.VideoUrlObject
}

func (s *DetectVideoLivingFaceAdvanceRequest) SetVideoUrlObject(v io.Reader) *DetectVideoLivingFaceAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

func (s *DetectVideoLivingFaceAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
