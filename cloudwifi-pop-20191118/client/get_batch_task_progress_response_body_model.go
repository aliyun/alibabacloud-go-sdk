// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchTaskProgressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *GetBatchTaskProgressResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v int32) *GetBatchTaskProgressResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *GetBatchTaskProgressResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *GetBatchTaskProgressResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *GetBatchTaskProgressResponseBody
	GetRequestId() *string
}

type GetBatchTaskProgressResponseBody struct {
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetBatchTaskProgressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskProgressResponseBody) GoString() string {
	return s.String()
}

func (s *GetBatchTaskProgressResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetBatchTaskProgressResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *GetBatchTaskProgressResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetBatchTaskProgressResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetBatchTaskProgressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBatchTaskProgressResponseBody) SetData(v map[string]interface{}) *GetBatchTaskProgressResponseBody {
	s.Data = v
	return s
}

func (s *GetBatchTaskProgressResponseBody) SetErrorCode(v int32) *GetBatchTaskProgressResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetBatchTaskProgressResponseBody) SetErrorMessage(v string) *GetBatchTaskProgressResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetBatchTaskProgressResponseBody) SetIsSuccess(v bool) *GetBatchTaskProgressResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetBatchTaskProgressResponseBody) SetRequestId(v string) *GetBatchTaskProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBatchTaskProgressResponseBody) Validate() error {
	return dara.Validate(s)
}
