// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecurityIdentifyResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListSecurityIdentifyResultsRequestListQuery) *ListSecurityIdentifyResultsRequest
	GetListQuery() *ListSecurityIdentifyResultsRequestListQuery
	SetOpTenantId(v int64) *ListSecurityIdentifyResultsRequest
	GetOpTenantId() *int64
}

type ListSecurityIdentifyResultsRequest struct {
	// Query conditions.
	ListQuery *ListSecurityIdentifyResultsRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// Tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListSecurityIdentifyResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityIdentifyResultsRequest) GoString() string {
	return s.String()
}

func (s *ListSecurityIdentifyResultsRequest) GetListQuery() *ListSecurityIdentifyResultsRequestListQuery {
	return s.ListQuery
}

func (s *ListSecurityIdentifyResultsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListSecurityIdentifyResultsRequest) SetListQuery(v *ListSecurityIdentifyResultsRequestListQuery) *ListSecurityIdentifyResultsRequest {
	s.ListQuery = v
	return s
}

func (s *ListSecurityIdentifyResultsRequest) SetOpTenantId(v int64) *ListSecurityIdentifyResultsRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListSecurityIdentifyResultsRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSecurityIdentifyResultsRequestListQuery struct {
	// Business unit names.
	BizUnitNameList []*string `json:"BizUnitNameList,omitempty" xml:"BizUnitNameList,omitempty" type:"Repeated"`
	// Data classification ID.
	//
	// example:
	//
	// 1
	ClassifyId *int64 `json:"ClassifyId,omitempty" xml:"ClassifyId,omitempty"`
	// Data source names.
	DatasourceNameList []*string `json:"DatasourceNameList,omitempty" xml:"DatasourceNameList,omitempty" type:"Repeated"`
	// Environment identifier of the data source, project, or business unit.
	//
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// Whether it is locked.
	IsLocked *bool `json:"IsLocked,omitempty" xml:"IsLocked,omitempty"`
	// Keyword filter. Supports matching by table catalog, table name, table display name, and field name.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// Page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Number of records per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Project names.
	ProjectNameList []*string `json:"ProjectNameList,omitempty" xml:"ProjectNameList,omitempty" type:"Repeated"`
	// Effective status filter: ENABLE or DISABLE.
	//
	// example:
	//
	// ENABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListSecurityIdentifyResultsRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityIdentifyResultsRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListSecurityIdentifyResultsRequestListQuery) GetBizUnitNameList() []*string {
	return s.BizUnitNameList
}

func (s *ListSecurityIdentifyResultsRequestListQuery) GetClassifyId() *int64 {
	return s.ClassifyId
}

func (s *ListSecurityIdentifyResultsRequestListQuery) GetDatasourceNameList() []*string {
	return s.DatasourceNameList
}

func (s *ListSecurityIdentifyResultsRequestListQuery) GetEnv() *string {
	return s.Env
}

func (s *ListSecurityIdentifyResultsRequestListQuery) GetIsLocked() *bool {
	return s.IsLocked
}

func (s *ListSecurityIdentifyResultsRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListSecurityIdentifyResultsRequestListQuery) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListSecurityIdentifyResultsRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSecurityIdentifyResultsRequestListQuery) GetProjectNameList() []*string {
	return s.ProjectNameList
}

func (s *ListSecurityIdentifyResultsRequestListQuery) GetStatus() *string {
	return s.Status
}

func (s *ListSecurityIdentifyResultsRequestListQuery) SetBizUnitNameList(v []*string) *ListSecurityIdentifyResultsRequestListQuery {
	s.BizUnitNameList = v
	return s
}

func (s *ListSecurityIdentifyResultsRequestListQuery) SetClassifyId(v int64) *ListSecurityIdentifyResultsRequestListQuery {
	s.ClassifyId = &v
	return s
}

func (s *ListSecurityIdentifyResultsRequestListQuery) SetDatasourceNameList(v []*string) *ListSecurityIdentifyResultsRequestListQuery {
	s.DatasourceNameList = v
	return s
}

func (s *ListSecurityIdentifyResultsRequestListQuery) SetEnv(v string) *ListSecurityIdentifyResultsRequestListQuery {
	s.Env = &v
	return s
}

func (s *ListSecurityIdentifyResultsRequestListQuery) SetIsLocked(v bool) *ListSecurityIdentifyResultsRequestListQuery {
	s.IsLocked = &v
	return s
}

func (s *ListSecurityIdentifyResultsRequestListQuery) SetKeyword(v string) *ListSecurityIdentifyResultsRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListSecurityIdentifyResultsRequestListQuery) SetPageNo(v int32) *ListSecurityIdentifyResultsRequestListQuery {
	s.PageNo = &v
	return s
}

func (s *ListSecurityIdentifyResultsRequestListQuery) SetPageSize(v int32) *ListSecurityIdentifyResultsRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListSecurityIdentifyResultsRequestListQuery) SetProjectNameList(v []*string) *ListSecurityIdentifyResultsRequestListQuery {
	s.ProjectNameList = v
	return s
}

func (s *ListSecurityIdentifyResultsRequestListQuery) SetStatus(v string) *ListSecurityIdentifyResultsRequestListQuery {
	s.Status = &v
	return s
}

func (s *ListSecurityIdentifyResultsRequestListQuery) Validate() error {
	return dara.Validate(s)
}
