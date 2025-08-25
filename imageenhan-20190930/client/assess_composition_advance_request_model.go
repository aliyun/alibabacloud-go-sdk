// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iAssessCompositionAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *AssessCompositionAdvanceRequest
	GetImageURLObject() io.Reader
}

type AssessCompositionAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/AssessComposition/AssessComposition1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s AssessCompositionAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s AssessCompositionAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AssessCompositionAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *AssessCompositionAdvanceRequest) SetImageURLObject(v io.Reader) *AssessCompositionAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *AssessCompositionAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
