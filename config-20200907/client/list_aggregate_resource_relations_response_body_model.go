// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateResourceRelationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListAggregateResourceRelationsResponseBody
	GetRequestId() *string
	SetResourceRelations(v *ListAggregateResourceRelationsResponseBodyResourceRelations) *ListAggregateResourceRelationsResponseBody
	GetResourceRelations() *ListAggregateResourceRelationsResponseBodyResourceRelations
}

type ListAggregateResourceRelationsResponseBody struct {
	RequestId         *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceRelations *ListAggregateResourceRelationsResponseBodyResourceRelations `json:"ResourceRelations,omitempty" xml:"ResourceRelations,omitempty" type:"Struct"`
}

func (s ListAggregateResourceRelationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateResourceRelationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAggregateResourceRelationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAggregateResourceRelationsResponseBody) GetResourceRelations() *ListAggregateResourceRelationsResponseBodyResourceRelations {
	return s.ResourceRelations
}

func (s *ListAggregateResourceRelationsResponseBody) SetRequestId(v string) *ListAggregateResourceRelationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAggregateResourceRelationsResponseBody) SetResourceRelations(v *ListAggregateResourceRelationsResponseBodyResourceRelations) *ListAggregateResourceRelationsResponseBody {
	s.ResourceRelations = v
	return s
}

func (s *ListAggregateResourceRelationsResponseBody) Validate() error {
	if s.ResourceRelations != nil {
		if err := s.ResourceRelations.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAggregateResourceRelationsResponseBodyResourceRelations struct {
	MaxResults           *int32                                                                             `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken            *string                                                                            `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceRelationList []*ListAggregateResourceRelationsResponseBodyResourceRelationsResourceRelationList `json:"ResourceRelationList,omitempty" xml:"ResourceRelationList,omitempty" type:"Repeated"`
}

func (s ListAggregateResourceRelationsResponseBodyResourceRelations) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateResourceRelationsResponseBodyResourceRelations) GoString() string {
	return s.String()
}

func (s *ListAggregateResourceRelationsResponseBodyResourceRelations) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAggregateResourceRelationsResponseBodyResourceRelations) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAggregateResourceRelationsResponseBodyResourceRelations) GetResourceRelationList() []*ListAggregateResourceRelationsResponseBodyResourceRelationsResourceRelationList {
	return s.ResourceRelationList
}

func (s *ListAggregateResourceRelationsResponseBodyResourceRelations) SetMaxResults(v int32) *ListAggregateResourceRelationsResponseBodyResourceRelations {
	s.MaxResults = &v
	return s
}

func (s *ListAggregateResourceRelationsResponseBodyResourceRelations) SetNextToken(v string) *ListAggregateResourceRelationsResponseBodyResourceRelations {
	s.NextToken = &v
	return s
}

func (s *ListAggregateResourceRelationsResponseBodyResourceRelations) SetResourceRelationList(v []*ListAggregateResourceRelationsResponseBodyResourceRelationsResourceRelationList) *ListAggregateResourceRelationsResponseBodyResourceRelations {
	s.ResourceRelationList = v
	return s
}

func (s *ListAggregateResourceRelationsResponseBodyResourceRelations) Validate() error {
	if s.ResourceRelationList != nil {
		for _, item := range s.ResourceRelationList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAggregateResourceRelationsResponseBodyResourceRelationsResourceRelationList struct {
	AccountId              *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	RelationType           *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	SourceResourceId       *string `json:"SourceResourceId,omitempty" xml:"SourceResourceId,omitempty"`
	SourceResourceRegionId *string `json:"SourceResourceRegionId,omitempty" xml:"SourceResourceRegionId,omitempty"`
	SourceResourceType     *string `json:"SourceResourceType,omitempty" xml:"SourceResourceType,omitempty"`
	TargetResourceId       *string `json:"TargetResourceId,omitempty" xml:"TargetResourceId,omitempty"`
	TargetResourceType     *string `json:"TargetResourceType,omitempty" xml:"TargetResourceType,omitempty"`
}

func (s ListAggregateResourceRelationsResponseBodyResourceRelationsResourceRelationList) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateResourceRelationsResponseBodyResourceRelationsResourceRelationList) GoString() string {
	return s.String()
}

func (s *ListAggregateResourceRelationsResponseBodyResourceRelationsResourceRelationList) GetAccountId() *int64 {
	return s.AccountId
}

func (s *ListAggregateResourceRelationsResponseBodyResourceRelationsResourceRelationList) GetRelationType() *string {
	return s.RelationType
}

func (s *ListAggregateResourceRelationsResponseBodyResourceRelationsResourceRelationList) GetSourceResourceId() *string {
	return s.SourceResourceId
}

func (s *ListAggregateResourceRelationsResponseBodyResourceRelationsResourceRelationList) GetSourceResourceRegionId() *string {
	return s.SourceResourceRegionId
}

func (s *ListAggregateResourceRelationsResponseBodyResourceRelationsResourceRelationList) GetSourceResourceType() *string {
	return s.SourceResourceType
}

func (s *ListAggregateResourceRelationsResponseBodyResourceRelationsResourceRelationList) GetTargetResourceId() *string {
	return s.TargetResourceId
}

func (s *ListAggregateResourceRelationsResponseBodyResourceRelationsResourceRelationList) GetTargetResourceType() *string {
	return s.TargetResourceType
}

func (s *ListAggregateResourceRelationsResponseBodyResourceRelationsResourceRelationList) SetAccountId(v int64) *ListAggregateResourceRelationsResponseBodyResourceRelationsResourceRelationList {
	s.AccountId = &v
	return s
}

func (s *ListAggregateResourceRelationsResponseBodyResourceRelationsResourceRelationList) SetRelationType(v string) *ListAggregateResourceRelationsResponseBodyResourceRelationsResourceRelationList {
	s.RelationType = &v
	return s
}

func (s *ListAggregateResourceRelationsResponseBodyResourceRelationsResourceRelationList) SetSourceResourceId(v string) *ListAggregateResourceRelationsResponseBodyResourceRelationsResourceRelationList {
	s.SourceResourceId = &v
	return s
}

func (s *ListAggregateResourceRelationsResponseBodyResourceRelationsResourceRelationList) SetSourceResourceRegionId(v string) *ListAggregateResourceRelationsResponseBodyResourceRelationsResourceRelationList {
	s.SourceResourceRegionId = &v
	return s
}

func (s *ListAggregateResourceRelationsResponseBodyResourceRelationsResourceRelationList) SetSourceResourceType(v string) *ListAggregateResourceRelationsResponseBodyResourceRelationsResourceRelationList {
	s.SourceResourceType = &v
	return s
}

func (s *ListAggregateResourceRelationsResponseBodyResourceRelationsResourceRelationList) SetTargetResourceId(v string) *ListAggregateResourceRelationsResponseBodyResourceRelationsResourceRelationList {
	s.TargetResourceId = &v
	return s
}

func (s *ListAggregateResourceRelationsResponseBodyResourceRelationsResourceRelationList) SetTargetResourceType(v string) *ListAggregateResourceRelationsResponseBodyResourceRelationsResourceRelationList {
	s.TargetResourceType = &v
	return s
}

func (s *ListAggregateResourceRelationsResponseBodyResourceRelationsResourceRelationList) Validate() error {
	return dara.Validate(s)
}
