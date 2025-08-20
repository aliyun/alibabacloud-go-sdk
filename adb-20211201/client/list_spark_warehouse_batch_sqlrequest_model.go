// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSparkWarehouseBatchSQLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ListSparkWarehouseBatchSQLRequest
	GetDBClusterId() *string
	SetPageNumber(v int64) *ListSparkWarehouseBatchSQLRequest
	GetPageNumber() *int64
	SetPageSize(v string) *ListSparkWarehouseBatchSQLRequest
	GetPageSize() *string
	SetResourceGroupName(v string) *ListSparkWarehouseBatchSQLRequest
	GetResourceGroupName() *string
}

type ListSparkWarehouseBatchSQLRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1ub9grke1****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the interactive resource group for which the Spark engine is enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s ListSparkWarehouseBatchSQLRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSparkWarehouseBatchSQLRequest) GoString() string {
	return s.String()
}

func (s *ListSparkWarehouseBatchSQLRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ListSparkWarehouseBatchSQLRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListSparkWarehouseBatchSQLRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListSparkWarehouseBatchSQLRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *ListSparkWarehouseBatchSQLRequest) SetDBClusterId(v string) *ListSparkWarehouseBatchSQLRequest {
	s.DBClusterId = &v
	return s
}

func (s *ListSparkWarehouseBatchSQLRequest) SetPageNumber(v int64) *ListSparkWarehouseBatchSQLRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSparkWarehouseBatchSQLRequest) SetPageSize(v string) *ListSparkWarehouseBatchSQLRequest {
	s.PageSize = &v
	return s
}

func (s *ListSparkWarehouseBatchSQLRequest) SetResourceGroupName(v string) *ListSparkWarehouseBatchSQLRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *ListSparkWarehouseBatchSQLRequest) Validate() error {
	return dara.Validate(s)
}
