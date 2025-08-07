// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterServerlessConfRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowShutDown(v string) *ModifyDBClusterServerlessConfRequest
	GetAllowShutDown() *string
	SetCrontabJobId(v string) *ModifyDBClusterServerlessConfRequest
	GetCrontabJobId() *string
	SetDBClusterId(v string) *ModifyDBClusterServerlessConfRequest
	GetDBClusterId() *string
	SetFromTimeService(v bool) *ModifyDBClusterServerlessConfRequest
	GetFromTimeService() *bool
	SetOwnerAccount(v string) *ModifyDBClusterServerlessConfRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBClusterServerlessConfRequest
	GetOwnerId() *int64
	SetPlannedEndTime(v string) *ModifyDBClusterServerlessConfRequest
	GetPlannedEndTime() *string
	SetPlannedStartTime(v string) *ModifyDBClusterServerlessConfRequest
	GetPlannedStartTime() *string
	SetResourceOwnerAccount(v string) *ModifyDBClusterServerlessConfRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBClusterServerlessConfRequest
	GetResourceOwnerId() *int64
	SetScaleApRoNumMax(v string) *ModifyDBClusterServerlessConfRequest
	GetScaleApRoNumMax() *string
	SetScaleApRoNumMin(v string) *ModifyDBClusterServerlessConfRequest
	GetScaleApRoNumMin() *string
	SetScaleMax(v string) *ModifyDBClusterServerlessConfRequest
	GetScaleMax() *string
	SetScaleMin(v string) *ModifyDBClusterServerlessConfRequest
	GetScaleMin() *string
	SetScaleRoNumMax(v string) *ModifyDBClusterServerlessConfRequest
	GetScaleRoNumMax() *string
	SetScaleRoNumMin(v string) *ModifyDBClusterServerlessConfRequest
	GetScaleRoNumMin() *string
	SetSecondsUntilAutoPause(v string) *ModifyDBClusterServerlessConfRequest
	GetSecondsUntilAutoPause() *string
	SetServerlessRuleCpuEnlargeThreshold(v string) *ModifyDBClusterServerlessConfRequest
	GetServerlessRuleCpuEnlargeThreshold() *string
	SetServerlessRuleCpuShrinkThreshold(v string) *ModifyDBClusterServerlessConfRequest
	GetServerlessRuleCpuShrinkThreshold() *string
	SetServerlessRuleMode(v string) *ModifyDBClusterServerlessConfRequest
	GetServerlessRuleMode() *string
	SetTaskId(v string) *ModifyDBClusterServerlessConfRequest
	GetTaskId() *string
}

