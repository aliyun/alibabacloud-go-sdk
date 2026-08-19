// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCrawlersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceIds(v []*int64) *ListCrawlersRequest
	GetDataSourceIds() []*int64
	SetEnvType(v string) *ListCrawlersRequest
	GetEnvType() *string
	SetName(v string) *ListCrawlersRequest
	GetName() *string
	SetOwner(v string) *ListCrawlersRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListCrawlersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCrawlersRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListCrawlersRequest
	GetProjectId() *int64
	SetType(v string) *ListCrawlersRequest
	GetType() *string
}

type ListCrawlersRequest struct {
	// The list of data source IDs. Up to 10 IDs are supported.
	DataSourceIds []*int64 `json:"DataSourceIds,omitempty" xml:"DataSourceIds,omitempty" type:"Repeated"`
	// The DataWorks environment type. Dev indicates the development environment. Prod indicates the production environment.
	//
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The metadata crawler name. Supports fuzzy match.
	//
	// example:
	//
	// example_crawler
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The DataWorks user ID of the crawler owner.
	//
	// example:
	//
	// 1000
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The page number. Starts from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The crawler type. Call GetCrawlerTypeCapabilities to query the valid values supported in the current region.
	//
	// example:
	//
	// starrocks
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListCrawlersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCrawlersRequest) GoString() string {
	return s.String()
}

func (s *ListCrawlersRequest) GetDataSourceIds() []*int64 {
	return s.DataSourceIds
}

func (s *ListCrawlersRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *ListCrawlersRequest) GetName() *string {
	return s.Name
}

func (s *ListCrawlersRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListCrawlersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCrawlersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCrawlersRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListCrawlersRequest) GetType() *string {
	return s.Type
}

func (s *ListCrawlersRequest) SetDataSourceIds(v []*int64) *ListCrawlersRequest {
	s.DataSourceIds = v
	return s
}

func (s *ListCrawlersRequest) SetEnvType(v string) *ListCrawlersRequest {
	s.EnvType = &v
	return s
}

func (s *ListCrawlersRequest) SetName(v string) *ListCrawlersRequest {
	s.Name = &v
	return s
}

func (s *ListCrawlersRequest) SetOwner(v string) *ListCrawlersRequest {
	s.Owner = &v
	return s
}

func (s *ListCrawlersRequest) SetPageNumber(v int32) *ListCrawlersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCrawlersRequest) SetPageSize(v int32) *ListCrawlersRequest {
	s.PageSize = &v
	return s
}

func (s *ListCrawlersRequest) SetProjectId(v int64) *ListCrawlersRequest {
	s.ProjectId = &v
	return s
}

func (s *ListCrawlersRequest) SetType(v string) *ListCrawlersRequest {
	s.Type = &v
	return s
}

func (s *ListCrawlersRequest) Validate() error {
	return dara.Validate(s)
}
