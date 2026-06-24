// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpenSourcePermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateOpenSourcePermissionResponseBody
	GetCode() *int32
	SetMessage(v string) *CreateOpenSourcePermissionResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateOpenSourcePermissionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateOpenSourcePermissionResponseBody
	GetSuccess() *bool
}

type CreateOpenSourcePermissionResponseBody struct {
	// The return code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8BFB1C9D-08A2-4859-A47C-403C9EFA2***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateOpenSourcePermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOpenSourcePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOpenSourcePermissionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateOpenSourcePermissionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateOpenSourcePermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOpenSourcePermissionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateOpenSourcePermissionResponseBody) SetCode(v int32) *CreateOpenSourcePermissionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateOpenSourcePermissionResponseBody) SetMessage(v string) *CreateOpenSourcePermissionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateOpenSourcePermissionResponseBody) SetRequestId(v string) *CreateOpenSourcePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOpenSourcePermissionResponseBody) SetSuccess(v bool) *CreateOpenSourcePermissionResponseBody {
	s.Success = &v
	return s
}

func (s *CreateOpenSourcePermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
