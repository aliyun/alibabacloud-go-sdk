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
	// The list of access policy names.
	//
	// > A maximum of 100 entries are returned.
	AppPolicyList []*ListAppPoliciesForIdentityResponseBodyAppPolicyList `json:"AppPolicyList,omitempty" xml:"AppPolicyList,omitempty" type:"Repeated"`
	// The request ID.
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
	if s.AppPolicyList != nil {
		for _, item := range s.AppPolicyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAppPoliciesForIdentityResponseBodyAppPolicyList struct {
	// The application ID.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The time when the application policy was granted to the role. Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// example:
	//
	// 2019-01-01T01:01:01Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The policy description.
	//
	// example:
	//
	// App full access permission
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the application policy granted to the role was last modified. Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// example:
	//
	// 2019-01-01T01:08:01Z
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// The policy name.
	//
	// example:
	//
	// VODAppFullAccess
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The policy type. Valid values:
	//
	// - **System**: system policy.
	//
	// - **Custom**: user-defined policy.
	//
	// example:
	//
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The policy value.
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
