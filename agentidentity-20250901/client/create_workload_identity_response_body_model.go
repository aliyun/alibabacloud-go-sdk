// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkloadIdentityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateWorkloadIdentityResponseBody
	GetRequestId() *string
	SetWorkloadIdentity(v *CreateWorkloadIdentityResponseBodyWorkloadIdentity) *CreateWorkloadIdentityResponseBody
	GetWorkloadIdentity() *CreateWorkloadIdentityResponseBodyWorkloadIdentity
}

type CreateWorkloadIdentityResponseBody struct {
	// example:
	//
	// D325DF9D-7CA8-5C47-BA0C-9A74873F2BE3
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WorkloadIdentity *CreateWorkloadIdentityResponseBodyWorkloadIdentity `json:"WorkloadIdentity,omitempty" xml:"WorkloadIdentity,omitempty" type:"Struct"`
}

func (s CreateWorkloadIdentityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkloadIdentityResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkloadIdentityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWorkloadIdentityResponseBody) GetWorkloadIdentity() *CreateWorkloadIdentityResponseBodyWorkloadIdentity {
	return s.WorkloadIdentity
}

func (s *CreateWorkloadIdentityResponseBody) SetRequestId(v string) *CreateWorkloadIdentityResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkloadIdentityResponseBody) SetWorkloadIdentity(v *CreateWorkloadIdentityResponseBodyWorkloadIdentity) *CreateWorkloadIdentityResponseBody {
	s.WorkloadIdentity = v
	return s
}

func (s *CreateWorkloadIdentityResponseBody) Validate() error {
	if s.WorkloadIdentity != nil {
		if err := s.WorkloadIdentity.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWorkloadIdentityResponseBodyWorkloadIdentity struct {
	AllowedResourceOAuth2ReturnURLs []*string `json:"AllowedResourceOAuth2ReturnURLs,omitempty" xml:"AllowedResourceOAuth2ReturnURLs,omitempty" type:"Repeated"`
	// example:
	//
	// 2025-12-18T06:19:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// example agent
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// identity-provider-okta
	IdentityProviderName *string `json:"IdentityProviderName,omitempty" xml:"IdentityProviderName,omitempty"`
	// example:
	//
	// acs:ram::123456:role/agent-101-role
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// example:
	//
	// 2025-12-18T06:19:17Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// acs:agentidentity:cn-beijing:123456:workloadidentitydirectory/default/workloadidentity/agent-101
	WorkloadIdentityArn *string `json:"WorkloadIdentityArn,omitempty" xml:"WorkloadIdentityArn,omitempty"`
	// example:
	//
	// agent-101
	WorkloadIdentityName *string `json:"WorkloadIdentityName,omitempty" xml:"WorkloadIdentityName,omitempty"`
}

func (s CreateWorkloadIdentityResponseBodyWorkloadIdentity) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkloadIdentityResponseBodyWorkloadIdentity) GoString() string {
	return s.String()
}

func (s *CreateWorkloadIdentityResponseBodyWorkloadIdentity) GetAllowedResourceOAuth2ReturnURLs() []*string {
	return s.AllowedResourceOAuth2ReturnURLs
}

func (s *CreateWorkloadIdentityResponseBodyWorkloadIdentity) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateWorkloadIdentityResponseBodyWorkloadIdentity) GetDescription() *string {
	return s.Description
}

func (s *CreateWorkloadIdentityResponseBodyWorkloadIdentity) GetIdentityProviderName() *string {
	return s.IdentityProviderName
}

func (s *CreateWorkloadIdentityResponseBodyWorkloadIdentity) GetRoleArn() *string {
	return s.RoleArn
}

func (s *CreateWorkloadIdentityResponseBodyWorkloadIdentity) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *CreateWorkloadIdentityResponseBodyWorkloadIdentity) GetWorkloadIdentityArn() *string {
	return s.WorkloadIdentityArn
}

func (s *CreateWorkloadIdentityResponseBodyWorkloadIdentity) GetWorkloadIdentityName() *string {
	return s.WorkloadIdentityName
}

func (s *CreateWorkloadIdentityResponseBodyWorkloadIdentity) SetAllowedResourceOAuth2ReturnURLs(v []*string) *CreateWorkloadIdentityResponseBodyWorkloadIdentity {
	s.AllowedResourceOAuth2ReturnURLs = v
	return s
}

func (s *CreateWorkloadIdentityResponseBodyWorkloadIdentity) SetCreateTime(v string) *CreateWorkloadIdentityResponseBodyWorkloadIdentity {
	s.CreateTime = &v
	return s
}

func (s *CreateWorkloadIdentityResponseBodyWorkloadIdentity) SetDescription(v string) *CreateWorkloadIdentityResponseBodyWorkloadIdentity {
	s.Description = &v
	return s
}

func (s *CreateWorkloadIdentityResponseBodyWorkloadIdentity) SetIdentityProviderName(v string) *CreateWorkloadIdentityResponseBodyWorkloadIdentity {
	s.IdentityProviderName = &v
	return s
}

func (s *CreateWorkloadIdentityResponseBodyWorkloadIdentity) SetRoleArn(v string) *CreateWorkloadIdentityResponseBodyWorkloadIdentity {
	s.RoleArn = &v
	return s
}

func (s *CreateWorkloadIdentityResponseBodyWorkloadIdentity) SetUpdateTime(v string) *CreateWorkloadIdentityResponseBodyWorkloadIdentity {
	s.UpdateTime = &v
	return s
}

func (s *CreateWorkloadIdentityResponseBodyWorkloadIdentity) SetWorkloadIdentityArn(v string) *CreateWorkloadIdentityResponseBodyWorkloadIdentity {
	s.WorkloadIdentityArn = &v
	return s
}

func (s *CreateWorkloadIdentityResponseBodyWorkloadIdentity) SetWorkloadIdentityName(v string) *CreateWorkloadIdentityResponseBodyWorkloadIdentity {
	s.WorkloadIdentityName = &v
	return s
}

func (s *CreateWorkloadIdentityResponseBodyWorkloadIdentity) Validate() error {
	return dara.Validate(s)
}
