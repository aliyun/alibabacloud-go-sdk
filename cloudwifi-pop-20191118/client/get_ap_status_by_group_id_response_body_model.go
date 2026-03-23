// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApStatusByGroupIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *GetApStatusByGroupIdResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v int32) *GetApStatusByGroupIdResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *GetApStatusByGroupIdResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *GetApStatusByGroupIdResponseBody
	GetIsSuccess() *bool
}

type GetApStatusByGroupIdResponseBody struct {
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
}

func (s GetApStatusByGroupIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApStatusByGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetApStatusByGroupIdResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetApStatusByGroupIdResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *GetApStatusByGroupIdResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetApStatusByGroupIdResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetApStatusByGroupIdResponseBody) SetData(v map[string]interface{}) *GetApStatusByGroupIdResponseBody {
	s.Data = v
	return s
}

func (s *GetApStatusByGroupIdResponseBody) SetErrorCode(v int32) *GetApStatusByGroupIdResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetApStatusByGroupIdResponseBody) SetErrorMessage(v string) *GetApStatusByGroupIdResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetApStatusByGroupIdResponseBody) SetIsSuccess(v bool) *GetApStatusByGroupIdResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetApStatusByGroupIdResponseBody) Validate() error {
	return dara.Validate(s)
}
