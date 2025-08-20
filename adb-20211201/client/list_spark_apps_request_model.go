// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSparkAppsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ListSparkAppsRequest
	GetDBClusterId() *string
	SetFilters(v string) *ListSparkAppsRequest
	GetFilters() *string
	SetPageNumber(v int64) *ListSparkAppsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListSparkAppsRequest
	GetPageSize() *int64
	SetResourceGroupName(v string) *ListSparkAppsRequest
	GetResourceGroupName() *string
}

type ListSparkAppsRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Filters     *string `json:"Filters,omitempty" xml:"Filters,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: **1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10. Valid values:
	//
	// - **10**
	//
	// - **50**
	//
	// - **100**
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the job resource group.
	//
	// example:
	//
	// test_instance
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s ListSparkAppsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSparkAppsRequest) GoString() string {
	return s.String()
}

func (s *ListSparkAppsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ListSparkAppsRequest) GetFilters() *string {
	return s.Filters
}

func (s *ListSparkAppsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListSparkAppsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSparkAppsRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *ListSparkAppsRequest) SetDBClusterId(v string) *ListSparkAppsRequest {
	s.DBClusterId = &v
	return s
}

func (s *ListSparkAppsRequest) SetFilters(v string) *ListSparkAppsRequest {
	s.Filters = &v
	return s
}

func (s *ListSparkAppsRequest) SetPageNumber(v int64) *ListSparkAppsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSparkAppsRequest) SetPageSize(v int64) *ListSparkAppsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSparkAppsRequest) SetResourceGroupName(v string) *ListSparkAppsRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *ListSparkAppsRequest) Validate() error {
	return dara.Validate(s)
}
