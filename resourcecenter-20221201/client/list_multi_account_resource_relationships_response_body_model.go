// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultiAccountResourceRelationshipsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListMultiAccountResourceRelationshipsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListMultiAccountResourceRelationshipsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMultiAccountResourceRelationshipsResponseBody
	GetRequestId() *string
	SetResourceRelationships(v []*ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships) *ListMultiAccountResourceRelationshipsResponseBody
	GetResourceRelationships() []*ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships
	SetScope(v string) *ListMultiAccountResourceRelationshipsResponseBody
	GetScope() *string
}

type ListMultiAccountResourceRelationshipsResponseBody struct {
	// The maximum number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BCAB07BA-82FA-5DC0-9322-FB7ED726481D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource relationships.
	ResourceRelationships []*ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships `json:"ResourceRelationships,omitempty" xml:"ResourceRelationships,omitempty" type:"Repeated"`
	// The search scope. Valid values:
	//
	// 	- ID of a resource directory: Resources within the management account and all members of the resource directory are searched.
	//
	// 	- ID of the Root folder: Resources within all members in the Root folder and the subfolders of the Root folder are searched.
	//
	// 	- ID of a folder: Resources within all members in the folder are searched.
	//
	// 	- ID of a member: Resources within the member are searched.
	//
	// example:
	//
	// rd-r4****
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s ListMultiAccountResourceRelationshipsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMultiAccountResourceRelationshipsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMultiAccountResourceRelationshipsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMultiAccountResourceRelationshipsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMultiAccountResourceRelationshipsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMultiAccountResourceRelationshipsResponseBody) GetResourceRelationships() []*ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships {
	return s.ResourceRelationships
}

func (s *ListMultiAccountResourceRelationshipsResponseBody) GetScope() *string {
	return s.Scope
}

func (s *ListMultiAccountResourceRelationshipsResponseBody) SetMaxResults(v int32) *ListMultiAccountResourceRelationshipsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsResponseBody) SetNextToken(v string) *ListMultiAccountResourceRelationshipsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsResponseBody) SetRequestId(v string) *ListMultiAccountResourceRelationshipsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsResponseBody) SetResourceRelationships(v []*ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships) *ListMultiAccountResourceRelationshipsResponseBody {
	s.ResourceRelationships = v
	return s
}

func (s *ListMultiAccountResourceRelationshipsResponseBody) SetScope(v string) *ListMultiAccountResourceRelationshipsResponseBody {
	s.Scope = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsResponseBody) Validate() error {
	if s.ResourceRelationships != nil {
		for _, item := range s.ResourceRelationships {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships struct {
	// The ID of the management account or member.
	//
	// example:
	//
	// 193396142051****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The region ID of the resource.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the associated resource.
	//
	// example:
	//
	// vpc-uf6m5okksddm6c9lh7***
	RelatedResourceId *string `json:"RelatedResourceId,omitempty" xml:"RelatedResourceId,omitempty"`
	// The region ID of the associated resource.
	//
	// example:
	//
	// cn-shanghai
	RelatedResourceRegionId *string `json:"RelatedResourceRegionId,omitempty" xml:"RelatedResourceRegionId,omitempty"`
	// The type of the associated resource.
	//
	// example:
	//
	// ACS::VPC::VPC
	RelatedResourceType *string `json:"RelatedResourceType,omitempty" xml:"RelatedResourceType,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// m-eb3hji****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// ACS::ACK::Cluster
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships) String() string {
	return dara.Prettify(s)
}

func (s ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships) GoString() string {
	return s.String()
}

func (s *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships) GetAccountId() *string {
	return s.AccountId
}

func (s *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships) GetRegionId() *string {
	return s.RegionId
}

func (s *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships) GetRelatedResourceId() *string {
	return s.RelatedResourceId
}

func (s *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships) GetRelatedResourceRegionId() *string {
	return s.RelatedResourceRegionId
}

func (s *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships) GetRelatedResourceType() *string {
	return s.RelatedResourceType
}

func (s *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships) SetAccountId(v string) *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships {
	s.AccountId = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships) SetRegionId(v string) *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships {
	s.RegionId = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships) SetRelatedResourceId(v string) *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships {
	s.RelatedResourceId = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships) SetRelatedResourceRegionId(v string) *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships {
	s.RelatedResourceRegionId = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships) SetRelatedResourceType(v string) *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships {
	s.RelatedResourceType = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships) SetResourceId(v string) *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships {
	s.ResourceId = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships) SetResourceType(v string) *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships {
	s.ResourceType = &v
	return s
}

func (s *ListMultiAccountResourceRelationshipsResponseBodyResourceRelationships) Validate() error {
	return dara.Validate(s)
}
