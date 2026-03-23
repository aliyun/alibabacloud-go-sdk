// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddApListToApgroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *AddApListToApgroupResponseBody
	GetData() *string
	SetErrorCode(v int32) *AddApListToApgroupResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *AddApListToApgroupResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *AddApListToApgroupResponseBody
	GetIsSuccess() *bool
}

type AddApListToApgroupResponseBody struct {
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
}

func (s AddApListToApgroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddApListToApgroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddApListToApgroupResponseBody) GetData() *string {
	return s.Data
}

func (s *AddApListToApgroupResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *AddApListToApgroupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AddApListToApgroupResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *AddApListToApgroupResponseBody) SetData(v string) *AddApListToApgroupResponseBody {
	s.Data = &v
	return s
}

func (s *AddApListToApgroupResponseBody) SetErrorCode(v int32) *AddApListToApgroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddApListToApgroupResponseBody) SetErrorMessage(v string) *AddApListToApgroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddApListToApgroupResponseBody) SetIsSuccess(v bool) *AddApListToApgroupResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *AddApListToApgroupResponseBody) Validate() error {
	return dara.Validate(s)
}
