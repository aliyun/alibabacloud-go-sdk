// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePolicyV2ResponseBody
	GetCode() *string
	SetMessage(v string) *CreatePolicyV2ResponseBody
	GetMessage() *string
	SetPolicyId(v string) *CreatePolicyV2ResponseBody
	GetPolicyId() *string
	SetRequestId(v string) *CreatePolicyV2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreatePolicyV2ResponseBody
	GetSuccess() *bool
}

type CreatePolicyV2ResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned. If the call is successful, "successful" is returned. If the call fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the backup policy.
	//
	// example:
	//
	// po-000000zemnuyx2li3y9y
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// EEC65C22-2152-5E31-8AD6-D6CBF1BFF49F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- true: The call is successful.
	//
	// 	- false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePolicyV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyV2ResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolicyV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePolicyV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePolicyV2ResponseBody) GetPolicyId() *string {
	return s.PolicyId
}

func (s *CreatePolicyV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePolicyV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreatePolicyV2ResponseBody) SetCode(v string) *CreatePolicyV2ResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePolicyV2ResponseBody) SetMessage(v string) *CreatePolicyV2ResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePolicyV2ResponseBody) SetPolicyId(v string) *CreatePolicyV2ResponseBody {
	s.PolicyId = &v
	return s
}

func (s *CreatePolicyV2ResponseBody) SetRequestId(v string) *CreatePolicyV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePolicyV2ResponseBody) SetSuccess(v bool) *CreatePolicyV2ResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePolicyV2ResponseBody) Validate() error {
	return dara.Validate(s)
}
