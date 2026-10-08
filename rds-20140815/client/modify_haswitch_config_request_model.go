// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHASwitchConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyHASwitchConfigRequest
	GetDBInstanceId() *string
	SetHAConfig(v string) *ModifyHASwitchConfigRequest
	GetHAConfig() *string
	SetManualHATime(v string) *ModifyHASwitchConfigRequest
	GetManualHATime() *string
	SetOwnerId(v int64) *ModifyHASwitchConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyHASwitchConfigRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyHASwitchConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyHASwitchConfigRequest
	GetResourceOwnerId() *int64
}

type ModifyHASwitchConfigRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The primary/secondary switchover setting. Valid values:
	//
	// 	- **Auto**: The system automatically switches over between the primary and secondary instances upon a fault.
	//
	// 	- **Manual**: Temporarily disables automatic switchover.
	//
	// Default value: **Auto**.
	//
	// >If you set this parameter to **Manual**, you must also specify the **ManualHATime*	- parameter.
	//
	// example:
	//
	// Manual
	HAConfig *string `json:"HAConfig,omitempty" xml:"HAConfig,omitempty"`
	// The deadline for temporarily disabling automatic switchover. You can set this parameter to a point in time up to 23:59:59 seven days later. Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// >This parameter takes effect only when **HAConfig*	- is set to **Manual**.
	//
	// example:
	//
	// 2019-08-29T15:00:00Z
	ManualHATime *string `json:"ManualHATime,omitempty" xml:"ManualHATime,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyHASwitchConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHASwitchConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyHASwitchConfigRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyHASwitchConfigRequest) GetHAConfig() *string {
	return s.HAConfig
}

func (s *ModifyHASwitchConfigRequest) GetManualHATime() *string {
	return s.ManualHATime
}

func (s *ModifyHASwitchConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyHASwitchConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyHASwitchConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyHASwitchConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyHASwitchConfigRequest) SetDBInstanceId(v string) *ModifyHASwitchConfigRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyHASwitchConfigRequest) SetHAConfig(v string) *ModifyHASwitchConfigRequest {
	s.HAConfig = &v
	return s
}

func (s *ModifyHASwitchConfigRequest) SetManualHATime(v string) *ModifyHASwitchConfigRequest {
	s.ManualHATime = &v
	return s
}

func (s *ModifyHASwitchConfigRequest) SetOwnerId(v int64) *ModifyHASwitchConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyHASwitchConfigRequest) SetRegionId(v string) *ModifyHASwitchConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHASwitchConfigRequest) SetResourceOwnerAccount(v string) *ModifyHASwitchConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyHASwitchConfigRequest) SetResourceOwnerId(v int64) *ModifyHASwitchConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyHASwitchConfigRequest) Validate() error {
	return dara.Validate(s)
}
