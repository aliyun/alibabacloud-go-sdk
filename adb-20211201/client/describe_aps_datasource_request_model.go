// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApsDatasourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeApsDatasourceRequest
	GetDBClusterId() *string
	SetDatasourceId(v int64) *DescribeApsDatasourceRequest
	GetDatasourceId() *int64
	SetRegionId(v string) *DescribeApsDatasourceRequest
	GetRegionId() *string
}

type DescribeApsDatasourceRequest struct {
	// The ID of the AnalyticDB for MySQL cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-*******
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DatasourceId *int64 `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeApsDatasourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsDatasourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeApsDatasourceRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeApsDatasourceRequest) GetDatasourceId() *int64 {
	return s.DatasourceId
}

func (s *DescribeApsDatasourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApsDatasourceRequest) SetDBClusterId(v string) *DescribeApsDatasourceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeApsDatasourceRequest) SetDatasourceId(v int64) *DescribeApsDatasourceRequest {
	s.DatasourceId = &v
	return s
}

func (s *DescribeApsDatasourceRequest) SetRegionId(v string) *DescribeApsDatasourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApsDatasourceRequest) Validate() error {
	return dara.Validate(s)
}
