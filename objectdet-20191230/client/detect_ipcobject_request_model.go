// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectIPCObjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *DetectIPCObjectRequest
	GetImageURL() *string
}

type DetectIPCObjectRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/objectdet/detect-ipc-xxxx.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectIPCObjectRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectIPCObjectRequest) GoString() string {
	return s.String()
}

func (s *DetectIPCObjectRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *DetectIPCObjectRequest) SetImageURL(v string) *DetectIPCObjectRequest {
	s.ImageURL = &v
	return s
}

func (s *DetectIPCObjectRequest) Validate() error {
	return dara.Validate(s)
}
