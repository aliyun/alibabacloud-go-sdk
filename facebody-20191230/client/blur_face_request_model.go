// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBlurFaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *BlurFaceRequest
	GetImageURL() *string
}

type BlurFaceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/BlurFace/BlurFace1.png
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s BlurFaceRequest) String() string {
	return dara.Prettify(s)
}

func (s BlurFaceRequest) GoString() string {
	return s.String()
}

func (s *BlurFaceRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *BlurFaceRequest) SetImageURL(v string) *BlurFaceRequest {
	s.ImageURL = &v
	return s
}

func (s *BlurFaceRequest) Validate() error {
	return dara.Validate(s)
}
