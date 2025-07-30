// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBInstanceAttributeRequest
	GetDBInstanceId() *string
	SetInstanceAttributeType(v string) *ModifyDBInstanceAttributeRequest
	GetInstanceAttributeType() *string
	SetRegionId(v string) *ModifyDBInstanceAttributeRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceAttributeRequest
	GetResourceOwnerId() *int64
	SetValue(v string) *ModifyDBInstanceAttributeRequest
	GetValue() *string
}

type ModifyDBInstanceAttributeRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The instance parameter to be modified. Valid values:
	//
	// 	- **MaintainTime**: Modify the maintenance window of the instance in the hh:mm-hh:mm format.
	//
	// 	- **DBInstanceDescription**: Modify the description of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// DBInstanceDescription
	InstanceAttributeType *string `json:"InstanceAttributeType,omitempty" xml:"InstanceAttributeType,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The new value of the instance parameter to be modified. Examples:
	//
	// 	- If InstanceAttributeType is set to MaintainTime, you can set Value to 00:00-06:00.
	//
	// 	- If InstanceAttributeType is set to DBInstanceDescription, you can set Value to testdb.
	//
	// This parameter is required.
	//
	// example:
	//
	// testdb01
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyDBInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceAttributeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceAttributeRequest) GetInstanceAttributeType() *string {
	return s.InstanceAttributeType
}

func (s *ModifyDBInstanceAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBInstanceAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceAttributeRequest) GetValue() *string {
	return s.Value
}

func (s *ModifyDBInstanceAttributeRequest) SetDBInstanceId(v string) *ModifyDBInstanceAttributeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceAttributeRequest) SetInstanceAttributeType(v string) *ModifyDBInstanceAttributeRequest {
	s.InstanceAttributeType = &v
	return s
}

func (s *ModifyDBInstanceAttributeRequest) SetRegionId(v string) *ModifyDBInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBInstanceAttributeRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceAttributeRequest) SetValue(v string) *ModifyDBInstanceAttributeRequest {
	s.Value = &v
	return s
}

func (s *ModifyDBInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
