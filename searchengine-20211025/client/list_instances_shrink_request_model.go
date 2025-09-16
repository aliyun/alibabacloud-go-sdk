// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *ListInstancesShrinkRequest
	GetCatalog() *string
	SetDataSourceType(v string) *ListInstancesShrinkRequest
	GetDataSourceType() *string
	SetDatabase(v string) *ListInstancesShrinkRequest
	GetDatabase() *string
	SetDescription(v string) *ListInstancesShrinkRequest
	GetDescription() *string
	SetEdition(v string) *ListInstancesShrinkRequest
	GetEdition() *string
	SetInstanceId(v string) *ListInstancesShrinkRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListInstancesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstancesShrinkRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *ListInstancesShrinkRequest
	GetResourceGroupId() *string
	SetTable(v string) *ListInstancesShrinkRequest
	GetTable() *string
	SetTagsShrink(v string) *ListInstancesShrinkRequest
	GetTagsShrink() *string
}

type ListInstancesShrinkRequest struct {
	Catalog        *string `json:"catalog,omitempty" xml:"catalog,omitempty"`
	DataSourceType *string `json:"dataSourceType,omitempty" xml:"dataSourceType,omitempty"`
	Database       *string `json:"database,omitempty" xml:"database,omitempty"`
	// The description of the instance. You can use this description to filter instances. Fuzzy match is supported.
	//
	// example:
	//
	// Havenask instance
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The instance type. Valid values: vector: OpenSearch Vector Search Edition instance. engine: OpenSearch Retrieval Engine Edition instance.
	//
	// example:
	//
	// vector
	Edition *string `json:"edition,omitempty" xml:"edition,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// ha-cn-83570439y0n
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 50. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-aekzgpiswzbksdi
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	Table           *string `json:"table,omitempty" xml:"table,omitempty"`
	// The tags of the instance.
	TagsShrink *string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s ListInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesShrinkRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *ListInstancesShrinkRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *ListInstancesShrinkRequest) GetDatabase() *string {
	return s.Database
}

func (s *ListInstancesShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *ListInstancesShrinkRequest) GetEdition() *string {
	return s.Edition
}

func (s *ListInstancesShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstancesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstancesShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstancesShrinkRequest) GetTable() *string {
	return s.Table
}

func (s *ListInstancesShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ListInstancesShrinkRequest) SetCatalog(v string) *ListInstancesShrinkRequest {
	s.Catalog = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetDataSourceType(v string) *ListInstancesShrinkRequest {
	s.DataSourceType = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetDatabase(v string) *ListInstancesShrinkRequest {
	s.Database = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetDescription(v string) *ListInstancesShrinkRequest {
	s.Description = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetEdition(v string) *ListInstancesShrinkRequest {
	s.Edition = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetInstanceId(v string) *ListInstancesShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetPageNumber(v int32) *ListInstancesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetPageSize(v int32) *ListInstancesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetResourceGroupId(v string) *ListInstancesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetTable(v string) *ListInstancesShrinkRequest {
	s.Table = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetTagsShrink(v string) *ListInstancesShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
