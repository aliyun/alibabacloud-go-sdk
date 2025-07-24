// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenericSearchResult interface {
	dara.Model
	String() string
	GoString() string
	SetPageItems(v []*ScorePageItem) *GenericSearchResult
	GetPageItems() []*ScorePageItem
	SetQueryContext(v *QueryContext) *GenericSearchResult
	GetQueryContext() *QueryContext
	SetRequestId(v string) *GenericSearchResult
	GetRequestId() *string
	SetSceneItems(v []*SceneItem) *GenericSearchResult
	GetSceneItems() []*SceneItem
	SetSearchInformation(v *SearchInformation) *GenericSearchResult
	GetSearchInformation() *SearchInformation
	SetWeiboItems(v []*WeiboItem) *GenericSearchResult
	GetWeiboItems() []*WeiboItem
}

type GenericSearchResult struct {
	PageItems    []*ScorePageItem `json:"pageItems,omitempty" xml:"pageItems,omitempty" type:"Repeated"`
	QueryContext *QueryContext    `json:"queryContext,omitempty" xml:"queryContext,omitempty"`
	// example:
	//
	// 123456
	RequestId         *string            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SceneItems        []*SceneItem       `json:"sceneItems,omitempty" xml:"sceneItems,omitempty" type:"Repeated"`
	SearchInformation *SearchInformation `json:"searchInformation,omitempty" xml:"searchInformation,omitempty"`
	WeiboItems        []*WeiboItem       `json:"weiboItems,omitempty" xml:"weiboItems,omitempty" type:"Repeated"`
}

func (s GenericSearchResult) String() string {
	return dara.Prettify(s)
}

func (s GenericSearchResult) GoString() string {
	return s.String()
}

func (s *GenericSearchResult) GetPageItems() []*ScorePageItem {
	return s.PageItems
}

func (s *GenericSearchResult) GetQueryContext() *QueryContext {
	return s.QueryContext
}

func (s *GenericSearchResult) GetRequestId() *string {
	return s.RequestId
}

func (s *GenericSearchResult) GetSceneItems() []*SceneItem {
	return s.SceneItems
}

func (s *GenericSearchResult) GetSearchInformation() *SearchInformation {
	return s.SearchInformation
}

func (s *GenericSearchResult) GetWeiboItems() []*WeiboItem {
	return s.WeiboItems
}

func (s *GenericSearchResult) SetPageItems(v []*ScorePageItem) *GenericSearchResult {
	s.PageItems = v
	return s
}

func (s *GenericSearchResult) SetQueryContext(v *QueryContext) *GenericSearchResult {
	s.QueryContext = v
	return s
}

func (s *GenericSearchResult) SetRequestId(v string) *GenericSearchResult {
	s.RequestId = &v
	return s
}

func (s *GenericSearchResult) SetSceneItems(v []*SceneItem) *GenericSearchResult {
	s.SceneItems = v
	return s
}

func (s *GenericSearchResult) SetSearchInformation(v *SearchInformation) *GenericSearchResult {
	s.SearchInformation = v
	return s
}

func (s *GenericSearchResult) SetWeiboItems(v []*WeiboItem) *GenericSearchResult {
	s.WeiboItems = v
	return s
}

func (s *GenericSearchResult) Validate() error {
	return dara.Validate(s)
}
