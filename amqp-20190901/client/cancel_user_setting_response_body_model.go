// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelUserSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CancelUserSettingResponseBody
	GetCode() *int32
	SetData(v string) *CancelUserSettingResponseBody
	GetData() *string
	SetMessage(v string) *CancelUserSettingResponseBody
	GetMessage() *string
	SetRequestId(v string) *CancelUserSettingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CancelUserSettingResponseBody
	GetSuccess() *bool
}

type CancelUserSettingResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelUserSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelUserSettingResponseBody) GoString() string {
	return s.String()
}

func (s *CancelUserSettingResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CancelUserSettingResponseBody) GetData() *string {
	return s.Data
}

func (s *CancelUserSettingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelUserSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelUserSettingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelUserSettingResponseBody) SetCode(v int32) *CancelUserSettingResponseBody {
	s.Code = &v
	return s
}

func (s *CancelUserSettingResponseBody) SetData(v string) *CancelUserSettingResponseBody {
	s.Data = &v
	return s
}

func (s *CancelUserSettingResponseBody) SetMessage(v string) *CancelUserSettingResponseBody {
	s.Message = &v
	return s
}

func (s *CancelUserSettingResponseBody) SetRequestId(v string) *CancelUserSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelUserSettingResponseBody) SetSuccess(v bool) *CancelUserSettingResponseBody {
	s.Success = &v
	return s
}

func (s *CancelUserSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
