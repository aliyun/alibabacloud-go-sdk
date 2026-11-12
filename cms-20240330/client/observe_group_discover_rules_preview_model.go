// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObserveGroupDiscoverRulesPreview interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogRevision(v string) *ObserveGroupDiscoverRulesPreview
	GetCatalogRevision() *string
	SetCount(v int32) *ObserveGroupDiscoverRulesPreview
	GetCount() *int32
	SetEntitySummaries(v []*ObserveGroupEntitySummary) *ObserveGroupDiscoverRulesPreview
	GetEntitySummaries() []*ObserveGroupEntitySummary
	SetFilteredCount(v int32) *ObserveGroupDiscoverRulesPreview
	GetFilteredCount() *int32
	SetInstances(v []*ObserveGroupInstance) *ObserveGroupDiscoverRulesPreview
	GetInstances() []*ObserveGroupInstance
	SetMainCount(v int64) *ObserveGroupDiscoverRulesPreview
	GetMainCount() *int64
	SetMaxResults(v int32) *ObserveGroupDiscoverRulesPreview
	GetMaxResults() *int32
	SetNextToken(v string) *ObserveGroupDiscoverRulesPreview
	GetNextToken() *string
	SetRelatedCount(v int64) *ObserveGroupDiscoverRulesPreview
	GetRelatedCount() *int64
	SetRuleCounts(v []*ObserveGroupRuleCount) *ObserveGroupDiscoverRulesPreview
	GetRuleCounts() []*ObserveGroupRuleCount
	SetTotalCount(v int64) *ObserveGroupDiscoverRulesPreview
	GetTotalCount() *int64
}

type ObserveGroupDiscoverRulesPreview struct {
	// The global association catalog version used for this preview.
	//
	// example:
	//
	// 03ac8b5b92cb9c2675119d64c31d2520c069647aa9b88979df1bb115472c02b4
	CatalogRevision *string `json:"catalogRevision,omitempty" xml:"catalogRevision,omitempty"`
	// The number of entities returned on this page.
	//
	// example:
	//
	// 10
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The primary entities, associated sub-entities, and total count by entity type.
	EntitySummaries []*ObserveGroupEntitySummary `json:"entitySummaries,omitempty" xml:"entitySummaries,omitempty" type:"Repeated"`
	// The total count after filtering by entity role, type, and keyword.
	//
	// example:
	//
	// 12
	FilteredCount *int32 `json:"filteredCount,omitempty" xml:"filteredCount,omitempty"`
	// The list of matched entities on this page.
	Instances []*ObserveGroupInstance `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
	// The deduplicated count of primary entities.
	//
	// example:
	//
	// 3
	MainCount *int64 `json:"mainCount,omitempty" xml:"mainCount,omitempty"`
	// The maximum number of entities returned on this page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The next page token, which is empty if there is no next page.
	//
	// example:
	//
	// eyJvZmZzZXQiOjEwLCJyZXF1ZXN0SGFzaCI6IjAzYWM4YjViOTJjYjljMjY3NTExOWQ2NGMzMWQyNTIwYzA2OTY0N2FhOWI4ODk3OWRmMWJiMTE1NDcyYzAyYjQifQ==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The deduplicated count of associated sub-entities.
	//
	// example:
	//
	// 9
	RelatedCount *int64 `json:"relatedCount,omitempty" xml:"relatedCount,omitempty"`
	// The primary entities, associated sub-entities, and total count by rule.
	RuleCounts []*ObserveGroupRuleCount `json:"ruleCounts,omitempty" xml:"ruleCounts,omitempty" type:"Repeated"`
	// The deduplicated total count of entities.
	//
	// example:
	//
	// 12
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ObserveGroupDiscoverRulesPreview) String() string {
	return dara.Prettify(s)
}

func (s ObserveGroupDiscoverRulesPreview) GoString() string {
	return s.String()
}

func (s *ObserveGroupDiscoverRulesPreview) GetCatalogRevision() *string {
	return s.CatalogRevision
}

func (s *ObserveGroupDiscoverRulesPreview) GetCount() *int32 {
	return s.Count
}

func (s *ObserveGroupDiscoverRulesPreview) GetEntitySummaries() []*ObserveGroupEntitySummary {
	return s.EntitySummaries
}

func (s *ObserveGroupDiscoverRulesPreview) GetFilteredCount() *int32 {
	return s.FilteredCount
}

func (s *ObserveGroupDiscoverRulesPreview) GetInstances() []*ObserveGroupInstance {
	return s.Instances
}

func (s *ObserveGroupDiscoverRulesPreview) GetMainCount() *int64 {
	return s.MainCount
}

func (s *ObserveGroupDiscoverRulesPreview) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ObserveGroupDiscoverRulesPreview) GetNextToken() *string {
	return s.NextToken
}

func (s *ObserveGroupDiscoverRulesPreview) GetRelatedCount() *int64 {
	return s.RelatedCount
}

func (s *ObserveGroupDiscoverRulesPreview) GetRuleCounts() []*ObserveGroupRuleCount {
	return s.RuleCounts
}

func (s *ObserveGroupDiscoverRulesPreview) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ObserveGroupDiscoverRulesPreview) SetCatalogRevision(v string) *ObserveGroupDiscoverRulesPreview {
	s.CatalogRevision = &v
	return s
}

func (s *ObserveGroupDiscoverRulesPreview) SetCount(v int32) *ObserveGroupDiscoverRulesPreview {
	s.Count = &v
	return s
}

func (s *ObserveGroupDiscoverRulesPreview) SetEntitySummaries(v []*ObserveGroupEntitySummary) *ObserveGroupDiscoverRulesPreview {
	s.EntitySummaries = v
	return s
}

func (s *ObserveGroupDiscoverRulesPreview) SetFilteredCount(v int32) *ObserveGroupDiscoverRulesPreview {
	s.FilteredCount = &v
	return s
}

func (s *ObserveGroupDiscoverRulesPreview) SetInstances(v []*ObserveGroupInstance) *ObserveGroupDiscoverRulesPreview {
	s.Instances = v
	return s
}

func (s *ObserveGroupDiscoverRulesPreview) SetMainCount(v int64) *ObserveGroupDiscoverRulesPreview {
	s.MainCount = &v
	return s
}

func (s *ObserveGroupDiscoverRulesPreview) SetMaxResults(v int32) *ObserveGroupDiscoverRulesPreview {
	s.MaxResults = &v
	return s
}

func (s *ObserveGroupDiscoverRulesPreview) SetNextToken(v string) *ObserveGroupDiscoverRulesPreview {
	s.NextToken = &v
	return s
}

func (s *ObserveGroupDiscoverRulesPreview) SetRelatedCount(v int64) *ObserveGroupDiscoverRulesPreview {
	s.RelatedCount = &v
	return s
}

func (s *ObserveGroupDiscoverRulesPreview) SetRuleCounts(v []*ObserveGroupRuleCount) *ObserveGroupDiscoverRulesPreview {
	s.RuleCounts = v
	return s
}

func (s *ObserveGroupDiscoverRulesPreview) SetTotalCount(v int64) *ObserveGroupDiscoverRulesPreview {
	s.TotalCount = &v
	return s
}

func (s *ObserveGroupDiscoverRulesPreview) Validate() error {
	if s.EntitySummaries != nil {
		for _, item := range s.EntitySummaries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RuleCounts != nil {
		for _, item := range s.RuleCounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
