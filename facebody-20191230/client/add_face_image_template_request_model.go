// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFaceImageTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *AddFaceImageTemplateRequest
	GetImageURL() *string
}

type AddFaceImageTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// https://invi-label.oss-cn-shanghai.aliyuncs.com/label/temp/faceswap/img_facefusion/template/sucai1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s AddFaceImageTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFaceImageTemplateRequest) GoString() string {
	return s.String()
}

func (s *AddFaceImageTemplateRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *AddFaceImageTemplateRequest) SetImageURL(v string) *AddFaceImageTemplateRequest {
	s.ImageURL = &v
	return s
}

func (s *AddFaceImageTemplateRequest) Validate() error {
	return dara.Validate(s)
}
