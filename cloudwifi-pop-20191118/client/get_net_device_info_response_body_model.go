// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetDeviceInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []map[string]interface{}) *GetNetDeviceInfoResponseBody
	GetData() []map[string]interface{}
	SetErrorCode(v int32) *GetNetDeviceInfoResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *GetNetDeviceInfoResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *GetNetDeviceInfoResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *GetNetDeviceInfoResponseBody
	GetRequestId() *string
}

type GetNetDeviceInfoResponseBody struct {
	Data         []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode    *int32                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                    `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId    *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNetDeviceInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNetDeviceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetNetDeviceInfoResponseBody) GetData() []map[string]interface{} {
	return s.Data
}

func (s *GetNetDeviceInfoResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *GetNetDeviceInfoResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetNetDeviceInfoResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetNetDeviceInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNetDeviceInfoResponseBody) SetData(v []map[string]interface{}) *GetNetDeviceInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetNetDeviceInfoResponseBody) SetErrorCode(v int32) *GetNetDeviceInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetNetDeviceInfoResponseBody) SetErrorMessage(v string) *GetNetDeviceInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetNetDeviceInfoResponseBody) SetIsSuccess(v bool) *GetNetDeviceInfoResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetNetDeviceInfoResponseBody) SetRequestId(v string) *GetNetDeviceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNetDeviceInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
