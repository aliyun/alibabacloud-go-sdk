// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iDetectIPCObjectAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *DetectIPCObjectAdvanceRequest
	GetImageURLObject() io.Reader
}

type DetectIPCObjectAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/objectdet/detect-ipc-xxxx.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectIPCObjectAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectIPCObjectAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectIPCObjectAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *DetectIPCObjectAdvanceRequest) SetImageURLObject(v io.Reader) *DetectIPCObjectAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *DetectIPCObjectAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
