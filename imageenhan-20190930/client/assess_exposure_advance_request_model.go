// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iAssessExposureAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *AssessExposureAdvanceRequest
	GetImageURLObject() io.Reader
}

type AssessExposureAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/AssessExposure/AssessExposure1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s AssessExposureAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s AssessExposureAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AssessExposureAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *AssessExposureAdvanceRequest) SetImageURLObject(v io.Reader) *AssessExposureAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *AssessExposureAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
