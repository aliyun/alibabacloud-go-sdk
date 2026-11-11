// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserLoginPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicies(v []*DescribeUserLoginPoliciesResponseBodyPolicies) *DescribeUserLoginPoliciesResponseBody
	GetPolicies() []*DescribeUserLoginPoliciesResponseBodyPolicies
	SetRequestId(v string) *DescribeUserLoginPoliciesResponseBody
	GetRequestId() *string
}

type DescribeUserLoginPoliciesResponseBody struct {
	// The list of policies.
	Policies []*DescribeUserLoginPoliciesResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserLoginPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserLoginPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserLoginPoliciesResponseBody) GetPolicies() []*DescribeUserLoginPoliciesResponseBodyPolicies {
	return s.Policies
}

func (s *DescribeUserLoginPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserLoginPoliciesResponseBody) SetPolicies(v []*DescribeUserLoginPoliciesResponseBodyPolicies) *DescribeUserLoginPoliciesResponseBody {
	s.Policies = v
	return s
}

func (s *DescribeUserLoginPoliciesResponseBody) SetRequestId(v string) *DescribeUserLoginPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserLoginPoliciesResponseBody) Validate() error {
	if s.Policies != nil {
		for _, item := range s.Policies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUserLoginPoliciesResponseBodyPolicies struct {
	// The policy name.
	//
	// example:
	//
	// deveoper-policy
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// up-89unlabdl****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The associated user group information.
	UserGroups []*DescribeUserLoginPoliciesResponseBodyPoliciesUserGroups `json:"UserGroups,omitempty" xml:"UserGroups,omitempty" type:"Repeated"`
}

func (s DescribeUserLoginPoliciesResponseBodyPolicies) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserLoginPoliciesResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *DescribeUserLoginPoliciesResponseBodyPolicies) GetName() *string {
	return s.Name
}

func (s *DescribeUserLoginPoliciesResponseBodyPolicies) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribeUserLoginPoliciesResponseBodyPolicies) GetUserGroups() []*DescribeUserLoginPoliciesResponseBodyPoliciesUserGroups {
	return s.UserGroups
}

func (s *DescribeUserLoginPoliciesResponseBodyPolicies) SetName(v string) *DescribeUserLoginPoliciesResponseBodyPolicies {
	s.Name = &v
	return s
}

func (s *DescribeUserLoginPoliciesResponseBodyPolicies) SetPolicyId(v string) *DescribeUserLoginPoliciesResponseBodyPolicies {
	s.PolicyId = &v
	return s
}

func (s *DescribeUserLoginPoliciesResponseBodyPolicies) SetUserGroups(v []*DescribeUserLoginPoliciesResponseBodyPoliciesUserGroups) *DescribeUserLoginPoliciesResponseBodyPolicies {
	s.UserGroups = v
	return s
}

func (s *DescribeUserLoginPoliciesResponseBodyPolicies) Validate() error {
	if s.UserGroups != nil {
		for _, item := range s.UserGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUserLoginPoliciesResponseBodyPoliciesUserGroups struct {
	// The user group ID.
	//
	// example:
	//
	// ug-ilu90mhsduf****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The user group name.
	//
	// example:
	//
	// Developers
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s DescribeUserLoginPoliciesResponseBodyPoliciesUserGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserLoginPoliciesResponseBodyPoliciesUserGroups) GoString() string {
	return s.String()
}

func (s *DescribeUserLoginPoliciesResponseBodyPoliciesUserGroups) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeUserLoginPoliciesResponseBodyPoliciesUserGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeUserLoginPoliciesResponseBodyPoliciesUserGroups) SetGroupId(v string) *DescribeUserLoginPoliciesResponseBodyPoliciesUserGroups {
	s.GroupId = &v
	return s
}

func (s *DescribeUserLoginPoliciesResponseBodyPoliciesUserGroups) SetGroupName(v string) *DescribeUserLoginPoliciesResponseBodyPoliciesUserGroups {
	s.GroupName = &v
	return s
}

func (s *DescribeUserLoginPoliciesResponseBodyPoliciesUserGroups) Validate() error {
	return dara.Validate(s)
}
