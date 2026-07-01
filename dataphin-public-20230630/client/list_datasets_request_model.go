// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetQuery(v *ListDatasetsRequestDatasetQuery) *ListDatasetsRequest
	GetDatasetQuery() *ListDatasetsRequestDatasetQuery
	SetOpTenantId(v int64) *ListDatasetsRequest
	GetOpTenantId() *int64
}

type ListDatasetsRequest struct {
	// The request body.
	DatasetQuery *ListDatasetsRequestDatasetQuery `json:"DatasetQuery,omitempty" xml:"DatasetQuery,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListDatasetsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsRequest) GoString() string {
	return s.String()
}

func (s *ListDatasetsRequest) GetDatasetQuery() *ListDatasetsRequestDatasetQuery {
	return s.DatasetQuery
}

func (s *ListDatasetsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListDatasetsRequest) SetDatasetQuery(v *ListDatasetsRequestDatasetQuery) *ListDatasetsRequest {
	s.DatasetQuery = v
	return s
}

func (s *ListDatasetsRequest) SetOpTenantId(v int64) *ListDatasetsRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListDatasetsRequest) Validate() error {
	if s.DatasetQuery != nil {
		if err := s.DatasetQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDatasetsRequestDatasetQuery struct {
	// The content type.
	//
	// example:
	//
	// GENERAL
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The data domain ID.
	//
	// example:
	//
	// 74280
	DataUnit *string `json:"DataUnit,omitempty" xml:"DataUnit,omitempty"`
	// Specifies whether to include version details. Default value: FALSE.
	IncludeVersionList *bool `json:"IncludeVersionList,omitempty" xml:"IncludeVersionList,omitempty"`
	// The keyword.
	//
	// example:
	//
	// rds
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The owner ID.
	//
	// example:
	//
	// 300001391
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7273382541481536
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The dataset scenario. Valid values:
	//
	// - OFFLINE: offline.
	//
	// - REALTIME: real-time.
	//
	// example:
	//
	// OFFLINE
	Scenario *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	// The storage type.
	//
	// example:
	//
	// OSS
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 30001011
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The dataset type. If left empty, all types are queried.
	TypeList []*string `json:"TypeList,omitempty" xml:"TypeList,omitempty" type:"Repeated"`
}

func (s ListDatasetsRequestDatasetQuery) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsRequestDatasetQuery) GoString() string {
	return s.String()
}

func (s *ListDatasetsRequestDatasetQuery) GetContentType() *string {
	return s.ContentType
}

func (s *ListDatasetsRequestDatasetQuery) GetDataUnit() *string {
	return s.DataUnit
}

func (s *ListDatasetsRequestDatasetQuery) GetIncludeVersionList() *bool {
	return s.IncludeVersionList
}

func (s *ListDatasetsRequestDatasetQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListDatasetsRequestDatasetQuery) GetOwner() *string {
	return s.Owner
}

func (s *ListDatasetsRequestDatasetQuery) GetPage() *int32 {
	return s.Page
}

func (s *ListDatasetsRequestDatasetQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDatasetsRequestDatasetQuery) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDatasetsRequestDatasetQuery) GetScenario() *string {
	return s.Scenario
}

func (s *ListDatasetsRequestDatasetQuery) GetStorageType() *string {
	return s.StorageType
}

func (s *ListDatasetsRequestDatasetQuery) GetTenantId() *int64 {
	return s.TenantId
}

func (s *ListDatasetsRequestDatasetQuery) GetTypeList() []*string {
	return s.TypeList
}

func (s *ListDatasetsRequestDatasetQuery) SetContentType(v string) *ListDatasetsRequestDatasetQuery {
	s.ContentType = &v
	return s
}

func (s *ListDatasetsRequestDatasetQuery) SetDataUnit(v string) *ListDatasetsRequestDatasetQuery {
	s.DataUnit = &v
	return s
}

func (s *ListDatasetsRequestDatasetQuery) SetIncludeVersionList(v bool) *ListDatasetsRequestDatasetQuery {
	s.IncludeVersionList = &v
	return s
}

func (s *ListDatasetsRequestDatasetQuery) SetKeyword(v string) *ListDatasetsRequestDatasetQuery {
	s.Keyword = &v
	return s
}

func (s *ListDatasetsRequestDatasetQuery) SetOwner(v string) *ListDatasetsRequestDatasetQuery {
	s.Owner = &v
	return s
}

func (s *ListDatasetsRequestDatasetQuery) SetPage(v int32) *ListDatasetsRequestDatasetQuery {
	s.Page = &v
	return s
}

func (s *ListDatasetsRequestDatasetQuery) SetPageSize(v int32) *ListDatasetsRequestDatasetQuery {
	s.PageSize = &v
	return s
}

func (s *ListDatasetsRequestDatasetQuery) SetProjectId(v int64) *ListDatasetsRequestDatasetQuery {
	s.ProjectId = &v
	return s
}

func (s *ListDatasetsRequestDatasetQuery) SetScenario(v string) *ListDatasetsRequestDatasetQuery {
	s.Scenario = &v
	return s
}

func (s *ListDatasetsRequestDatasetQuery) SetStorageType(v string) *ListDatasetsRequestDatasetQuery {
	s.StorageType = &v
	return s
}

func (s *ListDatasetsRequestDatasetQuery) SetTenantId(v int64) *ListDatasetsRequestDatasetQuery {
	s.TenantId = &v
	return s
}

func (s *ListDatasetsRequestDatasetQuery) SetTypeList(v []*string) *ListDatasetsRequestDatasetQuery {
	s.TypeList = v
	return s
}

func (s *ListDatasetsRequestDatasetQuery) Validate() error {
	return dara.Validate(s)
}
