// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iDetectVideoIPCObjectAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallbackOnlyHasObject(v bool) *DetectVideoIPCObjectAdvanceRequest
	GetCallbackOnlyHasObject() *bool
	SetStartTimestamp(v int64) *DetectVideoIPCObjectAdvanceRequest
	GetStartTimestamp() *int64
	SetVideoURLObject(v io.Reader) *DetectVideoIPCObjectAdvanceRequest
	GetVideoURLObject() io.Reader
}

type DetectVideoIPCObjectAdvanceRequest struct {
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
	VideoURLObject io.Reader `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s DetectVideoIPCObjectAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectVideoIPCObjectAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectVideoIPCObjectAdvanceRequest) GetCallbackOnlyHasObject() *bool {
	return s.CallbackOnlyHasObject
}

func (s *DetectVideoIPCObjectAdvanceRequest) GetStartTimestamp() *int64 {
	return s.StartTimestamp
}

func (s *DetectVideoIPCObjectAdvanceRequest) GetVideoURLObject() io.Reader {
	return s.VideoURLObject
}

func (s *DetectVideoIPCObjectAdvanceRequest) SetCallbackOnlyHasObject(v bool) *DetectVideoIPCObjectAdvanceRequest {
	s.CallbackOnlyHasObject = &v
	return s
}

func (s *DetectVideoIPCObjectAdvanceRequest) SetStartTimestamp(v int64) *DetectVideoIPCObjectAdvanceRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DetectVideoIPCObjectAdvanceRequest) SetVideoURLObject(v io.Reader) *DetectVideoIPCObjectAdvanceRequest {
	s.VideoURLObject = v
	return s
}

func (s *DetectVideoIPCObjectAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
