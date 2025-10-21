// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeletePolicyResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeletePolicyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeletePolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeletePolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeletePolicyResponseBody
	GetSuccess() *bool
}

type DeletePolicyResponseBody struct {
	// Status code, 00000 indicates success; others indicate failure.
	//
	// example:
	//
	// 00000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code description
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Error message
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
	// Indicates whether the operation was successful. true means success, false means failure.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeletePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeletePolicyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeletePolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeletePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeletePolicyResponseBody) SetCode(v string) *DeletePolicyResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePolicyResponseBody) SetHttpStatusCode(v int32) *DeletePolicyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeletePolicyResponseBody) SetMessage(v string) *DeletePolicyResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePolicyResponseBody) SetRequestId(v string) *DeletePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePolicyResponseBody) SetSuccess(v bool) *DeletePolicyResponseBody {
	s.Success = &v
	return s
}

func (s *DeletePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
