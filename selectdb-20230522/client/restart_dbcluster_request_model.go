// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDBClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *RestartDBClusterRequest
	GetDBClusterId() *string
	SetDBInstanceId(v string) *RestartDBClusterRequest
	GetDBInstanceId() *string
	SetParallelOperation(v bool) *RestartDBClusterRequest
	GetParallelOperation() *bool
	SetRegionId(v string) *RestartDBClusterRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *RestartDBClusterRequest
	GetResourceGroupId() *string
	SetResourceOwnerId(v int64) *RestartDBClusterRequest
	GetResourceOwnerId() *int64
}

type RestartDBClusterRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213c8yvv09-be
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Specifies whether to perform parallel operations on the cluster node.
	//
	// example:
	//
	// false
	ParallelOperation *bool `json:"ParallelOperation,omitempty" xml:"ParallelOperation,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-4690g37929****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RestartDBClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartDBClusterRequest) GoString() string {
	return s.String()
}

func (s *RestartDBClusterRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *RestartDBClusterRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *RestartDBClusterRequest) GetParallelOperation() *bool {
	return s.ParallelOperation
}

func (s *RestartDBClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RestartDBClusterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *RestartDBClusterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RestartDBClusterRequest) SetDBClusterId(v string) *RestartDBClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *RestartDBClusterRequest) SetDBInstanceId(v string) *RestartDBClusterRequest {
	s.DBInstanceId = &v
	return s
}

func (s *RestartDBClusterRequest) SetParallelOperation(v bool) *RestartDBClusterRequest {
	s.ParallelOperation = &v
	return s
}

func (s *RestartDBClusterRequest) SetRegionId(v string) *RestartDBClusterRequest {
	s.RegionId = &v
	return s
}

func (s *RestartDBClusterRequest) SetResourceGroupId(v string) *RestartDBClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *RestartDBClusterRequest) SetResourceOwnerId(v int64) *RestartDBClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RestartDBClusterRequest) Validate() error {
	return dara.Validate(s)
}
