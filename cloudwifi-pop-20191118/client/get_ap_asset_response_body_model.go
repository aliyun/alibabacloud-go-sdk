// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApAssetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *GetApAssetResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v int32) *GetApAssetResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *GetApAssetResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *GetApAssetResponseBody
	GetIsSuccess() *bool
}

type GetApAssetResponseBody struct {
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 0
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
}

func (s GetApAssetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApAssetResponseBody) GoString() string {
	return s.String()
}

func (s *GetApAssetResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetApAssetResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *GetApAssetResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetApAssetResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetApAssetResponseBody) SetData(v map[string]interface{}) *GetApAssetResponseBody {
	s.Data = v
	return s
}

func (s *GetApAssetResponseBody) SetErrorCode(v int32) *GetApAssetResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetApAssetResponseBody) SetErrorMessage(v string) *GetApAssetResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetApAssetResponseBody) SetIsSuccess(v bool) *GetApAssetResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetApAssetResponseBody) Validate() error {
	return dara.Validate(s)
}
