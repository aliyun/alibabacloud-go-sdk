// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomAuthIdentityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteCustomAuthIdentityResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteCustomAuthIdentityResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCustomAuthIdentityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCustomAuthIdentityResponseBody
	GetSuccess() *bool
}

type DeleteCustomAuthIdentityResponseBody struct {
	// The status code. A value of 200 indicates success, and other status codes indicate errors.
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
	// 11568B5B-13A8-4E72-9DBA-3A14F7D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values:
	//
	// - true: The operation is successful.
	//
	// - false: The operation fails.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCustomAuthIdentityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomAuthIdentityResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomAuthIdentityResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteCustomAuthIdentityResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCustomAuthIdentityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomAuthIdentityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCustomAuthIdentityResponseBody) SetCode(v int32) *DeleteCustomAuthIdentityResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCustomAuthIdentityResponseBody) SetMessage(v string) *DeleteCustomAuthIdentityResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCustomAuthIdentityResponseBody) SetRequestId(v string) *DeleteCustomAuthIdentityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomAuthIdentityResponseBody) SetSuccess(v bool) *DeleteCustomAuthIdentityResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCustomAuthIdentityResponseBody) Validate() error {
	return dara.Validate(s)
}
