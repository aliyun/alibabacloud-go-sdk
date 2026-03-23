// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApgroupIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *GetApgroupIdResponseBody
	GetData() *string
	SetErrorCode(v int32) *GetApgroupIdResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *GetApgroupIdResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *GetApgroupIdResponseBody
	GetIsSuccess() *bool
}

type GetApgroupIdResponseBody struct {
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
}

func (s GetApgroupIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApgroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetApgroupIdResponseBody) GetData() *string {
	return s.Data
}

func (s *GetApgroupIdResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *GetApgroupIdResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetApgroupIdResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetApgroupIdResponseBody) SetData(v string) *GetApgroupIdResponseBody {
	s.Data = &v
	return s
}

func (s *GetApgroupIdResponseBody) SetErrorCode(v int32) *GetApgroupIdResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetApgroupIdResponseBody) SetErrorMessage(v string) *GetApgroupIdResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetApgroupIdResponseBody) SetIsSuccess(v bool) *GetApgroupIdResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetApgroupIdResponseBody) Validate() error {
	return dara.Validate(s)
}
