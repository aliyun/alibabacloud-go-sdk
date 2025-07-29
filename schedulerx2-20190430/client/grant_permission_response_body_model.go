// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GrantPermissionResponseBody
	GetCode() *int32
	SetMessage(v string) *GrantPermissionResponseBody
	GetMessage() *string
	SetRequestId(v string) *GrantPermissionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GrantPermissionResponseBody
	GetSuccess() *bool
}

type GrantPermissionResponseBody struct {
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

func (s GrantPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GrantPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *GrantPermissionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GrantPermissionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GrantPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GrantPermissionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GrantPermissionResponseBody) SetCode(v int32) *GrantPermissionResponseBody {
	s.Code = &v
	return s
}

func (s *GrantPermissionResponseBody) SetMessage(v string) *GrantPermissionResponseBody {
	s.Message = &v
	return s
}

func (s *GrantPermissionResponseBody) SetRequestId(v string) *GrantPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantPermissionResponseBody) SetSuccess(v bool) *GrantPermissionResponseBody {
	s.Success = &v
	return s
}

func (s *GrantPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
