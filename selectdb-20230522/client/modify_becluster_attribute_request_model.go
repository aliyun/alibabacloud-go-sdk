// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBEClusterAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyBEClusterAttributeRequest
	GetDBClusterId() *string
	SetDBInstanceId(v string) *ModifyBEClusterAttributeRequest
	GetDBInstanceId() *string
	SetInstanceAttributeType(v string) *ModifyBEClusterAttributeRequest
	GetInstanceAttributeType() *string
	SetRegionId(v string) *ModifyBEClusterAttributeRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *ModifyBEClusterAttributeRequest
	GetResourceOwnerId() *int64
	SetValue(v string) *ModifyBEClusterAttributeRequest
	GetValue() *string
}

type ModifyBEClusterAttributeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// selectdb-xxxb9f2w-be
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The attribute type of the instance. Set this parameter to DBInstanceDescription.
	//
	// Valid values:
	//
	// 	- MaintainTime
	//
	// 	- DBInstanceDescription
	//
	// This parameter is required.
	//
	// example:
	//
	// DBInstanceDescription
	InstanceAttributeType *string `json:"InstanceAttributeType,omitempty" xml:"InstanceAttributeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The new name of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// testdb
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyBEClusterAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBEClusterAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyBEClusterAttributeRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyBEClusterAttributeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyBEClusterAttributeRequest) GetInstanceAttributeType() *string {
	return s.InstanceAttributeType
}

func (s *ModifyBEClusterAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyBEClusterAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyBEClusterAttributeRequest) GetValue() *string {
	return s.Value
}

func (s *ModifyBEClusterAttributeRequest) SetDBClusterId(v string) *ModifyBEClusterAttributeRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyBEClusterAttributeRequest) SetDBInstanceId(v string) *ModifyBEClusterAttributeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyBEClusterAttributeRequest) SetInstanceAttributeType(v string) *ModifyBEClusterAttributeRequest {
	s.InstanceAttributeType = &v
	return s
}

func (s *ModifyBEClusterAttributeRequest) SetRegionId(v string) *ModifyBEClusterAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyBEClusterAttributeRequest) SetResourceOwnerId(v int64) *ModifyBEClusterAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyBEClusterAttributeRequest) SetValue(v string) *ModifyBEClusterAttributeRequest {
	s.Value = &v
	return s
}

func (s *ModifyBEClusterAttributeRequest) Validate() error {
	return dara.Validate(s)
}