type ModifyDBClusterServerlessConfRequest struct {
	// Specifies whether to enable No-activity Suspension. Default value: false. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	AllowShutDown *string `json:"AllowShutDown,omitempty" xml:"AllowShutDown,omitempty"`
	// Cycle policy ID.
	//
	// example:
	//
	// 143f8e9f-2566-4dff-be47-bed79f28fc78
	CrontabJobId *string `json:"CrontabJobId,omitempty" xml:"CrontabJobId,omitempty"`
	// The ID of the serverless cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-bp10gr51qasnl****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies an immediate or scheduled task to modify parameters and restart the cluster. Valid values:
	//
	// 	- false: scheduled task
	//
	// 	- true: immediate task
	//
	// example:
	//
	// false
	FromTimeService *bool   `json:"FromTimeService,omitempty" xml:"FromTimeService,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The latest start time for upgrading the specifications within the scheduled time period. Specify the time in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC.
	//
	// > 	- The value of this parameter must be at least 30 minutes later than the value of PlannedStartTime.
	//
	// >	- If you specify PlannedStartTime but do not specify PlannedEndTime, the latest start time of the task is set to a value that is calculated by using the following formula: `PlannedEndTime value + 30 minutes`. For example, if you set PlannedStartTime to `2021-01-14T09:00:00Z` and you do not specify PlannedEndTime, the latest start time of the task is set to `2021-01-14T09:30:00Z`.
	//
	// example:
	//
	// 2021-01-14T09:30:00Z
	PlannedEndTime *string `json:"PlannedEndTime,omitempty" xml:"PlannedEndTime,omitempty"`
	// The earliest start time of the scheduled task for adding the read-only node. The scheduled task specifies that the task is run in the required period. Specify the time in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC.
	//
	// > 	- The earliest start time of the scheduled task can be a point in time within the next 24 hours. For example, if the current time is `2021-01-14T09:00:00Z`, you can specify a point in time between `2021-01-14T09:00:00Z` and `2021-01-15T09:00:00Z`.
	//
	// >	- If you leave this parameter empty, the task for adding the read-only node is immediately run by default.
	//
	// example:
	//
	// 2021-01-14T09:00:00Z
	PlannedStartTime     *string `json:"PlannedStartTime,omitempty" xml:"PlannedStartTime,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The maximum number of stable AP read-only nodes. Valid values: 0 to 7.
	//
	// example:
	//
	// 1
	ScaleApRoNumMax *string `json:"ScaleApRoNumMax,omitempty" xml:"ScaleApRoNumMax,omitempty"`
	// The minimum number of stable AP read-only nodes. Valid values: 0 to 7.
	//
	// example:
	//
	// 1
	ScaleApRoNumMin *string `json:"ScaleApRoNumMin,omitempty" xml:"ScaleApRoNumMin,omitempty"`
	// The maximum number of PCUs per node for scaling. Valid values: 1 PCU to 32 PCUs.
	//
	// example:
	//
	// 10
	ScaleMax *string `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// The minimum number of PCUs per node for scaling. Valid values: 1 PCU to 31 PCUs.
	//
	// example:
	//
	// 1
	ScaleMin *string `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	// The maximum number of read-only nodes for scaling. Valid values: 0 to 15.
	//
	// example:
	//
	// 2
	ScaleRoNumMax *string `json:"ScaleRoNumMax,omitempty" xml:"ScaleRoNumMax,omitempty"`
	// The minimum number of read-only nodes for scaling. Valid values: 0 to 15.
	//
	// example:
	//
	// 1
	ScaleRoNumMin *string `json:"ScaleRoNumMin,omitempty" xml:"ScaleRoNumMin,omitempty"`
	// The detection period for No-activity Suspension. Valid values: 5 to 1440. Unit: minutes. The detection duration must be a multiple of 5 minutes.
	//
	// example:
	//
	// 10
	SecondsUntilAutoPause *string `json:"SecondsUntilAutoPause,omitempty" xml:"SecondsUntilAutoPause,omitempty"`
	// CPU burst threshold
	//
	// example:
	//
	// 80
	ServerlessRuleCpuEnlargeThreshold *string `json:"ServerlessRuleCpuEnlargeThreshold,omitempty" xml:"ServerlessRuleCpuEnlargeThreshold,omitempty"`
	// CPU downscale threshold
	//
	// example:
	//
	// 50
	ServerlessRuleCpuShrinkThreshold *string `json:"ServerlessRuleCpuShrinkThreshold,omitempty" xml:"ServerlessRuleCpuShrinkThreshold,omitempty"`
	// Elastic sensitivity. Values: - normal: standard - flexible: sensitive
	//
	// example:
	//
	// normal
	ServerlessRuleMode *string `json:"ServerlessRuleMode,omitempty" xml:"ServerlessRuleMode,omitempty"`
	// Asynchronous task ID.
	//
	// example:
	//
	// 143f8e9f-2566-4dff-be47-bed79f28fc78
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyDBClusterServerlessConfRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterServerlessConfRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterServerlessConfRequest) GetAllowShutDown() *string {
	return s.AllowShutDown
}

func (s *ModifyDBClusterServerlessConfRequest) GetCrontabJobId() *string {
	return s.CrontabJobId
}

func (s *ModifyDBClusterServerlessConfRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterServerlessConfRequest) GetFromTimeService() *bool {
	return s.FromTimeService
}

func (s *ModifyDBClusterServerlessConfRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBClusterServerlessConfRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBClusterServerlessConfRequest) GetPlannedEndTime() *string {
	return s.PlannedEndTime
}

func (s *ModifyDBClusterServerlessConfRequest) GetPlannedStartTime() *string {
	return s.PlannedStartTime
}

func (s *ModifyDBClusterServerlessConfRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBClusterServerlessConfRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBClusterServerlessConfRequest) GetScaleApRoNumMax() *string {
	return s.ScaleApRoNumMax
}

func (s *ModifyDBClusterServerlessConfRequest) GetScaleApRoNumMin() *string {
	return s.ScaleApRoNumMin
}

func (s *ModifyDBClusterServerlessConfRequest) GetScaleMax() *string {
	return s.ScaleMax
}

func (s *ModifyDBClusterServerlessConfRequest) GetScaleMin() *string {
	return s.ScaleMin
}

func (s *ModifyDBClusterServerlessConfRequest) GetScaleRoNumMax() *string {
	return s.ScaleRoNumMax
}

func (s *ModifyDBClusterServerlessConfRequest) GetScaleRoNumMin() *string {
	return s.ScaleRoNumMin
}

