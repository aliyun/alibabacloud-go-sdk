// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iDetectObjectAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *DetectObjectAdvanceRequest
	GetImageURLObject() io.Reader
}

type DetectObjectAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/objectdet/DetectObject/DetectObject1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectObjectAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectObjectAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectObjectAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *DetectObjectAdvanceRequest) SetImageURLObject(v io.Reader) *DetectObjectAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *DetectObjectAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
