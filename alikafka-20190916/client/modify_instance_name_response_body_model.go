// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ModifyInstanceNameResponseBody
	GetCode() *int32
	SetMessage(v string) *ModifyInstanceNameResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyInstanceNameResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyInstanceNameResponseBody
	GetSuccess() *bool
}

type ModifyInstanceNameResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the call is successful.
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
	// The request ID.
	//
	// example:
	//
	// 06084011-E093-46F3-A51F-4B19A8AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyInstanceNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceNameResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNameResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ModifyInstanceNameResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyInstanceNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyInstanceNameResponseBody) SetCode(v int32) *ModifyInstanceNameResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyInstanceNameResponseBody) SetMessage(v string) *ModifyInstanceNameResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyInstanceNameResponseBody) SetRequestId(v string) *ModifyInstanceNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceNameResponseBody) SetSuccess(v bool) *ModifyInstanceNameResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyInstanceNameResponseBody) Validate() error {
	return dara.Validate(s)
}
