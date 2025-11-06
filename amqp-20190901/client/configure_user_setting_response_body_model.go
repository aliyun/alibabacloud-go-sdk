// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureUserSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ConfigureUserSettingResponseBody
	GetCode() *int32
	SetData(v string) *ConfigureUserSettingResponseBody
	GetData() *string
	SetMessage(v string) *ConfigureUserSettingResponseBody
	GetMessage() *string
	SetRequestId(v string) *ConfigureUserSettingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ConfigureUserSettingResponseBody
	GetSuccess() *bool
}

type ConfigureUserSettingResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfigureUserSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigureUserSettingResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigureUserSettingResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ConfigureUserSettingResponseBody) GetData() *string {
	return s.Data
}

func (s *ConfigureUserSettingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ConfigureUserSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigureUserSettingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ConfigureUserSettingResponseBody) SetCode(v int32) *ConfigureUserSettingResponseBody {
	s.Code = &v
	return s
}

func (s *ConfigureUserSettingResponseBody) SetData(v string) *ConfigureUserSettingResponseBody {
	s.Data = &v
	return s
}

func (s *ConfigureUserSettingResponseBody) SetMessage(v string) *ConfigureUserSettingResponseBody {
	s.Message = &v
	return s
}

func (s *ConfigureUserSettingResponseBody) SetRequestId(v string) *ConfigureUserSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigureUserSettingResponseBody) SetSuccess(v bool) *ConfigureUserSettingResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigureUserSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
