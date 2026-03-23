// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKickStaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *KickStaResponseBody
	GetData() *string
	SetErrorCode(v int32) *KickStaResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *KickStaResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *KickStaResponseBody
	GetIsSuccess() *bool
}

type KickStaResponseBody struct {
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
}

func (s KickStaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s KickStaResponseBody) GoString() string {
	return s.String()
}

func (s *KickStaResponseBody) GetData() *string {
	return s.Data
}

func (s *KickStaResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *KickStaResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *KickStaResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *KickStaResponseBody) SetData(v string) *KickStaResponseBody {
	s.Data = &v
	return s
}

func (s *KickStaResponseBody) SetErrorCode(v int32) *KickStaResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *KickStaResponseBody) SetErrorMessage(v string) *KickStaResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *KickStaResponseBody) SetIsSuccess(v bool) *KickStaResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *KickStaResponseBody) Validate() error {
	return dara.Validate(s)
}
