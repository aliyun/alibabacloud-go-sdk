// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityArchiveTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListQualityArchiveTablesRequestListQuery) *ListQualityArchiveTablesRequest
	GetListQuery() *ListQualityArchiveTablesRequestListQuery
	SetOpTenantId(v int64) *ListQualityArchiveTablesRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListQualityArchiveTablesRequest
	GetOpUserId() *string
}

type ListQualityArchiveTablesRequest struct {
	// The input parameters for querying the anomaly archived table list.
	//
	// This parameter is required.
	ListQuery *ListQualityArchiveTablesRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
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

func (s ListQualityArchiveTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQualityArchiveTablesRequest) GoString() string {
	return s.String()
}

func (s *ListQualityArchiveTablesRequest) GetListQuery() *ListQualityArchiveTablesRequestListQuery {
	return s.ListQuery
}

func (s *ListQualityArchiveTablesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListQualityArchiveTablesRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListQualityArchiveTablesRequest) SetListQuery(v *ListQualityArchiveTablesRequestListQuery) *ListQualityArchiveTablesRequest {
	s.ListQuery = v
	return s
}

func (s *ListQualityArchiveTablesRequest) SetOpTenantId(v int64) *ListQualityArchiveTablesRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListQualityArchiveTablesRequest) SetOpUserId(v string) *ListQualityArchiveTablesRequest {
	s.OpUserId = &v
	return s
}

func (s *ListQualityArchiveTablesRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListQualityArchiveTablesRequestListQuery struct {
	// The ID of the monitored object.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	WatchId *int64 `json:"WatchId,omitempty" xml:"WatchId,omitempty"`
}

func (s ListQualityArchiveTablesRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListQualityArchiveTablesRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListQualityArchiveTablesRequestListQuery) GetWatchId() *int64 {
	return s.WatchId
}

func (s *ListQualityArchiveTablesRequestListQuery) SetWatchId(v int64) *ListQualityArchiveTablesRequestListQuery {
	s.WatchId = &v
	return s
}

func (s *ListQualityArchiveTablesRequestListQuery) Validate() error {
	return dara.Validate(s)
}
