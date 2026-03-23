// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetApNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *SetApNameResponseBody
	GetData() *string
	SetErrorCode(v int32) *SetApNameResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *SetApNameResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *SetApNameResponseBody
	GetIsSuccess() *bool
}

type SetApNameResponseBody struct {
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
}

func (s SetApNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetApNameResponseBody) GoString() string {
	return s.String()
}

func (s *SetApNameResponseBody) GetData() *string {
	return s.Data
}

func (s *SetApNameResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *SetApNameResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SetApNameResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *SetApNameResponseBody) SetData(v string) *SetApNameResponseBody {
	s.Data = &v
	return s
}

func (s *SetApNameResponseBody) SetErrorCode(v int32) *SetApNameResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SetApNameResponseBody) SetErrorMessage(v string) *SetApNameResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SetApNameResponseBody) SetIsSuccess(v bool) *SetApNameResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *SetApNameResponseBody) Validate() error {
	return dara.Validate(s)
}
