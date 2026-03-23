// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnRegisterApAssetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *UnRegisterApAssetResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v int32) *UnRegisterApAssetResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *UnRegisterApAssetResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *UnRegisterApAssetResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *UnRegisterApAssetResponseBody
	GetRequestId() *string
}

type UnRegisterApAssetResponseBody struct {
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnRegisterApAssetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnRegisterApAssetResponseBody) GoString() string {
	return s.String()
}

func (s *UnRegisterApAssetResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *UnRegisterApAssetResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *UnRegisterApAssetResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UnRegisterApAssetResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *UnRegisterApAssetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnRegisterApAssetResponseBody) SetData(v map[string]interface{}) *UnRegisterApAssetResponseBody {
	s.Data = v
	return s
}

func (s *UnRegisterApAssetResponseBody) SetErrorCode(v int32) *UnRegisterApAssetResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UnRegisterApAssetResponseBody) SetErrorMessage(v string) *UnRegisterApAssetResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UnRegisterApAssetResponseBody) SetIsSuccess(v bool) *UnRegisterApAssetResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UnRegisterApAssetResponseBody) SetRequestId(v string) *UnRegisterApAssetResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnRegisterApAssetResponseBody) Validate() error {
	return dara.Validate(s)
}
