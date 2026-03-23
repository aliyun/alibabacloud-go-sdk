// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterApAssetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *RegisterApAssetResponseBody
	GetData() *string
	SetErrorCode(v int32) *RegisterApAssetResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *RegisterApAssetResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *RegisterApAssetResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *RegisterApAssetResponseBody
	GetRequestId() *string
}

type RegisterApAssetResponseBody struct {
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterApAssetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegisterApAssetResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterApAssetResponseBody) GetData() *string {
	return s.Data
}

func (s *RegisterApAssetResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *RegisterApAssetResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RegisterApAssetResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *RegisterApAssetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegisterApAssetResponseBody) SetData(v string) *RegisterApAssetResponseBody {
	s.Data = &v
	return s
}

func (s *RegisterApAssetResponseBody) SetErrorCode(v int32) *RegisterApAssetResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RegisterApAssetResponseBody) SetErrorMessage(v string) *RegisterApAssetResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RegisterApAssetResponseBody) SetIsSuccess(v bool) *RegisterApAssetResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *RegisterApAssetResponseBody) SetRequestId(v string) *RegisterApAssetResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterApAssetResponseBody) Validate() error {
	return dara.Validate(s)
}
