// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomAuthPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateCustomAuthPermissionResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateCustomAuthPermissionResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateCustomAuthPermissionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCustomAuthPermissionResponseBody
	GetSuccess() *bool
}

type UpdateCustomAuthPermissionResponseBody struct {
	// Error code returned upon failed invocation. For more information, see Error Codes.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Information
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 8CC04203-679B-4DED-89D9-E7C2E979****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. true: Call succeeded. false: Call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCustomAuthPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAuthPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomAuthPermissionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateCustomAuthPermissionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateCustomAuthPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCustomAuthPermissionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCustomAuthPermissionResponseBody) SetCode(v int32) *UpdateCustomAuthPermissionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateCustomAuthPermissionResponseBody) SetMessage(v string) *UpdateCustomAuthPermissionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCustomAuthPermissionResponseBody) SetRequestId(v string) *UpdateCustomAuthPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCustomAuthPermissionResponseBody) SetSuccess(v bool) *UpdateCustomAuthPermissionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCustomAuthPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
