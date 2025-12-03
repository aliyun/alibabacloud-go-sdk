// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryVMDeployMachineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *RetryVMDeployMachineResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *RetryVMDeployMachineResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *RetryVMDeployMachineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RetryVMDeployMachineResponseBody
	GetSuccess() *bool
}

type RetryVMDeployMachineResponseBody struct {
	// example:
	//
	// ”“
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ”“
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RetryVMDeployMachineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetryVMDeployMachineResponseBody) GoString() string {
	return s.String()
}

func (s *RetryVMDeployMachineResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RetryVMDeployMachineResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RetryVMDeployMachineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetryVMDeployMachineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RetryVMDeployMachineResponseBody) SetErrorCode(v string) *RetryVMDeployMachineResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RetryVMDeployMachineResponseBody) SetErrorMessage(v string) *RetryVMDeployMachineResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RetryVMDeployMachineResponseBody) SetRequestId(v string) *RetryVMDeployMachineResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryVMDeployMachineResponseBody) SetSuccess(v bool) *RetryVMDeployMachineResponseBody {
	s.Success = &v
	return s
}

func (s *RetryVMDeployMachineResponseBody) Validate() error {
	return dara.Validate(s)
}
