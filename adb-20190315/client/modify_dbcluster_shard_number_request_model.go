// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterShardNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyDBClusterShardNumberRequest
	GetDBClusterId() *string
	SetDryRun(v bool) *ModifyDBClusterShardNumberRequest
	GetDryRun() *bool
	SetNewShardNumber(v int64) *ModifyDBClusterShardNumberRequest
	GetNewShardNumber() *int64
	SetOwnerAccount(v string) *ModifyDBClusterShardNumberRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBClusterShardNumberRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyDBClusterShardNumberRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyDBClusterShardNumberRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBClusterShardNumberRequest
	GetResourceOwnerId() *int64
	SetSwitchTime(v string) *ModifyDBClusterShardNumberRequest
	GetSwitchTime() *string
	SetSwitchTimeMode(v int64) *ModifyDBClusterShardNumberRequest
	GetSwitchTimeMode() *int64
}

type ModifyDBClusterShardNumberRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the information about all AnalyticDB for MySQL Data Warehouse Edition clusters within a region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1xxxxxxxx47
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to perform only a dry run. Valid values:
	//
	// 	- true: sends a request to check whether the cluster meets the prerequisites for changing the number of shards and whether the desired number of shards is valid, but **does not*	- perform the change operation.
	//
	// 	- false (default): sends a request to perform a check and trigger the change operation.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The desired number of shards.
	//
	// example:
	//
	// 256
	NewShardNumber *int64  `json:"NewShardNumber,omitempty" xml:"NewShardNumber,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The point in time when you want the system to perform the change operation. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// example:
	//
	// 2021-07-09T22:00:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The mode in which you want the change operation to be performed. Valid values:
	//
	// 	- **Immediate*	- (default): immediately performs the change operation.
	//
	// 	- **MaintainTime**: performs the change operation within the maintenance window of the cluster. You can call the ModifyDBInstanceMaintainTime operation to change the maintenance window.
	//
	// 	- **ScheduleTime**: performs the change operation at the point in time that you specify. If you specify this value, you must also specify **SwitchTime**.
	//
	// example:
	//
	// Immediate
	SwitchTimeMode *int64 `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
}

func (s ModifyDBClusterShardNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterShardNumberRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterShardNumberRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterShardNumberRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyDBClusterShardNumberRequest) GetNewShardNumber() *int64 {
	return s.NewShardNumber
}

func (s *ModifyDBClusterShardNumberRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBClusterShardNumberRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBClusterShardNumberRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBClusterShardNumberRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBClusterShardNumberRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBClusterShardNumberRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *ModifyDBClusterShardNumberRequest) GetSwitchTimeMode() *int64 {
	return s.SwitchTimeMode
}

func (s *ModifyDBClusterShardNumberRequest) SetDBClusterId(v string) *ModifyDBClusterShardNumberRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterShardNumberRequest) SetDryRun(v bool) *ModifyDBClusterShardNumberRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyDBClusterShardNumberRequest) SetNewShardNumber(v int64) *ModifyDBClusterShardNumberRequest {
	s.NewShardNumber = &v
	return s
}

func (s *ModifyDBClusterShardNumberRequest) SetOwnerAccount(v string) *ModifyDBClusterShardNumberRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterShardNumberRequest) SetOwnerId(v int64) *ModifyDBClusterShardNumberRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterShardNumberRequest) SetRegionId(v string) *ModifyDBClusterShardNumberRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBClusterShardNumberRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterShardNumberRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterShardNumberRequest) SetResourceOwnerId(v int64) *ModifyDBClusterShardNumberRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterShardNumberRequest) SetSwitchTime(v string) *ModifyDBClusterShardNumberRequest {
	s.SwitchTime = &v
	return s
}

func (s *ModifyDBClusterShardNumberRequest) SetSwitchTimeMode(v int64) *ModifyDBClusterShardNumberRequest {
	s.SwitchTimeMode = &v
	return s
}

func (s *ModifyDBClusterShardNumberRequest) Validate() error {
	return dara.Validate(s)
}
