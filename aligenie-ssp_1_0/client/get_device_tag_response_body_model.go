// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetDeviceTagResponseBody
	GetCode() *int32
	SetMessage(v string) *GetDeviceTagResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDeviceTagResponseBody
	GetRequestId() *string
	SetResult(v *GetDeviceTagResponseBodyResult) *GetDeviceTagResponseBody
	GetResult() *GetDeviceTagResponseBodyResult
}

type GetDeviceTagResponseBody struct {
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
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetDeviceTagResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetDeviceTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceTagResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceTagResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetDeviceTagResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDeviceTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeviceTagResponseBody) GetResult() *GetDeviceTagResponseBodyResult {
	return s.Result
}

func (s *GetDeviceTagResponseBody) SetCode(v int32) *GetDeviceTagResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceTagResponseBody) SetMessage(v string) *GetDeviceTagResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceTagResponseBody) SetRequestId(v string) *GetDeviceTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceTagResponseBody) SetResult(v *GetDeviceTagResponseBodyResult) *GetDeviceTagResponseBody {
	s.Result = v
	return s
}

func (s *GetDeviceTagResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDeviceTagResponseBodyResult struct {
	// example:
	//
	// {       "antest1": "antest1",       "antest": "a"     }
	DeviceTags map[string]interface{} `json:"DeviceTags,omitempty" xml:"DeviceTags,omitempty"`
}

func (s GetDeviceTagResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceTagResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetDeviceTagResponseBodyResult) GetDeviceTags() map[string]interface{} {
	return s.DeviceTags
}

func (s *GetDeviceTagResponseBodyResult) SetDeviceTags(v map[string]interface{}) *GetDeviceTagResponseBodyResult {
	s.DeviceTags = v
	return s
}

func (s *GetDeviceTagResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
