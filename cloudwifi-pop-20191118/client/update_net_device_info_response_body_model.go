// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNetDeviceInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*int64) *UpdateNetDeviceInfoResponseBody
	GetData() []*int64
	SetErrorCode(v int32) *UpdateNetDeviceInfoResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *UpdateNetDeviceInfoResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *UpdateNetDeviceInfoResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *UpdateNetDeviceInfoResponseBody
	GetRequestId() *string
}

type UpdateNetDeviceInfoResponseBody struct {
	Data         []*int64 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode    *int32   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool    `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId    *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateNetDeviceInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetDeviceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNetDeviceInfoResponseBody) GetData() []*int64 {
	return s.Data
}

func (s *UpdateNetDeviceInfoResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *UpdateNetDeviceInfoResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateNetDeviceInfoResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *UpdateNetDeviceInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNetDeviceInfoResponseBody) SetData(v []*int64) *UpdateNetDeviceInfoResponseBody {
	s.Data = v
	return s
}

func (s *UpdateNetDeviceInfoResponseBody) SetErrorCode(v int32) *UpdateNetDeviceInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateNetDeviceInfoResponseBody) SetErrorMessage(v string) *UpdateNetDeviceInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateNetDeviceInfoResponseBody) SetIsSuccess(v bool) *UpdateNetDeviceInfoResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UpdateNetDeviceInfoResponseBody) SetRequestId(v string) *UpdateNetDeviceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNetDeviceInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
