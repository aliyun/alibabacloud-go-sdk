// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNewNetDeviceInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*int64) *NewNetDeviceInfoResponseBody
	GetData() []*int64
	SetErrorCode(v int32) *NewNetDeviceInfoResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *NewNetDeviceInfoResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *NewNetDeviceInfoResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *NewNetDeviceInfoResponseBody
	GetRequestId() *string
}

type NewNetDeviceInfoResponseBody struct {
	Data         []*int64 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode    *int32   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool    `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId    *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s NewNetDeviceInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s NewNetDeviceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *NewNetDeviceInfoResponseBody) GetData() []*int64 {
	return s.Data
}

func (s *NewNetDeviceInfoResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *NewNetDeviceInfoResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *NewNetDeviceInfoResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *NewNetDeviceInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *NewNetDeviceInfoResponseBody) SetData(v []*int64) *NewNetDeviceInfoResponseBody {
	s.Data = v
	return s
}

func (s *NewNetDeviceInfoResponseBody) SetErrorCode(v int32) *NewNetDeviceInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *NewNetDeviceInfoResponseBody) SetErrorMessage(v string) *NewNetDeviceInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *NewNetDeviceInfoResponseBody) SetIsSuccess(v bool) *NewNetDeviceInfoResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *NewNetDeviceInfoResponseBody) SetRequestId(v string) *NewNetDeviceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *NewNetDeviceInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
