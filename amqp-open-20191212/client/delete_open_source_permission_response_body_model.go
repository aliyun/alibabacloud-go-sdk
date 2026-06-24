// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOpenSourcePermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteOpenSourcePermissionResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteOpenSourcePermissionResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteOpenSourcePermissionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteOpenSourcePermissionResponseBody
	GetSuccess() *bool
}

type DeleteOpenSourcePermissionResponseBody struct {
	// The return code. A value of 200 indicates success.
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
	// The request ID.
	//
	// example:
	//
	// 92385FD2-624A-48C9-8FB5-753F2AFA***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteOpenSourcePermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOpenSourcePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOpenSourcePermissionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteOpenSourcePermissionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteOpenSourcePermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOpenSourcePermissionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteOpenSourcePermissionResponseBody) SetCode(v int32) *DeleteOpenSourcePermissionResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteOpenSourcePermissionResponseBody) SetMessage(v string) *DeleteOpenSourcePermissionResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteOpenSourcePermissionResponseBody) SetRequestId(v string) *DeleteOpenSourcePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOpenSourcePermissionResponseBody) SetSuccess(v bool) *DeleteOpenSourcePermissionResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteOpenSourcePermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
