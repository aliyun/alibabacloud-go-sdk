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
	// The instance check result code. A value of Success indicates that the check is passed. For other values, refer to ErrorMessage for troubleshooting.
	//
	// example:
	//
	// Success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The instance check result message. If the check fails, the specific cause is returned.
	//
	// example:
	//
	// The instance is not in the running state
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID. You can use the ID to locate this API call.
	//
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
