// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTaskFlowEdgesByConditionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteTaskFlowEdgesByConditionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteTaskFlowEdgesByConditionResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteTaskFlowEdgesByConditionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteTaskFlowEdgesByConditionResponseBody
	GetSuccess() *bool
}

type DeleteTaskFlowEdgesByConditionResponseBody struct {
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
	// The ID of the request. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 21234B66-6859-5558-9E5B-006EFE915CD0
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

func (s DeleteTaskFlowEdgesByConditionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTaskFlowEdgesByConditionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTaskFlowEdgesByConditionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteTaskFlowEdgesByConditionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteTaskFlowEdgesByConditionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTaskFlowEdgesByConditionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteTaskFlowEdgesByConditionResponseBody) SetErrorCode(v string) *DeleteTaskFlowEdgesByConditionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteTaskFlowEdgesByConditionResponseBody) SetErrorMessage(v string) *DeleteTaskFlowEdgesByConditionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteTaskFlowEdgesByConditionResponseBody) SetRequestId(v string) *DeleteTaskFlowEdgesByConditionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTaskFlowEdgesByConditionResponseBody) SetSuccess(v bool) *DeleteTaskFlowEdgesByConditionResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteTaskFlowEdgesByConditionResponseBody) Validate() error {
	return dara.Validate(s)
}
