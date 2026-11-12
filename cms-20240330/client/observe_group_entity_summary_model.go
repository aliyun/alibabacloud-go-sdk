// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObserveGroupEntitySummary interface {
	dara.Model
	String() string
	GoString() string
	SetEntityCategory(v string) *ObserveGroupEntitySummary
	GetEntityCategory() *string
	SetEntityCount(v int32) *ObserveGroupEntitySummary
	GetEntityCount() *int32
	SetEntityDomain(v string) *ObserveGroupEntitySummary
	GetEntityDomain() *string
	SetEntityType(v string) *ObserveGroupEntitySummary
	GetEntityType() *string
	SetMainCount(v int64) *ObserveGroupEntitySummary
	GetMainCount() *int64
	SetRelatedCount(v int64) *ObserveGroupEntitySummary
	GetRelatedCount() *int64
}

type ObserveGroupEntitySummary struct {
	// The entity category.
	//
	// example:
	//
	// acs
	EntityCategory *string `json:"entityCategory,omitempty" xml:"entityCategory,omitempty"`
	// The total count of entities of this type.
	//
	// example:
	//
	// 9
	EntityCount *int32 `json:"entityCount,omitempty" xml:"entityCount,omitempty"`
	// The entity domain.
	//
	// example:
	//
	// acs
	EntityDomain *string `json:"entityDomain,omitempty" xml:"entityDomain,omitempty"`
	// The entity type.
	//
	// example:
	//
	// acs.rocketmq.topic
	EntityType *string `json:"entityType,omitempty" xml:"entityType,omitempty"`
	// The count of primary entities of this type.
	//
	// example:
	//
	// 0
	MainCount *int64 `json:"mainCount,omitempty" xml:"mainCount,omitempty"`
	// The count of associated child entities of this type.
	//
	// example:
	//
	// 9
	RelatedCount *int64 `json:"relatedCount,omitempty" xml:"relatedCount,omitempty"`
}

func (s ObserveGroupEntitySummary) String() string {
	return dara.Prettify(s)
}

func (s ObserveGroupEntitySummary) GoString() string {
	return s.String()
}

func (s *ObserveGroupEntitySummary) GetEntityCategory() *string {
	return s.EntityCategory
}

func (s *ObserveGroupEntitySummary) GetEntityCount() *int32 {
	return s.EntityCount
}

func (s *ObserveGroupEntitySummary) GetEntityDomain() *string {
	return s.EntityDomain
}

func (s *ObserveGroupEntitySummary) GetEntityType() *string {
	return s.EntityType
}

func (s *ObserveGroupEntitySummary) GetMainCount() *int64 {
	return s.MainCount
}

func (s *ObserveGroupEntitySummary) GetRelatedCount() *int64 {
	return s.RelatedCount
}

func (s *ObserveGroupEntitySummary) SetEntityCategory(v string) *ObserveGroupEntitySummary {
	s.EntityCategory = &v
	return s
}

func (s *ObserveGroupEntitySummary) SetEntityCount(v int32) *ObserveGroupEntitySummary {
	s.EntityCount = &v
	return s
}

func (s *ObserveGroupEntitySummary) SetEntityDomain(v string) *ObserveGroupEntitySummary {
	s.EntityDomain = &v
	return s
}

func (s *ObserveGroupEntitySummary) SetEntityType(v string) *ObserveGroupEntitySummary {
	s.EntityType = &v
	return s
}

func (s *ObserveGroupEntitySummary) SetMainCount(v int64) *ObserveGroupEntitySummary {
	s.MainCount = &v
	return s
}

func (s *ObserveGroupEntitySummary) SetRelatedCount(v int64) *ObserveGroupEntitySummary {
	s.RelatedCount = &v
	return s
}

func (s *ObserveGroupEntitySummary) Validate() error {
	return dara.Validate(s)
}
