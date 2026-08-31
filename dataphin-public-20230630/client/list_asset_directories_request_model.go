// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssetDirectoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListAssetDirectoriesRequestListQuery) *ListAssetDirectoriesRequest
	GetListQuery() *ListAssetDirectoriesRequestListQuery
	SetOpTenantId(v int64) *ListAssetDirectoriesRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListAssetDirectoriesRequest
	GetOpUserId() *string
}

type ListAssetDirectoriesRequest struct {
	// The query parameters.
	//
	// This parameter is required.
	ListQuery *ListAssetDirectoriesRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListAssetDirectoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAssetDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *ListAssetDirectoriesRequest) GetListQuery() *ListAssetDirectoriesRequestListQuery {
	return s.ListQuery
}

func (s *ListAssetDirectoriesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListAssetDirectoriesRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListAssetDirectoriesRequest) SetListQuery(v *ListAssetDirectoriesRequestListQuery) *ListAssetDirectoriesRequest {
	s.ListQuery = v
	return s
}

func (s *ListAssetDirectoriesRequest) SetOpTenantId(v int64) *ListAssetDirectoriesRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListAssetDirectoriesRequest) SetOpUserId(v string) *ListAssetDirectoriesRequest {
	s.OpUserId = &v
	return s
}

func (s *ListAssetDirectoriesRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAssetDirectoriesRequestListQuery struct {
	// The folder name keyword. Maximum length: 128 characters.
	//
	// example:
	//
	// Core Metrics
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The maximum number of expansion levels. This parameter takes effect only in browse mode. Valid values: 1 to 10.
	//
	// example:
	//
	// 6470568
	MaxLevel *int32 `json:"MaxLevel,omitempty" xml:"MaxLevel,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 696844
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page. Default value: 50. Valid values: 1 to 200.
	//
	// example:
	//
	// 7576639
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The parent folder ID. This parameter takes effect only in browse mode.
	//
	// example:
	//
	// 466096149777
	ParentDirectoryId *int64 `json:"ParentDirectoryId,omitempty" xml:"ParentDirectoryId,omitempty"`
	// The topic ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 796027234512
	TopicId *int64 `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
}

func (s ListAssetDirectoriesRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListAssetDirectoriesRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListAssetDirectoriesRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListAssetDirectoriesRequestListQuery) GetMaxLevel() *int32 {
	return s.MaxLevel
}

func (s *ListAssetDirectoriesRequestListQuery) GetPage() *int32 {
	return s.Page
}

func (s *ListAssetDirectoriesRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAssetDirectoriesRequestListQuery) GetParentDirectoryId() *int64 {
	return s.ParentDirectoryId
}

func (s *ListAssetDirectoriesRequestListQuery) GetTopicId() *int64 {
	return s.TopicId
}

func (s *ListAssetDirectoriesRequestListQuery) SetKeyword(v string) *ListAssetDirectoriesRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListAssetDirectoriesRequestListQuery) SetMaxLevel(v int32) *ListAssetDirectoriesRequestListQuery {
	s.MaxLevel = &v
	return s
}

func (s *ListAssetDirectoriesRequestListQuery) SetPage(v int32) *ListAssetDirectoriesRequestListQuery {
	s.Page = &v
	return s
}

func (s *ListAssetDirectoriesRequestListQuery) SetPageSize(v int32) *ListAssetDirectoriesRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListAssetDirectoriesRequestListQuery) SetParentDirectoryId(v int64) *ListAssetDirectoriesRequestListQuery {
	s.ParentDirectoryId = &v
	return s
}

func (s *ListAssetDirectoriesRequestListQuery) SetTopicId(v int64) *ListAssetDirectoriesRequestListQuery {
	s.TopicId = &v
	return s
}

func (s *ListAssetDirectoriesRequestListQuery) Validate() error {
	return dara.Validate(s)
}
