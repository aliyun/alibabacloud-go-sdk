// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceRelationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListResourceRelationsResponseBody
	GetRequestId() *string
	SetResourceRelations(v *ListResourceRelationsResponseBodyResourceRelations) *ListResourceRelationsResponseBody
	GetResourceRelations() *ListResourceRelationsResponseBodyResourceRelations
}

type ListResourceRelationsResponseBody struct {
	RequestId         *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceRelations *ListResourceRelationsResponseBodyResourceRelations `json:"ResourceRelations,omitempty" xml:"ResourceRelations,omitempty" type:"Struct"`
}

func (s ListResourceRelationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceRelationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceRelationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceRelationsResponseBody) GetResourceRelations() *ListResourceRelationsResponseBodyResourceRelations {
	return s.ResourceRelations
}

func (s *ListResourceRelationsResponseBody) SetRequestId(v string) *ListResourceRelationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceRelationsResponseBody) SetResourceRelations(v *ListResourceRelationsResponseBodyResourceRelations) *ListResourceRelationsResponseBody {
	s.ResourceRelations = v
	return s
}

func (s *ListResourceRelationsResponseBody) Validate() error {
	if s.ResourceRelations != nil {
		if err := s.ResourceRelations.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListResourceRelationsResponseBodyResourceRelations struct {
	MaxResults           *int32                                                                    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken            *string                                                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceRelationList []*ListResourceRelationsResponseBodyResourceRelationsResourceRelationList `json:"ResourceRelationList,omitempty" xml:"ResourceRelationList,omitempty" type:"Repeated"`
}

func (s ListResourceRelationsResponseBodyResourceRelations) String() string {
	return dara.Prettify(s)
}

func (s ListResourceRelationsResponseBodyResourceRelations) GoString() string {
	return s.String()
}

func (s *ListResourceRelationsResponseBodyResourceRelations) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResourceRelationsResponseBodyResourceRelations) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceRelationsResponseBodyResourceRelations) GetResourceRelationList() []*ListResourceRelationsResponseBodyResourceRelationsResourceRelationList {
	return s.ResourceRelationList
}

func (s *ListResourceRelationsResponseBodyResourceRelations) SetMaxResults(v int32) *ListResourceRelationsResponseBodyResourceRelations {
	s.MaxResults = &v
	return s
}

func (s *ListResourceRelationsResponseBodyResourceRelations) SetNextToken(v string) *ListResourceRelationsResponseBodyResourceRelations {
	s.NextToken = &v
	return s
}

func (s *ListResourceRelationsResponseBodyResourceRelations) SetResourceRelationList(v []*ListResourceRelationsResponseBodyResourceRelationsResourceRelationList) *ListResourceRelationsResponseBodyResourceRelations {
	s.ResourceRelationList = v
	return s
}

func (s *ListResourceRelationsResponseBodyResourceRelations) Validate() error {
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

type ListResourceRelationsResponseBodyResourceRelationsResourceRelationList struct {
	AccountId              *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	RelationType           *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	SourceResourceId       *string `json:"SourceResourceId,omitempty" xml:"SourceResourceId,omitempty"`
	SourceResourceRegionId *string `json:"SourceResourceRegionId,omitempty" xml:"SourceResourceRegionId,omitempty"`
	SourceResourceType     *string `json:"SourceResourceType,omitempty" xml:"SourceResourceType,omitempty"`
	TargetResourceId       *string `json:"TargetResourceId,omitempty" xml:"TargetResourceId,omitempty"`
	TargetResourceType     *string `json:"TargetResourceType,omitempty" xml:"TargetResourceType,omitempty"`
}

func (s ListResourceRelationsResponseBodyResourceRelationsResourceRelationList) String() string {
	return dara.Prettify(s)
}

func (s ListResourceRelationsResponseBodyResourceRelationsResourceRelationList) GoString() string {
	return s.String()
}

func (s *ListResourceRelationsResponseBodyResourceRelationsResourceRelationList) GetAccountId() *int64 {
	return s.AccountId
}

func (s *ListResourceRelationsResponseBodyResourceRelationsResourceRelationList) GetRelationType() *string {
	return s.RelationType
}

func (s *ListResourceRelationsResponseBodyResourceRelationsResourceRelationList) GetSourceResourceId() *string {
	return s.SourceResourceId
}

func (s *ListResourceRelationsResponseBodyResourceRelationsResourceRelationList) GetSourceResourceRegionId() *string {
	return s.SourceResourceRegionId
}

func (s *ListResourceRelationsResponseBodyResourceRelationsResourceRelationList) GetSourceResourceType() *string {
	return s.SourceResourceType
}

func (s *ListResourceRelationsResponseBodyResourceRelationsResourceRelationList) GetTargetResourceId() *string {
	return s.TargetResourceId
}

func (s *ListResourceRelationsResponseBodyResourceRelationsResourceRelationList) GetTargetResourceType() *string {
	return s.TargetResourceType
}

func (s *ListResourceRelationsResponseBodyResourceRelationsResourceRelationList) SetAccountId(v int64) *ListResourceRelationsResponseBodyResourceRelationsResourceRelationList {
	s.AccountId = &v
	return s
}

func (s *ListResourceRelationsResponseBodyResourceRelationsResourceRelationList) SetRelationType(v string) *ListResourceRelationsResponseBodyResourceRelationsResourceRelationList {
	s.RelationType = &v
	return s
}

func (s *ListResourceRelationsResponseBodyResourceRelationsResourceRelationList) SetSourceResourceId(v string) *ListResourceRelationsResponseBodyResourceRelationsResourceRelationList {
	s.SourceResourceId = &v
	return s
}

func (s *ListResourceRelationsResponseBodyResourceRelationsResourceRelationList) SetSourceResourceRegionId(v string) *ListResourceRelationsResponseBodyResourceRelationsResourceRelationList {
	s.SourceResourceRegionId = &v
	return s
}

func (s *ListResourceRelationsResponseBodyResourceRelationsResourceRelationList) SetSourceResourceType(v string) *ListResourceRelationsResponseBodyResourceRelationsResourceRelationList {
	s.SourceResourceType = &v
	return s
}

func (s *ListResourceRelationsResponseBodyResourceRelationsResourceRelationList) SetTargetResourceId(v string) *ListResourceRelationsResponseBodyResourceRelationsResourceRelationList {
	s.TargetResourceId = &v
	return s
}

func (s *ListResourceRelationsResponseBodyResourceRelationsResourceRelationList) SetTargetResourceType(v string) *ListResourceRelationsResponseBodyResourceRelationsResourceRelationList {
	s.TargetResourceType = &v
	return s
}

func (s *ListResourceRelationsResponseBodyResourceRelationsResourceRelationList) Validate() error {
	return dara.Validate(s)
}
