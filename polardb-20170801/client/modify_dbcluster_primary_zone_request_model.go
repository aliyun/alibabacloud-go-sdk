// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterPrimaryZoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyDBClusterPrimaryZoneRequest
	GetDBClusterId() *string
	SetFromTimeService(v bool) *ModifyDBClusterPrimaryZoneRequest
	GetFromTimeService() *bool
	SetIsSwitchOverForDisaster(v string) *ModifyDBClusterPrimaryZoneRequest
	GetIsSwitchOverForDisaster() *string
	SetOwnerAccount(v string) *ModifyDBClusterPrimaryZoneRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBClusterPrimaryZoneRequest
	GetOwnerId() *int64
	SetPlannedEndTime(v string) *ModifyDBClusterPrimaryZoneRequest
	GetPlannedEndTime() *string
	SetPlannedStartTime(v string) *ModifyDBClusterPrimaryZoneRequest
	GetPlannedStartTime() *string
	SetResourceOwnerAccount(v string) *ModifyDBClusterPrimaryZoneRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBClusterPrimaryZoneRequest
	GetResourceOwnerId() *int64
	SetVPCId(v string) *ModifyDBClusterPrimaryZoneRequest
	GetVPCId() *string
	SetVSwitchId(v string) *ModifyDBClusterPrimaryZoneRequest
	GetVSwitchId() *string
	SetZoneId(v string) *ModifyDBClusterPrimaryZoneRequest
	GetZoneId() *string
	SetZoneType(v string) *ModifyDBClusterPrimaryZoneRequest
	GetZoneType() *string
}

