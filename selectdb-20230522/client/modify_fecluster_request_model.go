// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFEClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyFEClusterRequest
	GetDBInstanceId() *string
	SetFEClassCode(v string) *ModifyFEClusterRequest
	GetFEClassCode() *string
	SetFEClusterId(v string) *ModifyFEClusterRequest
	GetFEClusterId() *string
	SetFENodeCount(v string) *ModifyFEClusterRequest
	GetFENodeCount() *string
	SetProduct(v string) *ModifyFEClusterRequest
	GetProduct() *string
	SetRegionId(v string) *ModifyFEClusterRequest
	GetRegionId() *string
}

type ModifyFEClusterRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The specifications of the FE cluster. Valid values:
	//
	// - selectdb_fe.xlarge16: 4 cores and 16 GB.
	//
	// - selectdb_fe.2xlarge16: 8 cores and 32 GB.
	//
	// - selectdb_fe.4xlarge16: 16 cores and 64 GB.
	//
	// - selectdb_fe.8xlarge16: 32 cores and 128 GB.
	//
	// example:
	//
	// selectdb_fe.2xlarge16
	FEClassCode *string `json:"FEClassCode,omitempty" xml:"FEClassCode,omitempty"`
	// The FE cluster ID. Default value: {instance ID}-fe.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****-fe
	FEClusterId *string `json:"FEClusterId,omitempty" xml:"FEClusterId,omitempty"`
	// The number of nodes in the FE cluster. The value must be an odd number. Valid values for locally redundant instances: 1 to 11. Valid values for zone-redundant instances: 3 to 11.
	//
	// example:
	//
	// 5
	FENodeCount *string `json:"FENodeCount,omitempty" xml:"FENodeCount,omitempty"`
	// The product ID. Default value: selectdb.
	//
	// example:
	//
	// selectdb
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyFEClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyFEClusterRequest) GoString() string {
	return s.String()
}

func (s *ModifyFEClusterRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyFEClusterRequest) GetFEClassCode() *string {
	return s.FEClassCode
}

func (s *ModifyFEClusterRequest) GetFEClusterId() *string {
	return s.FEClusterId
}

func (s *ModifyFEClusterRequest) GetFENodeCount() *string {
	return s.FENodeCount
}

func (s *ModifyFEClusterRequest) GetProduct() *string {
	return s.Product
}

func (s *ModifyFEClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyFEClusterRequest) SetDBInstanceId(v string) *ModifyFEClusterRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyFEClusterRequest) SetFEClassCode(v string) *ModifyFEClusterRequest {
	s.FEClassCode = &v
	return s
}

func (s *ModifyFEClusterRequest) SetFEClusterId(v string) *ModifyFEClusterRequest {
	s.FEClusterId = &v
	return s
}

func (s *ModifyFEClusterRequest) SetFENodeCount(v string) *ModifyFEClusterRequest {
	s.FENodeCount = &v
	return s
}

func (s *ModifyFEClusterRequest) SetProduct(v string) *ModifyFEClusterRequest {
	s.Product = &v
	return s
}

func (s *ModifyFEClusterRequest) SetRegionId(v string) *ModifyFEClusterRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyFEClusterRequest) Validate() error {
	return dara.Validate(s)
}
