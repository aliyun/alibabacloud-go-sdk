// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iDetectWhiteBaseImageAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *DetectWhiteBaseImageAdvanceRequest
	GetImageURLObject() io.Reader
}

type DetectWhiteBaseImageAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/objectdet/DetectWhiteBaseImage/DetectWhiteBaseImage1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectWhiteBaseImageAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectWhiteBaseImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectWhiteBaseImageAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *DetectWhiteBaseImageAdvanceRequest) SetImageURLObject(v io.Reader) *DetectWhiteBaseImageAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *DetectWhiteBaseImageAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
