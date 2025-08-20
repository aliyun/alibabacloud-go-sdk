// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApsDatasoureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteApsDatasoureRequest
	GetDBClusterId() *string
	SetDatasourceId(v int64) *DeleteApsDatasoureRequest
	GetDatasourceId() *int64
	SetRegionId(v string) *DeleteApsDatasoureRequest
	GetRegionId() *string
}

type DeleteApsDatasoureRequest struct {
	// The ID of the AnalyticDB for MySQL cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-******
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
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteApsDatasoureRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApsDatasoureRequest) GoString() string {
	return s.String()
}

func (s *DeleteApsDatasoureRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteApsDatasoureRequest) GetDatasourceId() *int64 {
	return s.DatasourceId
}

func (s *DeleteApsDatasoureRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteApsDatasoureRequest) SetDBClusterId(v string) *DeleteApsDatasoureRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteApsDatasoureRequest) SetDatasourceId(v int64) *DeleteApsDatasoureRequest {
	s.DatasourceId = &v
	return s
}

func (s *DeleteApsDatasoureRequest) SetRegionId(v string) *DeleteApsDatasoureRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteApsDatasoureRequest) Validate() error {
	return dara.Validate(s)
}
