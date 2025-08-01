// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateAuthPolicyResponseBody
	GetCode() *int32
	SetData(v string) *UpdateAuthPolicyResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *UpdateAuthPolicyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateAuthPolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAuthPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAuthPolicyResponseBody
	GetSuccess() *bool
}

type UpdateAuthPolicyResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 500
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the data.
	//
	// example:
	//
	// {}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 83E3909D-D26F-5D97-B73B-407A26***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true: The request was successful. false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAuthPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAuthPolicyResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateAuthPolicyResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateAuthPolicyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateAuthPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAuthPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAuthPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAuthPolicyResponseBody) SetCode(v int32) *UpdateAuthPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAuthPolicyResponseBody) SetData(v string) *UpdateAuthPolicyResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateAuthPolicyResponseBody) SetHttpStatusCode(v int32) *UpdateAuthPolicyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateAuthPolicyResponseBody) SetMessage(v string) *UpdateAuthPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAuthPolicyResponseBody) SetRequestId(v string) *UpdateAuthPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAuthPolicyResponseBody) SetSuccess(v bool) *UpdateAuthPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAuthPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
