// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCrawlersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceIdsShrink(v string) *ListCrawlersShrinkRequest
	GetDataSourceIdsShrink() *string
	SetEnvType(v string) *ListCrawlersShrinkRequest
	GetEnvType() *string
	SetName(v string) *ListCrawlersShrinkRequest
	GetName() *string
	SetOwner(v string) *ListCrawlersShrinkRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListCrawlersShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCrawlersShrinkRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListCrawlersShrinkRequest
	GetProjectId() *int64
	SetType(v string) *ListCrawlersShrinkRequest
	GetType() *string
}

type ListCrawlersShrinkRequest struct {
	DataSourceIdsShrink *string `json:"DataSourceIds,omitempty" xml:"DataSourceIds,omitempty"`
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// example:
	//
	// example_crawler
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1000
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// starrocks
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListCrawlersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCrawlersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListCrawlersShrinkRequest) GetDataSourceIdsShrink() *string {
	return s.DataSourceIdsShrink
}

func (s *ListCrawlersShrinkRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *ListCrawlersShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListCrawlersShrinkRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListCrawlersShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCrawlersShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCrawlersShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListCrawlersShrinkRequest) GetType() *string {
	return s.Type
}

func (s *ListCrawlersShrinkRequest) SetDataSourceIdsShrink(v string) *ListCrawlersShrinkRequest {
	s.DataSourceIdsShrink = &v
	return s
}

func (s *ListCrawlersShrinkRequest) SetEnvType(v string) *ListCrawlersShrinkRequest {
	s.EnvType = &v
	return s
}

func (s *ListCrawlersShrinkRequest) SetName(v string) *ListCrawlersShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListCrawlersShrinkRequest) SetOwner(v string) *ListCrawlersShrinkRequest {
	s.Owner = &v
	return s
}

func (s *ListCrawlersShrinkRequest) SetPageNumber(v int32) *ListCrawlersShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCrawlersShrinkRequest) SetPageSize(v int32) *ListCrawlersShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListCrawlersShrinkRequest) SetProjectId(v int64) *ListCrawlersShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListCrawlersShrinkRequest) SetType(v string) *ListCrawlersShrinkRequest {
	s.Type = &v
	return s
}

func (s *ListCrawlersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