func (s *ModifyDBClusterServerlessConfRequest) GetSecondsUntilAutoPause() *string {
	return s.SecondsUntilAutoPause
}

func (s *ModifyDBClusterServerlessConfRequest) GetServerlessRuleCpuEnlargeThreshold() *string {
	return s.ServerlessRuleCpuEnlargeThreshold
}

func (s *ModifyDBClusterServerlessConfRequest) GetServerlessRuleCpuShrinkThreshold() *string {
	return s.ServerlessRuleCpuShrinkThreshold
}

func (s *ModifyDBClusterServerlessConfRequest) GetServerlessRuleMode() *string {
	return s.ServerlessRuleMode
}

func (s *ModifyDBClusterServerlessConfRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifyDBClusterServerlessConfRequest) SetAllowShutDown(v string) *ModifyDBClusterServerlessConfRequest {
	s.AllowShutDown = &v
	return s
}

func (s *ModifyDBClusterServerlessConfRequest) SetCrontabJobId(v string) *ModifyDBClusterServerlessConfRequest {
	s.CrontabJobId = &v
	return s
}

func (s *ModifyDBClusterServerlessConfRequest) SetDBClusterId(v string) *ModifyDBClusterServerlessConfRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterServerlessConfRequest) SetFromTimeService(v bool) *ModifyDBClusterServerlessConfRequest {
	s.FromTimeService = &v
	return s
}

func (s *ModifyDBClusterServerlessConfRequest) SetOwnerAccount(v string) *ModifyDBClusterServerlessConfRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterServerlessConfRequest) SetOwnerId(v int64) *ModifyDBClusterServerlessConfRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterServerlessConfRequest) SetPlannedEndTime(v string) *ModifyDBClusterServerlessConfRequest {
	s.PlannedEndTime = &v
	return s
}

func (s *ModifyDBClusterServerlessConfRequest) SetPlannedStartTime(v string) *ModifyDBClusterServerlessConfRequest {
	s.PlannedStartTime = &v
	return s
}

func (s *ModifyDBClusterServerlessConfRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterServerlessConfRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterServerlessConfRequest) SetResourceOwnerId(v int64) *ModifyDBClusterServerlessConfRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterServerlessConfRequest) SetScaleApRoNumMax(v string) *ModifyDBClusterServerlessConfRequest {
	s.ScaleApRoNumMax = &v
	return s
}

func (s *ModifyDBClusterServerlessConfRequest) SetScaleApRoNumMin(v string) *ModifyDBClusterServerlessConfRequest {
	s.ScaleApRoNumMin = &v
	return s
}

func (s *ModifyDBClusterServerlessConfRequest) SetScaleMax(v string) *ModifyDBClusterServerlessConfRequest {
	s.ScaleMax = &v
	return s
}

func (s *ModifyDBClusterServerlessConfRequest) SetScaleMin(v string) *ModifyDBClusterServerlessConfRequest {
	s.ScaleMin = &v
	return s
}

func (s *ModifyDBClusterServerlessConfRequest) SetScaleRoNumMax(v string) *ModifyDBClusterServerlessConfRequest {
	s.ScaleRoNumMax = &v
	return s
}

func (s *ModifyDBClusterServerlessConfRequest) SetScaleRoNumMin(v string) *ModifyDBClusterServerlessConfRequest {
	s.ScaleRoNumMin = &v
	return s
}

func (s *ModifyDBClusterServerlessConfRequest) SetSecondsUntilAutoPause(v string) *ModifyDBClusterServerlessConfRequest {
	s.SecondsUntilAutoPause = &v
	return s
}

func (s *ModifyDBClusterServerlessConfRequest) SetServerlessRuleCpuEnlargeThreshold(v string) *ModifyDBClusterServerlessConfRequest {
	s.ServerlessRuleCpuEnlargeThreshold = &v
	return s
}

func (s *ModifyDBClusterServerlessConfRequest) SetServerlessRuleCpuShrinkThreshold(v string) *ModifyDBClusterServerlessConfRequest {
	s.ServerlessRuleCpuShrinkThreshold = &v
	return s
}

func (s *ModifyDBClusterServerlessConfRequest) SetServerlessRuleMode(v string) *ModifyDBClusterServerlessConfRequest {
	s.ServerlessRuleMode = &v
	return s
}

func (s *ModifyDBClusterServerlessConfRequest) SetTaskId(v string) *ModifyDBClusterServerlessConfRequest {
	s.TaskId = &v
	return s
}

func (s *ModifyDBClusterServerlessConfRequest) Validate() error {
	return dara.Validate(s)
}
