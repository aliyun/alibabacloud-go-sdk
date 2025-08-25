// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssessSharpnessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *AssessSharpnessRequest
	GetImageURL() *string
}

type AssessSharpnessRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/AssessSharpness/AssessSharpness1.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s AssessSharpnessRequest) String() string {
	return dara.Prettify(s)
}

func (s AssessSharpnessRequest) GoString() string {
	return s.String()
}

func (s *AssessSharpnessRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *AssessSharpnessRequest) SetImageURL(v string) *AssessSharpnessRequest {
	s.ImageURL = &v
	return s
}

func (s *AssessSharpnessRequest) Validate() error {
	return dara.Validate(s)
}
