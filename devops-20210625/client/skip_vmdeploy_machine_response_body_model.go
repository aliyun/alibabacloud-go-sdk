// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSkipVMDeployMachineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *SkipVMDeployMachineResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *SkipVMDeployMachineResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *SkipVMDeployMachineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SkipVMDeployMachineResponseBody
	GetSuccess() *bool
}

type SkipVMDeployMachineResponseBody struct {
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

func (s SkipVMDeployMachineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SkipVMDeployMachineResponseBody) GoString() string {
	return s.String()
}

func (s *SkipVMDeployMachineResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SkipVMDeployMachineResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SkipVMDeployMachineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SkipVMDeployMachineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SkipVMDeployMachineResponseBody) SetErrorCode(v string) *SkipVMDeployMachineResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SkipVMDeployMachineResponseBody) SetErrorMessage(v string) *SkipVMDeployMachineResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SkipVMDeployMachineResponseBody) SetRequestId(v string) *SkipVMDeployMachineResponseBody {
	s.RequestId = &v
	return s
}

func (s *SkipVMDeployMachineResponseBody) SetSuccess(v bool) *SkipVMDeployMachineResponseBody {
	s.Success = &v
	return s
}

func (s *SkipVMDeployMachineResponseBody) Validate() error {
	return dara.Validate(s)
}
