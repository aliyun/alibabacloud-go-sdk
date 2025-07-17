// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePolicyResponseBody
	GetCode() *string
	SetData(v *CreatePolicyResponseBodyData) *CreatePolicyResponseBody
	GetData() *CreatePolicyResponseBodyData
	SetMessage(v string) *CreatePolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreatePolicyResponseBody
	GetRequestId() *string
}

type CreatePolicyResponseBody struct {
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Response data.
	Data *CreatePolicyResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E7406754***
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreatePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePolicyResponseBody) GetData() *CreatePolicyResponseBodyData {
	return s.Data
}

func (s *CreatePolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePolicyResponseBody) SetCode(v string) *CreatePolicyResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePolicyResponseBody) SetData(v *CreatePolicyResponseBodyData) *CreatePolicyResponseBody {
	s.Data = v
	return s
}

func (s *CreatePolicyResponseBody) SetMessage(v string) *CreatePolicyResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePolicyResponseBody) SetRequestId(v string) *CreatePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreatePolicyResponseBodyData struct {
	// Policy ID
	//
	// example:
	//
	// p-cq7l5s5lhtgi6qasr***
	PolicyId *string `json:"policyId,omitempty" xml:"policyId,omitempty"`
}

func (s CreatePolicyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePolicyResponseBodyData) GetPolicyId() *string {
	return s.PolicyId
}

func (s *CreatePolicyResponseBodyData) SetPolicyId(v string) *CreatePolicyResponseBodyData {
	s.PolicyId = &v
	return s
}

func (s *CreatePolicyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
