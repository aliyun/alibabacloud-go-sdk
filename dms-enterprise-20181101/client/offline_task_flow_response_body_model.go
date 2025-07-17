// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineTaskFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *OfflineTaskFlowResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *OfflineTaskFlowResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *OfflineTaskFlowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OfflineTaskFlowResponseBody
	GetSuccess() *bool
}

type OfflineTaskFlowResponseBody struct {
	// The error code returned if the request fails.
	//
	// example:
	//
	// UnknownError
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
	// A8FE12AA-300D-5FDF-806F-C2CB99161F32
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

func (s OfflineTaskFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OfflineTaskFlowResponseBody) GoString() string {
	return s.String()
}

func (s *OfflineTaskFlowResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *OfflineTaskFlowResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *OfflineTaskFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OfflineTaskFlowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OfflineTaskFlowResponseBody) SetErrorCode(v string) *OfflineTaskFlowResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *OfflineTaskFlowResponseBody) SetErrorMessage(v string) *OfflineTaskFlowResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *OfflineTaskFlowResponseBody) SetRequestId(v string) *OfflineTaskFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *OfflineTaskFlowResponseBody) SetSuccess(v bool) *OfflineTaskFlowResponseBody {
	s.Success = &v
	return s
}

func (s *OfflineTaskFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
