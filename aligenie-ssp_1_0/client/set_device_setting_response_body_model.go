// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDeviceSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *SetDeviceSettingResponseBody
	GetCode() *int32
	SetMessage(v string) *SetDeviceSettingResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetDeviceSettingResponseBody
	GetRequestId() *string
	SetResult(v bool) *SetDeviceSettingResponseBody
	GetResult() *bool
}

type SetDeviceSettingResponseBody struct {
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s SetDeviceSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDeviceSettingResponseBody) GoString() string {
	return s.String()
}

func (s *SetDeviceSettingResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *SetDeviceSettingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetDeviceSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDeviceSettingResponseBody) GetResult() *bool {
	return s.Result
}

func (s *SetDeviceSettingResponseBody) SetCode(v int32) *SetDeviceSettingResponseBody {
	s.Code = &v
	return s
}

func (s *SetDeviceSettingResponseBody) SetMessage(v string) *SetDeviceSettingResponseBody {
	s.Message = &v
	return s
}

func (s *SetDeviceSettingResponseBody) SetRequestId(v string) *SetDeviceSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDeviceSettingResponseBody) SetResult(v bool) *SetDeviceSettingResponseBody {
	s.Result = &v
	return s
}

func (s *SetDeviceSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
