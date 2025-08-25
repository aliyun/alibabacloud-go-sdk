// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iAssessSharpnessAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *AssessSharpnessAdvanceRequest
	GetImageURLObject() io.Reader
}

type AssessSharpnessAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/AssessSharpness/AssessSharpness1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s AssessSharpnessAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s AssessSharpnessAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AssessSharpnessAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *AssessSharpnessAdvanceRequest) SetImageURLObject(v io.Reader) *AssessSharpnessAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *AssessSharpnessAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
