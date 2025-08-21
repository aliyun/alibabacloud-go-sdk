// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceBasicInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListDeviceBasicInfoResponseBody
	GetCode() *int32
	SetMessage(v string) *ListDeviceBasicInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListDeviceBasicInfoResponseBody
	GetRequestId() *string
	SetResult(v map[string]*ResultValue) *ListDeviceBasicInfoResponseBody
	GetResult() map[string]*ResultValue
}

type ListDeviceBasicInfoResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    map[string]*ResultValue `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ListDeviceBasicInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceBasicInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceBasicInfoResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListDeviceBasicInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDeviceBasicInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDeviceBasicInfoResponseBody) GetResult() map[string]*ResultValue {
	return s.Result
}

func (s *ListDeviceBasicInfoResponseBody) SetCode(v int32) *ListDeviceBasicInfoResponseBody {
	s.Code = &v
	return s
}

func (s *ListDeviceBasicInfoResponseBody) SetMessage(v string) *ListDeviceBasicInfoResponseBody {
	s.Message = &v
	return s
}

func (s *ListDeviceBasicInfoResponseBody) SetRequestId(v string) *ListDeviceBasicInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeviceBasicInfoResponseBody) SetResult(v map[string]*ResultValue) *ListDeviceBasicInfoResponseBody {
	s.Result = v
	return s
}

func (s *ListDeviceBasicInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
