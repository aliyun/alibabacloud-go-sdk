// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssessCompositionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *AssessCompositionRequest
	GetImageURL() *string
}

type AssessCompositionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/AssessComposition/AssessComposition1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s AssessCompositionRequest) String() string {
	return dara.Prettify(s)
}

func (s AssessCompositionRequest) GoString() string {
	return s.String()
}

func (s *AssessCompositionRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *AssessCompositionRequest) SetImageURL(v string) *AssessCompositionRequest {
	s.ImageURL = &v
	return s
}

func (s *AssessCompositionRequest) Validate() error {
	return dara.Validate(s)
}
