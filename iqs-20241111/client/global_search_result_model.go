// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalSearchResult interface {
	dara.Model
	String() string
	GoString() string
	SetPageItems(v []*GlobalPageItem) *GlobalSearchResult
	GetPageItems() []*GlobalPageItem
	SetQueryContext(v *GlobalQueryContext) *GlobalSearchResult
	GetQueryContext() *GlobalQueryContext
	SetRequestId(v string) *GlobalSearchResult
	GetRequestId() *string
	SetSceneItems(v []*GlobalSceneItem) *GlobalSearchResult
	GetSceneItems() []*GlobalSceneItem
	SetSearchInformation(v *GlobalSearchInformation) *GlobalSearchResult
	GetSearchInformation() *GlobalSearchInformation
}

type GlobalSearchResult struct {
	PageItems    []*GlobalPageItem   `json:"pageItems,omitempty" xml:"pageItems,omitempty" type:"Repeated"`
	QueryContext *GlobalQueryContext `json:"queryContext,omitempty" xml:"queryContext,omitempty"`
	// example:
	//
	// 123456
	RequestId         *string                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SceneItems        []*GlobalSceneItem       `json:"sceneItems,omitempty" xml:"sceneItems,omitempty" type:"Repeated"`
	SearchInformation *GlobalSearchInformation `json:"searchInformation,omitempty" xml:"searchInformation,omitempty"`
}

func (s GlobalSearchResult) String() string {
	return dara.Prettify(s)
}

func (s GlobalSearchResult) GoString() string {
	return s.String()
}

func (s *GlobalSearchResult) GetPageItems() []*GlobalPageItem {
	return s.PageItems
}

func (s *GlobalSearchResult) GetQueryContext() *GlobalQueryContext {
	return s.QueryContext
}

func (s *GlobalSearchResult) GetRequestId() *string {
	return s.RequestId
}

func (s *GlobalSearchResult) GetSceneItems() []*GlobalSceneItem {
	return s.SceneItems
}

func (s *GlobalSearchResult) GetSearchInformation() *GlobalSearchInformation {
	return s.SearchInformation
}

func (s *GlobalSearchResult) SetPageItems(v []*GlobalPageItem) *GlobalSearchResult {
	s.PageItems = v
	return s
}

func (s *GlobalSearchResult) SetQueryContext(v *GlobalQueryContext) *GlobalSearchResult {
	s.QueryContext = v
	return s
}

func (s *GlobalSearchResult) SetRequestId(v string) *GlobalSearchResult {
	s.RequestId = &v
	return s
}

func (s *GlobalSearchResult) SetSceneItems(v []*GlobalSceneItem) *GlobalSearchResult {
	s.SceneItems = v
	return s
}

func (s *GlobalSearchResult) SetSearchInformation(v *GlobalSearchInformation) *GlobalSearchResult {
	s.SearchInformation = v
	return s
}

func (s *GlobalSearchResult) Validate() error {
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
