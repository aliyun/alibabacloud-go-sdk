// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPermissionPoliciesInAccessConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPermissionPolicies(v []*ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies) *ListPermissionPoliciesInAccessConfigurationResponseBody
	GetPermissionPolicies() []*ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies
	SetRequestId(v string) *ListPermissionPoliciesInAccessConfigurationResponseBody
	GetRequestId() *string
	SetTotalCounts(v int32) *ListPermissionPoliciesInAccessConfigurationResponseBody
	GetTotalCounts() *int32
}

type ListPermissionPoliciesInAccessConfigurationResponseBody struct {
	// The policies.
	PermissionPolicies []*ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies `json:"PermissionPolicies,omitempty" xml:"PermissionPolicies,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 3A5E771F-1F5A-5555-A64E-579748AAFD98
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of policies.
	//
	// example:
	//
	// 2
	TotalCounts *int32 `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
}

func (s ListPermissionPoliciesInAccessConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionPoliciesInAccessConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *ListPermissionPoliciesInAccessConfigurationResponseBody) GetPermissionPolicies() []*ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies {
	return s.PermissionPolicies
}

func (s *ListPermissionPoliciesInAccessConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPermissionPoliciesInAccessConfigurationResponseBody) GetTotalCounts() *int32 {
	return s.TotalCounts
}

func (s *ListPermissionPoliciesInAccessConfigurationResponseBody) SetPermissionPolicies(v []*ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies) *ListPermissionPoliciesInAccessConfigurationResponseBody {
	s.PermissionPolicies = v
	return s
}

func (s *ListPermissionPoliciesInAccessConfigurationResponseBody) SetRequestId(v string) *ListPermissionPoliciesInAccessConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPermissionPoliciesInAccessConfigurationResponseBody) SetTotalCounts(v int32) *ListPermissionPoliciesInAccessConfigurationResponseBody {
	s.TotalCounts = &v
	return s
}

func (s *ListPermissionPoliciesInAccessConfigurationResponseBody) Validate() error {
	if s.PermissionPolicies != nil {
		for _, item := range s.PermissionPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies struct {
	// The time when the policy was created for the access configuration.
	//
	// example:
	//
	// 2021-11-03T06:37:25Z
	AddTime *string `json:"AddTime,omitempty" xml:"AddTime,omitempty"`
	// The configurations of the inline policy.
	//
	// >  This parameter is returned only when the value of the PermissionPolicyType parameter is Inline.
	//
	// example:
	//
	// {\\"Statement\\": [{\\"Action\\": \\"*\\",\\"Effect\\": \\"Allow\\",\\"Resource\\": \\"*\\"}],\\"Version\\": \\"1\\"}
	PermissionPolicyDocument *string `json:"PermissionPolicyDocument,omitempty" xml:"PermissionPolicyDocument,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// AliyunECSFullAccess
	PermissionPolicyName *string `json:"PermissionPolicyName,omitempty" xml:"PermissionPolicyName,omitempty"`
	// The type of the policy.
	//
	// example:
	//
	// System
	PermissionPolicyType *string `json:"PermissionPolicyType,omitempty" xml:"PermissionPolicyType,omitempty"`
}

func (s ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies) GoString() string {
	return s.String()
}

func (s *ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies) GetAddTime() *string {
	return s.AddTime
}

func (s *ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies) GetPermissionPolicyDocument() *string {
	return s.PermissionPolicyDocument
}

func (s *ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies) GetPermissionPolicyName() *string {
	return s.PermissionPolicyName
}

func (s *ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies) GetPermissionPolicyType() *string {
	return s.PermissionPolicyType
}

func (s *ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies) SetAddTime(v string) *ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies {
	s.AddTime = &v
	return s
}

func (s *ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies) SetPermissionPolicyDocument(v string) *ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies {
	s.PermissionPolicyDocument = &v
	return s
}

func (s *ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies) SetPermissionPolicyName(v string) *ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies {
	s.PermissionPolicyName = &v
	return s
}

func (s *ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies) SetPermissionPolicyType(v string) *ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies {
	s.PermissionPolicyType = &v
	return s
}

func (s *ListPermissionPoliciesInAccessConfigurationResponseBodyPermissionPolicies) Validate() error {
	return dara.Validate(s)
}
