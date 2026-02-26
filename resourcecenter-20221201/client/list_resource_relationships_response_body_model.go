// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceRelationshipsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListResourceRelationshipsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListResourceRelationshipsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListResourceRelationshipsResponseBody
	GetRequestId() *string
	SetResourceRelationships(v []*ListResourceRelationshipsResponseBodyResourceRelationships) *ListResourceRelationshipsResponseBody
	GetResourceRelationships() []*ListResourceRelationshipsResponseBodyResourceRelationships
}

type ListResourceRelationshipsResponseBody struct {
	MaxResults            *int32                                                        `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken             *string                                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId             *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceRelationships []*ListResourceRelationshipsResponseBodyResourceRelationships `json:"ResourceRelationships,omitempty" xml:"ResourceRelationships,omitempty" type:"Repeated"`
}

func (s ListResourceRelationshipsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceRelationshipsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceRelationshipsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResourceRelationshipsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceRelationshipsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceRelationshipsResponseBody) GetResourceRelationships() []*ListResourceRelationshipsResponseBodyResourceRelationships {
	return s.ResourceRelationships
}

func (s *ListResourceRelationshipsResponseBody) SetMaxResults(v int32) *ListResourceRelationshipsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListResourceRelationshipsResponseBody) SetNextToken(v string) *ListResourceRelationshipsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListResourceRelationshipsResponseBody) SetRequestId(v string) *ListResourceRelationshipsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceRelationshipsResponseBody) SetResourceRelationships(v []*ListResourceRelationshipsResponseBodyResourceRelationships) *ListResourceRelationshipsResponseBody {
	s.ResourceRelationships = v
	return s
}

func (s *ListResourceRelationshipsResponseBody) Validate() error {
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

type ListResourceRelationshipsResponseBodyResourceRelationships struct {
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RelatedResourceId       *string `json:"RelatedResourceId,omitempty" xml:"RelatedResourceId,omitempty"`
	RelatedResourceRegionId *string `json:"RelatedResourceRegionId,omitempty" xml:"RelatedResourceRegionId,omitempty"`
	RelatedResourceType     *string `json:"RelatedResourceType,omitempty" xml:"RelatedResourceType,omitempty"`
	ResourceId              *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType            *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListResourceRelationshipsResponseBodyResourceRelationships) String() string {
	return dara.Prettify(s)
}

func (s ListResourceRelationshipsResponseBodyResourceRelationships) GoString() string {
	return s.String()
}

func (s *ListResourceRelationshipsResponseBodyResourceRelationships) GetRegionId() *string {
	return s.RegionId
}

func (s *ListResourceRelationshipsResponseBodyResourceRelationships) GetRelatedResourceId() *string {
	return s.RelatedResourceId
}

func (s *ListResourceRelationshipsResponseBodyResourceRelationships) GetRelatedResourceRegionId() *string {
	return s.RelatedResourceRegionId
}

func (s *ListResourceRelationshipsResponseBodyResourceRelationships) GetRelatedResourceType() *string {
	return s.RelatedResourceType
}

func (s *ListResourceRelationshipsResponseBodyResourceRelationships) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListResourceRelationshipsResponseBodyResourceRelationships) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourceRelationshipsResponseBodyResourceRelationships) SetRegionId(v string) *ListResourceRelationshipsResponseBodyResourceRelationships {
	s.RegionId = &v
	return s
}

func (s *ListResourceRelationshipsResponseBodyResourceRelationships) SetRelatedResourceId(v string) *ListResourceRelationshipsResponseBodyResourceRelationships {
	s.RelatedResourceId = &v
	return s
}

func (s *ListResourceRelationshipsResponseBodyResourceRelationships) SetRelatedResourceRegionId(v string) *ListResourceRelationshipsResponseBodyResourceRelationships {
	s.RelatedResourceRegionId = &v
	return s
}

func (s *ListResourceRelationshipsResponseBodyResourceRelationships) SetRelatedResourceType(v string) *ListResourceRelationshipsResponseBodyResourceRelationships {
	s.RelatedResourceType = &v
	return s
}

func (s *ListResourceRelationshipsResponseBodyResourceRelationships) SetResourceId(v string) *ListResourceRelationshipsResponseBodyResourceRelationships {
	s.ResourceId = &v
	return s
}

func (s *ListResourceRelationshipsResponseBodyResourceRelationships) SetResourceType(v string) *ListResourceRelationshipsResponseBodyResourceRelationships {
	s.ResourceType = &v
	return s
}

func (s *ListResourceRelationshipsResponseBodyResourceRelationships) Validate() error {
	return dara.Validate(s)
}
