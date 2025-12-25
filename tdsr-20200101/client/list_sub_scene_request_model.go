// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNum(v int64) *ListSubSceneRequest
	GetPageNum() *int64
	SetPageSize(v int64) *ListSubSceneRequest
	GetPageSize() *int64
	SetSceneId(v string) *ListSubSceneRequest
	GetSceneId() *string
	SetShowLayoutData(v bool) *ListSubSceneRequest
	GetShowLayoutData() *bool
	SetSortField(v string) *ListSubSceneRequest
	GetSortField() *string
}

type ListSubSceneRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// L2omaCMmQZZkEg4pE****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// true
	ShowLayoutData *bool `json:"ShowLayoutData,omitempty" xml:"ShowLayoutData,omitempty"`
	// example:
	//
	// NAME
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
}

func (s ListSubSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSubSceneRequest) GoString() string {
	return s.String()
}

func (s *ListSubSceneRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *ListSubSceneRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSubSceneRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *ListSubSceneRequest) GetShowLayoutData() *bool {
	return s.ShowLayoutData
}

func (s *ListSubSceneRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListSubSceneRequest) SetPageNum(v int64) *ListSubSceneRequest {
	s.PageNum = &v
	return s
}

func (s *ListSubSceneRequest) SetPageSize(v int64) *ListSubSceneRequest {
	s.PageSize = &v
	return s
}

func (s *ListSubSceneRequest) SetSceneId(v string) *ListSubSceneRequest {
	s.SceneId = &v
	return s
}

func (s *ListSubSceneRequest) SetShowLayoutData(v bool) *ListSubSceneRequest {
	s.ShowLayoutData = &v
	return s
}

func (s *ListSubSceneRequest) SetSortField(v string) *ListSubSceneRequest {
	s.SortField = &v
	return s
}

func (s *ListSubSceneRequest) Validate() error {
	return dara.Validate(s)
}
