// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateTaskResponseBody
	GetCode() *int32
	SetDetails(v string) *UpdateTaskResponseBody
	GetDetails() *string
	SetErrorCode(v string) *UpdateTaskResponseBody
	GetErrorCode() *string
	SetMessage(v string) *UpdateTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTaskResponseBody
	GetSuccess() *bool
}

type UpdateTaskResponseBody struct {
	// Total amount of data under the conditions of this request. This parameter is optional and does not need to be returned by default.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details
	//
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// error code
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Return message of the request
	//
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// is succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateTaskResponseBody) GetDetails() *string {
	return s.Details
}

func (s *UpdateTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTaskResponseBody) SetCode(v int32) *UpdateTaskResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTaskResponseBody) SetDetails(v string) *UpdateTaskResponseBody {
	s.Details = &v
	return s
}

func (s *UpdateTaskResponseBody) SetErrorCode(v string) *UpdateTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTaskResponseBody) SetMessage(v string) *UpdateTaskResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTaskResponseBody) SetRequestId(v string) *UpdateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTaskResponseBody) SetSuccess(v bool) *UpdateTaskResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
