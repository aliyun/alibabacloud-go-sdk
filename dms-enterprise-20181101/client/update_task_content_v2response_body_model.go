// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskContentV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateTaskContentV2ResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateTaskContentV2ResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateTaskContentV2ResponseBody
	GetRequestId() *string
	SetSuccess(v string) *UpdateTaskContentV2ResponseBody
	GetSuccess() *string
}

type UpdateTaskContentV2ResponseBody struct {
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
	// Id of the request.
	//
	// example:
	//
	// B5FD0BC8-2D90-4478-B8EC-A0E92E0B1773
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
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTaskContentV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskContentV2ResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskContentV2ResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateTaskContentV2ResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateTaskContentV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTaskContentV2ResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *UpdateTaskContentV2ResponseBody) SetErrorCode(v string) *UpdateTaskContentV2ResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTaskContentV2ResponseBody) SetErrorMessage(v string) *UpdateTaskContentV2ResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateTaskContentV2ResponseBody) SetRequestId(v string) *UpdateTaskContentV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTaskContentV2ResponseBody) SetSuccess(v string) *UpdateTaskContentV2ResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTaskContentV2ResponseBody) Validate() error {
	return dara.Validate(s)
}
