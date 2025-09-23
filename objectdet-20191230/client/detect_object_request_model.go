// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectObjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *DetectObjectRequest
	GetImageURL() *string
}

type DetectObjectRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/objectdet/DetectObject/DetectObject1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectObjectRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectObjectRequest) GoString() string {
	return s.String()
}

func (s *DetectObjectRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *DetectObjectRequest) SetImageURL(v string) *DetectObjectRequest {
	s.ImageURL = &v
	return s
}

func (s *DetectObjectRequest) Validate() error {
	return dara.Validate(s)
}
