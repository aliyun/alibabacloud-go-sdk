// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutoDisposeEntitiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoDisposeRecordIds(v []*string) *ListAutoDisposeEntitiesRequest
	GetAutoDisposeRecordIds() []*string
	SetCurrentPage(v string) *ListAutoDisposeEntitiesRequest
	GetCurrentPage() *string
	SetDataSourceType(v string) *ListAutoDisposeEntitiesRequest
	GetDataSourceType() *string
	SetLang(v string) *ListAutoDisposeEntitiesRequest
	GetLang() *string
	SetMaxResults(v int32) *ListAutoDisposeEntitiesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAutoDisposeEntitiesRequest
	GetNextToken() *string
	SetPageSize(v string) *ListAutoDisposeEntitiesRequest
	GetPageSize() *string
	SetUuid(v string) *ListAutoDisposeEntitiesRequest
	GetUuid() *string
}

type ListAutoDisposeEntitiesRequest struct {
	AutoDisposeRecordIds []*string `json:"AutoDisposeRecordIds,omitempty" xml:"AutoDisposeRecordIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// alibaba_cloud_sas
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// b2491e39-ddf2-478a-8c07-*****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListAutoDisposeEntitiesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAutoDisposeEntitiesRequest) GoString() string {
	return s.String()
}

func (s *ListAutoDisposeEntitiesRequest) GetAutoDisposeRecordIds() []*string {
	return s.AutoDisposeRecordIds
}

func (s *ListAutoDisposeEntitiesRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *ListAutoDisposeEntitiesRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *ListAutoDisposeEntitiesRequest) GetLang() *string {
	return s.Lang
}

func (s *ListAutoDisposeEntitiesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAutoDisposeEntitiesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAutoDisposeEntitiesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListAutoDisposeEntitiesRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ListAutoDisposeEntitiesRequest) SetAutoDisposeRecordIds(v []*string) *ListAutoDisposeEntitiesRequest {
	s.AutoDisposeRecordIds = v
	return s
}

func (s *ListAutoDisposeEntitiesRequest) SetCurrentPage(v string) *ListAutoDisposeEntitiesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListAutoDisposeEntitiesRequest) SetDataSourceType(v string) *ListAutoDisposeEntitiesRequest {
	s.DataSourceType = &v
	return s
}

func (s *ListAutoDisposeEntitiesRequest) SetLang(v string) *ListAutoDisposeEntitiesRequest {
	s.Lang = &v
	return s
}

func (s *ListAutoDisposeEntitiesRequest) SetMaxResults(v int32) *ListAutoDisposeEntitiesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAutoDisposeEntitiesRequest) SetNextToken(v string) *ListAutoDisposeEntitiesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAutoDisposeEntitiesRequest) SetPageSize(v string) *ListAutoDisposeEntitiesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAutoDisposeEntitiesRequest) SetUuid(v string) *ListAutoDisposeEntitiesRequest {
	s.Uuid = &v
	return s
}

func (s *ListAutoDisposeEntitiesRequest) Validate() error {
	return dara.Validate(s)
}
