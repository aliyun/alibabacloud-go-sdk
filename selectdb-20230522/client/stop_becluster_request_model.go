// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopBEClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *StopBEClusterRequest
	GetDBClusterId() *string
	SetDBInstanceId(v string) *StopBEClusterRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *StopBEClusterRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *StopBEClusterRequest
	GetResourceOwnerId() *int64
}

type StopBEClusterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s StopBEClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s StopBEClusterRequest) GoString() string {
	return s.String()
}

func (s *StopBEClusterRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *StopBEClusterRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *StopBEClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopBEClusterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *StopBEClusterRequest) SetDBClusterId(v string) *StopBEClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *StopBEClusterRequest) SetDBInstanceId(v string) *StopBEClusterRequest {
	s.DBInstanceId = &v
	return s
}

func (s *StopBEClusterRequest) SetRegionId(v string) *StopBEClusterRequest {
	s.RegionId = &v
	return s
}

func (s *StopBEClusterRequest) SetResourceOwnerId(v int64) *StopBEClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StopBEClusterRequest) Validate() error {
	return dara.Validate(s)
}
