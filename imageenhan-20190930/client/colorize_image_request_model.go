// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iColorizeImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *ColorizeImageRequest
	GetImageURL() *string
}

type ColorizeImageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/ColorizeImage/ColorizeImage1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s ColorizeImageRequest) String() string {
	return dara.Prettify(s)
}

func (s ColorizeImageRequest) GoString() string {
	return s.String()
}

func (s *ColorizeImageRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *ColorizeImageRequest) SetImageURL(v string) *ColorizeImageRequest {
	s.ImageURL = &v
	return s
}

func (s *ColorizeImageRequest) Validate() error {
	return dara.Validate(s)
}
