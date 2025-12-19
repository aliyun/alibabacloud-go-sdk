// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkloadIdentityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetWorkloadIdentityResponseBody
	GetRequestId() *string
	SetWorkloadIdentity(v *GetWorkloadIdentityResponseBodyWorkloadIdentity) *GetWorkloadIdentityResponseBody
	GetWorkloadIdentity() *GetWorkloadIdentityResponseBodyWorkloadIdentity
}

type GetWorkloadIdentityResponseBody struct {
	// example:
	//
	// 5EEF5C1D-E951-5C0D-B0BC-5FF1B626CFC6
	RequestId        *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WorkloadIdentity *GetWorkloadIdentityResponseBodyWorkloadIdentity `json:"WorkloadIdentity,omitempty" xml:"WorkloadIdentity,omitempty" type:"Struct"`
}

func (s GetWorkloadIdentityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkloadIdentityResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkloadIdentityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkloadIdentityResponseBody) GetWorkloadIdentity() *GetWorkloadIdentityResponseBodyWorkloadIdentity {
	return s.WorkloadIdentity
}

func (s *GetWorkloadIdentityResponseBody) SetRequestId(v string) *GetWorkloadIdentityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkloadIdentityResponseBody) SetWorkloadIdentity(v *GetWorkloadIdentityResponseBodyWorkloadIdentity) *GetWorkloadIdentityResponseBody {
	s.WorkloadIdentity = v
	return s
}

func (s *GetWorkloadIdentityResponseBody) Validate() error {
	if s.WorkloadIdentity != nil {
		if err := s.WorkloadIdentity.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkloadIdentityResponseBodyWorkloadIdentity struct {
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

func (s GetWorkloadIdentityResponseBodyWorkloadIdentity) String() string {
	return dara.Prettify(s)
}

func (s GetWorkloadIdentityResponseBodyWorkloadIdentity) GoString() string {
	return s.String()
}

func (s *GetWorkloadIdentityResponseBodyWorkloadIdentity) GetAllowedResourceOAuth2ReturnURLs() []*string {
	return s.AllowedResourceOAuth2ReturnURLs
}

func (s *GetWorkloadIdentityResponseBodyWorkloadIdentity) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetWorkloadIdentityResponseBodyWorkloadIdentity) GetDescription() *string {
	return s.Description
}

func (s *GetWorkloadIdentityResponseBodyWorkloadIdentity) GetIdentityProviderName() *string {
	return s.IdentityProviderName
}

func (s *GetWorkloadIdentityResponseBodyWorkloadIdentity) GetRoleArn() *string {
	return s.RoleArn
}

func (s *GetWorkloadIdentityResponseBodyWorkloadIdentity) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetWorkloadIdentityResponseBodyWorkloadIdentity) GetWorkloadIdentityArn() *string {
	return s.WorkloadIdentityArn
}

func (s *GetWorkloadIdentityResponseBodyWorkloadIdentity) GetWorkloadIdentityName() *string {
	return s.WorkloadIdentityName
}

func (s *GetWorkloadIdentityResponseBodyWorkloadIdentity) SetAllowedResourceOAuth2ReturnURLs(v []*string) *GetWorkloadIdentityResponseBodyWorkloadIdentity {
	s.AllowedResourceOAuth2ReturnURLs = v
	return s
}

func (s *GetWorkloadIdentityResponseBodyWorkloadIdentity) SetCreateTime(v string) *GetWorkloadIdentityResponseBodyWorkloadIdentity {
	s.CreateTime = &v
	return s
}

func (s *GetWorkloadIdentityResponseBodyWorkloadIdentity) SetDescription(v string) *GetWorkloadIdentityResponseBodyWorkloadIdentity {
	s.Description = &v
	return s
}

func (s *GetWorkloadIdentityResponseBodyWorkloadIdentity) SetIdentityProviderName(v string) *GetWorkloadIdentityResponseBodyWorkloadIdentity {
	s.IdentityProviderName = &v
	return s
}

func (s *GetWorkloadIdentityResponseBodyWorkloadIdentity) SetRoleArn(v string) *GetWorkloadIdentityResponseBodyWorkloadIdentity {
	s.RoleArn = &v
	return s
}

func (s *GetWorkloadIdentityResponseBodyWorkloadIdentity) SetUpdateTime(v string) *GetWorkloadIdentityResponseBodyWorkloadIdentity {
	s.UpdateTime = &v
	return s
}

func (s *GetWorkloadIdentityResponseBodyWorkloadIdentity) SetWorkloadIdentityArn(v string) *GetWorkloadIdentityResponseBodyWorkloadIdentity {
	s.WorkloadIdentityArn = &v
	return s
}

func (s *GetWorkloadIdentityResponseBodyWorkloadIdentity) SetWorkloadIdentityName(v string) *GetWorkloadIdentityResponseBodyWorkloadIdentity {
	s.WorkloadIdentityName = &v
	return s
}

func (s *GetWorkloadIdentityResponseBodyWorkloadIdentity) Validate() error {
	return dara.Validate(s)
}
