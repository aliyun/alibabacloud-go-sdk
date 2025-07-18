// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStreamingDataSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ListStreamingDataSourcesRequest
	GetDBInstanceId() *string
	SetPageNumber(v int32) *ListStreamingDataSourcesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListStreamingDataSourcesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListStreamingDataSourcesRequest
	GetRegionId() *string
}

type ListStreamingDataSourcesRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp10g78o9807yv9h3
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The page number. Pages start from page 1. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListStreamingDataSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStreamingDataSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListStreamingDataSourcesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListStreamingDataSourcesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListStreamingDataSourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListStreamingDataSourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListStreamingDataSourcesRequest) SetDBInstanceId(v string) *ListStreamingDataSourcesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ListStreamingDataSourcesRequest) SetPageNumber(v int32) *ListStreamingDataSourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStreamingDataSourcesRequest) SetPageSize(v int32) *ListStreamingDataSourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListStreamingDataSourcesRequest) SetRegionId(v string) *ListStreamingDataSourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListStreamingDataSourcesRequest) Validate() error {
	return dara.Validate(s)
}
