// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmsDataSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListMmsDataSourcesRequest
	GetName() *string
	SetPageNum(v int32) *ListMmsDataSourcesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListMmsDataSourcesRequest
	GetPageSize() *int32
	SetRegion(v string) *ListMmsDataSourcesRequest
	GetRegion() *string
	SetType(v string) *ListMmsDataSourcesRequest
	GetType() *string
}

type ListMmsDataSourcesRequest struct {
	// The name of the data source.
	//
	// example:
	//
	// demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The page number. If you leave this parameter empty, all data sources are returned.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// The number of entries to return on each page. If you leave this parameter empty, all data sources are returned.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The type of the data source.
	//
	// example:
	//
	// BIGQUERY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListMmsDataSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMmsDataSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListMmsDataSourcesRequest) GetName() *string {
	return s.Name
}

func (s *ListMmsDataSourcesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMmsDataSourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMmsDataSourcesRequest) GetRegion() *string {
	return s.Region
}

func (s *ListMmsDataSourcesRequest) GetType() *string {
	return s.Type
}

func (s *ListMmsDataSourcesRequest) SetName(v string) *ListMmsDataSourcesRequest {
	s.Name = &v
	return s
}

func (s *ListMmsDataSourcesRequest) SetPageNum(v int32) *ListMmsDataSourcesRequest {
	s.PageNum = &v
	return s
}

func (s *ListMmsDataSourcesRequest) SetPageSize(v int32) *ListMmsDataSourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListMmsDataSourcesRequest) SetRegion(v string) *ListMmsDataSourcesRequest {
	s.Region = &v
	return s
}

func (s *ListMmsDataSourcesRequest) SetType(v string) *ListMmsDataSourcesRequest {
	s.Type = &v
	return s
}

func (s *ListMmsDataSourcesRequest) Validate() error {
	return dara.Validate(s)
}
