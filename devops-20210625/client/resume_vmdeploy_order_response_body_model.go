// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeVMDeployOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ResumeVMDeployOrderResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ResumeVMDeployOrderResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ResumeVMDeployOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ResumeVMDeployOrderResponseBody
	GetSuccess() *bool
}

type ResumeVMDeployOrderResponseBody struct {
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

func (s ResumeVMDeployOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumeVMDeployOrderResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeVMDeployOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ResumeVMDeployOrderResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ResumeVMDeployOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumeVMDeployOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ResumeVMDeployOrderResponseBody) SetErrorCode(v string) *ResumeVMDeployOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ResumeVMDeployOrderResponseBody) SetErrorMessage(v string) *ResumeVMDeployOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ResumeVMDeployOrderResponseBody) SetRequestId(v string) *ResumeVMDeployOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeVMDeployOrderResponseBody) SetSuccess(v bool) *ResumeVMDeployOrderResponseBody {
	s.Success = &v
	return s
}

func (s *ResumeVMDeployOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
