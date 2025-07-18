// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExternalDataServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ListExternalDataServicesRequest
	GetDBInstanceId() *string
	SetPageNumber(v int32) *ListExternalDataServicesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListExternalDataServicesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListExternalDataServicesRequest
	GetRegionId() *string
}

type ListExternalDataServicesRequest struct {
	// The instance ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp10g78o9807yv9h3
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListExternalDataServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExternalDataServicesRequest) GoString() string {
	return s.String()
}

func (s *ListExternalDataServicesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListExternalDataServicesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListExternalDataServicesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListExternalDataServicesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListExternalDataServicesRequest) SetDBInstanceId(v string) *ListExternalDataServicesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ListExternalDataServicesRequest) SetPageNumber(v int32) *ListExternalDataServicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListExternalDataServicesRequest) SetPageSize(v int32) *ListExternalDataServicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListExternalDataServicesRequest) SetRegionId(v string) *ListExternalDataServicesRequest {
	s.RegionId = &v
	return s
}

func (s *ListExternalDataServicesRequest) Validate() error {
	return dara.Validate(s)
}
