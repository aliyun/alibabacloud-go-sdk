// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetDeviceInfoWithSizeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetNetDeviceInfoWithSizeResponseBodyData) *GetNetDeviceInfoWithSizeResponseBody
	GetData() *GetNetDeviceInfoWithSizeResponseBodyData
	SetErrorCode(v int32) *GetNetDeviceInfoWithSizeResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *GetNetDeviceInfoWithSizeResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *GetNetDeviceInfoWithSizeResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *GetNetDeviceInfoWithSizeResponseBody
	GetRequestId() *string
}

type GetNetDeviceInfoWithSizeResponseBody struct {
	Data         *GetNetDeviceInfoWithSizeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode    *int32                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                                     `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNetDeviceInfoWithSizeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNetDeviceInfoWithSizeResponseBody) GoString() string {
	return s.String()
}

func (s *GetNetDeviceInfoWithSizeResponseBody) GetData() *GetNetDeviceInfoWithSizeResponseBodyData {
	return s.Data
}

func (s *GetNetDeviceInfoWithSizeResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *GetNetDeviceInfoWithSizeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetNetDeviceInfoWithSizeResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetNetDeviceInfoWithSizeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNetDeviceInfoWithSizeResponseBody) SetData(v *GetNetDeviceInfoWithSizeResponseBodyData) *GetNetDeviceInfoWithSizeResponseBody {
	s.Data = v
	return s
}

func (s *GetNetDeviceInfoWithSizeResponseBody) SetErrorCode(v int32) *GetNetDeviceInfoWithSizeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetNetDeviceInfoWithSizeResponseBody) SetErrorMessage(v string) *GetNetDeviceInfoWithSizeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetNetDeviceInfoWithSizeResponseBody) SetIsSuccess(v bool) *GetNetDeviceInfoWithSizeResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetNetDeviceInfoWithSizeResponseBody) SetRequestId(v string) *GetNetDeviceInfoWithSizeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNetDeviceInfoWithSizeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNetDeviceInfoWithSizeResponseBodyData struct {
	Count *int64                   `json:"Count,omitempty" xml:"Count,omitempty"`
	Data  []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s GetNetDeviceInfoWithSizeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetNetDeviceInfoWithSizeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNetDeviceInfoWithSizeResponseBodyData) GetCount() *int64 {
	return s.Count
}

func (s *GetNetDeviceInfoWithSizeResponseBodyData) GetData() []map[string]interface{} {
	return s.Data
}

func (s *GetNetDeviceInfoWithSizeResponseBodyData) SetCount(v int64) *GetNetDeviceInfoWithSizeResponseBodyData {
	s.Count = &v
	return s
}

func (s *GetNetDeviceInfoWithSizeResponseBodyData) SetData(v []map[string]interface{}) *GetNetDeviceInfoWithSizeResponseBodyData {
	s.Data = v
	return s
}

func (s *GetNetDeviceInfoWithSizeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
