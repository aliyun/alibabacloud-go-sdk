// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetApAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *SetApAddressResponseBody
	GetData() *string
	SetErrorCode(v int32) *SetApAddressResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *SetApAddressResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *SetApAddressResponseBody
	GetIsSuccess() *bool
}

type SetApAddressResponseBody struct {
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
}

func (s SetApAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetApAddressResponseBody) GoString() string {
	return s.String()
}

func (s *SetApAddressResponseBody) GetData() *string {
	return s.Data
}

func (s *SetApAddressResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *SetApAddressResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SetApAddressResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *SetApAddressResponseBody) SetData(v string) *SetApAddressResponseBody {
	s.Data = &v
	return s
}

func (s *SetApAddressResponseBody) SetErrorCode(v int32) *SetApAddressResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SetApAddressResponseBody) SetErrorMessage(v string) *SetApAddressResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SetApAddressResponseBody) SetIsSuccess(v bool) *SetApAddressResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *SetApAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
