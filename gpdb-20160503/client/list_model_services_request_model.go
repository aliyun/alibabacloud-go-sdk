// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ListModelServicesRequest
	GetDBInstanceId() *string
	SetPageNumber(v int32) *ListModelServicesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModelServicesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListModelServicesRequest
	GetRegionId() *string
}

type ListModelServicesRequest struct {
	// The ID of the instance.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **20**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// Default value: **20**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the cluster.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListModelServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelServicesRequest) GoString() string {
	return s.String()
}

func (s *ListModelServicesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListModelServicesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelServicesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelServicesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListModelServicesRequest) SetDBInstanceId(v string) *ListModelServicesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ListModelServicesRequest) SetPageNumber(v int32) *ListModelServicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelServicesRequest) SetPageSize(v int32) *ListModelServicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelServicesRequest) SetRegionId(v string) *ListModelServicesRequest {
	s.RegionId = &v
	return s
}

func (s *ListModelServicesRequest) Validate() error {
	return dara.Validate(s)
}
