// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicy(v *CreatePolicyResponseBodyPolicy) *CreatePolicyResponseBody
	GetPolicy() *CreatePolicyResponseBodyPolicy
	SetRequestId(v string) *CreatePolicyResponseBody
	GetRequestId() *string
}

type CreatePolicyResponseBody struct {
	// The information about the permission policy.
	Policy *CreatePolicyResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolicyResponseBody) GetPolicy() *CreatePolicyResponseBodyPolicy {
	return s.Policy
}

func (s *CreatePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePolicyResponseBody) SetPolicy(v *CreatePolicyResponseBodyPolicy) *CreatePolicyResponseBody {
	s.Policy = v
	return s
}

func (s *CreatePolicyResponseBody) SetRequestId(v string) *CreatePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePolicyResponseBody) Validate() error {
	if s.Policy != nil {
		if err := s.Policy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePolicyResponseBodyPolicy struct {
	// The time when the permission policy was created.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The version number of the permission policy. Default value: v1.
	//
	// example:
	//
	// v1
	DefaultVersion *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	// The description of the permission policy.
	//
	// example:
	//
	// OSS administrator
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the permission policy.
	//
	// example:
	//
	// OSS-Administrator
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the permission policy. Valid values:
	//
	// 	- Custom
	//
	// 	- System
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s CreatePolicyResponseBodyPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyResponseBodyPolicy) GoString() string {
	return s.String()
}

func (s *CreatePolicyResponseBodyPolicy) GetCreateDate() *string {
	return s.CreateDate
}

func (s *CreatePolicyResponseBodyPolicy) GetDefaultVersion() *string {
	return s.DefaultVersion
}

func (s *CreatePolicyResponseBodyPolicy) GetDescription() *string {
	return s.Description
}

func (s *CreatePolicyResponseBodyPolicy) GetPolicyName() *string {
	return s.PolicyName
}

func (s *CreatePolicyResponseBodyPolicy) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreatePolicyResponseBodyPolicy) SetCreateDate(v string) *CreatePolicyResponseBodyPolicy {
	s.CreateDate = &v
	return s
}

func (s *CreatePolicyResponseBodyPolicy) SetDefaultVersion(v string) *CreatePolicyResponseBodyPolicy {
	s.DefaultVersion = &v
	return s
}

func (s *CreatePolicyResponseBodyPolicy) SetDescription(v string) *CreatePolicyResponseBodyPolicy {
	s.Description = &v
	return s
}

func (s *CreatePolicyResponseBodyPolicy) SetPolicyName(v string) *CreatePolicyResponseBodyPolicy {
	s.PolicyName = &v
	return s
}

func (s *CreatePolicyResponseBodyPolicy) SetPolicyType(v string) *CreatePolicyResponseBodyPolicy {
	s.PolicyType = &v
	return s
}

func (s *CreatePolicyResponseBodyPolicy) Validate() error {
	return dara.Validate(s)
}
