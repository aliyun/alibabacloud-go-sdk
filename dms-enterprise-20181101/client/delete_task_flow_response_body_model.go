// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTaskFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteTaskFlowResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteTaskFlowResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteTaskFlowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteTaskFlowResponseBody
	GetSuccess() *bool
}

type DeleteTaskFlowResponseBody struct {
	// The error code returned if the request fails.
	//
	// example:
	//
	// 403
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9AFE13F6-A4FD-581B-BFDE-B63B1CDC2336
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**: The request is successful.
	//
	// 	- **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteTaskFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTaskFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTaskFlowResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteTaskFlowResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteTaskFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTaskFlowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteTaskFlowResponseBody) SetErrorCode(v string) *DeleteTaskFlowResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteTaskFlowResponseBody) SetErrorMessage(v string) *DeleteTaskFlowResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteTaskFlowResponseBody) SetRequestId(v string) *DeleteTaskFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTaskFlowResponseBody) SetSuccess(v bool) *DeleteTaskFlowResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteTaskFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
