// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApproveOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ApproveOrderResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ApproveOrderResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ApproveOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ApproveOrderResponseBody
	GetSuccess() *bool
}

type ApproveOrderResponseBody struct {
	// The error code returned.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 427688B8-ADFB-4C4E-9D45-EF5C1FD6E23D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ApproveOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApproveOrderResponseBody) GoString() string {
	return s.String()
}

func (s *ApproveOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ApproveOrderResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ApproveOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApproveOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ApproveOrderResponseBody) SetErrorCode(v string) *ApproveOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ApproveOrderResponseBody) SetErrorMessage(v string) *ApproveOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ApproveOrderResponseBody) SetRequestId(v string) *ApproveOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApproveOrderResponseBody) SetSuccess(v bool) *ApproveOrderResponseBody {
	s.Success = &v
	return s
}

func (s *ApproveOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
