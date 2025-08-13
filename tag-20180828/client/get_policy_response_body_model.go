// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicy(v *GetPolicyResponseBodyPolicy) *GetPolicyResponseBody
	GetPolicy() *GetPolicyResponseBodyPolicy
	SetRequestId(v string) *GetPolicyResponseBody
	GetRequestId() *string
}

type GetPolicyResponseBody struct {
	// The details of the tag policy.
	Policy *GetPolicyResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 1E7A4802-AB57-570A-9860-F15B60E1586B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetPolicyResponseBody) GetPolicy() *GetPolicyResponseBodyPolicy {
	return s.Policy
}

func (s *GetPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPolicyResponseBody) SetPolicy(v *GetPolicyResponseBodyPolicy) *GetPolicyResponseBody {
	s.Policy = v
	return s
}

func (s *GetPolicyResponseBody) SetRequestId(v string) *GetPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetPolicyResponseBodyPolicy struct {
	// The document of the tag policy.
	//
	// example:
	//
	// {\\"tags\\":{\\"CostCenter\\":{\\"tag_value\\":{\\"@@assign\\":[\\"Beijing\\",\\"Shanghai\\"]},\\"tag_key\\":{\\"@@assign\\":\\"CostCenter\\"}}}}
	PolicyContent *string `json:"PolicyContent,omitempty" xml:"PolicyContent,omitempty"`
	// The description of the tag policy.
	//
	// example:
	//
	// This is a tag policy example.
	PolicyDesc *string `json:"PolicyDesc,omitempty" xml:"PolicyDesc,omitempty"`
	// The name of the tag policy.
	//
	// example:
	//
	// test
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The mode of the Tag Policy feature. Valid values:
	//
	// 	- USER: single-account mode
	//
	// 	- RD: multi-account mode
	//
	// For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](https://help.aliyun.com/document_detail/417434.html).
	//
	// example:
	//
	// USER
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s GetPolicyResponseBodyPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyResponseBodyPolicy) GoString() string {
	return s.String()
}

func (s *GetPolicyResponseBodyPolicy) GetPolicyContent() *string {
	return s.PolicyContent
}

func (s *GetPolicyResponseBodyPolicy) GetPolicyDesc() *string {
	return s.PolicyDesc
}

func (s *GetPolicyResponseBodyPolicy) GetPolicyName() *string {
	return s.PolicyName
}

func (s *GetPolicyResponseBodyPolicy) GetUserType() *string {
	return s.UserType
}

func (s *GetPolicyResponseBodyPolicy) SetPolicyContent(v string) *GetPolicyResponseBodyPolicy {
	s.PolicyContent = &v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetPolicyDesc(v string) *GetPolicyResponseBodyPolicy {
	s.PolicyDesc = &v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetPolicyName(v string) *GetPolicyResponseBodyPolicy {
	s.PolicyName = &v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetUserType(v string) *GetPolicyResponseBodyPolicy {
	s.UserType = &v
	return s
}

func (s *GetPolicyResponseBodyPolicy) Validate() error {
	return dara.Validate(s)
}
