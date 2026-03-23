// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApDetailStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *GetApDetailStatusResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v int32) *GetApDetailStatusResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *GetApDetailStatusResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *GetApDetailStatusResponseBody
	GetIsSuccess() *bool
}

type GetApDetailStatusResponseBody struct {
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
}

func (s GetApDetailStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApDetailStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetApDetailStatusResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetApDetailStatusResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *GetApDetailStatusResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetApDetailStatusResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetApDetailStatusResponseBody) SetData(v map[string]interface{}) *GetApDetailStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetApDetailStatusResponseBody) SetErrorCode(v int32) *GetApDetailStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetApDetailStatusResponseBody) SetErrorMessage(v string) *GetApDetailStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetApDetailStatusResponseBody) SetIsSuccess(v bool) *GetApDetailStatusResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetApDetailStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
