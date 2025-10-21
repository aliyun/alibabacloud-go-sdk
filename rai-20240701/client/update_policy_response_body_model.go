// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdatePolicyResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdatePolicyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdatePolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdatePolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdatePolicyResponseBody
	GetSuccess() *bool
}

type UpdatePolicyResponseBody struct {
	// Status code, 00000 indicates success; others indicate failure.
	//
	// example:
	//
	// 00000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Return message.
	//
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the operation was successful. true indicates success, false indicates failure.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdatePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdatePolicyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdatePolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdatePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdatePolicyResponseBody) SetCode(v string) *UpdatePolicyResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePolicyResponseBody) SetHttpStatusCode(v int32) *UpdatePolicyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdatePolicyResponseBody) SetMessage(v string) *UpdatePolicyResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePolicyResponseBody) SetRequestId(v string) *UpdatePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePolicyResponseBody) SetSuccess(v bool) *UpdatePolicyResponseBody {
	s.Success = &v
	return s
}

func (s *UpdatePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
