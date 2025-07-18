// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStreamingDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeleteStreamingDataSourceRequest
	GetDBInstanceId() *string
	SetDataSourceId(v int32) *DeleteStreamingDataSourceRequest
	GetDataSourceId() *int32
	SetRegionId(v string) *DeleteStreamingDataSourceRequest
	GetRegionId() *string
}

type DeleteStreamingDataSourceRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp10g78o9807yv9h3
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DataSourceId *int32 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteStreamingDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStreamingDataSourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteStreamingDataSourceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteStreamingDataSourceRequest) GetDataSourceId() *int32 {
	return s.DataSourceId
}

func (s *DeleteStreamingDataSourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteStreamingDataSourceRequest) SetDBInstanceId(v string) *DeleteStreamingDataSourceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteStreamingDataSourceRequest) SetDataSourceId(v int32) *DeleteStreamingDataSourceRequest {
	s.DataSourceId = &v
	return s
}

func (s *DeleteStreamingDataSourceRequest) SetRegionId(v string) *DeleteStreamingDataSourceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteStreamingDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
