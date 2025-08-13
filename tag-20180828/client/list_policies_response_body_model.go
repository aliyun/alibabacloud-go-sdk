// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListPoliciesResponseBody
	GetNextToken() *string
	SetPolicyList(v []*ListPoliciesResponseBodyPolicyList) *ListPoliciesResponseBody
	GetPolicyList() []*ListPoliciesResponseBodyPolicyList
	SetRequestId(v string) *ListPoliciesResponseBody
	GetRequestId() *string
}

type ListPoliciesResponseBody struct {
	// Indicates whether the next query is required.
	//
	// 	- If the value of this parameter is empty (`"NextToken": ""`), all results are returned, and the next query is not required.
	//
	// 	- If the value of this parameter is not empty, the next query is required, and the value is the token used to start the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The tag policies.
	PolicyList []*ListPoliciesResponseBodyPolicyList `json:"PolicyList,omitempty" xml:"PolicyList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 72477CFF-5B24-5E30-9861-3DD9C4BD46E5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPoliciesResponseBody) GetPolicyList() []*ListPoliciesResponseBodyPolicyList {
	return s.PolicyList
}

func (s *ListPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPoliciesResponseBody) SetNextToken(v string) *ListPoliciesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPoliciesResponseBody) SetPolicyList(v []*ListPoliciesResponseBodyPolicyList) *ListPoliciesResponseBody {
	s.PolicyList = v
	return s
}

func (s *ListPoliciesResponseBody) SetRequestId(v string) *ListPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPoliciesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPoliciesResponseBodyPolicyList struct {
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
	// The ID of the tag policy.
	//
	// example:
	//
	// p-de62a0bf400e4b69****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the tag policy.
	//
	// example:
	//
	// example
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

func (s ListPoliciesResponseBodyPolicyList) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesResponseBodyPolicyList) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBodyPolicyList) GetPolicyContent() *string {
	return s.PolicyContent
}

func (s *ListPoliciesResponseBodyPolicyList) GetPolicyDesc() *string {
	return s.PolicyDesc
}

func (s *ListPoliciesResponseBodyPolicyList) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListPoliciesResponseBodyPolicyList) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListPoliciesResponseBodyPolicyList) GetUserType() *string {
	return s.UserType
}

func (s *ListPoliciesResponseBodyPolicyList) SetPolicyContent(v string) *ListPoliciesResponseBodyPolicyList {
	s.PolicyContent = &v
	return s
}

func (s *ListPoliciesResponseBodyPolicyList) SetPolicyDesc(v string) *ListPoliciesResponseBodyPolicyList {
	s.PolicyDesc = &v
	return s
}

func (s *ListPoliciesResponseBodyPolicyList) SetPolicyId(v string) *ListPoliciesResponseBodyPolicyList {
	s.PolicyId = &v
	return s
}

func (s *ListPoliciesResponseBodyPolicyList) SetPolicyName(v string) *ListPoliciesResponseBodyPolicyList {
	s.PolicyName = &v
	return s
}

func (s *ListPoliciesResponseBodyPolicyList) SetUserType(v string) *ListPoliciesResponseBodyPolicyList {
	s.UserType = &v
	return s
}

func (s *ListPoliciesResponseBodyPolicyList) Validate() error {
	return dara.Validate(s)
}
