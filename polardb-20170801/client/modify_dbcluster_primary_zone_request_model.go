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
	// The ID of the cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/173433.html) operation to query information about all clusters that are deployed in a specified region, such as the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to change the primary zone immediately. Valid values:
	//
	// 	- false (default): changes the primary zone as scheduled.
	//
	// 	- true: changes the primary zone immediately.
	//
	// example:
	//
	// false
	FromTimeService *bool `json:"FromTimeService,omitempty" xml:"FromTimeService,omitempty"`
	// Specifies whether to switch back to the original primary zone.
	//
	// 	- true: switches back to the original primary zone.
	//
	// 	- false: does not switch back to the original primary zone.
	//
	// example:
	//
	// false
	IsSwitchOverForDisaster *string `json:"IsSwitchOverForDisaster,omitempty" xml:"IsSwitchOverForDisaster,omitempty"`
	OwnerAccount            *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                 *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The latest start time to switch the primary zone within the scheduled time period. Specify the time in the ISO 8601 standard in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC.
	//
	// > 	- The latest start time must be at least 30 minutes later than the earliest start time.
	//
	// >	- If you specify the `PlannedStartTime` parameter but do not specify the PlannedEndTime parameter, the latest start time of the task is set to a value that is calculated by `the value of the PlannedEndTime parameter + 30 minutes` by default. For example, if you set the `PlannedStartTime` parameter to `2021-01-14T09:00:00Z` and you do not specify the PlannedEndTime parameter, the latest start time of the task is set to `2021-01-14T09:30:00Z`.
	//
	// example:
	//
	// 2021-01-14T09:30:00Z
	PlannedEndTime *string `json:"PlannedEndTime,omitempty" xml:"PlannedEndTime,omitempty"`
	// The start time to change the primary zone within the scheduled time period. Specify the time in the ISO 8601 standard in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC.
	//
	// > 	- The start time of the task can be a point in time within the next 24 hours. For example, if the current time is `2021-01-14T09:00:00Z`, you can specify a point in time from `2021-01-14T09:00:00Z` to `2021-01-15T09:00:00Z`.
	//
	// >	- If you leave this parameter empty, the primary zone is immediately changed.
	//
	// example:
	//
	// 2021-01-14T09:00:00Z
	PlannedStartTime     *string `json:"PlannedStartTime,omitempty" xml:"PlannedStartTime,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The virtual private cloud (VPC) ID of the destination primary zone.
	//
	// example:
	//
	// vpc-**********
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The ID of the vSwitch in the destination primary zone.
	//
	// > 	- For a PolarDB for PostgreSQL (Compatible with Oracle) cluster or a PolarDB for PostgreSQL cluster, this parameter is required.
	//
	// >	- For a PolarDB for MySQL cluster, the default vSwitch is used if no vSwitches are created in the destination zone. If a vSwitch is in the destination zone, this parameter is required.
	//
	// example:
	//
	// vsw-**************
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the destination primary zone.
	//
	// >  You can call the DescribeRegions operation to query available zones.[](~~98041~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The zone type. Valid values:
	//
	// 	- **Primary**: primary zone
	//
	// 	- **Standby**: secondary zone
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
