// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolicyDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicy(v *UpdatePolicyDescriptionResponseBodyPolicy) *UpdatePolicyDescriptionResponseBody
	GetPolicy() *UpdatePolicyDescriptionResponseBodyPolicy
	SetRequestId(v string) *UpdatePolicyDescriptionResponseBody
	GetRequestId() *string
}

type UpdatePolicyDescriptionResponseBody struct {
	// The information about the policy.
	Policy *UpdatePolicyDescriptionResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7572DEBD-0ECE-518E-8682-D8CB82F8FE8E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePolicyDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePolicyDescriptionResponseBody) GetPolicy() *UpdatePolicyDescriptionResponseBodyPolicy {
	return s.Policy
}

func (s *UpdatePolicyDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePolicyDescriptionResponseBody) SetPolicy(v *UpdatePolicyDescriptionResponseBodyPolicy) *UpdatePolicyDescriptionResponseBody {
	s.Policy = v
	return s
}

func (s *UpdatePolicyDescriptionResponseBody) SetRequestId(v string) *UpdatePolicyDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePolicyDescriptionResponseBody) Validate() error {
	if s.Policy != nil {
		if err := s.Policy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePolicyDescriptionResponseBodyPolicy struct {
	// The time when the policy was created.
	//
	// example:
	//
	// 2022-02-28T07:04:15Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The version of the policy. Default value: v1.
	//
	// example:
	//
	// v1
	DefaultVersion *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	// The description of the policy.
	//
	// example:
	//
	// This is a test policy.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// TestPolicy
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the policy. Valid values:
	//
	// 	- Custom
	//
	// 	- System
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The time when the policy was modified.
	//
	// example:
	//
	// 2022-02-28T07:05:37Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s UpdatePolicyDescriptionResponseBodyPolicy) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyDescriptionResponseBodyPolicy) GoString() string {
	return s.String()
}

func (s *UpdatePolicyDescriptionResponseBodyPolicy) GetCreateDate() *string {
	return s.CreateDate
}

func (s *UpdatePolicyDescriptionResponseBodyPolicy) GetDefaultVersion() *string {
	return s.DefaultVersion
}

func (s *UpdatePolicyDescriptionResponseBodyPolicy) GetDescription() *string {
	return s.Description
}

func (s *UpdatePolicyDescriptionResponseBodyPolicy) GetPolicyName() *string {
	return s.PolicyName
}

func (s *UpdatePolicyDescriptionResponseBodyPolicy) GetPolicyType() *string {
	return s.PolicyType
}

func (s *UpdatePolicyDescriptionResponseBodyPolicy) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *UpdatePolicyDescriptionResponseBodyPolicy) SetCreateDate(v string) *UpdatePolicyDescriptionResponseBodyPolicy {
	s.CreateDate = &v
	return s
}

func (s *UpdatePolicyDescriptionResponseBodyPolicy) SetDefaultVersion(v string) *UpdatePolicyDescriptionResponseBodyPolicy {
	s.DefaultVersion = &v
	return s
}

func (s *UpdatePolicyDescriptionResponseBodyPolicy) SetDescription(v string) *UpdatePolicyDescriptionResponseBodyPolicy {
	s.Description = &v
	return s
}

func (s *UpdatePolicyDescriptionResponseBodyPolicy) SetPolicyName(v string) *UpdatePolicyDescriptionResponseBodyPolicy {
	s.PolicyName = &v
	return s
}

func (s *UpdatePolicyDescriptionResponseBodyPolicy) SetPolicyType(v string) *UpdatePolicyDescriptionResponseBodyPolicy {
	s.PolicyType = &v
	return s
}

func (s *UpdatePolicyDescriptionResponseBodyPolicy) SetUpdateDate(v string) *UpdatePolicyDescriptionResponseBodyPolicy {
	s.UpdateDate = &v
	return s
}

func (s *UpdatePolicyDescriptionResponseBodyPolicy) Validate() error {
	return dara.Validate(s)
}
