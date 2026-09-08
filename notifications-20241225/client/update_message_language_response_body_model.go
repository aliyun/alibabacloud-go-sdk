// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMessageLanguageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateMessageLanguageResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateMessageLanguageResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateMessageLanguageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMessageLanguageResponseBody
	GetSuccess() *bool
}

type UpdateMessageLanguageResponseBody struct {
	// The error code returned if the call failed. For more information, see error codes.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 73FD6AE8-898F-5D09-9763-69B8A875488A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Valid values: true: The operation was successful. false: The operation failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMessageLanguageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMessageLanguageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMessageLanguageResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateMessageLanguageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateMessageLanguageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMessageLanguageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMessageLanguageResponseBody) SetCode(v string) *UpdateMessageLanguageResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMessageLanguageResponseBody) SetMessage(v string) *UpdateMessageLanguageResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMessageLanguageResponseBody) SetRequestId(v string) *UpdateMessageLanguageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMessageLanguageResponseBody) SetSuccess(v bool) *UpdateMessageLanguageResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMessageLanguageResponseBody) Validate() error {
	return dara.Validate(s)
}
