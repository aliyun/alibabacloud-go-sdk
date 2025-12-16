// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchStrategy interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *SearchStrategy
	GetDescription() *string
	SetIsDefault(v bool) *SearchStrategy
	GetIsDefault() *bool
	SetMergeConfig(v *SearchStrategyMergeConfig) *SearchStrategy
	GetMergeConfig() *SearchStrategyMergeConfig
	SetName(v string) *SearchStrategy
	GetName() *string
	SetSearchConfigs(v []*SearchStrategySearchConfigs) *SearchStrategy
	GetSearchConfigs() []*SearchStrategySearchConfigs
}

type SearchStrategy struct {
	Description   *string                        `json:"description,omitempty" xml:"description,omitempty"`
	IsDefault     *bool                          `json:"isDefault,omitempty" xml:"isDefault,omitempty"`
	MergeConfig   *SearchStrategyMergeConfig     `json:"mergeConfig,omitempty" xml:"mergeConfig,omitempty" type:"Struct"`
	Name          *string                        `json:"name,omitempty" xml:"name,omitempty"`
	SearchConfigs []*SearchStrategySearchConfigs `json:"searchConfigs,omitempty" xml:"searchConfigs,omitempty" type:"Repeated"`
}

func (s SearchStrategy) String() string {
	return dara.Prettify(s)
}

func (s SearchStrategy) GoString() string {
	return s.String()
}

func (s *SearchStrategy) GetDescription() *string {
	return s.Description
}

func (s *SearchStrategy) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *SearchStrategy) GetMergeConfig() *SearchStrategyMergeConfig {
	return s.MergeConfig
}

func (s *SearchStrategy) GetName() *string {
	return s.Name
}

func (s *SearchStrategy) GetSearchConfigs() []*SearchStrategySearchConfigs {
	return s.SearchConfigs
}

func (s *SearchStrategy) SetDescription(v string) *SearchStrategy {
	s.Description = &v
	return s
}

func (s *SearchStrategy) SetIsDefault(v bool) *SearchStrategy {
	s.IsDefault = &v
	return s
}

func (s *SearchStrategy) SetMergeConfig(v *SearchStrategyMergeConfig) *SearchStrategy {
	s.MergeConfig = v
	return s
}

func (s *SearchStrategy) SetName(v string) *SearchStrategy {
	s.Name = &v
	return s
}

func (s *SearchStrategy) SetSearchConfigs(v []*SearchStrategySearchConfigs) *SearchStrategy {
	s.SearchConfigs = v
	return s
}

func (s *SearchStrategy) Validate() error {
	if s.MergeConfig != nil {
		if err := s.MergeConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SearchConfigs != nil {
		for _, item := range s.SearchConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchStrategyMergeConfig struct {
	DocCount *int32  `json:"docCount,omitempty" xml:"docCount,omitempty"`
	RankName *string `json:"rankName,omitempty" xml:"rankName,omitempty"`
}

func (s SearchStrategyMergeConfig) String() string {
	return dara.Prettify(s)
}

func (s SearchStrategyMergeConfig) GoString() string {
	return s.String()
}

func (s *SearchStrategyMergeConfig) GetDocCount() *int32 {
	return s.DocCount
}

func (s *SearchStrategyMergeConfig) GetRankName() *string {
	return s.RankName
}

func (s *SearchStrategyMergeConfig) SetDocCount(v int32) *SearchStrategyMergeConfig {
	s.DocCount = &v
	return s
}

func (s *SearchStrategyMergeConfig) SetRankName(v string) *SearchStrategyMergeConfig {
	s.RankName = &v
	return s
}

func (s *SearchStrategyMergeConfig) Validate() error {
	return dara.Validate(s)
}

type SearchStrategySearchConfigs struct {
	FirstRankName   *string `json:"firstRankName,omitempty" xml:"firstRankName,omitempty"`
	MergeProportion *int32  `json:"mergeProportion,omitempty" xml:"mergeProportion,omitempty"`
	// example:
	//
	// keyword: 关键字查询 vector: 向量查询
	QueryType      *string `json:"queryType,omitempty" xml:"queryType,omitempty"`
	SecondRankName *string `json:"secondRankName,omitempty" xml:"secondRankName,omitempty"`
}

func (s SearchStrategySearchConfigs) String() string {
	return dara.Prettify(s)
}

func (s SearchStrategySearchConfigs) GoString() string {
	return s.String()
}

func (s *SearchStrategySearchConfigs) GetFirstRankName() *string {
	return s.FirstRankName
}

func (s *SearchStrategySearchConfigs) GetMergeProportion() *int32 {
	return s.MergeProportion
}

func (s *SearchStrategySearchConfigs) GetQueryType() *string {
	return s.QueryType
}

func (s *SearchStrategySearchConfigs) GetSecondRankName() *string {
	return s.SecondRankName
}

func (s *SearchStrategySearchConfigs) SetFirstRankName(v string) *SearchStrategySearchConfigs {
	s.FirstRankName = &v
	return s
}

func (s *SearchStrategySearchConfigs) SetMergeProportion(v int32) *SearchStrategySearchConfigs {
	s.MergeProportion = &v
	return s
}

func (s *SearchStrategySearchConfigs) SetQueryType(v string) *SearchStrategySearchConfigs {
	s.QueryType = &v
	return s
}

func (s *SearchStrategySearchConfigs) SetSecondRankName(v string) *SearchStrategySearchConfigs {
	s.SecondRankName = &v
	return s
}

func (s *SearchStrategySearchConfigs) Validate() error {
	return dara.Validate(s)
}
