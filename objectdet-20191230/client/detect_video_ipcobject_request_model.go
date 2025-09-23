// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectVideoIPCObjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallbackOnlyHasObject(v bool) *DetectVideoIPCObjectRequest
	GetCallbackOnlyHasObject() *bool
	SetStartTimestamp(v int64) *DetectVideoIPCObjectRequest
	GetStartTimestamp() *int64
	SetVideoURL(v string) *DetectVideoIPCObjectRequest
	GetVideoURL() *string
}

type DetectVideoIPCObjectRequest struct {
	// example:
	//
	// true
	CallbackOnlyHasObject *bool `json:"CallbackOnlyHasObject,omitempty" xml:"CallbackOnlyHasObject,omitempty"`
	// example:
	//
	// 1629086400
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/objectdet/DetectVideoIPCObject/DetectVideoIPCObject1.mp4
	VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s DetectVideoIPCObjectRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectVideoIPCObjectRequest) GoString() string {
	return s.String()
}

func (s *DetectVideoIPCObjectRequest) GetCallbackOnlyHasObject() *bool {
	return s.CallbackOnlyHasObject
}

func (s *DetectVideoIPCObjectRequest) GetStartTimestamp() *int64 {
	return s.StartTimestamp
}

func (s *DetectVideoIPCObjectRequest) GetVideoURL() *string {
	return s.VideoURL
}

func (s *DetectVideoIPCObjectRequest) SetCallbackOnlyHasObject(v bool) *DetectVideoIPCObjectRequest {
	s.CallbackOnlyHasObject = &v
	return s
}

func (s *DetectVideoIPCObjectRequest) SetStartTimestamp(v int64) *DetectVideoIPCObjectRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DetectVideoIPCObjectRequest) SetVideoURL(v string) *DetectVideoIPCObjectRequest {
	s.VideoURL = &v
	return s
}

func (s *DetectVideoIPCObjectRequest) Validate() error {
	return dara.Validate(s)
}
