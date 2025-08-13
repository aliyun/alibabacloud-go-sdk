// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoliciesForTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListPoliciesForTargetResponseBodyData) *ListPoliciesForTargetResponseBody
	GetData() []*ListPoliciesForTargetResponseBodyData
	SetNextToken(v string) *ListPoliciesForTargetResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListPoliciesForTargetResponseBody
	GetRequestId() *string
}

type ListPoliciesForTargetResponseBody struct {
	// The tag policies that are attached to the object.
	Data []*ListPoliciesForTargetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// The ID of the request.
	//
	// example:
	//
	// 8C962146-AB38-516C-818C-695D4E9F2EA2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPoliciesForTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesForTargetResponseBody) GoString() string {
	return s.String()
}

func (s *ListPoliciesForTargetResponseBody) GetData() []*ListPoliciesForTargetResponseBodyData {
	return s.Data
}

func (s *ListPoliciesForTargetResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPoliciesForTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPoliciesForTargetResponseBody) SetData(v []*ListPoliciesForTargetResponseBodyData) *ListPoliciesForTargetResponseBody {
	s.Data = v
	return s
}

func (s *ListPoliciesForTargetResponseBody) SetNextToken(v string) *ListPoliciesForTargetResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPoliciesForTargetResponseBody) SetRequestId(v string) *ListPoliciesForTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPoliciesForTargetResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPoliciesForTargetResponseBodyData struct {
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

func (s ListPoliciesForTargetResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesForTargetResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPoliciesForTargetResponseBodyData) GetPolicyContent() *string {
	return s.PolicyContent
}

func (s *ListPoliciesForTargetResponseBodyData) GetPolicyDesc() *string {
	return s.PolicyDesc
}

func (s *ListPoliciesForTargetResponseBodyData) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListPoliciesForTargetResponseBodyData) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListPoliciesForTargetResponseBodyData) GetUserType() *string {
	return s.UserType
}

func (s *ListPoliciesForTargetResponseBodyData) SetPolicyContent(v string) *ListPoliciesForTargetResponseBodyData {
	s.PolicyContent = &v
	return s
}

func (s *ListPoliciesForTargetResponseBodyData) SetPolicyDesc(v string) *ListPoliciesForTargetResponseBodyData {
	s.PolicyDesc = &v
	return s
}

func (s *ListPoliciesForTargetResponseBodyData) SetPolicyId(v string) *ListPoliciesForTargetResponseBodyData {
	s.PolicyId = &v
	return s
}

func (s *ListPoliciesForTargetResponseBodyData) SetPolicyName(v string) *ListPoliciesForTargetResponseBodyData {
	s.PolicyName = &v
	return s
}

func (s *ListPoliciesForTargetResponseBodyData) SetUserType(v string) *ListPoliciesForTargetResponseBodyData {
	s.UserType = &v
	return s
}

func (s *ListPoliciesForTargetResponseBodyData) Validate() error {
	return dara.Validate(s)
}
