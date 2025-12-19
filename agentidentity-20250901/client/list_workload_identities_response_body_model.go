// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkloadIdentitiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListWorkloadIdentitiesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListWorkloadIdentitiesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListWorkloadIdentitiesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListWorkloadIdentitiesResponseBody
	GetTotalCount() *int32
	SetWorkloadIdentities(v []*ListWorkloadIdentitiesResponseBodyWorkloadIdentities) *ListWorkloadIdentitiesResponseBody
	GetWorkloadIdentities() []*ListWorkloadIdentitiesResponseBodyWorkloadIdentities
}

type ListWorkloadIdentitiesResponseBody struct {
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAcCoknY19uiBwPrAe1W7XMikkA6+rCQddIGHccphiDPypD8zCxQkHV2pg8CkZvyRKg==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 173C69C9-9C07-5B25-9477-603A33E5DA32
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 33
	TotalCount         *int32                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	WorkloadIdentities []*ListWorkloadIdentitiesResponseBodyWorkloadIdentities `json:"WorkloadIdentities,omitempty" xml:"WorkloadIdentities,omitempty" type:"Repeated"`
}

func (s ListWorkloadIdentitiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkloadIdentitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkloadIdentitiesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkloadIdentitiesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkloadIdentitiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkloadIdentitiesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListWorkloadIdentitiesResponseBody) GetWorkloadIdentities() []*ListWorkloadIdentitiesResponseBodyWorkloadIdentities {
	return s.WorkloadIdentities
}

func (s *ListWorkloadIdentitiesResponseBody) SetMaxResults(v int32) *ListWorkloadIdentitiesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWorkloadIdentitiesResponseBody) SetNextToken(v string) *ListWorkloadIdentitiesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWorkloadIdentitiesResponseBody) SetRequestId(v string) *ListWorkloadIdentitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkloadIdentitiesResponseBody) SetTotalCount(v int32) *ListWorkloadIdentitiesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWorkloadIdentitiesResponseBody) SetWorkloadIdentities(v []*ListWorkloadIdentitiesResponseBodyWorkloadIdentities) *ListWorkloadIdentitiesResponseBody {
	s.WorkloadIdentities = v
	return s
}

func (s *ListWorkloadIdentitiesResponseBody) Validate() error {
	if s.WorkloadIdentities != nil {
		for _, item := range s.WorkloadIdentities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkloadIdentitiesResponseBodyWorkloadIdentities struct {
	AllowedResourceOAuth2ReturnURLs []*string `json:"AllowedResourceOAuth2ReturnURLs,omitempty" xml:"AllowedResourceOAuth2ReturnURLs,omitempty" type:"Repeated"`
	// example:
	//
	// 2025-12-18T06:19:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// example agent
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IdentityProviderName *string `json:"IdentityProviderName,omitempty" xml:"IdentityProviderName,omitempty"`
	// example:
	//
	// acs:ram::1953507478506681:role/test-rrsa-cb5ca90a20f854671adbac6ed4559a654
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

func (s ListWorkloadIdentitiesResponseBodyWorkloadIdentities) String() string {
	return dara.Prettify(s)
}

func (s ListWorkloadIdentitiesResponseBodyWorkloadIdentities) GoString() string {
	return s.String()
}

func (s *ListWorkloadIdentitiesResponseBodyWorkloadIdentities) GetAllowedResourceOAuth2ReturnURLs() []*string {
	return s.AllowedResourceOAuth2ReturnURLs
}

func (s *ListWorkloadIdentitiesResponseBodyWorkloadIdentities) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListWorkloadIdentitiesResponseBodyWorkloadIdentities) GetDescription() *string {
	return s.Description
}

func (s *ListWorkloadIdentitiesResponseBodyWorkloadIdentities) GetIdentityProviderName() *string {
	return s.IdentityProviderName
}

func (s *ListWorkloadIdentitiesResponseBodyWorkloadIdentities) GetRoleArn() *string {
	return s.RoleArn
}

func (s *ListWorkloadIdentitiesResponseBodyWorkloadIdentities) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListWorkloadIdentitiesResponseBodyWorkloadIdentities) GetWorkloadIdentityArn() *string {
	return s.WorkloadIdentityArn
}

func (s *ListWorkloadIdentitiesResponseBodyWorkloadIdentities) GetWorkloadIdentityName() *string {
	return s.WorkloadIdentityName
}

func (s *ListWorkloadIdentitiesResponseBodyWorkloadIdentities) SetAllowedResourceOAuth2ReturnURLs(v []*string) *ListWorkloadIdentitiesResponseBodyWorkloadIdentities {
	s.AllowedResourceOAuth2ReturnURLs = v
	return s
}

func (s *ListWorkloadIdentitiesResponseBodyWorkloadIdentities) SetCreateTime(v string) *ListWorkloadIdentitiesResponseBodyWorkloadIdentities {
	s.CreateTime = &v
	return s
}

func (s *ListWorkloadIdentitiesResponseBodyWorkloadIdentities) SetDescription(v string) *ListWorkloadIdentitiesResponseBodyWorkloadIdentities {
	s.Description = &v
	return s
}

func (s *ListWorkloadIdentitiesResponseBodyWorkloadIdentities) SetIdentityProviderName(v string) *ListWorkloadIdentitiesResponseBodyWorkloadIdentities {
	s.IdentityProviderName = &v
	return s
}

func (s *ListWorkloadIdentitiesResponseBodyWorkloadIdentities) SetRoleArn(v string) *ListWorkloadIdentitiesResponseBodyWorkloadIdentities {
	s.RoleArn = &v
	return s
}

func (s *ListWorkloadIdentitiesResponseBodyWorkloadIdentities) SetUpdateTime(v string) *ListWorkloadIdentitiesResponseBodyWorkloadIdentities {
	s.UpdateTime = &v
	return s
}

func (s *ListWorkloadIdentitiesResponseBodyWorkloadIdentities) SetWorkloadIdentityArn(v string) *ListWorkloadIdentitiesResponseBodyWorkloadIdentities {
	s.WorkloadIdentityArn = &v
	return s
}

func (s *ListWorkloadIdentitiesResponseBodyWorkloadIdentities) SetWorkloadIdentityName(v string) *ListWorkloadIdentitiesResponseBodyWorkloadIdentities {
	s.WorkloadIdentityName = &v
	return s
}

func (s *ListWorkloadIdentitiesResponseBodyWorkloadIdentities) Validate() error {
	return dara.Validate(s)
}
