// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOpenSourcePermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateOpenSourcePermissionResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateOpenSourcePermissionResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateOpenSourcePermissionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateOpenSourcePermissionResponseBody
	GetSuccess() *bool
}

type UpdateOpenSourcePermissionResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 021788F6-E50C-4BD6-9F80-66B0A19A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateOpenSourcePermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateOpenSourcePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOpenSourcePermissionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateOpenSourcePermissionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateOpenSourcePermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateOpenSourcePermissionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateOpenSourcePermissionResponseBody) SetCode(v int32) *UpdateOpenSourcePermissionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateOpenSourcePermissionResponseBody) SetMessage(v string) *UpdateOpenSourcePermissionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateOpenSourcePermissionResponseBody) SetRequestId(v string) *UpdateOpenSourcePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOpenSourcePermissionResponseBody) SetSuccess(v bool) *UpdateOpenSourcePermissionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateOpenSourcePermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
