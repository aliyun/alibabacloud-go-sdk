// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokePermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *RevokePermissionResponseBody
	GetCode() *int32
	SetMessage(v string) *RevokePermissionResponseBody
	GetMessage() *string
	SetRequestId(v string) *RevokePermissionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RevokePermissionResponseBody
	GetSuccess() *bool
}

type RevokePermissionResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 400
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message that is returned only if the corresponding error occurs.
	//
	// example:
	//
	// Your request is denied as lack of ssl protect.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RevokePermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *RevokePermissionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *RevokePermissionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RevokePermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokePermissionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RevokePermissionResponseBody) SetCode(v int32) *RevokePermissionResponseBody {
	s.Code = &v
	return s
}

func (s *RevokePermissionResponseBody) SetMessage(v string) *RevokePermissionResponseBody {
	s.Message = &v
	return s
}

func (s *RevokePermissionResponseBody) SetRequestId(v string) *RevokePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokePermissionResponseBody) SetSuccess(v bool) *RevokePermissionResponseBody {
	s.Success = &v
	return s
}

func (s *RevokePermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
