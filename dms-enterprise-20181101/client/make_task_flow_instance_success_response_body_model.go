// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMakeTaskFlowInstanceSuccessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *MakeTaskFlowInstanceSuccessResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *MakeTaskFlowInstanceSuccessResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *MakeTaskFlowInstanceSuccessResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MakeTaskFlowInstanceSuccessResponseBody
	GetSuccess() *bool
}

type MakeTaskFlowInstanceSuccessResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 028BF827-3801-5869-8548-F4A039256304
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MakeTaskFlowInstanceSuccessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MakeTaskFlowInstanceSuccessResponseBody) GoString() string {
	return s.String()
}

func (s *MakeTaskFlowInstanceSuccessResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *MakeTaskFlowInstanceSuccessResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *MakeTaskFlowInstanceSuccessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MakeTaskFlowInstanceSuccessResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MakeTaskFlowInstanceSuccessResponseBody) SetErrorCode(v string) *MakeTaskFlowInstanceSuccessResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *MakeTaskFlowInstanceSuccessResponseBody) SetErrorMessage(v string) *MakeTaskFlowInstanceSuccessResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *MakeTaskFlowInstanceSuccessResponseBody) SetRequestId(v string) *MakeTaskFlowInstanceSuccessResponseBody {
	s.RequestId = &v
	return s
}

func (s *MakeTaskFlowInstanceSuccessResponseBody) SetSuccess(v bool) *MakeTaskFlowInstanceSuccessResponseBody {
	s.Success = &v
	return s
}

func (s *MakeTaskFlowInstanceSuccessResponseBody) Validate() error {
	return dara.Validate(s)
}
