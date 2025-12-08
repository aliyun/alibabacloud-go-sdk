// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iAddFaceImageTemplateAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *AddFaceImageTemplateAdvanceRequest
	GetImageURLObject() io.Reader
}

type AddFaceImageTemplateAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// https://invi-label.oss-cn-shanghai.aliyuncs.com/label/temp/faceswap/img_facefusion/template/sucai1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s AddFaceImageTemplateAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFaceImageTemplateAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AddFaceImageTemplateAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *AddFaceImageTemplateAdvanceRequest) SetImageURLObject(v io.Reader) *AddFaceImageTemplateAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *AddFaceImageTemplateAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
