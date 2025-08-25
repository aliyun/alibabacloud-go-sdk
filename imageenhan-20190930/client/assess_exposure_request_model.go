// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssessExposureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *AssessExposureRequest
	GetImageURL() *string
}

type AssessExposureRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/AssessExposure/AssessExposure1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s AssessExposureRequest) String() string {
	return dara.Prettify(s)
}

func (s AssessExposureRequest) GoString() string {
	return s.String()
}

func (s *AssessExposureRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *AssessExposureRequest) SetImageURL(v string) *AssessExposureRequest {
	s.ImageURL = &v
	return s
}

func (s *AssessExposureRequest) Validate() error {
	return dara.Validate(s)
}
