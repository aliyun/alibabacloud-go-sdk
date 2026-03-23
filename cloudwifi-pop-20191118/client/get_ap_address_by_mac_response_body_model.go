// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApAddressByMacResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *GetApAddressByMacResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v int32) *GetApAddressByMacResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *GetApAddressByMacResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *GetApAddressByMacResponseBody
	GetIsSuccess() *bool
}

type GetApAddressByMacResponseBody struct {
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
}

func (s GetApAddressByMacResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApAddressByMacResponseBody) GoString() string {
	return s.String()
}

func (s *GetApAddressByMacResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetApAddressByMacResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *GetApAddressByMacResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetApAddressByMacResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetApAddressByMacResponseBody) SetData(v map[string]interface{}) *GetApAddressByMacResponseBody {
	s.Data = v
	return s
}

func (s *GetApAddressByMacResponseBody) SetErrorCode(v int32) *GetApAddressByMacResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetApAddressByMacResponseBody) SetErrorMessage(v string) *GetApAddressByMacResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetApAddressByMacResponseBody) SetIsSuccess(v bool) *GetApAddressByMacResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetApAddressByMacResponseBody) Validate() error {
	return dara.Validate(s)
}
