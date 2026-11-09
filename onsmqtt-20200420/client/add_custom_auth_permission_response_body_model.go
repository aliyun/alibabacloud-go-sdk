// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomAuthPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddCustomAuthPermissionResponseBody
	GetCode() *int32
	SetMessage(v string) *AddCustomAuthPermissionResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddCustomAuthPermissionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddCustomAuthPermissionResponseBody
	GetSuccess() *bool
}

type AddCustomAuthPermissionResponseBody struct {
	// The status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82B9E503-F4A1-4F30-976F-C6999FF9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddCustomAuthPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCustomAuthPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *AddCustomAuthPermissionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddCustomAuthPermissionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddCustomAuthPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCustomAuthPermissionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddCustomAuthPermissionResponseBody) SetCode(v int32) *AddCustomAuthPermissionResponseBody {
	s.Code = &v
	return s
}

func (s *AddCustomAuthPermissionResponseBody) SetMessage(v string) *AddCustomAuthPermissionResponseBody {
	s.Message = &v
	return s
}

func (s *AddCustomAuthPermissionResponseBody) SetRequestId(v string) *AddCustomAuthPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCustomAuthPermissionResponseBody) SetSuccess(v bool) *AddCustomAuthPermissionResponseBody {
	s.Success = &v
	return s
}

func (s *AddCustomAuthPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
