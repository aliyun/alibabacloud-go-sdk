// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetDeviceSettingResponseBody
	GetCode() *int32
	SetMessage(v string) *GetDeviceSettingResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDeviceSettingResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *GetDeviceSettingResponseBody
	GetResult() map[string]interface{}
}

type GetDeviceSettingResponseBody struct {
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
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    map[string]interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s GetDeviceSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceSettingResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceSettingResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetDeviceSettingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDeviceSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeviceSettingResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *GetDeviceSettingResponseBody) SetCode(v int32) *GetDeviceSettingResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeviceSettingResponseBody) SetMessage(v string) *GetDeviceSettingResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeviceSettingResponseBody) SetRequestId(v string) *GetDeviceSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceSettingResponseBody) SetResult(v map[string]interface{}) *GetDeviceSettingResponseBody {
	s.Result = v
	return s
}

func (s *GetDeviceSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
