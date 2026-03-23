// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutAppConfigAndSaveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *PutAppConfigAndSaveResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v int32) *PutAppConfigAndSaveResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *PutAppConfigAndSaveResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *PutAppConfigAndSaveResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *PutAppConfigAndSaveResponseBody
	GetRequestId() *string
}

type PutAppConfigAndSaveResponseBody struct {
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutAppConfigAndSaveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutAppConfigAndSaveResponseBody) GoString() string {
	return s.String()
}

func (s *PutAppConfigAndSaveResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *PutAppConfigAndSaveResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *PutAppConfigAndSaveResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *PutAppConfigAndSaveResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *PutAppConfigAndSaveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutAppConfigAndSaveResponseBody) SetData(v map[string]interface{}) *PutAppConfigAndSaveResponseBody {
	s.Data = v
	return s
}

func (s *PutAppConfigAndSaveResponseBody) SetErrorCode(v int32) *PutAppConfigAndSaveResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *PutAppConfigAndSaveResponseBody) SetErrorMessage(v string) *PutAppConfigAndSaveResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *PutAppConfigAndSaveResponseBody) SetIsSuccess(v bool) *PutAppConfigAndSaveResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *PutAppConfigAndSaveResponseBody) SetRequestId(v string) *PutAppConfigAndSaveResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutAppConfigAndSaveResponseBody) Validate() error {
	return dara.Validate(s)
}
