// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppPoliciesForIdentityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppPolicyList(v []*ListAppPoliciesForIdentityResponseBodyAppPolicyList) *ListAppPoliciesForIdentityResponseBody
	GetAppPolicyList() []*ListAppPoliciesForIdentityResponseBodyAppPolicyList
	SetRequestId(v string) *ListAppPoliciesForIdentityResponseBody
	GetRequestId() *string
}

type ListAppPoliciesForIdentityResponseBody struct {
	// The details of each policy.
	//
	// > A maximum of 100 entries can be returned.
	AppPolicyList []*ListAppPoliciesForIdentityResponseBodyAppPolicyList `json:"AppPolicyList,omitempty" xml:"AppPolicyList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// C9F3E715-B3B8-4D*****27-3A70346F0E04
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAppPoliciesForIdentityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppPoliciesForIdentityResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppPoliciesForIdentityResponseBody) GetAppPolicyList() []*ListAppPoliciesForIdentityResponseBodyAppPolicyList {
	return s.AppPolicyList
}

func (s *ListAppPoliciesForIdentityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppPoliciesForIdentityResponseBody) SetAppPolicyList(v []*ListAppPoliciesForIdentityResponseBodyAppPolicyList) *ListAppPoliciesForIdentityResponseBody {
	s.AppPolicyList = v
	return s
}

func (s *ListAppPoliciesForIdentityResponseBody) SetRequestId(v string) *ListAppPoliciesForIdentityResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppPoliciesForIdentityResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAppPoliciesForIdentityResponseBodyAppPolicyList struct {
	// The ID of the application.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The time when the application policy was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-01-01T01:01:01Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the policy.
	//
	// example:
	//
	// App full access permission
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The last time when the application policy was modified. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-01-01T01:08:01Z
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// VODAppFullAccess
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the policy. Valid values:
	//
	// 	- **System**
	//
	// 	- **Custom**
	//
	// example:
	//
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The content of the policy.
	//
	// example:
	//
	// ****
	PolicyValue *string `json:"PolicyValue,omitempty" xml:"PolicyValue,omitempty"`
}

func (s ListAppPoliciesForIdentityResponseBodyAppPolicyList) String() string {
	return dara.Prettify(s)
}

func (s ListAppPoliciesForIdentityResponseBodyAppPolicyList) GoString() string {
	return s.String()
}

func (s *ListAppPoliciesForIdentityResponseBodyAppPolicyList) GetAppId() *string {
	return s.AppId
}

func (s *ListAppPoliciesForIdentityResponseBodyAppPolicyList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListAppPoliciesForIdentityResponseBodyAppPolicyList) GetDescription() *string {
	return s.Description
}

func (s *ListAppPoliciesForIdentityResponseBodyAppPolicyList) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *ListAppPoliciesForIdentityResponseBodyAppPolicyList) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListAppPoliciesForIdentityResponseBodyAppPolicyList) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListAppPoliciesForIdentityResponseBodyAppPolicyList) GetPolicyValue() *string {
	return s.PolicyValue
}

func (s *ListAppPoliciesForIdentityResponseBodyAppPolicyList) SetAppId(v string) *ListAppPoliciesForIdentityResponseBodyAppPolicyList {
	s.AppId = &v
	return s
}

func (s *ListAppPoliciesForIdentityResponseBodyAppPolicyList) SetCreationTime(v string) *ListAppPoliciesForIdentityResponseBodyAppPolicyList {
	s.CreationTime = &v
	return s
}

func (s *ListAppPoliciesForIdentityResponseBodyAppPolicyList) SetDescription(v string) *ListAppPoliciesForIdentityResponseBodyAppPolicyList {
	s.Description = &v
	return s
}

func (s *ListAppPoliciesForIdentityResponseBodyAppPolicyList) SetModificationTime(v string) *ListAppPoliciesForIdentityResponseBodyAppPolicyList {
	s.ModificationTime = &v
	return s
}

func (s *ListAppPoliciesForIdentityResponseBodyAppPolicyList) SetPolicyName(v string) *ListAppPoliciesForIdentityResponseBodyAppPolicyList {
	s.PolicyName = &v
	return s
}

func (s *ListAppPoliciesForIdentityResponseBodyAppPolicyList) SetPolicyType(v string) *ListAppPoliciesForIdentityResponseBodyAppPolicyList {
	s.PolicyType = &v
	return s
}

func (s *ListAppPoliciesForIdentityResponseBodyAppPolicyList) SetPolicyValue(v string) *ListAppPoliciesForIdentityResponseBodyAppPolicyList {
	s.PolicyValue = &v
	return s
}

func (s *ListAppPoliciesForIdentityResponseBodyAppPolicyList) Validate() error {
	return dara.Validate(s)
}
