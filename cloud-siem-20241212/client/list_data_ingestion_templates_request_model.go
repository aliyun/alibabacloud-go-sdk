// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataIngestionTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataIngestionTemplateStatus(v string) *ListDataIngestionTemplatesRequest
	GetDataIngestionTemplateStatus() *string
	SetDataSourceTemplateIds(v string) *ListDataIngestionTemplatesRequest
	GetDataSourceTemplateIds() *string
	SetLang(v string) *ListDataIngestionTemplatesRequest
	GetLang() *string
	SetPageNumber(v string) *ListDataIngestionTemplatesRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListDataIngestionTemplatesRequest
	GetPageSize() *string
	SetProductId(v string) *ListDataIngestionTemplatesRequest
	GetProductId() *string
	SetRegionId(v string) *ListDataIngestionTemplatesRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListDataIngestionTemplatesRequest
	GetRoleFor() *int64
}

type ListDataIngestionTemplatesRequest struct {
	// example:
	//
	// running。
	DataIngestionTemplateStatus *string `json:"DataIngestionTemplateStatus,omitempty" xml:"DataIngestionTemplateStatus,omitempty"`
	// example:
	//
	// alibaba_cloud_sas_account_snapshot_log_173326*******。
	DataSourceTemplateIds *string `json:"DataSourceTemplateIds,omitempty" xml:"DataSourceTemplateIds,omitempty"`
	// example:
	//
	// zh。
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 1。
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10。
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// alibaba_cloud_sas。
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// cn-hangzhou。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 173326*******。
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s ListDataIngestionTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataIngestionTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListDataIngestionTemplatesRequest) GetDataIngestionTemplateStatus() *string {
	return s.DataIngestionTemplateStatus
}

func (s *ListDataIngestionTemplatesRequest) GetDataSourceTemplateIds() *string {
	return s.DataSourceTemplateIds
}

func (s *ListDataIngestionTemplatesRequest) GetLang() *string {
	return s.Lang
}

func (s *ListDataIngestionTemplatesRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListDataIngestionTemplatesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListDataIngestionTemplatesRequest) GetProductId() *string {
	return s.ProductId
}

func (s *ListDataIngestionTemplatesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDataIngestionTemplatesRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListDataIngestionTemplatesRequest) SetDataIngestionTemplateStatus(v string) *ListDataIngestionTemplatesRequest {
	s.DataIngestionTemplateStatus = &v
	return s
}

func (s *ListDataIngestionTemplatesRequest) SetDataSourceTemplateIds(v string) *ListDataIngestionTemplatesRequest {
	s.DataSourceTemplateIds = &v
	return s
}

func (s *ListDataIngestionTemplatesRequest) SetLang(v string) *ListDataIngestionTemplatesRequest {
	s.Lang = &v
	return s
}

func (s *ListDataIngestionTemplatesRequest) SetPageNumber(v string) *ListDataIngestionTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataIngestionTemplatesRequest) SetPageSize(v string) *ListDataIngestionTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataIngestionTemplatesRequest) SetProductId(v string) *ListDataIngestionTemplatesRequest {
	s.ProductId = &v
	return s
}

func (s *ListDataIngestionTemplatesRequest) SetRegionId(v string) *ListDataIngestionTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *ListDataIngestionTemplatesRequest) SetRoleFor(v int64) *ListDataIngestionTemplatesRequest {
	s.RoleFor = &v
	return s
}

func (s *ListDataIngestionTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
