// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *ListInstancesRequest
	GetCatalog() *string
	SetDataSourceType(v string) *ListInstancesRequest
	GetDataSourceType() *string
	SetDatabase(v string) *ListInstancesRequest
	GetDatabase() *string
	SetDescription(v string) *ListInstancesRequest
	GetDescription() *string
	SetEdition(v string) *ListInstancesRequest
	GetEdition() *string
	SetInstanceId(v string) *ListInstancesRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstancesRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *ListInstancesRequest
	GetResourceGroupId() *string
	SetTable(v string) *ListInstancesRequest
	GetTable() *string
	SetTags(v []*ListInstancesRequestTags) *ListInstancesRequest
	GetTags() []*ListInstancesRequestTags
}

type ListInstancesRequest struct {
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
	Tags []*ListInstancesRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s ListInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *ListInstancesRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *ListInstancesRequest) GetDatabase() *string {
	return s.Database
}

func (s *ListInstancesRequest) GetDescription() *string {
	return s.Description
}

func (s *ListInstancesRequest) GetEdition() *string {
	return s.Edition
}

func (s *ListInstancesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstancesRequest) GetTable() *string {
	return s.Table
}

func (s *ListInstancesRequest) GetTags() []*ListInstancesRequestTags {
	return s.Tags
}

func (s *ListInstancesRequest) SetCatalog(v string) *ListInstancesRequest {
	s.Catalog = &v
	return s
}

func (s *ListInstancesRequest) SetDataSourceType(v string) *ListInstancesRequest {
	s.DataSourceType = &v
	return s
}

func (s *ListInstancesRequest) SetDatabase(v string) *ListInstancesRequest {
	s.Database = &v
	return s
}

func (s *ListInstancesRequest) SetDescription(v string) *ListInstancesRequest {
	s.Description = &v
	return s
}

func (s *ListInstancesRequest) SetEdition(v string) *ListInstancesRequest {
	s.Edition = &v
	return s
}

func (s *ListInstancesRequest) SetInstanceId(v string) *ListInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesRequest) SetPageNumber(v int32) *ListInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesRequest) SetPageSize(v int32) *ListInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesRequest) SetResourceGroupId(v string) *ListInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesRequest) SetTable(v string) *ListInstancesRequest {
	s.Table = &v
	return s
}

func (s *ListInstancesRequest) SetTags(v []*ListInstancesRequestTags) *ListInstancesRequest {
	s.Tags = v
	return s
}

func (s *ListInstancesRequest) Validate() error {
	return dara.Validate(s)
}

type ListInstancesRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// backup
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// oboms-disk
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListInstancesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesRequestTags) GoString() string {
	return s.String()
}

func (s *ListInstancesRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListInstancesRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListInstancesRequestTags) SetKey(v string) *ListInstancesRequestTags {
	s.Key = &v
	return s
}

func (s *ListInstancesRequestTags) SetValue(v string) *ListInstancesRequestTags {
	s.Value = &v
	return s
}

func (s *ListInstancesRequestTags) Validate() error {
	return dara.Validate(s)
}
