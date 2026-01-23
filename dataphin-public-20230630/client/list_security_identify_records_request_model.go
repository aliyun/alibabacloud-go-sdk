// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecurityIdentifyRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListSecurityIdentifyRecordsRequestListQuery) *ListSecurityIdentifyRecordsRequest
	GetListQuery() *ListSecurityIdentifyRecordsRequestListQuery
	SetOpTenantId(v int64) *ListSecurityIdentifyRecordsRequest
	GetOpTenantId() *int64
}

type ListSecurityIdentifyRecordsRequest struct {
	// This parameter is required.
	ListQuery *ListSecurityIdentifyRecordsRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListSecurityIdentifyRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityIdentifyRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListSecurityIdentifyRecordsRequest) GetListQuery() *ListSecurityIdentifyRecordsRequestListQuery {
	return s.ListQuery
}

func (s *ListSecurityIdentifyRecordsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListSecurityIdentifyRecordsRequest) SetListQuery(v *ListSecurityIdentifyRecordsRequestListQuery) *ListSecurityIdentifyRecordsRequest {
	s.ListQuery = v
	return s
}

func (s *ListSecurityIdentifyRecordsRequest) SetOpTenantId(v int64) *ListSecurityIdentifyRecordsRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListSecurityIdentifyRecordsRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSecurityIdentifyRecordsRequestListQuery struct {
	// example:
	//
	// DEV
	DatasourceEnv *string `json:"DatasourceEnv,omitempty" xml:"DatasourceEnv,omitempty"`
	// example:
	//
	// test
	DatasourceName *string `json:"DatasourceName,omitempty" xml:"DatasourceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// col1
	FieldName         *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	IsDatasourceTable *bool   `json:"IsDatasourceTable,omitempty" xml:"IsDatasourceTable,omitempty"`
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	TableCatalog *string `json:"TableCatalog,omitempty" xml:"TableCatalog,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// t_test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListSecurityIdentifyRecordsRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityIdentifyRecordsRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListSecurityIdentifyRecordsRequestListQuery) GetDatasourceEnv() *string {
	return s.DatasourceEnv
}

func (s *ListSecurityIdentifyRecordsRequestListQuery) GetDatasourceName() *string {
	return s.DatasourceName
}

func (s *ListSecurityIdentifyRecordsRequestListQuery) GetFieldName() *string {
	return s.FieldName
}

func (s *ListSecurityIdentifyRecordsRequestListQuery) GetIsDatasourceTable() *bool {
	return s.IsDatasourceTable
}

func (s *ListSecurityIdentifyRecordsRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListSecurityIdentifyRecordsRequestListQuery) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListSecurityIdentifyRecordsRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSecurityIdentifyRecordsRequestListQuery) GetTableCatalog() *string {
	return s.TableCatalog
}

func (s *ListSecurityIdentifyRecordsRequestListQuery) GetTableName() *string {
	return s.TableName
}

func (s *ListSecurityIdentifyRecordsRequestListQuery) SetDatasourceEnv(v string) *ListSecurityIdentifyRecordsRequestListQuery {
	s.DatasourceEnv = &v
	return s
}

func (s *ListSecurityIdentifyRecordsRequestListQuery) SetDatasourceName(v string) *ListSecurityIdentifyRecordsRequestListQuery {
	s.DatasourceName = &v
	return s
}

func (s *ListSecurityIdentifyRecordsRequestListQuery) SetFieldName(v string) *ListSecurityIdentifyRecordsRequestListQuery {
	s.FieldName = &v
	return s
}

func (s *ListSecurityIdentifyRecordsRequestListQuery) SetIsDatasourceTable(v bool) *ListSecurityIdentifyRecordsRequestListQuery {
	s.IsDatasourceTable = &v
	return s
}

func (s *ListSecurityIdentifyRecordsRequestListQuery) SetKeyword(v string) *ListSecurityIdentifyRecordsRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListSecurityIdentifyRecordsRequestListQuery) SetPageNo(v int32) *ListSecurityIdentifyRecordsRequestListQuery {
	s.PageNo = &v
	return s
}

func (s *ListSecurityIdentifyRecordsRequestListQuery) SetPageSize(v int32) *ListSecurityIdentifyRecordsRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListSecurityIdentifyRecordsRequestListQuery) SetTableCatalog(v string) *ListSecurityIdentifyRecordsRequestListQuery {
	s.TableCatalog = &v
	return s
}

func (s *ListSecurityIdentifyRecordsRequestListQuery) SetTableName(v string) *ListSecurityIdentifyRecordsRequestListQuery {
	s.TableName = &v
	return s
}

func (s *ListSecurityIdentifyRecordsRequestListQuery) Validate() error {
	return dara.Validate(s)
}
