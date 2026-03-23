// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetDeviceInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*int64) *DeleteNetDeviceInfoResponseBody
	GetData() []*int64
	SetErrorCode(v int32) *DeleteNetDeviceInfoResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *DeleteNetDeviceInfoResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *DeleteNetDeviceInfoResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *DeleteNetDeviceInfoResponseBody
	GetRequestId() *string
}

type DeleteNetDeviceInfoResponseBody struct {
	Data         []*int64 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode    *int32   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool    `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId    *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNetDeviceInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetDeviceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNetDeviceInfoResponseBody) GetData() []*int64 {
	return s.Data
}

func (s *DeleteNetDeviceInfoResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *DeleteNetDeviceInfoResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteNetDeviceInfoResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteNetDeviceInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNetDeviceInfoResponseBody) SetData(v []*int64) *DeleteNetDeviceInfoResponseBody {
	s.Data = v
	return s
}

func (s *DeleteNetDeviceInfoResponseBody) SetErrorCode(v int32) *DeleteNetDeviceInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteNetDeviceInfoResponseBody) SetErrorMessage(v string) *DeleteNetDeviceInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteNetDeviceInfoResponseBody) SetIsSuccess(v bool) *DeleteNetDeviceInfoResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteNetDeviceInfoResponseBody) SetRequestId(v string) *DeleteNetDeviceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNetDeviceInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
