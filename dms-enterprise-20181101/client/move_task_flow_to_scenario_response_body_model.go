// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveTaskFlowToScenarioResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *MoveTaskFlowToScenarioResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *MoveTaskFlowToScenarioResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *MoveTaskFlowToScenarioResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MoveTaskFlowToScenarioResponseBody
	GetSuccess() *bool
}

type MoveTaskFlowToScenarioResponseBody struct {
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
	// D85FD18C-4322-5D49-8C14-5A10E668F86C
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

func (s MoveTaskFlowToScenarioResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MoveTaskFlowToScenarioResponseBody) GoString() string {
	return s.String()
}

func (s *MoveTaskFlowToScenarioResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *MoveTaskFlowToScenarioResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *MoveTaskFlowToScenarioResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MoveTaskFlowToScenarioResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MoveTaskFlowToScenarioResponseBody) SetErrorCode(v string) *MoveTaskFlowToScenarioResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *MoveTaskFlowToScenarioResponseBody) SetErrorMessage(v string) *MoveTaskFlowToScenarioResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *MoveTaskFlowToScenarioResponseBody) SetRequestId(v string) *MoveTaskFlowToScenarioResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveTaskFlowToScenarioResponseBody) SetSuccess(v bool) *MoveTaskFlowToScenarioResponseBody {
	s.Success = &v
	return s
}

func (s *MoveTaskFlowToScenarioResponseBody) Validate() error {
	return dara.Validate(s)
}
