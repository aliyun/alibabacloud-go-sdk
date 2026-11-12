// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListDatasourcesRequest
	GetClusterId() *string
	SetMaxResults(v int32) *ListDatasourcesRequest
	GetMaxResults() *int32
	SetName(v string) *ListDatasourcesRequest
	GetName() *string
	SetNextToken(v string) *ListDatasourcesRequest
	GetNextToken() *string
	SetPageNum(v int32) *ListDatasourcesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListDatasourcesRequest
	GetPageSize() *int32
	SetType(v int32) *ListDatasourcesRequest
	GetType() *int32
}

type ListDatasourcesRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The maximum number of entries to return for this request. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The task name.
	//
	// example:
	//
	// Corrected voucher statistics
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The token that indicates the position from which the current call returns data. An empty value indicates that all data has been read. You do not need to specify this parameter for the first query.
	//
	// example:
	//
	// eCKqVlS5FKF5EWGGOo8EgQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The Grafana dashboard type. Valid values:
	//
	// - dash-db: Dashboard.
	//
	// - dash-folder: Folder, which can contain dashboards.
	//
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDatasourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatasourcesRequest) GoString() string {
	return s.String()
}

func (s *ListDatasourcesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListDatasourcesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDatasourcesRequest) GetName() *string {
	return s.Name
}

func (s *ListDatasourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDatasourcesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListDatasourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDatasourcesRequest) GetType() *int32 {
	return s.Type
}

func (s *ListDatasourcesRequest) SetClusterId(v string) *ListDatasourcesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListDatasourcesRequest) SetMaxResults(v int32) *ListDatasourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDatasourcesRequest) SetName(v string) *ListDatasourcesRequest {
	s.Name = &v
	return s
}

func (s *ListDatasourcesRequest) SetNextToken(v string) *ListDatasourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListDatasourcesRequest) SetPageNum(v int32) *ListDatasourcesRequest {
	s.PageNum = &v
	return s
}

func (s *ListDatasourcesRequest) SetPageSize(v int32) *ListDatasourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDatasourcesRequest) SetType(v int32) *ListDatasourcesRequest {
	s.Type = &v
	return s
}

func (s *ListDatasourcesRequest) Validate() error {
	return dara.Validate(s)
}
