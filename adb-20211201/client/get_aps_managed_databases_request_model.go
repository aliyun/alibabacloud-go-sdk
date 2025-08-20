// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApsManagedDatabasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *GetApsManagedDatabasesRequest
	GetDBClusterId() *string
	SetRegionId(v string) *GetApsManagedDatabasesRequest
	GetRegionId() *string
}

type GetApsManagedDatabasesRequest struct {
	// The ID of the AnalyticDB for MySQL cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-******
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetApsManagedDatabasesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApsManagedDatabasesRequest) GoString() string {
	return s.String()
}

func (s *GetApsManagedDatabasesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetApsManagedDatabasesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetApsManagedDatabasesRequest) SetDBClusterId(v string) *GetApsManagedDatabasesRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetApsManagedDatabasesRequest) SetRegionId(v string) *GetApsManagedDatabasesRequest {
	s.RegionId = &v
	return s
}

func (s *GetApsManagedDatabasesRequest) Validate() error {
	return dara.Validate(s)
}
