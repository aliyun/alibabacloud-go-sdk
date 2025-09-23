// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectWhiteBaseImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *DetectWhiteBaseImageRequest
	GetImageURL() *string
}

type DetectWhiteBaseImageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/objectdet/DetectWhiteBaseImage/DetectWhiteBaseImage1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectWhiteBaseImageRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectWhiteBaseImageRequest) GoString() string {
	return s.String()
}

func (s *DetectWhiteBaseImageRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *DetectWhiteBaseImageRequest) SetImageURL(v string) *DetectWhiteBaseImageRequest {
	s.ImageURL = &v
	return s
}

func (s *DetectWhiteBaseImageRequest) Validate() error {
	return dara.Validate(s)
}
