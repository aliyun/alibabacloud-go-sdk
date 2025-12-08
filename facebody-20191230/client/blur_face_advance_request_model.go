// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iBlurFaceAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *BlurFaceAdvanceRequest
	GetImageURLObject() io.Reader
}

type BlurFaceAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/BlurFace/BlurFace1.png
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s BlurFaceAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s BlurFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *BlurFaceAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *BlurFaceAdvanceRequest) SetImageURLObject(v io.Reader) *BlurFaceAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *BlurFaceAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
