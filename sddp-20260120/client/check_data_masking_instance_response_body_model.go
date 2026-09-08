// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDataMaskingInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CheckDataMaskingInstanceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CheckDataMaskingInstanceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CheckDataMaskingInstanceResponseBody
	GetRequestId() *string
}

type CheckDataMaskingInstanceResponseBody struct {
	// example:
	//
	// Success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 实例未处于运行状态
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 7C6D8E9F-1234-5678-ABCD-0123456789AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckDataMaskingInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckDataMaskingInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDataMaskingInstanceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CheckDataMaskingInstanceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CheckDataMaskingInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckDataMaskingInstanceResponseBody) SetErrorCode(v string) *CheckDataMaskingInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CheckDataMaskingInstanceResponseBody) SetErrorMessage(v string) *CheckDataMaskingInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CheckDataMaskingInstanceResponseBody) SetRequestId(v string) *CheckDataMaskingInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDataMaskingInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
