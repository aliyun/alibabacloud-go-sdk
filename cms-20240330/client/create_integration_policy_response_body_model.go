// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIntegrationPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreated(v bool) *CreateIntegrationPolicyResponseBody
	GetCreated() *bool
	SetPolicy(v *CreateIntegrationPolicyResponseBodyPolicy) *CreateIntegrationPolicyResponseBody
	GetPolicy() *CreateIntegrationPolicyResponseBodyPolicy
	SetRequestId(v string) *CreateIntegrationPolicyResponseBody
	GetRequestId() *string
}

type CreateIntegrationPolicyResponseBody struct {
	// Whether it was created.
	//
	// example:
	//
	// true
	Created *bool `json:"created,omitempty" xml:"created,omitempty"`
	// Uploaded policy.
	Policy *CreateIntegrationPolicyResponseBodyPolicy `json:"policy,omitempty" xml:"policy,omitempty" type:"Struct"`
	// Request ID.
	//
	// example:
	//
	// CD8BA7D6-995D-578D-9941-78B0FECD14B5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateIntegrationPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIntegrationPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIntegrationPolicyResponseBody) GetCreated() *bool {
	return s.Created
}

func (s *CreateIntegrationPolicyResponseBody) GetPolicy() *CreateIntegrationPolicyResponseBodyPolicy {
	return s.Policy
}

func (s *CreateIntegrationPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIntegrationPolicyResponseBody) SetCreated(v bool) *CreateIntegrationPolicyResponseBody {
	s.Created = &v
	return s
}

func (s *CreateIntegrationPolicyResponseBody) SetPolicy(v *CreateIntegrationPolicyResponseBodyPolicy) *CreateIntegrationPolicyResponseBody {
	s.Policy = v
	return s
}

func (s *CreateIntegrationPolicyResponseBody) SetRequestId(v string) *CreateIntegrationPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIntegrationPolicyResponseBody) Validate() error {
	if s.Policy != nil {
		if err := s.Policy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateIntegrationPolicyResponseBodyPolicy struct {
	// Entity group ID.
	//
	// example:
	//
	// eg-b79f65d11fb94e779867cf937c3a3002
	EntityGroupId *string `json:"entityGroupId,omitempty" xml:"entityGroupId,omitempty"`
	// Policy ID.
	//
	// example:
	//
	// policy-14c8e9a29b0a46da843f8781471062ff
	PolicyId *string `json:"policyId,omitempty" xml:"policyId,omitempty"`
	// Policy name.
	//
	// example:
	//
	// metrics-inner-manage
	PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty"`
	// Policy type.
	//
	// example:
	//
	// CS
	PolicyType *string `json:"policyType,omitempty" xml:"policyType,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// User ID.
	//
	// example:
	//
	// u1234567
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// The workspace where the Policy resides.
	//
	// example:
	//
	// prometheus
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateIntegrationPolicyResponseBodyPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateIntegrationPolicyResponseBodyPolicy) GoString() string {
	return s.String()
}

func (s *CreateIntegrationPolicyResponseBodyPolicy) GetEntityGroupId() *string {
	return s.EntityGroupId
}

func (s *CreateIntegrationPolicyResponseBodyPolicy) GetPolicyId() *string {
	return s.PolicyId
}

func (s *CreateIntegrationPolicyResponseBodyPolicy) GetPolicyName() *string {
	return s.PolicyName
}

func (s *CreateIntegrationPolicyResponseBodyPolicy) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreateIntegrationPolicyResponseBodyPolicy) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateIntegrationPolicyResponseBodyPolicy) GetUserId() *string {
	return s.UserId
}

func (s *CreateIntegrationPolicyResponseBodyPolicy) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateIntegrationPolicyResponseBodyPolicy) SetEntityGroupId(v string) *CreateIntegrationPolicyResponseBodyPolicy {
	s.EntityGroupId = &v
	return s
}

func (s *CreateIntegrationPolicyResponseBodyPolicy) SetPolicyId(v string) *CreateIntegrationPolicyResponseBodyPolicy {
	s.PolicyId = &v
	return s
}

func (s *CreateIntegrationPolicyResponseBodyPolicy) SetPolicyName(v string) *CreateIntegrationPolicyResponseBodyPolicy {
	s.PolicyName = &v
	return s
}

func (s *CreateIntegrationPolicyResponseBodyPolicy) SetPolicyType(v string) *CreateIntegrationPolicyResponseBodyPolicy {
	s.PolicyType = &v
	return s
}

func (s *CreateIntegrationPolicyResponseBodyPolicy) SetRegionId(v string) *CreateIntegrationPolicyResponseBodyPolicy {
	s.RegionId = &v
	return s
}

func (s *CreateIntegrationPolicyResponseBodyPolicy) SetUserId(v string) *CreateIntegrationPolicyResponseBodyPolicy {
	s.UserId = &v
	return s
}

func (s *CreateIntegrationPolicyResponseBodyPolicy) SetWorkspace(v string) *CreateIntegrationPolicyResponseBodyPolicy {
	s.Workspace = &v
	return s
}

func (s *CreateIntegrationPolicyResponseBodyPolicy) Validate() error {
	return dara.Validate(s)
}
