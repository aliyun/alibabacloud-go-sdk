// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAuthPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddAuthPolicyResponseBody
	GetCode() *int32
	SetData(v string) *AddAuthPolicyResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *AddAuthPolicyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddAuthPolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddAuthPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddAuthPolicyResponseBody
	GetSuccess() *bool
}

type AddAuthPolicyResponseBody struct {
	// example:
	//
	// 200
	Code *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// dc63-465d-8ef5-20dc18af****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddAuthPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddAuthPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *AddAuthPolicyResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddAuthPolicyResponseBody) GetData() *string {
	return s.Data
}

func (s *AddAuthPolicyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddAuthPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddAuthPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddAuthPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddAuthPolicyResponseBody) SetCode(v int32) *AddAuthPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *AddAuthPolicyResponseBody) SetData(v string) *AddAuthPolicyResponseBody {
	s.Data = &v
	return s
}

func (s *AddAuthPolicyResponseBody) SetHttpStatusCode(v int32) *AddAuthPolicyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddAuthPolicyResponseBody) SetMessage(v string) *AddAuthPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *AddAuthPolicyResponseBody) SetRequestId(v string) *AddAuthPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAuthPolicyResponseBody) SetSuccess(v bool) *AddAuthPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *AddAuthPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