type ModifyDBClusterPrimaryZoneRequest struct {
	// The cluster ID.
	//
	// > Call the [DescribeDBClusters](https://help.aliyun.com/document_detail/173433.html) operation to query the details of all clusters in a destination region, including their cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to perform the zone change immediately or at a scheduled time. Valid values:
	//
	// - false (default): The zone change is performed at the scheduled time.
	//
	// - true: The zone change is performed immediately.
	//
	// example:
	//
	// false
	FromTimeService *bool `json:"FromTimeService,omitempty" xml:"FromTimeService,omitempty"`
	// Specifies whether to fail back to the original zone. Valid values:
	//
	// - true: Fails back to the original zone.
	//
	// - false: Does not fail back to the original zone.
	//
	// example:
	//
	// false
	IsSwitchOverForDisaster *string `json:"IsSwitchOverForDisaster,omitempty" xml:"IsSwitchOverForDisaster,omitempty"`
	OwnerAccount            *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                 *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The latest time to start the scheduled task. Specify the time in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC.
	//
	// > - The latest start time must be at least 30 minutes later than the earliest start time.
	//
	// >
	//
	// > - If you specify `PlannedStartTime` but not this parameter, the latest start time of the task is the value of `PlannedStartTime` plus 30 minutes by default. For example, if you set `PlannedStartTime` to `2021-01-14T09:00:00Z` and leave this parameter empty, the task starts no later than `2021-01-14T09:30:00Z`.
	//
	// example:
	//
	// 2021-01-14T09:30:00Z
	PlannedEndTime *string `json:"PlannedEndTime,omitempty" xml:"PlannedEndTime,omitempty"`
	// The earliest time to start the scheduled task to change the zone. Specify the time in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC.
	//
	// > - The start time must be a point in time within the next 24 hours. For example, if the current time is `2021-01-14T09:00:00Z`, you can set the start time to a value from `2021-01-14T09:00:00Z` to `2021-01-15T09:00:00Z`.
	//
	// >
	//
	// > - If you do not specify this parameter, the zone change task is performed immediately.
	//
	// example:
	//
	// 2021-01-14T09:00:00Z
	PlannedStartTime     *string `json:"PlannedStartTime,omitempty" xml:"PlannedStartTime,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-**********
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The ID of the vSwitch in the destination zone.
	//
	// > - This parameter is required for PolarDB for Oracle and PolarDB for PostgreSQL clusters.
	//
	// >
	//
	// > - For PolarDB for MySQL clusters, this parameter is required if a vSwitch is created in the destination zone. If no vSwitch is created, the default vSwitch is used and this parameter is optional.
	//
	// example:
	//
	// vsw-**************
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the new zone.
	//
	// > Call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query available zones.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The type of the zone. Valid values:
	//
	// - **Primary**: The primary zone.
	//
	// - **Standby**: The secondary zone.
	//
	// example:
	//
	// Primary
	ZoneType *string `json:"ZoneType,omitempty" xml:"ZoneType,omitempty"`
}

func (s ModifyDBClusterPrimaryZoneRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterPrimaryZoneRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterPrimaryZoneRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterPrimaryZoneRequest) GetFromTimeService() *bool {
	return s.FromTimeService
}

func (s *ModifyDBClusterPrimaryZoneRequest) GetIsSwitchOverForDisaster() *string {
	return s.IsSwitchOverForDisaster
}

func (s *ModifyDBClusterPrimaryZoneRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBClusterPrimaryZoneRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBClusterPrimaryZoneRequest) GetPlannedEndTime() *string {
	return s.PlannedEndTime
}

func (s *ModifyDBClusterPrimaryZoneRequest) GetPlannedStartTime() *string {
	return s.PlannedStartTime
}

func (s *ModifyDBClusterPrimaryZoneRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBClusterPrimaryZoneRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBClusterPrimaryZoneRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *ModifyDBClusterPrimaryZoneRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ModifyDBClusterPrimaryZoneRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ModifyDBClusterPrimaryZoneRequest) GetZoneType() *string {
	return s.ZoneType
}

func (s *ModifyDBClusterPrimaryZoneRequest) SetDBClusterId(v string) *ModifyDBClusterPrimaryZoneRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterPrimaryZoneRequest) SetFromTimeService(v bool) *ModifyDBClusterPrimaryZoneRequest {
	s.FromTimeService = &v
	return s
}

func (s *ModifyDBClusterPrimaryZoneRequest) SetIsSwitchOverForDisaster(v string) *ModifyDBClusterPrimaryZoneRequest {
	s.IsSwitchOverForDisaster = &v
	return s
}

func (s *ModifyDBClusterPrimaryZoneRequest) SetOwnerAccount(v string) *ModifyDBClusterPrimaryZoneRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterPrimaryZoneRequest) SetOwnerId(v int64) *ModifyDBClusterPrimaryZoneRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterPrimaryZoneRequest) SetPlannedEndTime(v string) *ModifyDBClusterPrimaryZoneRequest {
	s.PlannedEndTime = &v
	return s
}

func (s *ModifyDBClusterPrimaryZoneRequest) SetPlannedStartTime(v string) *ModifyDBClusterPrimaryZoneRequest {
	s.PlannedStartTime = &v
	return s
}

func (s *ModifyDBClusterPrimaryZoneRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterPrimaryZoneRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterPrimaryZoneRequest) SetResourceOwnerId(v int64) *ModifyDBClusterPrimaryZoneRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterPrimaryZoneRequest) SetVPCId(v string) *ModifyDBClusterPrimaryZoneRequest {
	s.VPCId = &v
	return s
}

func (s *ModifyDBClusterPrimaryZoneRequest) SetVSwitchId(v string) *ModifyDBClusterPrimaryZoneRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyDBClusterPrimaryZoneRequest) SetZoneId(v string) *ModifyDBClusterPrimaryZoneRequest {
	s.ZoneId = &v
	return s
}

func (s *ModifyDBClusterPrimaryZoneRequest) SetZoneType(v string) *ModifyDBClusterPrimaryZoneRequest {
	s.ZoneType = &v
	return s
}

func (s *ModifyDBClusterPrimaryZoneRequest) Validate() error {
	return dara.Validate(s)
}
