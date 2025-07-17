// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTryRunTaskFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *TryRunTaskFlowResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *TryRunTaskFlowResponseBody
	GetErrorMessage() *string
	SetInstanceId(v string) *TryRunTaskFlowResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *TryRunTaskFlowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TryRunTaskFlowResponseBody
	GetSuccess() *bool
}

type TryRunTaskFlowResponseBody struct {
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 169****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 9997630E-1993-5E6D-9DF1-4EFEE755FE31
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TryRunTaskFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TryRunTaskFlowResponseBody) GoString() string {
	return s.String()
}

func (s *TryRunTaskFlowResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *TryRunTaskFlowResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *TryRunTaskFlowResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *TryRunTaskFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TryRunTaskFlowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TryRunTaskFlowResponseBody) SetErrorCode(v string) *TryRunTaskFlowResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TryRunTaskFlowResponseBody) SetErrorMessage(v string) *TryRunTaskFlowResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *TryRunTaskFlowResponseBody) SetInstanceId(v string) *TryRunTaskFlowResponseBody {
	s.InstanceId = &v
	return s
}

func (s *TryRunTaskFlowResponseBody) SetRequestId(v string) *TryRunTaskFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *TryRunTaskFlowResponseBody) SetSuccess(v bool) *TryRunTaskFlowResponseBody {
	s.Success = &v
	return s
}

func (s *TryRunTaskFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
