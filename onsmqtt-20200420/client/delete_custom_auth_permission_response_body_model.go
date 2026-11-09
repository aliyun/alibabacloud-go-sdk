// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomAuthPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteCustomAuthPermissionResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteCustomAuthPermissionResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCustomAuthPermissionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCustomAuthPermissionResponseBody
	GetSuccess() *bool
}

type DeleteCustomAuthPermissionResponseBody struct {
	// The code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned by the request.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 34063BCA-0946-49C1-B824-2ED2C905****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - true: The request is successful.
	//
	// - false: The request fails.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCustomAuthPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomAuthPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomAuthPermissionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteCustomAuthPermissionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCustomAuthPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomAuthPermissionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCustomAuthPermissionResponseBody) SetCode(v int32) *DeleteCustomAuthPermissionResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCustomAuthPermissionResponseBody) SetMessage(v string) *DeleteCustomAuthPermissionResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCustomAuthPermissionResponseBody) SetRequestId(v string) *DeleteCustomAuthPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomAuthPermissionResponseBody) SetSuccess(v bool) *DeleteCustomAuthPermissionResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCustomAuthPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
