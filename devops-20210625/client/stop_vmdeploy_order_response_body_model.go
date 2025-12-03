// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopVMDeployOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *StopVMDeployOrderResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *StopVMDeployOrderResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *StopVMDeployOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopVMDeployOrderResponseBody
	GetSuccess() *bool
}

type StopVMDeployOrderResponseBody struct {
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

func (s StopVMDeployOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopVMDeployOrderResponseBody) GoString() string {
	return s.String()
}

func (s *StopVMDeployOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *StopVMDeployOrderResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *StopVMDeployOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopVMDeployOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopVMDeployOrderResponseBody) SetErrorCode(v string) *StopVMDeployOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StopVMDeployOrderResponseBody) SetErrorMessage(v string) *StopVMDeployOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StopVMDeployOrderResponseBody) SetRequestId(v string) *StopVMDeployOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopVMDeployOrderResponseBody) SetSuccess(v bool) *StopVMDeployOrderResponseBody {
	s.Success = &v
	return s
}

func (s *StopVMDeployOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
