// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApsDatasourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeApsDatasourcesRequest
	GetDBClusterId() *string
	SetDatasourceName(v string) *DescribeApsDatasourcesRequest
	GetDatasourceName() *string
	SetDatasourceType(v string) *DescribeApsDatasourcesRequest
	GetDatasourceType() *string
	SetEndTime(v string) *DescribeApsDatasourcesRequest
	GetEndTime() *string
	SetPageNumber(v int32) *DescribeApsDatasourcesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApsDatasourcesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeApsDatasourcesRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeApsDatasourcesRequest
	GetStartTime() *string
}

type DescribeApsDatasourcesRequest struct {
	// The ID of the AnalyticDB for MySQL cluster.
	//
	// example:
	//
	// amv-******
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// test
	DatasourceName *string `json:"DatasourceName,omitempty" xml:"DatasourceName,omitempty"`
	// The type of the data source.
	//
	// example:
	//
	// SLS
	DatasourceType *string `json:"DatasourceType,omitempty" xml:"DatasourceType,omitempty"`
	// The end of the time range to query.
	//
	// example:
	//
	// 2024-01-30
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query.
	//
	// example:
	//
	// 2024-01-01
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeApsDatasourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsDatasourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeApsDatasourcesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeApsDatasourcesRequest) GetDatasourceName() *string {
	return s.DatasourceName
}

func (s *DescribeApsDatasourcesRequest) GetDatasourceType() *string {
	return s.DatasourceType
}

func (s *DescribeApsDatasourcesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeApsDatasourcesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApsDatasourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApsDatasourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApsDatasourcesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeApsDatasourcesRequest) SetDBClusterId(v string) *DescribeApsDatasourcesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeApsDatasourcesRequest) SetDatasourceName(v string) *DescribeApsDatasourcesRequest {
	s.DatasourceName = &v
	return s
}

func (s *DescribeApsDatasourcesRequest) SetDatasourceType(v string) *DescribeApsDatasourcesRequest {
	s.DatasourceType = &v
	return s
}

func (s *DescribeApsDatasourcesRequest) SetEndTime(v string) *DescribeApsDatasourcesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeApsDatasourcesRequest) SetPageNumber(v int32) *DescribeApsDatasourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApsDatasourcesRequest) SetPageSize(v int32) *DescribeApsDatasourcesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApsDatasourcesRequest) SetRegionId(v string) *DescribeApsDatasourcesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApsDatasourcesRequest) SetStartTime(v string) *DescribeApsDatasourcesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeApsDatasourcesRequest) Validate() error {
	return dara.Validate(s)
}
