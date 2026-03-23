// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNewApgroupConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *NewApgroupConfigResponseBody
	GetData() *string
	SetErrorCode(v int32) *NewApgroupConfigResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *NewApgroupConfigResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *NewApgroupConfigResponseBody
	GetIsSuccess() *bool
}

type NewApgroupConfigResponseBody struct {
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
}

func (s NewApgroupConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s NewApgroupConfigResponseBody) GoString() string {
	return s.String()
}

func (s *NewApgroupConfigResponseBody) GetData() *string {
	return s.Data
}

func (s *NewApgroupConfigResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *NewApgroupConfigResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *NewApgroupConfigResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *NewApgroupConfigResponseBody) SetData(v string) *NewApgroupConfigResponseBody {
	s.Data = &v
	return s
}

func (s *NewApgroupConfigResponseBody) SetErrorCode(v int32) *NewApgroupConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *NewApgroupConfigResponseBody) SetErrorMessage(v string) *NewApgroupConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *NewApgroupConfigResponseBody) SetIsSuccess(v bool) *NewApgroupConfigResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *NewApgroupConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
