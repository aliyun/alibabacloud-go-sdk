// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicesForUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListPolicesForUserGroupResponseBody
	GetRequestId() *string
	SetUserGroups(v []*ListPolicesForUserGroupResponseBodyUserGroups) *ListPolicesForUserGroupResponseBody
	GetUserGroups() []*ListPolicesForUserGroupResponseBodyUserGroups
}

type ListPolicesForUserGroupResponseBody struct {
	// example:
	//
	// 5F04DFBD-3F48-5F70-AE72-474026670128
	RequestId  *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserGroups []*ListPolicesForUserGroupResponseBodyUserGroups `json:"UserGroups,omitempty" xml:"UserGroups,omitempty" type:"Repeated"`
}

func (s ListPolicesForUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPolicesForUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListPolicesForUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPolicesForUserGroupResponseBody) GetUserGroups() []*ListPolicesForUserGroupResponseBodyUserGroups {
	return s.UserGroups
}

func (s *ListPolicesForUserGroupResponseBody) SetRequestId(v string) *ListPolicesForUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPolicesForUserGroupResponseBody) SetUserGroups(v []*ListPolicesForUserGroupResponseBodyUserGroups) *ListPolicesForUserGroupResponseBody {
	s.UserGroups = v
	return s
}

func (s *ListPolicesForUserGroupResponseBody) Validate() error {
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

type ListPolicesForUserGroupResponseBodyUserGroups struct {
	Polices []*ListPolicesForUserGroupResponseBodyUserGroupsPolices `json:"Polices,omitempty" xml:"Polices,omitempty" type:"Repeated"`
	// example:
	//
	// usergroup-6f1ef2fc56b6****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s ListPolicesForUserGroupResponseBodyUserGroups) String() string {
	return dara.Prettify(s)
}

func (s ListPolicesForUserGroupResponseBodyUserGroups) GoString() string {
	return s.String()
}

func (s *ListPolicesForUserGroupResponseBodyUserGroups) GetPolices() []*ListPolicesForUserGroupResponseBodyUserGroupsPolices {
	return s.Polices
}

func (s *ListPolicesForUserGroupResponseBodyUserGroups) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *ListPolicesForUserGroupResponseBodyUserGroups) SetPolices(v []*ListPolicesForUserGroupResponseBodyUserGroupsPolices) *ListPolicesForUserGroupResponseBodyUserGroups {
	s.Polices = v
	return s
}

func (s *ListPolicesForUserGroupResponseBodyUserGroups) SetUserGroupId(v string) *ListPolicesForUserGroupResponseBodyUserGroups {
	s.UserGroupId = &v
	return s
}

func (s *ListPolicesForUserGroupResponseBodyUserGroups) Validate() error {
	if s.Polices != nil {
		for _, item := range s.Polices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPolicesForUserGroupResponseBodyUserGroupsPolices struct {
	// example:
	//
	// private_access_policy_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// pa-policy-ce2bf7236fab****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// PrivateAccess
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ListPolicesForUserGroupResponseBodyUserGroupsPolices) String() string {
	return dara.Prettify(s)
}

func (s ListPolicesForUserGroupResponseBodyUserGroupsPolices) GoString() string {
	return s.String()
}

func (s *ListPolicesForUserGroupResponseBodyUserGroupsPolices) GetName() *string {
	return s.Name
}

func (s *ListPolicesForUserGroupResponseBodyUserGroupsPolices) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListPolicesForUserGroupResponseBodyUserGroupsPolices) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListPolicesForUserGroupResponseBodyUserGroupsPolices) SetName(v string) *ListPolicesForUserGroupResponseBodyUserGroupsPolices {
	s.Name = &v
	return s
}

func (s *ListPolicesForUserGroupResponseBodyUserGroupsPolices) SetPolicyId(v string) *ListPolicesForUserGroupResponseBodyUserGroupsPolices {
	s.PolicyId = &v
	return s
}

func (s *ListPolicesForUserGroupResponseBodyUserGroupsPolices) SetPolicyType(v string) *ListPolicesForUserGroupResponseBodyUserGroupsPolices {
	s.PolicyType = &v
	return s
}

func (s *ListPolicesForUserGroupResponseBodyUserGroupsPolices) Validate() error {
	return dara.Validate(s)
}
