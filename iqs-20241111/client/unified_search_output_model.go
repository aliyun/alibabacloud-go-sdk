// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnifiedSearchOutput interface {
	dara.Model
	String() string
	GoString() string
	SetCostCredits(v *UnifiedCostCredits) *UnifiedSearchOutput
	GetCostCredits() *UnifiedCostCredits
	SetPageItems(v []*UnifiedPageItem) *UnifiedSearchOutput
	GetPageItems() []*UnifiedPageItem
	SetQueryContext(v *UnifiedQueryContext) *UnifiedSearchOutput
	GetQueryContext() *UnifiedQueryContext
	SetRequestId(v string) *UnifiedSearchOutput
	GetRequestId() *string
	SetSceneItems(v []*UnifiedSceneItem) *UnifiedSearchOutput
	GetSceneItems() []*UnifiedSceneItem
	SetSearchInformation(v *UnifiedSearchInformation) *UnifiedSearchOutput
	GetSearchInformation() *UnifiedSearchInformation
}

type UnifiedSearchOutput struct {
	CostCredits       *UnifiedCostCredits       `json:"costCredits,omitempty" xml:"costCredits,omitempty"`
	PageItems         []*UnifiedPageItem        `json:"pageItems,omitempty" xml:"pageItems,omitempty" type:"Repeated"`
	QueryContext      *UnifiedQueryContext      `json:"queryContext,omitempty" xml:"queryContext,omitempty"`
	RequestId         *string                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SceneItems        []*UnifiedSceneItem       `json:"sceneItems,omitempty" xml:"sceneItems,omitempty" type:"Repeated"`
	SearchInformation *UnifiedSearchInformation `json:"searchInformation,omitempty" xml:"searchInformation,omitempty"`
}

func (s UnifiedSearchOutput) String() string {
	return dara.Prettify(s)
}

func (s UnifiedSearchOutput) GoString() string {
	return s.String()
}

func (s *UnifiedSearchOutput) GetCostCredits() *UnifiedCostCredits {
	return s.CostCredits
}

func (s *UnifiedSearchOutput) GetPageItems() []*UnifiedPageItem {
	return s.PageItems
}

func (s *UnifiedSearchOutput) GetQueryContext() *UnifiedQueryContext {
	return s.QueryContext
}

func (s *UnifiedSearchOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *UnifiedSearchOutput) GetSceneItems() []*UnifiedSceneItem {
	return s.SceneItems
}

func (s *UnifiedSearchOutput) GetSearchInformation() *UnifiedSearchInformation {
	return s.SearchInformation
}

func (s *UnifiedSearchOutput) SetCostCredits(v *UnifiedCostCredits) *UnifiedSearchOutput {
	s.CostCredits = v
	return s
}

func (s *UnifiedSearchOutput) SetPageItems(v []*UnifiedPageItem) *UnifiedSearchOutput {
	s.PageItems = v
	return s
}

func (s *UnifiedSearchOutput) SetQueryContext(v *UnifiedQueryContext) *UnifiedSearchOutput {
	s.QueryContext = v
	return s
}

func (s *UnifiedSearchOutput) SetRequestId(v string) *UnifiedSearchOutput {
	s.RequestId = &v
	return s
}

func (s *UnifiedSearchOutput) SetSceneItems(v []*UnifiedSceneItem) *UnifiedSearchOutput {
	s.SceneItems = v
	return s
}

func (s *UnifiedSearchOutput) SetSearchInformation(v *UnifiedSearchInformation) *UnifiedSearchOutput {
	s.SearchInformation = v
	return s
}

func (s *UnifiedSearchOutput) Validate() error {
	if s.CostCredits != nil {
		if err := s.CostCredits.Validate(); err != nil {
			return err
		}
	}
	if s.PageItems != nil {
		for _, item := range s.PageItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.QueryContext != nil {
		if err := s.QueryContext.Validate(); err != nil {
			return err
		}
	}
	if s.SceneItems != nil {
		for _, item := range s.SceneItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SearchInformation != nil {
		if err := s.SearchInformation.Validate(); err != nil {
			return err
		}
	}
	return nil
}
